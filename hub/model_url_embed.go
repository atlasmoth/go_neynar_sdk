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

// checks if the UrlEmbed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UrlEmbed{}

// UrlEmbed struct for UrlEmbed
type UrlEmbed struct {
	Url string `json:"url"`
	AdditionalProperties map[string]interface{}
}

type _UrlEmbed UrlEmbed

// NewUrlEmbed instantiates a new UrlEmbed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUrlEmbed(url string) *UrlEmbed {
	this := UrlEmbed{}
	this.Url = url
	return &this
}

// NewUrlEmbedWithDefaults instantiates a new UrlEmbed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlEmbedWithDefaults() *UrlEmbed {
	this := UrlEmbed{}
	return &this
}

// GetUrl returns the Url field value
func (o *UrlEmbed) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UrlEmbed) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UrlEmbed) SetUrl(v string) {
	o.Url = v
}

func (o UrlEmbed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlEmbed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UrlEmbed) UnmarshalJSON(data []byte) (err error) {
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

	varUrlEmbed := _UrlEmbed{}

	err = json.Unmarshal(data, &varUrlEmbed)

	if err != nil {
		return err
	}

	*o = UrlEmbed(varUrlEmbed)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUrlEmbed struct {
	value *UrlEmbed
	isSet bool
}

func (v NullableUrlEmbed) Get() *UrlEmbed {
	return v.value
}

func (v *NullableUrlEmbed) Set(val *UrlEmbed) {
	v.value = val
	v.isSet = true
}

func (v NullableUrlEmbed) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlEmbed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlEmbed(val *UrlEmbed) *NullableUrlEmbed {
	return &NullableUrlEmbed{value: val, isSet: true}
}

func (v NullableUrlEmbed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlEmbed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


