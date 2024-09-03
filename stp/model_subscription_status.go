/*
STP API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SubscriptionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionStatus{}

// SubscriptionStatus struct for SubscriptionStatus
type SubscriptionStatus struct {
	Object string `json:"object"`
	Status bool `json:"status"`
	ExpiresAt NullableInt64 `json:"expires_at"`
	SubscribedAt NullableInt64 `json:"subscribed_at"`
	Tier SubscriptionTier `json:"tier"`
}

type _SubscriptionStatus SubscriptionStatus

// NewSubscriptionStatus instantiates a new SubscriptionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionStatus(object string, status bool, expiresAt NullableInt64, subscribedAt NullableInt64, tier SubscriptionTier) *SubscriptionStatus {
	this := SubscriptionStatus{}
	this.Object = object
	this.Status = status
	this.ExpiresAt = expiresAt
	this.SubscribedAt = subscribedAt
	this.Tier = tier
	return &this
}

// NewSubscriptionStatusWithDefaults instantiates a new SubscriptionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionStatusWithDefaults() *SubscriptionStatus {
	this := SubscriptionStatus{}
	return &this
}

// GetObject returns the Object field value
func (o *SubscriptionStatus) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SubscriptionStatus) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SubscriptionStatus) SetObject(v string) {
	o.Object = v
}

// GetStatus returns the Status field value
func (o *SubscriptionStatus) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SubscriptionStatus) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SubscriptionStatus) SetStatus(v bool) {
	o.Status = v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *SubscriptionStatus) GetExpiresAt() int64 {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionStatus) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *SubscriptionStatus) SetExpiresAt(v int64) {
	o.ExpiresAt.Set(&v)
}

// GetSubscribedAt returns the SubscribedAt field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *SubscriptionStatus) GetSubscribedAt() int64 {
	if o == nil || o.SubscribedAt.Get() == nil {
		var ret int64
		return ret
	}

	return *o.SubscribedAt.Get()
}

// GetSubscribedAtOk returns a tuple with the SubscribedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubscriptionStatus) GetSubscribedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscribedAt.Get(), o.SubscribedAt.IsSet()
}

// SetSubscribedAt sets field value
func (o *SubscriptionStatus) SetSubscribedAt(v int64) {
	o.SubscribedAt.Set(&v)
}

// GetTier returns the Tier field value
func (o *SubscriptionStatus) GetTier() SubscriptionTier {
	if o == nil {
		var ret SubscriptionTier
		return ret
	}

	return o.Tier
}

// GetTierOk returns a tuple with the Tier field value
// and a boolean to check if the value has been set.
func (o *SubscriptionStatus) GetTierOk() (*SubscriptionTier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tier, true
}

// SetTier sets field value
func (o *SubscriptionStatus) SetTier(v SubscriptionTier) {
	o.Tier = v
}

func (o SubscriptionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["status"] = o.Status
	toSerialize["expires_at"] = o.ExpiresAt.Get()
	toSerialize["subscribed_at"] = o.SubscribedAt.Get()
	toSerialize["tier"] = o.Tier
	return toSerialize, nil
}

func (o *SubscriptionStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"status",
		"expires_at",
		"subscribed_at",
		"tier",
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

	varSubscriptionStatus := _SubscriptionStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionStatus)

	if err != nil {
		return err
	}

	*o = SubscriptionStatus(varSubscriptionStatus)

	return err
}

type NullableSubscriptionStatus struct {
	value *SubscriptionStatus
	isSet bool
}

func (v NullableSubscriptionStatus) Get() *SubscriptionStatus {
	return v.value
}

func (v *NullableSubscriptionStatus) Set(val *SubscriptionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionStatus(val *SubscriptionStatus) *NullableSubscriptionStatus {
	return &NullableSubscriptionStatus{value: val, isSet: true}
}

func (v NullableSubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


