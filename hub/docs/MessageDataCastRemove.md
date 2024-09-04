# MessageDataCastRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**CastRemoveBody** | Pointer to [**CastRemoveBody**](CastRemoveBody.md) |  | [optional] 

## Methods

### NewMessageDataCastRemove

`func NewMessageDataCastRemove() *MessageDataCastRemove`

NewMessageDataCastRemove instantiates a new MessageDataCastRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataCastRemoveWithDefaults

`func NewMessageDataCastRemoveWithDefaults() *MessageDataCastRemove`

NewMessageDataCastRemoveWithDefaults instantiates a new MessageDataCastRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataCastRemove) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataCastRemove) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataCastRemove) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataCastRemove) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataCastRemove) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataCastRemove) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataCastRemove) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataCastRemove) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataCastRemove) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataCastRemove) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataCastRemove) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataCastRemove) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCastRemoveBody

`func (o *MessageDataCastRemove) GetCastRemoveBody() CastRemoveBody`

GetCastRemoveBody returns the CastRemoveBody field if non-nil, zero value otherwise.

### GetCastRemoveBodyOk

`func (o *MessageDataCastRemove) GetCastRemoveBodyOk() (*CastRemoveBody, bool)`

GetCastRemoveBodyOk returns a tuple with the CastRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRemoveBody

`func (o *MessageDataCastRemove) SetCastRemoveBody(v CastRemoveBody)`

SetCastRemoveBody sets CastRemoveBody field to given value.

### HasCastRemoveBody

`func (o *MessageDataCastRemove) HasCastRemoveBody() bool`

HasCastRemoveBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


