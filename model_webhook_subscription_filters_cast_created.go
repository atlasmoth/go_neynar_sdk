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

// checks if the WebhookSubscriptionFiltersCastCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookSubscriptionFiltersCastCreated{}

// WebhookSubscriptionFiltersCastCreated struct for WebhookSubscriptionFiltersCastCreated
type WebhookSubscriptionFiltersCastCreated struct {
	// Exclude casts that matches these authors. **Note:** This is applied as an AND operation against rest of the filters. Rest of the filters are bundled as an OR operation. 
	ExcludeAuthorFids []int32 `json:"exclude_author_fids,omitempty"`
	AuthorFids []int32 `json:"author_fids,omitempty"`
	MentionedFids []int32 `json:"mentioned_fids,omitempty"`
	ParentUrls []string `json:"parent_urls,omitempty"`
	RootParentUrls []string `json:"root_parent_urls,omitempty"`
	ParentHashes []string `json:"parent_hashes,omitempty"`
	ParentAuthorFids []int32 `json:"parent_author_fids,omitempty"`
	// Regex pattern to match the text key of the cast. **Note:**  1) Regex must be parsed by Go's RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: (?i)\\\\$degen should be written as (?i)\\\\\\\\$degen 
	Text *string `json:"text,omitempty"`
	// Regex pattern to match the embeded_url (key embeds) of the cast. **Note:**  1) Regex must be parsed by Go's RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: \\\\b(farcaster|neynar)\\\\b should be written as \\\\\\\\b(farcaster|neynar)\\\\\\\\b 
	Embeds *string `json:"embeds,omitempty"`
}

// NewWebhookSubscriptionFiltersCastCreated instantiates a new WebhookSubscriptionFiltersCastCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSubscriptionFiltersCastCreated() *WebhookSubscriptionFiltersCastCreated {
	this := WebhookSubscriptionFiltersCastCreated{}
	return &this
}

// NewWebhookSubscriptionFiltersCastCreatedWithDefaults instantiates a new WebhookSubscriptionFiltersCastCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSubscriptionFiltersCastCreatedWithDefaults() *WebhookSubscriptionFiltersCastCreated {
	this := WebhookSubscriptionFiltersCastCreated{}
	return &this
}

// GetExcludeAuthorFids returns the ExcludeAuthorFids field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetExcludeAuthorFids() []int32 {
	if o == nil || IsNil(o.ExcludeAuthorFids) {
		var ret []int32
		return ret
	}
	return o.ExcludeAuthorFids
}

// GetExcludeAuthorFidsOk returns a tuple with the ExcludeAuthorFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetExcludeAuthorFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ExcludeAuthorFids) {
		return nil, false
	}
	return o.ExcludeAuthorFids, true
}

// HasExcludeAuthorFids returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasExcludeAuthorFids() bool {
	if o != nil && !IsNil(o.ExcludeAuthorFids) {
		return true
	}

	return false
}

// SetExcludeAuthorFids gets a reference to the given []int32 and assigns it to the ExcludeAuthorFids field.
func (o *WebhookSubscriptionFiltersCastCreated) SetExcludeAuthorFids(v []int32) {
	o.ExcludeAuthorFids = v
}

// GetAuthorFids returns the AuthorFids field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetAuthorFids() []int32 {
	if o == nil || IsNil(o.AuthorFids) {
		var ret []int32
		return ret
	}
	return o.AuthorFids
}

// GetAuthorFidsOk returns a tuple with the AuthorFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetAuthorFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.AuthorFids) {
		return nil, false
	}
	return o.AuthorFids, true
}

// HasAuthorFids returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasAuthorFids() bool {
	if o != nil && !IsNil(o.AuthorFids) {
		return true
	}

	return false
}

// SetAuthorFids gets a reference to the given []int32 and assigns it to the AuthorFids field.
func (o *WebhookSubscriptionFiltersCastCreated) SetAuthorFids(v []int32) {
	o.AuthorFids = v
}

// GetMentionedFids returns the MentionedFids field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetMentionedFids() []int32 {
	if o == nil || IsNil(o.MentionedFids) {
		var ret []int32
		return ret
	}
	return o.MentionedFids
}

// GetMentionedFidsOk returns a tuple with the MentionedFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetMentionedFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.MentionedFids) {
		return nil, false
	}
	return o.MentionedFids, true
}

// HasMentionedFids returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasMentionedFids() bool {
	if o != nil && !IsNil(o.MentionedFids) {
		return true
	}

	return false
}

// SetMentionedFids gets a reference to the given []int32 and assigns it to the MentionedFids field.
func (o *WebhookSubscriptionFiltersCastCreated) SetMentionedFids(v []int32) {
	o.MentionedFids = v
}

// GetParentUrls returns the ParentUrls field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetParentUrls() []string {
	if o == nil || IsNil(o.ParentUrls) {
		var ret []string
		return ret
	}
	return o.ParentUrls
}

// GetParentUrlsOk returns a tuple with the ParentUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetParentUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentUrls) {
		return nil, false
	}
	return o.ParentUrls, true
}

// HasParentUrls returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasParentUrls() bool {
	if o != nil && !IsNil(o.ParentUrls) {
		return true
	}

	return false
}

