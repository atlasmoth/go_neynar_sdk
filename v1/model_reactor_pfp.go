/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ReactorPfp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactorPfp{}

// ReactorPfp struct for ReactorPfp
type ReactorPfp struct {
	// The URL of the reactor's profile picture.
	Url *string `json:"url,omitempty"`
}

// NewReactorPfp instantiates a new ReactorPfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactorPfp() *ReactorPfp {
	this := ReactorPfp{}
	return &this
}

// NewReactorPfpWithDefaults instantiates a new ReactorPfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactorPfpWithDefaults() *ReactorPfp {
	this := ReactorPfp{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ReactorPfp) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactorPfp) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ReactorPfp) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ReactorPfp) SetUrl(v string) {
	o.Url = &v
}

func (o ReactorPfp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactorPfp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableReactorPfp struct {
	value *ReactorPfp
	isSet bool
}

func (v NullableReactorPfp) Get() *ReactorPfp {
	return v.value
}

func (v *NullableReactorPfp) Set(val *ReactorPfp) {
	v.value = val
	v.isSet = true
}

func (v NullableReactorPfp) IsSet() bool {
	return v.isSet
}

func (v *NullableReactorPfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactorPfp(val *ReactorPfp) *NullableReactorPfp {
	return &NullableReactorPfp{value: val, isSet: true}
}

func (v NullableReactorPfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactorPfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


