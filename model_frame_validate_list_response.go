/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the FrameValidateListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameValidateListResponse{}

// FrameValidateListResponse struct for FrameValidateListResponse
type FrameValidateListResponse struct {
	Frames []string `json:"frames"`
}

type _FrameValidateListResponse FrameValidateListResponse

// NewFrameValidateListResponse instantiates a new FrameValidateListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameValidateListResponse(frames []string) *FrameValidateListResponse {
	this := FrameValidateListResponse{}
	this.Frames = frames
	return &this
}

// NewFrameValidateListResponseWithDefaults instantiates a new FrameValidateListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameValidateListResponseWithDefaults() *FrameValidateListResponse {
	this := FrameValidateListResponse{}
	return &this
}

// GetFrames returns the Frames field value
func (o *FrameValidateListResponse) GetFrames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Frames
}

// GetFramesOk returns a tuple with the Frames field value
// and a boolean to check if the value has been set.
func (o *FrameValidateListResponse) GetFramesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Frames, true
}

// SetFrames sets field value
func (o *FrameValidateListResponse) SetFrames(v []string) {
	o.Frames = v
}

func (o FrameValidateListResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameValidateListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["frames"] = o.Frames
	return toSerialize, nil
}

func (o *FrameValidateListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"frames",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFrameValidateListResponse := _FrameValidateListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameValidateListResponse)

	if err != nil {
		return err
	}

	*o = FrameValidateListResponse(varFrameValidateListResponse)

	return err
}

type NullableFrameValidateListResponse struct {
	value *FrameValidateListResponse
	isSet bool
}

func (v NullableFrameValidateListResponse) Get() *FrameValidateListResponse {
	return v.value
}

func (v *NullableFrameValidateListResponse) Set(val *FrameValidateListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameValidateListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameValidateListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameValidateListResponse(val *FrameValidateListResponse) *NullableFrameValidateListResponse {
	return &NullableFrameValidateListResponse{value: val, isSet: true}
}

func (v NullableFrameValidateListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameValidateListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
