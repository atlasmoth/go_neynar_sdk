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

// checks if the EmbedUrlMetadataVideoStreamInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbedUrlMetadataVideoStreamInner{}

// EmbedUrlMetadataVideoStreamInner struct for EmbedUrlMetadataVideoStreamInner
type EmbedUrlMetadataVideoStreamInner struct {
	CodecName *string `json:"codec_name,omitempty"`
	HeightPx *int32 `json:"height_px,omitempty"`
	WidthPx *int32 `json:"width_px,omitempty"`
}

// NewEmbedUrlMetadataVideoStreamInner instantiates a new EmbedUrlMetadataVideoStreamInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbedUrlMetadataVideoStreamInner() *EmbedUrlMetadataVideoStreamInner {
	this := EmbedUrlMetadataVideoStreamInner{}
	return &this
}

// NewEmbedUrlMetadataVideoStreamInnerWithDefaults instantiates a new EmbedUrlMetadataVideoStreamInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbedUrlMetadataVideoStreamInnerWithDefaults() *EmbedUrlMetadataVideoStreamInner {
	this := EmbedUrlMetadataVideoStreamInner{}
	return &this
}

// GetCodecName returns the CodecName field value if set, zero value otherwise.
func (o *EmbedUrlMetadataVideoStreamInner) GetCodecName() string {
	if o == nil || IsNil(o.CodecName) {
		var ret string
		return ret
	}
	return *o.CodecName
}

// GetCodecNameOk returns a tuple with the CodecName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbedUrlMetadataVideoStreamInner) GetCodecNameOk() (*string, bool) {
	if o == nil || IsNil(o.CodecName) {
		return nil, false
	}
	return o.CodecName, true
}

// HasCodecName returns a boolean if a field has been set.
func (o *EmbedUrlMetadataVideoStreamInner) HasCodecName() bool {
	if o != nil && !IsNil(o.CodecName) {
		return true
	}

	return false
}

// SetCodecName gets a reference to the given string and assigns it to the CodecName field.
func (o *EmbedUrlMetadataVideoStreamInner) SetCodecName(v string) {
	o.CodecName = &v
}

// GetHeightPx returns the HeightPx field value if set, zero value otherwise.
func (o *EmbedUrlMetadataVideoStreamInner) GetHeightPx() int32 {
	if o == nil || IsNil(o.HeightPx) {
		var ret int32
		return ret
	}
	return *o.HeightPx
}

// GetHeightPxOk returns a tuple with the HeightPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbedUrlMetadataVideoStreamInner) GetHeightPxOk() (*int32, bool) {
	if o == nil || IsNil(o.HeightPx) {
		return nil, false
	}
	return o.HeightPx, true
}

// HasHeightPx returns a boolean if a field has been set.
func (o *EmbedUrlMetadataVideoStreamInner) HasHeightPx() bool {
	if o != nil && !IsNil(o.HeightPx) {
		return true
	}

	return false
}

// SetHeightPx gets a reference to the given int32 and assigns it to the HeightPx field.
func (o *EmbedUrlMetadataVideoStreamInner) SetHeightPx(v int32) {
	o.HeightPx = &v
}

// GetWidthPx returns the WidthPx field value if set, zero value otherwise.
func (o *EmbedUrlMetadataVideoStreamInner) GetWidthPx() int32 {
	if o == nil || IsNil(o.WidthPx) {
		var ret int32
		return ret
	}
	return *o.WidthPx
}

// GetWidthPxOk returns a tuple with the WidthPx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbedUrlMetadataVideoStreamInner) GetWidthPxOk() (*int32, bool) {
	if o == nil || IsNil(o.WidthPx) {
		return nil, false
	}
	return o.WidthPx, true
}

// HasWidthPx returns a boolean if a field has been set.
func (o *EmbedUrlMetadataVideoStreamInner) HasWidthPx() bool {
	if o != nil && !IsNil(o.WidthPx) {
		return true
	}

	return false
}

// SetWidthPx gets a reference to the given int32 and assigns it to the WidthPx field.
func (o *EmbedUrlMetadataVideoStreamInner) SetWidthPx(v int32) {
	o.WidthPx = &v
}

func (o EmbedUrlMetadataVideoStreamInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbedUrlMetadataVideoStreamInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CodecName) {
		toSerialize["codec_name"] = o.CodecName
	}
	if !IsNil(o.HeightPx) {
		toSerialize["height_px"] = o.HeightPx
	}
	if !IsNil(o.WidthPx) {
		toSerialize["width_px"] = o.WidthPx
	}
	return toSerialize, nil
}

type NullableEmbedUrlMetadataVideoStreamInner struct {
	value *EmbedUrlMetadataVideoStreamInner
	isSet bool
}

func (v NullableEmbedUrlMetadataVideoStreamInner) Get() *EmbedUrlMetadataVideoStreamInner {
	return v.value
}

func (v *NullableEmbedUrlMetadataVideoStreamInner) Set(val *EmbedUrlMetadataVideoStreamInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbedUrlMetadataVideoStreamInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbedUrlMetadataVideoStreamInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbedUrlMetadataVideoStreamInner(val *EmbedUrlMetadataVideoStreamInner) *NullableEmbedUrlMetadataVideoStreamInner {
	return &NullableEmbedUrlMetadataVideoStreamInner{value: val, isSet: true}
}

func (v NullableEmbedUrlMetadataVideoStreamInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbedUrlMetadataVideoStreamInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

