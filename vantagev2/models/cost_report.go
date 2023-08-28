// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CostReport CostReport model
//
// swagger:model CostReport
type CostReport struct {

	// The date and time, in UTC, the report was created. ISO 8601 Formatted.
	// Example: 2021-07-09T00:00:00Z
	CreatedAt string `json:"created_at,omitempty"`

	// The filter applied to the CostReport.
	Filter string `json:"filter,omitempty"`

	// The token for the Folder the CostReport is a part of.
	FolderToken string `json:"folder_token,omitempty"`

	// The tokens for the SavedFilters assigned to the CostReport.
	SavedFilterTokens []string `json:"saved_filter_tokens"`

	// The title of the CostReport.
	// Example: Production Environment
	Title string `json:"title,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// The token for the Workspace the CostReport is a part of.
	WorkspaceToken string `json:"workspace_token,omitempty"`
}

// Validate validates this cost report
func (m *CostReport) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cost report based on context it is used
func (m *CostReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CostReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CostReport) UnmarshalBinary(b []byte) error {
	var res CostReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
