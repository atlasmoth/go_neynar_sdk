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

// checks if the NeynarPageButton type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarPageButton{}

// NeynarPageButton struct for NeynarPageButton
type NeynarPageButton struct {
	// The title of the button.
	Title string `json:"title"`
	// The index of the button, first button should have index 1 and so on.
	Index int32 `json:"index"`
	// The type of action that the button performs.
	ActionType string `json:"action_type"`
	NextPage *NeynarPageButtonNextPage `json:"next_page,omitempty"`
}

type _NeynarPageButton NeynarPageButton

// NewNeynarPageButton instantiates a new NeynarPageButton object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarPageButton(title string, index int32, actionType string) *NeynarPageButton {
	this := NeynarPageButton{}
	this.Title = title
	this.Index = index
	this.ActionType = actionType
	return &this
}

// NewNeynarPageButtonWithDefaults instantiates a new NeynarPageButton object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarPageButtonWithDefaults() *NeynarPageButton {
	this := NeynarPageButton{}
	return &this
}

// GetTitle returns the Title field value
func (o *NeynarPageButton) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *NeynarPageButton) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *NeynarPageButton) SetTitle(v string) {
	o.Title = v
}

// GetIndex returns the Index field value
func (o *NeynarPageButton) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *NeynarPageButton) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *NeynarPageButton) SetIndex(v int32) {
	o.Index = v
}

// GetActionType returns the ActionType field value
func (o *NeynarPageButton) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *NeynarPageButton) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *NeynarPageButton) SetActionType(v string) {
	o.ActionType = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *NeynarPageButton) GetNextPage() NeynarPageButtonNextPage {
	if o == nil || IsNil(o.NextPage) {
		var ret NeynarPageButtonNextPage
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeynarPageButton) GetNextPageOk() (*NeynarPageButtonNextPage, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *NeynarPageButton) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given NeynarPageButtonNextPage and assigns it to the NextPage field.
func (o *NeynarPageButton) SetNextPage(v NeynarPageButtonNextPage) {
	o.NextPage = &v
}

func (o NeynarPageButton) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarPageButton) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["index"] = o.Index
	toSerialize["action_type"] = o.ActionType
	if !IsNil(o.NextPage) {
		toSerialize["next_page"] = o.NextPage
	}
	return toSerialize, nil
}

func (o *NeynarPageButton) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"index",
		"action_type",
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

	varNeynarPageButton := _NeynarPageButton{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNeynarPageButton)

	if err != nil {
		return err
	}

	*o = NeynarPageButton(varNeynarPageButton)

	return err
}

type NullableNeynarPageButton struct {
	value *NeynarPageButton
	isSet bool
}

func (v NullableNeynarPageButton) Get() *NeynarPageButton {
	return v.value
}

func (v *NullableNeynarPageButton) Set(val *NeynarPageButton) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarPageButton) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarPageButton) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarPageButton(val *NeynarPageButton) *NullableNeynarPageButton {
	return &NullableNeynarPageButton{value: val, isSet: true}
}

func (v NullableNeynarPageButton) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarPageButton) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


