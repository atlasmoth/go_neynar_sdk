/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CastLikesResponseResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastLikesResponseResult{}

// CastLikesResponseResult struct for CastLikesResponseResult
type CastLikesResponseResult struct {
	Likes []Reaction `json:"likes"`
	Next NextCursor `json:"next"`
}

type _CastLikesResponseResult CastLikesResponseResult

// NewCastLikesResponseResult instantiates a new CastLikesResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastLikesResponseResult(likes []Reaction, next NextCursor) *CastLikesResponseResult {
	this := CastLikesResponseResult{}
	this.Likes = likes
	this.Next = next
	return &this
}

// NewCastLikesResponseResultWithDefaults instantiates a new CastLikesResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastLikesResponseResultWithDefaults() *CastLikesResponseResult {
	this := CastLikesResponseResult{}
	return &this
}

// GetLikes returns the Likes field value
func (o *CastLikesResponseResult) GetLikes() []Reaction {
	if o == nil {
		var ret []Reaction
		return ret
	}

	return o.Likes
}

// GetLikesOk returns a tuple with the Likes field value
// and a boolean to check if the value has been set.
func (o *CastLikesResponseResult) GetLikesOk() ([]Reaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Likes, true
}

// SetLikes sets field value
func (o *CastLikesResponseResult) SetLikes(v []Reaction) {
	o.Likes = v
}

// GetNext returns the Next field value
func (o *CastLikesResponseResult) GetNext() NextCursor {
	if o == nil {
		var ret NextCursor
		return ret
	}

	return o.Next
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
func (o *CastLikesResponseResult) GetNextOk() (*NextCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Next, true
}

// SetNext sets field value
func (o *CastLikesResponseResult) SetNext(v NextCursor) {
	o.Next = v
}

func (o CastLikesResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastLikesResponseResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["likes"] = o.Likes
	toSerialize["next"] = o.Next
	return toSerialize, nil
}

func (o *CastLikesResponseResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"likes",
		"next",
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

	varCastLikesResponseResult := _CastLikesResponseResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastLikesResponseResult)

	if err != nil {
		return err
	}

	*o = CastLikesResponseResult(varCastLikesResponseResult)

	return err
}

type NullableCastLikesResponseResult struct {
	value *CastLikesResponseResult
	isSet bool
}

func (v NullableCastLikesResponseResult) Get() *CastLikesResponseResult {
	return v.value
}

func (v *NullableCastLikesResponseResult) Set(val *CastLikesResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCastLikesResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCastLikesResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastLikesResponseResult(val *CastLikesResponseResult) *NullableCastLikesResponseResult {
	return &NullableCastLikesResponseResult{value: val, isSet: true}
}

func (v NullableCastLikesResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastLikesResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

