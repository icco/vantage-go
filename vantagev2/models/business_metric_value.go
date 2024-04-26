// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BusinessMetricValue business metric value
//
// swagger:model BusinessMetricValue
type BusinessMetricValue struct {

	// The amount of the Business Metric Value as a string to ensure precision.
	// Example: 100.00
	Amount string `json:"amount,omitempty"`

	// The date of the Business Metric Value. ISO 8601 formatted.
	// Example: 2024-03-01+00:00
	Date string `json:"date,omitempty"`

	// The label of the Business Metric Value.
	// Example: Cost Center A
	Label string `json:"label,omitempty"`
}

// Validate validates this business metric value
func (m *BusinessMetricValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this business metric value based on context it is used
func (m *BusinessMetricValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BusinessMetricValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessMetricValue) UnmarshalBinary(b []byte) error {
	var res BusinessMetricValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
