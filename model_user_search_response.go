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

// checks if the UserSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSearchResponse{}

// UserSearchResponse struct for UserSearchResponse
type UserSearchResponse struct {
	Result UserSearchResponseResult `json:"result"`
}

type _UserSearchResponse UserSearchResponse

// NewUserSearchResponse instantiates a new UserSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSearchResponse(result UserSearchResponseResult) *UserSearchResponse {
	this := UserSearchResponse{}
	this.Result = result
	return &this
}

// NewUserSearchResponseWithDefaults instantiates a new UserSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSearchResponseWithDefaults() *UserSearchResponse {
	this := UserSearchResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *UserSearchResponse) GetResult() UserSearchResponseResult {
	if o == nil {
		var ret UserSearchResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *UserSearchResponse) GetResultOk() (*UserSearchResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *UserSearchResponse) SetResult(v UserSearchResponseResult) {
	o.Result = v
}

func (o UserSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *UserSearchResponse) UnmarshalJSON(data []byte) (err error) {
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

	varUserSearchResponse := _UserSearchResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserSearchResponse)

	if err != nil {
		return err
	}

	*o = UserSearchResponse(varUserSearchResponse)

	return err
}

type NullableUserSearchResponse struct {
	value *UserSearchResponse
	isSet bool
}

func (v NullableUserSearchResponse) Get() *UserSearchResponse {
	return v.value
}

func (v *NullableUserSearchResponse) Set(val *UserSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSearchResponse(val *UserSearchResponse) *NullableUserSearchResponse {
	return &NullableUserSearchResponse{value: val, isSet: true}
}

func (v NullableUserSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


