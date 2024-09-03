/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DbStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DbStats{}

// DbStats struct for DbStats
type DbStats struct {
	NumMessages int32 `json:"numMessages"`
	NumFidEvents int32 `json:"numFidEvents"`
	NumFnameEvents int32 `json:"numFnameEvents"`
}

type _DbStats DbStats

// NewDbStats instantiates a new DbStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbStats(numMessages int32, numFidEvents int32, numFnameEvents int32) *DbStats {
	this := DbStats{}
	this.NumMessages = numMessages
	this.NumFidEvents = numFidEvents
	this.NumFnameEvents = numFnameEvents
	return &this
}

// NewDbStatsWithDefaults instantiates a new DbStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbStatsWithDefaults() *DbStats {
	this := DbStats{}
	return &this
}

// GetNumMessages returns the NumMessages field value
func (o *DbStats) GetNumMessages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumMessages
}

// GetNumMessagesOk returns a tuple with the NumMessages field value
// and a boolean to check if the value has been set.
func (o *DbStats) GetNumMessagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumMessages, true
}

// SetNumMessages sets field value
func (o *DbStats) SetNumMessages(v int32) {
	o.NumMessages = v
}

// GetNumFidEvents returns the NumFidEvents field value
func (o *DbStats) GetNumFidEvents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumFidEvents
}

// GetNumFidEventsOk returns a tuple with the NumFidEvents field value
// and a boolean to check if the value has been set.
func (o *DbStats) GetNumFidEventsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumFidEvents, true
}

// SetNumFidEvents sets field value
func (o *DbStats) SetNumFidEvents(v int32) {
	o.NumFidEvents = v
}

// GetNumFnameEvents returns the NumFnameEvents field value
func (o *DbStats) GetNumFnameEvents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumFnameEvents
}

// GetNumFnameEventsOk returns a tuple with the NumFnameEvents field value
// and a boolean to check if the value has been set.
func (o *DbStats) GetNumFnameEventsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumFnameEvents, true
}

// SetNumFnameEvents sets field value
func (o *DbStats) SetNumFnameEvents(v int32) {
	o.NumFnameEvents = v
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
	toSerialize["numMessages"] = o.NumMessages
	toSerialize["numFidEvents"] = o.NumFidEvents
	toSerialize["numFnameEvents"] = o.NumFnameEvents
	return toSerialize, nil
}

func (o *DbStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"numMessages",
		"numFidEvents",
		"numFnameEvents",
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

	varDbStats := _DbStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDbStats)

	if err != nil {
		return err
	}

	*o = DbStats(varDbStats)

	return err
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

