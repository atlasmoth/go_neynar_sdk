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

// checks if the UserProfileBio type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfileBio{}

// UserProfileBio struct for UserProfileBio
type UserProfileBio struct {
	Text *string `json:"text,omitempty"`
	MentionedProfiles []string `json:"mentioned_profiles,omitempty"`
}

// NewUserProfileBio instantiates a new UserProfileBio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfileBio() *UserProfileBio {
	this := UserProfileBio{}
	return &this
}

// NewUserProfileBioWithDefaults instantiates a new UserProfileBio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileBioWithDefaults() *UserProfileBio {
	this := UserProfileBio{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *UserProfileBio) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileBio) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *UserProfileBio) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *UserProfileBio) SetText(v string) {
	o.Text = &v
}

// GetMentionedProfiles returns the MentionedProfiles field value if set, zero value otherwise.
func (o *UserProfileBio) GetMentionedProfiles() []string {
	if o == nil || IsNil(o.MentionedProfiles) {
		var ret []string
		return ret
	}
	return o.MentionedProfiles
}

// GetMentionedProfilesOk returns a tuple with the MentionedProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfileBio) GetMentionedProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.MentionedProfiles) {
		return nil, false
	}
	return o.MentionedProfiles, true
}

// HasMentionedProfiles returns a boolean if a field has been set.
func (o *UserProfileBio) HasMentionedProfiles() bool {
	if o != nil && !IsNil(o.MentionedProfiles) {
		return true
	}

	return false
}

// SetMentionedProfiles gets a reference to the given []string and assigns it to the MentionedProfiles field.
func (o *UserProfileBio) SetMentionedProfiles(v []string) {
	o.MentionedProfiles = v
}

func (o UserProfileBio) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfileBio) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.MentionedProfiles) {
		toSerialize["mentioned_profiles"] = o.MentionedProfiles
	}
	return toSerialize, nil
}

type NullableUserProfileBio struct {
	value *UserProfileBio
	isSet bool
}

func (v NullableUserProfileBio) Get() *UserProfileBio {
	return v.value
}

func (v *NullableUserProfileBio) Set(val *UserProfileBio) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfileBio) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfileBio) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfileBio(val *UserProfileBio) *NullableUserProfileBio {
	return &NullableUserProfileBio{value: val, isSet: true}
}

func (v NullableUserProfileBio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfileBio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


