# LinkRemoveAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**LinkBody** | Pointer to [**LinkBody**](LinkBody.md) |  | [optional] 
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewLinkRemoveAllOfData

`func NewLinkRemoveAllOfData() *LinkRemoveAllOfData`

NewLinkRemoveAllOfData instantiates a new LinkRemoveAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkRemoveAllOfDataWithDefaults

`func NewLinkRemoveAllOfDataWithDefaults() *LinkRemoveAllOfData`

NewLinkRemoveAllOfDataWithDefaults instantiates a new LinkRemoveAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *LinkRemoveAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *LinkRemoveAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *LinkRemoveAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *LinkRemoveAllOfData) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *LinkRemoveAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *LinkRemoveAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *LinkRemoveAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *LinkRemoveAllOfData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *LinkRemoveAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *LinkRemoveAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *LinkRemoveAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *LinkRemoveAllOfData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetLinkBody

`func (o *LinkRemoveAllOfData) GetLinkBody() LinkBody`

GetLinkBody returns the LinkBody field if non-nil, zero value otherwise.

### GetLinkBodyOk

`func (o *LinkRemoveAllOfData) GetLinkBodyOk() (*LinkBody, bool)`

GetLinkBodyOk returns a tuple with the LinkBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBody

`func (o *LinkRemoveAllOfData) SetLinkBody(v LinkBody)`

SetLinkBody sets LinkBody field to given value.

### HasLinkBody

`func (o *LinkRemoveAllOfData) HasLinkBody() bool`

HasLinkBody returns a boolean if a field has been set.

### GetType

`func (o *LinkRemoveAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkRemoveAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkRemoveAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkRemoveAllOfData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


