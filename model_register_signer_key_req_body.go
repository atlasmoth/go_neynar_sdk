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

// checks if the RegisterSignerKeyReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterSignerKeyReqBody{}

// RegisterSignerKeyReqBody struct for RegisterSignerKeyReqBody
type RegisterSignerKeyReqBody struct {
	// UUID of the signer
	SignerUuid string `json:"signer_uuid"`
	// Signature generated by the custody address of the app. Signed data includes app_fid, deadline, signer’s public key
	Signature string `json:"signature"`
	// User identifier (unsigned integer)
	AppFid int32 `json:"app_fid"`
	// unix timestamp in seconds that controls how long the signed key request is valid for. (24 hours from now is recommended)
	Deadline int32 `json:"deadline"`
}

type _RegisterSignerKeyReqBody RegisterSignerKeyReqBody

// NewRegisterSignerKeyReqBody instantiates a new RegisterSignerKeyReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterSignerKeyReqBody(signerUuid string, signature string, appFid int32, deadline int32) *RegisterSignerKeyReqBody {
	this := RegisterSignerKeyReqBody{}
	this.SignerUuid = signerUuid
	this.Signature = signature
	this.AppFid = appFid
	this.Deadline = deadline
	return &this
}

// NewRegisterSignerKeyReqBodyWithDefaults instantiates a new RegisterSignerKeyReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterSignerKeyReqBodyWithDefaults() *RegisterSignerKeyReqBody {
	this := RegisterSignerKeyReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value
func (o *RegisterSignerKeyReqBody) GetSignerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value
// and a boolean to check if the value has been set.
func (o *RegisterSignerKeyReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerUuid, true
}

// SetSignerUuid sets field value
func (o *RegisterSignerKeyReqBody) SetSignerUuid(v string) {
	o.SignerUuid = v
}

// GetSignature returns the Signature field value
func (o *RegisterSignerKeyReqBody) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *RegisterSignerKeyReqBody) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *RegisterSignerKeyReqBody) SetSignature(v string) {
	o.Signature = v
}

// GetAppFid returns the AppFid field value
func (o *RegisterSignerKeyReqBody) GetAppFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppFid
}

// GetAppFidOk returns a tuple with the AppFid field value
// and a boolean to check if the value has been set.
func (o *RegisterSignerKeyReqBody) GetAppFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppFid, true
}

// SetAppFid sets field value
func (o *RegisterSignerKeyReqBody) SetAppFid(v int32) {
	o.AppFid = v
}

// GetDeadline returns the Deadline field value
func (o *RegisterSignerKeyReqBody) GetDeadline() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value
// and a boolean to check if the value has been set.
func (o *RegisterSignerKeyReqBody) GetDeadlineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadline, true
}

// SetDeadline sets field value
func (o *RegisterSignerKeyReqBody) SetDeadline(v int32) {
	o.Deadline = v
}

func (o RegisterSignerKeyReqBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterSignerKeyReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signer_uuid"] = o.SignerUuid
	toSerialize["signature"] = o.Signature
	toSerialize["app_fid"] = o.AppFid
	toSerialize["deadline"] = o.Deadline
	return toSerialize, nil
}

func (o *RegisterSignerKeyReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signer_uuid",
		"signature",
		"app_fid",
		"deadline",
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

	varRegisterSignerKeyReqBody := _RegisterSignerKeyReqBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterSignerKeyReqBody)

	if err != nil {
		return err
	}

	*o = RegisterSignerKeyReqBody(varRegisterSignerKeyReqBody)

	return err
}

type NullableRegisterSignerKeyReqBody struct {
	value *RegisterSignerKeyReqBody
	isSet bool
}

func (v NullableRegisterSignerKeyReqBody) Get() *RegisterSignerKeyReqBody {
	return v.value
}

func (v *NullableRegisterSignerKeyReqBody) Set(val *RegisterSignerKeyReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterSignerKeyReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterSignerKeyReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterSignerKeyReqBody(val *RegisterSignerKeyReqBody) *NullableRegisterSignerKeyReqBody {
	return &NullableRegisterSignerKeyReqBody{value: val, isSet: true}
}

func (v NullableRegisterSignerKeyReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterSignerKeyReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
