// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChannelInfo channel info
// swagger:model channelInfo
type ChannelInfo struct {

	// channel ID
	// Required: true
	ChannelID *string `json:"channelID"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this channel info
func (m *ChannelInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelInfo) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channelID", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelInfo) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelInfo) UnmarshalBinary(b []byte) error {
	var res ChannelInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
