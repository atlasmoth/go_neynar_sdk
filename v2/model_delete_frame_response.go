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

// checks if the DeleteFrameResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteFrameResponse{}

// DeleteFrameResponse struct for DeleteFrameResponse
type DeleteFrameResponse struct {
	Success *bool `json:"success,omitempty"`
	Uuid *string `json:"uuid,omitempty"`
}

// NewDeleteFrameResponse instantiates a new DeleteFrameResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteFrameResponse() *DeleteFrameResponse {
	this := DeleteFrameResponse{}
	return &this
}

// NewDeleteFrameResponseWithDefaults instantiates a new DeleteFrameResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteFrameResponseWithDefaults() *DeleteFrameResponse {
	this := DeleteFrameResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteFrameResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFrameResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteFrameResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteFrameResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *DeleteFrameResponse) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteFrameResponse) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *DeleteFrameResponse) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *DeleteFrameResponse) SetUuid(v string) {
	o.Uuid = &v
}

func (o DeleteFrameResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteFrameResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableDeleteFrameResponse struct {
	value *DeleteFrameResponse
	isSet bool
}

func (v NullableDeleteFrameResponse) Get() *DeleteFrameResponse {
	return v.value
}

func (v *NullableDeleteFrameResponse) Set(val *DeleteFrameResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteFrameResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteFrameResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteFrameResponse(val *DeleteFrameResponse) *NullableDeleteFrameResponse {
	return &NullableDeleteFrameResponse{value: val, isSet: true}
}

func (v NullableDeleteFrameResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteFrameResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

