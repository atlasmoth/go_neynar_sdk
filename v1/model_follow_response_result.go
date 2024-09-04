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

// checks if the FollowResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponseResult{}

// FollowResponseResult struct for FollowResponseResult
type FollowResponseResult struct {
	Users []FollowResponseUser `json:"users,omitempty"`
	Next *NextCursor `json:"next,omitempty"`
}

// NewFollowResponseResult instantiates a new FollowResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponseResult() *FollowResponseResult {
	this := FollowResponseResult{}
	return &this
}

// NewFollowResponseResultWithDefaults instantiates a new FollowResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseResultWithDefaults() *FollowResponseResult {
	this := FollowResponseResult{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *FollowResponseResult) GetUsers() []FollowResponseUser {
	if o == nil || IsNil(o.Users) {
		var ret []FollowResponseUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowResponseResult) GetUsersOk() ([]FollowResponseUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *FollowResponseResult) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []FollowResponseUser and assigns it to the Users field.
func (o *FollowResponseResult) SetUsers(v []FollowResponseUser) {
	o.Users = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *FollowResponseResult) GetNext() NextCursor {
	if o == nil || IsNil(o.Next) {
		var ret NextCursor
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *FollowResponseResult) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextCursor and assigns it to the Next field.
func (o *FollowResponseResult) SetNext(v NextCursor) {
	o.Next = &v
}

func (o FollowResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableFollowResponseResult struct {
	value *FollowResponseResult
	isSet bool
}

func (v NullableFollowResponseResult) Get() *FollowResponseResult {
	return v.value
}

func (v *NullableFollowResponseResult) Set(val *FollowResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponseResult(val *FollowResponseResult) *NullableFollowResponseResult {
	return &NullableFollowResponseResult{value: val, isSet: true}
}

func (v NullableFollowResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


