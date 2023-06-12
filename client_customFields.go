/**
 * Copyright (c) 2023 Sebastian Borchers
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package wego

import "context"

// GetAllCustomFields performs a get_all_custom_fields request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_custom_fields
func (c *Client) GetAllCustomFields(ctx context.Context, boardID string) (fields []GetAllCustomField, err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &fields)
	if err != nil {
		return
	}

	return
}

// NewCustomField performs a new_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#new_custom_field
func (c *Client) NewCustomField(ctx context.Context, boardID string, data NewCustomFieldRequest) (r NewCustomFieldResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields")

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

// GetCustomField performs a get_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#get_all_custom_fields
func (c *Client) GetCustomField(ctx context.Context, boardID string) (resp []GetCustomField, err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields")

	req, err := c.newAuthenticatedGETRequest(ctx, endpoint)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &resp)
	if err != nil {
		return
	}

	return
}

// EditCustomField performs a edit_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#edit_custom_field
func (c *Client) EditCustomField(ctx context.Context, boardID string, data EditCustomFieldRequest) (r EditCustomFieldResponse, err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields")

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, data)
	if err != nil {
		return
	}

	err = c.doSimpleRequest(req, &r)
	if err != nil {
		return
	}

	return
}

// DeleteCustomField performs a delete_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_custom_field
func (c *Client) DeleteCustomField(ctx context.Context, boardID, customFieldID string) (err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields", customFieldID)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// AddCustomFieldDropdownItems performs a add_custom_field_dropdown_items request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#add_custom_field_dropdown_items
func (c *Client) AddCustomFieldDropdownItems(ctx context.Context, boardID, customFieldID string, items []string) (err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields", customFieldID, "dropdown-items")

	req, err := c.newAuthenticatedPOSTRequest(ctx, endpoint, addCustomFieldDropdownItemsRequest{
		Items: items,
	})
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// EditCustomFieldDropdownItems performs a edit_custom_field_dropdown_items request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#edit_custom_field_dropdown_items
func (c *Client) EditCustomFieldDropdownItems(ctx context.Context, boardID, customFieldID, dropdownItem, name string) (err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields", customFieldID, "dropdown-items", dropdownItem)

	req, err := c.newAuthenticatedPUTRequest(ctx, endpoint, editCustomFieldDropdownItemsRequest{
		Name: name,
	})
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

// DeleteCustomFieldDropdownItem performs a delete_custom_field request against the Wekan server.
// See https://wekan.github.io/api/v5.13/#delete_custom_field_dropdown_item
func (c *Client) DeleteCustomFieldDropdownItem(ctx context.Context, boardID, customFieldID, dropdownItem string) (err error) {
	endpoint := c.endpoint("boards", boardID, "custom-fields", customFieldID, "dropdown-items", dropdownItem)

	req, err := c.newAuthenticatedDELETERequest(ctx, endpoint)
	if err != nil {
		return
	}

	return c.doSimpleRequest(req, nil)
}

//#############//
//### Types ###//
//#############//

type GetAllCustomField struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type NewCustomFieldRequest struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	Settings            string `json:"settings"`
	ShowOnCard          bool   `json:"showOnCard"`
	AutomaticallyOnCard bool   `json:"automaticallyOnCard"`
	ShowLabelOnMiniCard bool   `json:"showLabelOnMiniCard"`
	AuthorId            string `json:"authorId"`
}

type NewCustomFieldResponse struct {
	ID string `json:"_id"`
}

type GetCustomField struct {
	ID       string `json:"_id"`
	BoardIDs string `json:"boardIds"`
}

type EditCustomFieldRequest struct {
	Name                string `json:"name"`
	Type                string `json:"type"`
	Settings            string `json:"settings"`
	ShowOnCard          bool   `json:"showOnCard"`
	AutomaticallyOnCard bool   `json:"automaticallyOnCard"`
	AlwaysOnCard        bool   `json:"alwaysOnCard"`
	ShowLabelOnMiniCard bool   `json:"showLabelOnMiniCard"`
}

type EditCustomFieldResponse struct {
	ID string `json:"_id"`
}

type addCustomFieldDropdownItemsRequest struct {
	Items []string `json:"items"`
}

type editCustomFieldDropdownItemsRequest struct {
	Name string `json:"name"`
}
