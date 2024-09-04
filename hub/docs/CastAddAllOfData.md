# CastAddAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**CastAddBody** | Pointer to [**CastAddBody**](CastAddBody.md) |  | [optional] 
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewCastAddAllOfData

`func NewCastAddAllOfData() *CastAddAllOfData`

NewCastAddAllOfData instantiates a new CastAddAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastAddAllOfDataWithDefaults

`func NewCastAddAllOfDataWithDefaults() *CastAddAllOfData`

NewCastAddAllOfDataWithDefaults instantiates a new CastAddAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *CastAddAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CastAddAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CastAddAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *CastAddAllOfData) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *CastAddAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastAddAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastAddAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CastAddAllOfData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *CastAddAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CastAddAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CastAddAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CastAddAllOfData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCastAddBody

`func (o *CastAddAllOfData) GetCastAddBody() CastAddBody`

GetCastAddBody returns the CastAddBody field if non-nil, zero value otherwise.

### GetCastAddBodyOk

`func (o *CastAddAllOfData) GetCastAddBodyOk() (*CastAddBody, bool)`

GetCastAddBodyOk returns a tuple with the CastAddBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastAddBody

`func (o *CastAddAllOfData) SetCastAddBody(v CastAddBody)`

SetCastAddBody sets CastAddBody field to given value.

### HasCastAddBody

`func (o *CastAddAllOfData) HasCastAddBody() bool`

HasCastAddBody returns a boolean if a field has been set.

### GetType

`func (o *CastAddAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastAddAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastAddAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *CastAddAllOfData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


