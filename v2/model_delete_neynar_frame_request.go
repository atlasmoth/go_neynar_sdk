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

// checks if the DeleteNeynarFrameRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteNeynarFrameRequest{}

// DeleteNeynarFrameRequest struct for DeleteNeynarFrameRequest
type DeleteNeynarFrameRequest struct {
	Uuid *string `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteNeynarFrameRequest DeleteNeynarFrameRequest

// NewDeleteNeynarFrameRequest instantiates a new DeleteNeynarFrameRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNeynarFrameRequest() *DeleteNeynarFrameRequest {
	this := DeleteNeynarFrameRequest{}
	return &this
}

// NewDeleteNeynarFrameRequestWithDefaults instantiates a new DeleteNeynarFrameRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNeynarFrameRequestWithDefaults() *DeleteNeynarFrameRequest {
	this := DeleteNeynarFrameRequest{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeleteNeynarFrameRequest) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteNeynarFrameRequest) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeleteNeynarFrameRequest) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeleteNeynarFrameRequest) SetUuid(v string) {
	o.Uuid = &v
}

func (o DeleteNeynarFrameRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteNeynarFrameRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteNeynarFrameRequest) UnmarshalJSON(data []byte) (err error) {
	varDeleteNeynarFrameRequest := _DeleteNeynarFrameRequest{}

	err = json.Unmarshal(data, &varDeleteNeynarFrameRequest)

	if err != nil {
		return err
	}

	*o = DeleteNeynarFrameRequest(varDeleteNeynarFrameRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteNeynarFrameRequest struct {
	value *DeleteNeynarFrameRequest
	isSet bool
}

func (v NullableDeleteNeynarFrameRequest) Get() *DeleteNeynarFrameRequest {
	return v.value
}

func (v *NullableDeleteNeynarFrameRequest) Set(val *DeleteNeynarFrameRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNeynarFrameRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNeynarFrameRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNeynarFrameRequest(val *DeleteNeynarFrameRequest) *NullableDeleteNeynarFrameRequest {
	return &NullableDeleteNeynarFrameRequest{value: val, isSet: true}
}

func (v NullableDeleteNeynarFrameRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteNeynarFrameRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


