/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DehydratedChannel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DehydratedChannel{}

// DehydratedChannel struct for DehydratedChannel
type DehydratedChannel struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Object *string `json:"object,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
}

// NewDehydratedChannel instantiates a new DehydratedChannel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDehydratedChannel() *DehydratedChannel {
	this := DehydratedChannel{}
	return &this
}

// NewDehydratedChannelWithDefaults instantiates a new DehydratedChannel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDehydratedChannelWithDefaults() *DehydratedChannel {
	this := DehydratedChannel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DehydratedChannel) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DehydratedChannel) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DehydratedChannel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DehydratedChannel) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DehydratedChannel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DehydratedChannel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DehydratedChannel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DehydratedChannel) SetName(v string) {
	o.Name = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *DehydratedChannel) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DehydratedChannel) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *DehydratedChannel) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *DehydratedChannel) SetObject(v string) {
	o.Object = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *DehydratedChannel) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DehydratedChannel) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *DehydratedChannel) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *DehydratedChannel) SetImageUrl(v string) {
	o.ImageUrl = &v
}

func (o DehydratedChannel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DehydratedChannel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	return toSerialize, nil
}

type NullableDehydratedChannel struct {
	value *DehydratedChannel
	isSet bool
}

func (v NullableDehydratedChannel) Get() *DehydratedChannel {
	return v.value
}

func (v *NullableDehydratedChannel) Set(val *DehydratedChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableDehydratedChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableDehydratedChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDehydratedChannel(val *DehydratedChannel) *NullableDehydratedChannel {
	return &NullableDehydratedChannel{value: val, isSet: true}
}

func (v NullableDehydratedChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDehydratedChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


