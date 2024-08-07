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

// SubscriptionProviders The provider of the subscription.
type SubscriptionProviders string

// List of SubscriptionProviders
const (
	SUBSCRIPTIONPROVIDERS_FABRIC_STP SubscriptionProviders = "fabric_stp"
	SUBSCRIPTIONPROVIDERS_PARAGRAPH  SubscriptionProviders = "paragraph"
)

// All allowed values of SubscriptionProviders enum
var AllowedSubscriptionProvidersEnumValues = []SubscriptionProviders{
	"fabric_stp",
	"paragraph",
}

func (v *SubscriptionProviders) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionProviders(value)
	for _, existing := range AllowedSubscriptionProvidersEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionProviders", value)
}

// NewSubscriptionProvidersFromValue returns a pointer to a valid SubscriptionProviders
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionProvidersFromValue(v string) (*SubscriptionProviders, error) {
	ev := SubscriptionProviders(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionProviders: valid values are %v", v, AllowedSubscriptionProvidersEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionProviders) IsValid() bool {
	for _, existing := range AllowedSubscriptionProvidersEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionProviders value
func (v SubscriptionProviders) Ptr() *SubscriptionProviders {
	return &v
}

type NullableSubscriptionProviders struct {
	value *SubscriptionProviders
	isSet bool
}

func (v NullableSubscriptionProviders) Get() *SubscriptionProviders {
	return v.value
}

func (v *NullableSubscriptionProviders) Set(val *SubscriptionProviders) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionProviders) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionProviders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionProviders(val *SubscriptionProviders) *NullableSubscriptionProviders {
	return &NullableSubscriptionProviders{value: val, isSet: true}
}

func (v NullableSubscriptionProviders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionProviders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
