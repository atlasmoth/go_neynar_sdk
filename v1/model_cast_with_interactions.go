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

// checks if the CastWithInteractions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastWithInteractions{}

// CastWithInteractions struct for CastWithInteractions
type CastWithInteractions struct {
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
	Reactions *CastWithInteractionsReactions `json:"reactions,omitempty"`
	Recasts *CastWithInteractionsRecasts `json:"recasts,omitempty"`
	Recasters []string `json:"recasters,omitempty"`
	ViewerContext *ViewerContext `json:"viewerContext,omitempty"`
	Replies *CastWithInteractionsReplies `json:"replies,omitempty"`
}

// NewCastWithInteractions instantiates a new CastWithInteractions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastWithInteractions() *CastWithInteractions {
	this := CastWithInteractions{}
	return &this
}

// NewCastWithInteractionsWithDefaults instantiates a new CastWithInteractions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithInteractionsWithDefaults() *CastWithInteractions {
	this := CastWithInteractions{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *CastWithInteractions) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *CastWithInteractions) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *CastWithInteractions) SetHash(v string) {
	o.Hash = &v
}

// GetParentHash returns the ParentHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CastWithInteractions) GetParentHash() string {
	if o == nil || IsNil(o.ParentHash.Get()) {
		var ret string
		return ret
	}
	return *o.ParentHash.Get()
}

// GetParentHashOk returns a tuple with the ParentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentHash.Get(), o.ParentHash.IsSet()
}

// HasParentHash returns a boolean if a field has been set.
func (o *CastWithInteractions) HasParentHash() bool {
	if o != nil && o.ParentHash.IsSet() {
		return true
	}

	return false
}

// SetParentHash gets a reference to the given NullableString and assigns it to the ParentHash field.
func (o *CastWithInteractions) SetParentHash(v string) {
	o.ParentHash.Set(&v)
}
// SetParentHashNil sets the value for ParentHash to be an explicit nil
func (o *CastWithInteractions) SetParentHashNil() {
	o.ParentHash.Set(nil)
}

// UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
func (o *CastWithInteractions) UnsetParentHash() {
	o.ParentHash.Unset()
}

// GetParentUrl returns the ParentUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CastWithInteractions) GetParentUrl() string {
	if o == nil || IsNil(o.ParentUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ParentUrl.Get()
}

// GetParentUrlOk returns a tuple with the ParentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractions) GetParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentUrl.Get(), o.ParentUrl.IsSet()
}

// HasParentUrl returns a boolean if a field has been set.
func (o *CastWithInteractions) HasParentUrl() bool {
	if o != nil && o.ParentUrl.IsSet() {
		return true
	}

	return false
}

// SetParentUrl gets a reference to the given NullableString and assigns it to the ParentUrl field.
func (o *CastWithInteractions) SetParentUrl(v string) {
	o.ParentUrl.Set(&v)
}
// SetParentUrlNil sets the value for ParentUrl to be an explicit nil
func (o *CastWithInteractions) SetParentUrlNil() {
	o.ParentUrl.Set(nil)
}

// UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
func (o *CastWithInteractions) UnsetParentUrl() {
	o.ParentUrl.Unset()
}

// GetThreadHash returns the ThreadHash field value if set, zero value otherwise.
func (o *CastWithInteractions) GetThreadHash() string {
	if o == nil || IsNil(o.ThreadHash) {
		var ret string
		return ret
	}
	return *o.ThreadHash
}

// GetThreadHashOk returns a tuple with the ThreadHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetThreadHashOk() (*string, bool) {
	if o == nil || IsNil(o.ThreadHash) {
		return nil, false
	}
	return o.ThreadHash, true
}

// HasThreadHash returns a boolean if a field has been set.
func (o *CastWithInteractions) HasThreadHash() bool {
	if o != nil && !IsNil(o.ThreadHash) {
		return true
	}

	return false
}

// SetThreadHash gets a reference to the given string and assigns it to the ThreadHash field.
func (o *CastWithInteractions) SetThreadHash(v string) {
	o.ThreadHash = &v
}

