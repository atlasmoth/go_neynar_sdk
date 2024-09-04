# VerificationAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**VerificationAddEthAddressBody** | Pointer to [**VerificationAddEthAddressBody**](VerificationAddEthAddressBody.md) |  | [optional] 
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewVerificationAllOfData

`func NewVerificationAllOfData() *VerificationAllOfData`

NewVerificationAllOfData instantiates a new VerificationAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationAllOfDataWithDefaults

`func NewVerificationAllOfDataWithDefaults() *VerificationAllOfData`

NewVerificationAllOfDataWithDefaults instantiates a new VerificationAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *VerificationAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *VerificationAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *VerificationAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *VerificationAllOfData) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *VerificationAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VerificationAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VerificationAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VerificationAllOfData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *VerificationAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VerificationAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VerificationAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VerificationAllOfData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVerificationAddEthAddressBody

`func (o *VerificationAllOfData) GetVerificationAddEthAddressBody() VerificationAddEthAddressBody`

GetVerificationAddEthAddressBody returns the VerificationAddEthAddressBody field if non-nil, zero value otherwise.

### GetVerificationAddEthAddressBodyOk

`func (o *VerificationAllOfData) GetVerificationAddEthAddressBodyOk() (*VerificationAddEthAddressBody, bool)`

GetVerificationAddEthAddressBodyOk returns a tuple with the VerificationAddEthAddressBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAddEthAddressBody

`func (o *VerificationAllOfData) SetVerificationAddEthAddressBody(v VerificationAddEthAddressBody)`

SetVerificationAddEthAddressBody sets VerificationAddEthAddressBody field to given value.

### HasVerificationAddEthAddressBody

`func (o *VerificationAllOfData) HasVerificationAddEthAddressBody() bool`

HasVerificationAddEthAddressBody returns a boolean if a field has been set.

### GetType

`func (o *VerificationAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VerificationAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VerificationAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *VerificationAllOfData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


