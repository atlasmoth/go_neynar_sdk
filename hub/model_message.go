/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message *  A Message is a delta operation on the Farcaster network. The message protobuf is an envelope  that wraps a MessageData object and contains a hash and signature which can verify its authenticity.
type Message struct {
	Hash string `json:"hash" validate:"regexp=^0x[0-9a-fA-F]{40}$"`
	HashScheme HashScheme `json:"hashScheme"`
	Signature string `json:"signature" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	SignatureScheme SignatureScheme `json:"signatureScheme"`
	Signer string `json:"signer" validate:"regexp=^0x[0-9a-fA-F]+$"`
	Data MessageAllOfData `json:"data"`
}

type _Message Message

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage(hash string, hashScheme HashScheme, signature string, signatureScheme SignatureScheme, signer string, data MessageAllOfData) *Message {
	this := Message{}
	this.Hash = hash
	this.HashScheme = hashScheme
	this.Signature = signature
	this.SignatureScheme = signatureScheme
	this.Signer = signer
	this.Data = data
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	var hashScheme HashScheme = HASHSCHEME_HASH_SCHEME_BLAKE3
	this.HashScheme = hashScheme
	var signatureScheme SignatureScheme = SIGNATURESCHEME_ED25519
	this.SignatureScheme = signatureScheme
	return &this
}

// GetHash returns the Hash field value
func (o *Message) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Message) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Message) SetHash(v string) {
	o.Hash = v
}

// GetHashScheme returns the HashScheme field value
func (o *Message) GetHashScheme() HashScheme {
	if o == nil {
		var ret HashScheme
		return ret
	}

	return o.HashScheme
}

// GetHashSchemeOk returns a tuple with the HashScheme field value
// and a boolean to check if the value has been set.
func (o *Message) GetHashSchemeOk() (*HashScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HashScheme, true
}

// SetHashScheme sets field value
func (o *Message) SetHashScheme(v HashScheme) {
	o.HashScheme = v
}

// GetSignature returns the Signature field value
func (o *Message) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *Message) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *Message) SetSignature(v string) {
	o.Signature = v
}

// GetSignatureScheme returns the SignatureScheme field value
func (o *Message) GetSignatureScheme() SignatureScheme {
	if o == nil {
		var ret SignatureScheme
		return ret
	}

	return o.SignatureScheme
}

// GetSignatureSchemeOk returns a tuple with the SignatureScheme field value
// and a boolean to check if the value has been set.
func (o *Message) GetSignatureSchemeOk() (*SignatureScheme, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignatureScheme, true
}

// SetSignatureScheme sets field value
func (o *Message) SetSignatureScheme(v SignatureScheme) {
	o.SignatureScheme = v
}

// GetSigner returns the Signer field value
func (o *Message) GetSigner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signer
}

// GetSignerOk returns a tuple with the Signer field value
// and a boolean to check if the value has been set.
func (o *Message) GetSignerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signer, true
}

// SetSigner sets field value
func (o *Message) SetSigner(v string) {
	o.Signer = v
}

// GetData returns the Data field value
func (o *Message) GetData() MessageAllOfData {
	if o == nil {
		var ret MessageAllOfData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Message) GetDataOk() (*MessageAllOfData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Message) SetData(v MessageAllOfData) {
	o.Data = v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["hashScheme"] = o.HashScheme
	toSerialize["signature"] = o.Signature
	toSerialize["signatureScheme"] = o.SignatureScheme
	toSerialize["signer"] = o.Signer
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *Message) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"hashScheme",
		"signature",
		"signatureScheme",
		"signer",
		"data",
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

	varMessage := _Message{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessage)

	if err != nil {
		return err
	}

	*o = Message(varMessage)

	return err
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


