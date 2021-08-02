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

// ConsulExposeConfig struct for ConsulExposeConfig
type ConsulExposeConfig struct {
	Path *[]ConsulExposePath `json:"Path,omitempty"`
}

// NewConsulExposeConfig instantiates a new ConsulExposeConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulExposeConfig() *ConsulExposeConfig {
	this := ConsulExposeConfig{}
	return &this
}

// NewConsulExposeConfigWithDefaults instantiates a new ConsulExposeConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulExposeConfigWithDefaults() *ConsulExposeConfig {
	this := ConsulExposeConfig{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ConsulExposeConfig) GetPath() []ConsulExposePath {
	if o == nil || o.Path == nil {
		var ret []ConsulExposePath
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulExposeConfig) GetPathOk() (*[]ConsulExposePath, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ConsulExposeConfig) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []ConsulExposePath and assigns it to the Path field.
func (o *ConsulExposeConfig) SetPath(v []ConsulExposePath) {
	o.Path = &v
}

func (o ConsulExposeConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableConsulExposeConfig struct {
	value *ConsulExposeConfig
	isSet bool
}

func (v NullableConsulExposeConfig) Get() *ConsulExposeConfig {
	return v.value
}

func (v *NullableConsulExposeConfig) Set(val *ConsulExposeConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulExposeConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulExposeConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulExposeConfig(val *ConsulExposeConfig) *NullableConsulExposeConfig {
	return &NullableConsulExposeConfig{value: val, isSet: true}
}

func (v NullableConsulExposeConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulExposeConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


