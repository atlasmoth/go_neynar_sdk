/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CastAddBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastAddBody{}

// CastAddBody Adds a new Cast
type CastAddBody struct {
	EmbedsDeprecated []string `json:"embedsDeprecated"`
	Mentions []int32 `json:"mentions"`
	ParentCastId *CastId `json:"parentCastId,omitempty"`
	ParentUrl *string `json:"parentUrl,omitempty"`
	Text string `json:"text"`
	MentionsPositions []int64 `json:"mentionsPositions"`
	Embeds []Embed `json:"embeds"`
}

type _CastAddBody CastAddBody

// NewCastAddBody instantiates a new CastAddBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastAddBody(embedsDeprecated []string, mentions []int32, text string, mentionsPositions []int64, embeds []Embed) *CastAddBody {
	this := CastAddBody{}
	this.EmbedsDeprecated = embedsDeprecated
	this.Mentions = mentions
	this.Text = text
	this.MentionsPositions = mentionsPositions
	this.Embeds = embeds
	return &this
}

// NewCastAddBodyWithDefaults instantiates a new CastAddBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastAddBodyWithDefaults() *CastAddBody {
	this := CastAddBody{}
	return &this
}

// GetEmbedsDeprecated returns the EmbedsDeprecated field value
func (o *CastAddBody) GetEmbedsDeprecated() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmbedsDeprecated
}

// GetEmbedsDeprecatedOk returns a tuple with the EmbedsDeprecated field value
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetEmbedsDeprecatedOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmbedsDeprecated, true
}

// SetEmbedsDeprecated sets field value
func (o *CastAddBody) SetEmbedsDeprecated(v []string) {
	o.EmbedsDeprecated = v
}

// GetMentions returns the Mentions field value
func (o *CastAddBody) GetMentions() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetMentionsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mentions, true
}

// SetMentions sets field value
func (o *CastAddBody) SetMentions(v []int32) {
	o.Mentions = v
}

// GetParentCastId returns the ParentCastId field value if set, zero value otherwise.
func (o *CastAddBody) GetParentCastId() CastId {
	if o == nil || IsNil(o.ParentCastId) {
		var ret CastId
		return ret
	}
	return *o.ParentCastId
}

// GetParentCastIdOk returns a tuple with the ParentCastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetParentCastIdOk() (*CastId, bool) {
	if o == nil || IsNil(o.ParentCastId) {
		return nil, false
	}
	return o.ParentCastId, true
}

// HasParentCastId returns a boolean if a field has been set.
func (o *CastAddBody) HasParentCastId() bool {
	if o != nil && !IsNil(o.ParentCastId) {
		return true
	}

	return false
}

// SetParentCastId gets a reference to the given CastId and assigns it to the ParentCastId field.
func (o *CastAddBody) SetParentCastId(v CastId) {
	o.ParentCastId = &v
}

// GetParentUrl returns the ParentUrl field value if set, zero value otherwise.
func (o *CastAddBody) GetParentUrl() string {
	if o == nil || IsNil(o.ParentUrl) {
		var ret string
		return ret
	}
	return *o.ParentUrl
}

// GetParentUrlOk returns a tuple with the ParentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetParentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ParentUrl) {
		return nil, false
	}
	return o.ParentUrl, true
}

// HasParentUrl returns a boolean if a field has been set.
func (o *CastAddBody) HasParentUrl() bool {
	if o != nil && !IsNil(o.ParentUrl) {
		return true
	}

	return false
}

// SetParentUrl gets a reference to the given string and assigns it to the ParentUrl field.
func (o *CastAddBody) SetParentUrl(v string) {
	o.ParentUrl = &v
}

// GetText returns the Text field value
func (o *CastAddBody) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *CastAddBody) SetText(v string) {
	o.Text = v
}

// GetMentionsPositions returns the MentionsPositions field value
func (o *CastAddBody) GetMentionsPositions() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.MentionsPositions
}

// GetMentionsPositionsOk returns a tuple with the MentionsPositions field value
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetMentionsPositionsOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionsPositions, true
}

// SetMentionsPositions sets field value
func (o *CastAddBody) SetMentionsPositions(v []int64) {
	o.MentionsPositions = v
}

// GetEmbeds returns the Embeds field value
func (o *CastAddBody) GetEmbeds() []Embed {
	if o == nil {
		var ret []Embed
		return ret
	}

	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value
// and a boolean to check if the value has been set.
func (o *CastAddBody) GetEmbedsOk() ([]Embed, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// SetEmbeds sets field value
func (o *CastAddBody) SetEmbeds(v []Embed) {
	o.Embeds = v
}

func (o CastAddBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastAddBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["embedsDeprecated"] = o.EmbedsDeprecated
	toSerialize["mentions"] = o.Mentions
	if !IsNil(o.ParentCastId) {
		toSerialize["parentCastId"] = o.ParentCastId
	}
	if !IsNil(o.ParentUrl) {
		toSerialize["parentUrl"] = o.ParentUrl
	}
	toSerialize["text"] = o.Text
	toSerialize["mentionsPositions"] = o.MentionsPositions
	toSerialize["embeds"] = o.Embeds
	return toSerialize, nil
}

func (o *CastAddBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"embedsDeprecated",
		"mentions",
		"text",
		"mentionsPositions",
		"embeds",
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

	varCastAddBody := _CastAddBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastAddBody)

	if err != nil {
		return err
	}

	*o = CastAddBody(varCastAddBody)

	return err
}

type NullableCastAddBody struct {
	value *CastAddBody
	isSet bool
}

func (v NullableCastAddBody) Get() *CastAddBody {
	return v.value
}

func (v *NullableCastAddBody) Set(val *CastAddBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCastAddBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCastAddBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastAddBody(val *CastAddBody) *NullableCastAddBody {
	return &NullableCastAddBody{value: val, isSet: true}
}

func (v NullableCastAddBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastAddBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


