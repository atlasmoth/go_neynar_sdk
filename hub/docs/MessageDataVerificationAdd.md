# MessageDataVerificationAdd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**VerificationAddEthAddressBody** | Pointer to [**VerificationAddEthAddressBody**](VerificationAddEthAddressBody.md) |  | [optional] 

## Methods

### NewMessageDataVerificationAdd

`func NewMessageDataVerificationAdd() *MessageDataVerificationAdd`

NewMessageDataVerificationAdd instantiates a new MessageDataVerificationAdd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageDataVerificationAddWithDefaults

`func NewMessageDataVerificationAddWithDefaults() *MessageDataVerificationAdd`

NewMessageDataVerificationAddWithDefaults instantiates a new MessageDataVerificationAdd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageDataVerificationAdd) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageDataVerificationAdd) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageDataVerificationAdd) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *MessageDataVerificationAdd) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *MessageDataVerificationAdd) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageDataVerificationAdd) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageDataVerificationAdd) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MessageDataVerificationAdd) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *MessageDataVerificationAdd) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageDataVerificationAdd) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageDataVerificationAdd) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MessageDataVerificationAdd) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVerificationAddEthAddressBody

`func (o *MessageDataVerificationAdd) GetVerificationAddEthAddressBody() VerificationAddEthAddressBody`

GetVerificationAddEthAddressBody returns the VerificationAddEthAddressBody field if non-nil, zero value otherwise.

### GetVerificationAddEthAddressBodyOk

`func (o *MessageDataVerificationAdd) GetVerificationAddEthAddressBodyOk() (*VerificationAddEthAddressBody, bool)`

GetVerificationAddEthAddressBodyOk returns a tuple with the VerificationAddEthAddressBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAddEthAddressBody

`func (o *MessageDataVerificationAdd) SetVerificationAddEthAddressBody(v VerificationAddEthAddressBody)`

SetVerificationAddEthAddressBody sets VerificationAddEthAddressBody field to given value.

### HasVerificationAddEthAddressBody

`func (o *MessageDataVerificationAdd) HasVerificationAddEthAddressBody() bool`

HasVerificationAddEthAddressBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


