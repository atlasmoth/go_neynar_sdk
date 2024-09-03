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

// checks if the ValidateMessageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateMessageResponse{}

// ValidateMessageResponse struct for ValidateMessageResponse
type ValidateMessageResponse struct {
	Valid bool `json:"valid"`
	Message Message `json:"message"`
}

type _ValidateMessageResponse ValidateMessageResponse

// NewValidateMessageResponse instantiates a new ValidateMessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateMessageResponse(valid bool, message Message) *ValidateMessageResponse {
	this := ValidateMessageResponse{}
	this.Valid = valid
	this.Message = message
	return &this
}

// NewValidateMessageResponseWithDefaults instantiates a new ValidateMessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateMessageResponseWithDefaults() *ValidateMessageResponse {
	this := ValidateMessageResponse{}
	return &this
}

// GetValid returns the Valid field value
func (o *ValidateMessageResponse) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *ValidateMessageResponse) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *ValidateMessageResponse) SetValid(v bool) {
	o.Valid = v
}

// GetMessage returns the Message field value
func (o *ValidateMessageResponse) GetMessage() Message {
	if o == nil {
		var ret Message
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ValidateMessageResponse) GetMessageOk() (*Message, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ValidateMessageResponse) SetMessage(v Message) {
	o.Message = v
}

func (o ValidateMessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateMessageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valid"] = o.Valid
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ValidateMessageResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"valid",
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

	varValidateMessageResponse := _ValidateMessageResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValidateMessageResponse)

	if err != nil {
		return err
	}

	*o = ValidateMessageResponse(varValidateMessageResponse)

	return err
}

type NullableValidateMessageResponse struct {
	value *ValidateMessageResponse
	isSet bool
}

func (v NullableValidateMessageResponse) Get() *ValidateMessageResponse {
	return v.value
}

func (v *NullableValidateMessageResponse) Set(val *ValidateMessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateMessageResponse(val *ValidateMessageResponse) *NullableValidateMessageResponse {
	return &NullableValidateMessageResponse{value: val, isSet: true}
}

func (v NullableValidateMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


