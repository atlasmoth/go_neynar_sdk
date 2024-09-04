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

// checks if the CastDehydrated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastDehydrated{}

// CastDehydrated struct for CastDehydrated
type CastDehydrated struct {
	Hash *string `json:"hash,omitempty"`
	Object *string `json:"object,omitempty"`
}

// NewCastDehydrated instantiates a new CastDehydrated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastDehydrated() *CastDehydrated {
	this := CastDehydrated{}
	return &this
}

// NewCastDehydratedWithDefaults instantiates a new CastDehydrated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastDehydratedWithDefaults() *CastDehydrated {
	this := CastDehydrated{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *CastDehydrated) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastDehydrated) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *CastDehydrated) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *CastDehydrated) SetHash(v string) {
	o.Hash = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *CastDehydrated) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastDehydrated) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *CastDehydrated) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *CastDehydrated) SetObject(v string) {
	o.Object = &v
}

func (o CastDehydrated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastDehydrated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	return toSerialize, nil
}

type NullableCastDehydrated struct {
	value *CastDehydrated
	isSet bool
}

func (v NullableCastDehydrated) Get() *CastDehydrated {
	return v.value
}

func (v *NullableCastDehydrated) Set(val *CastDehydrated) {
	v.value = val
	v.isSet = true
}

func (v NullableCastDehydrated) IsSet() bool {
	return v.isSet
}

func (v *NullableCastDehydrated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastDehydrated(val *CastDehydrated) *NullableCastDehydrated {
	return &NullableCastDehydrated{value: val, isSet: true}
}

func (v NullableCastDehydrated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastDehydrated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


