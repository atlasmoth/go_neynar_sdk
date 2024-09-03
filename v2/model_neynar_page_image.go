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

// checks if the NeynarPageImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarPageImage{}

// NeynarPageImage struct for NeynarPageImage
type NeynarPageImage struct {
	// The URL of the page's image.
	Url string `json:"url"`
	// The aspect ratio of the image.
	AspectRatio string `json:"aspect_ratio"`
	AdditionalProperties map[string]interface{}
}

type _NeynarPageImage NeynarPageImage

// NewNeynarPageImage instantiates a new NeynarPageImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarPageImage(url string, aspectRatio string) *NeynarPageImage {
	this := NeynarPageImage{}
	this.Url = url
	this.AspectRatio = aspectRatio
	return &this
}

// NewNeynarPageImageWithDefaults instantiates a new NeynarPageImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarPageImageWithDefaults() *NeynarPageImage {
	this := NeynarPageImage{}
	return &this
}

// GetUrl returns the Url field value
func (o *NeynarPageImage) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *NeynarPageImage) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *NeynarPageImage) SetUrl(v string) {
	o.Url = v
}

// GetAspectRatio returns the AspectRatio field value
func (o *NeynarPageImage) GetAspectRatio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AspectRatio
}

// GetAspectRatioOk returns a tuple with the AspectRatio field value
// and a boolean to check if the value has been set.
func (o *NeynarPageImage) GetAspectRatioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AspectRatio, true
}

// SetAspectRatio sets field value
func (o *NeynarPageImage) SetAspectRatio(v string) {
	o.AspectRatio = v
}

func (o NeynarPageImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarPageImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["aspect_ratio"] = o.AspectRatio

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NeynarPageImage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"url",
		"aspect_ratio",
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

	varNeynarPageImage := _NeynarPageImage{}

	err = json.Unmarshal(data, &varNeynarPageImage)

	if err != nil {
		return err
	}

	*o = NeynarPageImage(varNeynarPageImage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "aspect_ratio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNeynarPageImage struct {
	value *NeynarPageImage
	isSet bool
}

func (v NullableNeynarPageImage) Get() *NeynarPageImage {
	return v.value
}

func (v *NullableNeynarPageImage) Set(val *NeynarPageImage) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarPageImage) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarPageImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarPageImage(val *NeynarPageImage) *NullableNeynarPageImage {
	return &NullableNeynarPageImage{value: val, isSet: true}
}

func (v NullableNeynarPageImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarPageImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


