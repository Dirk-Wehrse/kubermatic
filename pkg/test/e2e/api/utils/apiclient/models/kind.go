// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Kind Kind specifies the resource Kind and APIGroup
//
// swagger:model Kind
type Kind struct {

	// APIGroups specifies the APIGroup of the resource
	APIGroups string `json:"apiGroups,omitempty"`

	// Kinds specifies the kind of the resource
	Kinds string `json:"kinds,omitempty"`
}

// Validate validates this kind
func (m *Kind) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Kind) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Kind) UnmarshalBinary(b []byte) error {
	var res Kind
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
