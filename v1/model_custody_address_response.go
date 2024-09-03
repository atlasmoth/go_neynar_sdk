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

// checks if the CustodyAddressResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustodyAddressResponse{}

// CustodyAddressResponse struct for CustodyAddressResponse
type CustodyAddressResponse struct {
	Result CustodyAddressResponseResult `json:"result"`
}

type _CustodyAddressResponse CustodyAddressResponse

// NewCustodyAddressResponse instantiates a new CustodyAddressResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustodyAddressResponse(result CustodyAddressResponseResult) *CustodyAddressResponse {
	this := CustodyAddressResponse{}
	this.Result = result
	return &this
}

// NewCustodyAddressResponseWithDefaults instantiates a new CustodyAddressResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustodyAddressResponseWithDefaults() *CustodyAddressResponse {
	this := CustodyAddressResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *CustodyAddressResponse) GetResult() CustodyAddressResponseResult {
	if o == nil {
		var ret CustodyAddressResponseResult
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CustodyAddressResponse) GetResultOk() (*CustodyAddressResponseResult, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CustodyAddressResponse) SetResult(v CustodyAddressResponseResult) {
	o.Result = v
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
	toSerialize["result"] = o.Result
	return toSerialize, nil
}

func (o *CustodyAddressResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result",
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

	varCustodyAddressResponse := _CustodyAddressResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustodyAddressResponse)

	if err != nil {
		return err
	}

	*o = CustodyAddressResponse(varCustodyAddressResponse)

	return err
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


