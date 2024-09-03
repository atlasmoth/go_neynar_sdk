# VerificationRemoveAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** |  | 
**Timestamp** | **int64** |  | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_MAINNET]
**VerificationRemoveBody** | [**VerificationRemoveBody**](VerificationRemoveBody.md) |  | 
**Type** | [**MessageType**](MessageType.md) |  | [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewVerificationRemoveAllOfData

`func NewVerificationRemoveAllOfData(fid int32, timestamp int64, network FarcasterNetwork, verificationRemoveBody VerificationRemoveBody, type_ MessageType, ) *VerificationRemoveAllOfData`

NewVerificationRemoveAllOfData instantiates a new VerificationRemoveAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRemoveAllOfDataWithDefaults

`func NewVerificationRemoveAllOfDataWithDefaults() *VerificationRemoveAllOfData`

NewVerificationRemoveAllOfDataWithDefaults instantiates a new VerificationRemoveAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *VerificationRemoveAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *VerificationRemoveAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *VerificationRemoveAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *VerificationRemoveAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VerificationRemoveAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VerificationRemoveAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *VerificationRemoveAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VerificationRemoveAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VerificationRemoveAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetVerificationRemoveBody

`func (o *VerificationRemoveAllOfData) GetVerificationRemoveBody() VerificationRemoveBody`

GetVerificationRemoveBody returns the VerificationRemoveBody field if non-nil, zero value otherwise.

### GetVerificationRemoveBodyOk

`func (o *VerificationRemoveAllOfData) GetVerificationRemoveBodyOk() (*VerificationRemoveBody, bool)`

GetVerificationRemoveBodyOk returns a tuple with the VerificationRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRemoveBody

`func (o *VerificationRemoveAllOfData) SetVerificationRemoveBody(v VerificationRemoveBody)`

SetVerificationRemoveBody sets VerificationRemoveBody field to given value.


### GetType

`func (o *VerificationRemoveAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VerificationRemoveAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VerificationRemoveAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

