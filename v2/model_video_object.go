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

// checks if the VideoObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VideoObject{}

// VideoObject struct for VideoObject
type VideoObject struct {
	Height *int32 `json:"height,omitempty"`
	Type *string `json:"type,omitempty"`
	Url string `json:"url"`
	Width *int32 `json:"width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VideoObject VideoObject

// NewVideoObject instantiates a new VideoObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideoObject(url string) *VideoObject {
	this := VideoObject{}
	this.Url = url
	return &this
}

// NewVideoObjectWithDefaults instantiates a new VideoObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoObjectWithDefaults() *VideoObject {
	this := VideoObject{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *VideoObject) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoObject) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *VideoObject) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *VideoObject) SetHeight(v int32) {
	o.Height = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VideoObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VideoObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VideoObject) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value
func (o *VideoObject) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VideoObject) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VideoObject) SetUrl(v string) {
	o.Url = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *VideoObject) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VideoObject) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *VideoObject) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *VideoObject) SetWidth(v int32) {
	o.Width = &v
}

func (o VideoObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VideoObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VideoObject) UnmarshalJSON(data []byte) (err error) {
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

	varVideoObject := _VideoObject{}

	err = json.Unmarshal(data, &varVideoObject)

	if err != nil {
		return err
	}

	*o = VideoObject(varVideoObject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "height")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		delete(additionalProperties, "width")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVideoObject struct {
	value *VideoObject
	isSet bool
}

func (v NullableVideoObject) Get() *VideoObject {
	return v.value
}

func (v *NullableVideoObject) Set(val *VideoObject) {
	v.value = val
	v.isSet = true
}

func (v NullableVideoObject) IsSet() bool {
	return v.isSet
}

func (v *NullableVideoObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideoObject(val *VideoObject) *NullableVideoObject {
	return &NullableVideoObject{value: val, isSet: true}
}

func (v NullableVideoObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideoObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


