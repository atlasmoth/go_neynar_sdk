# MessageDataReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**ReactionBody** | Pointer to [**ReactionBody**](ReactionBody.md) |  | [optional] 

## Methods

### NewMessageDataReaction

`func NewMessageDataReaction() *MessageDataReaction`

NewMessageDataReaction instantiates a new MessageDataReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataReactionWithDefaults

`func NewMessageDataReactionWithDefaults() *MessageDataReaction`

NewMessageDataReactionWithDefaults instantiates a new MessageDataReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataReaction) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataReaction) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataReaction) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataReaction) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataReaction) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataReaction) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataReaction) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataReaction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataReaction) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataReaction) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataReaction) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataReaction) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetReactionBody

`func (o *MessageDataReaction) GetReactionBody() ReactionBody`

GetReactionBody returns the ReactionBody field if non-nil, zero value otherwise.

### GetReactionBodyOk

`func (o *MessageDataReaction) GetReactionBodyOk() (*ReactionBody, bool)`

GetReactionBodyOk returns a tuple with the ReactionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionBody

`func (o *MessageDataReaction) SetReactionBody(v ReactionBody)`

SetReactionBody sets ReactionBody field to given value.

### HasReactionBody

`func (o *MessageDataReaction) HasReactionBody() bool`

HasReactionBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


