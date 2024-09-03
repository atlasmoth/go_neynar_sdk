/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ValidatedFrameActionTappedButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidatedFrameActionTappedButton{}

// ValidatedFrameActionTappedButton struct for ValidatedFrameActionTappedButton
type ValidatedFrameActionTappedButton struct {
	Index int32 `json:"index"`
	AdditionalProperties map[string]interface{}
}

type _ValidatedFrameActionTappedButton ValidatedFrameActionTappedButton

// NewValidatedFrameActionTappedButton instantiates a new ValidatedFrameActionTappedButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatedFrameActionTappedButton(index int32) *ValidatedFrameActionTappedButton {
	this := ValidatedFrameActionTappedButton{}
	this.Index = index
	return &this
}

// NewValidatedFrameActionTappedButtonWithDefaults instantiates a new ValidatedFrameActionTappedButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatedFrameActionTappedButtonWithDefaults() *ValidatedFrameActionTappedButton {
	this := ValidatedFrameActionTappedButton{}
	return &this
}

// GetIndex returns the Index field value
func (o *ValidatedFrameActionTappedButton) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameActionTappedButton) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *ValidatedFrameActionTappedButton) SetIndex(v int32) {
	o.Index = v
}

func (o ValidatedFrameActionTappedButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidatedFrameActionTappedButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidatedFrameActionTappedButton) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"index",
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

	varValidatedFrameActionTappedButton := _ValidatedFrameActionTappedButton{}

	err = json.Unmarshal(data, &varValidatedFrameActionTappedButton)

	if err != nil {
		return err
	}

	*o = ValidatedFrameActionTappedButton(varValidatedFrameActionTappedButton)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "index")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidatedFrameActionTappedButton struct {
	value *ValidatedFrameActionTappedButton
	isSet bool
}

func (v NullableValidatedFrameActionTappedButton) Get() *ValidatedFrameActionTappedButton {
	return v.value
}

func (v *NullableValidatedFrameActionTappedButton) Set(val *ValidatedFrameActionTappedButton) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatedFrameActionTappedButton) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatedFrameActionTappedButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatedFrameActionTappedButton(val *ValidatedFrameActionTappedButton) *NullableValidatedFrameActionTappedButton {
	return &NullableValidatedFrameActionTappedButton{value: val, isSet: true}
}

func (v NullableValidatedFrameActionTappedButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatedFrameActionTappedButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


