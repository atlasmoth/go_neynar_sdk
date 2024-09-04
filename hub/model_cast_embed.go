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

// checks if the CastEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastEmbed{}

// CastEmbed struct for CastEmbed
type CastEmbed struct {
	CastId *CastId `json:"castId,omitempty"`
}

// NewCastEmbed instantiates a new CastEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastEmbed() *CastEmbed {
	this := CastEmbed{}
	return &this
}

// NewCastEmbedWithDefaults instantiates a new CastEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastEmbedWithDefaults() *CastEmbed {
	this := CastEmbed{}
	return &this
}

// GetCastId returns the CastId field value if set, zero value otherwise.
func (o *CastEmbed) GetCastId() CastId {
	if o == nil || IsNil(o.CastId) {
		var ret CastId
		return ret
	}
	return *o.CastId
}

// GetCastIdOk returns a tuple with the CastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastEmbed) GetCastIdOk() (*CastId, bool) {
	if o == nil || IsNil(o.CastId) {
		return nil, false
	}
	return o.CastId, true
}

// HasCastId returns a boolean if a field has been set.
func (o *CastEmbed) HasCastId() bool {
	if o != nil && !IsNil(o.CastId) {
		return true
	}

	return false
}

// SetCastId gets a reference to the given CastId and assigns it to the CastId field.
func (o *CastEmbed) SetCastId(v CastId) {
	o.CastId = &v
}

func (o CastEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CastId) {
		toSerialize["castId"] = o.CastId
	}
	return toSerialize, nil
}

type NullableCastEmbed struct {
	value *CastEmbed
	isSet bool
}

func (v NullableCastEmbed) Get() *CastEmbed {
	return v.value
}

func (v *NullableCastEmbed) Set(val *CastEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableCastEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableCastEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastEmbed(val *CastEmbed) *NullableCastEmbed {
	return &NullableCastEmbed{value: val, isSet: true}
}

func (v NullableCastEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


