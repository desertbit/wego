/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import (
	"context"
	"errors"
	"net/http"
	"path/filepath"
	"sync"
	"time"

	"github.com/desertbit/closer/v3"
	"github.com/rs/zerolog/log"
)

const (
	mimeJSON = "application/json"
	mimeURL  = "application/x-www-form-urlencoded"
)

type Options struct {
	// Mandataroy fields.

	// The address of the wekan server the client should connect to.
	RemoteAddr string
	// The username of the user that should be used to log in.
	Username string
	// The password of the user that should be used to log in.
	Password string

	// Optional fields.

	// The HTTP client that should be used.
	// If nil, a default client is used.
	Client *http.Client

	// The time the client waits between login attempts.
	// Can not be shorter than 1 second.
	TimeBetweenLoginAttemps time.Duration

	// The closer used to manage all routines of the client.
	// If nil, a default closer is created.
	Closer closer.Closer
}

type Client struct {
	closer.Closer

	opts Options

	httpc *http.Client

	// Unbuffered channel that used to distribute API tokens to the request methods.
	authChan chan chan string

	mx       sync.Mutex
	mxUserID string
}

func NewClient(opts Options) (*Client, error) {
	c := &Client{
		Closer:   opts.Closer,
		opts:     opts,
		httpc:    opts.Client,
		authChan: make(chan chan string),
	}

	// Assign default values.
	if opts.Client == nil {
		c.httpc = &http.Client{
			Timeout: 30 * time.Second,
		}
	}
	if opts.TimeBetweenLoginAttemps < time.Second {
		c.opts.TimeBetweenLoginAttemps = time.Second
	}
	if opts.Closer == nil {
		c.Closer = closer.New()
	}

	// Start routines.
	ctx, cancel := c.Context()
	defer cancel()

	// Request the first token.
	// Error can only be a context.ErrCanceled.
	token, tokenExpires, err := c.loginUntilSuccess(ctx)
	if err != nil {
		return nil, err
	}

	c.startConnectionRoutine(token, tokenExpires)

	return c, nil
}

func (c *Client) startConnectionRoutine(token string, tokenExpires time.Time) {
	c.CloserAddWait(1)
	go c.connectionRoutine(token, tokenExpires)
}

func (c *Client) connectionRoutine(token string, tokenExpires time.Time) {
	defer c.CloseAndDone_()

	ctx, cancel := c.Context()
	defer cancel()

	var (
		err error

		closingChan = c.ClosingChan()
	)

	// Start a timer so we renew our token.
	expires := time.NewTimer(time.Until(tokenExpires) - 5*time.Second)
	defer expires.Stop()

	for {
		select {
		case <-closingChan:
			return

		case <-expires.C:
			// Token is expired, login to retrieve a new one.
			token, tokenExpires, err = c.loginUntilSuccess(ctx)
			if err != nil {
				if !errors.Is(err, context.Canceled) {
					log.Error().Err(err).Msg("connectionRoutine")
				}
				return
			}

			// Restart the timer to renew our token.
			expires.Reset(time.Until(tokenExpires) - 5*time.Second)

		case tokenChan := <-c.authChan:
			// Buffered channel, no select needed.
			tokenChan <- token
		}
	}
}

// loginUntilSuccess attempts to login over and over again until successful.
// If a login succeeds, the userID is saved in c and the auth token gets returned.
// The login process is aborted, when the provided context closes.
func (c *Client) loginUntilSuccess(ctx context.Context) (token string, tokenExpires time.Time, err error) {
	var resp LoginResponse
	for {
		resp, err = c.Login(ctx, c.opts.Username, c.opts.Password)
		if err != nil {
			if ctx.Err() != nil {
				err = ctx.Err()
				return
			}

			log.Error().Err(err).Msg("connectionRoutine: login")
			time.Sleep(c.opts.TimeBetweenLoginAttemps)
			continue
		}

		// Successfully logged in.
		token = resp.Token
		tokenExpires = resp.TokenExpires

		// Save the user's id.
		c.mx.Lock()
		c.mxUserID = resp.ID
		c.mx.Unlock()
		return
	}
}

func (c *Client) authenticateRequest(ctx context.Context, req *http.Request) error {
	token, err := c.token(ctx)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	return nil
}

func (c *Client) token(ctx context.Context) (string, error) {
	// Buffered so the connection routine can immediately resume its work.
	tokenChan := make(chan string, 1)

	select {
	case <-c.ClosingChan():
		return "", closer.ErrClosed
	case <-ctx.Done():
		return "", ctx.Err()
	case c.authChan <- tokenChan:
	}

	select {
	case <-c.ClosingChan():
		return "", closer.ErrClosed
	case <-ctx.Done():
		return "", ctx.Err()
	case token := <-tokenChan:
		return token, nil
	}
}

func (c *Client) endpoint(segments ...string) string {
	return "/api/" + filepath.Join(segments...)
}
