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

// GetAllComments performs a get_all_comments request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_comments
func (c *Client) GetAllComments(ctx context.Context, boardID, cardID string) (comments []GetAllComment, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "comments")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &comments)
	if err != nil {
		return
	}

	return
}

// NewComment performs a new_comment request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_comment
func (c *Client) NewComment(ctx context.Context, boardID, cardID string, data NewCommentRequest) (r NewCommentResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "comments")

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

// GetComment performs a get_comment request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_comment
//
// Returns ErrNotFound, if the comment could not be found.
func (c *Client) GetComment(ctx context.Context, boardID, cardID, commentID string) (comment GetComment, err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "comments", commentID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &comment)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = ErrNotFound
		}
		return
	}

	return
}

// DeleteComment performs a delete_comment request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_comment
func (c *Client) DeleteComment(ctx context.Context, boardID, cardID, commentID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "cards", cardID, "comments", commentID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type GetAllComment struct {
	ID       string `json:"_id"`
	Comment  string `json:"comment"`
	AuthorID string `json:"authorId"`
}

type NewCommentRequest struct {
	AuthorID string `json:"authorId"`
	Comment  string `json:"comment"`
}

type NewCommentResponse struct {
	ID string `json:"_id"`
}

type GetComment struct {
	BoardID    string `json:"boardId"`
	CardID     string `json:"cardId"`
	Text       string `json:"text"`
	CreatedAt  string `json:"createdAt"`
	ModifiedAt string `json:"modifiedAt"`
	UserID     string `json:"userId"`
}
