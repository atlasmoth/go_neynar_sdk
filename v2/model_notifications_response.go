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

// checks if the NotificationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsResponse{}

// NotificationsResponse struct for NotificationsResponse
type NotificationsResponse struct {
	UnseenNotificationsCount *int32 `json:"unseen_notifications_count,omitempty"`
	Notifications []Notification `json:"notifications,omitempty"`
	Next *NextCursor `json:"next,omitempty"`
}

// NewNotificationsResponse instantiates a new NotificationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsResponse() *NotificationsResponse {
	this := NotificationsResponse{}
	return &this
}

// NewNotificationsResponseWithDefaults instantiates a new NotificationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsResponseWithDefaults() *NotificationsResponse {
	this := NotificationsResponse{}
	return &this
}

// GetUnseenNotificationsCount returns the UnseenNotificationsCount field value if set, zero value otherwise.
func (o *NotificationsResponse) GetUnseenNotificationsCount() int32 {
	if o == nil || IsNil(o.UnseenNotificationsCount) {
		var ret int32
		return ret
	}
	return *o.UnseenNotificationsCount
}

// GetUnseenNotificationsCountOk returns a tuple with the UnseenNotificationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetUnseenNotificationsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UnseenNotificationsCount) {
		return nil, false
	}
	return o.UnseenNotificationsCount, true
}

// HasUnseenNotificationsCount returns a boolean if a field has been set.
func (o *NotificationsResponse) HasUnseenNotificationsCount() bool {
	if o != nil && !IsNil(o.UnseenNotificationsCount) {
		return true
	}

	return false
}

// SetUnseenNotificationsCount gets a reference to the given int32 and assigns it to the UnseenNotificationsCount field.
func (o *NotificationsResponse) SetUnseenNotificationsCount(v int32) {
	o.UnseenNotificationsCount = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *NotificationsResponse) GetNotifications() []Notification {
	if o == nil || IsNil(o.Notifications) {
		var ret []Notification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetNotificationsOk() ([]Notification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *NotificationsResponse) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *NotificationsResponse) SetNotifications(v []Notification) {
	o.Notifications = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *NotificationsResponse) GetNext() NextCursor {
	if o == nil || IsNil(o.Next) {
		var ret NextCursor
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsResponse) GetNextOk() (*NextCursor, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *NotificationsResponse) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextCursor and assigns it to the Next field.
func (o *NotificationsResponse) SetNext(v NextCursor) {
	o.Next = &v
}

func (o NotificationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnseenNotificationsCount) {
		toSerialize["unseen_notifications_count"] = o.UnseenNotificationsCount
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
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


