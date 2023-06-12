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

// GetCurrentUser performs a get_current_user request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_current_user
func (c *Client) GetCurrentUser(ctx context.Context) (u UserDetail, err error) {
	const endpoint = "/api/user"

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

//#############//
//### Types ###//
//#############//

type UserDetail struct {
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