// GetParentAuthor returns the ParentAuthor field value if set, zero value otherwise.
func (o *CastWithInteractions) GetParentAuthor() CastParentAuthor {
	if o == nil || IsNil(o.ParentAuthor) {
		var ret CastParentAuthor
		return ret
	}
	return *o.ParentAuthor
}

// GetParentAuthorOk returns a tuple with the ParentAuthor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetParentAuthorOk() (*CastParentAuthor, bool) {
	if o == nil || IsNil(o.ParentAuthor) {
		return nil, false
	}
	return o.ParentAuthor, true
}

// HasParentAuthor returns a boolean if a field has been set.
func (o *CastWithInteractions) HasParentAuthor() bool {
	if o != nil && !IsNil(o.ParentAuthor) {
		return true
	}

	return false
}

// SetParentAuthor gets a reference to the given CastParentAuthor and assigns it to the ParentAuthor field.
func (o *CastWithInteractions) SetParentAuthor(v CastParentAuthor) {
	o.ParentAuthor = &v
}

// GetMentionedProfiles returns the MentionedProfiles field value if set, zero value otherwise.
func (o *CastWithInteractions) GetMentionedProfiles() []User {
	if o == nil || IsNil(o.MentionedProfiles) {
		var ret []User
		return ret
	}
	return o.MentionedProfiles
}

// GetMentionedProfilesOk returns a tuple with the MentionedProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetMentionedProfilesOk() ([]User, bool) {
	if o == nil || IsNil(o.MentionedProfiles) {
		return nil, false
	}
	return o.MentionedProfiles, true
}

// HasMentionedProfiles returns a boolean if a field has been set.
func (o *CastWithInteractions) HasMentionedProfiles() bool {
	if o != nil && !IsNil(o.MentionedProfiles) {
		return true
	}

	return false
}

// SetMentionedProfiles gets a reference to the given []User and assigns it to the MentionedProfiles field.
func (o *CastWithInteractions) SetMentionedProfiles(v []User) {
	o.MentionedProfiles = v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *CastWithInteractions) GetAuthor() CastAuthor {
	if o == nil || IsNil(o.Author) {
		var ret CastAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetAuthorOk() (*CastAuthor, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *CastWithInteractions) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given CastAuthor and assigns it to the Author field.
func (o *CastWithInteractions) SetAuthor(v CastAuthor) {
	o.Author = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CastWithInteractions) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CastWithInteractions) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CastWithInteractions) SetText(v string) {
	o.Text = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CastWithInteractions) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CastWithInteractions) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *CastWithInteractions) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise.
