# MessageDataVerificationRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**VerificationRemoveBody** | Pointer to [**VerificationRemoveBody**](VerificationRemoveBody.md) |  | [optional] 

## Methods

### NewMessageDataVerificationRemove

`func NewMessageDataVerificationRemove() *MessageDataVerificationRemove`

NewMessageDataVerificationRemove instantiates a new MessageDataVerificationRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataVerificationRemoveWithDefaults

`func NewMessageDataVerificationRemoveWithDefaults() *MessageDataVerificationRemove`

NewMessageDataVerificationRemoveWithDefaults instantiates a new MessageDataVerificationRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataVerificationRemove) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataVerificationRemove) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataVerificationRemove) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataVerificationRemove) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataVerificationRemove) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataVerificationRemove) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataVerificationRemove) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataVerificationRemove) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataVerificationRemove) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataVerificationRemove) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataVerificationRemove) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataVerificationRemove) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVerificationRemoveBody

`func (o *MessageDataVerificationRemove) GetVerificationRemoveBody() VerificationRemoveBody`

GetVerificationRemoveBody returns the VerificationRemoveBody field if non-nil, zero value otherwise.

### GetVerificationRemoveBodyOk

`func (o *MessageDataVerificationRemove) GetVerificationRemoveBodyOk() (*VerificationRemoveBody, bool)`

GetVerificationRemoveBodyOk returns a tuple with the VerificationRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRemoveBody

`func (o *MessageDataVerificationRemove) SetVerificationRemoveBody(v VerificationRemoveBody)`

SetVerificationRemoveBody sets VerificationRemoveBody field to given value.

### HasVerificationRemoveBody

`func (o *MessageDataVerificationRemove) HasVerificationRemoveBody() bool`

HasVerificationRemoveBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


