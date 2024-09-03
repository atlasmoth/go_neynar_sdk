/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the RecasterPfp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecasterPfp{}

// RecasterPfp struct for RecasterPfp
type RecasterPfp struct {
	// The URL of the recaster's profile picture.
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _RecasterPfp RecasterPfp

// NewRecasterPfp instantiates a new RecasterPfp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecasterPfp(url string) *RecasterPfp {
	this := RecasterPfp{}
	this.Url = url
	return &this
}

// NewRecasterPfpWithDefaults instantiates a new RecasterPfp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecasterPfpWithDefaults() *RecasterPfp {
	this := RecasterPfp{}
	return &this
}

// GetUrl returns the Url field value
func (o *RecasterPfp) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *RecasterPfp) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *RecasterPfp) SetUrl(v string) {
	o.Url = v
}

func (o RecasterPfp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecasterPfp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecasterPfp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
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

	varRecasterPfp := _RecasterPfp{}

	err = json.Unmarshal(data, &varRecasterPfp)

	if err != nil {
		return err
	}

	*o = RecasterPfp(varRecasterPfp)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecasterPfp struct {
	value *RecasterPfp
	isSet bool
}

func (v NullableRecasterPfp) Get() *RecasterPfp {
	return v.value
}

func (v *NullableRecasterPfp) Set(val *RecasterPfp) {
	v.value = val
	v.isSet = true
}

func (v NullableRecasterPfp) IsSet() bool {
	return v.isSet
}

func (v *NullableRecasterPfp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecasterPfp(val *RecasterPfp) *NullableRecasterPfp {
	return &NullableRecasterPfp{value: val, isSet: true}
}

func (v NullableRecasterPfp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecasterPfp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


