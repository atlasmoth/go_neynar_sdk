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

// checks if the UserDataAdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataAdd{}

// UserDataAdd struct for UserDataAdd
type UserDataAdd struct {
	Hash *string `json:"hash,omitempty" validate:"regexp=^0x[0-9a-fA-F]{40}$"`
	HashScheme *HashScheme `json:"hashScheme,omitempty"`
	Signature *string `json:"signature,omitempty" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	SignatureScheme *SignatureScheme `json:"signatureScheme,omitempty"`
	Signer *string `json:"signer,omitempty" validate:"regexp=^0x[0-9a-fA-F]+$"`
	Data *UserDataAddAllOfData `json:"data,omitempty"`
}

// NewUserDataAdd instantiates a new UserDataAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataAdd() *UserDataAdd {
	this := UserDataAdd{}
	var hashScheme HashScheme = HASHSCHEME_HASH_SCHEME_BLAKE3
	this.HashScheme = &hashScheme
	var signatureScheme SignatureScheme = SIGNATURESCHEME_ED25519
	this.SignatureScheme = &signatureScheme
	return &this
}

// NewUserDataAddWithDefaults instantiates a new UserDataAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataAddWithDefaults() *UserDataAdd {
	this := UserDataAdd{}
	var hashScheme HashScheme = HASHSCHEME_HASH_SCHEME_BLAKE3
	this.HashScheme = &hashScheme
	var signatureScheme SignatureScheme = SIGNATURESCHEME_ED25519
	this.SignatureScheme = &signatureScheme
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *UserDataAdd) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataAdd) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *UserDataAdd) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *UserDataAdd) SetHash(v string) {
	o.Hash = &v
}

// GetHashScheme returns the HashScheme field value if set, zero value otherwise.
func (o *UserDataAdd) GetHashScheme() HashScheme {
	if o == nil || IsNil(o.HashScheme) {
		var ret HashScheme
		return ret
	}
	return *o.HashScheme
}

// GetHashSchemeOk returns a tuple with the HashScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataAdd) GetHashSchemeOk() (*HashScheme, bool) {
	if o == nil || IsNil(o.HashScheme) {
		return nil, false
	}
	return o.HashScheme, true
}

// HasHashScheme returns a boolean if a field has been set.
func (o *UserDataAdd) HasHashScheme() bool {
	if o != nil && !IsNil(o.HashScheme) {
		return true
	}

	return false
}

// SetHashScheme gets a reference to the given HashScheme and assigns it to the HashScheme field.
func (o *UserDataAdd) SetHashScheme(v HashScheme) {
	o.HashScheme = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *UserDataAdd) GetSignature() string {
	if o == nil || IsNil(o.Signature) {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataAdd) GetSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *UserDataAdd) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *UserDataAdd) SetSignature(v string) {
	o.Signature = &v
}

// GetSignatureScheme returns the SignatureScheme field value if set, zero value otherwise.
func (o *UserDataAdd) GetSignatureScheme() SignatureScheme {
	if o == nil || IsNil(o.SignatureScheme) {
		var ret SignatureScheme
		return ret
	}
	return *o.SignatureScheme
}

// GetSignatureSchemeOk returns a tuple with the SignatureScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataAdd) GetSignatureSchemeOk() (*SignatureScheme, bool) {
	if o == nil || IsNil(o.SignatureScheme) {
		return nil, false
	}
	return o.SignatureScheme, true
}

// HasSignatureScheme returns a boolean if a field has been set.
func (o *UserDataAdd) HasSignatureScheme() bool {
	if o != nil && !IsNil(o.SignatureScheme) {
		return true
	}

	return false
}

// SetSignatureScheme gets a reference to the given SignatureScheme and assigns it to the SignatureScheme field.
func (o *UserDataAdd) SetSignatureScheme(v SignatureScheme) {
	o.SignatureScheme = &v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *UserDataAdd) GetSigner() string {
	if o == nil || IsNil(o.Signer) {
		var ret string
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataAdd) GetSignerOk() (*string, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *UserDataAdd) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given string and assigns it to the Signer field.
func (o *UserDataAdd) SetSigner(v string) {
	o.Signer = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserDataAdd) GetData() UserDataAddAllOfData {
	if o == nil || IsNil(o.Data) {
		var ret UserDataAddAllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataAdd) GetDataOk() (*UserDataAddAllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserDataAdd) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UserDataAddAllOfData and assigns it to the Data field.
func (o *UserDataAdd) SetData(v UserDataAddAllOfData) {
	o.Data = &v
}

func (o UserDataAdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataAdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.HashScheme) {
		toSerialize["hashScheme"] = o.HashScheme
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	if !IsNil(o.SignatureScheme) {
		toSerialize["signatureScheme"] = o.SignatureScheme
	}
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUserDataAdd struct {
	value *UserDataAdd
	isSet bool
}

func (v NullableUserDataAdd) Get() *UserDataAdd {
	return v.value
}

func (v *NullableUserDataAdd) Set(val *UserDataAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataAdd(val *UserDataAdd) *NullableUserDataAdd {
	return &NullableUserDataAdd{value: val, isSet: true}
}

func (v NullableUserDataAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


