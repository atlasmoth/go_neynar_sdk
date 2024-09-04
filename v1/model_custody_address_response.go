/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CustodyAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustodyAddressResponse{}

// CustodyAddressResponse struct for CustodyAddressResponse
type CustodyAddressResponse struct {
	Result *CustodyAddressResponseResult `json:"result,omitempty"`
}

// NewCustodyAddressResponse instantiates a new CustodyAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustodyAddressResponse() *CustodyAddressResponse {
	this := CustodyAddressResponse{}
	return &this
}

// NewCustodyAddressResponseWithDefaults instantiates a new CustodyAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustodyAddressResponseWithDefaults() *CustodyAddressResponse {
	this := CustodyAddressResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CustodyAddressResponse) GetResult() CustodyAddressResponseResult {
	if o == nil || IsNil(o.Result) {
		var ret CustodyAddressResponseResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustodyAddressResponse) GetResultOk() (*CustodyAddressResponseResult, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CustodyAddressResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given CustodyAddressResponseResult and assigns it to the Result field.
func (o *CustodyAddressResponse) SetResult(v CustodyAddressResponseResult) {
	o.Result = &v
}

func (o CustodyAddressResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustodyAddressResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCustodyAddressResponse struct {
	value *CustodyAddressResponse
	isSet bool
}

func (v NullableCustodyAddressResponse) Get() *CustodyAddressResponse {
	return v.value
}

func (v *NullableCustodyAddressResponse) Set(val *CustodyAddressResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustodyAddressResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustodyAddressResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustodyAddressResponse(val *CustodyAddressResponse) *NullableCustodyAddressResponse {
	return &NullableCustodyAddressResponse{value: val, isSet: true}
}

func (v NullableCustodyAddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustodyAddressResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


