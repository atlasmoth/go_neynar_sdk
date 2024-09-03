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

// checks if the MessageDataLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageDataLink{}

// MessageDataLink struct for MessageDataLink
type MessageDataLink struct {
	Fid int32 `json:"fid"`
	Timestamp int64 `json:"timestamp"`
	Network FarcasterNetwork `json:"network"`
	LinkBody LinkBody `json:"linkBody"`
}

type _MessageDataLink MessageDataLink

// NewMessageDataLink instantiates a new MessageDataLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageDataLink(fid int32, timestamp int64, network FarcasterNetwork, linkBody LinkBody) *MessageDataLink {
	this := MessageDataLink{}
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.LinkBody = linkBody
	return &this
}

// NewMessageDataLinkWithDefaults instantiates a new MessageDataLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageDataLinkWithDefaults() *MessageDataLink {
	this := MessageDataLink{}
	var network FarcasterNetwork = FARCASTERNETWORK_MAINNET
	this.Network = network
	return &this
}

// GetFid returns the Fid field value
func (o *MessageDataLink) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *MessageDataLink) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *MessageDataLink) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *MessageDataLink) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *MessageDataLink) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *MessageDataLink) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *MessageDataLink) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *MessageDataLink) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *MessageDataLink) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetLinkBody returns the LinkBody field value
func (o *MessageDataLink) GetLinkBody() LinkBody {
	if o == nil {
		var ret LinkBody
		return ret
	}

	return o.LinkBody
}

// GetLinkBodyOk returns a tuple with the LinkBody field value
// and a boolean to check if the value has been set.
func (o *MessageDataLink) GetLinkBodyOk() (*LinkBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkBody, true
}

// SetLinkBody sets field value
func (o *MessageDataLink) SetLinkBody(v LinkBody) {
	o.LinkBody = v
}

func (o MessageDataLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageDataLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["network"] = o.Network
	toSerialize["linkBody"] = o.LinkBody
	return toSerialize, nil
}

func (o *MessageDataLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"timestamp",
		"network",
		"linkBody",
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

	varMessageDataLink := _MessageDataLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMessageDataLink)

	if err != nil {
		return err
	}

	*o = MessageDataLink(varMessageDataLink)

	return err
}

type NullableMessageDataLink struct {
	value *MessageDataLink
	isSet bool
}

func (v NullableMessageDataLink) Get() *MessageDataLink {
	return v.value
}

func (v *NullableMessageDataLink) Set(val *MessageDataLink) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageDataLink) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageDataLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageDataLink(val *MessageDataLink) *NullableMessageDataLink {
	return &NullableMessageDataLink{value: val, isSet: true}
}

func (v NullableMessageDataLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageDataLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


