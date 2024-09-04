/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the Cast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Cast{}

// Cast struct for Cast
type Cast struct {
	Hash *string `json:"hash,omitempty"`
	ParentHash NullableString `json:"parentHash,omitempty"`
	ParentUrl NullableString `json:"parentUrl,omitempty"`
	ThreadHash *string `json:"threadHash,omitempty"`
	ParentAuthor *CastParentAuthor `json:"parentAuthor,omitempty"`
	MentionedProfiles []User `json:"mentionedProfiles,omitempty"`
	Author *CastAuthor `json:"author,omitempty"`
	Text *string `json:"text,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
	Embeds []EmbedUrl `json:"embeds,omitempty"`
	Type *CastType `json:"type,omitempty"`
}

// NewCast instantiates a new Cast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCast() *Cast {
	this := Cast{}
	return &this
}

// NewCastWithDefaults instantiates a new Cast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithDefaults() *Cast {
	this := Cast{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *Cast) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *Cast) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *Cast) SetHash(v string) {
	o.Hash = &v
}

// GetParentHash returns the ParentHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cast) GetParentHash() string {
	if o == nil || IsNil(o.ParentHash.Get()) {
		var ret string
		return ret
	}
	return *o.ParentHash.Get()
}

// GetParentHashOk returns a tuple with the ParentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cast) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentHash.Get(), o.ParentHash.IsSet()
}

// HasParentHash returns a boolean if a field has been set.
func (o *Cast) HasParentHash() bool {
	if o != nil && o.ParentHash.IsSet() {
		return true
	}

	return false
}

// SetParentHash gets a reference to the given NullableString and assigns it to the ParentHash field.
func (o *Cast) SetParentHash(v string) {
	o.ParentHash.Set(&v)
}
// SetParentHashNil sets the value for ParentHash to be an explicit nil
func (o *Cast) SetParentHashNil() {
	o.ParentHash.Set(nil)
}

// UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
func (o *Cast) UnsetParentHash() {
	o.ParentHash.Unset()
}

// GetParentUrl returns the ParentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Cast) GetParentUrl() string {
	if o == nil || IsNil(o.ParentUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ParentUrl.Get()
}

// GetParentUrlOk returns a tuple with the ParentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cast) GetParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentUrl.Get(), o.ParentUrl.IsSet()
}

// HasParentUrl returns a boolean if a field has been set.
func (o *Cast) HasParentUrl() bool {
	if o != nil && o.ParentUrl.IsSet() {
		return true
	}

	return false
}

// SetParentUrl gets a reference to the given NullableString and assigns it to the ParentUrl field.
func (o *Cast) SetParentUrl(v string) {
	o.ParentUrl.Set(&v)
}
// SetParentUrlNil sets the value for ParentUrl to be an explicit nil
func (o *Cast) SetParentUrlNil() {
	o.ParentUrl.Set(nil)
}

// UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
func (o *Cast) UnsetParentUrl() {
	o.ParentUrl.Unset()
}

// GetThreadHash returns the ThreadHash field value if set, zero value otherwise.
func (o *Cast) GetThreadHash() string {
	if o == nil || IsNil(o.ThreadHash) {
		var ret string
		return ret
	}
	return *o.ThreadHash
}

// GetThreadHashOk returns a tuple with the ThreadHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetThreadHashOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadHash) {
		return nil, false
	}
	return o.ThreadHash, true
}

// HasThreadHash returns a boolean if a field has been set.
func (o *Cast) HasThreadHash() bool {
	if o != nil && !IsNil(o.ThreadHash) {
		return true
	}

	return false
}

// SetThreadHash gets a reference to the given string and assigns it to the ThreadHash field.
func (o *Cast) SetThreadHash(v string) {
	o.ThreadHash = &v
}

// GetParentAuthor returns the ParentAuthor field value if set, zero value otherwise.
func (o *Cast) GetParentAuthor() CastParentAuthor {
	if o == nil || IsNil(o.ParentAuthor) {
		var ret CastParentAuthor
		return ret
	}
	return *o.ParentAuthor
}

// GetParentAuthorOk returns a tuple with the ParentAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetParentAuthorOk() (*CastParentAuthor, bool) {
	if o == nil || IsNil(o.ParentAuthor) {
		return nil, false
	}
	return o.ParentAuthor, true
}

// HasParentAuthor returns a boolean if a field has been set.
func (o *Cast) HasParentAuthor() bool {
	if o != nil && !IsNil(o.ParentAuthor) {
		return true
	}

	return false
}

// SetParentAuthor gets a reference to the given CastParentAuthor and assigns it to the ParentAuthor field.
func (o *Cast) SetParentAuthor(v CastParentAuthor) {
	o.ParentAuthor = &v
}

// GetMentionedProfiles returns the MentionedProfiles field value if set, zero value otherwise.
func (o *Cast) GetMentionedProfiles() []User {
	if o == nil || IsNil(o.MentionedProfiles) {
		var ret []User
		return ret
	}
	return o.MentionedProfiles
}

// GetMentionedProfilesOk returns a tuple with the MentionedProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetMentionedProfilesOk() ([]User, bool) {
	if o == nil || IsNil(o.MentionedProfiles) {
		return nil, false
	}
	return o.MentionedProfiles, true
}

// HasMentionedProfiles returns a boolean if a field has been set.
func (o *Cast) HasMentionedProfiles() bool {
	if o != nil && !IsNil(o.MentionedProfiles) {
		return true
	}

	return false
}

// SetMentionedProfiles gets a reference to the given []User and assigns it to the MentionedProfiles field.
func (o *Cast) SetMentionedProfiles(v []User) {
	o.MentionedProfiles = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *Cast) GetAuthor() CastAuthor {
	if o == nil || IsNil(o.Author) {
		var ret CastAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetAuthorOk() (*CastAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *Cast) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CastAuthor and assigns it to the Author field.
func (o *Cast) SetAuthor(v CastAuthor) {
	o.Author = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Cast) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Cast) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Cast) SetText(v string) {
	o.Text = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *Cast) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *Cast) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *Cast) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise.
func (o *Cast) GetEmbeds() []EmbedUrl {
	if o == nil || IsNil(o.Embeds) {
		var ret []EmbedUrl
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetEmbedsOk() ([]EmbedUrl, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *Cast) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []EmbedUrl and assigns it to the Embeds field.
func (o *Cast) SetEmbeds(v []EmbedUrl) {
	o.Embeds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Cast) GetType() CastType {
	if o == nil || IsNil(o.Type) {
		var ret CastType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Cast) GetTypeOk() (*CastType, bool) {
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

// SetType gets a reference to the given CastType and assigns it to the Type field.
func (o *Cast) SetType(v CastType) {
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
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if o.ParentHash.IsSet() {
		toSerialize["parentHash"] = o.ParentHash.Get()
	}
	if o.ParentUrl.IsSet() {
		toSerialize["parentUrl"] = o.ParentUrl.Get()
	}
	if !IsNil(o.ThreadHash) {
		toSerialize["threadHash"] = o.ThreadHash
	}
	if !IsNil(o.ParentAuthor) {
		toSerialize["parentAuthor"] = o.ParentAuthor
	}
	if !IsNil(o.MentionedProfiles) {
		toSerialize["mentionedProfiles"] = o.MentionedProfiles
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Embeds) {
		toSerialize["embeds"] = o.Embeds
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
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


