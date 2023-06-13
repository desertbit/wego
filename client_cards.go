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
//
// Returns ErrNotFound, if the card could not be found.
func (c *Client) GetCard(ctx context.Context, boardID, listID, cardID string) (card GetCard, err error) {
	var endpoint = c.endpoint("boards", boardID, "lists", listID, "cards", cardID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &card)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = ErrNotFound
		}
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

	err = c.doSimpleRequest(req, &r)
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
	MemberIDs []string `json:"members,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
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
	Question             string   `json:"question,omitempty"`
	Positive             []string `json:"positive,omitempty"`
	Negative             []string `json:"negative,omitempty"`
	End                  string   `json:"end,omitempty"`
	Public               bool     `json:"public,omitempty"`
	AllowNonBoardMembers bool     `json:"allowNonBoardMembers,omitempty"`
}

type Poker struct {
	Question             bool     `json:"question,omitempty"`
	One                  []string `json:"one,omitempty"`
	Two                  []string `json:"two,omitempty"`
	Three                []string `json:"three,omitempty"`
	Five                 []string `json:"five,omitempty"`
	Eight                []string `json:"eight,omitempty"`
	Thirteen             []string `json:"thirteen,omitempty"`
	Twenty               []string `json:"twenty,omitempty"`
	Forty                []string `json:"forty,omitempty"`
	OneHundred           []string `json:"oneHundred,omitempty"`
	Unsure               []string `json:"unsure,omitempty"`
	End                  string   `json:"end,omitempty"`
	AllowNonBoardMembers bool     `json:"allowNonBoardMembers,omitempty"`
	Estimation           int      `json:"estimation,omitempty"`
}

type EditCardOptions struct {
	Title        string            `json:"title,omitempty"`
	Sort         string            `json:"sort,omitempty"`
	ParentID     string            `json:"parentId,omitempty"`
	Description  string            `json:"description,omitempty"`
	Color        string            `json:"color,omitempty"`
	Vote         *Vote             `json:"vote,omitempty"`
	Poker        *Poker            `json:"poker,omitempty"`
	LabelIDs     []string          `json:"labelIds,omitempty"`
	RequestedBy  string            `json:"requestedBy,omitempty"`
	AssignedBy   string            `json:"assignedBy,omitempty"`
	ReceivedAt   *time.Time        `json:"receivedAt,omitempty"`
	StartAt      *time.Time        `json:"startAt,omitempty"`
	DueAt        *time.Time        `json:"dueAt,omitempty"`
	EndAt        *time.Time        `json:"endAt,omitempty"`
	SpentTime    string            `json:"spentTime,omitempty"`
	IsOverTime   bool              `json:"isOverTime,omitempty"`
	CustomFields []CardCustomField `json:"customFields,omitempty"`
	Members      []string          `json:"members,omitempty"`
	Assignees    []string          `json:"assignees,omitempty"`
	SwimlaneID   string            `json:"swimlaneId,omitempty"`
	ListID       string            `json:"listId,omitempty"`
	AuthorID     string            `json:"authorId,omitempty"`
}

type EditCardResponse struct {
	ID string `json:"_id"`
}
