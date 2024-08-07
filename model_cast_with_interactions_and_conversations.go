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
	"bytes"
	"fmt"
)

// checks if the CastWithInteractionsAndConversations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CastWithInteractionsAndConversations{}

// CastWithInteractionsAndConversations struct for CastWithInteractionsAndConversations
type CastWithInteractionsAndConversations struct {
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
	Frames []Frame `json:"frames,omitempty"`
	Reactions CastWithInteractionsReactions `json:"reactions"`
	Replies CastWithInteractionsReplies `json:"replies"`
	ThreadHash NullableString `json:"thread_hash"`
	MentionedProfiles []User `json:"mentioned_profiles"`
	Channel ChannelOrDehydratedChannel `json:"channel"`
	ViewerContext *CastViewerContext `json:"viewer_context,omitempty"`
	// note: This is recursive. It contains the direct replies to the cast and their direct replies up to n reply_depth.
	DirectReplies []CastWithInteractionsAndConversations `json:"direct_replies"`
}

type _CastWithInteractionsAndConversations CastWithInteractionsAndConversations

// NewCastWithInteractionsAndConversations instantiates a new CastWithInteractionsAndConversations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCastWithInteractionsAndConversations(hash string, parentHash NullableString, parentUrl NullableString, rootParentUrl NullableString, parentAuthor CastParentAuthor, author User, text string, timestamp time.Time, embeds []EmbeddedCast, reactions CastWithInteractionsReactions, replies CastWithInteractionsReplies, threadHash NullableString, mentionedProfiles []User, channel ChannelOrDehydratedChannel, directReplies []CastWithInteractionsAndConversations) *CastWithInteractionsAndConversations {
	this := CastWithInteractionsAndConversations{}
	this.Hash = hash
	this.ParentHash = parentHash
	this.ParentUrl = parentUrl
	this.RootParentUrl = rootParentUrl
	this.ParentAuthor = parentAuthor
	this.Author = author
	this.Text = text
	this.Timestamp = timestamp
	this.Embeds = embeds
	this.Reactions = reactions
	this.Replies = replies
	this.ThreadHash = threadHash
	this.MentionedProfiles = mentionedProfiles
	this.Channel = channel
	this.DirectReplies = directReplies
	return &this
}

// NewCastWithInteractionsAndConversationsWithDefaults instantiates a new CastWithInteractionsAndConversations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCastWithInteractionsAndConversationsWithDefaults() *CastWithInteractionsAndConversations {
	this := CastWithInteractionsAndConversations{}
	return &this
}

// GetHash returns the Hash field value
func (o *CastWithInteractionsAndConversations) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *CastWithInteractionsAndConversations) SetHash(v string) {
	o.Hash = v
}

// GetParentHash returns the ParentHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractionsAndConversations) GetParentHash() string {
	if o == nil || o.ParentHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentHash.Get()
}

// GetParentHashOk returns a tuple with the ParentHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractionsAndConversations) GetParentHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentHash.Get(), o.ParentHash.IsSet()
}

// SetParentHash sets field value
func (o *CastWithInteractionsAndConversations) SetParentHash(v string) {
	o.ParentHash.Set(&v)
}

// GetParentUrl returns the ParentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractionsAndConversations) GetParentUrl() string {
	if o == nil || o.ParentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ParentUrl.Get()
}

// GetParentUrlOk returns a tuple with the ParentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractionsAndConversations) GetParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentUrl.Get(), o.ParentUrl.IsSet()
}

// SetParentUrl sets field value
func (o *CastWithInteractionsAndConversations) SetParentUrl(v string) {
	o.ParentUrl.Set(&v)
}

// GetRootParentUrl returns the RootParentUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractionsAndConversations) GetRootParentUrl() string {
	if o == nil || o.RootParentUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.RootParentUrl.Get()
}

// GetRootParentUrlOk returns a tuple with the RootParentUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractionsAndConversations) GetRootParentUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RootParentUrl.Get(), o.RootParentUrl.IsSet()
}

// SetRootParentUrl sets field value
func (o *CastWithInteractionsAndConversations) SetRootParentUrl(v string) {
	o.RootParentUrl.Set(&v)
}

