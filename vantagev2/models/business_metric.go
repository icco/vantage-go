// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BusinessMetric BusinessMetric model
//
// swagger:model BusinessMetric
type BusinessMetric struct {

	// The tokens for any CostReports that use the BusinessMetric, and the unit scale.
	CostReportTokensWithMetadata []*AttachedCostReportForBusinessMetric `json:"cost_report_tokens_with_metadata"`

	// The token of the User who created the BusinessMetric.
	// Example: usr_1234
	CreatedByToken string `json:"created_by_token,omitempty"`

	// The title of the BusinessMetric.
	// Example: Total Revenue
	Title string `json:"title,omitempty"`

	// The token of the BusinessMetric.
	// Example: bsnss_mtrc_1234
	Token string `json:"token,omitempty"`

	// The dates, amounts, and (optional) labels for the BusinessMetric.
	Values []*BusinessMetricValue `json:"values"`
}

// Validate validates this business metric
func (m *BusinessMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCostReportTokensWithMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BusinessMetric) validateCostReportTokensWithMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.CostReportTokensWithMetadata) { // not required
		return nil
	}

	for i := 0; i < len(m.CostReportTokensWithMetadata); i++ {
		if swag.IsZero(m.CostReportTokensWithMetadata[i]) { // not required
			continue
		}

		if m.CostReportTokensWithMetadata[i] != nil {
			if err := m.CostReportTokensWithMetadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BusinessMetric) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {
		if swag.IsZero(m.Values[i]) { // not required
			continue
		}

		if m.Values[i] != nil {
			if err := m.Values[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this business metric based on the context it is used
func (m *BusinessMetric) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCostReportTokensWithMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BusinessMetric) contextValidateCostReportTokensWithMetadata(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CostReportTokensWithMetadata); i++ {

		if m.CostReportTokensWithMetadata[i] != nil {

			if swag.IsZero(m.CostReportTokensWithMetadata[i]) { // not required
				return nil
			}

			if err := m.CostReportTokensWithMetadata[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cost_report_tokens_with_metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BusinessMetric) contextValidateValues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Values); i++ {

		if m.Values[i] != nil {

			if swag.IsZero(m.Values[i]) { // not required
				return nil
			}

			if err := m.Values[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("values" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("values" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BusinessMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessMetric) UnmarshalBinary(b []byte) error {
	var res BusinessMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
