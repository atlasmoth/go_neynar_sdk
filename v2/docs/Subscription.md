# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**ProviderName** | Pointer to **string** |  | [optional] 
**ContractAddress** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to [**SubscriptionMetadata**](SubscriptionMetadata.md) |  | [optional] 
**OwnerAddress** | Pointer to **string** |  | [optional] 
**Price** | Pointer to [**SubscriptionPrice**](SubscriptionPrice.md) |  | [optional] 
**Tiers** | Pointer to [**[]SubscriptionTier**](SubscriptionTier.md) |  | [optional] 
**ProtocolVersion** | Pointer to **int32** |  | [optional] 
**Token** | Pointer to [**SubscriptionToken**](SubscriptionToken.md) |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Subscription) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Subscription) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Subscription) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Subscription) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetProviderName

`func (o *Subscription) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *Subscription) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *Subscription) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *Subscription) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### GetContractAddress

`func (o *Subscription) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *Subscription) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *Subscription) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *Subscription) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetChain

`func (o *Subscription) GetChain() int32`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Subscription) GetChainOk() (*int32, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Subscription) SetChain(v int32)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Subscription) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetMetadata

`func (o *Subscription) GetMetadata() SubscriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Subscription) GetMetadataOk() (*SubscriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Subscription) SetMetadata(v SubscriptionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Subscription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *Subscription) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *Subscription) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *Subscription) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *Subscription) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetPrice

`func (o *Subscription) GetPrice() SubscriptionPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Subscription) GetPriceOk() (*SubscriptionPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Subscription) SetPrice(v SubscriptionPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Subscription) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTiers

`func (o *Subscription) GetTiers() []SubscriptionTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *Subscription) GetTiersOk() (*[]SubscriptionTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *Subscription) SetTiers(v []SubscriptionTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *Subscription) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetProtocolVersion

`func (o *Subscription) GetProtocolVersion() int32`

GetProtocolVersion returns the ProtocolVersion field if non-nil, zero value otherwise.

### GetProtocolVersionOk

`func (o *Subscription) GetProtocolVersionOk() (*int32, bool)`

GetProtocolVersionOk returns a tuple with the ProtocolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolVersion

`func (o *Subscription) SetProtocolVersion(v int32)`

SetProtocolVersion sets ProtocolVersion field to given value.

### HasProtocolVersion

`func (o *Subscription) HasProtocolVersion() bool`

HasProtocolVersion returns a boolean if a field has been set.

### GetToken

`func (o *Subscription) GetToken() SubscriptionToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Subscription) GetTokenOk() (*SubscriptionToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Subscription) SetToken(v SubscriptionToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *Subscription) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


