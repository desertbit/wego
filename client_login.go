/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Login performs a login request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#wekan-rest-api-login
//
// Note: The client ensures to authenticate against the API on its own.
// It is not required to call this method for normal usage.
func (c *Client) Login(ctx context.Context, username, password string) (r LoginResponse, err error) {
	const endpoint = "/users/login"

	// Create the url encoded params.
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)

	return c.loginOrRegister(ctx, endpoint, params)
}

// Register performs a register request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#register
func (c *Client) Register(ctx context.Context, username, password, email string) (r LoginResponse, err error) {
	const endpoint = "/users/register"

	// Create the url encoded params.
	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)
	params.Set("email", email)

	return c.loginOrRegister(ctx, endpoint, params)
}

//################//
//### Internal ###//
//################//

// loginOrRegister is an internal helper that performs a login or register request, since they
// are almost the same in the Wekan API.
func (c *Client) loginOrRegister(ctx context.Context, endpoint string, params url.Values) (r LoginResponse, err error) {
	// Create the HTTP request.
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.opts.RemoteAddr+endpoint, strings.NewReader(params.Encode()))
	if err != nil {
		err = fmt.Errorf("failed to create new request: %v", err)
		return
	}
	req.Header.Set("Content-Type", mimeURL)
	req.Header.Set("Accept", mimeJSON)
	resp, err := c.httpc.Do(req)
	if err != nil {
		err = fmt.Errorf("failed to send POST request: %v", err)
		return
	} else if resp.StatusCode == http.StatusBadRequest {
		var respData badRequestResponse
		err = parseResponse(resp, &respData)
		if err != nil {
			err = fmt.Errorf("failed to parse response of bad request: %v", err)
			return
		}

		err = fmt.Errorf("bad request: %s (%d)", respData.Reason, respData.Error)
		return
	} else if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unexpected status code '%d' received", resp.StatusCode)
		return
	}

	// Parse the response.
	var respData loginResponse
	err = parseResponse(resp, &respData)
	if err != nil {
		err = fmt.Errorf("failed to parse response: %v", err)
		return
	}

	// Load the response into our public type.
	r.load(respData)
	return
}

//#############//
//### Types ###//
//#############//

type badRequestResponse struct {
	Error  int    `json:"error"`
	Reason string `json:"reason"`
}

type loginResponse struct {
	ID           string `json:"id"`
	Token        string `json:"token"`
	TokenExpires string `json:"tokenExpires"`
}

type LoginResponse struct {
	ID           string
	Token        string
	TokenExpires time.Time
}

func (r *LoginResponse) load(l loginResponse) (err error) {
	r.ID = l.ID
	r.Token = l.Token

	r.TokenExpires, err = time.Parse(time.RFC3339, l.TokenExpires)
	if err != nil {
		return fmt.Errorf("failed to parse token expires time stamp: %v", err)
	}

	return nil
}
