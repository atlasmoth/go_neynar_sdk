# FollowReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | Pointer to **string** | UUID of the signer | [optional] 
**TargetFids** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewFollowReqBody

`func NewFollowReqBody() *FollowReqBody`

NewFollowReqBody instantiates a new FollowReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowReqBodyWithDefaults

`func NewFollowReqBodyWithDefaults() *FollowReqBody`

NewFollowReqBodyWithDefaults instantiates a new FollowReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *FollowReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *FollowReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *FollowReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.

### HasSignerUuid

`func (o *FollowReqBody) HasSignerUuid() bool`

HasSignerUuid returns a boolean if a field has been set.

### GetTargetFids

`func (o *FollowReqBody) GetTargetFids() []int32`

GetTargetFids returns the TargetFids field if non-nil, zero value otherwise.

### GetTargetFidsOk

`func (o *FollowReqBody) GetTargetFidsOk() (*[]int32, bool)`

GetTargetFidsOk returns a tuple with the TargetFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFids

`func (o *FollowReqBody) SetTargetFids(v []int32)`

SetTargetFids sets TargetFids field to given value.

### HasTargetFids

`func (o *FollowReqBody) HasTargetFids() bool`

HasTargetFids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


