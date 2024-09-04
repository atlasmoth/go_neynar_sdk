/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OnChainEventCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnChainEventCommon{}

// OnChainEventCommon struct for OnChainEventCommon
type OnChainEventCommon struct {
	Type *string `json:"type,omitempty"`
	ChainId *int32 `json:"chainId,omitempty"`
	BlockNumber *int32 `json:"blockNumber,omitempty"`
	BlockHash *string `json:"blockHash,omitempty"`
	BlockTimestamp *int32 `json:"blockTimestamp,omitempty"`
	TransactionHash *string `json:"transactionHash,omitempty"`
	LogIndex *int32 `json:"logIndex,omitempty"`
	TxIndex *int32 `json:"txIndex,omitempty"`
	Fid *int32 `json:"fid,omitempty"`
}

// NewOnChainEventCommon instantiates a new OnChainEventCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainEventCommon() *OnChainEventCommon {
	this := OnChainEventCommon{}
	return &this
}

// NewOnChainEventCommonWithDefaults instantiates a new OnChainEventCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainEventCommonWithDefaults() *OnChainEventCommon {
	this := OnChainEventCommon{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OnChainEventCommon) SetType(v string) {
	o.Type = &v
}

// GetChainId returns the ChainId field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetChainId() int32 {
	if o == nil || IsNil(o.ChainId) {
		var ret int32
		return ret
	}
	return *o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetChainIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ChainId) {
		return nil, false
	}
	return o.ChainId, true
}

// HasChainId returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasChainId() bool {
	if o != nil && !IsNil(o.ChainId) {
		return true
	}

	return false
}

// SetChainId gets a reference to the given int32 and assigns it to the ChainId field.
func (o *OnChainEventCommon) SetChainId(v int32) {
	o.ChainId = &v
}

// GetBlockNumber returns the BlockNumber field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetBlockNumber() int32 {
	if o == nil || IsNil(o.BlockNumber) {
		var ret int32
		return ret
	}
	return *o.BlockNumber
}

// GetBlockNumberOk returns a tuple with the BlockNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetBlockNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockNumber) {
		return nil, false
	}
	return o.BlockNumber, true
}

// HasBlockNumber returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasBlockNumber() bool {
	if o != nil && !IsNil(o.BlockNumber) {
		return true
	}

	return false
}

// SetBlockNumber gets a reference to the given int32 and assigns it to the BlockNumber field.
func (o *OnChainEventCommon) SetBlockNumber(v int32) {
	o.BlockNumber = &v
}

// GetBlockHash returns the BlockHash field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetBlockHash() string {
	if o == nil || IsNil(o.BlockHash) {
		var ret string
		return ret
	}
	return *o.BlockHash
}

// GetBlockHashOk returns a tuple with the BlockHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetBlockHashOk() (*string, bool) {
	if o == nil || IsNil(o.BlockHash) {
		return nil, false
	}
	return o.BlockHash, true
}

// HasBlockHash returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasBlockHash() bool {
	if o != nil && !IsNil(o.BlockHash) {
		return true
	}

	return false
}

// SetBlockHash gets a reference to the given string and assigns it to the BlockHash field.
func (o *OnChainEventCommon) SetBlockHash(v string) {
	o.BlockHash = &v
}

// GetBlockTimestamp returns the BlockTimestamp field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetBlockTimestamp() int32 {
	if o == nil || IsNil(o.BlockTimestamp) {
		var ret int32
		return ret
	}
	return *o.BlockTimestamp
}

// GetBlockTimestampOk returns a tuple with the BlockTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetBlockTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.BlockTimestamp) {
		return nil, false
	}
	return o.BlockTimestamp, true
}

// HasBlockTimestamp returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasBlockTimestamp() bool {
	if o != nil && !IsNil(o.BlockTimestamp) {
		return true
	}

	return false
}

