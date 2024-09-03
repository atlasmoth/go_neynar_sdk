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

// checks if the VerificationAllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VerificationAllOfData{}

// VerificationAllOfData struct for VerificationAllOfData
type VerificationAllOfData struct {
	Fid int32 `json:"fid"`
	Timestamp int64 `json:"timestamp"`
	Network FarcasterNetwork `json:"network"`
	VerificationAddEthAddressBody VerificationAddEthAddressBody `json:"verificationAddEthAddressBody"`
	Type MessageType `json:"type"`
}

type _VerificationAllOfData VerificationAllOfData

// NewVerificationAllOfData instantiates a new VerificationAllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerificationAllOfData(fid int32, timestamp int64, network FarcasterNetwork, verificationAddEthAddressBody VerificationAddEthAddressBody, type_ MessageType) *VerificationAllOfData {
	this := VerificationAllOfData{}
	this.Fid = fid
	this.Timestamp = timestamp
	this.Network = network
	this.VerificationAddEthAddressBody = verificationAddEthAddressBody
	this.Type = type_
	return &this
}

// NewVerificationAllOfDataWithDefaults instantiates a new VerificationAllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerificationAllOfDataWithDefaults() *VerificationAllOfData {
	this := VerificationAllOfData{}
	var network FarcasterNetwork = FARCASTERNETWORK_MAINNET
	this.Network = network
	var type_ MessageType = MESSAGETYPE_CAST_ADD
	this.Type = type_
	return &this
}

// GetFid returns the Fid field value
func (o *VerificationAllOfData) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *VerificationAllOfData) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *VerificationAllOfData) SetFid(v int32) {
	o.Fid = v
}

// GetTimestamp returns the Timestamp field value
func (o *VerificationAllOfData) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *VerificationAllOfData) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *VerificationAllOfData) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetNetwork returns the Network field value
func (o *VerificationAllOfData) GetNetwork() FarcasterNetwork {
	if o == nil {
		var ret FarcasterNetwork
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *VerificationAllOfData) GetNetworkOk() (*FarcasterNetwork, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *VerificationAllOfData) SetNetwork(v FarcasterNetwork) {
	o.Network = v
}

// GetVerificationAddEthAddressBody returns the VerificationAddEthAddressBody field value
func (o *VerificationAllOfData) GetVerificationAddEthAddressBody() VerificationAddEthAddressBody {
	if o == nil {
		var ret VerificationAddEthAddressBody
		return ret
	}

	return o.VerificationAddEthAddressBody
}

// GetVerificationAddEthAddressBodyOk returns a tuple with the VerificationAddEthAddressBody field value
// and a boolean to check if the value has been set.
func (o *VerificationAllOfData) GetVerificationAddEthAddressBodyOk() (*VerificationAddEthAddressBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerificationAddEthAddressBody, true
}

// SetVerificationAddEthAddressBody sets field value
func (o *VerificationAllOfData) SetVerificationAddEthAddressBody(v VerificationAddEthAddressBody) {
	o.VerificationAddEthAddressBody = v
}

// GetType returns the Type field value
func (o *VerificationAllOfData) GetType() MessageType {
	if o == nil {
		var ret MessageType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VerificationAllOfData) GetTypeOk() (*MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VerificationAllOfData) SetType(v MessageType) {
	o.Type = v
}

func (o VerificationAllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VerificationAllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["network"] = o.Network
	toSerialize["verificationAddEthAddressBody"] = o.VerificationAddEthAddressBody
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *VerificationAllOfData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"timestamp",
		"network",
		"verificationAddEthAddressBody",
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

	varVerificationAllOfData := _VerificationAllOfData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVerificationAllOfData)

	if err != nil {
		return err
	}

	*o = VerificationAllOfData(varVerificationAllOfData)

	return err
}

type NullableVerificationAllOfData struct {
	value *VerificationAllOfData
	isSet bool
}

func (v NullableVerificationAllOfData) Get() *VerificationAllOfData {
	return v.value
}

func (v *NullableVerificationAllOfData) Set(val *VerificationAllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableVerificationAllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableVerificationAllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerificationAllOfData(val *VerificationAllOfData) *NullableVerificationAllOfData {
	return &NullableVerificationAllOfData{value: val, isSet: true}
}

func (v NullableVerificationAllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerificationAllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

