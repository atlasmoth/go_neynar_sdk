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

// checks if the RecentCastsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentCastsResponse{}

// RecentCastsResponse struct for RecentCastsResponse
type RecentCastsResponse struct {
	Result CastsResponseResult `json:"result"`
}

type _RecentCastsResponse RecentCastsResponse

// NewRecentCastsResponse instantiates a new RecentCastsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentCastsResponse(result CastsResponseResult) *RecentCastsResponse {
	this := RecentCastsResponse{}
	this.Result = result
	return &this
}

// NewRecentCastsResponseWithDefaults instantiates a new RecentCastsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentCastsResponseWithDefaults() *RecentCastsResponse {
	this := RecentCastsResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *RecentCastsResponse) GetResult() CastsResponseResult {
	if o == nil {
		var ret CastsResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *RecentCastsResponse) GetResultOk() (*CastsResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *RecentCastsResponse) SetResult(v CastsResponseResult) {
	o.Result = v
}

func (o RecentCastsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentCastsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *RecentCastsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varRecentCastsResponse := _RecentCastsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRecentCastsResponse)

	if err != nil {
		return err
	}

	*o = RecentCastsResponse(varRecentCastsResponse)

	return err
}

type NullableRecentCastsResponse struct {
	value *RecentCastsResponse
	isSet bool
}

func (v NullableRecentCastsResponse) Get() *RecentCastsResponse {
	return v.value
}

func (v *NullableRecentCastsResponse) Set(val *RecentCastsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentCastsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentCastsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentCastsResponse(val *RecentCastsResponse) *NullableRecentCastsResponse {
	return &NullableRecentCastsResponse{value: val, isSet: true}
}

func (v NullableRecentCastsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentCastsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

