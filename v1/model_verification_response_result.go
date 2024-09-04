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

// checks if the VerificationResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationResponseResult{}

// VerificationResponseResult struct for VerificationResponseResult
type VerificationResponseResult struct {
	Fid *string `json:"fid,omitempty"`
	Username *string `json:"username,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Verifications []string `json:"verifications,omitempty"`
}

// NewVerificationResponseResult instantiates a new VerificationResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationResponseResult() *VerificationResponseResult {
	this := VerificationResponseResult{}
	return &this
}

// NewVerificationResponseResultWithDefaults instantiates a new VerificationResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationResponseResultWithDefaults() *VerificationResponseResult {
	this := VerificationResponseResult{}
	return &this
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *VerificationResponseResult) GetFid() string {
	if o == nil || IsNil(o.Fid) {
		var ret string
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResponseResult) GetFidOk() (*string, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *VerificationResponseResult) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given string and assigns it to the Fid field.
func (o *VerificationResponseResult) SetFid(v string) {
	o.Fid = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *VerificationResponseResult) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResponseResult) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *VerificationResponseResult) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *VerificationResponseResult) SetUsername(v string) {
	o.Username = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *VerificationResponseResult) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResponseResult) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *VerificationResponseResult) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *VerificationResponseResult) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetVerifications returns the Verifications field value if set, zero value otherwise.
func (o *VerificationResponseResult) GetVerifications() []string {
	if o == nil || IsNil(o.Verifications) {
		var ret []string
		return ret
	}
	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerificationResponseResult) GetVerificationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Verifications) {
		return nil, false
	}
	return o.Verifications, true
}

// HasVerifications returns a boolean if a field has been set.
func (o *VerificationResponseResult) HasVerifications() bool {
	if o != nil && !IsNil(o.Verifications) {
		return true
	}

	return false
}

// SetVerifications gets a reference to the given []string and assigns it to the Verifications field.
func (o *VerificationResponseResult) SetVerifications(v []string) {
	o.Verifications = v
}

func (o VerificationResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.Verifications) {
		toSerialize["verifications"] = o.Verifications
	}
	return toSerialize, nil
}

type NullableVerificationResponseResult struct {
	value *VerificationResponseResult
	isSet bool
}

func (v NullableVerificationResponseResult) Get() *VerificationResponseResult {
	return v.value
}

func (v *NullableVerificationResponseResult) Set(val *VerificationResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationResponseResult(val *VerificationResponseResult) *NullableVerificationResponseResult {
	return &NullableVerificationResponseResult{value: val, isSet: true}
}

func (v NullableVerificationResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


