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
	"github.com/go-openapi/validate"
)

// SAPProfiles s a p profiles
//
// swagger:model SAPProfiles
type SAPProfiles struct {

	// SAP Profiles
	// Required: true
	Profiles []*SAPProfile `json:"profiles"`
}

// Validate validates this s a p profiles
func (m *SAPProfiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAPProfiles) validateProfiles(formats strfmt.Registry) error {

	if err := validate.Required("profiles", "body", m.Profiles); err != nil {
		return err
	}

	for i := 0; i < len(m.Profiles); i++ {
		if swag.IsZero(m.Profiles[i]) { // not required
			continue
		}

		if m.Profiles[i] != nil {
			if err := m.Profiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this s a p profiles based on the context it is used
func (m *SAPProfiles) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProfiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAPProfiles) contextValidateProfiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Profiles); i++ {

		if m.Profiles[i] != nil {

			if swag.IsZero(m.Profiles[i]) { // not required
				return nil
			}

			if err := m.Profiles[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profiles" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("profiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SAPProfiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAPProfiles) UnmarshalBinary(b []byte) error {
	var res SAPProfiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
