/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// UserNameType the model 'UserNameType'
type UserNameType string

// List of UserNameType
const (
	USERNAMETYPE_FNAME UserNameType = "USERNAME_TYPE_FNAME"
	USERNAMETYPE_ENS_L1 UserNameType = "USERNAME_TYPE_ENS_L1"
)

// All allowed values of UserNameType enum
var AllowedUserNameTypeEnumValues = []UserNameType{
	"USERNAME_TYPE_FNAME",
	"USERNAME_TYPE_ENS_L1",
}

func (v *UserNameType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserNameType(value)
	for _, existing := range AllowedUserNameTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserNameType", value)
}

// NewUserNameTypeFromValue returns a pointer to a valid UserNameType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserNameTypeFromValue(v string) (*UserNameType, error) {
	ev := UserNameType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserNameType: valid values are %v", v, AllowedUserNameTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserNameType) IsValid() bool {
	for _, existing := range AllowedUserNameTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserNameType value
func (v UserNameType) Ptr() *UserNameType {
	return &v
}

type NullableUserNameType struct {
	value *UserNameType
	isSet bool
}

func (v NullableUserNameType) Get() *UserNameType {
	return v.value
}

func (v *NullableUserNameType) Set(val *UserNameType) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNameType) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNameType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNameType(val *UserNameType) *NullableUserNameType {
	return &NullableUserNameType{value: val, isSet: true}
}

func (v NullableUserNameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNameType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

