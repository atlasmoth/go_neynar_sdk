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

// checks if the MessageDataVerificationAdd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDataVerificationAdd{}

// MessageDataVerificationAdd struct for MessageDataVerificationAdd
type MessageDataVerificationAdd struct {
	Fid int32 `json:"fid"`
	Timestamp int64 `json:"timestamp"`
	Network FarcasterNetwork `json:"network"`
	VerificationAddEthAddressBody VerificationAddEthAddressBody `json:"verificationAddEthAddressBody"`
}

type _MessageDataVerificationAdd MessageDataVerificationAdd

// NewMessageDataVerificationAdd instantiates a new MessageDataVerificationAdd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDataVerificationAdd(fid int32, timestamp int64, network FarcasterNetwork, verificationAddEthAddressBody VerificationAddEthAddressBody) *MessageDataVerificationAdd {
	this := MessageDataVerificationAdd{}
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.VerificationAddEthAddressBody = verificationAddEthAddressBody
	return &this
}

// NewMessageDataVerificationAddWithDefaults instantiates a new MessageDataVerificationAdd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDataVerificationAddWithDefaults() *MessageDataVerificationAdd {
	this := MessageDataVerificationAdd{}
	var network FarcasterNetwork = FARCASTERNETWORK_MAINNET
	this.Network = network
	return &this
}

// GetFid returns the Fid field value
func (o *MessageDataVerificationAdd) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationAdd) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *MessageDataVerificationAdd) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *MessageDataVerificationAdd) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationAdd) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MessageDataVerificationAdd) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *MessageDataVerificationAdd) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationAdd) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *MessageDataVerificationAdd) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetVerificationAddEthAddressBody returns the VerificationAddEthAddressBody field value
func (o *MessageDataVerificationAdd) GetVerificationAddEthAddressBody() VerificationAddEthAddressBody {
	if o == nil {
		var ret VerificationAddEthAddressBody
		return ret
	}

	return o.VerificationAddEthAddressBody
}

// GetVerificationAddEthAddressBodyOk returns a tuple with the VerificationAddEthAddressBody field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationAdd) GetVerificationAddEthAddressBodyOk() (*VerificationAddEthAddressBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationAddEthAddressBody, true
}

// SetVerificationAddEthAddressBody sets field value
func (o *MessageDataVerificationAdd) SetVerificationAddEthAddressBody(v VerificationAddEthAddressBody) {
	o.VerificationAddEthAddressBody = v
}

func (o MessageDataVerificationAdd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDataVerificationAdd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["network"] = o.Network
	toSerialize["verificationAddEthAddressBody"] = o.VerificationAddEthAddressBody
	return toSerialize, nil
}

func (o *MessageDataVerificationAdd) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"timestamp",
		"network",
		"verificationAddEthAddressBody",
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

	varMessageDataVerificationAdd := _MessageDataVerificationAdd{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDataVerificationAdd)

	if err != nil {
		return err
	}

	*o = MessageDataVerificationAdd(varMessageDataVerificationAdd)

	return err
}

type NullableMessageDataVerificationAdd struct {
	value *MessageDataVerificationAdd
	isSet bool
}

func (v NullableMessageDataVerificationAdd) Get() *MessageDataVerificationAdd {
	return v.value
}

func (v *NullableMessageDataVerificationAdd) Set(val *MessageDataVerificationAdd) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDataVerificationAdd) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDataVerificationAdd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDataVerificationAdd(val *MessageDataVerificationAdd) *NullableMessageDataVerificationAdd {
	return &NullableMessageDataVerificationAdd{value: val, isSet: true}
}

func (v NullableMessageDataVerificationAdd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDataVerificationAdd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

