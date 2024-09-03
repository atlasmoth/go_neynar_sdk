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

// checks if the ListOnChainEventsByFid200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOnChainEventsByFid200Response{}

// ListOnChainEventsByFid200Response struct for ListOnChainEventsByFid200Response
type ListOnChainEventsByFid200Response struct {
	Events []OnChainEvent `json:"events"`
}

type _ListOnChainEventsByFid200Response ListOnChainEventsByFid200Response

// NewListOnChainEventsByFid200Response instantiates a new ListOnChainEventsByFid200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOnChainEventsByFid200Response(events []OnChainEvent) *ListOnChainEventsByFid200Response {
	this := ListOnChainEventsByFid200Response{}
	this.Events = events
	return &this
}

// NewListOnChainEventsByFid200ResponseWithDefaults instantiates a new ListOnChainEventsByFid200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOnChainEventsByFid200ResponseWithDefaults() *ListOnChainEventsByFid200Response {
	this := ListOnChainEventsByFid200Response{}
	return &this
}

// GetEvents returns the Events field value
func (o *ListOnChainEventsByFid200Response) GetEvents() []OnChainEvent {
	if o == nil {
		var ret []OnChainEvent
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *ListOnChainEventsByFid200Response) GetEventsOk() ([]OnChainEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *ListOnChainEventsByFid200Response) SetEvents(v []OnChainEvent) {
	o.Events = v
}

func (o ListOnChainEventsByFid200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOnChainEventsByFid200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

func (o *ListOnChainEventsByFid200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"events",
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

	varListOnChainEventsByFid200Response := _ListOnChainEventsByFid200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListOnChainEventsByFid200Response)

	if err != nil {
		return err
	}

	*o = ListOnChainEventsByFid200Response(varListOnChainEventsByFid200Response)

	return err
}

type NullableListOnChainEventsByFid200Response struct {
	value *ListOnChainEventsByFid200Response
	isSet bool
}

func (v NullableListOnChainEventsByFid200Response) Get() *ListOnChainEventsByFid200Response {
	return v.value
}

func (v *NullableListOnChainEventsByFid200Response) Set(val *ListOnChainEventsByFid200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListOnChainEventsByFid200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListOnChainEventsByFid200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOnChainEventsByFid200Response(val *ListOnChainEventsByFid200Response) *NullableListOnChainEventsByFid200Response {
	return &NullableListOnChainEventsByFid200Response{value: val, isSet: true}
}

func (v NullableListOnChainEventsByFid200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOnChainEventsByFid200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

