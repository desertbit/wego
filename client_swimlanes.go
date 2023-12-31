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
	"io"
)

// GetAllSwimlanes performs a get_all_swimlanes request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_swimlanes
func (c *Client) GetAllSwimlanes(ctx context.Context, boardID string) (swimlanes []GetAllSwimlane, err error) {
	endpoint := c.endpoint("boards", boardID, "swimlanes")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &swimlanes)
	if err != nil {
		return
	}

	return
}

// NewSwimlane performs a new_swimlane request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_swimlane
func (c *Client) NewSwimlane(ctx context.Context, boardID, title string) (r NewSwimlaneResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "swimlanes")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, newSwimlaneRequest{Title: title})
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetSwimlane performs a get_swimlane request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_swimlane
//
// Returns ErrNotFound, if the swimlane could not be found.
func (c *Client) GetSwimlane(ctx context.Context, boardID, swimlaneID string) (swimlane GetSwimlane, err error) {
	endpoint := c.endpoint("boards", boardID, "swimlanes", swimlaneID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &swimlane)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = ErrNotFound
		}
		return
	}

	return
}

// DeleteSwimlane performs a delete_swimlane request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_swimlane
func (c *Client) DeleteSwimlane(ctx context.Context, boardID, swimlaneID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "swimlanes", swimlaneID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type GetAllSwimlane struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type newSwimlaneRequest struct {
	Title string `json:"title"`
}

type NewSwimlaneResponse struct {
	ID string `json:"_id"`
}

type GetSwimlane struct {
	Title      string `json:"title"`
	Archived   bool   `json:"archived"`
	ArchivedAt string `json:"archivedAt"`
	BoardID    string `json:"boardId"`
	CreatedAt  string `json:"createdAt"`
	Sort       int    `json:"sort"`
	Color      string `json:"color"`
	UpdatedAt  string `json:"updatedAt"`
	ModifiedAt string `json:"modifiedAt"`
	Type       string `json:"type"`
}
