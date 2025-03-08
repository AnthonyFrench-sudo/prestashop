/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// IntegrationKeyOut struct for IntegrationKeyOut
type IntegrationKeyOut struct {
	Key string `json:"key"`
}

// NewIntegrationKeyOut instantiates a new IntegrationKeyOut object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationKeyOut(key string) *IntegrationKeyOut {
	this := IntegrationKeyOut{}
	this.Key = key
	return &this
}

// NewIntegrationKeyOutWithDefaults instantiates a new IntegrationKeyOut object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationKeyOutWithDefaults() *IntegrationKeyOut {
	this := IntegrationKeyOut{}
	return &this
}

// GetKey returns the Key field value
func (o *IntegrationKeyOut) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *IntegrationKeyOut) GetKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *IntegrationKeyOut) SetKey(v string) {
	o.Key = v
}

func (o IntegrationKeyOut) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key"] = o.Key
	}
	return json.Marshal(toSerialize)
}

type NullableIntegrationKeyOut struct {
	value *IntegrationKeyOut
	isSet bool
}

func (v NullableIntegrationKeyOut) Get() *IntegrationKeyOut {
	return v.value
}

func (v *NullableIntegrationKeyOut) Set(val *IntegrationKeyOut) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationKeyOut) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationKeyOut) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationKeyOut(val *IntegrationKeyOut) *NullableIntegrationKeyOut {
	return &NullableIntegrationKeyOut{value: val, isSet: true}
}

func (v NullableIntegrationKeyOut) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationKeyOut) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


