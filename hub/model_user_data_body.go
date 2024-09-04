/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserDataBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataBody{}

// UserDataBody Adds metadata about a user
type UserDataBody struct {
	Type *UserDataType `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewUserDataBody instantiates a new UserDataBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataBody() *UserDataBody {
	this := UserDataBody{}
	var type_ UserDataType = USERDATATYPE_PFP
	this.Type = &type_
	return &this
}

// NewUserDataBodyWithDefaults instantiates a new UserDataBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataBodyWithDefaults() *UserDataBody {
	this := UserDataBody{}
	var type_ UserDataType = USERDATATYPE_PFP
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserDataBody) GetType() UserDataType {
	if o == nil || IsNil(o.Type) {
		var ret UserDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataBody) GetTypeOk() (*UserDataType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserDataBody) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given UserDataType and assigns it to the Type field.
func (o *UserDataBody) SetType(v UserDataType) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserDataBody) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataBody) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserDataBody) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserDataBody) SetValue(v string) {
	o.Value = &v
}

func (o UserDataBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableUserDataBody struct {
	value *UserDataBody
	isSet bool
}

func (v NullableUserDataBody) Get() *UserDataBody {
	return v.value
}

func (v *NullableUserDataBody) Set(val *UserDataBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataBody(val *UserDataBody) *NullableUserDataBody {
	return &NullableUserDataBody{value: val, isSet: true}
}

func (v NullableUserDataBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


