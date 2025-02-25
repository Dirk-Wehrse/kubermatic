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

// PacketSize PacketSize is the object representing Packet VM sizes.
//
// swagger:model PacketSize
type PacketSize struct {

	// c p us
	CPUs []*PacketCPU `json:"cpus"`

	// drives
	Drives []*PacketDrive `json:"drives"`

	// memory
	Memory string `json:"memory,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this packet size
func (m *PacketSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrives(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketSize) validateCPUs(formats strfmt.Registry) error {
	if swag.IsZero(m.CPUs) { // not required
		return nil
	}

	for i := 0; i < len(m.CPUs); i++ {
		if swag.IsZero(m.CPUs[i]) { // not required
			continue
		}

		if m.CPUs[i] != nil {
			if err := m.CPUs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PacketSize) validateDrives(formats strfmt.Registry) error {
	if swag.IsZero(m.Drives) { // not required
		return nil
	}

	for i := 0; i < len(m.Drives); i++ {
		if swag.IsZero(m.Drives[i]) { // not required
			continue
		}

		if m.Drives[i] != nil {
			if err := m.Drives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this packet size based on the context it is used
func (m *PacketSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCPUs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PacketSize) contextValidateCPUs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CPUs); i++ {

		if m.CPUs[i] != nil {
			if err := m.CPUs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cpus" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PacketSize) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Drives); i++ {

		if m.Drives[i] != nil {
			if err := m.Drives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PacketSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacketSize) UnmarshalBinary(b []byte) error {
	var res PacketSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
