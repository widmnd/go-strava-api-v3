// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdatableActivity updatable activity
//
// swagger:model updatableActivity
type UpdatableActivity struct {

	// Whether this activity is a commute
	Commute bool `json:"commute,omitempty"`

	// The description of the activity
	Description string `json:"description,omitempty"`

	// Identifier for the gear associated with the activity. ‘none’ clears gear from activity
	GearID string `json:"gear_id,omitempty"`

	// The name of the activity
	Name string `json:"name,omitempty"`

	// Whether this activity was recorded on a training machine
	Trainer bool `json:"trainer,omitempty"`

	// type
	Type ActivityType `json:"type,omitempty"`
}

// Validate validates this updatable activity
func (m *UpdatableActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatableActivity) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatableActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatableActivity) UnmarshalBinary(b []byte) error {
	var res UpdatableActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
