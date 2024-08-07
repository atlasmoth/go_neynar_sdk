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
)

// checks if the ReactionsCastResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionsCastResponse{}

// ReactionsCastResponse struct for ReactionsCastResponse
type ReactionsCastResponse struct {
	Reactions []ReactionForCast `json:"reactions"`
	Next      NextCursor        `json:"next"`
}

type _ReactionsCastResponse ReactionsCastResponse

// NewReactionsCastResponse instantiates a new ReactionsCastResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionsCastResponse(reactions []ReactionForCast, next NextCursor) *ReactionsCastResponse {
	this := ReactionsCastResponse{}
	this.Reactions = reactions
	this.Next = next
	return &this
}

// NewReactionsCastResponseWithDefaults instantiates a new ReactionsCastResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionsCastResponseWithDefaults() *ReactionsCastResponse {
	this := ReactionsCastResponse{}
	return &this
}

// GetReactions returns the Reactions field value
func (o *ReactionsCastResponse) GetReactions() []ReactionForCast {
	if o == nil {
		var ret []ReactionForCast
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *ReactionsCastResponse) GetReactionsOk() ([]ReactionForCast, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reactions, true
}

// SetReactions sets field value
func (o *ReactionsCastResponse) SetReactions(v []ReactionForCast) {
	o.Reactions = v
}

// GetNext returns the Next field value
func (o *ReactionsCastResponse) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *ReactionsCastResponse) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *ReactionsCastResponse) SetNext(v NextCursor) {
	o.Next = v
}

func (o ReactionsCastResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionsCastResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reactions"] = o.Reactions
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *ReactionsCastResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reactions",
		"next",
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

	varReactionsCastResponse := _ReactionsCastResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReactionsCastResponse)

	if err != nil {
		return err
	}

	*o = ReactionsCastResponse(varReactionsCastResponse)

	return err
}

type NullableReactionsCastResponse struct {
	value *ReactionsCastResponse
	isSet bool
}

func (v NullableReactionsCastResponse) Get() *ReactionsCastResponse {
	return v.value
}

func (v *NullableReactionsCastResponse) Set(val *ReactionsCastResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionsCastResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionsCastResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionsCastResponse(val *ReactionsCastResponse) *NullableReactionsCastResponse {
	return &NullableReactionsCastResponse{value: val, isSet: true}
}

func (v NullableReactionsCastResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionsCastResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
