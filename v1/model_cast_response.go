/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CastResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastResponse{}

// CastResponse struct for CastResponse
type CastResponse struct {
	Result *CastResponseResult `json:"result,omitempty"`
}

// NewCastResponse instantiates a new CastResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastResponse() *CastResponse {
	this := CastResponse{}
	return &this
}

// NewCastResponseWithDefaults instantiates a new CastResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastResponseWithDefaults() *CastResponse {
	this := CastResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CastResponse) GetResult() CastResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret CastResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastResponse) GetResultOk() (*CastResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CastResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CastResponseResult and assigns it to the Result field.
func (o *CastResponse) SetResult(v CastResponseResult) {
	o.Result = &v
}

func (o CastResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCastResponse struct {
	value *CastResponse
	isSet bool
}

func (v NullableCastResponse) Get() *CastResponse {
	return v.value
}

func (v *NullableCastResponse) Set(val *CastResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCastResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCastResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastResponse(val *CastResponse) *NullableCastResponse {
	return &NullableCastResponse{value: val, isSet: true}
}

func (v NullableCastResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


