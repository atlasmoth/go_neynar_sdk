/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FrameValidateAnalyticsTotalInteractors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameValidateAnalyticsTotalInteractors{}

// FrameValidateAnalyticsTotalInteractors struct for FrameValidateAnalyticsTotalInteractors
type FrameValidateAnalyticsTotalInteractors struct {
	TotalInteractors *float32 `json:"total_interactors,omitempty"`
}

// NewFrameValidateAnalyticsTotalInteractors instantiates a new FrameValidateAnalyticsTotalInteractors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameValidateAnalyticsTotalInteractors() *FrameValidateAnalyticsTotalInteractors {
	this := FrameValidateAnalyticsTotalInteractors{}
	return &this
}

// NewFrameValidateAnalyticsTotalInteractorsWithDefaults instantiates a new FrameValidateAnalyticsTotalInteractors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameValidateAnalyticsTotalInteractorsWithDefaults() *FrameValidateAnalyticsTotalInteractors {
	this := FrameValidateAnalyticsTotalInteractors{}
	return &this
}

// GetTotalInteractors returns the TotalInteractors field value if set, zero value otherwise.
func (o *FrameValidateAnalyticsTotalInteractors) GetTotalInteractors() float32 {
	if o == nil || IsNil(o.TotalInteractors) {
		var ret float32
		return ret
	}
	return *o.TotalInteractors
}

// GetTotalInteractorsOk returns a tuple with the TotalInteractors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameValidateAnalyticsTotalInteractors) GetTotalInteractorsOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalInteractors) {
		return nil, false
	}
	return o.TotalInteractors, true
}

// HasTotalInteractors returns a boolean if a field has been set.
func (o *FrameValidateAnalyticsTotalInteractors) HasTotalInteractors() bool {
	if o != nil && !IsNil(o.TotalInteractors) {
		return true
	}

	return false
}

// SetTotalInteractors gets a reference to the given float32 and assigns it to the TotalInteractors field.
func (o *FrameValidateAnalyticsTotalInteractors) SetTotalInteractors(v float32) {
	o.TotalInteractors = &v
}

func (o FrameValidateAnalyticsTotalInteractors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameValidateAnalyticsTotalInteractors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalInteractors) {
		toSerialize["total_interactors"] = o.TotalInteractors
	}
	return toSerialize, nil
}

type NullableFrameValidateAnalyticsTotalInteractors struct {
	value *FrameValidateAnalyticsTotalInteractors
	isSet bool
}

func (v NullableFrameValidateAnalyticsTotalInteractors) Get() *FrameValidateAnalyticsTotalInteractors {
	return v.value
}

func (v *NullableFrameValidateAnalyticsTotalInteractors) Set(val *FrameValidateAnalyticsTotalInteractors) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameValidateAnalyticsTotalInteractors) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameValidateAnalyticsTotalInteractors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameValidateAnalyticsTotalInteractors(val *FrameValidateAnalyticsTotalInteractors) *NullableFrameValidateAnalyticsTotalInteractors {
	return &NullableFrameValidateAnalyticsTotalInteractors{value: val, isSet: true}
}

func (v NullableFrameValidateAnalyticsTotalInteractors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameValidateAnalyticsTotalInteractors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


