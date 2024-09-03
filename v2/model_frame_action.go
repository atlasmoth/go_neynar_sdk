/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FrameAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameAction{}

// FrameAction struct for FrameAction
type FrameAction struct {
	Version *string `json:"version,omitempty"`
	Title *string `json:"title,omitempty"`
	Image *string `json:"image,omitempty"`
	Button FrameActionButton `json:"button"`
	Input *FrameInput `json:"input,omitempty"`
	State *FrameState `json:"state,omitempty"`
	Transaction *FrameTransaction `json:"transaction,omitempty"`
	// The connected wallet address of the interacting user.
	Address *string `json:"address,omitempty"`
	// URL of the frames
	FramesUrl string `json:"frames_url"`
	// URL of the post to get the next frame
	PostUrl string `json:"post_url"`
}

type _FrameAction FrameAction

// NewFrameAction instantiates a new FrameAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameAction(button FrameActionButton, framesUrl string, postUrl string) *FrameAction {
	this := FrameAction{}
	this.Button = button
	this.FramesUrl = framesUrl
	this.PostUrl = postUrl
	return &this
}

// NewFrameActionWithDefaults instantiates a new FrameAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameActionWithDefaults() *FrameAction {
	this := FrameAction{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FrameAction) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FrameAction) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FrameAction) SetVersion(v string) {
	o.Version = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FrameAction) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FrameAction) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FrameAction) SetTitle(v string) {
	o.Title = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *FrameAction) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *FrameAction) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *FrameAction) SetImage(v string) {
	o.Image = &v
}

// GetButton returns the Button field value
func (o *FrameAction) GetButton() FrameActionButton {
	if o == nil {
		var ret FrameActionButton
		return ret
	}

	return o.Button
}

// GetButtonOk returns a tuple with the Button field value
// and a boolean to check if the value has been set.
func (o *FrameAction) GetButtonOk() (*FrameActionButton, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Button, true
}

// SetButton sets field value
func (o *FrameAction) SetButton(v FrameActionButton) {
	o.Button = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *FrameAction) GetInput() FrameInput {
	if o == nil || IsNil(o.Input) {
		var ret FrameInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetInputOk() (*FrameInput, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *FrameAction) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given FrameInput and assigns it to the Input field.
func (o *FrameAction) SetInput(v FrameInput) {
	o.Input = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FrameAction) GetState() FrameState {
	if o == nil || IsNil(o.State) {
		var ret FrameState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetStateOk() (*FrameState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FrameAction) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given FrameState and assigns it to the State field.
func (o *FrameAction) SetState(v FrameState) {
	o.State = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *FrameAction) GetTransaction() FrameTransaction {
	if o == nil || IsNil(o.Transaction) {
		var ret FrameTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetTransactionOk() (*FrameTransaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *FrameAction) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given FrameTransaction and assigns it to the Transaction field.
func (o *FrameAction) SetTransaction(v FrameTransaction) {
	o.Transaction = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FrameAction) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameAction) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FrameAction) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *FrameAction) SetAddress(v string) {
	o.Address = &v
}

// GetFramesUrl returns the FramesUrl field value
func (o *FrameAction) GetFramesUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FramesUrl
}

// GetFramesUrlOk returns a tuple with the FramesUrl field value
// and a boolean to check if the value has been set.
func (o *FrameAction) GetFramesUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FramesUrl, true
}

// SetFramesUrl sets field value
func (o *FrameAction) SetFramesUrl(v string) {
	o.FramesUrl = v
}

// GetPostUrl returns the PostUrl field value
func (o *FrameAction) GetPostUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostUrl
}

// GetPostUrlOk returns a tuple with the PostUrl field value
// and a boolean to check if the value has been set.
func (o *FrameAction) GetPostUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostUrl, true
}

// SetPostUrl sets field value
func (o *FrameAction) SetPostUrl(v string) {
	o.PostUrl = v
}

func (o FrameAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	toSerialize["button"] = o.Button
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	toSerialize["frames_url"] = o.FramesUrl
	toSerialize["post_url"] = o.PostUrl
	return toSerialize, nil
}

func (o *FrameAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"button",
		"frames_url",
		"post_url",
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

	varFrameAction := _FrameAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameAction)

	if err != nil {
		return err
	}

	*o = FrameAction(varFrameAction)

	return err
}

type NullableFrameAction struct {
	value *FrameAction
	isSet bool
}

func (v NullableFrameAction) Get() *FrameAction {
	return v.value
}

func (v *NullableFrameAction) Set(val *FrameAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameAction(val *FrameAction) *NullableFrameAction {
	return &NullableFrameAction{value: val, isSet: true}
}

func (v NullableFrameAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

