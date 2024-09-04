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

// checks if the BulkFollowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkFollowResponse{}

// BulkFollowResponse struct for BulkFollowResponse
type BulkFollowResponse struct {
	Success *bool `json:"success,omitempty"`
	Details []FollowResponse `json:"details,omitempty"`
}

// NewBulkFollowResponse instantiates a new BulkFollowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkFollowResponse() *BulkFollowResponse {
	this := BulkFollowResponse{}
	return &this
}

// NewBulkFollowResponseWithDefaults instantiates a new BulkFollowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkFollowResponseWithDefaults() *BulkFollowResponse {
	this := BulkFollowResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkFollowResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkFollowResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkFollowResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *BulkFollowResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BulkFollowResponse) GetDetails() []FollowResponse {
	if o == nil || IsNil(o.Details) {
		var ret []FollowResponse
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkFollowResponse) GetDetailsOk() ([]FollowResponse, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BulkFollowResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []FollowResponse and assigns it to the Details field.
func (o *BulkFollowResponse) SetDetails(v []FollowResponse) {
	o.Details = v
}

func (o BulkFollowResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkFollowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableBulkFollowResponse struct {
	value *BulkFollowResponse
	isSet bool
}

func (v NullableBulkFollowResponse) Get() *BulkFollowResponse {
	return v.value
}

func (v *NullableBulkFollowResponse) Set(val *BulkFollowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkFollowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkFollowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkFollowResponse(val *BulkFollowResponse) *NullableBulkFollowResponse {
	return &NullableBulkFollowResponse{value: val, isSet: true}
}

func (v NullableBulkFollowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkFollowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


