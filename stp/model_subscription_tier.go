/*
STP API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SubscriptionTier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionTier{}

// SubscriptionTier struct for SubscriptionTier
type SubscriptionTier struct {
	Id *int32 `json:"id,omitempty"`
	Price *SubscriptionTierPrice `json:"price,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionTier SubscriptionTier

// NewSubscriptionTier instantiates a new SubscriptionTier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionTier() *SubscriptionTier {
	this := SubscriptionTier{}
	return &this
}

// NewSubscriptionTierWithDefaults instantiates a new SubscriptionTier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionTierWithDefaults() *SubscriptionTier {
	this := SubscriptionTier{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SubscriptionTier) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionTier) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SubscriptionTier) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SubscriptionTier) SetId(v int32) {
	o.Id = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *SubscriptionTier) GetPrice() SubscriptionTierPrice {
	if o == nil || IsNil(o.Price) {
		var ret SubscriptionTierPrice
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionTier) GetPriceOk() (*SubscriptionTierPrice, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *SubscriptionTier) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given SubscriptionTierPrice and assigns it to the Price field.
func (o *SubscriptionTier) SetPrice(v SubscriptionTierPrice) {
	o.Price = &v
}

func (o SubscriptionTier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionTier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionTier) UnmarshalJSON(data []byte) (err error) {
	varSubscriptionTier := _SubscriptionTier{}

	err = json.Unmarshal(data, &varSubscriptionTier)

	if err != nil {
		return err
	}

	*o = SubscriptionTier(varSubscriptionTier)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "price")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionTier struct {
	value *SubscriptionTier
	isSet bool
}

func (v NullableSubscriptionTier) Get() *SubscriptionTier {
	return v.value
}

func (v *NullableSubscriptionTier) Set(val *SubscriptionTier) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionTier) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionTier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionTier(val *SubscriptionTier) *NullableSubscriptionTier {
	return &NullableSubscriptionTier{value: val, isSet: true}
}

func (v NullableSubscriptionTier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionTier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


