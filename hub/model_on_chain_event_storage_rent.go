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

// checks if the OnChainEventStorageRent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnChainEventStorageRent{}

// OnChainEventStorageRent struct for OnChainEventStorageRent
type OnChainEventStorageRent struct {
	Type string `json:"type"`
	ChainId int32 `json:"chainId"`
	BlockNumber int32 `json:"blockNumber"`
	BlockHash string `json:"blockHash"`
	BlockTimestamp int32 `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
	LogIndex int32 `json:"logIndex"`
	TxIndex int32 `json:"txIndex"`
	Fid int32 `json:"fid"`
	StorageRentEventBody StorageRentEventBody `json:"storageRentEventBody"`
}

type _OnChainEventStorageRent OnChainEventStorageRent

// NewOnChainEventStorageRent instantiates a new OnChainEventStorageRent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainEventStorageRent(type_ string, chainId int32, blockNumber int32, blockHash string, blockTimestamp int32, transactionHash string, logIndex int32, txIndex int32, fid int32, storageRentEventBody StorageRentEventBody) *OnChainEventStorageRent {
	this := OnChainEventStorageRent{}
	this.Type = type_
	this.ChainId = chainId
	this.BlockNumber = blockNumber
	this.BlockHash = blockHash
	this.BlockTimestamp = blockTimestamp
	this.TransactionHash = transactionHash
	this.LogIndex = logIndex
	this.TxIndex = txIndex
	this.Fid = fid
	this.StorageRentEventBody = storageRentEventBody
	return &this
}

// NewOnChainEventStorageRentWithDefaults instantiates a new OnChainEventStorageRent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainEventStorageRentWithDefaults() *OnChainEventStorageRent {
	this := OnChainEventStorageRent{}
	return &this
}

// GetType returns the Type field value
func (o *OnChainEventStorageRent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OnChainEventStorageRent) SetType(v string) {
	o.Type = v
}

// GetChainId returns the ChainId field value
func (o *OnChainEventStorageRent) GetChainId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetChainIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *OnChainEventStorageRent) SetChainId(v int32) {
	o.ChainId = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *OnChainEventStorageRent) GetBlockNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetBlockNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *OnChainEventStorageRent) SetBlockNumber(v int32) {
	o.BlockNumber = v
}

// GetBlockHash returns the BlockHash field value
func (o *OnChainEventStorageRent) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *OnChainEventStorageRent) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockTimestamp returns the BlockTimestamp field value
func (o *OnChainEventStorageRent) GetBlockTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockTimestamp
}

// GetBlockTimestampOk returns a tuple with the BlockTimestamp field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetBlockTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimestamp, true
}

// SetBlockTimestamp sets field value
func (o *OnChainEventStorageRent) SetBlockTimestamp(v int32) {
	o.BlockTimestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *OnChainEventStorageRent) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *OnChainEventStorageRent) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetLogIndex returns the LogIndex field value
func (o *OnChainEventStorageRent) GetLogIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LogIndex
}

// GetLogIndexOk returns a tuple with the LogIndex field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetLogIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogIndex, true
}

// SetLogIndex sets field value
func (o *OnChainEventStorageRent) SetLogIndex(v int32) {
	o.LogIndex = v
}

// GetTxIndex returns the TxIndex field value
func (o *OnChainEventStorageRent) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *OnChainEventStorageRent) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetFid returns the Fid field value
func (o *OnChainEventStorageRent) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *OnChainEventStorageRent) SetFid(v int32) {
	o.Fid = v
}

// GetStorageRentEventBody returns the StorageRentEventBody field value
func (o *OnChainEventStorageRent) GetStorageRentEventBody() StorageRentEventBody {
	if o == nil {
		var ret StorageRentEventBody
		return ret
	}

	return o.StorageRentEventBody
}

// GetStorageRentEventBodyOk returns a tuple with the StorageRentEventBody field value
// and a boolean to check if the value has been set.
func (o *OnChainEventStorageRent) GetStorageRentEventBodyOk() (*StorageRentEventBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageRentEventBody, true
}

// SetStorageRentEventBody sets field value
func (o *OnChainEventStorageRent) SetStorageRentEventBody(v StorageRentEventBody) {
	o.StorageRentEventBody = v
}

func (o OnChainEventStorageRent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnChainEventStorageRent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["chainId"] = o.ChainId
	toSerialize["blockNumber"] = o.BlockNumber
	toSerialize["blockHash"] = o.BlockHash
	toSerialize["blockTimestamp"] = o.BlockTimestamp
	toSerialize["transactionHash"] = o.TransactionHash
	toSerialize["logIndex"] = o.LogIndex
	toSerialize["txIndex"] = o.TxIndex
	toSerialize["fid"] = o.Fid
	toSerialize["storageRentEventBody"] = o.StorageRentEventBody
	return toSerialize, nil
}

func (o *OnChainEventStorageRent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"chainId",
		"blockNumber",
		"blockHash",
		"blockTimestamp",
		"transactionHash",
		"logIndex",
		"txIndex",
		"fid",
		"storageRentEventBody",
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

	varOnChainEventStorageRent := _OnChainEventStorageRent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnChainEventStorageRent)

	if err != nil {
		return err
	}

	*o = OnChainEventStorageRent(varOnChainEventStorageRent)

	return err
}

type NullableOnChainEventStorageRent struct {
	value *OnChainEventStorageRent
	isSet bool
}

func (v NullableOnChainEventStorageRent) Get() *OnChainEventStorageRent {
	return v.value
}

func (v *NullableOnChainEventStorageRent) Set(val *OnChainEventStorageRent) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainEventStorageRent) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainEventStorageRent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainEventStorageRent(val *OnChainEventStorageRent) *NullableOnChainEventStorageRent {
	return &NullableOnChainEventStorageRent{value: val, isSet: true}
}

func (v NullableOnChainEventStorageRent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainEventStorageRent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

