/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the NeynarFrame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarFrame{}

// NeynarFrame struct for NeynarFrame
type NeynarFrame struct {
	// Unique identifier for the frame.
	Uuid string `json:"uuid"`
	// Name of the frame.
	Name string `json:"name"`
	// Generated link for the frame's first page.
	Link  string            `json:"link"`
	Pages []NeynarFramePage `json:"pages"`
	// Indicates if the frame is valid.
	Valid *bool `json:"valid,omitempty"`
}

type _NeynarFrame NeynarFrame

// NewNeynarFrame instantiates a new NeynarFrame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarFrame(uuid string, name string, link string, pages []NeynarFramePage) *NeynarFrame {
	this := NeynarFrame{}
	this.Uuid = uuid
	this.Name = name
	this.Link = link
	this.Pages = pages
	return &this
}

// NewNeynarFrameWithDefaults instantiates a new NeynarFrame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarFrameWithDefaults() *NeynarFrame {
	this := NeynarFrame{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *NeynarFrame) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *NeynarFrame) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *NeynarFrame) SetUuid(v string) {
	o.Uuid = v
}

// GetName returns the Name field value
func (o *NeynarFrame) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NeynarFrame) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NeynarFrame) SetName(v string) {
	o.Name = v
}

// GetLink returns the Link field value
func (o *NeynarFrame) GetLink() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Link
}

// GetLinkOk returns a tuple with the Link field value
// and a boolean to check if the value has been set.
func (o *NeynarFrame) GetLinkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Link, true
}

// SetLink sets field value
func (o *NeynarFrame) SetLink(v string) {
	o.Link = v
}

// GetPages returns the Pages field value
func (o *NeynarFrame) GetPages() []NeynarFramePage {
	if o == nil {
		var ret []NeynarFramePage
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *NeynarFrame) GetPagesOk() ([]NeynarFramePage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *NeynarFrame) SetPages(v []NeynarFramePage) {
	o.Pages = v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *NeynarFrame) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NeynarFrame) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *NeynarFrame) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *NeynarFrame) SetValid(v bool) {
	o.Valid = &v
}

func (o NeynarFrame) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarFrame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	toSerialize["name"] = o.Name
	toSerialize["link"] = o.Link
	toSerialize["pages"] = o.Pages
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

func (o *NeynarFrame) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
		"name",
		"link",
		"pages",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNeynarFrame := _NeynarFrame{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNeynarFrame)

	if err != nil {
		return err
	}

	*o = NeynarFrame(varNeynarFrame)

	return err
}

type NullableNeynarFrame struct {
	value *NeynarFrame
	isSet bool
}

func (v NullableNeynarFrame) Get() *NeynarFrame {
	return v.value
}

func (v *NullableNeynarFrame) Set(val *NeynarFrame) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarFrame) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarFrame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarFrame(val *NeynarFrame) *NullableNeynarFrame {
	return &NullableNeynarFrame{value: val, isSet: true}
}

func (v NullableNeynarFrame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarFrame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
