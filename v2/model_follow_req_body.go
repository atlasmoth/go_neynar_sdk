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

// checks if the FollowReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowReqBody{}

// FollowReqBody struct for FollowReqBody
type FollowReqBody struct {
	// UUID of the signer
	SignerUuid *string `json:"signer_uuid,omitempty"`
	TargetFids []int32 `json:"target_fids,omitempty"`
}

// NewFollowReqBody instantiates a new FollowReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowReqBody() *FollowReqBody {
	this := FollowReqBody{}
	return &this
}

// NewFollowReqBodyWithDefaults instantiates a new FollowReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowReqBodyWithDefaults() *FollowReqBody {
	this := FollowReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value if set, zero value otherwise.
func (o *FollowReqBody) GetSignerUuid() string {
	if o == nil || IsNil(o.SignerUuid) {
		var ret string
		return ret
	}
	return *o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.SignerUuid) {
		return nil, false
	}
	return o.SignerUuid, true
}

// HasSignerUuid returns a boolean if a field has been set.
func (o *FollowReqBody) HasSignerUuid() bool {
	if o != nil && !IsNil(o.SignerUuid) {
		return true
	}

	return false
}

// SetSignerUuid gets a reference to the given string and assigns it to the SignerUuid field.
func (o *FollowReqBody) SetSignerUuid(v string) {
	o.SignerUuid = &v
}

// GetTargetFids returns the TargetFids field value if set, zero value otherwise.
func (o *FollowReqBody) GetTargetFids() []int32 {
	if o == nil || IsNil(o.TargetFids) {
		var ret []int32
		return ret
	}
	return o.TargetFids
}

// GetTargetFidsOk returns a tuple with the TargetFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowReqBody) GetTargetFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TargetFids) {
		return nil, false
	}
	return o.TargetFids, true
}

// HasTargetFids returns a boolean if a field has been set.
func (o *FollowReqBody) HasTargetFids() bool {
	if o != nil && !IsNil(o.TargetFids) {
		return true
	}

	return false
}

// SetTargetFids gets a reference to the given []int32 and assigns it to the TargetFids field.
func (o *FollowReqBody) SetTargetFids(v []int32) {
	o.TargetFids = v
}

func (o FollowReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignerUuid) {
		toSerialize["signer_uuid"] = o.SignerUuid
	}
	if !IsNil(o.TargetFids) {
		toSerialize["target_fids"] = o.TargetFids
	}
	return toSerialize, nil
}

type NullableFollowReqBody struct {
	value *FollowReqBody
	isSet bool
}

func (v NullableFollowReqBody) Get() *FollowReqBody {
	return v.value
}

func (v *NullableFollowReqBody) Set(val *FollowReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowReqBody(val *FollowReqBody) *NullableFollowReqBody {
	return &NullableFollowReqBody{value: val, isSet: true}
}

func (v NullableFollowReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


