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

// checks if the MessageDataVerificationRemove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDataVerificationRemove{}

// MessageDataVerificationRemove struct for MessageDataVerificationRemove
type MessageDataVerificationRemove struct {
	Fid int32 `json:"fid"`
	Timestamp int64 `json:"timestamp"`
	Network FarcasterNetwork `json:"network"`
	VerificationRemoveBody VerificationRemoveBody `json:"verificationRemoveBody"`
}

type _MessageDataVerificationRemove MessageDataVerificationRemove

// NewMessageDataVerificationRemove instantiates a new MessageDataVerificationRemove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDataVerificationRemove(fid int32, timestamp int64, network FarcasterNetwork, verificationRemoveBody VerificationRemoveBody) *MessageDataVerificationRemove {
	this := MessageDataVerificationRemove{}
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.VerificationRemoveBody = verificationRemoveBody
	return &this
}

// NewMessageDataVerificationRemoveWithDefaults instantiates a new MessageDataVerificationRemove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDataVerificationRemoveWithDefaults() *MessageDataVerificationRemove {
	this := MessageDataVerificationRemove{}
	var network FarcasterNetwork = FARCASTERNETWORK_MAINNET
	this.Network = network
	return &this
}

// GetFid returns the Fid field value
func (o *MessageDataVerificationRemove) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationRemove) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *MessageDataVerificationRemove) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *MessageDataVerificationRemove) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationRemove) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MessageDataVerificationRemove) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *MessageDataVerificationRemove) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationRemove) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *MessageDataVerificationRemove) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetVerificationRemoveBody returns the VerificationRemoveBody field value
func (o *MessageDataVerificationRemove) GetVerificationRemoveBody() VerificationRemoveBody {
	if o == nil {
		var ret VerificationRemoveBody
		return ret
	}

	return o.VerificationRemoveBody
}

// GetVerificationRemoveBodyOk returns a tuple with the VerificationRemoveBody field value
// and a boolean to check if the value has been set.
func (o *MessageDataVerificationRemove) GetVerificationRemoveBodyOk() (*VerificationRemoveBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationRemoveBody, true
}

// SetVerificationRemoveBody sets field value
func (o *MessageDataVerificationRemove) SetVerificationRemoveBody(v VerificationRemoveBody) {
	o.VerificationRemoveBody = v
}

func (o MessageDataVerificationRemove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDataVerificationRemove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["network"] = o.Network
	toSerialize["verificationRemoveBody"] = o.VerificationRemoveBody
	return toSerialize, nil
}

func (o *MessageDataVerificationRemove) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"timestamp",
		"network",
		"verificationRemoveBody",
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

	varMessageDataVerificationRemove := _MessageDataVerificationRemove{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDataVerificationRemove)

	if err != nil {
		return err
	}

	*o = MessageDataVerificationRemove(varMessageDataVerificationRemove)

	return err
}

type NullableMessageDataVerificationRemove struct {
	value *MessageDataVerificationRemove
	isSet bool
}

func (v NullableMessageDataVerificationRemove) Get() *MessageDataVerificationRemove {
	return v.value
}

func (v *NullableMessageDataVerificationRemove) Set(val *MessageDataVerificationRemove) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDataVerificationRemove) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDataVerificationRemove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDataVerificationRemove(val *MessageDataVerificationRemove) *NullableMessageDataVerificationRemove {
	return &NullableMessageDataVerificationRemove{value: val, isSet: true}
}

func (v NullableMessageDataVerificationRemove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDataVerificationRemove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

