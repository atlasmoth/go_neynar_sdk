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

// checks if the UserPowerLiteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPowerLiteResponse{}

// UserPowerLiteResponse struct for UserPowerLiteResponse
type UserPowerLiteResponse struct {
	Result *UserPowerLiteResponseResult `json:"result,omitempty"`
}

// NewUserPowerLiteResponse instantiates a new UserPowerLiteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPowerLiteResponse() *UserPowerLiteResponse {
	this := UserPowerLiteResponse{}
	return &this
}

// NewUserPowerLiteResponseWithDefaults instantiates a new UserPowerLiteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPowerLiteResponseWithDefaults() *UserPowerLiteResponse {
	this := UserPowerLiteResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UserPowerLiteResponse) GetResult() UserPowerLiteResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret UserPowerLiteResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPowerLiteResponse) GetResultOk() (*UserPowerLiteResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UserPowerLiteResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given UserPowerLiteResponseResult and assigns it to the Result field.
func (o *UserPowerLiteResponse) SetResult(v UserPowerLiteResponseResult) {
	o.Result = &v
}

func (o UserPowerLiteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPowerLiteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableUserPowerLiteResponse struct {
	value *UserPowerLiteResponse
	isSet bool
}

func (v NullableUserPowerLiteResponse) Get() *UserPowerLiteResponse {
	return v.value
}

func (v *NullableUserPowerLiteResponse) Set(val *UserPowerLiteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPowerLiteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPowerLiteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPowerLiteResponse(val *UserPowerLiteResponse) *NullableUserPowerLiteResponse {
	return &NullableUserPowerLiteResponse{value: val, isSet: true}
}

func (v NullableUserPowerLiteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPowerLiteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


