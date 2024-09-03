/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CastParentAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastParentAuthor{}

// CastParentAuthor struct for CastParentAuthor
type CastParentAuthor struct {
	Fid NullableString `json:"fid"`
}

type _CastParentAuthor CastParentAuthor

// NewCastParentAuthor instantiates a new CastParentAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastParentAuthor(fid NullableString) *CastParentAuthor {
	this := CastParentAuthor{}
	this.Fid = fid
	return &this
}

// NewCastParentAuthorWithDefaults instantiates a new CastParentAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastParentAuthorWithDefaults() *CastParentAuthor {
	this := CastParentAuthor{}
	return &this
}

// GetFid returns the Fid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastParentAuthor) GetFid() string {
	if o == nil || o.Fid.Get() == nil {
		var ret string
		return ret
	}

	return *o.Fid.Get()
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastParentAuthor) GetFidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fid.Get(), o.Fid.IsSet()
}

// SetFid sets field value
func (o *CastParentAuthor) SetFid(v string) {
	o.Fid.Set(&v)
}

func (o CastParentAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastParentAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid.Get()
	return toSerialize, nil
}

func (o *CastParentAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
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

	varCastParentAuthor := _CastParentAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastParentAuthor)

	if err != nil {
		return err
	}

	*o = CastParentAuthor(varCastParentAuthor)

	return err
}

type NullableCastParentAuthor struct {
	value *CastParentAuthor
	isSet bool
}

func (v NullableCastParentAuthor) Get() *CastParentAuthor {
	return v.value
}

func (v *NullableCastParentAuthor) Set(val *CastParentAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableCastParentAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableCastParentAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastParentAuthor(val *CastParentAuthor) *NullableCastParentAuthor {
	return &NullableCastParentAuthor{value: val, isSet: true}
}

func (v NullableCastParentAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastParentAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


