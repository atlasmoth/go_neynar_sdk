# MessageDataCastAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**CastAddBody** | Pointer to [**CastAddBody**](CastAddBody.md) |  | [optional] 

## Methods

### NewMessageDataCastAdd

`func NewMessageDataCastAdd() *MessageDataCastAdd`

NewMessageDataCastAdd instantiates a new MessageDataCastAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataCastAddWithDefaults

`func NewMessageDataCastAddWithDefaults() *MessageDataCastAdd`

NewMessageDataCastAddWithDefaults instantiates a new MessageDataCastAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataCastAdd) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataCastAdd) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataCastAdd) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataCastAdd) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataCastAdd) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataCastAdd) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataCastAdd) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataCastAdd) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataCastAdd) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataCastAdd) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataCastAdd) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataCastAdd) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCastAddBody

`func (o *MessageDataCastAdd) GetCastAddBody() CastAddBody`

GetCastAddBody returns the CastAddBody field if non-nil, zero value otherwise.

### GetCastAddBodyOk

`func (o *MessageDataCastAdd) GetCastAddBodyOk() (*CastAddBody, bool)`

GetCastAddBodyOk returns a tuple with the CastAddBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastAddBody

`func (o *MessageDataCastAdd) SetCastAddBody(v CastAddBody)`

SetCastAddBody sets CastAddBody field to given value.

### HasCastAddBody

`func (o *MessageDataCastAdd) HasCastAddBody() bool`

HasCastAddBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


