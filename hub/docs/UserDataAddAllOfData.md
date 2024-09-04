# UserDataAddAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to [**FarcasterNetwork**](FarcasterNetwork.md) |  | [optional] [default to FARCASTERNETWORK_MAINNET]
**UserDataBody** | Pointer to [**UserDataBody**](UserDataBody.md) |  | [optional] 
**Type** | Pointer to [**MessageType**](MessageType.md) |  | [optional] [default to MESSAGETYPE_CAST_ADD]

## Methods

### NewUserDataAddAllOfData

`func NewUserDataAddAllOfData() *UserDataAddAllOfData`

NewUserDataAddAllOfData instantiates a new UserDataAddAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataAddAllOfDataWithDefaults

`func NewUserDataAddAllOfDataWithDefaults() *UserDataAddAllOfData`

NewUserDataAddAllOfDataWithDefaults instantiates a new UserDataAddAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *UserDataAddAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *UserDataAddAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *UserDataAddAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *UserDataAddAllOfData) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetTimestamp

`func (o *UserDataAddAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserDataAddAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserDataAddAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserDataAddAllOfData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNetwork

`func (o *UserDataAddAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *UserDataAddAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *UserDataAddAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *UserDataAddAllOfData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUserDataBody

`func (o *UserDataAddAllOfData) GetUserDataBody() UserDataBody`

GetUserDataBody returns the UserDataBody field if non-nil, zero value otherwise.

### GetUserDataBodyOk

`func (o *UserDataAddAllOfData) GetUserDataBodyOk() (*UserDataBody, bool)`

GetUserDataBodyOk returns a tuple with the UserDataBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBody

`func (o *UserDataAddAllOfData) SetUserDataBody(v UserDataBody)`

SetUserDataBody sets UserDataBody field to given value.

### HasUserDataBody

`func (o *UserDataAddAllOfData) HasUserDataBody() bool`

HasUserDataBody returns a boolean if a field has been set.

### GetType

`func (o *UserDataAddAllOfData) GetType() MessageType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDataAddAllOfData) GetTypeOk() (*MessageType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDataAddAllOfData) SetType(v MessageType)`

SetType sets Type field to given value.

### HasType

`func (o *UserDataAddAllOfData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


