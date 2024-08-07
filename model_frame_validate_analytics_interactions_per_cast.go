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

// checks if the FrameValidateAnalyticsInteractionsPerCast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameValidateAnalyticsInteractionsPerCast{}

// FrameValidateAnalyticsInteractionsPerCast struct for FrameValidateAnalyticsInteractionsPerCast
type FrameValidateAnalyticsInteractionsPerCast struct {
	InteractionsPerCast []FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner `json:"interactions_per_cast"`
}

type _FrameValidateAnalyticsInteractionsPerCast FrameValidateAnalyticsInteractionsPerCast

// NewFrameValidateAnalyticsInteractionsPerCast instantiates a new FrameValidateAnalyticsInteractionsPerCast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameValidateAnalyticsInteractionsPerCast(interactionsPerCast []FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner) *FrameValidateAnalyticsInteractionsPerCast {
	this := FrameValidateAnalyticsInteractionsPerCast{}
	this.InteractionsPerCast = interactionsPerCast
	return &this
}

// NewFrameValidateAnalyticsInteractionsPerCastWithDefaults instantiates a new FrameValidateAnalyticsInteractionsPerCast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameValidateAnalyticsInteractionsPerCastWithDefaults() *FrameValidateAnalyticsInteractionsPerCast {
	this := FrameValidateAnalyticsInteractionsPerCast{}
	return &this
}

// GetInteractionsPerCast returns the InteractionsPerCast field value
func (o *FrameValidateAnalyticsInteractionsPerCast) GetInteractionsPerCast() []FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner {
	if o == nil {
		var ret []FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner
		return ret
	}

	return o.InteractionsPerCast
}

// GetInteractionsPerCastOk returns a tuple with the InteractionsPerCast field value
// and a boolean to check if the value has been set.
func (o *FrameValidateAnalyticsInteractionsPerCast) GetInteractionsPerCastOk() ([]FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.InteractionsPerCast, true
}

// SetInteractionsPerCast sets field value
func (o *FrameValidateAnalyticsInteractionsPerCast) SetInteractionsPerCast(v []FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner) {
	o.InteractionsPerCast = v
}

func (o FrameValidateAnalyticsInteractionsPerCast) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameValidateAnalyticsInteractionsPerCast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interactions_per_cast"] = o.InteractionsPerCast
	return toSerialize, nil
}

func (o *FrameValidateAnalyticsInteractionsPerCast) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interactions_per_cast",
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

	varFrameValidateAnalyticsInteractionsPerCast := _FrameValidateAnalyticsInteractionsPerCast{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameValidateAnalyticsInteractionsPerCast)

	if err != nil {
		return err
	}

	*o = FrameValidateAnalyticsInteractionsPerCast(varFrameValidateAnalyticsInteractionsPerCast)

	return err
}

type NullableFrameValidateAnalyticsInteractionsPerCast struct {
	value *FrameValidateAnalyticsInteractionsPerCast
	isSet bool
}

func (v NullableFrameValidateAnalyticsInteractionsPerCast) Get() *FrameValidateAnalyticsInteractionsPerCast {
	return v.value
}

func (v *NullableFrameValidateAnalyticsInteractionsPerCast) Set(val *FrameValidateAnalyticsInteractionsPerCast) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameValidateAnalyticsInteractionsPerCast) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameValidateAnalyticsInteractionsPerCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameValidateAnalyticsInteractionsPerCast(val *FrameValidateAnalyticsInteractionsPerCast) *NullableFrameValidateAnalyticsInteractionsPerCast {
	return &NullableFrameValidateAnalyticsInteractionsPerCast{value: val, isSet: true}
}

func (v NullableFrameValidateAnalyticsInteractionsPerCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameValidateAnalyticsInteractionsPerCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
