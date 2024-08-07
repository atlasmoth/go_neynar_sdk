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

// VerificationChainId Chain ID for farcaster verifications. 0 for EOA verifications, 1 or 10 for contract verifications
type VerificationChainId int32

// List of VerificationChainId
const (
	VERIFICATIONCHAINID__0 VerificationChainId = 0
	VERIFICATIONCHAINID__1 VerificationChainId = 1
	VERIFICATIONCHAINID__10 VerificationChainId = 10
)

// All allowed values of VerificationChainId enum
var AllowedVerificationChainIdEnumValues = []VerificationChainId{
	0,
	1,
	10,
}

func (v *VerificationChainId) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VerificationChainId(value)
	for _, existing := range AllowedVerificationChainIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VerificationChainId", value)
}

// NewVerificationChainIdFromValue returns a pointer to a valid VerificationChainId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVerificationChainIdFromValue(v int32) (*VerificationChainId, error) {
	ev := VerificationChainId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VerificationChainId: valid values are %v", v, AllowedVerificationChainIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VerificationChainId) IsValid() bool {
	for _, existing := range AllowedVerificationChainIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VerificationChainId value
func (v VerificationChainId) Ptr() *VerificationChainId {
	return &v
}

type NullableVerificationChainId struct {
	value *VerificationChainId
	isSet bool
}

func (v NullableVerificationChainId) Get() *VerificationChainId {
	return v.value
}

func (v *NullableVerificationChainId) Set(val *VerificationChainId) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationChainId) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationChainId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationChainId(val *VerificationChainId) *NullableVerificationChainId {
	return &NullableVerificationChainId{value: val, isSet: true}
}

func (v NullableVerificationChainId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationChainId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

