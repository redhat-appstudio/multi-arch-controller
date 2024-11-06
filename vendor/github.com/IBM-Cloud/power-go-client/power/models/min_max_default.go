// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MinMaxDefault min max default
//
// swagger:model MinMaxDefault
type MinMaxDefault struct {

	// default value
	// Required: true
	Default *float64 `json:"default"`

	// max value
	// Required: true
	Max *float64 `json:"max"`

	// min value
	// Required: true
	Min *float64 `json:"min"`
}

// Validate validates this min max default
func (m *MinMaxDefault) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MinMaxDefault) validateDefault(formats strfmt.Registry) error {

	if err := validate.Required("default", "body", m.Default); err != nil {
		return err
	}

	return nil
}

func (m *MinMaxDefault) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	return nil
}

func (m *MinMaxDefault) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this min max default based on context it is used
func (m *MinMaxDefault) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MinMaxDefault) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MinMaxDefault) UnmarshalBinary(b []byte) error {
	var res MinMaxDefault
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
