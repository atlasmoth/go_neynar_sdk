/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the Cast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cast{}

// Cast struct for Cast
type Cast struct {
	Hash string `json:"hash"`
	ParentHash NullableString `json:"parent_hash"`
	ParentUrl NullableString `json:"parent_url"`
	RootParentUrl NullableString `json:"root_parent_url"`
	ParentAuthor CastParentAuthor `json:"parent_author"`
	Author User `json:"author"`
	Text string `json:"text"`
	Timestamp time.Time `json:"timestamp"`
	Embeds []EmbeddedCast `json:"embeds"`
	Type *CastNotificationType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Cast Cast

// NewCast instantiates a new Cast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCast(hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []EmbeddedCast) *Cast {
	this := Cast{}
	this.Hash = hash
	this.ParentHash = parentHash
	this.ParentUrl = parentUrl
	this.RootParentUrl = rootParentUrl
	this.ParentAuthor = parentAuthor
	this.Author = author
	this.Text = text
	this.Timestamp = timestamp
	this.Embeds = embeds
	return &this
}

// NewCastWithDefaults instantiates a new Cast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithDefaults() *Cast {
	this := Cast{}
	return &this
}

// GetHash returns the Hash field value
func (o *Cast) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *Cast) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *Cast) SetHash(v string) {
	o.Hash = v
}

// GetParentHash returns the ParentHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Cast) GetParentHash() string {
	if o == nil || o.ParentHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentHash.Get()
}

// GetParentHashOk returns a tuple with the ParentHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cast) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentHash.Get(), o.ParentHash.IsSet()
}

// SetParentHash sets field value
func (o *Cast) SetParentHash(v string) {
	o.ParentHash.Set(&v)
}

// GetParentUrl returns the ParentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Cast) GetParentUrl() string {
	if o == nil || o.ParentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentUrl.Get()
}

// GetParentUrlOk returns a tuple with the ParentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cast) GetParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentUrl.Get(), o.ParentUrl.IsSet()
}

// SetParentUrl sets field value
func (o *Cast) SetParentUrl(v string) {
	o.ParentUrl.Set(&v)
}

// GetRootParentUrl returns the RootParentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Cast) GetRootParentUrl() string {
	if o == nil || o.RootParentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.RootParentUrl.Get()
}

// GetRootParentUrlOk returns a tuple with the RootParentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cast) GetRootParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootParentUrl.Get(), o.RootParentUrl.IsSet()
}

// SetRootParentUrl sets field value
func (o *Cast) SetRootParentUrl(v string) {
	o.RootParentUrl.Set(&v)
}

// GetParentAuthor returns the ParentAuthor field value
func (o *Cast) GetParentAuthor() CastParentAuthor {
	if o == nil {
		var ret CastParentAuthor
		return ret
	}

	return o.ParentAuthor
}

// GetParentAuthorOk returns a tuple with the ParentAuthor field value
// and a boolean to check if the value has been set.
func (o *Cast) GetParentAuthorOk() (*CastParentAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentAuthor, true
}

// SetParentAuthor sets field value
func (o *Cast) SetParentAuthor(v CastParentAuthor) {
	o.ParentAuthor = v
}

// GetAuthor returns the Author field value
func (o *Cast) GetAuthor() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *Cast) GetAuthorOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *Cast) SetAuthor(v User) {
	o.Author = v
}

// GetText returns the Text field value
func (o *Cast) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *Cast) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *Cast) SetText(v string) {
	o.Text = v
}

// GetTimestamp returns the Timestamp field value
func (o *Cast) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Cast) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Cast) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEmbeds returns the Embeds field value
func (o *Cast) GetEmbeds() []EmbeddedCast {
	if o == nil {
		var ret []EmbeddedCast
		return ret
	}

	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value
// and a boolean to check if the value has been set.
func (o *Cast) GetEmbedsOk() ([]EmbeddedCast, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// SetEmbeds sets field value
func (o *Cast) SetEmbeds(v []EmbeddedCast) {
	o.Embeds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Cast) GetType() CastNotificationType {
	if o == nil || IsNil(o.Type) {
		var ret CastNotificationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetTypeOk() (*CastNotificationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Cast) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CastNotificationType and assigns it to the Type field.
func (o *Cast) SetType(v CastNotificationType) {
	o.Type = &v
}

func (o Cast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Cast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["parent_hash"] = o.ParentHash.Get()
	toSerialize["parent_url"] = o.ParentUrl.Get()
	toSerialize["root_parent_url"] = o.RootParentUrl.Get()
	toSerialize["parent_author"] = o.ParentAuthor
	toSerialize["author"] = o.Author
	toSerialize["text"] = o.Text
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["embeds"] = o.Embeds
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Cast) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"parent_hash",
		"parent_url",
		"root_parent_url",
		"parent_author",
		"author",
		"text",
		"timestamp",
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

	varCast := _Cast{}

	err = json.Unmarshal(data, &varCast)

	if err != nil {
		return err
	}

	*o = Cast(varCast)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "hash")
		delete(additionalProperties, "parent_hash")
		delete(additionalProperties, "parent_url")
		delete(additionalProperties, "root_parent_url")
		delete(additionalProperties, "parent_author")
		delete(additionalProperties, "author")
		delete(additionalProperties, "text")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "embeds")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCast struct {
	value *Cast
	isSet bool
}

func (v NullableCast) Get() *Cast {
	return v.value
}

func (v *NullableCast) Set(val *Cast) {
	v.value = val
	v.isSet = true
}

func (v NullableCast) IsSet() bool {
	return v.isSet
}

func (v *NullableCast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCast(val *Cast) *NullableCast {
	return &NullableCast{value: val, isSet: true}
}

func (v NullableCast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


