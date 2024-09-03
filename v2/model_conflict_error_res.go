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

// checks if the ConflictErrorRes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConflictErrorRes{}

// ConflictErrorRes Details for the conflict error response
type ConflictErrorRes struct {
	Code *string `json:"code,omitempty"`
	Message string `json:"message"`
	Property *string `json:"property,omitempty"`
	Key *string `json:"key,omitempty"`
}

type _ConflictErrorRes ConflictErrorRes

// NewConflictErrorRes instantiates a new ConflictErrorRes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConflictErrorRes(message string) *ConflictErrorRes {
	this := ConflictErrorRes{}
	this.Message = message
	return &this
}

// NewConflictErrorResWithDefaults instantiates a new ConflictErrorRes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConflictErrorResWithDefaults() *ConflictErrorRes {
	this := ConflictErrorRes{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ConflictErrorRes) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictErrorRes) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ConflictErrorRes) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ConflictErrorRes) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value
func (o *ConflictErrorRes) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ConflictErrorRes) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ConflictErrorRes) SetMessage(v string) {
	o.Message = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ConflictErrorRes) GetProperty() string {
	if o == nil || IsNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictErrorRes) GetPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ConflictErrorRes) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *ConflictErrorRes) SetProperty(v string) {
	o.Property = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ConflictErrorRes) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConflictErrorRes) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ConflictErrorRes) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ConflictErrorRes) SetKey(v string) {
	o.Key = &v
}

func (o ConflictErrorRes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConflictErrorRes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["message"] = o.Message
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

func (o *ConflictErrorRes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varConflictErrorRes := _ConflictErrorRes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConflictErrorRes)

	if err != nil {
		return err
	}

	*o = ConflictErrorRes(varConflictErrorRes)

	return err
}

type NullableConflictErrorRes struct {
	value *ConflictErrorRes
	isSet bool
}

func (v NullableConflictErrorRes) Get() *ConflictErrorRes {
	return v.value
}

func (v *NullableConflictErrorRes) Set(val *ConflictErrorRes) {
	v.value = val
	v.isSet = true
}

func (v NullableConflictErrorRes) IsSet() bool {
	return v.isSet
}

func (v *NullableConflictErrorRes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConflictErrorRes(val *ConflictErrorRes) *NullableConflictErrorRes {
	return &NullableConflictErrorRes{value: val, isSet: true}
}

func (v NullableConflictErrorRes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConflictErrorRes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


