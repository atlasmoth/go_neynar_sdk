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

// checks if the Channel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Channel{}

// Channel struct for Channel
type Channel struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Object string `json:"object"`
	// Epoch timestamp in seconds.
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Number of followers the channel has.
	FollowerCount *float32 `json:"follower_count,omitempty"`
	ImageUrl *string `json:"image_url,omitempty"`
	ParentUrl *string `json:"parent_url,omitempty"`
	Lead *User `json:"lead,omitempty"`
	Moderator *User `json:"moderator,omitempty"`
	Hosts []User `json:"hosts,omitempty"`
	ViewerContext *ChannelViewerContext `json:"viewer_context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Channel Channel

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel(id string, url string, object string) *Channel {
	this := Channel{}
	this.Id = id
	this.Url = url
	this.Object = object
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetId returns the Id field value
func (o *Channel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Channel) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Channel) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Channel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Channel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Channel) SetUrl(v string) {
	o.Url = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Channel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Channel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Channel) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Channel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Channel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Channel) SetDescription(v string) {
	o.Description = &v
}

// GetObject returns the Object field value
func (o *Channel) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *Channel) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *Channel) SetObject(v string) {
	o.Object = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Channel) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Channel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Channel) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetFollowerCount returns the FollowerCount field value if set, zero value otherwise.
func (o *Channel) GetFollowerCount() float32 {
	if o == nil || IsNil(o.FollowerCount) {
		var ret float32
		return ret
	}
	return *o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetFollowerCountOk() (*float32, bool) {
	if o == nil || IsNil(o.FollowerCount) {
		return nil, false
	}
	return o.FollowerCount, true
}

// HasFollowerCount returns a boolean if a field has been set.
func (o *Channel) HasFollowerCount() bool {
	if o != nil && !IsNil(o.FollowerCount) {
		return true
	}

	return false
}

// SetFollowerCount gets a reference to the given float32 and assigns it to the FollowerCount field.
func (o *Channel) SetFollowerCount(v float32) {
	o.FollowerCount = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *Channel) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Channel) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *Channel) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetParentUrl returns the ParentUrl field value if set, zero value otherwise.
func (o *Channel) GetParentUrl() string {
	if o == nil || IsNil(o.ParentUrl) {
		var ret string
		return ret
	}
	return *o.ParentUrl
}

// GetParentUrlOk returns a tuple with the ParentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetParentUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ParentUrl) {
		return nil, false
	}
	return o.ParentUrl, true
}

// HasParentUrl returns a boolean if a field has been set.
func (o *Channel) HasParentUrl() bool {
	if o != nil && !IsNil(o.ParentUrl) {
		return true
	}

	return false
}

// SetParentUrl gets a reference to the given string and assigns it to the ParentUrl field.
func (o *Channel) SetParentUrl(v string) {
	o.ParentUrl = &v
}

// GetLead returns the Lead field value if set, zero value otherwise.
func (o *Channel) GetLead() User {
	if o == nil || IsNil(o.Lead) {
		var ret User
		return ret
	}
	return *o.Lead
}

// GetLeadOk returns a tuple with the Lead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetLeadOk() (*User, bool) {
	if o == nil || IsNil(o.Lead) {
		return nil, false
	}
	return o.Lead, true
}

// HasLead returns a boolean if a field has been set.
func (o *Channel) HasLead() bool {
	if o != nil && !IsNil(o.Lead) {
		return true
	}

	return false
}

// SetLead gets a reference to the given User and assigns it to the Lead field.
func (o *Channel) SetLead(v User) {
	o.Lead = &v
}

// GetModerator returns the Moderator field value if set, zero value otherwise.
func (o *Channel) GetModerator() User {
	if o == nil || IsNil(o.Moderator) {
		var ret User
		return ret
	}
	return *o.Moderator
}

// GetModeratorOk returns a tuple with the Moderator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetModeratorOk() (*User, bool) {
	if o == nil || IsNil(o.Moderator) {
		return nil, false
	}
	return o.Moderator, true
}

// HasModerator returns a boolean if a field has been set.
func (o *Channel) HasModerator() bool {
	if o != nil && !IsNil(o.Moderator) {
		return true
	}

	return false
}

// SetModerator gets a reference to the given User and assigns it to the Moderator field.
func (o *Channel) SetModerator(v User) {
	o.Moderator = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *Channel) GetHosts() []User {
	if o == nil || IsNil(o.Hosts) {
		var ret []User
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetHostsOk() ([]User, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *Channel) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []User and assigns it to the Hosts field.
func (o *Channel) SetHosts(v []User) {
	o.Hosts = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *Channel) GetViewerContext() ChannelViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret ChannelViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetViewerContextOk() (*ChannelViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *Channel) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given ChannelViewerContext and assigns it to the ViewerContext field.
func (o *Channel) SetViewerContext(v ChannelViewerContext) {
	o.ViewerContext = &v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Channel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["object"] = o.Object
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.FollowerCount) {
		toSerialize["follower_count"] = o.FollowerCount
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.ParentUrl) {
		toSerialize["parent_url"] = o.ParentUrl
	}
	if !IsNil(o.Lead) {
		toSerialize["lead"] = o.Lead
	}
	if !IsNil(o.Moderator) {
		toSerialize["moderator"] = o.Moderator
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.ViewerContext) {
		toSerialize["viewer_context"] = o.ViewerContext
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Channel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"object",
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

	varChannel := _Channel{}

	err = json.Unmarshal(data, &varChannel)

	if err != nil {
		return err
	}

	*o = Channel(varChannel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "object")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "follower_count")
		delete(additionalProperties, "image_url")
		delete(additionalProperties, "parent_url")
		delete(additionalProperties, "lead")
		delete(additionalProperties, "moderator")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "viewer_context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


