# FrameActionReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | Pointer to **string** | UUID of the signer | [optional] 
**CastHash** | Pointer to **string** | Cast Hash | [optional] [default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be"]
**Action** | Pointer to [**FrameAction**](FrameAction.md) |  | [optional] 

## Methods

### NewFrameActionReqBody

`func NewFrameActionReqBody() *FrameActionReqBody`

NewFrameActionReqBody instantiates a new FrameActionReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameActionReqBodyWithDefaults

`func NewFrameActionReqBodyWithDefaults() *FrameActionReqBody`

NewFrameActionReqBodyWithDefaults instantiates a new FrameActionReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *FrameActionReqBody) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *FrameActionReqBody) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *FrameActionReqBody) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.

### HasSignerUuid

`func (o *FrameActionReqBody) HasSignerUuid() bool`

HasSignerUuid returns a boolean if a field has been set.

### GetCastHash

`func (o *FrameActionReqBody) GetCastHash() string`

GetCastHash returns the CastHash field if non-nil, zero value otherwise.

### GetCastHashOk

`func (o *FrameActionReqBody) GetCastHashOk() (*string, bool)`

GetCastHashOk returns a tuple with the CastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastHash

`func (o *FrameActionReqBody) SetCastHash(v string)`

SetCastHash sets CastHash field to given value.

### HasCastHash

`func (o *FrameActionReqBody) HasCastHash() bool`

HasCastHash returns a boolean if a field has been set.

### GetAction

`func (o *FrameActionReqBody) GetAction() FrameAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FrameActionReqBody) GetActionOk() (*FrameAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FrameActionReqBody) SetAction(v FrameAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *FrameActionReqBody) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


