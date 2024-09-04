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

// checks if the BulkUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUsersResponse{}

// BulkUsersResponse struct for BulkUsersResponse
type BulkUsersResponse struct {
	Users []User `json:"users,omitempty"`
}

// NewBulkUsersResponse instantiates a new BulkUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUsersResponse() *BulkUsersResponse {
	this := BulkUsersResponse{}
	return &this
}

// NewBulkUsersResponseWithDefaults instantiates a new BulkUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUsersResponseWithDefaults() *BulkUsersResponse {
	this := BulkUsersResponse{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *BulkUsersResponse) GetUsers() []User {
	if o == nil || IsNil(o.Users) {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkUsersResponse) GetUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BulkUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *BulkUsersResponse) SetUsers(v []User) {
	o.Users = v
}

func (o BulkUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableBulkUsersResponse struct {
	value *BulkUsersResponse
	isSet bool
}

func (v NullableBulkUsersResponse) Get() *BulkUsersResponse {
	return v.value
}

func (v *NullableBulkUsersResponse) Set(val *BulkUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUsersResponse(val *BulkUsersResponse) *NullableBulkUsersResponse {
	return &NullableBulkUsersResponse{value: val, isSet: true}
}

func (v NullableBulkUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


