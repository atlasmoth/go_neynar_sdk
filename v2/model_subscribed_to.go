/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the SubscribedTo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscribedTo{}

// SubscribedTo struct for SubscribedTo
type SubscribedTo struct {
	Object string `json:"object"`
	ProviderName *string `json:"provider_name,omitempty"`
	ContractAddress string `json:"contract_address"`
	Chain int32 `json:"chain"`
	Metadata SubscriptionMetadata `json:"metadata"`
	OwnerAddress string `json:"owner_address"`
	Price SubscriptionPrice `json:"price"`
	Tiers []SubscriptionTier `json:"tiers,omitempty"`
	ProtocolVersion int32 `json:"protocol_version"`
	Token SubscriptionToken `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	SubscribedAt time.Time `json:"subscribed_at"`
	Tier SubscriptionTier `json:"tier"`
	Creator User `json:"creator"`
	AdditionalProperties map[string]interface{}
}

type _SubscribedTo SubscribedTo

// NewSubscribedTo instantiates a new SubscribedTo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribedTo(object string, contractAddress string, chain int32, metadata SubscriptionMetadata, ownerAddress string, price SubscriptionPrice, protocolVersion int32, token SubscriptionToken, expiresAt time.Time, subscribedAt time.Time, tier SubscriptionTier, creator User) *SubscribedTo {
	this := SubscribedTo{}
	this.Object = object
	this.ContractAddress = contractAddress
	this.Chain = chain
	this.Metadata = metadata
	this.OwnerAddress = ownerAddress
	this.Price = price
	this.ProtocolVersion = protocolVersion
	this.Token = token
	this.ExpiresAt = expiresAt
	this.SubscribedAt = subscribedAt
	this.Tier = tier
	this.Creator = creator
	return &this
}

// NewSubscribedToWithDefaults instantiates a new SubscribedTo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribedToWithDefaults() *SubscribedTo {
	this := SubscribedTo{}
	return &this
}

// GetObject returns the Object field value
func (o *SubscribedTo) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SubscribedTo) SetObject(v string) {
	o.Object = v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *SubscribedTo) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *SubscribedTo) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *SubscribedTo) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetContractAddress returns the ContractAddress field value
func (o *SubscribedTo) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *SubscribedTo) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetChain returns the Chain field value
func (o *SubscribedTo) GetChain() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Chain
}

// GetChainOk returns a tuple with the Chain field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetChainOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chain, true
}

// SetChain sets field value
func (o *SubscribedTo) SetChain(v int32) {
	o.Chain = v
}

// GetMetadata returns the Metadata field value
func (o *SubscribedTo) GetMetadata() SubscriptionMetadata {
	if o == nil {
		var ret SubscriptionMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetMetadataOk() (*SubscriptionMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SubscribedTo) SetMetadata(v SubscriptionMetadata) {
	o.Metadata = v
}

// GetOwnerAddress returns the OwnerAddress field value
func (o *SubscribedTo) GetOwnerAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetOwnerAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerAddress, true
}

// SetOwnerAddress sets field value
func (o *SubscribedTo) SetOwnerAddress(v string) {
	o.OwnerAddress = v
}

// GetPrice returns the Price field value
func (o *SubscribedTo) GetPrice() SubscriptionPrice {
	if o == nil {
		var ret SubscriptionPrice
		return ret
	}

	return o.Price
}

// GetPriceOk returns a tuple with the Price field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetPriceOk() (*SubscriptionPrice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Price, true
}

// SetPrice sets field value
func (o *SubscribedTo) SetPrice(v SubscriptionPrice) {
	o.Price = v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *SubscribedTo) GetTiers() []SubscriptionTier {
	if o == nil || IsNil(o.Tiers) {
		var ret []SubscriptionTier
		return ret
	}
	return o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetTiersOk() ([]SubscriptionTier, bool) {
	if o == nil || IsNil(o.Tiers) {
		return nil, false
	}
	return o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *SubscribedTo) HasTiers() bool {
	if o != nil && !IsNil(o.Tiers) {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []SubscriptionTier and assigns it to the Tiers field.
func (o *SubscribedTo) SetTiers(v []SubscriptionTier) {
	o.Tiers = v
}

// GetProtocolVersion returns the ProtocolVersion field value
func (o *SubscribedTo) GetProtocolVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProtocolVersion
}

// GetProtocolVersionOk returns a tuple with the ProtocolVersion field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetProtocolVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtocolVersion, true
}

// SetProtocolVersion sets field value
func (o *SubscribedTo) SetProtocolVersion(v int32) {
	o.ProtocolVersion = v
}

// GetToken returns the Token field value
func (o *SubscribedTo) GetToken() SubscriptionToken {
	if o == nil {
		var ret SubscriptionToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetTokenOk() (*SubscriptionToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SubscribedTo) SetToken(v SubscriptionToken) {
	o.Token = v
}

// GetExpiresAt returns the ExpiresAt field value
func (o *SubscribedTo) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value
func (o *SubscribedTo) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetSubscribedAt returns the SubscribedAt field value
func (o *SubscribedTo) GetSubscribedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SubscribedAt
}

// GetSubscribedAtOk returns a tuple with the SubscribedAt field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetSubscribedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscribedAt, true
}

// SetSubscribedAt sets field value
func (o *SubscribedTo) SetSubscribedAt(v time.Time) {
	o.SubscribedAt = v
}

// GetTier returns the Tier field value
func (o *SubscribedTo) GetTier() SubscriptionTier {
	if o == nil {
		var ret SubscriptionTier
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetTierOk() (*SubscriptionTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *SubscribedTo) SetTier(v SubscriptionTier) {
	o.Tier = v
}

// GetCreator returns the Creator field value
func (o *SubscribedTo) GetCreator() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
func (o *SubscribedTo) GetCreatorOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creator, true
}

// SetCreator sets field value
func (o *SubscribedTo) SetCreator(v User) {
	o.Creator = v
}

func (o SubscribedTo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribedTo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	if !IsNil(o.ProviderName) {
		toSerialize["provider_name"] = o.ProviderName
	}
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["chain"] = o.Chain
	toSerialize["metadata"] = o.Metadata
	toSerialize["owner_address"] = o.OwnerAddress
	toSerialize["price"] = o.Price
	if !IsNil(o.Tiers) {
		toSerialize["tiers"] = o.Tiers
	}
	toSerialize["protocol_version"] = o.ProtocolVersion
	toSerialize["token"] = o.Token
	toSerialize["expires_at"] = o.ExpiresAt
	toSerialize["subscribed_at"] = o.SubscribedAt
	toSerialize["tier"] = o.Tier
	toSerialize["creator"] = o.Creator

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribedTo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"contract_address",
		"chain",
		"metadata",
		"owner_address",
		"price",
		"protocol_version",
		"token",
		"expires_at",
		"subscribed_at",
		"tier",
		"creator",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSubscribedTo := _SubscribedTo{}

	err = json.Unmarshal(data, &varSubscribedTo)

	if err != nil {
		return err
	}

	*o = SubscribedTo(varSubscribedTo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "provider_name")
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "chain")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "owner_address")
		delete(additionalProperties, "price")
		delete(additionalProperties, "tiers")
		delete(additionalProperties, "protocol_version")
		delete(additionalProperties, "token")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "subscribed_at")
		delete(additionalProperties, "tier")
		delete(additionalProperties, "creator")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribedTo struct {
	value *SubscribedTo
	isSet bool
}

func (v NullableSubscribedTo) Get() *SubscribedTo {
	return v.value
}

func (v *NullableSubscribedTo) Set(val *SubscribedTo) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedTo) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedTo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedTo(val *SubscribedTo) *NullableSubscribedTo {
	return &NullableSubscribedTo{value: val, isSet: true}
}

func (v NullableSubscribedTo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedTo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


