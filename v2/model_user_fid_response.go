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

// checks if the UserFIDResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFIDResponse{}

// UserFIDResponse struct for UserFIDResponse
type UserFIDResponse struct {
	Fid *int32 `json:"fid,omitempty"`
}

// NewUserFIDResponse instantiates a new UserFIDResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFIDResponse() *UserFIDResponse {
	this := UserFIDResponse{}
	return &this
}

// NewUserFIDResponseWithDefaults instantiates a new UserFIDResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFIDResponseWithDefaults() *UserFIDResponse {
	this := UserFIDResponse{}
	return &this
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *UserFIDResponse) GetFid() int32 {
	if o == nil || IsNil(o.Fid) {
		var ret int32
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFIDResponse) GetFidOk() (*int32, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *UserFIDResponse) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given int32 and assigns it to the Fid field.
func (o *UserFIDResponse) SetFid(v int32) {
	o.Fid = &v
}

func (o UserFIDResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFIDResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	return toSerialize, nil
}

type NullableUserFIDResponse struct {
	value *UserFIDResponse
	isSet bool
}

func (v NullableUserFIDResponse) Get() *UserFIDResponse {
	return v.value
}

func (v *NullableUserFIDResponse) Set(val *UserFIDResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFIDResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFIDResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFIDResponse(val *UserFIDResponse) *NullableUserFIDResponse {
	return &NullableUserFIDResponse{value: val, isSet: true}
}

func (v NullableUserFIDResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFIDResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


