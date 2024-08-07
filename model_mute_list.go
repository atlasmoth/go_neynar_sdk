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
	"time"
)

// checks if the MuteList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteList{}

// MuteList struct for MuteList
type MuteList struct {
	Object  string    `json:"object"`
	Muted   User      `json:"muted"`
	MutedAt time.Time `json:"muted_at"`
}

type _MuteList MuteList

// NewMuteList instantiates a new MuteList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteList(object string, muted User, mutedAt time.Time) *MuteList {
	this := MuteList{}
	this.Object = object
	this.Muted = muted
	this.MutedAt = mutedAt
	return &this
}

// NewMuteListWithDefaults instantiates a new MuteList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteListWithDefaults() *MuteList {
	this := MuteList{}
	return &this
}

// GetObject returns the Object field value
func (o *MuteList) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *MuteList) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *MuteList) SetObject(v string) {
	o.Object = v
}

// GetMuted returns the Muted field value
func (o *MuteList) GetMuted() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Muted
}

// GetMutedOk returns a tuple with the Muted field value
// and a boolean to check if the value has been set.
func (o *MuteList) GetMutedOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Muted, true
}

// SetMuted sets field value
func (o *MuteList) SetMuted(v User) {
	o.Muted = v
}

// GetMutedAt returns the MutedAt field value
func (o *MuteList) GetMutedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.MutedAt
}

// GetMutedAtOk returns a tuple with the MutedAt field value
// and a boolean to check if the value has been set.
func (o *MuteList) GetMutedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MutedAt, true
}

// SetMutedAt sets field value
func (o *MuteList) SetMutedAt(v time.Time) {
	o.MutedAt = v
}

func (o MuteList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["muted"] = o.Muted
	toSerialize["muted_at"] = o.MutedAt
	return toSerialize, nil
}

func (o *MuteList) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"muted",
		"muted_at",
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

	varMuteList := _MuteList{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMuteList)

	if err != nil {
		return err
	}

	*o = MuteList(varMuteList)

	return err
}

type NullableMuteList struct {
	value *MuteList
	isSet bool
}

func (v NullableMuteList) Get() *MuteList {
	return v.value
}

func (v *NullableMuteList) Set(val *MuteList) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteList) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteList(val *MuteList) *NullableMuteList {
	return &NullableMuteList{value: val, isSet: true}
}

func (v NullableMuteList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
