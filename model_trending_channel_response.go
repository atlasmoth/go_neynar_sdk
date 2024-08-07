/*
Farcaster API V2

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

// checks if the TrendingChannelResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrendingChannelResponse{}

// TrendingChannelResponse struct for TrendingChannelResponse
type TrendingChannelResponse struct {
	Channels []ChannelActivity `json:"channels"`
	Next NextCursor `json:"next"`
}

type _TrendingChannelResponse TrendingChannelResponse

// NewTrendingChannelResponse instantiates a new TrendingChannelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrendingChannelResponse(channels []ChannelActivity, next NextCursor) *TrendingChannelResponse {
	this := TrendingChannelResponse{}
	this.Channels = channels
	this.Next = next
	return &this
}

// NewTrendingChannelResponseWithDefaults instantiates a new TrendingChannelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrendingChannelResponseWithDefaults() *TrendingChannelResponse {
	this := TrendingChannelResponse{}
	return &this
}

// GetChannels returns the Channels field value
func (o *TrendingChannelResponse) GetChannels() []ChannelActivity {
	if o == nil {
		var ret []ChannelActivity
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *TrendingChannelResponse) GetChannelsOk() ([]ChannelActivity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Channels, true
}

// SetChannels sets field value
func (o *TrendingChannelResponse) SetChannels(v []ChannelActivity) {
	o.Channels = v
}

// GetNext returns the Next field value
func (o *TrendingChannelResponse) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *TrendingChannelResponse) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *TrendingChannelResponse) SetNext(v NextCursor) {
	o.Next = v
}

func (o TrendingChannelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrendingChannelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channels"] = o.Channels
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *TrendingChannelResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channels",
		"next",
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

	varTrendingChannelResponse := _TrendingChannelResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTrendingChannelResponse)

	if err != nil {
		return err
	}

	*o = TrendingChannelResponse(varTrendingChannelResponse)

	return err
}

type NullableTrendingChannelResponse struct {
	value *TrendingChannelResponse
	isSet bool
}

func (v NullableTrendingChannelResponse) Get() *TrendingChannelResponse {
	return v.value
}

func (v *NullableTrendingChannelResponse) Set(val *TrendingChannelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendingChannelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendingChannelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendingChannelResponse(val *TrendingChannelResponse) *NullableTrendingChannelResponse {
	return &NullableTrendingChannelResponse{value: val, isSet: true}
}

func (v NullableTrendingChannelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrendingChannelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


