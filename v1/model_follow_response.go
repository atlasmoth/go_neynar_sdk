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

// checks if the FollowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponse{}

// FollowResponse struct for FollowResponse
type FollowResponse struct {
	Result *FollowResponseResult `json:"result,omitempty"`
}

// NewFollowResponse instantiates a new FollowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponse() *FollowResponse {
	this := FollowResponse{}
	return &this
}

// NewFollowResponseWithDefaults instantiates a new FollowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseWithDefaults() *FollowResponse {
	this := FollowResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *FollowResponse) GetResult() FollowResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret FollowResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowResponse) GetResultOk() (*FollowResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *FollowResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given FollowResponseResult and assigns it to the Result field.
func (o *FollowResponse) SetResult(v FollowResponseResult) {
	o.Result = &v
}

func (o FollowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableFollowResponse struct {
	value *FollowResponse
	isSet bool
}

func (v NullableFollowResponse) Get() *FollowResponse {
	return v.value
}

func (v *NullableFollowResponse) Set(val *FollowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponse(val *FollowResponse) *NullableFollowResponse {
	return &NullableFollowResponse{value: val, isSet: true}
}

func (v NullableFollowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


