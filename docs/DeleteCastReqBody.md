# DeleteCastReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | **string** | UUID of the signer | 
**TargetHash** | **string** | Ethereum address | 

## Methods

### NewDeleteCastReqBody

`func NewDeleteCastReqBody(signerUuid string, targetHash string, ) *DeleteCastReqBody`

NewDeleteCastReqBody instantiates a new DeleteCastReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteCastReqBodyWithDefaults

`func NewDeleteCastReqBodyWithDefaults() *DeleteCastReqBody`

NewDeleteCastReqBodyWithDefaults instantiates a new DeleteCastReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *DeleteCastReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *DeleteCastReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *DeleteCastReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.


### GetTargetHash

`func (o *DeleteCastReqBody) GetTargetHash() string`

GetTargetHash returns the TargetHash field if non-nil, zero value otherwise.

### GetTargetHashOk

`func (o *DeleteCastReqBody) GetTargetHashOk() (*string, bool)`

GetTargetHashOk returns a tuple with the TargetHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetHash

`func (o *DeleteCastReqBody) SetTargetHash(v string)`

SetTargetHash sets TargetHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


