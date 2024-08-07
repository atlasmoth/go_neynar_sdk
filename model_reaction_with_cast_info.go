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

// checks if the ReactionWithCastInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionWithCastInfo{}

// ReactionWithCastInfo struct for ReactionWithCastInfo
type ReactionWithCastInfo struct {
	ReactionType      string               `json:"reaction_type"`
	Cast              CastWithInteractions `json:"cast"`
	ReactionTimestamp time.Time            `json:"reaction_timestamp"`
	Object            string               `json:"object"`
	User              UserDehydrated       `json:"user"`
}

type _ReactionWithCastInfo ReactionWithCastInfo

// NewReactionWithCastInfo instantiates a new ReactionWithCastInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionWithCastInfo(reactionType string, cast CastWithInteractions, reactionTimestamp time.Time, object string, user UserDehydrated) *ReactionWithCastInfo {
	this := ReactionWithCastInfo{}
	this.ReactionType = reactionType
	this.Cast = cast
	this.ReactionTimestamp = reactionTimestamp
	this.Object = object
	this.User = user
	return &this
}

// NewReactionWithCastInfoWithDefaults instantiates a new ReactionWithCastInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionWithCastInfoWithDefaults() *ReactionWithCastInfo {
	this := ReactionWithCastInfo{}
	return &this
}

// GetReactionType returns the ReactionType field value
func (o *ReactionWithCastInfo) GetReactionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReactionType
}

// GetReactionTypeOk returns a tuple with the ReactionType field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastInfo) GetReactionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactionType, true
}

// SetReactionType sets field value
func (o *ReactionWithCastInfo) SetReactionType(v string) {
	o.ReactionType = v
}

// GetCast returns the Cast field value
func (o *ReactionWithCastInfo) GetCast() CastWithInteractions {
	if o == nil {
		var ret CastWithInteractions
		return ret
	}

	return o.Cast
}

// GetCastOk returns a tuple with the Cast field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastInfo) GetCastOk() (*CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cast, true
}

// SetCast sets field value
func (o *ReactionWithCastInfo) SetCast(v CastWithInteractions) {
	o.Cast = v
}

// GetReactionTimestamp returns the ReactionTimestamp field value
func (o *ReactionWithCastInfo) GetReactionTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReactionTimestamp
}

// GetReactionTimestampOk returns a tuple with the ReactionTimestamp field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastInfo) GetReactionTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReactionTimestamp, true
}

// SetReactionTimestamp sets field value
func (o *ReactionWithCastInfo) SetReactionTimestamp(v time.Time) {
	o.ReactionTimestamp = v
}

// GetObject returns the Object field value
func (o *ReactionWithCastInfo) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastInfo) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ReactionWithCastInfo) SetObject(v string) {
	o.Object = v
}

// GetUser returns the User field value
func (o *ReactionWithCastInfo) GetUser() UserDehydrated {
	if o == nil {
		var ret UserDehydrated
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ReactionWithCastInfo) GetUserOk() (*UserDehydrated, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ReactionWithCastInfo) SetUser(v UserDehydrated) {
	o.User = v
}

func (o ReactionWithCastInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionWithCastInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reaction_type"] = o.ReactionType
	toSerialize["cast"] = o.Cast
	toSerialize["reaction_timestamp"] = o.ReactionTimestamp
	toSerialize["object"] = o.Object
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *ReactionWithCastInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reaction_type",
		"cast",
		"reaction_timestamp",
		"object",
		"user",
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

	varReactionWithCastInfo := _ReactionWithCastInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReactionWithCastInfo)

	if err != nil {
		return err
	}

	*o = ReactionWithCastInfo(varReactionWithCastInfo)

	return err
}

type NullableReactionWithCastInfo struct {
	value *ReactionWithCastInfo
	isSet bool
}

func (v NullableReactionWithCastInfo) Get() *ReactionWithCastInfo {
	return v.value
}

func (v *NullableReactionWithCastInfo) Set(val *ReactionWithCastInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionWithCastInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionWithCastInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionWithCastInfo(val *ReactionWithCastInfo) *NullableReactionWithCastInfo {
	return &NullableReactionWithCastInfo{value: val, isSet: true}
}

func (v NullableReactionWithCastInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionWithCastInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
