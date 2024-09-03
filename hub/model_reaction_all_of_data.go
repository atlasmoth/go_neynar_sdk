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

// checks if the ReactionAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionAllOfData{}

// ReactionAllOfData struct for ReactionAllOfData
type ReactionAllOfData struct {
	Fid int32 `json:"fid"`
	Timestamp int64 `json:"timestamp"`
	Network FarcasterNetwork `json:"network"`
	ReactionBody ReactionBody `json:"reactionBody"`
	Type MessageType `json:"type"`
}

type _ReactionAllOfData ReactionAllOfData

// NewReactionAllOfData instantiates a new ReactionAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionAllOfData(fid int32, timestamp int64, network FarcasterNetwork, reactionBody ReactionBody, type_ MessageType) *ReactionAllOfData {
	this := ReactionAllOfData{}
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.ReactionBody = reactionBody
	this.Type = type_
	return &this
}

// NewReactionAllOfDataWithDefaults instantiates a new ReactionAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionAllOfDataWithDefaults() *ReactionAllOfData {
	this := ReactionAllOfData{}
	var network FarcasterNetwork = FARCASTERNETWORK_MAINNET
	this.Network = network
	var type_ MessageType = MESSAGETYPE_CAST_ADD
	this.Type = type_
	return &this
}

// GetFid returns the Fid field value
func (o *ReactionAllOfData) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *ReactionAllOfData) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *ReactionAllOfData) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *ReactionAllOfData) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ReactionAllOfData) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ReactionAllOfData) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *ReactionAllOfData) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *ReactionAllOfData) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *ReactionAllOfData) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetReactionBody returns the ReactionBody field value
func (o *ReactionAllOfData) GetReactionBody() ReactionBody {
	if o == nil {
		var ret ReactionBody
		return ret
	}

	return o.ReactionBody
}

// GetReactionBodyOk returns a tuple with the ReactionBody field value
// and a boolean to check if the value has been set.
func (o *ReactionAllOfData) GetReactionBodyOk() (*ReactionBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactionBody, true
}

// SetReactionBody sets field value
func (o *ReactionAllOfData) SetReactionBody(v ReactionBody) {
	o.ReactionBody = v
}

// GetType returns the Type field value
func (o *ReactionAllOfData) GetType() MessageType {
	if o == nil {
		var ret MessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReactionAllOfData) GetTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReactionAllOfData) SetType(v MessageType) {
	o.Type = v
}

func (o ReactionAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["network"] = o.Network
	toSerialize["reactionBody"] = o.ReactionBody
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *ReactionAllOfData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"timestamp",
		"network",
		"reactionBody",
		"type",
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

	varReactionAllOfData := _ReactionAllOfData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReactionAllOfData)

	if err != nil {
		return err
	}

	*o = ReactionAllOfData(varReactionAllOfData)

	return err
}

type NullableReactionAllOfData struct {
	value *ReactionAllOfData
	isSet bool
}

func (v NullableReactionAllOfData) Get() *ReactionAllOfData {
	return v.value
}

func (v *NullableReactionAllOfData) Set(val *ReactionAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionAllOfData(val *ReactionAllOfData) *NullableReactionAllOfData {
	return &NullableReactionAllOfData{value: val, isSet: true}
}

func (v NullableReactionAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


