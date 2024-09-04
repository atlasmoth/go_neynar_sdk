# OnChainEventIdRegister

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**ChainId** | Pointer to **int32** |  | [optional] 
**BlockNumber** | Pointer to **int32** |  | [optional] 
**BlockHash** | Pointer to **string** |  | [optional] 
**BlockTimestamp** | Pointer to **int32** |  | [optional] 
**TransactionHash** | Pointer to **string** |  | [optional] 
**LogIndex** | Pointer to **int32** |  | [optional] 
**TxIndex** | Pointer to **int32** |  | [optional] 
**Fid** | Pointer to **int32** |  | [optional] 
**IdRegisterEventBody** | Pointer to [**IdRegisterEventBody**](IdRegisterEventBody.md) |  | [optional] 

## Methods

### NewOnChainEventIdRegister

`func NewOnChainEventIdRegister() *OnChainEventIdRegister`

NewOnChainEventIdRegister instantiates a new OnChainEventIdRegister object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnChainEventIdRegisterWithDefaults

`func NewOnChainEventIdRegisterWithDefaults() *OnChainEventIdRegister`

NewOnChainEventIdRegisterWithDefaults instantiates a new OnChainEventIdRegister object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OnChainEventIdRegister) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OnChainEventIdRegister) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OnChainEventIdRegister) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OnChainEventIdRegister) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChainId

`func (o *OnChainEventIdRegister) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *OnChainEventIdRegister) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *OnChainEventIdRegister) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *OnChainEventIdRegister) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetBlockNumber

`func (o *OnChainEventIdRegister) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *OnChainEventIdRegister) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *OnChainEventIdRegister) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *OnChainEventIdRegister) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockHash

`func (o *OnChainEventIdRegister) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *OnChainEventIdRegister) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *OnChainEventIdRegister) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *OnChainEventIdRegister) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockTimestamp

`func (o *OnChainEventIdRegister) GetBlockTimestamp() int32`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *OnChainEventIdRegister) GetBlockTimestampOk() (*int32, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *OnChainEventIdRegister) SetBlockTimestamp(v int32)`

SetBlockTimestamp sets BlockTimestamp field to given value.

### HasBlockTimestamp

`func (o *OnChainEventIdRegister) HasBlockTimestamp() bool`

HasBlockTimestamp returns a boolean if a field has been set.

### GetTransactionHash

`func (o *OnChainEventIdRegister) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *OnChainEventIdRegister) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *OnChainEventIdRegister) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *OnChainEventIdRegister) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetLogIndex

`func (o *OnChainEventIdRegister) GetLogIndex() int32`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *OnChainEventIdRegister) GetLogIndexOk() (*int32, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *OnChainEventIdRegister) SetLogIndex(v int32)`

SetLogIndex sets LogIndex field to given value.

### HasLogIndex

`func (o *OnChainEventIdRegister) HasLogIndex() bool`

HasLogIndex returns a boolean if a field has been set.

### GetTxIndex

`func (o *OnChainEventIdRegister) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *OnChainEventIdRegister) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *OnChainEventIdRegister) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *OnChainEventIdRegister) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetFid

`func (o *OnChainEventIdRegister) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *OnChainEventIdRegister) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *OnChainEventIdRegister) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *OnChainEventIdRegister) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetIdRegisterEventBody

`func (o *OnChainEventIdRegister) GetIdRegisterEventBody() IdRegisterEventBody`

GetIdRegisterEventBody returns the IdRegisterEventBody field if non-nil, zero value otherwise.

### GetIdRegisterEventBodyOk

`func (o *OnChainEventIdRegister) GetIdRegisterEventBodyOk() (*IdRegisterEventBody, bool)`

GetIdRegisterEventBodyOk returns a tuple with the IdRegisterEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdRegisterEventBody

`func (o *OnChainEventIdRegister) SetIdRegisterEventBody(v IdRegisterEventBody)`

SetIdRegisterEventBody sets IdRegisterEventBody field to given value.

### HasIdRegisterEventBody

`func (o *OnChainEventIdRegister) HasIdRegisterEventBody() bool`

HasIdRegisterEventBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


