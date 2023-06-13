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

// GetAllIntegrations performs a get_all_integrations request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_integrations
func (c *Client) GetAllIntegrations(ctx context.Context, boardID string) (integrations []Integration, err error) {
	endpoint := c.endpoint("boards", boardID, "integrations")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &integrations)
	if err != nil {
		return
	}

	return
}

// NewIntegration performs a new_integration request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_integration
func (c *Client) NewIntegration(ctx context.Context, boardID, url string) (r NewIntegrationResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "integrations")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, newIntegrationRequest{Url: url})
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// GetIntegration performs a get_integration request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_integration
//
// Returns ErrNotFound, if the integration could not be found.
func (c *Client) GetIntegration(ctx context.Context, boardID, integrationID string) (integration Integration, err error) {
	endpoint := c.endpoint("boards", boardID, "integrations", integrationID)

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &integration)
	if err != nil {
		if errors.Is(err, io.EOF) {
			err = ErrNotFound
		}
		return
	}

	return
}

// EditIntegration performs a edit_integration request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#edit_integration
func (c *Client) EditIntegration(ctx context.Context, boardID, integrationID string, data EditIntegrationOptions) (err error) {
	endpoint := c.endpoint("boards", boardID, "integrations", integrationID)

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, data)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// DeleteIntegration performs a delete_integration request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_integration
func (c *Client) DeleteIntegration(ctx context.Context, boardID, integrationID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "integrations", integrationID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// DeleteIntegrationActivities performs a delete_integration_activities request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_integration_activities
func (c *Client) DeleteIntegrationActivities(ctx context.Context, boardID, integrationID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "integrations", integrationID, "activities")

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// NewIntegrationActivities performs a new_integration_activities request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_integration_activities
func (c *Client) NewIntegrationActivities(ctx context.Context, boardID, integrationID string, activities []string) (integration Integration, err error) {
	endpoint := c.endpoint("boards", boardID, "integrations", integrationID, "activities")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, newIntegrationActivitiesRequest{Activities: activities})
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &integration)
	if err != nil {
		return
	}

	return
}

//#############//
//### Types ###//
//#############//

type Integration struct {
	Enabled    bool     `json:"enabled"`
	Title      string   `json:"title"`
	Type       string   `json:"type"`
	Activities []string `json:"activities"`
	Url        string   `json:"url"`
	Token      string   `json:"token"`
	BoardID    string   `json:"boardId"`
	CreatedAt  string   `json:"createdAt"`
	ModifiedAt string   `json:"modifiedAt"`
	UserID     string   `json:"userId"`
}

type newIntegrationRequest struct {
	Url string `json:"url"`
}

type NewIntegrationResponse struct {
	ID string `json:"_id"`
}

type EditIntegrationOptions struct {
	Enabled    bool     `json:"enabled"`
	Title      string   `json:"title"`
	Url        string   `json:"url"`
	Token      string   `json:"token"`
	Activities []string `json:"activities"`
}

type newIntegrationActivitiesRequest struct {
	Activities []string `json:"activities"`
}
