/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the VerificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationResponse{}

// VerificationResponse struct for VerificationResponse
type VerificationResponse struct {
	Result VerificationResponseResult `json:"result"`
}

type _VerificationResponse VerificationResponse

// NewVerificationResponse instantiates a new VerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationResponse(result VerificationResponseResult) *VerificationResponse {
	this := VerificationResponse{}
	this.Result = result
	return &this
}

// NewVerificationResponseWithDefaults instantiates a new VerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationResponseWithDefaults() *VerificationResponse {
	this := VerificationResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *VerificationResponse) GetResult() VerificationResponseResult {
	if o == nil {
		var ret VerificationResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *VerificationResponse) GetResultOk() (*VerificationResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *VerificationResponse) SetResult(v VerificationResponseResult) {
	o.Result = v
}

func (o VerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *VerificationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varVerificationResponse := _VerificationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerificationResponse)

	if err != nil {
		return err
	}

	*o = VerificationResponse(varVerificationResponse)

	return err
}

type NullableVerificationResponse struct {
	value *VerificationResponse
	isSet bool
}

func (v NullableVerificationResponse) Get() *VerificationResponse {
	return v.value
}

func (v *NullableVerificationResponse) Set(val *VerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationResponse(val *VerificationResponse) *NullableVerificationResponse {
	return &NullableVerificationResponse{value: val, isSet: true}
}

func (v NullableVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