// SetParentUrls gets a reference to the given []string and assigns it to the ParentUrls field.
func (o *WebhookSubscriptionFiltersCastCreated) SetParentUrls(v []string) {
	o.ParentUrls = v
}

// GetRootParentUrls returns the RootParentUrls field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetRootParentUrls() []string {
	if o == nil || IsNil(o.RootParentUrls) {
		var ret []string
		return ret
	}
	return o.RootParentUrls
}

// GetRootParentUrlsOk returns a tuple with the RootParentUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetRootParentUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.RootParentUrls) {
		return nil, false
	}
	return o.RootParentUrls, true
}

// HasRootParentUrls returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasRootParentUrls() bool {
	if o != nil && !IsNil(o.RootParentUrls) {
		return true
	}

	return false
}

// SetRootParentUrls gets a reference to the given []string and assigns it to the RootParentUrls field.
func (o *WebhookSubscriptionFiltersCastCreated) SetRootParentUrls(v []string) {
	o.RootParentUrls = v
}

// GetParentHashes returns the ParentHashes field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetParentHashes() []string {
	if o == nil || IsNil(o.ParentHashes) {
		var ret []string
		return ret
	}
	return o.ParentHashes
}

// GetParentHashesOk returns a tuple with the ParentHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetParentHashesOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentHashes) {
		return nil, false
	}
	return o.ParentHashes, true
}

// HasParentHashes returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasParentHashes() bool {
	if o != nil && !IsNil(o.ParentHashes) {
		return true
	}

	return false
}

// SetParentHashes gets a reference to the given []string and assigns it to the ParentHashes field.
func (o *WebhookSubscriptionFiltersCastCreated) SetParentHashes(v []string) {
	o.ParentHashes = v
}

// GetParentAuthorFids returns the ParentAuthorFids field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetParentAuthorFids() []int32 {
	if o == nil || IsNil(o.ParentAuthorFids) {
		var ret []int32
		return ret
	}
	return o.ParentAuthorFids
}

// GetParentAuthorFidsOk returns a tuple with the ParentAuthorFids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetParentAuthorFidsOk() ([]int32, bool) {
	if o == nil || IsNil(o.ParentAuthorFids) {
		return nil, false
	}
	return o.ParentAuthorFids, true
}

// HasParentAuthorFids returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasParentAuthorFids() bool {
	if o != nil && !IsNil(o.ParentAuthorFids) {
		return true
	}

	return false
}

// SetParentAuthorFids gets a reference to the given []int32 and assigns it to the ParentAuthorFids field.
func (o *WebhookSubscriptionFiltersCastCreated) SetParentAuthorFids(v []int32) {
	o.ParentAuthorFids = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *WebhookSubscriptionFiltersCastCreated) SetText(v string) {
	o.Text = &v
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise.
func (o *WebhookSubscriptionFiltersCastCreated) GetEmbeds() string {
	if o == nil || IsNil(o.Embeds) {
		var ret string
		return ret
	}
	return *o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookSubscriptionFiltersCastCreated) GetEmbedsOk() (*string, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *WebhookSubscriptionFiltersCastCreated) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given string and assigns it to the Embeds field.
func (o *WebhookSubscriptionFiltersCastCreated) SetEmbeds(v string) {
	o.Embeds = &v
}

func (o WebhookSubscriptionFiltersCastCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookSubscriptionFiltersCastCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExcludeAuthorFids) {
		toSerialize["exclude_author_fids"] = o.ExcludeAuthorFids
	}
	if !IsNil(o.AuthorFids) {
		toSerialize["author_fids"] = o.AuthorFids
	}
	if !IsNil(o.MentionedFids) {
		toSerialize["mentioned_fids"] = o.MentionedFids
	}
	if !IsNil(o.ParentUrls) {
		toSerialize["parent_urls"] = o.ParentUrls
	}
	if !IsNil(o.RootParentUrls) {
		toSerialize["root_parent_urls"] = o.RootParentUrls
	}
	if !IsNil(o.ParentHashes) {
		toSerialize["parent_hashes"] = o.ParentHashes
	}
	if !IsNil(o.ParentAuthorFids) {
		toSerialize["parent_author_fids"] = o.ParentAuthorFids
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Embeds) {
		toSerialize["embeds"] = o.Embeds
	}
	return toSerialize, nil
}

type NullableWebhookSubscriptionFiltersCastCreated struct {
	value *WebhookSubscriptionFiltersCastCreated
	isSet bool
}

func (v NullableWebhookSubscriptionFiltersCastCreated) Get() *WebhookSubscriptionFiltersCastCreated {
	return v.value
}

func (v *NullableWebhookSubscriptionFiltersCastCreated) Set(val *WebhookSubscriptionFiltersCastCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSubscriptionFiltersCastCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSubscriptionFiltersCastCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSubscriptionFiltersCastCreated(val *WebhookSubscriptionFiltersCastCreated) *NullableWebhookSubscriptionFiltersCastCreated {
	return &NullableWebhookSubscriptionFiltersCastCreated{value: val, isSet: true}
}

func (v NullableWebhookSubscriptionFiltersCastCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSubscriptionFiltersCastCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


