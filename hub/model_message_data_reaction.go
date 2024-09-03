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

// checks if the MessageDataReaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDataReaction{}

// MessageDataReaction struct for MessageDataReaction
type MessageDataReaction struct {
	Fid int32 `json:"fid"`
	Timestamp int64 `json:"timestamp"`
	Network FarcasterNetwork `json:"network"`
	ReactionBody ReactionBody `json:"reactionBody"`
}

type _MessageDataReaction MessageDataReaction

// NewMessageDataReaction instantiates a new MessageDataReaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDataReaction(fid int32, timestamp int64, network FarcasterNetwork, reactionBody ReactionBody) *MessageDataReaction {
	this := MessageDataReaction{}
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.ReactionBody = reactionBody
	return &this
}

// NewMessageDataReactionWithDefaults instantiates a new MessageDataReaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDataReactionWithDefaults() *MessageDataReaction {
	this := MessageDataReaction{}
	var network FarcasterNetwork = FARCASTERNETWORK_MAINNET
	this.Network = network
	return &this
}

// GetFid returns the Fid field value
func (o *MessageDataReaction) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *MessageDataReaction) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *MessageDataReaction) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *MessageDataReaction) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MessageDataReaction) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MessageDataReaction) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *MessageDataReaction) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *MessageDataReaction) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *MessageDataReaction) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetReactionBody returns the ReactionBody field value
func (o *MessageDataReaction) GetReactionBody() ReactionBody {
	if o == nil {
		var ret ReactionBody
		return ret
	}

	return o.ReactionBody
}

// GetReactionBodyOk returns a tuple with the ReactionBody field value
// and a boolean to check if the value has been set.
func (o *MessageDataReaction) GetReactionBodyOk() (*ReactionBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactionBody, true
}

// SetReactionBody sets field value
func (o *MessageDataReaction) SetReactionBody(v ReactionBody) {
	o.ReactionBody = v
}

func (o MessageDataReaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDataReaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["network"] = o.Network
	toSerialize["reactionBody"] = o.ReactionBody
	return toSerialize, nil
}

func (o *MessageDataReaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"timestamp",
		"network",
		"reactionBody",
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

	varMessageDataReaction := _MessageDataReaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDataReaction)

	if err != nil {
		return err
	}

	*o = MessageDataReaction(varMessageDataReaction)

	return err
}

type NullableMessageDataReaction struct {
	value *MessageDataReaction
	isSet bool
}

func (v NullableMessageDataReaction) Get() *MessageDataReaction {
	return v.value
}

func (v *NullableMessageDataReaction) Set(val *MessageDataReaction) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDataReaction) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDataReaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDataReaction(val *MessageDataReaction) *NullableMessageDataReaction {
	return &NullableMessageDataReaction{value: val, isSet: true}
}

func (v NullableMessageDataReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDataReaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

