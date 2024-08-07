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

// checks if the ValidatedFrameActionSigner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidatedFrameActionSigner{}

// ValidatedFrameActionSigner struct for ValidatedFrameActionSigner
type ValidatedFrameActionSigner struct {
	Client *User `json:"client,omitempty"`
}

// NewValidatedFrameActionSigner instantiates a new ValidatedFrameActionSigner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatedFrameActionSigner() *ValidatedFrameActionSigner {
	this := ValidatedFrameActionSigner{}
	return &this
}

// NewValidatedFrameActionSignerWithDefaults instantiates a new ValidatedFrameActionSigner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatedFrameActionSignerWithDefaults() *ValidatedFrameActionSigner {
	this := ValidatedFrameActionSigner{}
	return &this
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *ValidatedFrameActionSigner) GetClient() User {
	if o == nil || IsNil(o.Client) {
		var ret User
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedFrameActionSigner) GetClientOk() (*User, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *ValidatedFrameActionSigner) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given User and assigns it to the Client field.
func (o *ValidatedFrameActionSigner) SetClient(v User) {
	o.Client = &v
}

func (o ValidatedFrameActionSigner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidatedFrameActionSigner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	return toSerialize, nil
}

type NullableValidatedFrameActionSigner struct {
	value *ValidatedFrameActionSigner
	isSet bool
}

func (v NullableValidatedFrameActionSigner) Get() *ValidatedFrameActionSigner {
	return v.value
}

func (v *NullableValidatedFrameActionSigner) Set(val *ValidatedFrameActionSigner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatedFrameActionSigner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatedFrameActionSigner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatedFrameActionSigner(val *ValidatedFrameActionSigner) *NullableValidatedFrameActionSigner {
	return &NullableValidatedFrameActionSigner{value: val, isSet: true}
}

func (v NullableValidatedFrameActionSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatedFrameActionSigner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
