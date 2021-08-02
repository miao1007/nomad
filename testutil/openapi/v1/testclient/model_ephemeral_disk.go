/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testclient

import (
	"encoding/json"
)

// EphemeralDisk struct for EphemeralDisk
type EphemeralDisk struct {
	Migrate *bool `json:"Migrate,omitempty"`
	SizeMB *int32 `json:"SizeMB,omitempty"`
	Sticky *bool `json:"Sticky,omitempty"`
}

// NewEphemeralDisk instantiates a new EphemeralDisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEphemeralDisk() *EphemeralDisk {
	this := EphemeralDisk{}
	return &this
}

// NewEphemeralDiskWithDefaults instantiates a new EphemeralDisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEphemeralDiskWithDefaults() *EphemeralDisk {
	this := EphemeralDisk{}
	return &this
}

// GetMigrate returns the Migrate field value if set, zero value otherwise.
func (o *EphemeralDisk) GetMigrate() bool {
	if o == nil || o.Migrate == nil {
		var ret bool
		return ret
	}
	return *o.Migrate
}

// GetMigrateOk returns a tuple with the Migrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralDisk) GetMigrateOk() (*bool, bool) {
	if o == nil || o.Migrate == nil {
		return nil, false
	}
	return o.Migrate, true
}

// HasMigrate returns a boolean if a field has been set.
func (o *EphemeralDisk) HasMigrate() bool {
	if o != nil && o.Migrate != nil {
		return true
	}

	return false
}

// SetMigrate gets a reference to the given bool and assigns it to the Migrate field.
func (o *EphemeralDisk) SetMigrate(v bool) {
	o.Migrate = &v
}

// GetSizeMB returns the SizeMB field value if set, zero value otherwise.
func (o *EphemeralDisk) GetSizeMB() int32 {
	if o == nil || o.SizeMB == nil {
		var ret int32
		return ret
	}
	return *o.SizeMB
}

// GetSizeMBOk returns a tuple with the SizeMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralDisk) GetSizeMBOk() (*int32, bool) {
	if o == nil || o.SizeMB == nil {
		return nil, false
	}
	return o.SizeMB, true
}

// HasSizeMB returns a boolean if a field has been set.
func (o *EphemeralDisk) HasSizeMB() bool {
	if o != nil && o.SizeMB != nil {
		return true
	}

	return false
}

// SetSizeMB gets a reference to the given int32 and assigns it to the SizeMB field.
func (o *EphemeralDisk) SetSizeMB(v int32) {
	o.SizeMB = &v
}

// GetSticky returns the Sticky field value if set, zero value otherwise.
func (o *EphemeralDisk) GetSticky() bool {
	if o == nil || o.Sticky == nil {
		var ret bool
		return ret
	}
	return *o.Sticky
}

// GetStickyOk returns a tuple with the Sticky field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralDisk) GetStickyOk() (*bool, bool) {
	if o == nil || o.Sticky == nil {
		return nil, false
	}
	return o.Sticky, true
}

// HasSticky returns a boolean if a field has been set.
func (o *EphemeralDisk) HasSticky() bool {
	if o != nil && o.Sticky != nil {
		return true
	}

	return false
}

// SetSticky gets a reference to the given bool and assigns it to the Sticky field.
func (o *EphemeralDisk) SetSticky(v bool) {
	o.Sticky = &v
}

func (o EphemeralDisk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Migrate != nil {
		toSerialize["Migrate"] = o.Migrate
	}
	if o.SizeMB != nil {
		toSerialize["SizeMB"] = o.SizeMB
	}
	if o.Sticky != nil {
		toSerialize["Sticky"] = o.Sticky
	}
	return json.Marshal(toSerialize)
}

type NullableEphemeralDisk struct {
	value *EphemeralDisk
	isSet bool
}

func (v NullableEphemeralDisk) Get() *EphemeralDisk {
	return v.value
}

func (v *NullableEphemeralDisk) Set(val *EphemeralDisk) {
	v.value = val
	v.isSet = true
}

func (v NullableEphemeralDisk) IsSet() bool {
	return v.isSet
}

func (v *NullableEphemeralDisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEphemeralDisk(val *EphemeralDisk) *NullableEphemeralDisk {
	return &NullableEphemeralDisk{value: val, isSet: true}
}

func (v NullableEphemeralDisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEphemeralDisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


