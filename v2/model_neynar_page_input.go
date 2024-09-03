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

// checks if the NeynarPageInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarPageInput{}

// NeynarPageInput struct for NeynarPageInput
type NeynarPageInput struct {
	Text *NeynarPageInputText `json:"text,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NeynarPageInput NeynarPageInput

// NewNeynarPageInput instantiates a new NeynarPageInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarPageInput() *NeynarPageInput {
	this := NeynarPageInput{}
	return &this
}

// NewNeynarPageInputWithDefaults instantiates a new NeynarPageInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarPageInputWithDefaults() *NeynarPageInput {
	this := NeynarPageInput{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *NeynarPageInput) GetText() NeynarPageInputText {
	if o == nil || IsNil(o.Text) {
		var ret NeynarPageInputText
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeynarPageInput) GetTextOk() (*NeynarPageInputText, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *NeynarPageInput) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given NeynarPageInputText and assigns it to the Text field.
func (o *NeynarPageInput) SetText(v NeynarPageInputText) {
	o.Text = &v
}

func (o NeynarPageInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarPageInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NeynarPageInput) UnmarshalJSON(data []byte) (err error) {
	varNeynarPageInput := _NeynarPageInput{}

	err = json.Unmarshal(data, &varNeynarPageInput)

	if err != nil {
		return err
	}

	*o = NeynarPageInput(varNeynarPageInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "text")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNeynarPageInput struct {
	value *NeynarPageInput
	isSet bool
}

func (v NullableNeynarPageInput) Get() *NeynarPageInput {
	return v.value
}

func (v *NullableNeynarPageInput) Set(val *NeynarPageInput) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarPageInput) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarPageInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarPageInput(val *NeynarPageInput) *NullableNeynarPageInput {
	return &NullableNeynarPageInput{value: val, isSet: true}
}

func (v NullableNeynarPageInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarPageInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


