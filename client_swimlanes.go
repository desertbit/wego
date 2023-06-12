/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import "context"

// GetAllSwimlanes performs a get_all_swimlanes request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_swimlanes
func (c *Client) GetAllSwimlanes(ctx context.Context, boardID string) (swimlanes []Swimlane, err error) {
	var endpoint = "/api/boards/" + boardID + "/swimlanes"

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
	var endpoint = "/api/boards/" + boardID + "/swimlanes"

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
func (c *Client) GetSwimlane(ctx context.Context, boardID, swimlaneID string) (swimlane SwimlaneDetail, err error) {
	var endpoint = "/api/boards/" + boardID + "/swimlanes/" + swimlaneID

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &swimlane)
	if err != nil {
		return
	}

	return
}

// DeleteSwimlane performs a delete_swimlane request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_swimlane
func (c *Client) DeleteSwimlane(ctx context.Context, boardID, swimlaneID string) (err error) {
	var endpoint = "/api/boards/" + boardID + "/swimlanes/" + swimlaneID

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type Swimlane struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type newSwimlaneRequest struct {
	Title string `json:"title"`
}

type NewSwimlaneResponse struct {
	ID string `json:"_id"`
}

type SwimlaneDetail struct {
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
