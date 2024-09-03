# CastRemoveAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** |  | 
**Timestamp** | **int64** |  | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_MAINNET]
**CastRemoveBody** | [**CastRemoveBody**](CastRemoveBody.md) |  | 
**Type** | [**MessageType**](MessageType.md) |  | [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewCastRemoveAllOfData

`func NewCastRemoveAllOfData(fid int32, timestamp int64, network FarcasterNetwork, castRemoveBody CastRemoveBody, type_ MessageType, ) *CastRemoveAllOfData`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


