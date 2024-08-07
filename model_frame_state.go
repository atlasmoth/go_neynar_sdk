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

// checks if the FrameState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameState{}

// FrameState struct for FrameState
type FrameState struct {
	// State for the frame in a serialized format
	Serialized *string `json:"serialized,omitempty"`
}

// NewFrameState instantiates a new FrameState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameState() *FrameState {
	this := FrameState{}
	return &this
}

// NewFrameStateWithDefaults instantiates a new FrameState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameStateWithDefaults() *FrameState {
	this := FrameState{}
	return &this
}

// GetSerialized returns the Serialized field value if set, zero value otherwise.
func (o *FrameState) GetSerialized() string {
	if o == nil || IsNil(o.Serialized) {
		var ret string
		return ret
	}
	return *o.Serialized
}

// GetSerializedOk returns a tuple with the Serialized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameState) GetSerializedOk() (*string, bool) {
	if o == nil || IsNil(o.Serialized) {
		return nil, false
	}
	return o.Serialized, true
}

// HasSerialized returns a boolean if a field has been set.
func (o *FrameState) HasSerialized() bool {
	if o != nil && !IsNil(o.Serialized) {
		return true
	}

	return false
}

// SetSerialized gets a reference to the given string and assigns it to the Serialized field.
func (o *FrameState) SetSerialized(v string) {
	o.Serialized = &v
}

func (o FrameState) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serialized) {
		toSerialize["serialized"] = o.Serialized
	}
	return toSerialize, nil
}

type NullableFrameState struct {
	value *FrameState
	isSet bool
}

func (v NullableFrameState) Get() *FrameState {
	return v.value
}

func (v *NullableFrameState) Set(val *FrameState) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameState) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameState(val *FrameState) *NullableFrameState {
	return &NullableFrameState{value: val, isSet: true}
}

func (v NullableFrameState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
