/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) newAuthenticatedGETRequest(ctx context.Context, endpoint string) (req *http.Request, err error) {
	req, err = c.newGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	c.authenticateRequest(ctx, req)

	return
}

func (c *Client) newGETRequest(ctx context.Context, endpoint string) (req *http.Request, err error) {
	req, err = http.NewRequestWithContext(ctx, http.MethodGet, c.opts.RemoteAddr+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("new http GET request: %v", err)
	}

	// Set headers.
	req.Header.Set("Accept", "application/json")

	return
}

func (c *Client) newAuthenticatedPOSTRequest(ctx context.Context, endpoint string, body any) (req *http.Request, err error) {
	// Marshal the request data to JSON.
	reqData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %v", err)
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, c.opts.RemoteAddr+endpoint, bytes.NewReader(reqData))
	if err != nil {
		return nil, fmt.Errorf("new http POST request: %v", err)
	}

	// Set headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c.authenticateRequest(ctx, req)

	return
}

func (c *Client) newAuthenticatedPUTRequest(ctx context.Context, endpoint string, body any) (req *http.Request, err error) {
	// Marshal the request data to JSON.
	reqData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal json: %v", err)
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPut, c.opts.RemoteAddr+endpoint, strings.NewReader(string(reqData)))
	if err != nil {
		return nil, fmt.Errorf("new http POST request: %v", err)
	}

	// Set headers.
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	c.authenticateRequest(ctx, req)

	return
}

func (c *Client) newAuthenticatedDELETERequest(ctx context.Context, endpoint string) (req *http.Request, err error) {
	req, err = http.NewRequestWithContext(ctx, http.MethodDelete, c.opts.RemoteAddr+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("new http DELETE request: %v", err)
	}

	// Set headers.
	c.authenticateRequest(ctx, req)

	return
}

// doSimpleRequest is a helper that executes the given request and attempts to parse
// its JSON response into resp.
// The argument resp must be a pointer.
// If any other status code than 200 is received, an error is returned.
func (c *Client) doSimpleRequest(req *http.Request, resp any) error {
	r, err := c.httpc.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %v", err)
	} else if r.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code '%d' received", r.StatusCode)
	}

	// If no return value is expected, do not parse the response.
	if resp == nil {
		return nil
	}

	// Parse response.
	err = parseResponse(r, &resp)
	if err != nil {
		return fmt.Errorf("failed to parse response: %v", err)
	}

	return nil
}

func parseResponse(resp *http.Response, dst any) error {
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	err = json.Unmarshal(data, dst)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %v; raw response: %s", err, string(data))
	}

	return nil
}
