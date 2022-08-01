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

// DefaultSession IDTokenSession is a session container for the id token
//
// swagger:model DefaultSession
type DefaultSession struct {

	// expires at
	ExpiresAt map[string]strfmt.DateTime `json:"expires_at,omitempty"`

	// headers
	Headers *Headers `json:"headers,omitempty"`

	// id token claims
	IDTokenClaims *IDTokenClaims `json:"id_token_claims,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this default session
func (m *DefaultSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIDTokenClaims(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefaultSession) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	for k := range m.ExpiresAt {

		if err := validate.FormatOf("expires_at"+"."+k, "body", "date-time", m.ExpiresAt[k].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *DefaultSession) validateHeaders(formats strfmt.Registry) error {
	if swag.IsZero(m.Headers) { // not required
		return nil
	}

	if m.Headers != nil {
		if err := m.Headers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("headers")
			}
			return err
		}
	}

	return nil
}

func (m *DefaultSession) validateIDTokenClaims(formats strfmt.Registry) error {
	if swag.IsZero(m.IDTokenClaims) { // not required
		return nil
	}

	if m.IDTokenClaims != nil {
		if err := m.IDTokenClaims.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id_token_claims")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this default session based on the context it is used
func (m *DefaultSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHeaders(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIDTokenClaims(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DefaultSession) contextValidateHeaders(ctx context.Context, formats strfmt.Registry) error {

	if m.Headers != nil {
		if err := m.Headers.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("headers")
			}
			return err
		}
	}

	return nil
}

func (m *DefaultSession) contextValidateIDTokenClaims(ctx context.Context, formats strfmt.Registry) error {

	if m.IDTokenClaims != nil {
		if err := m.IDTokenClaims.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("id_token_claims")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DefaultSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DefaultSession) UnmarshalBinary(b []byte) error {
	var res DefaultSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
