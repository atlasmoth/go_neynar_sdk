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

// checks if the GetCastsReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCastsReqBody{}

// GetCastsReqBody struct for GetCastsReqBody
type GetCastsReqBody struct {
	Casts []IndividualHashObj `json:"casts,omitempty"`
}

// NewGetCastsReqBody instantiates a new GetCastsReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCastsReqBody() *GetCastsReqBody {
	this := GetCastsReqBody{}
	return &this
}

// NewGetCastsReqBodyWithDefaults instantiates a new GetCastsReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCastsReqBodyWithDefaults() *GetCastsReqBody {
	this := GetCastsReqBody{}
	return &this
}

// GetCasts returns the Casts field value if set, zero value otherwise.
func (o *GetCastsReqBody) GetCasts() []IndividualHashObj {
	if o == nil || IsNil(o.Casts) {
		var ret []IndividualHashObj
		return ret
	}
	return o.Casts
}

// GetCastsOk returns a tuple with the Casts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCastsReqBody) GetCastsOk() ([]IndividualHashObj, bool) {
	if o == nil || IsNil(o.Casts) {
		return nil, false
	}
	return o.Casts, true
}

// HasCasts returns a boolean if a field has been set.
func (o *GetCastsReqBody) HasCasts() bool {
	if o != nil && !IsNil(o.Casts) {
		return true
	}

	return false
}

// SetCasts gets a reference to the given []IndividualHashObj and assigns it to the Casts field.
func (o *GetCastsReqBody) SetCasts(v []IndividualHashObj) {
	o.Casts = v
}

func (o GetCastsReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCastsReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Casts) {
		toSerialize["casts"] = o.Casts
	}
	return toSerialize, nil
}

type NullableGetCastsReqBody struct {
	value *GetCastsReqBody
	isSet bool
}

func (v NullableGetCastsReqBody) Get() *GetCastsReqBody {
	return v.value
}

func (v *NullableGetCastsReqBody) Set(val *GetCastsReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCastsReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCastsReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCastsReqBody(val *GetCastsReqBody) *NullableGetCastsReqBody {
	return &NullableGetCastsReqBody{value: val, isSet: true}
}

func (v NullableGetCastsReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCastsReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


