# SubscriptionToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**Decimals** | Pointer to **int32** |  | [optional] 
**Erc20** | Pointer to **bool** |  | [optional] 

## Methods

### NewSubscriptionToken

`func NewSubscriptionToken() *SubscriptionToken`

NewSubscriptionToken instantiates a new SubscriptionToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTokenWithDefaults

`func NewSubscriptionTokenWithDefaults() *SubscriptionToken`

NewSubscriptionTokenWithDefaults instantiates a new SubscriptionToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SubscriptionToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SubscriptionToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SubscriptionToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SubscriptionToken) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetAddress

`func (o *SubscriptionToken) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SubscriptionToken) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SubscriptionToken) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SubscriptionToken) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *SubscriptionToken) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *SubscriptionToken) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetDecimals

`func (o *SubscriptionToken) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *SubscriptionToken) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *SubscriptionToken) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *SubscriptionToken) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetErc20

`func (o *SubscriptionToken) GetErc20() bool`

GetErc20 returns the Erc20 field if non-nil, zero value otherwise.

### GetErc20Ok

`func (o *SubscriptionToken) GetErc20Ok() (*bool, bool)`

GetErc20Ok returns a tuple with the Erc20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErc20

`func (o *SubscriptionToken) SetErc20(v bool)`

SetErc20 sets Erc20 field to given value.

### HasErc20

`func (o *SubscriptionToken) HasErc20() bool`

HasErc20 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


