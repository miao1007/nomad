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

// ConsulTerminatingConfigEntry struct for ConsulTerminatingConfigEntry
type ConsulTerminatingConfigEntry struct {
	Services *[]ConsulLinkedService `json:"Services,omitempty"`
}

// NewConsulTerminatingConfigEntry instantiates a new ConsulTerminatingConfigEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulTerminatingConfigEntry() *ConsulTerminatingConfigEntry {
	this := ConsulTerminatingConfigEntry{}
	return &this
}

// NewConsulTerminatingConfigEntryWithDefaults instantiates a new ConsulTerminatingConfigEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulTerminatingConfigEntryWithDefaults() *ConsulTerminatingConfigEntry {
	this := ConsulTerminatingConfigEntry{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *ConsulTerminatingConfigEntry) GetServices() []ConsulLinkedService {
	if o == nil || o.Services == nil {
		var ret []ConsulLinkedService
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulTerminatingConfigEntry) GetServicesOk() (*[]ConsulLinkedService, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *ConsulTerminatingConfigEntry) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []ConsulLinkedService and assigns it to the Services field.
func (o *ConsulTerminatingConfigEntry) SetServices(v []ConsulLinkedService) {
	o.Services = &v
}

func (o ConsulTerminatingConfigEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableConsulTerminatingConfigEntry struct {
	value *ConsulTerminatingConfigEntry
	isSet bool
}

func (v NullableConsulTerminatingConfigEntry) Get() *ConsulTerminatingConfigEntry {
	return v.value
}

func (v *NullableConsulTerminatingConfigEntry) Set(val *ConsulTerminatingConfigEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulTerminatingConfigEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulTerminatingConfigEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulTerminatingConfigEntry(val *ConsulTerminatingConfigEntry) *NullableConsulTerminatingConfigEntry {
	return &NullableConsulTerminatingConfigEntry{value: val, isSet: true}
}

func (v NullableConsulTerminatingConfigEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulTerminatingConfigEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


