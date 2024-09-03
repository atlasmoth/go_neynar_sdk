/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the SubscriptionToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionToken{}

// SubscriptionToken struct for SubscriptionToken
type SubscriptionToken struct {
	Symbol string `json:"symbol"`
	Address NullableString `json:"address"`
	Decimals int32 `json:"decimals"`
	Erc20 bool `json:"erc20"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionToken SubscriptionToken

// NewSubscriptionToken instantiates a new SubscriptionToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionToken(symbol string, address NullableString, decimals int32, erc20 bool) *SubscriptionToken {
	this := SubscriptionToken{}
	this.Symbol = symbol
	this.Address = address
	this.Decimals = decimals
	this.Erc20 = erc20
	return &this
}

// NewSubscriptionTokenWithDefaults instantiates a new SubscriptionToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionTokenWithDefaults() *SubscriptionToken {
	this := SubscriptionToken{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SubscriptionToken) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SubscriptionToken) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SubscriptionToken) SetSymbol(v string) {
	o.Symbol = v
}

// GetAddress returns the Address field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SubscriptionToken) GetAddress() string {
	if o == nil || o.Address.Get() == nil {
		var ret string
		return ret
	}

	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionToken) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// SetAddress sets field value
func (o *SubscriptionToken) SetAddress(v string) {
	o.Address.Set(&v)
}

// GetDecimals returns the Decimals field value
func (o *SubscriptionToken) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *SubscriptionToken) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *SubscriptionToken) SetDecimals(v int32) {
	o.Decimals = v
}

// GetErc20 returns the Erc20 field value
func (o *SubscriptionToken) GetErc20() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Erc20
}

// GetErc20Ok returns a tuple with the Erc20 field value
// and a boolean to check if the value has been set.
func (o *SubscriptionToken) GetErc20Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Erc20, true
}

// SetErc20 sets field value
func (o *SubscriptionToken) SetErc20(v bool) {
	o.Erc20 = v
}

func (o SubscriptionToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["address"] = o.Address.Get()
	toSerialize["decimals"] = o.Decimals
	toSerialize["erc20"] = o.Erc20

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"address",
		"decimals",
		"erc20",
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

	varSubscriptionToken := _SubscriptionToken{}

	err = json.Unmarshal(data, &varSubscriptionToken)

	if err != nil {
		return err
	}

	*o = SubscriptionToken(varSubscriptionToken)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "address")
		delete(additionalProperties, "decimals")
		delete(additionalProperties, "erc20")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionToken struct {
	value *SubscriptionToken
	isSet bool
}

func (v NullableSubscriptionToken) Get() *SubscriptionToken {
	return v.value
}

func (v *NullableSubscriptionToken) Set(val *SubscriptionToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionToken(val *SubscriptionToken) *NullableSubscriptionToken {
	return &NullableSubscriptionToken{value: val, isSet: true}
}

func (v NullableSubscriptionToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


