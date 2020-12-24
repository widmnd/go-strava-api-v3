// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActivityTotal A roll-up of metrics pertaining to a set of activities. Values are in seconds and meters.
//
// swagger:model activityTotal
type ActivityTotal struct {

	// The total number of achievements of the considered activities.
	AchievementCount int64 `json:"achievement_count,omitempty"`

	// The number of activities considered in this total.
	Count int64 `json:"count,omitempty"`

	// The total distance covered by the considered activities.
	Distance float32 `json:"distance,omitempty"`

	// The total elapsed time of the considered activities.
	ElapsedTime int64 `json:"elapsed_time,omitempty"`

	// The total elevation gain of the considered activities.
	ElevationGain float32 `json:"elevation_gain,omitempty"`

	// The total moving time of the considered activities.
	MovingTime int64 `json:"moving_time,omitempty"`
}

// Validate validates this activity total
func (m *ActivityTotal) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ActivityTotal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityTotal) UnmarshalBinary(b []byte) error {
	var res ActivityTotal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
