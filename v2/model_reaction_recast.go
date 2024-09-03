/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ReactionRecast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionRecast{}

// ReactionRecast struct for ReactionRecast
type ReactionRecast struct {
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	Fname string `json:"fname"`
	AdditionalProperties map[string]interface{}
}

type _ReactionRecast ReactionRecast

// NewReactionRecast instantiates a new ReactionRecast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionRecast(fid int32, fname string) *ReactionRecast {
	this := ReactionRecast{}
	this.Fid = fid
	this.Fname = fname
	return &this
}

// NewReactionRecastWithDefaults instantiates a new ReactionRecast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionRecastWithDefaults() *ReactionRecast {
	this := ReactionRecast{}
	return &this
}

// GetFid returns the Fid field value
func (o *ReactionRecast) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *ReactionRecast) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *ReactionRecast) SetFid(v int32) {
	o.Fid = v
}

// GetFname returns the Fname field value
func (o *ReactionRecast) GetFname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fname
}

// GetFnameOk returns a tuple with the Fname field value
// and a boolean to check if the value has been set.
func (o *ReactionRecast) GetFnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fname, true
}

// SetFname sets field value
func (o *ReactionRecast) SetFname(v string) {
	o.Fname = v
}

func (o ReactionRecast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionRecast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["fname"] = o.Fname

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReactionRecast) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"fname",
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

	varReactionRecast := _ReactionRecast{}

	err = json.Unmarshal(data, &varReactionRecast)

	if err != nil {
		return err
	}

	*o = ReactionRecast(varReactionRecast)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fid")
		delete(additionalProperties, "fname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReactionRecast struct {
	value *ReactionRecast
	isSet bool
}

func (v NullableReactionRecast) Get() *ReactionRecast {
	return v.value
}

func (v *NullableReactionRecast) Set(val *ReactionRecast) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionRecast) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionRecast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionRecast(val *ReactionRecast) *NullableReactionRecast {
	return &NullableReactionRecast{value: val, isSet: true}
}

func (v NullableReactionRecast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionRecast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