func (o *CastWithInteractions) GetEmbeds() []EmbedUrl {
	if o == nil || IsNil(o.Embeds) {
		var ret []EmbedUrl
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetEmbedsOk() ([]EmbedUrl, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *CastWithInteractions) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []EmbedUrl and assigns it to the Embeds field.
func (o *CastWithInteractions) SetEmbeds(v []EmbedUrl) {
	o.Embeds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CastWithInteractions) GetType() CastType {
	if o == nil || IsNil(o.Type) {
		var ret CastType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetTypeOk() (*CastType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CastWithInteractions) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CastType and assigns it to the Type field.
func (o *CastWithInteractions) SetType(v CastType) {
	o.Type = &v
}

// GetReactions returns the Reactions field value if set, zero value otherwise.
func (o *CastWithInteractions) GetReactions() CastWithInteractionsReactions {
	if o == nil || IsNil(o.Reactions) {
		var ret CastWithInteractionsReactions
		return ret
	}
	return *o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetReactionsOk() (*CastWithInteractionsReactions, bool) {
	if o == nil || IsNil(o.Reactions) {
		return nil, false
	}
	return o.Reactions, true
}

// HasReactions returns a boolean if a field has been set.
func (o *CastWithInteractions) HasReactions() bool {
	if o != nil && !IsNil(o.Reactions) {
		return true
	}

	return false
}

// SetReactions gets a reference to the given CastWithInteractionsReactions and assigns it to the Reactions field.
func (o *CastWithInteractions) SetReactions(v CastWithInteractionsReactions) {
	o.Reactions = &v
}

// GetRecasts returns the Recasts field value if set, zero value otherwise.
func (o *CastWithInteractions) GetRecasts() CastWithInteractionsRecasts {
	if o == nil || IsNil(o.Recasts) {
		var ret CastWithInteractionsRecasts
		return ret
	}
	return *o.Recasts
}

// GetRecastsOk returns a tuple with the Recasts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetRecastsOk() (*CastWithInteractionsRecasts, bool) {
	if o == nil || IsNil(o.Recasts) {
		return nil, false
	}
	return o.Recasts, true
}

// HasRecasts returns a boolean if a field has been set.
func (o *CastWithInteractions) HasRecasts() bool {
	if o != nil && !IsNil(o.Recasts) {
		return true
	}

	return false
}

// SetRecasts gets a reference to the given CastWithInteractionsRecasts and assigns it to the Recasts field.
func (o *CastWithInteractions) SetRecasts(v CastWithInteractionsRecasts) {
	o.Recasts = &v
}

// GetRecasters returns the Recasters field value if set, zero value otherwise.
func (o *CastWithInteractions) GetRecasters() []string {
	if o == nil || IsNil(o.Recasters) {
		var ret []string
		return ret
	}
	return o.Recasters
}

// GetRecastersOk returns a tuple with the Recasters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetRecastersOk() ([]string, bool) {
	if o == nil || IsNil(o.Recasters) {
		return nil, false
	}
	return o.Recasters, true
}

// HasRecasters returns a boolean if a field has been set.
func (o *CastWithInteractions) HasRecasters() bool {
	if o != nil && !IsNil(o.Recasters) {
		return true
	}

	return false
}

// SetRecasters gets a reference to the given []string and assigns it to the Recasters field.
func (o *CastWithInteractions) SetRecasters(v []string) {
	o.Recasters = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *CastWithInteractions) GetViewerContext() ViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret ViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetViewerContextOk() (*ViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *CastWithInteractions) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given ViewerContext and assigns it to the ViewerContext field.
func (o *CastWithInteractions) SetViewerContext(v ViewerContext) {
	o.ViewerContext = &v
}

// GetReplies returns the Replies field value if set, zero value otherwise.
func (o *CastWithInteractions) GetReplies() CastWithInteractionsReplies {
	if o == nil || IsNil(o.Replies) {
		var ret CastWithInteractionsReplies
		return ret
	}
	return *o.Replies
}

// GetRepliesOk returns a tuple with the Replies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractions) GetRepliesOk() (*CastWithInteractionsReplies, bool) {
	if o == nil || IsNil(o.Replies) {
		return nil, false
	}
	return o.Replies, true
}

// HasReplies returns a boolean if a field has been set.
func (o *CastWithInteractions) HasReplies() bool {
	if o != nil && !IsNil(o.Replies) {
		return true
	}

	return false
}

// SetReplies gets a reference to the given CastWithInteractionsReplies and assigns it to the Replies field.
func (o *CastWithInteractions) SetReplies(v CastWithInteractionsReplies) {
	o.Replies = &v
}

func (o CastWithInteractions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastWithInteractions) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Reactions) {
		toSerialize["reactions"] = o.Reactions
	}
	if !IsNil(o.Recasts) {
		toSerialize["recasts"] = o.Recasts
	}
	if !IsNil(o.Recasters) {
		toSerialize["recasters"] = o.Recasters
	}
	if !IsNil(o.ViewerContext) {
		toSerialize["viewerContext"] = o.ViewerContext
	}
	if !IsNil(o.Replies) {
		toSerialize["replies"] = o.Replies
	}
	return toSerialize, nil
}

type NullableCastWithInteractions struct {
	value *CastWithInteractions
	isSet bool
}

func (v NullableCastWithInteractions) Get() *CastWithInteractions {
	return v.value
}

func (v *NullableCastWithInteractions) Set(val *CastWithInteractions) {
	v.value = val
	v.isSet = true
}

func (v NullableCastWithInteractions) IsSet() bool {
	return v.isSet
}

func (v *NullableCastWithInteractions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastWithInteractions(val *CastWithInteractions) *NullableCastWithInteractions {
	return &NullableCastWithInteractions{value: val, isSet: true}
}

func (v NullableCastWithInteractions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastWithInteractions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


