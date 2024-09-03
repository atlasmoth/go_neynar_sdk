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

// checks if the CastsSearchResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastsSearchResponseResult{}

// CastsSearchResponseResult struct for CastsSearchResponseResult
type CastsSearchResponseResult struct {
	Casts []CastWithInteractions `json:"casts"`
	Next NextCursor `json:"next"`
}

type _CastsSearchResponseResult CastsSearchResponseResult

// NewCastsSearchResponseResult instantiates a new CastsSearchResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastsSearchResponseResult(casts []CastWithInteractions, next NextCursor) *CastsSearchResponseResult {
	this := CastsSearchResponseResult{}
	this.Casts = casts
	this.Next = next
	return &this
}

// NewCastsSearchResponseResultWithDefaults instantiates a new CastsSearchResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastsSearchResponseResultWithDefaults() *CastsSearchResponseResult {
	this := CastsSearchResponseResult{}
	return &this
}

// GetCasts returns the Casts field value
func (o *CastsSearchResponseResult) GetCasts() []CastWithInteractions {
	if o == nil {
		var ret []CastWithInteractions
		return ret
	}

	return o.Casts
}

// GetCastsOk returns a tuple with the Casts field value
// and a boolean to check if the value has been set.
func (o *CastsSearchResponseResult) GetCastsOk() ([]CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Casts, true
}

// SetCasts sets field value
func (o *CastsSearchResponseResult) SetCasts(v []CastWithInteractions) {
	o.Casts = v
}

// GetNext returns the Next field value
func (o *CastsSearchResponseResult) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *CastsSearchResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *CastsSearchResponseResult) SetNext(v NextCursor) {
	o.Next = v
}

func (o CastsSearchResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastsSearchResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["casts"] = o.Casts
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *CastsSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"casts",
		"next",
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

	varCastsSearchResponseResult := _CastsSearchResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastsSearchResponseResult)

	if err != nil {
		return err
	}

	*o = CastsSearchResponseResult(varCastsSearchResponseResult)

	return err
}

type NullableCastsSearchResponseResult struct {
	value *CastsSearchResponseResult
	isSet bool
}

func (v NullableCastsSearchResponseResult) Get() *CastsSearchResponseResult {
	return v.value
}

func (v *NullableCastsSearchResponseResult) Set(val *CastsSearchResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCastsSearchResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCastsSearchResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastsSearchResponseResult(val *CastsSearchResponseResult) *NullableCastsSearchResponseResult {
	return &NullableCastsSearchResponseResult{value: val, isSet: true}
}

func (v NullableCastsSearchResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastsSearchResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

