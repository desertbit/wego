/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import "context"

// GetChecklistItem performs a get_checklist_item request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_checklist_item
func (c *Client) GetChecklistItem(ctx context.Context, boardID, cardID, checklistID, itemID string) (item GetChecklistItem, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists", checklistID, "items", itemID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &item)
	if err != nil {
		return
	}

	return
}

// EditChecklistItem performs a edit_checklist_item request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#edit_checklist_item
func (c *Client) EditChecklistItem(ctx context.Context, boardID, cardID, checklistID, itemID string, data EditChecklistItemRequest) (err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists", checklistID, "items", itemID)

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, data)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// DeleteChecklistItem performs a delete_checklist_item request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_checklist_item
func (c *Client) DeleteChecklistItem(ctx context.Context, boardID, cardID, checklistID, itemID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "checklists", checklistID, "items", itemID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type GetChecklistItem struct {
	Title       string `json:"title"`
	Sort        int    `json:"sort"`
	IsFinished  bool   `json:"isFinished"`
	ChecklistID string `json:"checklistId"`
	CardID      string `json:"cardId"`
	CreatedAt   string `json:"createdAt"`
	ModifiedAt  string `json:"modifiedAt"`
}

type EditChecklistItemRequest struct {
	Title      string `json:"title"`
	IsFinished bool   `json:"isFinished"`
}
