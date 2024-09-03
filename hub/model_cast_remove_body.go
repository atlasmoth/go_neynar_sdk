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

// checks if the CastRemoveBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastRemoveBody{}

// CastRemoveBody Removes an existing Cast
type CastRemoveBody struct {
	TargetHash string `json:"targetHash" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	AdditionalProperties map[string]interface{}
}

type _CastRemoveBody CastRemoveBody

// NewCastRemoveBody instantiates a new CastRemoveBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastRemoveBody(targetHash string) *CastRemoveBody {
	this := CastRemoveBody{}
	this.TargetHash = targetHash
	return &this
}

// NewCastRemoveBodyWithDefaults instantiates a new CastRemoveBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastRemoveBodyWithDefaults() *CastRemoveBody {
	this := CastRemoveBody{}
	return &this
}

// GetTargetHash returns the TargetHash field value
func (o *CastRemoveBody) GetTargetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetHash
}

// GetTargetHashOk returns a tuple with the TargetHash field value
// and a boolean to check if the value has been set.
func (o *CastRemoveBody) GetTargetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetHash, true
}

// SetTargetHash sets field value
func (o *CastRemoveBody) SetTargetHash(v string) {
	o.TargetHash = v
}

func (o CastRemoveBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastRemoveBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetHash"] = o.TargetHash

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CastRemoveBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetHash",
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

	varCastRemoveBody := _CastRemoveBody{}

	err = json.Unmarshal(data, &varCastRemoveBody)

	if err != nil {
		return err
	}

	*o = CastRemoveBody(varCastRemoveBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "targetHash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCastRemoveBody struct {
	value *CastRemoveBody
	isSet bool
}

func (v NullableCastRemoveBody) Get() *CastRemoveBody {
	return v.value
}

func (v *NullableCastRemoveBody) Set(val *CastRemoveBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCastRemoveBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCastRemoveBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastRemoveBody(val *CastRemoveBody) *NullableCastRemoveBody {
	return &NullableCastRemoveBody{value: val, isSet: true}
}

func (v NullableCastRemoveBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastRemoveBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


