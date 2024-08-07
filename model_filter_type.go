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

// FilterType the model 'FilterType'
type FilterType string

// List of FilterType
const (
	FILTERTYPE_FIDS FilterType = "fids"
	FILTERTYPE_PARENT_URL FilterType = "parent_url"
	FILTERTYPE_CHANNEL_ID FilterType = "channel_id"
	FILTERTYPE_EMBED_URL FilterType = "embed_url"
	FILTERTYPE_GLOBAL_TRENDING FilterType = "global_trending"
)

// All allowed values of FilterType enum
var AllowedFilterTypeEnumValues = []FilterType{
	"fids",
	"parent_url",
	"channel_id",
	"embed_url",
	"global_trending",
}

func (v *FilterType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FilterType(value)
	for _, existing := range AllowedFilterTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FilterType", value)
}

// NewFilterTypeFromValue returns a pointer to a valid FilterType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFilterTypeFromValue(v string) (*FilterType, error) {
	ev := FilterType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FilterType: valid values are %v", v, AllowedFilterTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FilterType) IsValid() bool {
	for _, existing := range AllowedFilterTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FilterType value
func (v FilterType) Ptr() *FilterType {
	return &v
}

type NullableFilterType struct {
	value *FilterType
	isSet bool
}

func (v NullableFilterType) Get() *FilterType {
	return v.value
}

func (v *NullableFilterType) Set(val *FilterType) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterType) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterType(val *FilterType) *NullableFilterType {
	return &NullableFilterType{value: val, isSet: true}
}

func (v NullableFilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

