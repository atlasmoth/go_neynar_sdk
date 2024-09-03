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

// checks if the CastLikesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastLikesResponse{}

// CastLikesResponse struct for CastLikesResponse
type CastLikesResponse struct {
	Result CastLikesResponseResult `json:"result"`
}

type _CastLikesResponse CastLikesResponse

// NewCastLikesResponse instantiates a new CastLikesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastLikesResponse(result CastLikesResponseResult) *CastLikesResponse {
	this := CastLikesResponse{}
	this.Result = result
	return &this
}

// NewCastLikesResponseWithDefaults instantiates a new CastLikesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastLikesResponseWithDefaults() *CastLikesResponse {
	this := CastLikesResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *CastLikesResponse) GetResult() CastLikesResponseResult {
	if o == nil {
		var ret CastLikesResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CastLikesResponse) GetResultOk() (*CastLikesResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CastLikesResponse) SetResult(v CastLikesResponseResult) {
	o.Result = v
}

func (o CastLikesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastLikesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CastLikesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCastLikesResponse := _CastLikesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastLikesResponse)

	if err != nil {
		return err
	}

	*o = CastLikesResponse(varCastLikesResponse)

	return err
}

type NullableCastLikesResponse struct {
	value *CastLikesResponse
	isSet bool
}

func (v NullableCastLikesResponse) Get() *CastLikesResponse {
	return v.value
}

func (v *NullableCastLikesResponse) Set(val *CastLikesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCastLikesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCastLikesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastLikesResponse(val *CastLikesResponse) *NullableCastLikesResponse {
	return &NullableCastLikesResponse{value: val, isSet: true}
}

func (v NullableCastLikesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastLikesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