// GetParentAuthor returns the ParentAuthor field value
func (o *CastWithInteractionsAndConversations) GetParentAuthor() CastParentAuthor {
	if o == nil {
		var ret CastParentAuthor
		return ret
	}

	return o.ParentAuthor
}

// GetParentAuthorOk returns a tuple with the ParentAuthor field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetParentAuthorOk() (*CastParentAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentAuthor, true
}

// SetParentAuthor sets field value
func (o *CastWithInteractionsAndConversations) SetParentAuthor(v CastParentAuthor) {
	o.ParentAuthor = v
}

// GetAuthor returns the Author field value
func (o *CastWithInteractionsAndConversations) GetAuthor() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetAuthorOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *CastWithInteractionsAndConversations) SetAuthor(v User) {
	o.Author = v
}

// GetText returns the Text field value
func (o *CastWithInteractionsAndConversations) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *CastWithInteractionsAndConversations) SetText(v string) {
	o.Text = v
}

// GetTimestamp returns the Timestamp field value
func (o *CastWithInteractionsAndConversations) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *CastWithInteractionsAndConversations) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEmbeds returns the Embeds field value
func (o *CastWithInteractionsAndConversations) GetEmbeds() []EmbeddedCast {
	if o == nil {
		var ret []EmbeddedCast
		return ret
	}

	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetEmbedsOk() ([]EmbeddedCast, bool) {
	if o == nil {
		return nil, false
	}
	return o.Embeds, true
}

// SetEmbeds sets field value
func (o *CastWithInteractionsAndConversations) SetEmbeds(v []EmbeddedCast) {
	o.Embeds = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CastWithInteractionsAndConversations) GetType() CastNotificationType {
	if o == nil || IsNil(o.Type) {
		var ret CastNotificationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetTypeOk() (*CastNotificationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CastWithInteractionsAndConversations) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CastNotificationType and assigns it to the Type field.
func (o *CastWithInteractionsAndConversations) SetType(v CastNotificationType) {
	o.Type = &v
}

// GetFrames returns the Frames field value if set, zero value otherwise.
func (o *CastWithInteractionsAndConversations) GetFrames() []Frame {
	if o == nil || IsNil(o.Frames) {
		var ret []Frame
		return ret
	}
	return o.Frames
}

// GetFramesOk returns a tuple with the Frames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetFramesOk() ([]Frame, bool) {
	if o == nil || IsNil(o.Frames) {
		return nil, false
	}
	return o.Frames, true
}

// HasFrames returns a boolean if a field has been set.
func (o *CastWithInteractionsAndConversations) HasFrames() bool {
	if o != nil && !IsNil(o.Frames) {
		return true
	}

	return false
}

// SetFrames gets a reference to the given []Frame and assigns it to the Frames field.
func (o *CastWithInteractionsAndConversations) SetFrames(v []Frame) {
	o.Frames = v
}

// GetReactions returns the Reactions field value
func (o *CastWithInteractionsAndConversations) GetReactions() CastWithInteractionsReactions {
	if o == nil {
		var ret CastWithInteractionsReactions
		return ret
	}

	return o.Reactions
}

// GetReactionsOk returns a tuple with the Reactions field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetReactionsOk() (*CastWithInteractionsReactions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reactions, true
}

// SetReactions sets field value
func (o *CastWithInteractionsAndConversations) SetReactions(v CastWithInteractionsReactions) {
	o.Reactions = v
}

// GetReplies returns the Replies field value
func (o *CastWithInteractionsAndConversations) GetReplies() CastWithInteractionsReplies {
	if o == nil {
		var ret CastWithInteractionsReplies
		return ret
	}

	return o.Replies
}

// GetRepliesOk returns a tuple with the Replies field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetRepliesOk() (*CastWithInteractionsReplies, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replies, true
}

// SetReplies sets field value
func (o *CastWithInteractionsAndConversations) SetReplies(v CastWithInteractionsReplies) {
	o.Replies = v
}

// GetThreadHash returns the ThreadHash field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CastWithInteractionsAndConversations) GetThreadHash() string {
	if o == nil || o.ThreadHash.Get() == nil {
		var ret string
		return ret
	}

	return *o.ThreadHash.Get()
}

// GetThreadHashOk returns a tuple with the ThreadHash field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CastWithInteractionsAndConversations) GetThreadHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadHash.Get(), o.ThreadHash.IsSet()
}

