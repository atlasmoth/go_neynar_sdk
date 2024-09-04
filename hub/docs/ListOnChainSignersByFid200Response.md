# ListOnChainSignersByFid200Response

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
**SignerEventBody** | Pointer to [**SignerEventBody**](SignerEventBody.md) |  | [optional] 
**Events** | Pointer to [**[]OnChainEventSigner**](OnChainEventSigner.md) |  | [optional] 

## Methods

### NewListOnChainSignersByFid200Response

`func NewListOnChainSignersByFid200Response() *ListOnChainSignersByFid200Response`

NewListOnChainSignersByFid200Response instantiates a new ListOnChainSignersByFid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOnChainSignersByFid200ResponseWithDefaults

`func NewListOnChainSignersByFid200ResponseWithDefaults() *ListOnChainSignersByFid200Response`

NewListOnChainSignersByFid200ResponseWithDefaults instantiates a new ListOnChainSignersByFid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListOnChainSignersByFid200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOnChainSignersByFid200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOnChainSignersByFid200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListOnChainSignersByFid200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChainId

`func (o *ListOnChainSignersByFid200Response) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ListOnChainSignersByFid200Response) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ListOnChainSignersByFid200Response) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ListOnChainSignersByFid200Response) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetBlockNumber

`func (o *ListOnChainSignersByFid200Response) GetBlockNumber() int32`

GetBlockNumber returns the BlockNumber field if non-nil, zero value otherwise.

### GetBlockNumberOk

`func (o *ListOnChainSignersByFid200Response) GetBlockNumberOk() (*int32, bool)`

GetBlockNumberOk returns a tuple with the BlockNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockNumber

`func (o *ListOnChainSignersByFid200Response) SetBlockNumber(v int32)`

SetBlockNumber sets BlockNumber field to given value.

### HasBlockNumber

`func (o *ListOnChainSignersByFid200Response) HasBlockNumber() bool`

HasBlockNumber returns a boolean if a field has been set.

### GetBlockHash

`func (o *ListOnChainSignersByFid200Response) GetBlockHash() string`

GetBlockHash returns the BlockHash field if non-nil, zero value otherwise.

### GetBlockHashOk

`func (o *ListOnChainSignersByFid200Response) GetBlockHashOk() (*string, bool)`

GetBlockHashOk returns a tuple with the BlockHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockHash

`func (o *ListOnChainSignersByFid200Response) SetBlockHash(v string)`

SetBlockHash sets BlockHash field to given value.

### HasBlockHash

`func (o *ListOnChainSignersByFid200Response) HasBlockHash() bool`

HasBlockHash returns a boolean if a field has been set.

### GetBlockTimestamp

`func (o *ListOnChainSignersByFid200Response) GetBlockTimestamp() int32`

GetBlockTimestamp returns the BlockTimestamp field if non-nil, zero value otherwise.

### GetBlockTimestampOk

`func (o *ListOnChainSignersByFid200Response) GetBlockTimestampOk() (*int32, bool)`

GetBlockTimestampOk returns a tuple with the BlockTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTimestamp

`func (o *ListOnChainSignersByFid200Response) SetBlockTimestamp(v int32)`

SetBlockTimestamp sets BlockTimestamp field to given value.

### HasBlockTimestamp

`func (o *ListOnChainSignersByFid200Response) HasBlockTimestamp() bool`

HasBlockTimestamp returns a boolean if a field has been set.

### GetTransactionHash

`func (o *ListOnChainSignersByFid200Response) GetTransactionHash() string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ListOnChainSignersByFid200Response) GetTransactionHashOk() (*string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ListOnChainSignersByFid200Response) SetTransactionHash(v string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *ListOnChainSignersByFid200Response) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.

### GetLogIndex

`func (o *ListOnChainSignersByFid200Response) GetLogIndex() int32`

GetLogIndex returns the LogIndex field if non-nil, zero value otherwise.

### GetLogIndexOk

`func (o *ListOnChainSignersByFid200Response) GetLogIndexOk() (*int32, bool)`

GetLogIndexOk returns a tuple with the LogIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogIndex

`func (o *ListOnChainSignersByFid200Response) SetLogIndex(v int32)`

SetLogIndex sets LogIndex field to given value.

### HasLogIndex

`func (o *ListOnChainSignersByFid200Response) HasLogIndex() bool`

HasLogIndex returns a boolean if a field has been set.

### GetTxIndex

`func (o *ListOnChainSignersByFid200Response) GetTxIndex() int32`

GetTxIndex returns the TxIndex field if non-nil, zero value otherwise.

### GetTxIndexOk

`func (o *ListOnChainSignersByFid200Response) GetTxIndexOk() (*int32, bool)`

GetTxIndexOk returns a tuple with the TxIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxIndex

`func (o *ListOnChainSignersByFid200Response) SetTxIndex(v int32)`

SetTxIndex sets TxIndex field to given value.

### HasTxIndex

`func (o *ListOnChainSignersByFid200Response) HasTxIndex() bool`

HasTxIndex returns a boolean if a field has been set.

### GetFid

`func (o *ListOnChainSignersByFid200Response) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *ListOnChainSignersByFid200Response) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *ListOnChainSignersByFid200Response) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *ListOnChainSignersByFid200Response) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetSignerEventBody

`func (o *ListOnChainSignersByFid200Response) GetSignerEventBody() SignerEventBody`

GetSignerEventBody returns the SignerEventBody field if non-nil, zero value otherwise.

### GetSignerEventBodyOk

`func (o *ListOnChainSignersByFid200Response) GetSignerEventBodyOk() (*SignerEventBody, bool)`

GetSignerEventBodyOk returns a tuple with the SignerEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerEventBody

`func (o *ListOnChainSignersByFid200Response) SetSignerEventBody(v SignerEventBody)`

SetSignerEventBody sets SignerEventBody field to given value.

### HasSignerEventBody

`func (o *ListOnChainSignersByFid200Response) HasSignerEventBody() bool`

HasSignerEventBody returns a boolean if a field has been set.

### GetEvents

`func (o *ListOnChainSignersByFid200Response) GetEvents() []OnChainEventSigner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListOnChainSignersByFid200Response) GetEventsOk() (*[]OnChainEventSigner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListOnChainSignersByFid200Response) SetEvents(v []OnChainEventSigner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListOnChainSignersByFid200Response) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


