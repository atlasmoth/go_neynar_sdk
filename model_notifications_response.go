/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NotificationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsResponse{}

// NotificationsResponse struct for NotificationsResponse
type NotificationsResponse struct {
	Notifications []Notification `json:"notifications"`
	Next          NextCursor     `json:"next"`
}

type _NotificationsResponse NotificationsResponse

// NewNotificationsResponse instantiates a new NotificationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsResponse(notifications []Notification, next NextCursor) *NotificationsResponse {
	this := NotificationsResponse{}
	this.Notifications = notifications
	this.Next = next
	return &this
}

// NewNotificationsResponseWithDefaults instantiates a new NotificationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsResponseWithDefaults() *NotificationsResponse {
	this := NotificationsResponse{}
	return &this
}

// GetNotifications returns the Notifications field value
func (o *NotificationsResponse) GetNotifications() []Notification {
	if o == nil {
		var ret []Notification
		return ret
	}

	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetNotificationsOk() ([]Notification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notifications, true
}

// SetNotifications sets field value
func (o *NotificationsResponse) SetNotifications(v []Notification) {
	o.Notifications = v
}

// GetNext returns the Next field value
func (o *NotificationsResponse) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *NotificationsResponse) SetNext(v NextCursor) {
	o.Next = v
}

func (o NotificationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notifications"] = o.Notifications
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *NotificationsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"notifications",
		"next",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNotificationsResponse := _NotificationsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotificationsResponse)

	if err != nil {
		return err
	}

	*o = NotificationsResponse(varNotificationsResponse)

	return err
}

type NullableNotificationsResponse struct {
	value *NotificationsResponse
	isSet bool
}

func (v NullableNotificationsResponse) Get() *NotificationsResponse {
	return v.value
}

func (v *NullableNotificationsResponse) Set(val *NotificationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsResponse(val *NotificationsResponse) *NullableNotificationsResponse {
	return &NullableNotificationsResponse{value: val, isSet: true}
}

func (v NullableNotificationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
