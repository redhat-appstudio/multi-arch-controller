// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateStoragePool update storage pool
//
// swagger:model UpdateStoragePool
type UpdateStoragePool struct {

	// display name of storage pool
	DisplayName *string `json:"displayName,omitempty"`

	// indicates if the storage pool is disaster recovery (dr) enabled
	DrEnabled *bool `json:"drEnabled,omitempty"`

	// threshold override settings of a pool
	OverrideThresholds *Thresholds `json:"overrideThresholds,omitempty"`

	// state of storage pool
	// Enum: ["closed","opened"]
	State *string `json:"state,omitempty"`
}

// Validate validates this update storage pool
func (m *UpdateStoragePool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverrideThresholds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateStoragePool) validateOverrideThresholds(formats strfmt.Registry) error {
	if swag.IsZero(m.OverrideThresholds) { // not required
		return nil
	}

	if m.OverrideThresholds != nil {
		if err := m.OverrideThresholds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overrideThresholds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overrideThresholds")
			}
			return err
		}
	}

	return nil
}

var updateStoragePoolTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["closed","opened"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateStoragePoolTypeStatePropEnum = append(updateStoragePoolTypeStatePropEnum, v)
	}
}

const (

	// UpdateStoragePoolStateClosed captures enum value "closed"
	UpdateStoragePoolStateClosed string = "closed"

	// UpdateStoragePoolStateOpened captures enum value "opened"
	UpdateStoragePoolStateOpened string = "opened"
)

// prop value enum
func (m *UpdateStoragePool) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateStoragePoolTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateStoragePool) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update storage pool based on the context it is used
func (m *UpdateStoragePool) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOverrideThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateStoragePool) contextValidateOverrideThresholds(ctx context.Context, formats strfmt.Registry) error {

	if m.OverrideThresholds != nil {

		if swag.IsZero(m.OverrideThresholds) { // not required
			return nil
		}

		if err := m.OverrideThresholds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overrideThresholds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("overrideThresholds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateStoragePool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateStoragePool) UnmarshalBinary(b []byte) error {
	var res UpdateStoragePool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}