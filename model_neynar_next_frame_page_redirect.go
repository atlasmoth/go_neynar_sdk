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

// checks if the NeynarNextFramePageRedirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarNextFramePageRedirect{}

// NeynarNextFramePageRedirect struct for NeynarNextFramePageRedirect
type NeynarNextFramePageRedirect struct {
	// The URL to redirect to.
	RedirectUrl string `json:"redirect_url"`
}

type _NeynarNextFramePageRedirect NeynarNextFramePageRedirect

// NewNeynarNextFramePageRedirect instantiates a new NeynarNextFramePageRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarNextFramePageRedirect(redirectUrl string) *NeynarNextFramePageRedirect {
	this := NeynarNextFramePageRedirect{}
	this.RedirectUrl = redirectUrl
	return &this
}

// NewNeynarNextFramePageRedirectWithDefaults instantiates a new NeynarNextFramePageRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarNextFramePageRedirectWithDefaults() *NeynarNextFramePageRedirect {
	this := NeynarNextFramePageRedirect{}
	return &this
}

// GetRedirectUrl returns the RedirectUrl field value
func (o *NeynarNextFramePageRedirect) GetRedirectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value
// and a boolean to check if the value has been set.
func (o *NeynarNextFramePageRedirect) GetRedirectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RedirectUrl, true
}

// SetRedirectUrl sets field value
func (o *NeynarNextFramePageRedirect) SetRedirectUrl(v string) {
	o.RedirectUrl = v
}

func (o NeynarNextFramePageRedirect) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarNextFramePageRedirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["redirect_url"] = o.RedirectUrl
	return toSerialize, nil
}

func (o *NeynarNextFramePageRedirect) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"redirect_url",
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

	varNeynarNextFramePageRedirect := _NeynarNextFramePageRedirect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNeynarNextFramePageRedirect)

	if err != nil {
		return err
	}

	*o = NeynarNextFramePageRedirect(varNeynarNextFramePageRedirect)

	return err
}

type NullableNeynarNextFramePageRedirect struct {
	value *NeynarNextFramePageRedirect
	isSet bool
}

func (v NullableNeynarNextFramePageRedirect) Get() *NeynarNextFramePageRedirect {
	return v.value
}

func (v *NullableNeynarNextFramePageRedirect) Set(val *NeynarNextFramePageRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarNextFramePageRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarNextFramePageRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarNextFramePageRedirect(val *NeynarNextFramePageRedirect) *NullableNeynarNextFramePageRedirect {
	return &NullableNeynarNextFramePageRedirect{value: val, isSet: true}
}

func (v NullableNeynarNextFramePageRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarNextFramePageRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