// SetThreadHash sets field value
func (o *CastWithInteractionsAndConversations) SetThreadHash(v string) {
	o.ThreadHash.Set(&v)
}

// GetMentionedProfiles returns the MentionedProfiles field value
func (o *CastWithInteractionsAndConversations) GetMentionedProfiles() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.MentionedProfiles
}

// GetMentionedProfilesOk returns a tuple with the MentionedProfiles field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetMentionedProfilesOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.MentionedProfiles, true
}

// SetMentionedProfiles sets field value
func (o *CastWithInteractionsAndConversations) SetMentionedProfiles(v []User) {
	o.MentionedProfiles = v
}

// GetChannel returns the Channel field value
func (o *CastWithInteractionsAndConversations) GetChannel() ChannelOrDehydratedChannel {
	if o == nil {
		var ret ChannelOrDehydratedChannel
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetChannelOk() (*ChannelOrDehydratedChannel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *CastWithInteractionsAndConversations) SetChannel(v ChannelOrDehydratedChannel) {
	o.Channel = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *CastWithInteractionsAndConversations) GetViewerContext() CastViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret CastViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetViewerContextOk() (*CastViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *CastWithInteractionsAndConversations) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given CastViewerContext and assigns it to the ViewerContext field.
func (o *CastWithInteractionsAndConversations) SetViewerContext(v CastViewerContext) {
	o.ViewerContext = &v
}

// GetDirectReplies returns the DirectReplies field value
func (o *CastWithInteractionsAndConversations) GetDirectReplies() []CastWithInteractionsAndConversations {
	if o == nil {
		var ret []CastWithInteractionsAndConversations
		return ret
	}

	return o.DirectReplies
}

// GetDirectRepliesOk returns a tuple with the DirectReplies field value
// and a boolean to check if the value has been set.
func (o *CastWithInteractionsAndConversations) GetDirectRepliesOk() ([]CastWithInteractionsAndConversations, bool) {
	if o == nil {
		return nil, false
	}
	return o.DirectReplies, true
}

// SetDirectReplies sets field value
func (o *CastWithInteractionsAndConversations) SetDirectReplies(v []CastWithInteractionsAndConversations) {
	o.DirectReplies = v
}

func (o CastWithInteractionsAndConversations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CastWithInteractionsAndConversations) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Frames) {
		toSerialize["frames"] = o.Frames
	}
	toSerialize["reactions"] = o.Reactions
	toSerialize["replies"] = o.Replies
	toSerialize["thread_hash"] = o.ThreadHash.Get()
	toSerialize["mentioned_profiles"] = o.MentionedProfiles
	toSerialize["channel"] = o.Channel
	if !IsNil(o.ViewerContext) {
		toSerialize["viewer_context"] = o.ViewerContext
	}
	toSerialize["direct_replies"] = o.DirectReplies
	return toSerialize, nil
}

func (o *CastWithInteractionsAndConversations) UnmarshalJSON(data []byte) (err error) {
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
		"reactions",
		"replies",
		"thread_hash",
		"mentioned_profiles",
		"channel",
		"direct_replies",
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

	varCastWithInteractionsAndConversations := _CastWithInteractionsAndConversations{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCastWithInteractionsAndConversations)

	if err != nil {
		return err
	}

	*o = CastWithInteractionsAndConversations(varCastWithInteractionsAndConversations)

	return err
}

type NullableCastWithInteractionsAndConversations struct {
	value *CastWithInteractionsAndConversations
	isSet bool
}

func (v NullableCastWithInteractionsAndConversations) Get() *CastWithInteractionsAndConversations {
	return v.value
}

func (v *NullableCastWithInteractionsAndConversations) Set(val *CastWithInteractionsAndConversations) {
	v.value = val
	v.isSet = true
}

func (v NullableCastWithInteractionsAndConversations) IsSet() bool {
	return v.isSet
}

func (v *NullableCastWithInteractionsAndConversations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCastWithInteractionsAndConversations(val *CastWithInteractionsAndConversations) *NullableCastWithInteractionsAndConversations {
	return &NullableCastWithInteractionsAndConversations{value: val, isSet: true}
}

func (v NullableCastWithInteractionsAndConversations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCastWithInteractionsAndConversations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


