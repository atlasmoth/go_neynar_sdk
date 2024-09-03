/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CastRecasterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastRecasterResponse{}

// CastRecasterResponse struct for CastRecasterResponse
type CastRecasterResponse struct {
	Result CastRecasterResponseResult `json:"result"`
	AdditionalProperties map[string]interface{}
}

type _CastRecasterResponse CastRecasterResponse

// NewCastRecasterResponse instantiates a new CastRecasterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastRecasterResponse(result CastRecasterResponseResult) *CastRecasterResponse {
	this := CastRecasterResponse{}
	this.Result = result
	return &this
}

// NewCastRecasterResponseWithDefaults instantiates a new CastRecasterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastRecasterResponseWithDefaults() *CastRecasterResponse {
	this := CastRecasterResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *CastRecasterResponse) GetResult() CastRecasterResponseResult {
	if o == nil {
		var ret CastRecasterResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CastRecasterResponse) GetResultOk() (*CastRecasterResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CastRecasterResponse) SetResult(v CastRecasterResponseResult) {
	o.Result = v
}

func (o CastRecasterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastRecasterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CastRecasterResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varCastRecasterResponse := _CastRecasterResponse{}

	err = json.Unmarshal(data, &varCastRecasterResponse)

	if err != nil {
		return err
	}

	*o = CastRecasterResponse(varCastRecasterResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCastRecasterResponse struct {
	value *CastRecasterResponse
	isSet bool
}

func (v NullableCastRecasterResponse) Get() *CastRecasterResponse {
	return v.value
}

func (v *NullableCastRecasterResponse) Set(val *CastRecasterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCastRecasterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCastRecasterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastRecasterResponse(val *CastRecasterResponse) *NullableCastRecasterResponse {
	return &NullableCastRecasterResponse{value: val, isSet: true}
}

func (v NullableCastRecasterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastRecasterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


