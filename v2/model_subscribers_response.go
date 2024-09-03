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

// checks if the SubscribersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscribersResponse{}

// SubscribersResponse struct for SubscribersResponse
type SubscribersResponse struct {
	Subscribers []Subscriber `json:"subscribers,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SubscribersResponse SubscribersResponse

// NewSubscribersResponse instantiates a new SubscribersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribersResponse() *SubscribersResponse {
	this := SubscribersResponse{}
	return &this
}

// NewSubscribersResponseWithDefaults instantiates a new SubscribersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribersResponseWithDefaults() *SubscribersResponse {
	this := SubscribersResponse{}
	return &this
}

// GetSubscribers returns the Subscribers field value if set, zero value otherwise.
func (o *SubscribersResponse) GetSubscribers() []Subscriber {
	if o == nil || IsNil(o.Subscribers) {
		var ret []Subscriber
		return ret
	}
	return o.Subscribers
}

// GetSubscribersOk returns a tuple with the Subscribers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribersResponse) GetSubscribersOk() ([]Subscriber, bool) {
	if o == nil || IsNil(o.Subscribers) {
		return nil, false
	}
	return o.Subscribers, true
}

// HasSubscribers returns a boolean if a field has been set.
func (o *SubscribersResponse) HasSubscribers() bool {
	if o != nil && !IsNil(o.Subscribers) {
		return true
	}

	return false
}

// SetSubscribers gets a reference to the given []Subscriber and assigns it to the Subscribers field.
func (o *SubscribersResponse) SetSubscribers(v []Subscriber) {
	o.Subscribers = v
}

func (o SubscribersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscribersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscribers) {
		toSerialize["subscribers"] = o.Subscribers
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscribersResponse) UnmarshalJSON(data []byte) (err error) {
	varSubscribersResponse := _SubscribersResponse{}

	err = json.Unmarshal(data, &varSubscribersResponse)

	if err != nil {
		return err
	}

	*o = SubscribersResponse(varSubscribersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subscribers")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscribersResponse struct {
	value *SubscribersResponse
	isSet bool
}

func (v NullableSubscribersResponse) Get() *SubscribersResponse {
	return v.value
}

func (v *NullableSubscribersResponse) Set(val *SubscribersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribersResponse(val *SubscribersResponse) *NullableSubscribersResponse {
	return &NullableSubscribersResponse{value: val, isSet: true}
}

func (v NullableSubscribersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


