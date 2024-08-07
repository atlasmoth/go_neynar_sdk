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

// checks if the NeynarFrameCreationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NeynarFrameCreationRequest{}

// NeynarFrameCreationRequest struct for NeynarFrameCreationRequest
type NeynarFrameCreationRequest struct {
	// The name of the frame.
	Name string `json:"name"`
	Pages []NeynarFramePage `json:"pages"`
}

type _NeynarFrameCreationRequest NeynarFrameCreationRequest

// NewNeynarFrameCreationRequest instantiates a new NeynarFrameCreationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNeynarFrameCreationRequest(name string, pages []NeynarFramePage) *NeynarFrameCreationRequest {
	this := NeynarFrameCreationRequest{}
	this.Name = name
	this.Pages = pages
	return &this
}

// NewNeynarFrameCreationRequestWithDefaults instantiates a new NeynarFrameCreationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNeynarFrameCreationRequestWithDefaults() *NeynarFrameCreationRequest {
	this := NeynarFrameCreationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *NeynarFrameCreationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NeynarFrameCreationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NeynarFrameCreationRequest) SetName(v string) {
	o.Name = v
}

// GetPages returns the Pages field value
func (o *NeynarFrameCreationRequest) GetPages() []NeynarFramePage {
	if o == nil {
		var ret []NeynarFramePage
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *NeynarFrameCreationRequest) GetPagesOk() ([]NeynarFramePage, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pages, true
}

// SetPages sets field value
func (o *NeynarFrameCreationRequest) SetPages(v []NeynarFramePage) {
	o.Pages = v
}

func (o NeynarFrameCreationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NeynarFrameCreationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["pages"] = o.Pages
	return toSerialize, nil
}

func (o *NeynarFrameCreationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"pages",
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

	varNeynarFrameCreationRequest := _NeynarFrameCreationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNeynarFrameCreationRequest)

	if err != nil {
		return err
	}

	*o = NeynarFrameCreationRequest(varNeynarFrameCreationRequest)

	return err
}

type NullableNeynarFrameCreationRequest struct {
	value *NeynarFrameCreationRequest
	isSet bool
}

func (v NullableNeynarFrameCreationRequest) Get() *NeynarFrameCreationRequest {
	return v.value
}

func (v *NullableNeynarFrameCreationRequest) Set(val *NeynarFrameCreationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNeynarFrameCreationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNeynarFrameCreationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNeynarFrameCreationRequest(val *NeynarFrameCreationRequest) *NullableNeynarFrameCreationRequest {
	return &NullableNeynarFrameCreationRequest{value: val, isSet: true}
}

func (v NullableNeynarFrameCreationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNeynarFrameCreationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


