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

// checks if the WebhookSubscriptionFiltersFollow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSubscriptionFiltersFollow{}

// WebhookSubscriptionFiltersFollow struct for WebhookSubscriptionFiltersFollow
type WebhookSubscriptionFiltersFollow struct {
	Fids []int32 `json:"fids,omitempty"`
	TargetFids []int32 `json:"target_fids,omitempty"`
}

// NewWebhookSubscriptionFiltersFollow instantiates a new WebhookSubscriptionFiltersFollow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSubscriptionFiltersFollow() *WebhookSubscriptionFiltersFollow {
	this := WebhookSubscriptionFiltersFollow{}
	return &this
}

// NewWebhookSubscriptionFiltersFollowWithDefaults instantiates a new WebhookSubscriptionFiltersFollow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSubscriptionFiltersFollowWithDefaults() *WebhookSubscriptionFiltersFollow {
	this := WebhookSubscriptionFiltersFollow{}
	return &this
}

// GetFids returns the Fids field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersFollow) GetFids() []int32 {
	if o == nil || IsNil(o.Fids) {
		var ret []int32
		return ret
	}
	return o.Fids
}

// GetFidsOk returns a tuple with the Fids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersFollow) GetFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Fids) {
		return nil, false
	}
	return o.Fids, true
}

// HasFids returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersFollow) HasFids() bool {
	if o != nil && !IsNil(o.Fids) {
		return true
	}

	return false
}

// SetFids gets a reference to the given []int32 and assigns it to the Fids field.
func (o *WebhookSubscriptionFiltersFollow) SetFids(v []int32) {
	o.Fids = v
}

// GetTargetFids returns the TargetFids field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersFollow) GetTargetFids() []int32 {
	if o == nil || IsNil(o.TargetFids) {
		var ret []int32
		return ret
	}
	return o.TargetFids
}

// GetTargetFidsOk returns a tuple with the TargetFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersFollow) GetTargetFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TargetFids) {
		return nil, false
	}
	return o.TargetFids, true
}

// HasTargetFids returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersFollow) HasTargetFids() bool {
	if o != nil && !IsNil(o.TargetFids) {
		return true
	}

	return false
}

// SetTargetFids gets a reference to the given []int32 and assigns it to the TargetFids field.
func (o *WebhookSubscriptionFiltersFollow) SetTargetFids(v []int32) {
	o.TargetFids = v
}

func (o WebhookSubscriptionFiltersFollow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSubscriptionFiltersFollow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fids) {
		toSerialize["fids"] = o.Fids
	}
	if !IsNil(o.TargetFids) {
		toSerialize["target_fids"] = o.TargetFids
	}
	return toSerialize, nil
}

type NullableWebhookSubscriptionFiltersFollow struct {
	value *WebhookSubscriptionFiltersFollow
	isSet bool
}

func (v NullableWebhookSubscriptionFiltersFollow) Get() *WebhookSubscriptionFiltersFollow {
	return v.value
}

func (v *NullableWebhookSubscriptionFiltersFollow) Set(val *WebhookSubscriptionFiltersFollow) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSubscriptionFiltersFollow) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSubscriptionFiltersFollow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSubscriptionFiltersFollow(val *WebhookSubscriptionFiltersFollow) *NullableWebhookSubscriptionFiltersFollow {
	return &NullableWebhookSubscriptionFiltersFollow{value: val, isSet: true}
}

func (v NullableWebhookSubscriptionFiltersFollow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSubscriptionFiltersFollow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

