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

// checks if the CastRecasterResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastRecasterResponseResult{}

// CastRecasterResponseResult struct for CastRecasterResponseResult
type CastRecasterResponseResult struct {
	Users []Recaster `json:"users,omitempty"`
	Next *NextCursor `json:"next,omitempty"`
}

// NewCastRecasterResponseResult instantiates a new CastRecasterResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastRecasterResponseResult() *CastRecasterResponseResult {
	this := CastRecasterResponseResult{}
	return &this
}

// NewCastRecasterResponseResultWithDefaults instantiates a new CastRecasterResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastRecasterResponseResultWithDefaults() *CastRecasterResponseResult {
	this := CastRecasterResponseResult{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CastRecasterResponseResult) GetUsers() []Recaster {
	if o == nil || IsNil(o.Users) {
		var ret []Recaster
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastRecasterResponseResult) GetUsersOk() ([]Recaster, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CastRecasterResponseResult) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []Recaster and assigns it to the Users field.
func (o *CastRecasterResponseResult) SetUsers(v []Recaster) {
	o.Users = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *CastRecasterResponseResult) GetNext() NextCursor {
	if o == nil || IsNil(o.Next) {
		var ret NextCursor
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastRecasterResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *CastRecasterResponseResult) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextCursor and assigns it to the Next field.
func (o *CastRecasterResponseResult) SetNext(v NextCursor) {
	o.Next = &v
}

func (o CastRecasterResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastRecasterResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableCastRecasterResponseResult struct {
	value *CastRecasterResponseResult
	isSet bool
}

func (v NullableCastRecasterResponseResult) Get() *CastRecasterResponseResult {
	return v.value
}

func (v *NullableCastRecasterResponseResult) Set(val *CastRecasterResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCastRecasterResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCastRecasterResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastRecasterResponseResult(val *CastRecasterResponseResult) *NullableCastRecasterResponseResult {
	return &NullableCastRecasterResponseResult{value: val, isSet: true}
}

func (v NullableCastRecasterResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastRecasterResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


