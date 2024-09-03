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

// checks if the CastsResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastsResponseResult{}

// CastsResponseResult struct for CastsResponseResult
type CastsResponseResult struct {
	Casts []CastWithInteractions `json:"casts"`
	AdditionalProperties map[string]interface{}
}

type _CastsResponseResult CastsResponseResult

// NewCastsResponseResult instantiates a new CastsResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastsResponseResult(casts []CastWithInteractions) *CastsResponseResult {
	this := CastsResponseResult{}
	this.Casts = casts
	return &this
}

// NewCastsResponseResultWithDefaults instantiates a new CastsResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastsResponseResultWithDefaults() *CastsResponseResult {
	this := CastsResponseResult{}
	return &this
}

// GetCasts returns the Casts field value
func (o *CastsResponseResult) GetCasts() []CastWithInteractions {
	if o == nil {
		var ret []CastWithInteractions
		return ret
	}

	return o.Casts
}

// GetCastsOk returns a tuple with the Casts field value
// and a boolean to check if the value has been set.
func (o *CastsResponseResult) GetCastsOk() ([]CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Casts, true
}

// SetCasts sets field value
func (o *CastsResponseResult) SetCasts(v []CastWithInteractions) {
	o.Casts = v
}

func (o CastsResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastsResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["casts"] = o.Casts

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CastsResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"casts",
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

	varCastsResponseResult := _CastsResponseResult{}

	err = json.Unmarshal(data, &varCastsResponseResult)

	if err != nil {
		return err
	}

	*o = CastsResponseResult(varCastsResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "casts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCastsResponseResult struct {
	value *CastsResponseResult
	isSet bool
}

func (v NullableCastsResponseResult) Get() *CastsResponseResult {
	return v.value
}

func (v *NullableCastsResponseResult) Set(val *CastsResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCastsResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCastsResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastsResponseResult(val *CastsResponseResult) *NullableCastsResponseResult {
	return &NullableCastsResponseResult{value: val, isSet: true}
}

func (v NullableCastsResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastsResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


