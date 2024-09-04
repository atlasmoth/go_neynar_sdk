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

// checks if the RegisterDeveloperManagedSignedKeyReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterDeveloperManagedSignedKeyReqBody{}

// RegisterDeveloperManagedSignedKeyReqBody struct for RegisterDeveloperManagedSignedKeyReqBody
type RegisterDeveloperManagedSignedKeyReqBody struct {
	// Ed25519 public key
	PublicKey *string `json:"public_key,omitempty" validate:"regexp=^0x[a-fA-F0-9]{64}$"`
	// Signature generated by the custody address of the app. Signed data includes app_fid, deadline, signer’s public key
	Signature *string `json:"signature,omitempty"`
	// User identifier (unsigned integer)
	AppFid *int32 `json:"app_fid,omitempty"`
	// unix timestamp in seconds that controls how long the signed key request is valid for. (24 hours from now is recommended)
	Deadline *int32 `json:"deadline,omitempty"`
	Sponsor *RegisterSignerKeyReqBodySponsor `json:"sponsor,omitempty"`
}

// NewRegisterDeveloperManagedSignedKeyReqBody instantiates a new RegisterDeveloperManagedSignedKeyReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDeveloperManagedSignedKeyReqBody() *RegisterDeveloperManagedSignedKeyReqBody {
	this := RegisterDeveloperManagedSignedKeyReqBody{}
	return &this
}

// NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults instantiates a new RegisterDeveloperManagedSignedKeyReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults() *RegisterDeveloperManagedSignedKeyReqBody {
	this := RegisterDeveloperManagedSignedKeyReqBody{}
	return &this
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetSignature(v string) {
	o.Signature = &v
}

// GetAppFid returns the AppFid field value if set, zero value otherwise.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetAppFid() int32 {
	if o == nil || IsNil(o.AppFid) {
		var ret int32
		return ret
	}
	return *o.AppFid
}

// GetAppFidOk returns a tuple with the AppFid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetAppFidOk() (*int32, bool) {
	if o == nil || IsNil(o.AppFid) {
		return nil, false
	}
	return o.AppFid, true
}

// HasAppFid returns a boolean if a field has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) HasAppFid() bool {
	if o != nil && !IsNil(o.AppFid) {
		return true
	}

	return false
}

// SetAppFid gets a reference to the given int32 and assigns it to the AppFid field.
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetAppFid(v int32) {
	o.AppFid = &v
}

// GetDeadline returns the Deadline field value if set, zero value otherwise.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetDeadline() int32 {
	if o == nil || IsNil(o.Deadline) {
		var ret int32
		return ret
	}
	return *o.Deadline
}

// GetDeadlineOk returns a tuple with the Deadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) GetDeadlineOk() (*int32, bool) {
	if o == nil || IsNil(o.Deadline) {
		return nil, false
	}
	return o.Deadline, true
}

// HasDeadline returns a boolean if a field has been set.
func (o *RegisterDeveloperManagedSignedKeyReqBody) HasDeadline() bool {
	if o != nil && !IsNil(o.Deadline) {
		return true
	}

	return false
}

// SetDeadline gets a reference to the given int32 and assigns it to the Deadline field.
func (o *RegisterDeveloperManagedSignedKeyReqBody) SetDeadline(v int32) {
	o.Deadline = &v
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
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.AppFid) {
		toSerialize["app_fid"] = o.AppFid
	}
	if !IsNil(o.Deadline) {
		toSerialize["deadline"] = o.Deadline
	}
	if !IsNil(o.Sponsor) {
		toSerialize["sponsor"] = o.Sponsor
	}
	return toSerialize, nil
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


