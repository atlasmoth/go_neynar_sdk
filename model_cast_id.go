/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CastId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastId{}

// CastId struct for CastId
type CastId struct {
	// User identifier (unsigned integer)
	Fid  int32  `json:"fid"`
	Hash string `json:"hash"`
}

type _CastId CastId

// NewCastId instantiates a new CastId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastId(fid int32, hash string) *CastId {
	this := CastId{}
	this.Fid = fid
	this.Hash = hash
	return &this
}

// NewCastIdWithDefaults instantiates a new CastId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastIdWithDefaults() *CastId {
	this := CastId{}
	return &this
}

// GetFid returns the Fid field value
func (o *CastId) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *CastId) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *CastId) SetFid(v int32) {
	o.Fid = v
}

// GetHash returns the Hash field value
func (o *CastId) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *CastId) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *CastId) SetHash(v string) {
	o.Hash = v
}

func (o CastId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["hash"] = o.Hash
	return toSerialize, nil
}

func (o *CastId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"hash",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCastId := _CastId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastId)

	if err != nil {
		return err
	}

	*o = CastId(varCastId)

	return err
}

type NullableCastId struct {
	value *CastId
	isSet bool
}

func (v NullableCastId) Get() *CastId {
	return v.value
}

func (v *NullableCastId) Set(val *CastId) {
	v.value = val
	v.isSet = true
}

func (v NullableCastId) IsSet() bool {
	return v.isSet
}

func (v *NullableCastId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastId(val *CastId) *NullableCastId {
	return &NullableCastId{value: val, isSet: true}
}

func (v NullableCastId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
