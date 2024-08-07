/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FrameValidateAnalyticsInputTextInputTextsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameValidateAnalyticsInputTextInputTextsInner{}

// FrameValidateAnalyticsInputTextInputTextsInner struct for FrameValidateAnalyticsInputTextInputTextsInner
type FrameValidateAnalyticsInputTextInputTextsInner struct {
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	Username string `json:"username"`
	InputText string `json:"input_text"`
}

type _FrameValidateAnalyticsInputTextInputTextsInner FrameValidateAnalyticsInputTextInputTextsInner

// NewFrameValidateAnalyticsInputTextInputTextsInner instantiates a new FrameValidateAnalyticsInputTextInputTextsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameValidateAnalyticsInputTextInputTextsInner(fid int32, username string, inputText string) *FrameValidateAnalyticsInputTextInputTextsInner {
	this := FrameValidateAnalyticsInputTextInputTextsInner{}
	this.Fid = fid
	this.Username = username
	this.InputText = inputText
	return &this
}

// NewFrameValidateAnalyticsInputTextInputTextsInnerWithDefaults instantiates a new FrameValidateAnalyticsInputTextInputTextsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameValidateAnalyticsInputTextInputTextsInnerWithDefaults() *FrameValidateAnalyticsInputTextInputTextsInner {
	this := FrameValidateAnalyticsInputTextInputTextsInner{}
	return &this
}

// GetFid returns the Fid field value
func (o *FrameValidateAnalyticsInputTextInputTextsInner) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *FrameValidateAnalyticsInputTextInputTextsInner) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *FrameValidateAnalyticsInputTextInputTextsInner) SetFid(v int32) {
	o.Fid = v
}

// GetUsername returns the Username field value
func (o *FrameValidateAnalyticsInputTextInputTextsInner) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *FrameValidateAnalyticsInputTextInputTextsInner) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *FrameValidateAnalyticsInputTextInputTextsInner) SetUsername(v string) {
	o.Username = v
}

// GetInputText returns the InputText field value
func (o *FrameValidateAnalyticsInputTextInputTextsInner) GetInputText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InputText
}

// GetInputTextOk returns a tuple with the InputText field value
// and a boolean to check if the value has been set.
func (o *FrameValidateAnalyticsInputTextInputTextsInner) GetInputTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InputText, true
}

// SetInputText sets field value
func (o *FrameValidateAnalyticsInputTextInputTextsInner) SetInputText(v string) {
	o.InputText = v
}

func (o FrameValidateAnalyticsInputTextInputTextsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameValidateAnalyticsInputTextInputTextsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	toSerialize["input_text"] = o.InputText
	return toSerialize, nil
}

func (o *FrameValidateAnalyticsInputTextInputTextsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"username",
		"input_text",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFrameValidateAnalyticsInputTextInputTextsInner := _FrameValidateAnalyticsInputTextInputTextsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameValidateAnalyticsInputTextInputTextsInner)

	if err != nil {
		return err
	}

	*o = FrameValidateAnalyticsInputTextInputTextsInner(varFrameValidateAnalyticsInputTextInputTextsInner)

	return err
}

type NullableFrameValidateAnalyticsInputTextInputTextsInner struct {
	value *FrameValidateAnalyticsInputTextInputTextsInner
	isSet bool
}

func (v NullableFrameValidateAnalyticsInputTextInputTextsInner) Get() *FrameValidateAnalyticsInputTextInputTextsInner {
	return v.value
}

func (v *NullableFrameValidateAnalyticsInputTextInputTextsInner) Set(val *FrameValidateAnalyticsInputTextInputTextsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameValidateAnalyticsInputTextInputTextsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameValidateAnalyticsInputTextInputTextsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameValidateAnalyticsInputTextInputTextsInner(val *FrameValidateAnalyticsInputTextInputTextsInner) *NullableFrameValidateAnalyticsInputTextInputTextsInner {
	return &NullableFrameValidateAnalyticsInputTextInputTextsInner{value: val, isSet: true}
}

func (v NullableFrameValidateAnalyticsInputTextInputTextsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameValidateAnalyticsInputTextInputTextsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


