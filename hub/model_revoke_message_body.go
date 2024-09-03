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

// checks if the RevokeMessageBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RevokeMessageBody{}

// RevokeMessageBody struct for RevokeMessageBody
type RevokeMessageBody struct {
	Message Message `json:"message"`
}

type _RevokeMessageBody RevokeMessageBody

// NewRevokeMessageBody instantiates a new RevokeMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRevokeMessageBody(message Message) *RevokeMessageBody {
	this := RevokeMessageBody{}
	this.Message = message
	return &this
}

// NewRevokeMessageBodyWithDefaults instantiates a new RevokeMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRevokeMessageBodyWithDefaults() *RevokeMessageBody {
	this := RevokeMessageBody{}
	return &this
}

// GetMessage returns the Message field value
func (o *RevokeMessageBody) GetMessage() Message {
	if o == nil {
		var ret Message
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *RevokeMessageBody) GetMessageOk() (*Message, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *RevokeMessageBody) SetMessage(v Message) {
	o.Message = v
}

func (o RevokeMessageBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RevokeMessageBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *RevokeMessageBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varRevokeMessageBody := _RevokeMessageBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRevokeMessageBody)

	if err != nil {
		return err
	}

	*o = RevokeMessageBody(varRevokeMessageBody)

	return err
}

type NullableRevokeMessageBody struct {
	value *RevokeMessageBody
	isSet bool
}

func (v NullableRevokeMessageBody) Get() *RevokeMessageBody {
	return v.value
}

func (v *NullableRevokeMessageBody) Set(val *RevokeMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokeMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokeMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokeMessageBody(val *RevokeMessageBody) *NullableRevokeMessageBody {
	return &NullableRevokeMessageBody{value: val, isSet: true}
}

func (v NullableRevokeMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokeMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


