# MessageDataLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**LinkBody** | Pointer to [**LinkBody**](LinkBody.md) |  | [optional] 

## Methods

### NewMessageDataLink

`func NewMessageDataLink() *MessageDataLink`

NewMessageDataLink instantiates a new MessageDataLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataLinkWithDefaults

`func NewMessageDataLinkWithDefaults() *MessageDataLink`

NewMessageDataLinkWithDefaults instantiates a new MessageDataLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataLink) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataLink) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataLink) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataLink) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataLink) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataLink) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataLink) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataLink) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataLink) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataLink) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataLink) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataLink) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetLinkBody

`func (o *MessageDataLink) GetLinkBody() LinkBody`

GetLinkBody returns the LinkBody field if non-nil, zero value otherwise.

### GetLinkBodyOk

`func (o *MessageDataLink) GetLinkBodyOk() (*LinkBody, bool)`

GetLinkBodyOk returns a tuple with the LinkBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBody

`func (o *MessageDataLink) SetLinkBody(v LinkBody)`

SetLinkBody sets LinkBody field to given value.

### HasLinkBody

`func (o *MessageDataLink) HasLinkBody() bool`

HasLinkBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


