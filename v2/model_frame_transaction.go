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

// checks if the FrameTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameTransaction{}

// FrameTransaction struct for FrameTransaction
type FrameTransaction struct {
	// Transaction hash
	Hash *string `json:"hash,omitempty"`
}

// NewFrameTransaction instantiates a new FrameTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameTransaction() *FrameTransaction {
	this := FrameTransaction{}
	return &this
}

// NewFrameTransactionWithDefaults instantiates a new FrameTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameTransactionWithDefaults() *FrameTransaction {
	this := FrameTransaction{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *FrameTransaction) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameTransaction) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *FrameTransaction) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *FrameTransaction) SetHash(v string) {
	o.Hash = &v
}

func (o FrameTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return toSerialize, nil
}

type NullableFrameTransaction struct {
	value *FrameTransaction
	isSet bool
}

func (v NullableFrameTransaction) Get() *FrameTransaction {
	return v.value
}

func (v *NullableFrameTransaction) Set(val *FrameTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameTransaction(val *FrameTransaction) *NullableFrameTransaction {
	return &NullableFrameTransaction{value: val, isSet: true}
}

func (v NullableFrameTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


