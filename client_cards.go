/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import (
	"context"
	"encoding/json"
	"time"
)

// GetCardsByCustomField performs a get_cards_by_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_cards_by_custom_field
func (c *Client) GetCardsByCustomField(ctx context.Context, boardID, customField, customFieldValue string) (cards []CardExtended, err error) {
	var endpoint = c.endpoint("boards", boardID, "cardsByCustomField", customField, customFieldValue)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &cards)
	if err != nil {
		return
	}

	return
}

// GetAllCards performs a get_all_cards request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_cards
func (c *Client) GetAllCards(ctx context.Context, boardID, listID string) (cards []Card, err error) {
	var endpoint = c.endpoint("boards", boardID, "lists", listID, "cards")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &cards)
	if err != nil {
		return
	}

	return
}

// NewCard performs a new_card request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_card
func (c *Client) NewCard(ctx context.Context, boardID, listID string, request NewCardRequest) (r NewCardResponse, err error) {
	var endpoint = c.endpoint("boards", boardID, "lists", listID, "cards")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, request)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetCard performs a get_card request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_card
func (c *Client) GetCard(ctx context.Context, boardID, listID, cardID string) (card CardDetail, err error) {
	var endpoint = c.endpoint("boards", boardID, "lists", listID, "cards", cardID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &card)
	if err != nil {
		return
	}

	return
}

// EditCard performs a edit_card request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#edit_card
func (c *Client) EditCard(ctx context.Context, boardID, listID, cardID string, opts EditCardOptions) (r EditCardResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "lists", listID, "cards", cardID)

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, opts)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, nil)
	if err != nil {
		return
	}

	return
}

// DeleteCard performs a delete_card request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_card
func (c *Client) DeleteCard(ctx context.Context, boardID, cardID string) (err error) {
	var endpoint = "/api/boards/" + boardID + "/cards/" + cardID

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type Card struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CardExtended struct {
	Card `json:",inline"`

	ListID     string `json:"listId"`
	SwimlaneID string `json:"swimlaneId"`
}

type NewCardRequest struct {
	// Required
	AuthorID    string `json:"authorId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	SwimlaneID  string `json:"swimlaneId"`

	// Optional
	NewCardOptions `json:",inline"`
}

type NewCardOptions struct {
	MemberIDs []string `json:"members"`
	Assignees []string `json:"assignees"`
}

type NewCardResponse struct {
	ID string `json:"_id"`
}

type CardDetail struct {
	Title            string          `json:"title"`
	Archived         bool            `json:"archived"`
	ArchivedAt       string          `json:"archivedAt"`
	ParentID         string          `json:"parentId"`
	ListID           string          `json:"listId"`
	SwimlaneID       string          `json:"swimlaneId"`
	BoardID          string          `json:"boardId"`
	CoverID          string          `json:"coverId"`
	Color            string          `json:"color"`
	CreatedAt        string          `json:"createdAt"`
	ModifiedAt       string          `json:"modifiedAt"`
	CustomFields     json.RawMessage `json:"customFields"`
	DateLastActivity string          `json:"dateLastActivity"`
	Description      string          `json:"description"`
	RequestedBy      string          `json:"requestedBy"`
	AssignedBy       string          `json:"assignedBy"`
	LabelIDs         []string        `json:"labelIds"`
	Members          []string        `json:"members"`
	Assignees        []string        `json:"assignees"`
	ReceivedAt       string          `json:"receivedAt"`
	StartAt          string          `json:"startAt"`
	DueAt            string          `json:"dueAt"`
	EndAt            string          `json:"endAt"`
	SpentTime        int             `json:"spentTime"`
	IsOvertime       bool            `json:"isOvertime"`
	UserID           string          `json:"userId"`
	Sort             int             `json:"sort"`
	SubtaskSort      int             `json:"subtaskSort"`
	Type             string          `json:"type"`
	LinkedID         string          `json:"linkedId"`
	Vote             CardVote        `json:"vote"`
}

type CardVote struct {
	Question             string   `json:"question"`
	Positive             []string `json:"positive"`
	Negative             []string `json:"negative"`
	End                  string   `json:"end"`
	Public               bool     `json:"public"`
	AllowNonBoardMembers bool     `json:"allowNonBoardMembers"`
}

type EditCardOptions struct {
	Title        string    `json:"title"`
	Sort         string    `json:"sort"`
	ParentID     string    `json:"parentId"`
	Description  string    `json:"description"`
	Color        string    `json:"color"`
	Vote         CardVote  `json:"vote"`
	LabelIDs     []string  `json:"labelIds"`
	RequestedBy  string    `json:"requestedBy"`
	AssignedBy   string    `json:"assignedBy"`
	ReceivedAt   time.Time `json:"receivedAt"`
	StartAt      time.Time `json:"startAt"`
	DueAt        time.Time `json:"dueAt"`
	EndAt        time.Time `json:"endAt"`
	SpentTime    string    `json:"spentTime"`
	IsOverTime   bool      `json:"isOverTime"`
	CustomFields string    `json:"customFields"`
	Members      []string  `json:"members"`
	Assignees    []string  `json:"assignees"`
	SwimlaneID   string    `json:"swimlaneId"`
	ListID       string    `json:"listId"`
	AuthorID     string    `json:"authorId"`
}

type EditCardResponse struct {
	ID string `json:"_id"`
}
