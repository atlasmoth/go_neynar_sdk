/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the MessageCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageCommon{}

// MessageCommon struct for MessageCommon
type MessageCommon struct {
	Hash string `json:"hash" validate:"regexp=^0x[0-9a-fA-F]{40}$"`
	HashScheme HashScheme `json:"hashScheme"`
	Signature string `json:"signature" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	SignatureScheme SignatureScheme `json:"signatureScheme"`
	Signer string `json:"signer" validate:"regexp=^0x[0-9a-fA-F]+$"`
	AdditionalProperties map[string]interface{}
}

type _MessageCommon MessageCommon

// NewMessageCommon instantiates a new MessageCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageCommon(hash string, hashScheme HashScheme, signature string, signatureScheme SignatureScheme, signer string) *MessageCommon {
	this := MessageCommon{}
	this.Hash = hash
	this.HashScheme = hashScheme
	this.Signature = signature
	this.SignatureScheme = signatureScheme
	this.Signer = signer
	return &this
}

// NewMessageCommonWithDefaults instantiates a new MessageCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageCommonWithDefaults() *MessageCommon {
	this := MessageCommon{}
	var hashScheme HashScheme = HASHSCHEME_HASH_SCHEME_BLAKE3
	this.HashScheme = hashScheme
	var signatureScheme SignatureScheme = SIGNATURESCHEME_ED25519
	this.SignatureScheme = signatureScheme
	return &this
}

// GetHash returns the Hash field value
func (o *MessageCommon) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *MessageCommon) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *MessageCommon) SetHash(v string) {
	o.Hash = v
}

// GetHashScheme returns the HashScheme field value
func (o *MessageCommon) GetHashScheme() HashScheme {
	if o == nil {
		var ret HashScheme
		return ret
	}

	return o.HashScheme
}

// GetHashSchemeOk returns a tuple with the HashScheme field value
// and a boolean to check if the value has been set.
func (o *MessageCommon) GetHashSchemeOk() (*HashScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashScheme, true
}

// SetHashScheme sets field value
func (o *MessageCommon) SetHashScheme(v HashScheme) {
	o.HashScheme = v
}

// GetSignature returns the Signature field value
func (o *MessageCommon) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *MessageCommon) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *MessageCommon) SetSignature(v string) {
	o.Signature = v
}

// GetSignatureScheme returns the SignatureScheme field value
func (o *MessageCommon) GetSignatureScheme() SignatureScheme {
	if o == nil {
		var ret SignatureScheme
		return ret
	}

	return o.SignatureScheme
}

// GetSignatureSchemeOk returns a tuple with the SignatureScheme field value
// and a boolean to check if the value has been set.
func (o *MessageCommon) GetSignatureSchemeOk() (*SignatureScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureScheme, true
}

// SetSignatureScheme sets field value
func (o *MessageCommon) SetSignatureScheme(v SignatureScheme) {
	o.SignatureScheme = v
}

// GetSigner returns the Signer field value
func (o *MessageCommon) GetSigner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signer
}

// GetSignerOk returns a tuple with the Signer field value
// and a boolean to check if the value has been set.
func (o *MessageCommon) GetSignerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signer, true
}

// SetSigner sets field value
func (o *MessageCommon) SetSigner(v string) {
	o.Signer = v
}

func (o MessageCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["hashScheme"] = o.HashScheme
	toSerialize["signature"] = o.Signature
	toSerialize["signatureScheme"] = o.SignatureScheme
	toSerialize["signer"] = o.Signer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MessageCommon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"hashScheme",
		"signature",
		"signatureScheme",
		"signer",
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

	varMessageCommon := _MessageCommon{}

	err = json.Unmarshal(data, &varMessageCommon)

	if err != nil {
		return err
	}

	*o = MessageCommon(varMessageCommon)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hash")
		delete(additionalProperties, "hashScheme")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "signatureScheme")
		delete(additionalProperties, "signer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMessageCommon struct {
	value *MessageCommon
	isSet bool
}

func (v NullableMessageCommon) Get() *MessageCommon {
	return v.value
}

func (v *NullableMessageCommon) Set(val *MessageCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageCommon(val *MessageCommon) *NullableMessageCommon {
	return &NullableMessageCommon{value: val, isSet: true}
}

func (v NullableMessageCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


