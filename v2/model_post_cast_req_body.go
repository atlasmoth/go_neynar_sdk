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

// checks if the PostCastReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCastReqBody{}

// PostCastReqBody struct for PostCastReqBody
type PostCastReqBody struct {
	// UUID of the signer
	SignerUuid *string `json:"signer_uuid,omitempty"`
	Text *string `json:"text,omitempty"`
	Embeds []EmbeddedCast `json:"embeds,omitempty"`
	// parent_url of the channel the cast is in, or hash of the cast
	Parent *string `json:"parent,omitempty"`
	// Channel ID of the channel where the cast is to be posted. e.g. neynar, farcaster, warpcast
	ChannelId *string `json:"channel_id,omitempty"`
	// An Idempotency key is a unique identifier for the request. **Note:**  1) This is used to prevent duplicate requests. Use the same idem key on retry attempts. 2) This should be a unique identifier for each request. 3) Recommended format is a 16-character string generated by the developer at the time of making this request. 
	Idem *string `json:"idem,omitempty"`
	// User identifier (unsigned integer)
	ParentAuthorFid *int32 `json:"parent_author_fid,omitempty"`
}

// NewPostCastReqBody instantiates a new PostCastReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCastReqBody() *PostCastReqBody {
	this := PostCastReqBody{}
	return &this
}

// NewPostCastReqBodyWithDefaults instantiates a new PostCastReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCastReqBodyWithDefaults() *PostCastReqBody {
	this := PostCastReqBody{}
	return &this
}

// GetSignerUuid returns the SignerUuid field value if set, zero value otherwise.
func (o *PostCastReqBody) GetSignerUuid() string {
	if o == nil || IsNil(o.SignerUuid) {
		var ret string
		return ret
	}
	return *o.SignerUuid
}

// GetSignerUuidOk returns a tuple with the SignerUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetSignerUuidOk() (*string, bool) {
	if o == nil || IsNil(o.SignerUuid) {
		return nil, false
	}
	return o.SignerUuid, true
}

// HasSignerUuid returns a boolean if a field has been set.
func (o *PostCastReqBody) HasSignerUuid() bool {
	if o != nil && !IsNil(o.SignerUuid) {
		return true
	}

	return false
}

// SetSignerUuid gets a reference to the given string and assigns it to the SignerUuid field.
func (o *PostCastReqBody) SetSignerUuid(v string) {
	o.SignerUuid = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *PostCastReqBody) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *PostCastReqBody) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *PostCastReqBody) SetText(v string) {
	o.Text = &v
}

// GetEmbeds returns the Embeds field value if set, zero value otherwise.
func (o *PostCastReqBody) GetEmbeds() []EmbeddedCast {
	if o == nil || IsNil(o.Embeds) {
		var ret []EmbeddedCast
		return ret
	}
	return o.Embeds
}

// GetEmbedsOk returns a tuple with the Embeds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetEmbedsOk() ([]EmbeddedCast, bool) {
	if o == nil || IsNil(o.Embeds) {
		return nil, false
	}
	return o.Embeds, true
}

// HasEmbeds returns a boolean if a field has been set.
func (o *PostCastReqBody) HasEmbeds() bool {
	if o != nil && !IsNil(o.Embeds) {
		return true
	}

	return false
}

// SetEmbeds gets a reference to the given []EmbeddedCast and assigns it to the Embeds field.
func (o *PostCastReqBody) SetEmbeds(v []EmbeddedCast) {
	o.Embeds = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *PostCastReqBody) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *PostCastReqBody) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *PostCastReqBody) SetParent(v string) {
	o.Parent = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *PostCastReqBody) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *PostCastReqBody) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *PostCastReqBody) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetIdem returns the Idem field value if set, zero value otherwise.
func (o *PostCastReqBody) GetIdem() string {
	if o == nil || IsNil(o.Idem) {
		var ret string
		return ret
	}
	return *o.Idem
}

// GetIdemOk returns a tuple with the Idem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetIdemOk() (*string, bool) {
	if o == nil || IsNil(o.Idem) {
		return nil, false
	}
	return o.Idem, true
}

// HasIdem returns a boolean if a field has been set.
func (o *PostCastReqBody) HasIdem() bool {
	if o != nil && !IsNil(o.Idem) {
		return true
	}

	return false
}

// SetIdem gets a reference to the given string and assigns it to the Idem field.
func (o *PostCastReqBody) SetIdem(v string) {
	o.Idem = &v
}

// GetParentAuthorFid returns the ParentAuthorFid field value if set, zero value otherwise.
func (o *PostCastReqBody) GetParentAuthorFid() int32 {
	if o == nil || IsNil(o.ParentAuthorFid) {
		var ret int32
		return ret
	}
	return *o.ParentAuthorFid
}

// GetParentAuthorFidOk returns a tuple with the ParentAuthorFid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCastReqBody) GetParentAuthorFidOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentAuthorFid) {
		return nil, false
	}
	return o.ParentAuthorFid, true
}

// HasParentAuthorFid returns a boolean if a field has been set.
func (o *PostCastReqBody) HasParentAuthorFid() bool {
	if o != nil && !IsNil(o.ParentAuthorFid) {
		return true
	}

	return false
}

// SetParentAuthorFid gets a reference to the given int32 and assigns it to the ParentAuthorFid field.
func (o *PostCastReqBody) SetParentAuthorFid(v int32) {
	o.ParentAuthorFid = &v
}

func (o PostCastReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCastReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignerUuid) {
		toSerialize["signer_uuid"] = o.SignerUuid
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Embeds) {
		toSerialize["embeds"] = o.Embeds
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.Idem) {
		toSerialize["idem"] = o.Idem
	}
	if !IsNil(o.ParentAuthorFid) {
		toSerialize["parent_author_fid"] = o.ParentAuthorFid
	}
	return toSerialize, nil
}

type NullablePostCastReqBody struct {
	value *PostCastReqBody
	isSet bool
}

func (v NullablePostCastReqBody) Get() *PostCastReqBody {
	return v.value
}

func (v *NullablePostCastReqBody) Set(val *PostCastReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCastReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCastReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCastReqBody(val *PostCastReqBody) *NullablePostCastReqBody {
	return &NullablePostCastReqBody{value: val, isSet: true}
}

func (v NullablePostCastReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCastReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


