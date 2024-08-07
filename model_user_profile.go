/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfile{}

// UserProfile struct for UserProfile
type UserProfile struct {
	Bio UserProfileBio `json:"bio"`
}

type _UserProfile UserProfile

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile(bio UserProfileBio) *UserProfile {
	this := UserProfile{}
	this.Bio = bio
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetBio returns the Bio field value
func (o *UserProfile) GetBio() UserProfileBio {
	if o == nil {
		var ret UserProfileBio
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *UserProfile) GetBioOk() (*UserProfileBio, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *UserProfile) SetBio(v UserProfileBio) {
	o.Bio = v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bio"] = o.Bio
	return toSerialize, nil
}

func (o *UserProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bio",
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

	varUserProfile := _UserProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserProfile)

	if err != nil {
		return err
	}

	*o = UserProfile(varUserProfile)

	return err
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


