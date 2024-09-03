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

// checks if the SubscriptionMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionMetadata{}

// SubscriptionMetadata struct for SubscriptionMetadata
type SubscriptionMetadata struct {
	Title string `json:"title"`
	Symbol string `json:"symbol"`
	ArtUrl string `json:"art_url"`
	AdditionalProperties map[string]interface{}
}

type _SubscriptionMetadata SubscriptionMetadata

// NewSubscriptionMetadata instantiates a new SubscriptionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionMetadata(title string, symbol string, artUrl string) *SubscriptionMetadata {
	this := SubscriptionMetadata{}
	this.Title = title
	this.Symbol = symbol
	this.ArtUrl = artUrl
	return &this
}

// NewSubscriptionMetadataWithDefaults instantiates a new SubscriptionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionMetadataWithDefaults() *SubscriptionMetadata {
	this := SubscriptionMetadata{}
	return &this
}

// GetTitle returns the Title field value
func (o *SubscriptionMetadata) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SubscriptionMetadata) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SubscriptionMetadata) SetTitle(v string) {
	o.Title = v
}

// GetSymbol returns the Symbol field value
func (o *SubscriptionMetadata) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SubscriptionMetadata) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SubscriptionMetadata) SetSymbol(v string) {
	o.Symbol = v
}

// GetArtUrl returns the ArtUrl field value
func (o *SubscriptionMetadata) GetArtUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ArtUrl
}

// GetArtUrlOk returns a tuple with the ArtUrl field value
// and a boolean to check if the value has been set.
func (o *SubscriptionMetadata) GetArtUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtUrl, true
}

// SetArtUrl sets field value
func (o *SubscriptionMetadata) SetArtUrl(v string) {
	o.ArtUrl = v
}

func (o SubscriptionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["symbol"] = o.Symbol
	toSerialize["art_url"] = o.ArtUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SubscriptionMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"symbol",
		"art_url",
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

	varSubscriptionMetadata := _SubscriptionMetadata{}

	err = json.Unmarshal(data, &varSubscriptionMetadata)

	if err != nil {
		return err
	}

	*o = SubscriptionMetadata(varSubscriptionMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "art_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSubscriptionMetadata struct {
	value *SubscriptionMetadata
	isSet bool
}

func (v NullableSubscriptionMetadata) Get() *SubscriptionMetadata {
	return v.value
}

func (v *NullableSubscriptionMetadata) Set(val *SubscriptionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionMetadata(val *SubscriptionMetadata) *NullableSubscriptionMetadata {
	return &NullableSubscriptionMetadata{value: val, isSet: true}
}

func (v NullableSubscriptionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


