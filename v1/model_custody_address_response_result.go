/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the CustodyAddressResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustodyAddressResponseResult{}

// CustodyAddressResponseResult struct for CustodyAddressResponseResult
type CustodyAddressResponseResult struct {
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	CustodyAddress NullableString `json:"custodyAddress"`
	AdditionalProperties map[string]interface{}
}

type _CustodyAddressResponseResult CustodyAddressResponseResult

// NewCustodyAddressResponseResult instantiates a new CustodyAddressResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustodyAddressResponseResult(fid int32, custodyAddress NullableString) *CustodyAddressResponseResult {
	this := CustodyAddressResponseResult{}
	this.Fid = fid
	this.CustodyAddress = custodyAddress
	return &this
}

// NewCustodyAddressResponseResultWithDefaults instantiates a new CustodyAddressResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustodyAddressResponseResultWithDefaults() *CustodyAddressResponseResult {
	this := CustodyAddressResponseResult{}
	var fid int32 = 3
	this.Fid = fid
	return &this
}

// GetFid returns the Fid field value
func (o *CustodyAddressResponseResult) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *CustodyAddressResponseResult) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *CustodyAddressResponseResult) SetFid(v int32) {
	o.Fid = v
}

// GetCustodyAddress returns the CustodyAddress field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CustodyAddressResponseResult) GetCustodyAddress() string {
	if o == nil || o.CustodyAddress.Get() == nil {
		var ret string
		return ret
	}

	return *o.CustodyAddress.Get()
}

// GetCustodyAddressOk returns a tuple with the CustodyAddress field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustodyAddressResponseResult) GetCustodyAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustodyAddress.Get(), o.CustodyAddress.IsSet()
}

// SetCustodyAddress sets field value
func (o *CustodyAddressResponseResult) SetCustodyAddress(v string) {
	o.CustodyAddress.Set(&v)
}

func (o CustodyAddressResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustodyAddressResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["custodyAddress"] = o.CustodyAddress.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustodyAddressResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"custodyAddress",
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

	varCustodyAddressResponseResult := _CustodyAddressResponseResult{}

	err = json.Unmarshal(data, &varCustodyAddressResponseResult)

	if err != nil {
		return err
	}

	*o = CustodyAddressResponseResult(varCustodyAddressResponseResult)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fid")
		delete(additionalProperties, "custodyAddress")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustodyAddressResponseResult struct {
	value *CustodyAddressResponseResult
	isSet bool
}

func (v NullableCustodyAddressResponseResult) Get() *CustodyAddressResponseResult {
	return v.value
}

func (v *NullableCustodyAddressResponseResult) Set(val *CustodyAddressResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustodyAddressResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustodyAddressResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustodyAddressResponseResult(val *CustodyAddressResponseResult) *NullableCustodyAddressResponseResult {
	return &NullableCustodyAddressResponseResult{value: val, isSet: true}
}

func (v NullableCustodyAddressResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustodyAddressResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