// SetBlockTimestamp gets a reference to the given int32 and assigns it to the BlockTimestamp field.
func (o *OnChainEventCommon) SetBlockTimestamp(v int32) {
	o.BlockTimestamp = &v
}

// GetTransactionHash returns the TransactionHash field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetTransactionHash() string {
	if o == nil || IsNil(o.TransactionHash) {
		var ret string
		return ret
	}
	return *o.TransactionHash
}

// GetTransactionHashOk returns a tuple with the TransactionHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetTransactionHashOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionHash) {
		return nil, false
	}
	return o.TransactionHash, true
}

// HasTransactionHash returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasTransactionHash() bool {
	if o != nil && !IsNil(o.TransactionHash) {
		return true
	}

	return false
}

// SetTransactionHash gets a reference to the given string and assigns it to the TransactionHash field.
func (o *OnChainEventCommon) SetTransactionHash(v string) {
	o.TransactionHash = &v
}

// GetLogIndex returns the LogIndex field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetLogIndex() int32 {
	if o == nil || IsNil(o.LogIndex) {
		var ret int32
		return ret
	}
	return *o.LogIndex
}

// GetLogIndexOk returns a tuple with the LogIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetLogIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.LogIndex) {
		return nil, false
	}
	return o.LogIndex, true
}

// HasLogIndex returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasLogIndex() bool {
	if o != nil && !IsNil(o.LogIndex) {
		return true
	}

	return false
}

// SetLogIndex gets a reference to the given int32 and assigns it to the LogIndex field.
func (o *OnChainEventCommon) SetLogIndex(v int32) {
	o.LogIndex = &v
}

// GetTxIndex returns the TxIndex field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetTxIndex() int32 {
	if o == nil || IsNil(o.TxIndex) {
		var ret int32
		return ret
	}
	return *o.TxIndex
}

// GetTxIndexOk returns a tuple with the TxIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetTxIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.TxIndex) {
		return nil, false
	}
	return o.TxIndex, true
}

// HasTxIndex returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasTxIndex() bool {
	if o != nil && !IsNil(o.TxIndex) {
		return true
	}

	return false
}

// SetTxIndex gets a reference to the given int32 and assigns it to the TxIndex field.
func (o *OnChainEventCommon) SetTxIndex(v int32) {
	o.TxIndex = &v
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *OnChainEventCommon) GetFid() int32 {
	if o == nil || IsNil(o.Fid) {
		var ret int32
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnChainEventCommon) GetFidOk() (*int32, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *OnChainEventCommon) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given int32 and assigns it to the Fid field.
func (o *OnChainEventCommon) SetFid(v int32) {
	o.Fid = &v
}

func (o OnChainEventCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnChainEventCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ChainId) {
		toSerialize["chainId"] = o.ChainId
	}
	if !IsNil(o.BlockNumber) {
		toSerialize["blockNumber"] = o.BlockNumber
	}
	if !IsNil(o.BlockHash) {
		toSerialize["blockHash"] = o.BlockHash
	}
	if !IsNil(o.BlockTimestamp) {
		toSerialize["blockTimestamp"] = o.BlockTimestamp
	}
	if !IsNil(o.TransactionHash) {
		toSerialize["transactionHash"] = o.TransactionHash
	}
	if !IsNil(o.LogIndex) {
		toSerialize["logIndex"] = o.LogIndex
	}
	if !IsNil(o.TxIndex) {
		toSerialize["txIndex"] = o.TxIndex
	}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	return toSerialize, nil
}

type NullableOnChainEventCommon struct {
	value *OnChainEventCommon
	isSet bool
}

func (v NullableOnChainEventCommon) Get() *OnChainEventCommon {
	return v.value
}

func (v *NullableOnChainEventCommon) Set(val *OnChainEventCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainEventCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainEventCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainEventCommon(val *OnChainEventCommon) *NullableOnChainEventCommon {
	return &NullableOnChainEventCommon{value: val, isSet: true}
}

func (v NullableOnChainEventCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainEventCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


