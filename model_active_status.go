/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ActiveStatus The status of a user.   - active: The user is currently active.   - inactive: The user is not currently active. 
type ActiveStatus string

// List of ActiveStatus
const (
	ACTIVESTATUS_ACTIVE ActiveStatus = "active"
	ACTIVESTATUS_INACTIVE ActiveStatus = "inactive"
)

// All allowed values of ActiveStatus enum
var AllowedActiveStatusEnumValues = []ActiveStatus{
	"active",
	"inactive",
}

func (v *ActiveStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActiveStatus(value)
	for _, existing := range AllowedActiveStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActiveStatus", value)
}

// NewActiveStatusFromValue returns a pointer to a valid ActiveStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActiveStatusFromValue(v string) (*ActiveStatus, error) {
	ev := ActiveStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActiveStatus: valid values are %v", v, AllowedActiveStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActiveStatus) IsValid() bool {
	for _, existing := range AllowedActiveStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActiveStatus value
func (v ActiveStatus) Ptr() *ActiveStatus {
	return &v
}

type NullableActiveStatus struct {
	value *ActiveStatus
	isSet bool
}

func (v NullableActiveStatus) Get() *ActiveStatus {
	return v.value
}

func (v *NullableActiveStatus) Set(val *ActiveStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveStatus(val *ActiveStatus) *NullableActiveStatus {
	return &NullableActiveStatus{value: val, isSet: true}
}

func (v NullableActiveStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

