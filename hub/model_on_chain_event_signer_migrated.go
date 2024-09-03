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

// checks if the OnChainEventSignerMigrated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnChainEventSignerMigrated{}

// OnChainEventSignerMigrated struct for OnChainEventSignerMigrated
type OnChainEventSignerMigrated struct {
	Type string `json:"type"`
	ChainId int32 `json:"chainId"`
	BlockNumber int32 `json:"blockNumber"`
	BlockHash string `json:"blockHash"`
	BlockTimestamp int32 `json:"blockTimestamp"`
	TransactionHash string `json:"transactionHash"`
	LogIndex int32 `json:"logIndex"`
	TxIndex int32 `json:"txIndex"`
	Fid int32 `json:"fid"`
	SignerMigratedEventBody SignerMigratedEventBody `json:"signerMigratedEventBody"`
	AdditionalProperties map[string]interface{}
}

type _OnChainEventSignerMigrated OnChainEventSignerMigrated

// NewOnChainEventSignerMigrated instantiates a new OnChainEventSignerMigrated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainEventSignerMigrated(type_ string, chainId int32, blockNumber int32, blockHash string, blockTimestamp int32, transactionHash string, logIndex int32, txIndex int32, fid int32, signerMigratedEventBody SignerMigratedEventBody) *OnChainEventSignerMigrated {
	this := OnChainEventSignerMigrated{}
	this.Type = type_
	this.ChainId = chainId
	this.BlockNumber = blockNumber
	this.BlockHash = blockHash
	this.BlockTimestamp = blockTimestamp
	this.TransactionHash = transactionHash
	this.LogIndex = logIndex
	this.TxIndex = txIndex
	this.Fid = fid
	this.SignerMigratedEventBody = signerMigratedEventBody
	return &this
}

// NewOnChainEventSignerMigratedWithDefaults instantiates a new OnChainEventSignerMigrated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainEventSignerMigratedWithDefaults() *OnChainEventSignerMigrated {
	this := OnChainEventSignerMigrated{}
	return &this
}

// GetType returns the Type field value
func (o *OnChainEventSignerMigrated) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OnChainEventSignerMigrated) SetType(v string) {
	o.Type = v
}

// GetChainId returns the ChainId field value
func (o *OnChainEventSignerMigrated) GetChainId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetChainIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *OnChainEventSignerMigrated) SetChainId(v int32) {
	o.ChainId = v
}

// GetBlockNumber returns the BlockNumber field value
func (o *OnChainEventSignerMigrated) GetBlockNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetBlockNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockNumber, true
}

// SetBlockNumber sets field value
func (o *OnChainEventSignerMigrated) SetBlockNumber(v int32) {
	o.BlockNumber = v
}

// GetBlockHash returns the BlockHash field value
func (o *OnChainEventSignerMigrated) GetBlockHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetBlockHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockHash, true
}

// SetBlockHash sets field value
func (o *OnChainEventSignerMigrated) SetBlockHash(v string) {
	o.BlockHash = v
}

// GetBlockTimestamp returns the BlockTimestamp field value
func (o *OnChainEventSignerMigrated) GetBlockTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlockTimestamp
}

// GetBlockTimestampOk returns a tuple with the BlockTimestamp field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetBlockTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlockTimestamp, true
}

// SetBlockTimestamp sets field value
func (o *OnChainEventSignerMigrated) SetBlockTimestamp(v int32) {
	o.BlockTimestamp = v
}

// GetTransactionHash returns the TransactionHash field value
func (o *OnChainEventSignerMigrated) GetTransactionHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetTransactionHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionHash, true
}

// SetTransactionHash sets field value
func (o *OnChainEventSignerMigrated) SetTransactionHash(v string) {
	o.TransactionHash = v
}

// GetLogIndex returns the LogIndex field value
func (o *OnChainEventSignerMigrated) GetLogIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LogIndex
}

// GetLogIndexOk returns a tuple with the LogIndex field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetLogIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogIndex, true
}

// SetLogIndex sets field value
func (o *OnChainEventSignerMigrated) SetLogIndex(v int32) {
	o.LogIndex = v
}

// GetTxIndex returns the TxIndex field value
func (o *OnChainEventSignerMigrated) GetTxIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetTxIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxIndex, true
}

// SetTxIndex sets field value
func (o *OnChainEventSignerMigrated) SetTxIndex(v int32) {
	o.TxIndex = v
}

// GetFid returns the Fid field value
func (o *OnChainEventSignerMigrated) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *OnChainEventSignerMigrated) SetFid(v int32) {
	o.Fid = v
}

// GetSignerMigratedEventBody returns the SignerMigratedEventBody field value
func (o *OnChainEventSignerMigrated) GetSignerMigratedEventBody() SignerMigratedEventBody {
	if o == nil {
		var ret SignerMigratedEventBody
		return ret
	}

	return o.SignerMigratedEventBody
}

// GetSignerMigratedEventBodyOk returns a tuple with the SignerMigratedEventBody field value
// and a boolean to check if the value has been set.
func (o *OnChainEventSignerMigrated) GetSignerMigratedEventBodyOk() (*SignerMigratedEventBody, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignerMigratedEventBody, true
}

// SetSignerMigratedEventBody sets field value
func (o *OnChainEventSignerMigrated) SetSignerMigratedEventBody(v SignerMigratedEventBody) {
	o.SignerMigratedEventBody = v
}

func (o OnChainEventSignerMigrated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnChainEventSignerMigrated) ToMap() (map[string]interface{}, error) {
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
	toSerialize["signerMigratedEventBody"] = o.SignerMigratedEventBody

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnChainEventSignerMigrated) UnmarshalJSON(data []byte) (err error) {
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
		"signerMigratedEventBody",
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

	varOnChainEventSignerMigrated := _OnChainEventSignerMigrated{}

	err = json.Unmarshal(data, &varOnChainEventSignerMigrated)

	if err != nil {
		return err
	}

	*o = OnChainEventSignerMigrated(varOnChainEventSignerMigrated)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "chainId")
		delete(additionalProperties, "blockNumber")
		delete(additionalProperties, "blockHash")
		delete(additionalProperties, "blockTimestamp")
		delete(additionalProperties, "transactionHash")
		delete(additionalProperties, "logIndex")
		delete(additionalProperties, "txIndex")
		delete(additionalProperties, "fid")
		delete(additionalProperties, "signerMigratedEventBody")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnChainEventSignerMigrated struct {
	value *OnChainEventSignerMigrated
	isSet bool
}

func (v NullableOnChainEventSignerMigrated) Get() *OnChainEventSignerMigrated {
	return v.value
}

func (v *NullableOnChainEventSignerMigrated) Set(val *OnChainEventSignerMigrated) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainEventSignerMigrated) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainEventSignerMigrated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainEventSignerMigrated(val *OnChainEventSignerMigrated) *NullableOnChainEventSignerMigrated {
	return &NullableOnChainEventSignerMigrated{value: val, isSet: true}
}

func (v NullableOnChainEventSignerMigrated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainEventSignerMigrated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


