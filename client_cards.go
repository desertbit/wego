/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import (
	"context"
	"time"
)

// GetCardsByCustomField performs a get_cards_by_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_cards_by_custom_field
func (c *Client) GetCardsByCustomField(ctx context.Context, boardID, customField, customFieldValue string) (cards []GetCardByCustomField, err error) {
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
func (c *Client) GetAllCards(ctx context.Context, boardID, listID string) (cards []GetAllCard, err error) {
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
func (c *Client) GetCard(ctx context.Context, boardID, listID, cardID string) (card GetCard, err error) {
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

// GetSwimlaneCards performs a get_swimlane_cards request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_swimlane_cards
func (c *Client) GetSwimlaneCards(ctx context.Context, boardID, swimlaneID string) (cards []GetSwimlaneCard, err error) {
	var endpoint = c.endpoint("boards", boardID, "swimlanes", swimlaneID, "cards")

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

//#############//
//### Types ###//
//#############//

type GetAllCard struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetCardByCustomField struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ListID      string `json:"listId"`
	SwimlaneID  string `json:"swimlaneId"`
}

type GetSwimlaneCard struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ListID      string `json:"listId"`
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

type GetCard struct {
	Title            string            `json:"title"`
	Archived         bool              `json:"archived"`
	ArchivedAt       string            `json:"archivedAt"`
	ParentID         string            `json:"parentId"`
	ListID           string            `json:"listId"`
	SwimlaneID       string            `json:"swimlaneId"`
	BoardID          string            `json:"boardId"`
	CoverID          string            `json:"coverId"`
	Color            string            `json:"color"`
	CreatedAt        string            `json:"createdAt"`
	ModifiedAt       string            `json:"modifiedAt"`
	CustomFields     []CardCustomField `json:"customFields"`
	DateLastActivity string            `json:"dateLastActivity"`
	Description      string            `json:"description"`
	RequestedBy      string            `json:"requestedBy"`
	AssignedBy       string            `json:"assignedBy"`
	LabelIds         []string          `json:"labelIds"`
	Members          []string          `json:"members"`
	Assignees        []string          `json:"assignees"`
	ReceivedAt       string            `json:"receivedAt"`
	StartAt          string            `json:"startAt"`
	DueAt            string            `json:"dueAt"`
	EndAt            string            `json:"endAt"`
	SpentTime        int               `json:"spentTime"`
	IsOvertime       bool              `json:"isOvertime"`
	UserID           string            `json:"userId"`
	Sort             int               `json:"sort"`
	SubtaskSort      int               `json:"subtaskSort"`
	Type             string            `json:"type"`
	LinkedID         string            `json:"linkedId"`
	Vote             Vote              `json:"vote"`
	Poker            Poker             `json:"poker"`
	TargetIDGantt    []string          `json:"targetId_gantt"`
	LinkTypeGantt    []int             `json:"linkType_gantt"`
	LinkIDGantt      []string          `json:"linkId_gantt"`
	CardNumber       int               `json:"cardNumber"`
}

type CardCustomField struct {
	ID    string `json:"_id"`
	Value any    `json:"value"`
}

type Vote struct {
	Question             string   `json:"question"`
	Positive             []string `json:"positive"`
	Negative             []string `json:"negative"`
	End                  string   `json:"end"`
	Public               bool     `json:"public"`
	AllowNonBoardMembers bool     `json:"allowNonBoardMembers"`
}

type Poker struct {
	Question             bool     `json:"question"`
	One                  []string `json:"one"`
	Two                  []string `json:"two"`
	Three                []string `json:"three"`
	Five                 []string `json:"five"`
	Eight                []string `json:"eight"`
	Thirteen             []string `json:"thirteen"`
	Twenty               []string `json:"twenty"`
	Forty                []string `json:"forty"`
	OneHundred           []string `json:"oneHundred"`
	Unsure               []string `json:"unsure"`
	End                  string   `json:"end"`
	AllowNonBoardMembers bool     `json:"allowNonBoardMembers"`
	Estimation           int      `json:"estimation"`
}

type EditCardOptions struct {
	Title        string    `json:"title"`
	Sort         string    `json:"sort"`
	ParentID     string    `json:"parentId"`
	Description  string    `json:"description"`
	Color        string    `json:"color"`
	Vote         Vote      `json:"vote"`
	Poker        Poker     `json:"poker"`
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
