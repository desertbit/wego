/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import "context"

// GetAllChecklists performs a get_all_checklists request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_checklists
func (c *Client) GetAllChecklists(ctx context.Context, boardID, cardID string) (checklists []GetAllChecklist, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &checklists)
	if err != nil {
		return
	}

	return
}

// NewChecklist performs a new_checklist request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_checklist
func (c *Client) NewChecklist(ctx context.Context, boardID, cardID string, data NewChecklistRequest) (r NewChecklistResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, data)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetChecklist performs a get_checklist request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_checklist
func (c *Client) GetChecklist(ctx context.Context, boardID, cardID, checklistID string) (checklist GetChecklist, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists", checklistID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &checklist)
	if err != nil {
		return
	}

	return
}

// DeleteChecklist performs a delete_checklist request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_checklist
func (c *Client) DeleteChecklist(ctx context.Context, boardID, cardID, checklistID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists", checklistID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type GetAllChecklist struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type NewChecklistRequest struct {
	Title string   `json:"title"`
	Items []string `json:"items"`
}

type NewChecklistResponse struct {
	ID string `json:"_id"`
}

type GetChecklist struct {
	CardId     string          `json:"cardId"`
	Title      string          `json:"title"`
	FinishedAt string          `json:"finishedAt"`
	CreatedAt  string          `json:"createdAt"`
	Sort       int             `json:"sort"`
	Items      []ChecklistItem `json:"items"`
}

type ChecklistItem struct {
	ID         string `json:"_id"`
	Title      string `json:"title"`
	IsFinished bool   `json:"isFinished"`
}
