/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the RecentUsersResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentUsersResponseResult{}

// RecentUsersResponseResult struct for RecentUsersResponseResult
type RecentUsersResponseResult struct {
	Users []User `json:"users"`
	Next NextCursor `json:"next"`
	AdditionalProperties map[string]interface{}
}

type _RecentUsersResponseResult RecentUsersResponseResult

// NewRecentUsersResponseResult instantiates a new RecentUsersResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentUsersResponseResult(users []User, next NextCursor) *RecentUsersResponseResult {
	this := RecentUsersResponseResult{}
	this.Users = users
	this.Next = next
	return &this
}

// NewRecentUsersResponseResultWithDefaults instantiates a new RecentUsersResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentUsersResponseResultWithDefaults() *RecentUsersResponseResult {
	this := RecentUsersResponseResult{}
	return &this
}

// GetUsers returns the Users field value
func (o *RecentUsersResponseResult) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *RecentUsersResponseResult) GetUsersOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *RecentUsersResponseResult) SetUsers(v []User) {
	o.Users = v
}

// GetNext returns the Next field value
func (o *RecentUsersResponseResult) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *RecentUsersResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *RecentUsersResponseResult) SetNext(v NextCursor) {
	o.Next = v
}

func (o RecentUsersResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentUsersResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["next"] = o.Next

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecentUsersResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
		"next",
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

	varRecentUsersResponseResult := _RecentUsersResponseResult{}

	err = json.Unmarshal(data, &varRecentUsersResponseResult)

	if err != nil {
		return err
	}

	*o = RecentUsersResponseResult(varRecentUsersResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "users")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecentUsersResponseResult struct {
	value *RecentUsersResponseResult
	isSet bool
}

func (v NullableRecentUsersResponseResult) Get() *RecentUsersResponseResult {
	return v.value
}

func (v *NullableRecentUsersResponseResult) Set(val *RecentUsersResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentUsersResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentUsersResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentUsersResponseResult(val *RecentUsersResponseResult) *NullableRecentUsersResponseResult {
	return &NullableRecentUsersResponseResult{value: val, isSet: true}
}

func (v NullableRecentUsersResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentUsersResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


