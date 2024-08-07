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

// FollowSortType the model 'FollowSortType'
type FollowSortType string

// List of FollowSortType
const (
	FOLLOWSORTTYPE_DESC_CHRON FollowSortType = "desc_chron"
	FOLLOWSORTTYPE_ALGORITHMIC FollowSortType = "algorithmic"
)

// All allowed values of FollowSortType enum
var AllowedFollowSortTypeEnumValues = []FollowSortType{
	"desc_chron",
	"algorithmic",
}

func (v *FollowSortType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FollowSortType(value)
	for _, existing := range AllowedFollowSortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FollowSortType", value)
}

// NewFollowSortTypeFromValue returns a pointer to a valid FollowSortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFollowSortTypeFromValue(v string) (*FollowSortType, error) {
	ev := FollowSortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FollowSortType: valid values are %v", v, AllowedFollowSortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FollowSortType) IsValid() bool {
	for _, existing := range AllowedFollowSortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FollowSortType value
func (v FollowSortType) Ptr() *FollowSortType {
	return &v
}

type NullableFollowSortType struct {
	value *FollowSortType
	isSet bool
}

func (v NullableFollowSortType) Get() *FollowSortType {
	return v.value
}

func (v *NullableFollowSortType) Set(val *FollowSortType) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowSortType) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowSortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowSortType(val *FollowSortType) *NullableFollowSortType {
	return &NullableFollowSortType{value: val, isSet: true}
}

func (v NullableFollowSortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowSortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

