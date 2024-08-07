/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SubscriptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionsResponse{}

// SubscriptionsResponse struct for SubscriptionsResponse
type SubscriptionsResponse struct {
	SubscriptionsCreated []Subscriptions `json:"subscriptions_created,omitempty"`
}

// NewSubscriptionsResponse instantiates a new SubscriptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionsResponse() *SubscriptionsResponse {
	this := SubscriptionsResponse{}
	return &this
}

// NewSubscriptionsResponseWithDefaults instantiates a new SubscriptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsResponseWithDefaults() *SubscriptionsResponse {
	this := SubscriptionsResponse{}
	return &this
}

// GetSubscriptionsCreated returns the SubscriptionsCreated field value if set, zero value otherwise.
func (o *SubscriptionsResponse) GetSubscriptionsCreated() []Subscriptions {
	if o == nil || IsNil(o.SubscriptionsCreated) {
		var ret []Subscriptions
		return ret
	}
	return o.SubscriptionsCreated
}

// GetSubscriptionsCreatedOk returns a tuple with the SubscriptionsCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionsResponse) GetSubscriptionsCreatedOk() ([]Subscriptions, bool) {
	if o == nil || IsNil(o.SubscriptionsCreated) {
		return nil, false
	}
	return o.SubscriptionsCreated, true
}

// HasSubscriptionsCreated returns a boolean if a field has been set.
func (o *SubscriptionsResponse) HasSubscriptionsCreated() bool {
	if o != nil && !IsNil(o.SubscriptionsCreated) {
		return true
	}

	return false
}

// SetSubscriptionsCreated gets a reference to the given []Subscriptions and assigns it to the SubscriptionsCreated field.
func (o *SubscriptionsResponse) SetSubscriptionsCreated(v []Subscriptions) {
	o.SubscriptionsCreated = v
}

func (o SubscriptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionsCreated) {
		toSerialize["subscriptions_created"] = o.SubscriptionsCreated
	}
	return toSerialize, nil
}

type NullableSubscriptionsResponse struct {
	value *SubscriptionsResponse
	isSet bool
}

func (v NullableSubscriptionsResponse) Get() *SubscriptionsResponse {
	return v.value
}

func (v *NullableSubscriptionsResponse) Set(val *SubscriptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionsResponse(val *SubscriptionsResponse) *NullableSubscriptionsResponse {
	return &NullableSubscriptionsResponse{value: val, isSet: true}
}

func (v NullableSubscriptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

