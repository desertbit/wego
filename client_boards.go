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
	"errors"
	"io"
	"time"
)

// GetPublicBoards performs a get_public_boards request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_public_boards
func (c *Client) GetPublicBoards(ctx context.Context) (boards []GetPublicBoard, err error) {
	endpoint := c.endpoint("boards")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &boards)
	if err != nil {
		return
	}

	return
}

// NewBoard performs a new_board request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_board
//
// Note: Owner must be a userID, not an email or username.
func (c *Client) NewBoard(ctx context.Context, request NewBoardRequest) (r NewBoardResponse, err error) {
	endpoint := c.endpoint("boards")

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

// GetBoard performs a get_board request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_board
//
// Returns ErrNotFound, if the board could not be found.
func (c *Client) GetBoard(ctx context.Context, boardID string) (r GetBoard, err error) {
	endpoint := c.endpoint("boards", boardID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = ErrNotFound
		}
		return
	}

	return
}

// DeleteBoard performs a delete_board request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_board
func (c *Client) DeleteBoard(ctx context.Context, boardID string) (err error) {
	endpoint := c.endpoint("boards", boardID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// GetBoardAttachments performs a get_board_attachments request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_board_attachments
func (c *Client) GetBoardAttachments(ctx context.Context, boardID string) (attachments []BoardAttachment, err error) {
	endpoint := c.endpoint("boards", boardID, "attachments")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &attachments)
	if err != nil {
		return
	}

	return
}

// ExportJSON performs an export_json request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#exportjson
func (c *Client) ExportJSON(ctx context.Context, boardID string) (boardJSON json.RawMessage, err error) {
	token, err := c.token(ctx)
	if err != nil {
		return
	}

	endpoint := c.endpoint("boards", boardID, "export?authToken="+token)

	req, err := c.newGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &boardJSON)
	if err != nil {
		return
	}

	return
}

// AddBoardLabel performs an add_board_label request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#add_board_label
//
// Note: Currently broken
func (c *Client) AddBoardLabel(ctx context.Context, boardID, name, color string) (err error) {
	endpoint := c.endpoint("boards", boardID, "labels")

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, addBoardLabelRequest{
		Label: addBoardLabelRequestLabel{
			Name:  name,
			Color: color,
		},
	})
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// SetBoardMemberPermission performs an set_board_member_permission request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#set_board_member_permission
func (c *Client) SetBoardMemberPermission(ctx context.Context, boardID, memberID string, opts SetBoardMemberPermissionOptions) (err error) {
	endpoint := c.endpoint("boards", boardID, "members", memberID)

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, opts)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// GetBoardsCount performs a get_boards_count request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_boards_count
func (c *Client) GetBoardsCount(ctx context.Context) (r GetBoardsCountResponse, err error) {
	endpoint := c.endpoint("boards_count")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetBoardsFromUser performs a get_boards_from_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_boards_from_user
func (c *Client) GetBoardsFromUser(ctx context.Context, userID string) (r []GetBoardFromUser, err error) {
	endpoint := c.endpoint("users", userID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

//#############//
//### Types ###//
//#############//

type GetPublicBoard struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type GetBoardFromUser struct {
	ID    string `json:"_id"`
	Title string `json:"title"`
}

type GetBoard struct {
	Title                      string        `json:"title"`
	Slug                       string        `json:"slug"`
	Archived                   bool          `json:"archived"`
	ArchivedAt                 time.Time     `json:"archivedAt"`
	CreatedAt                  time.Time     `json:"createdAt"`
	ModifiedAt                 time.Time     `json:"modifiedAt"`
	Stars                      int           `json:"stars"`
	Labels                     []BoardLabel  `json:"labels"`
	Members                    []BoardMember `json:"members"`
	Permission                 string        `json:"permission"`
	Color                      string        `json:"color"`
	Description                string        `json:"description"`
	SubtasksDefaultBoardID     string        `json:"subtasksDefaultBoardId"`
	SubtasksDefaultListID      string        `json:"subtasksDefaultListId"`
	DateSettingsDefaultBoardID string        `json:"dateSettingsDefaultBoardId"`
	DateSettingsDefaultListID  string        `json:"dateSettingsDefaultListId"`
	AllowsSubtasks             bool          `json:"allowsSubtasks"`
	AllowsAttachments          bool          `json:"allowsAttachments"`
	AllowsChecklists           bool          `json:"allowsChecklists"`
	AllowsComments             bool          `json:"allowsComments"`
	AllowsDescriptionTitle     bool          `json:"allowsDescriptionTitle"`
	AllowsDescriptionText      bool          `json:"allowsDescriptionText"`
	AllowsActivities           bool          `json:"allowsActivities"`
	AllowsLabels               bool          `json:"allowsLabels"`
	AllowsAssignee             bool          `json:"allowsAssignee"`
	AllowsMembers              bool          `json:"allowsMembers"`
	AllowsRequestedBy          bool          `json:"allowsRequestedBy"`
	AllowsAssignedBy           bool          `json:"allowsAssignedBy"`
	AllowsReceivedDate         bool          `json:"allowsReceivedDate"`
	AllowsStartDate            bool          `json:"allowsStartDate"`
	AllowsEndDate              bool          `json:"allowsEndDate"`
	AllowsDueDate              bool          `json:"allowsDueDate"`
	PresentParentTask          string        `json:"presentParentTask"`
	StartAt                    string        `json:"startAt"`
	DueAt                      string        `json:"dueAt"`
	EndAt                      string        `json:"endAt"`
	SpentTime                  int           `json:"spentTime"`
	IsOvertime                 bool          `json:"isOvertime"`
	Type                       string        `json:"type"`
	Sort                       int           `json:"sort"`
}

type BoardLabel struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type BoardMember struct {
	UserID        string `json:"userId"`
	IsAdmin       bool   `json:"isAdmin"`
	IsActive      bool   `json:"isActive"`
	IsNoComments  bool   `json:"isNoComments"`
	IsCommentOnly bool   `json:"isCommentOnly"`
	IsWorker      bool   `json:"isWorker"`
}

type BoardAttachment struct {
	AttachmentID   string `json:"attachmentId"`
	AttachmentName string `json:"attachmentName"`
	AttachmentType string `json:"attachmentType"`
	CardID         string `json:"cardId"`
	ListID         string `json:"listId"`
	SwimlaneID     string `json:"swimlaneId"`
}

type NewBoardRequest struct {
	// Required
	Title string `json:"title"`
	Owner string `json:"owner"`

	// Optional
	NewBoardOptions `json:",inline"`
}

type NewBoardOptions struct {
	IsAdmin       bool   `json:"isAdmin"`
	IsActive      bool   `json:"isActive"`
	IsNoComments  bool   `json:"isNoComments"`
	IsCommentOnly bool   `json:"isCommentOnly"`
	IsWorker      bool   `json:"isWorker"`
	Permission    string `json:"permission"`
	Color         string `json:"color"`
}

type NewBoardResponse struct {
	ID                string `json:"_id"`
	DefaultSwimlaneID string `json:"defaultSwimlaneId"`
}

type addBoardLabelRequest struct {
	Label addBoardLabelRequestLabel `json:"label"`
}

type addBoardLabelRequestLabel struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type SetBoardMemberPermissionOptions struct {
	IsAdmin       bool `json:"isAdmin"`
	IsNoComments  bool `json:"isNoComments"`
	IsCommentOnly bool `json:"isCommentOnly"`
	IsWorker      bool `json:"isWorker"`
}

type GetBoardsCountResponse struct {
	Private int `json:"private"`
	Public  int `json:"public"`
}
