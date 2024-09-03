/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the ReactionBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReactionBody{}

// ReactionBody Adds or removes a Reaction from a Cast
type ReactionBody struct {
	Type ReactionType `json:"type"`
	TargetCastId *CastId `json:"targetCastId,omitempty"`
	TargetUrl *string `json:"targetUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReactionBody ReactionBody

// NewReactionBody instantiates a new ReactionBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReactionBody(type_ ReactionType) *ReactionBody {
	this := ReactionBody{}
	this.Type = type_
	return &this
}

// NewReactionBodyWithDefaults instantiates a new ReactionBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionBodyWithDefaults() *ReactionBody {
	this := ReactionBody{}
	var type_ ReactionType = REACTIONTYPE_LIKE
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ReactionBody) GetType() ReactionType {
	if o == nil {
		var ret ReactionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReactionBody) GetTypeOk() (*ReactionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReactionBody) SetType(v ReactionType) {
	o.Type = v
}

// GetTargetCastId returns the TargetCastId field value if set, zero value otherwise.
func (o *ReactionBody) GetTargetCastId() CastId {
	if o == nil || IsNil(o.TargetCastId) {
		var ret CastId
		return ret
	}
	return *o.TargetCastId
}

// GetTargetCastIdOk returns a tuple with the TargetCastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactionBody) GetTargetCastIdOk() (*CastId, bool) {
	if o == nil || IsNil(o.TargetCastId) {
		return nil, false
	}
	return o.TargetCastId, true
}

// HasTargetCastId returns a boolean if a field has been set.
func (o *ReactionBody) HasTargetCastId() bool {
	if o != nil && !IsNil(o.TargetCastId) {
		return true
	}

	return false
}

// SetTargetCastId gets a reference to the given CastId and assigns it to the TargetCastId field.
func (o *ReactionBody) SetTargetCastId(v CastId) {
	o.TargetCastId = &v
}

// GetTargetUrl returns the TargetUrl field value if set, zero value otherwise.
func (o *ReactionBody) GetTargetUrl() string {
	if o == nil || IsNil(o.TargetUrl) {
		var ret string
		return ret
	}
	return *o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactionBody) GetTargetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUrl) {
		return nil, false
	}
	return o.TargetUrl, true
}

// HasTargetUrl returns a boolean if a field has been set.
func (o *ReactionBody) HasTargetUrl() bool {
	if o != nil && !IsNil(o.TargetUrl) {
		return true
	}

	return false
}

// SetTargetUrl gets a reference to the given string and assigns it to the TargetUrl field.
func (o *ReactionBody) SetTargetUrl(v string) {
	o.TargetUrl = &v
}

func (o ReactionBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReactionBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.TargetCastId) {
		toSerialize["targetCastId"] = o.TargetCastId
	}
	if !IsNil(o.TargetUrl) {
		toSerialize["targetUrl"] = o.TargetUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReactionBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varReactionBody := _ReactionBody{}

	err = json.Unmarshal(data, &varReactionBody)

	if err != nil {
		return err
	}

	*o = ReactionBody(varReactionBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "targetCastId")
		delete(additionalProperties, "targetUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReactionBody struct {
	value *ReactionBody
	isSet bool
}

func (v NullableReactionBody) Get() *ReactionBody {
	return v.value
}

func (v *NullableReactionBody) Set(val *ReactionBody) {
	v.value = val
	v.isSet = true
}

func (v NullableReactionBody) IsSet() bool {
	return v.isSet
}

func (v *NullableReactionBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReactionBody(val *ReactionBody) *NullableReactionBody {
	return &NullableReactionBody{value: val, isSet: true}
}

func (v NullableReactionBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReactionBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


