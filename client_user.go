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
)

// GetCurrentUserID returns the id of the logged in user.
// This is an additional convenience method that has no pendant in the Wekan API.
func (c *Client) GetCurrentUserID() (id string) {
	c.mx.Lock()
	id = c.mxUserID
	c.mx.Unlock()
	return
}

// AddBoardMember performs a add_board_member request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#add_board_member
func (c *Client) AddBoardMember(ctx context.Context, boardID, userID string, data AddBoardMemberRequest) (err error) {
	endpoint := c.endpoint("boards", boardID, "members", userID, "add")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, data)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// RemoveBoardMember performs a remove_board_member request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#remove_board_member
func (c *Client) RemoveBoardMember(ctx context.Context, boardID, userID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "members", userID, "remove")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, removeBoardMemberRequest{Action: "remove"})
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// CreateUserToken performs a create_user_token request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#create_user_token
func (c *Client) CreateUserToken(ctx context.Context, userID string) (r CreateUserTokenResponse, err error) {
	endpoint := c.endpoint("createtoken", userID)

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, nil)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetCurrentUser performs a get_current_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_current_user
func (c *Client) GetCurrentUser(ctx context.Context) (u User, err error) {
	endpoint := c.endpoint("user")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &u)
	if err != nil {
		return
	}

	return
}

// GetAllUsers performs a get_all_users request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_users
func (c *Client) GetAllUsers(ctx context.Context) (users []GetAllUser, err error) {
	endpoint := c.endpoint("users")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &users)
	if err != nil {
		return
	}

	return
}

// NewUser performs a new_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_user
func (c *Client) NewUser(ctx context.Context, data NewUserRequest) (r NewUserResponse, err error) {
	endpoint := c.endpoint("users")

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

// GetUser performs a get_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_user
func (c *Client) GetUser(ctx context.Context, userID string) (user User, err error) {
	endpoint := c.endpoint("users", userID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &user)
	if err != nil {
		return
	}

	return
}

// EditUser performs a edit_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#edit_user
//
// Note: Possible values for action:
//   - takeOwnership: The admin takes the ownership of ALL boards of the user (archived and not archived) where the user is admin on.
//   - disableLogin:  Disable a user (the user is not allowed to login and his login tokens are purged)
//   - enableLogin:   Enable a user
func (c *Client) EditUser(ctx context.Context, userID, action string) (err error) {
	endpoint := c.endpoint("users", userID)

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, editUserRequest{Action: action})
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// DeleteUser performs a delete_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_user
func (c *Client) DeleteUser(ctx context.Context, userID string) (err error) {
	endpoint := c.endpoint("users", userID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type AddBoardMemberRequest struct {
	Action        string `json:"action"`
	IsAdmin       bool   `json:"isAdmin"`
	IsNoComments  bool   `json:"isNoComments"`
	IsCommentOnly bool   `json:"isCommentOnly"`
}

type removeBoardMemberRequest struct {
	Action string `json:"action"`
}

type CreateUserTokenResponse struct {
	ID string `json:"_id"`
}

type User struct {
	Username             string          `json:"username"`
	Emails               []UserEmail     `json:"emails"`
	CreatedAt            string          `json:"createdAt"`
	ModifiedAt           string          `json:"modifiedAt"`
	Profile              UserProfile     `json:"profile"`
	Services             json.RawMessage `json:"services"`
	Heartbeat            string          `json:"heartbeat"`
	IsAdmin              bool            `json:"isAdmin"`
	CreatedThroughApi    bool            `json:"createdThroughApi"`
	LoginDisabled        bool            `json:"loginDisabled"`
	AuthenticationMethod string          `json:"authenticationMethod"`
	SessionData          UserSessionData `json:"sessionData"`
	ImportUsernames      []string        `json:"importUsernames"`
}

type UserEmail struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
}

type UserSessionData struct {
	TotalHits int `json:"totalHits"`
	LastHit   int `json:"lastHit"`
}

type UserProfile struct {
	AvatarUrl                string          `json:"avatarUrl"`
	EmailBuffer              []string        `json:"emailBuffer"`
	Fullname                 string          `json:"fullname"`
	ShowDesktopDragHandles   bool            `json:"showDesktopDragHandles"`
	HideCheckedItems         bool            `json:"hideCheckedItems"`
	HiddenSystemMessages     bool            `json:"hiddenSystemMessages"`
	HiddenMinicardLabelText  bool            `json:"hiddenMinicardLabelText"`
	Initials                 string          `json:"initials"`
	InvitedBoards            []string        `json:"invitedBoards"`
	Language                 string          `json:"language"`
	Notifications            json.RawMessage `json:"notifications"`
	Activity                 string          `json:"activity"`
	Read                     string          `json:"read"`
	ShowCardsCountAt         int             `json:"showCardsCountAt"`
	StartDayOfWeek           int             `json:"startDayOfWeek"`
	StarredBoards            []string        `json:"starredBoards"`
	Icode                    string          `json:"icode"`
	BoardView                string          `json:"boardView"`
	ListSortBy               string          `json:"listSortBy"`
	TemplatesBoardID         string          `json:"templatesBoardId"`
	CardTemplatesSwimlaneID  string          `json:"cardTemplatesSwimlaneId"`
	ListTemplatesSwimlaneID  string          `json:"listTemplatesSwimlaneId"`
	BoardTemplatesSwimlaneID string          `json:"boardTemplatesSwimlaneId"`
}

type GetAllUser struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
}

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type NewUserResponse struct {
	ID string `json:"_id"`
}

type editUserRequest struct {
	Action string `json:"action"`
}
