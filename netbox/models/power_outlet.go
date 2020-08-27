// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PowerOutlet power outlet
//
// swagger:model PowerOutlet
type PowerOutlet struct {

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]string `json:"connected_endpoint,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// connection status
	ConnectionStatus *PowerOutletConnectionStatus `json:"connection_status,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// feed leg
	FeedLeg *PowerOutletFeedLeg `json:"feed_leg,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// power port
	PowerPort *NestedPowerPort `json:"power_port,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *PowerOutletType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this power outlet
func (m *PowerOutlet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeedLeg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerOutlet) validateCable(formats strfmt.Registry) error {

	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateConnectionStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionStatus) { // not required
		return nil
	}

	if m.ConnectionStatus != nil {
		if err := m.ConnectionStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection_status")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 200); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateFeedLeg(formats strfmt.Registry) error {

	if swag.IsZero(m.FeedLeg) { // not required
		return nil
	}

	if m.FeedLeg != nil {
		if err := m.FeedLeg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("feed_leg")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", string(m.Label), 64); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 64); err != nil {
		return err
	}

	return nil
}

func (m *PowerOutlet) validatePowerPort(formats strfmt.Registry) error {

	if swag.IsZero(m.PowerPort) { // not required
		return nil
	}

	if m.PowerPort != nil {
		if err := m.PowerPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("power_port")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PowerOutlet) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *PowerOutlet) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutlet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutlet) UnmarshalBinary(b []byte) error {
	var res PowerOutlet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletConnectionStatus Connection status
//
// swagger:model PowerOutletConnectionStatus
type PowerOutletConnectionStatus struct {

	// label
	// Required: true
	// Enum: [Not Connected Connected]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [false true]
	Value *bool `json:"value"`
}

// Validate validates this power outlet connection status
func (m *PowerOutletConnectionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerOutletConnectionStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Not Connected","Connected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletConnectionStatusTypeLabelPropEnum = append(powerOutletConnectionStatusTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletConnectionStatusLabelNotConnected captures enum value "Not Connected"
	PowerOutletConnectionStatusLabelNotConnected string = "Not Connected"

	// PowerOutletConnectionStatusLabelConnected captures enum value "Connected"
	PowerOutletConnectionStatusLabelConnected string = "Connected"
)

// prop value enum
func (m *PowerOutletConnectionStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletConnectionStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletConnectionStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("connection_status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletConnectionStatusTypeValuePropEnum []interface{}

func init() {
	var res []bool
	if err := json.Unmarshal([]byte(`[false,true]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletConnectionStatusTypeValuePropEnum = append(powerOutletConnectionStatusTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *PowerOutletConnectionStatus) validateValueEnum(path, location string, value bool) error {
	if err := validate.EnumCase(path, location, value, powerOutletConnectionStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletConnectionStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("connection_status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("connection_status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletConnectionStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletConnectionStatus) UnmarshalBinary(b []byte) error {
	var res PowerOutletConnectionStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletFeedLeg Feed leg
//
// swagger:model PowerOutletFeedLeg
type PowerOutletFeedLeg struct {

	// label
	// Required: true
	// Enum: [A B C]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [A B C]
	Value *string `json:"value"`
}

// Validate validates this power outlet feed leg
func (m *PowerOutletFeedLeg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerOutletFeedLegTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletFeedLegTypeLabelPropEnum = append(powerOutletFeedLegTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletFeedLegLabelA captures enum value "A"
	PowerOutletFeedLegLabelA string = "A"

	// PowerOutletFeedLegLabelB captures enum value "B"
	PowerOutletFeedLegLabelB string = "B"

	// PowerOutletFeedLegLabelC captures enum value "C"
	PowerOutletFeedLegLabelC string = "C"
)

// prop value enum
func (m *PowerOutletFeedLeg) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletFeedLegTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletFeedLeg) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("feed_leg"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletFeedLegTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","B","C"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletFeedLegTypeValuePropEnum = append(powerOutletFeedLegTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletFeedLegValueA captures enum value "A"
	PowerOutletFeedLegValueA string = "A"

	// PowerOutletFeedLegValueB captures enum value "B"
	PowerOutletFeedLegValueB string = "B"

	// PowerOutletFeedLegValueC captures enum value "C"
	PowerOutletFeedLegValueC string = "C"
)

// prop value enum
func (m *PowerOutletFeedLeg) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletFeedLegTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletFeedLeg) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("feed_leg"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("feed_leg"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletFeedLeg) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletFeedLeg) UnmarshalBinary(b []byte) error {
	var res PowerOutletFeedLeg
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PowerOutletType Type
//
// swagger:model PowerOutletType
type PowerOutletType struct {

	// label
	// Required: true
	// Enum: [C5 C7 C13 C15 C19 P+N+E 4H P+N+E 6H P+N+E 9H 2P+E 4H 2P+E 6H 2P+E 9H 3P+E 4H 3P+E 6H 3P+E 9H 3P+N+E 4H 3P+N+E 6H 3P+N+E 9H NEMA 1-15R NEMA 5-15R NEMA 5-20R NEMA 5-30R NEMA 5-50R NEMA 6-15R NEMA 6-20R NEMA 6-30R NEMA 6-50R NEMA 10-30R NEMA 10-50R NEMA 14-20R NEMA 14-30R NEMA 14-50R NEMA 14-60R NEMA 15-15R NEMA 15-20R NEMA 15-30R NEMA 15-50R NEMA 15-60R NEMA L1-15R NEMA L5-15R NEMA L5-20R NEMA L5-30R NEMA L5-50R NEMA L6-15R NEMA L6-20R NEMA L6-30R NEMA L6-50R NEMA L10-30R NEMA L14-20R NEMA L14-30R NEMA L14-50R NEMA L14-60R NEMA L15-20R NEMA L15-30R NEMA L15-50R NEMA L15-60R NEMA L21-20R NEMA L21-30R CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ITA Type E (CEE7/5) ITA Type F (CEE7/3) ITA Type G (BS 1363) ITA Type H ITA Type I ITA Type J ITA Type K ITA Type L (CEI 23-50) ITA Type M (BS 546) ITA Type N ITA Type O HDOT Cx]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [iec-60320-c5 iec-60320-c7 iec-60320-c13 iec-60320-c15 iec-60320-c19 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15r nema-5-15r nema-5-20r nema-5-30r nema-5-50r nema-6-15r nema-6-20r nema-6-30r nema-6-50r nema-10-30r nema-10-50r nema-14-20r nema-14-30r nema-14-50r nema-14-60r nema-15-15r nema-15-20r nema-15-30r nema-15-50r nema-15-60r nema-l1-15r nema-l5-15r nema-l5-20r nema-l5-30r nema-l5-50r nema-l6-15r nema-l6-20r nema-l6-30r nema-l6-50r nema-l10-30r nema-l14-20r nema-l14-30r nema-l14-50r nema-l14-60r nema-l15-20r nema-l15-30r nema-l15-50r nema-l15-60r nema-l21-20r nema-l21-30r CS6360C CS6364C CS8164C CS8264C CS8364C CS8464C ita-e ita-f ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o hdot-cx]
	Value *string `json:"value"`
}

// Validate validates this power outlet type
func (m *PowerOutletType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var powerOutletTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["C5","C7","C13","C15","C19","P+N+E 4H","P+N+E 6H","P+N+E 9H","2P+E 4H","2P+E 6H","2P+E 9H","3P+E 4H","3P+E 6H","3P+E 9H","3P+N+E 4H","3P+N+E 6H","3P+N+E 9H","NEMA 1-15R","NEMA 5-15R","NEMA 5-20R","NEMA 5-30R","NEMA 5-50R","NEMA 6-15R","NEMA 6-20R","NEMA 6-30R","NEMA 6-50R","NEMA 10-30R","NEMA 10-50R","NEMA 14-20R","NEMA 14-30R","NEMA 14-50R","NEMA 14-60R","NEMA 15-15R","NEMA 15-20R","NEMA 15-30R","NEMA 15-50R","NEMA 15-60R","NEMA L1-15R","NEMA L5-15R","NEMA L5-20R","NEMA L5-30R","NEMA L5-50R","NEMA L6-15R","NEMA L6-20R","NEMA L6-30R","NEMA L6-50R","NEMA L10-30R","NEMA L14-20R","NEMA L14-30R","NEMA L14-50R","NEMA L14-60R","NEMA L15-20R","NEMA L15-30R","NEMA L15-50R","NEMA L15-60R","NEMA L21-20R","NEMA L21-30R","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ITA Type E (CEE7/5)","ITA Type F (CEE7/3)","ITA Type G (BS 1363)","ITA Type H","ITA Type I","ITA Type J","ITA Type K","ITA Type L (CEI 23-50)","ITA Type M (BS 546)","ITA Type N","ITA Type O","HDOT Cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTypeTypeLabelPropEnum = append(powerOutletTypeTypeLabelPropEnum, v)
	}
}

const (

	// PowerOutletTypeLabelC5 captures enum value "C5"
	PowerOutletTypeLabelC5 string = "C5"

	// PowerOutletTypeLabelC7 captures enum value "C7"
	PowerOutletTypeLabelC7 string = "C7"

	// PowerOutletTypeLabelC13 captures enum value "C13"
	PowerOutletTypeLabelC13 string = "C13"

	// PowerOutletTypeLabelC15 captures enum value "C15"
	PowerOutletTypeLabelC15 string = "C15"

	// PowerOutletTypeLabelC19 captures enum value "C19"
	PowerOutletTypeLabelC19 string = "C19"

	// PowerOutletTypeLabelPNE4H captures enum value "P+N+E 4H"
	PowerOutletTypeLabelPNE4H string = "P+N+E 4H"

	// PowerOutletTypeLabelPNE6H captures enum value "P+N+E 6H"
	PowerOutletTypeLabelPNE6H string = "P+N+E 6H"

	// PowerOutletTypeLabelPNE9H captures enum value "P+N+E 9H"
	PowerOutletTypeLabelPNE9H string = "P+N+E 9H"

	// PowerOutletTypeLabelNr2PE4H captures enum value "2P+E 4H"
	PowerOutletTypeLabelNr2PE4H string = "2P+E 4H"

	// PowerOutletTypeLabelNr2PE6H captures enum value "2P+E 6H"
	PowerOutletTypeLabelNr2PE6H string = "2P+E 6H"

	// PowerOutletTypeLabelNr2PE9H captures enum value "2P+E 9H"
	PowerOutletTypeLabelNr2PE9H string = "2P+E 9H"

	// PowerOutletTypeLabelNr3PE4H captures enum value "3P+E 4H"
	PowerOutletTypeLabelNr3PE4H string = "3P+E 4H"

	// PowerOutletTypeLabelNr3PE6H captures enum value "3P+E 6H"
	PowerOutletTypeLabelNr3PE6H string = "3P+E 6H"

	// PowerOutletTypeLabelNr3PE9H captures enum value "3P+E 9H"
	PowerOutletTypeLabelNr3PE9H string = "3P+E 9H"

	// PowerOutletTypeLabelNr3PNE4H captures enum value "3P+N+E 4H"
	PowerOutletTypeLabelNr3PNE4H string = "3P+N+E 4H"

	// PowerOutletTypeLabelNr3PNE6H captures enum value "3P+N+E 6H"
	PowerOutletTypeLabelNr3PNE6H string = "3P+N+E 6H"

	// PowerOutletTypeLabelNr3PNE9H captures enum value "3P+N+E 9H"
	PowerOutletTypeLabelNr3PNE9H string = "3P+N+E 9H"

	// PowerOutletTypeLabelNEMA115R captures enum value "NEMA 1-15R"
	PowerOutletTypeLabelNEMA115R string = "NEMA 1-15R"

	// PowerOutletTypeLabelNEMA515R captures enum value "NEMA 5-15R"
	PowerOutletTypeLabelNEMA515R string = "NEMA 5-15R"

	// PowerOutletTypeLabelNEMA520R captures enum value "NEMA 5-20R"
	PowerOutletTypeLabelNEMA520R string = "NEMA 5-20R"

	// PowerOutletTypeLabelNEMA530R captures enum value "NEMA 5-30R"
	PowerOutletTypeLabelNEMA530R string = "NEMA 5-30R"

	// PowerOutletTypeLabelNEMA550R captures enum value "NEMA 5-50R"
	PowerOutletTypeLabelNEMA550R string = "NEMA 5-50R"

	// PowerOutletTypeLabelNEMA615R captures enum value "NEMA 6-15R"
	PowerOutletTypeLabelNEMA615R string = "NEMA 6-15R"

	// PowerOutletTypeLabelNEMA620R captures enum value "NEMA 6-20R"
	PowerOutletTypeLabelNEMA620R string = "NEMA 6-20R"

	// PowerOutletTypeLabelNEMA630R captures enum value "NEMA 6-30R"
	PowerOutletTypeLabelNEMA630R string = "NEMA 6-30R"

	// PowerOutletTypeLabelNEMA650R captures enum value "NEMA 6-50R"
	PowerOutletTypeLabelNEMA650R string = "NEMA 6-50R"

	// PowerOutletTypeLabelNEMA1030R captures enum value "NEMA 10-30R"
	PowerOutletTypeLabelNEMA1030R string = "NEMA 10-30R"

	// PowerOutletTypeLabelNEMA1050R captures enum value "NEMA 10-50R"
	PowerOutletTypeLabelNEMA1050R string = "NEMA 10-50R"

	// PowerOutletTypeLabelNEMA1420R captures enum value "NEMA 14-20R"
	PowerOutletTypeLabelNEMA1420R string = "NEMA 14-20R"

	// PowerOutletTypeLabelNEMA1430R captures enum value "NEMA 14-30R"
	PowerOutletTypeLabelNEMA1430R string = "NEMA 14-30R"

	// PowerOutletTypeLabelNEMA1450R captures enum value "NEMA 14-50R"
	PowerOutletTypeLabelNEMA1450R string = "NEMA 14-50R"

	// PowerOutletTypeLabelNEMA1460R captures enum value "NEMA 14-60R"
	PowerOutletTypeLabelNEMA1460R string = "NEMA 14-60R"

	// PowerOutletTypeLabelNEMA1515R captures enum value "NEMA 15-15R"
	PowerOutletTypeLabelNEMA1515R string = "NEMA 15-15R"

	// PowerOutletTypeLabelNEMA1520R captures enum value "NEMA 15-20R"
	PowerOutletTypeLabelNEMA1520R string = "NEMA 15-20R"

	// PowerOutletTypeLabelNEMA1530R captures enum value "NEMA 15-30R"
	PowerOutletTypeLabelNEMA1530R string = "NEMA 15-30R"

	// PowerOutletTypeLabelNEMA1550R captures enum value "NEMA 15-50R"
	PowerOutletTypeLabelNEMA1550R string = "NEMA 15-50R"

	// PowerOutletTypeLabelNEMA1560R captures enum value "NEMA 15-60R"
	PowerOutletTypeLabelNEMA1560R string = "NEMA 15-60R"

	// PowerOutletTypeLabelNEMAL115R captures enum value "NEMA L1-15R"
	PowerOutletTypeLabelNEMAL115R string = "NEMA L1-15R"

	// PowerOutletTypeLabelNEMAL515R captures enum value "NEMA L5-15R"
	PowerOutletTypeLabelNEMAL515R string = "NEMA L5-15R"

	// PowerOutletTypeLabelNEMAL520R captures enum value "NEMA L5-20R"
	PowerOutletTypeLabelNEMAL520R string = "NEMA L5-20R"

	// PowerOutletTypeLabelNEMAL530R captures enum value "NEMA L5-30R"
	PowerOutletTypeLabelNEMAL530R string = "NEMA L5-30R"

	// PowerOutletTypeLabelNEMAL550R captures enum value "NEMA L5-50R"
	PowerOutletTypeLabelNEMAL550R string = "NEMA L5-50R"

	// PowerOutletTypeLabelNEMAL615R captures enum value "NEMA L6-15R"
	PowerOutletTypeLabelNEMAL615R string = "NEMA L6-15R"

	// PowerOutletTypeLabelNEMAL620R captures enum value "NEMA L6-20R"
	PowerOutletTypeLabelNEMAL620R string = "NEMA L6-20R"

	// PowerOutletTypeLabelNEMAL630R captures enum value "NEMA L6-30R"
	PowerOutletTypeLabelNEMAL630R string = "NEMA L6-30R"

	// PowerOutletTypeLabelNEMAL650R captures enum value "NEMA L6-50R"
	PowerOutletTypeLabelNEMAL650R string = "NEMA L6-50R"

	// PowerOutletTypeLabelNEMAL1030R captures enum value "NEMA L10-30R"
	PowerOutletTypeLabelNEMAL1030R string = "NEMA L10-30R"

	// PowerOutletTypeLabelNEMAL1420R captures enum value "NEMA L14-20R"
	PowerOutletTypeLabelNEMAL1420R string = "NEMA L14-20R"

	// PowerOutletTypeLabelNEMAL1430R captures enum value "NEMA L14-30R"
	PowerOutletTypeLabelNEMAL1430R string = "NEMA L14-30R"

	// PowerOutletTypeLabelNEMAL1450R captures enum value "NEMA L14-50R"
	PowerOutletTypeLabelNEMAL1450R string = "NEMA L14-50R"

	// PowerOutletTypeLabelNEMAL1460R captures enum value "NEMA L14-60R"
	PowerOutletTypeLabelNEMAL1460R string = "NEMA L14-60R"

	// PowerOutletTypeLabelNEMAL1520R captures enum value "NEMA L15-20R"
	PowerOutletTypeLabelNEMAL1520R string = "NEMA L15-20R"

	// PowerOutletTypeLabelNEMAL1530R captures enum value "NEMA L15-30R"
	PowerOutletTypeLabelNEMAL1530R string = "NEMA L15-30R"

	// PowerOutletTypeLabelNEMAL1550R captures enum value "NEMA L15-50R"
	PowerOutletTypeLabelNEMAL1550R string = "NEMA L15-50R"

	// PowerOutletTypeLabelNEMAL1560R captures enum value "NEMA L15-60R"
	PowerOutletTypeLabelNEMAL1560R string = "NEMA L15-60R"

	// PowerOutletTypeLabelNEMAL2120R captures enum value "NEMA L21-20R"
	PowerOutletTypeLabelNEMAL2120R string = "NEMA L21-20R"

	// PowerOutletTypeLabelNEMAL2130R captures enum value "NEMA L21-30R"
	PowerOutletTypeLabelNEMAL2130R string = "NEMA L21-30R"

	// PowerOutletTypeLabelCS6360C captures enum value "CS6360C"
	PowerOutletTypeLabelCS6360C string = "CS6360C"

	// PowerOutletTypeLabelCS6364C captures enum value "CS6364C"
	PowerOutletTypeLabelCS6364C string = "CS6364C"

	// PowerOutletTypeLabelCS8164C captures enum value "CS8164C"
	PowerOutletTypeLabelCS8164C string = "CS8164C"

	// PowerOutletTypeLabelCS8264C captures enum value "CS8264C"
	PowerOutletTypeLabelCS8264C string = "CS8264C"

	// PowerOutletTypeLabelCS8364C captures enum value "CS8364C"
	PowerOutletTypeLabelCS8364C string = "CS8364C"

	// PowerOutletTypeLabelCS8464C captures enum value "CS8464C"
	PowerOutletTypeLabelCS8464C string = "CS8464C"

	// PowerOutletTypeLabelITATypeECEE75 captures enum value "ITA Type E (CEE7/5)"
	PowerOutletTypeLabelITATypeECEE75 string = "ITA Type E (CEE7/5)"

	// PowerOutletTypeLabelITATypeFCEE73 captures enum value "ITA Type F (CEE7/3)"
	PowerOutletTypeLabelITATypeFCEE73 string = "ITA Type F (CEE7/3)"

	// PowerOutletTypeLabelITATypeGBS1363 captures enum value "ITA Type G (BS 1363)"
	PowerOutletTypeLabelITATypeGBS1363 string = "ITA Type G (BS 1363)"

	// PowerOutletTypeLabelITATypeH captures enum value "ITA Type H"
	PowerOutletTypeLabelITATypeH string = "ITA Type H"

	// PowerOutletTypeLabelITATypeI captures enum value "ITA Type I"
	PowerOutletTypeLabelITATypeI string = "ITA Type I"

	// PowerOutletTypeLabelITATypeJ captures enum value "ITA Type J"
	PowerOutletTypeLabelITATypeJ string = "ITA Type J"

	// PowerOutletTypeLabelITATypeK captures enum value "ITA Type K"
	PowerOutletTypeLabelITATypeK string = "ITA Type K"

	// PowerOutletTypeLabelITATypeLCEI2350 captures enum value "ITA Type L (CEI 23-50)"
	PowerOutletTypeLabelITATypeLCEI2350 string = "ITA Type L (CEI 23-50)"

	// PowerOutletTypeLabelITATypeMBS546 captures enum value "ITA Type M (BS 546)"
	PowerOutletTypeLabelITATypeMBS546 string = "ITA Type M (BS 546)"

	// PowerOutletTypeLabelITATypeN captures enum value "ITA Type N"
	PowerOutletTypeLabelITATypeN string = "ITA Type N"

	// PowerOutletTypeLabelITATypeO captures enum value "ITA Type O"
	PowerOutletTypeLabelITATypeO string = "ITA Type O"

	// PowerOutletTypeLabelHDOTCx captures enum value "HDOT Cx"
	PowerOutletTypeLabelHDOTCx string = "HDOT Cx"
)

// prop value enum
func (m *PowerOutletType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var powerOutletTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c5","iec-60320-c7","iec-60320-c13","iec-60320-c15","iec-60320-c19","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15r","nema-5-15r","nema-5-20r","nema-5-30r","nema-5-50r","nema-6-15r","nema-6-20r","nema-6-30r","nema-6-50r","nema-10-30r","nema-10-50r","nema-14-20r","nema-14-30r","nema-14-50r","nema-14-60r","nema-15-15r","nema-15-20r","nema-15-30r","nema-15-50r","nema-15-60r","nema-l1-15r","nema-l5-15r","nema-l5-20r","nema-l5-30r","nema-l5-50r","nema-l6-15r","nema-l6-20r","nema-l6-30r","nema-l6-50r","nema-l10-30r","nema-l14-20r","nema-l14-30r","nema-l14-50r","nema-l14-60r","nema-l15-20r","nema-l15-30r","nema-l15-50r","nema-l15-60r","nema-l21-20r","nema-l21-30r","CS6360C","CS6364C","CS8164C","CS8264C","CS8364C","CS8464C","ita-e","ita-f","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","hdot-cx"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		powerOutletTypeTypeValuePropEnum = append(powerOutletTypeTypeValuePropEnum, v)
	}
}

const (

	// PowerOutletTypeValueIec60320C5 captures enum value "iec-60320-c5"
	PowerOutletTypeValueIec60320C5 string = "iec-60320-c5"

	// PowerOutletTypeValueIec60320C7 captures enum value "iec-60320-c7"
	PowerOutletTypeValueIec60320C7 string = "iec-60320-c7"

	// PowerOutletTypeValueIec60320C13 captures enum value "iec-60320-c13"
	PowerOutletTypeValueIec60320C13 string = "iec-60320-c13"

	// PowerOutletTypeValueIec60320C15 captures enum value "iec-60320-c15"
	PowerOutletTypeValueIec60320C15 string = "iec-60320-c15"

	// PowerOutletTypeValueIec60320C19 captures enum value "iec-60320-c19"
	PowerOutletTypeValueIec60320C19 string = "iec-60320-c19"

	// PowerOutletTypeValueIec60309pne4h captures enum value "iec-60309-p-n-e-4h"
	PowerOutletTypeValueIec60309pne4h string = "iec-60309-p-n-e-4h"

	// PowerOutletTypeValueIec60309pne6h captures enum value "iec-60309-p-n-e-6h"
	PowerOutletTypeValueIec60309pne6h string = "iec-60309-p-n-e-6h"

	// PowerOutletTypeValueIec60309pne9h captures enum value "iec-60309-p-n-e-9h"
	PowerOutletTypeValueIec60309pne9h string = "iec-60309-p-n-e-9h"

	// PowerOutletTypeValueIec603092pe4h captures enum value "iec-60309-2p-e-4h"
	PowerOutletTypeValueIec603092pe4h string = "iec-60309-2p-e-4h"

	// PowerOutletTypeValueIec603092pe6h captures enum value "iec-60309-2p-e-6h"
	PowerOutletTypeValueIec603092pe6h string = "iec-60309-2p-e-6h"

	// PowerOutletTypeValueIec603092pe9h captures enum value "iec-60309-2p-e-9h"
	PowerOutletTypeValueIec603092pe9h string = "iec-60309-2p-e-9h"

	// PowerOutletTypeValueIec603093pe4h captures enum value "iec-60309-3p-e-4h"
	PowerOutletTypeValueIec603093pe4h string = "iec-60309-3p-e-4h"

	// PowerOutletTypeValueIec603093pe6h captures enum value "iec-60309-3p-e-6h"
	PowerOutletTypeValueIec603093pe6h string = "iec-60309-3p-e-6h"

	// PowerOutletTypeValueIec603093pe9h captures enum value "iec-60309-3p-e-9h"
	PowerOutletTypeValueIec603093pe9h string = "iec-60309-3p-e-9h"

	// PowerOutletTypeValueIec603093pne4h captures enum value "iec-60309-3p-n-e-4h"
	PowerOutletTypeValueIec603093pne4h string = "iec-60309-3p-n-e-4h"

	// PowerOutletTypeValueIec603093pne6h captures enum value "iec-60309-3p-n-e-6h"
	PowerOutletTypeValueIec603093pne6h string = "iec-60309-3p-n-e-6h"

	// PowerOutletTypeValueIec603093pne9h captures enum value "iec-60309-3p-n-e-9h"
	PowerOutletTypeValueIec603093pne9h string = "iec-60309-3p-n-e-9h"

	// PowerOutletTypeValueNema115r captures enum value "nema-1-15r"
	PowerOutletTypeValueNema115r string = "nema-1-15r"

	// PowerOutletTypeValueNema515r captures enum value "nema-5-15r"
	PowerOutletTypeValueNema515r string = "nema-5-15r"

	// PowerOutletTypeValueNema520r captures enum value "nema-5-20r"
	PowerOutletTypeValueNema520r string = "nema-5-20r"

	// PowerOutletTypeValueNema530r captures enum value "nema-5-30r"
	PowerOutletTypeValueNema530r string = "nema-5-30r"

	// PowerOutletTypeValueNema550r captures enum value "nema-5-50r"
	PowerOutletTypeValueNema550r string = "nema-5-50r"

	// PowerOutletTypeValueNema615r captures enum value "nema-6-15r"
	PowerOutletTypeValueNema615r string = "nema-6-15r"

	// PowerOutletTypeValueNema620r captures enum value "nema-6-20r"
	PowerOutletTypeValueNema620r string = "nema-6-20r"

	// PowerOutletTypeValueNema630r captures enum value "nema-6-30r"
	PowerOutletTypeValueNema630r string = "nema-6-30r"

	// PowerOutletTypeValueNema650r captures enum value "nema-6-50r"
	PowerOutletTypeValueNema650r string = "nema-6-50r"

	// PowerOutletTypeValueNema1030r captures enum value "nema-10-30r"
	PowerOutletTypeValueNema1030r string = "nema-10-30r"

	// PowerOutletTypeValueNema1050r captures enum value "nema-10-50r"
	PowerOutletTypeValueNema1050r string = "nema-10-50r"

	// PowerOutletTypeValueNema1420r captures enum value "nema-14-20r"
	PowerOutletTypeValueNema1420r string = "nema-14-20r"

	// PowerOutletTypeValueNema1430r captures enum value "nema-14-30r"
	PowerOutletTypeValueNema1430r string = "nema-14-30r"

	// PowerOutletTypeValueNema1450r captures enum value "nema-14-50r"
	PowerOutletTypeValueNema1450r string = "nema-14-50r"

	// PowerOutletTypeValueNema1460r captures enum value "nema-14-60r"
	PowerOutletTypeValueNema1460r string = "nema-14-60r"

	// PowerOutletTypeValueNema1515r captures enum value "nema-15-15r"
	PowerOutletTypeValueNema1515r string = "nema-15-15r"

	// PowerOutletTypeValueNema1520r captures enum value "nema-15-20r"
	PowerOutletTypeValueNema1520r string = "nema-15-20r"

	// PowerOutletTypeValueNema1530r captures enum value "nema-15-30r"
	PowerOutletTypeValueNema1530r string = "nema-15-30r"

	// PowerOutletTypeValueNema1550r captures enum value "nema-15-50r"
	PowerOutletTypeValueNema1550r string = "nema-15-50r"

	// PowerOutletTypeValueNema1560r captures enum value "nema-15-60r"
	PowerOutletTypeValueNema1560r string = "nema-15-60r"

	// PowerOutletTypeValueNemaL115r captures enum value "nema-l1-15r"
	PowerOutletTypeValueNemaL115r string = "nema-l1-15r"

	// PowerOutletTypeValueNemaL515r captures enum value "nema-l5-15r"
	PowerOutletTypeValueNemaL515r string = "nema-l5-15r"

	// PowerOutletTypeValueNemaL520r captures enum value "nema-l5-20r"
	PowerOutletTypeValueNemaL520r string = "nema-l5-20r"

	// PowerOutletTypeValueNemaL530r captures enum value "nema-l5-30r"
	PowerOutletTypeValueNemaL530r string = "nema-l5-30r"

	// PowerOutletTypeValueNemaL550r captures enum value "nema-l5-50r"
	PowerOutletTypeValueNemaL550r string = "nema-l5-50r"

	// PowerOutletTypeValueNemaL615r captures enum value "nema-l6-15r"
	PowerOutletTypeValueNemaL615r string = "nema-l6-15r"

	// PowerOutletTypeValueNemaL620r captures enum value "nema-l6-20r"
	PowerOutletTypeValueNemaL620r string = "nema-l6-20r"

	// PowerOutletTypeValueNemaL630r captures enum value "nema-l6-30r"
	PowerOutletTypeValueNemaL630r string = "nema-l6-30r"

	// PowerOutletTypeValueNemaL650r captures enum value "nema-l6-50r"
	PowerOutletTypeValueNemaL650r string = "nema-l6-50r"

	// PowerOutletTypeValueNemaL1030r captures enum value "nema-l10-30r"
	PowerOutletTypeValueNemaL1030r string = "nema-l10-30r"

	// PowerOutletTypeValueNemaL1420r captures enum value "nema-l14-20r"
	PowerOutletTypeValueNemaL1420r string = "nema-l14-20r"

	// PowerOutletTypeValueNemaL1430r captures enum value "nema-l14-30r"
	PowerOutletTypeValueNemaL1430r string = "nema-l14-30r"

	// PowerOutletTypeValueNemaL1450r captures enum value "nema-l14-50r"
	PowerOutletTypeValueNemaL1450r string = "nema-l14-50r"

	// PowerOutletTypeValueNemaL1460r captures enum value "nema-l14-60r"
	PowerOutletTypeValueNemaL1460r string = "nema-l14-60r"

	// PowerOutletTypeValueNemaL1520r captures enum value "nema-l15-20r"
	PowerOutletTypeValueNemaL1520r string = "nema-l15-20r"

	// PowerOutletTypeValueNemaL1530r captures enum value "nema-l15-30r"
	PowerOutletTypeValueNemaL1530r string = "nema-l15-30r"

	// PowerOutletTypeValueNemaL1550r captures enum value "nema-l15-50r"
	PowerOutletTypeValueNemaL1550r string = "nema-l15-50r"

	// PowerOutletTypeValueNemaL1560r captures enum value "nema-l15-60r"
	PowerOutletTypeValueNemaL1560r string = "nema-l15-60r"

	// PowerOutletTypeValueNemaL2120r captures enum value "nema-l21-20r"
	PowerOutletTypeValueNemaL2120r string = "nema-l21-20r"

	// PowerOutletTypeValueNemaL2130r captures enum value "nema-l21-30r"
	PowerOutletTypeValueNemaL2130r string = "nema-l21-30r"

	// PowerOutletTypeValueCS6360C captures enum value "CS6360C"
	PowerOutletTypeValueCS6360C string = "CS6360C"

	// PowerOutletTypeValueCS6364C captures enum value "CS6364C"
	PowerOutletTypeValueCS6364C string = "CS6364C"

	// PowerOutletTypeValueCS8164C captures enum value "CS8164C"
	PowerOutletTypeValueCS8164C string = "CS8164C"

	// PowerOutletTypeValueCS8264C captures enum value "CS8264C"
	PowerOutletTypeValueCS8264C string = "CS8264C"

	// PowerOutletTypeValueCS8364C captures enum value "CS8364C"
	PowerOutletTypeValueCS8364C string = "CS8364C"

	// PowerOutletTypeValueCS8464C captures enum value "CS8464C"
	PowerOutletTypeValueCS8464C string = "CS8464C"

	// PowerOutletTypeValueItae captures enum value "ita-e"
	PowerOutletTypeValueItae string = "ita-e"

	// PowerOutletTypeValueItaf captures enum value "ita-f"
	PowerOutletTypeValueItaf string = "ita-f"

	// PowerOutletTypeValueItag captures enum value "ita-g"
	PowerOutletTypeValueItag string = "ita-g"

	// PowerOutletTypeValueItah captures enum value "ita-h"
	PowerOutletTypeValueItah string = "ita-h"

	// PowerOutletTypeValueItai captures enum value "ita-i"
	PowerOutletTypeValueItai string = "ita-i"

	// PowerOutletTypeValueItaj captures enum value "ita-j"
	PowerOutletTypeValueItaj string = "ita-j"

	// PowerOutletTypeValueItak captures enum value "ita-k"
	PowerOutletTypeValueItak string = "ita-k"

	// PowerOutletTypeValueItal captures enum value "ita-l"
	PowerOutletTypeValueItal string = "ita-l"

	// PowerOutletTypeValueItam captures enum value "ita-m"
	PowerOutletTypeValueItam string = "ita-m"

	// PowerOutletTypeValueItan captures enum value "ita-n"
	PowerOutletTypeValueItan string = "ita-n"

	// PowerOutletTypeValueItao captures enum value "ita-o"
	PowerOutletTypeValueItao string = "ita-o"

	// PowerOutletTypeValueHdotCx captures enum value "hdot-cx"
	PowerOutletTypeValueHdotCx string = "hdot-cx"
)

// prop value enum
func (m *PowerOutletType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, powerOutletTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PowerOutletType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerOutletType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerOutletType) UnmarshalBinary(b []byte) error {
	var res PowerOutletType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
