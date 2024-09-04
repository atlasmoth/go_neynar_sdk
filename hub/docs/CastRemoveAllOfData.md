# CastRemoveAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**CastRemoveBody** | Pointer to [**CastRemoveBody**](CastRemoveBody.md) |  | [optional] 
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewCastRemoveAllOfData

`func NewCastRemoveAllOfData() *CastRemoveAllOfData`

NewCastRemoveAllOfData instantiates a new CastRemoveAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastRemoveAllOfDataWithDefaults

`func NewCastRemoveAllOfDataWithDefaults() *CastRemoveAllOfData`

NewCastRemoveAllOfDataWithDefaults instantiates a new CastRemoveAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *CastRemoveAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CastRemoveAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CastRemoveAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *CastRemoveAllOfData) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *CastRemoveAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastRemoveAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastRemoveAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CastRemoveAllOfData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *CastRemoveAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CastRemoveAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CastRemoveAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CastRemoveAllOfData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCastRemoveBody

`func (o *CastRemoveAllOfData) GetCastRemoveBody() CastRemoveBody`

GetCastRemoveBody returns the CastRemoveBody field if non-nil, zero value otherwise.

### GetCastRemoveBodyOk

`func (o *CastRemoveAllOfData) GetCastRemoveBodyOk() (*CastRemoveBody, bool)`

GetCastRemoveBodyOk returns a tuple with the CastRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRemoveBody

`func (o *CastRemoveAllOfData) SetCastRemoveBody(v CastRemoveBody)`

SetCastRemoveBody sets CastRemoveBody field to given value.

### HasCastRemoveBody

`func (o *CastRemoveAllOfData) HasCastRemoveBody() bool`

HasCastRemoveBody returns a boolean if a field has been set.

### GetType

`func (o *CastRemoveAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastRemoveAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastRemoveAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *CastRemoveAllOfData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


