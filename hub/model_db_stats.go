/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DbStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbStats{}

// DbStats struct for DbStats
type DbStats struct {
	NumMessages *int32 `json:"numMessages,omitempty"`
	NumFidEvents *int32 `json:"numFidEvents,omitempty"`
	NumFnameEvents *int32 `json:"numFnameEvents,omitempty"`
}

// NewDbStats instantiates a new DbStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbStats() *DbStats {
	this := DbStats{}
	return &this
}

// NewDbStatsWithDefaults instantiates a new DbStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbStatsWithDefaults() *DbStats {
	this := DbStats{}
	return &this
}

// GetNumMessages returns the NumMessages field value if set, zero value otherwise.
func (o *DbStats) GetNumMessages() int32 {
	if o == nil || IsNil(o.NumMessages) {
		var ret int32
		return ret
	}
	return *o.NumMessages
}

// GetNumMessagesOk returns a tuple with the NumMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbStats) GetNumMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumMessages) {
		return nil, false
	}
	return o.NumMessages, true
}

// HasNumMessages returns a boolean if a field has been set.
func (o *DbStats) HasNumMessages() bool {
	if o != nil && !IsNil(o.NumMessages) {
		return true
	}

	return false
}

// SetNumMessages gets a reference to the given int32 and assigns it to the NumMessages field.
func (o *DbStats) SetNumMessages(v int32) {
	o.NumMessages = &v
}

// GetNumFidEvents returns the NumFidEvents field value if set, zero value otherwise.
func (o *DbStats) GetNumFidEvents() int32 {
	if o == nil || IsNil(o.NumFidEvents) {
		var ret int32
		return ret
	}
	return *o.NumFidEvents
}

// GetNumFidEventsOk returns a tuple with the NumFidEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbStats) GetNumFidEventsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumFidEvents) {
		return nil, false
	}
	return o.NumFidEvents, true
}

// HasNumFidEvents returns a boolean if a field has been set.
func (o *DbStats) HasNumFidEvents() bool {
	if o != nil && !IsNil(o.NumFidEvents) {
		return true
	}

	return false
}

// SetNumFidEvents gets a reference to the given int32 and assigns it to the NumFidEvents field.
func (o *DbStats) SetNumFidEvents(v int32) {
	o.NumFidEvents = &v
}

// GetNumFnameEvents returns the NumFnameEvents field value if set, zero value otherwise.
func (o *DbStats) GetNumFnameEvents() int32 {
	if o == nil || IsNil(o.NumFnameEvents) {
		var ret int32
		return ret
	}
	return *o.NumFnameEvents
}

// GetNumFnameEventsOk returns a tuple with the NumFnameEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DbStats) GetNumFnameEventsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumFnameEvents) {
		return nil, false
	}
	return o.NumFnameEvents, true
}

// HasNumFnameEvents returns a boolean if a field has been set.
func (o *DbStats) HasNumFnameEvents() bool {
	if o != nil && !IsNil(o.NumFnameEvents) {
		return true
	}

	return false
}

// SetNumFnameEvents gets a reference to the given int32 and assigns it to the NumFnameEvents field.
func (o *DbStats) SetNumFnameEvents(v int32) {
	o.NumFnameEvents = &v
}

func (o DbStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DbStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumMessages) {
		toSerialize["numMessages"] = o.NumMessages
	}
	if !IsNil(o.NumFidEvents) {
		toSerialize["numFidEvents"] = o.NumFidEvents
	}
	if !IsNil(o.NumFnameEvents) {
		toSerialize["numFnameEvents"] = o.NumFnameEvents
	}
	return toSerialize, nil
}

type NullableDbStats struct {
	value *DbStats
	isSet bool
}

func (v NullableDbStats) Get() *DbStats {
	return v.value
}

func (v *NullableDbStats) Set(val *DbStats) {
	v.value = val
	v.isSet = true
}

func (v NullableDbStats) IsSet() bool {
	return v.isSet
}

func (v *NullableDbStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbStats(val *DbStats) *NullableDbStats {
	return &NullableDbStats{value: val, isSet: true}
}

func (v NullableDbStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


