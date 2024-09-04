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

// checks if the ChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelResponse{}

// ChannelResponse struct for ChannelResponse
type ChannelResponse struct {
	Channel *Channel `json:"channel,omitempty"`
}

// NewChannelResponse instantiates a new ChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelResponse() *ChannelResponse {
	this := ChannelResponse{}
	return &this
}

// NewChannelResponseWithDefaults instantiates a new ChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelResponseWithDefaults() *ChannelResponse {
	this := ChannelResponse{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ChannelResponse) GetChannel() Channel {
	if o == nil || IsNil(o.Channel) {
		var ret Channel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelResponse) GetChannelOk() (*Channel, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ChannelResponse) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given Channel and assigns it to the Channel field.
func (o *ChannelResponse) SetChannel(v Channel) {
	o.Channel = &v
}

func (o ChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	return toSerialize, nil
}

type NullableChannelResponse struct {
	value *ChannelResponse
	isSet bool
}

func (v NullableChannelResponse) Get() *ChannelResponse {
	return v.value
}

func (v *NullableChannelResponse) Set(val *ChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelResponse(val *ChannelResponse) *NullableChannelResponse {
	return &NullableChannelResponse{value: val, isSet: true}
}

func (v NullableChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


