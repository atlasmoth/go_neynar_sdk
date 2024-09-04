# MessageDataUserDataAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**UserDataBody** | Pointer to [**UserDataBody**](UserDataBody.md) |  | [optional] 

## Methods

### NewMessageDataUserDataAdd

`func NewMessageDataUserDataAdd() *MessageDataUserDataAdd`

NewMessageDataUserDataAdd instantiates a new MessageDataUserDataAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataUserDataAddWithDefaults

`func NewMessageDataUserDataAddWithDefaults() *MessageDataUserDataAdd`

NewMessageDataUserDataAddWithDefaults instantiates a new MessageDataUserDataAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataUserDataAdd) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataUserDataAdd) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataUserDataAdd) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataUserDataAdd) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataUserDataAdd) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataUserDataAdd) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataUserDataAdd) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataUserDataAdd) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataUserDataAdd) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataUserDataAdd) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataUserDataAdd) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataUserDataAdd) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUserDataBody

`func (o *MessageDataUserDataAdd) GetUserDataBody() UserDataBody`

GetUserDataBody returns the UserDataBody field if non-nil, zero value otherwise.

### GetUserDataBodyOk

`func (o *MessageDataUserDataAdd) GetUserDataBodyOk() (*UserDataBody, bool)`

GetUserDataBodyOk returns a tuple with the UserDataBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBody

`func (o *MessageDataUserDataAdd) SetUserDataBody(v UserDataBody)`

SetUserDataBody sets UserDataBody field to given value.

### HasUserDataBody

`func (o *MessageDataUserDataAdd) HasUserDataBody() bool`

HasUserDataBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


