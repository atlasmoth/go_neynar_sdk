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

// checks if the ListEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListEvents200Response{}

// ListEvents200Response struct for ListEvents200Response
type ListEvents200Response struct {
	NextPageEventId *int32 `json:"nextPageEventId,omitempty"`
	Events []HubEvent `json:"events,omitempty"`
}

// NewListEvents200Response instantiates a new ListEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEvents200Response() *ListEvents200Response {
	this := ListEvents200Response{}
	return &this
}

// NewListEvents200ResponseWithDefaults instantiates a new ListEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEvents200ResponseWithDefaults() *ListEvents200Response {
	this := ListEvents200Response{}
	return &this
}

// GetNextPageEventId returns the NextPageEventId field value if set, zero value otherwise.
func (o *ListEvents200Response) GetNextPageEventId() int32 {
	if o == nil || IsNil(o.NextPageEventId) {
		var ret int32
		return ret
	}
	return *o.NextPageEventId
}

// GetNextPageEventIdOk returns a tuple with the NextPageEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEvents200Response) GetNextPageEventIdOk() (*int32, bool) {
	if o == nil || IsNil(o.NextPageEventId) {
		return nil, false
	}
	return o.NextPageEventId, true
}

// HasNextPageEventId returns a boolean if a field has been set.
func (o *ListEvents200Response) HasNextPageEventId() bool {
	if o != nil && !IsNil(o.NextPageEventId) {
		return true
	}

	return false
}

// SetNextPageEventId gets a reference to the given int32 and assigns it to the NextPageEventId field.
func (o *ListEvents200Response) SetNextPageEventId(v int32) {
	o.NextPageEventId = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ListEvents200Response) GetEvents() []HubEvent {
	if o == nil || IsNil(o.Events) {
		var ret []HubEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListEvents200Response) GetEventsOk() ([]HubEvent, bool) {
	if o == nil || IsNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ListEvents200Response) HasEvents() bool {
	if o != nil && !IsNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []HubEvent and assigns it to the Events field.
func (o *ListEvents200Response) SetEvents(v []HubEvent) {
	o.Events = v
}

func (o ListEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPageEventId) {
		toSerialize["nextPageEventId"] = o.NextPageEventId
	}
	if !IsNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	return toSerialize, nil
}

type NullableListEvents200Response struct {
	value *ListEvents200Response
	isSet bool
}

func (v NullableListEvents200Response) Get() *ListEvents200Response {
	return v.value
}

func (v *NullableListEvents200Response) Set(val *ListEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEvents200Response(val *ListEvents200Response) *NullableListEvents200Response {
	return &NullableListEvents200Response{value: val, isSet: true}
}

func (v NullableListEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


