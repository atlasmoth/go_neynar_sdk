/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the RegisterDeveloperManagedSignedKeyReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterDeveloperManagedSignedKeyReqBody{}

// RegisterDeveloperManagedSignedKeyReqBody struct for RegisterDeveloperManagedSignedKeyReqBody
type RegisterDeveloperManagedSignedKeyReqBody struct {
	// Ed25519 public key
	PublicKey string `json:"public_key" validate:"regexp=^0x[a-fA-F0-9]{64}$"`
	// Signature generated by the custody address of the app. Signed data includes app_fid, deadline, signer’s public key
	Signature string `json:"signature"`
	// User identifier (unsigned integer)
	AppFid int32 `json:"app_fid"`
	// unix timestamp in seconds that controls how long the signed key request is valid for. (24 hours from now is recommended)
	Deadline int32 `json:"deadline"`
	Sponsor *RegisterSignerKeyReqBodySponsor `json:"sponsor,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegisterDeveloperManagedSignedKeyReqBody RegisterDeveloperManagedSignedKeyReqBody

// NewRegisterDeveloperManagedSignedKeyReqBody instantiates a new RegisterDeveloperManagedSignedKeyReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDeveloperManagedSignedKeyReqBody(publicKey string, signature string, appFid int32, deadline int32) *RegisterDeveloperManagedSignedKeyReqBody {
	this := RegisterDeveloperManagedSignedKeyReqBody{}
	this.PublicKey = publicKey
	this.Signature = signature
	this.AppFid = appFid
	this.Deadline = deadline
	return &this
}

// NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults instantiates a new RegisterDeveloperManagedSignedKeyReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults() *RegisterDeveloperManagedSignedKeyReqBody {
	this := RegisterDeveloperManagedSignedKeyReqBody{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetSignature returns the Signature field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetSignature(v string) {
	o.Signature = v
}

// GetAppFid returns the AppFid field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetAppFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AppFid
}

// GetAppFidOk returns a tuple with the AppFid field value
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetAppFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppFid, true
}

// SetAppFid sets field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetAppFid(v int32) {
	o.AppFid = v
}

// GetDeadline returns the Deadline field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetDeadline() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetDeadlineOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deadline, true
}

// SetDeadline sets field value
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetDeadline(v int32) {
	o.Deadline = v
}

// GetSponsor returns the Sponsor field value if set, zero value otherwise.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSponsor() RegisterSignerKeyReqBodySponsor {
	if o == nil || IsNil(o.Sponsor) {
		var ret RegisterSignerKeyReqBodySponsor
		return ret
	}
	return *o.Sponsor
}

// GetSponsorOk returns a tuple with the Sponsor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSponsorOk() (*RegisterSignerKeyReqBodySponsor, bool) {
	if o == nil || IsNil(o.Sponsor) {
		return nil, false
	}
	return o.Sponsor, true
}

// HasSponsor returns a boolean if a field has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) HasSponsor() bool {
	if o != nil && !IsNil(o.Sponsor) {
		return true
	}

	return false
}

// SetSponsor gets a reference to the given RegisterSignerKeyReqBodySponsor and assigns it to the Sponsor field.
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetSponsor(v RegisterSignerKeyReqBodySponsor) {
	o.Sponsor = &v
}

func (o RegisterDeveloperManagedSignedKeyReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterDeveloperManagedSignedKeyReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["public_key"] = o.PublicKey
	toSerialize["signature"] = o.Signature
	toSerialize["app_fid"] = o.AppFid
	toSerialize["deadline"] = o.Deadline
	if !IsNil(o.Sponsor) {
		toSerialize["sponsor"] = o.Sponsor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegisterDeveloperManagedSignedKeyReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"public_key",
		"signature",
		"app_fid",
		"deadline",
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

	varRegisterDeveloperManagedSignedKeyReqBody := _RegisterDeveloperManagedSignedKeyReqBody{}

	err = json.Unmarshal(data, &varRegisterDeveloperManagedSignedKeyReqBody)

	if err != nil {
		return err
	}

	*o = RegisterDeveloperManagedSignedKeyReqBody(varRegisterDeveloperManagedSignedKeyReqBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "public_key")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "app_fid")
		delete(additionalProperties, "deadline")
		delete(additionalProperties, "sponsor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegisterDeveloperManagedSignedKeyReqBody struct {
	value *RegisterDeveloperManagedSignedKeyReqBody
	isSet bool
}

func (v NullableRegisterDeveloperManagedSignedKeyReqBody) Get() *RegisterDeveloperManagedSignedKeyReqBody {
	return v.value
}

func (v *NullableRegisterDeveloperManagedSignedKeyReqBody) Set(val *RegisterDeveloperManagedSignedKeyReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDeveloperManagedSignedKeyReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDeveloperManagedSignedKeyReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDeveloperManagedSignedKeyReqBody(val *RegisterDeveloperManagedSignedKeyReqBody) *NullableRegisterDeveloperManagedSignedKeyReqBody {
	return &NullableRegisterDeveloperManagedSignedKeyReqBody{value: val, isSet: true}
}

func (v NullableRegisterDeveloperManagedSignedKeyReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDeveloperManagedSignedKeyReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


