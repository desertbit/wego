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

// GetAllLists performs a get_all_lists request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_lists
func (c *Client) GetAllLists(ctx context.Context, boardID string) (lists []GetAllList, err error) {
	endpoint := c.endpoint("boards", boardID, "lists")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &lists)
	if err != nil {
		return
	}

	return
}

// NewList performs a new_list request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_list
func (c *Client) NewList(ctx context.Context, boardID, title string) (r NewListResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "lists")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, newListRequest{Title: title})
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetList performs a get_list request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_list
//
// Returns ErrNotFound, if the list could not be found.
func (c *Client) GetList(ctx context.Context, boardID, listID string) (list GetList, err error) {
	endpoint := c.endpoint("boards", boardID, "lists", listID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &list)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = ErrNotFound
		}
		return
	}

	return
}

// DeleteList performs a delete_list request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_list
func (c *Client) DeleteList(ctx context.Context, boardID, listID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "lists", listID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type GetAllList struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type newListRequest struct {
	Title string `json:"title"`
}

type NewListResponse struct {
	ID string `json:"_id"`
}

type GetList struct {
	Title      string       `json:"title"`
	Starred    bool         `json:"starred"`
	Archived   bool         `json:"archived"`
	ArchivedAt string       `json:"archivedAt"`
	BoardID    string       `json:"boardId"`
	SwimlaneID string       `json:"swimlaneId"`
	CreatedAt  string       `json:"createdAt"`
	Sort       int          `json:"sort"`
	UpdatedAt  string       `json:"updatedAt"`
	ModifiedAt string       `json:"modifiedAt"`
	WipLimit   ListWIPLimit `json:"wipLimit"`
	Color      string       `json:"color"`
	Type       string       `json:"type"`
}

type ListWIPLimit struct {
	Value   int  `json:"value"`
	Enabled bool `json:"enabled"`
	Soft    bool `json:"soft"`
}
