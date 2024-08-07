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

// checks if the NextCursor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NextCursor{}

// NextCursor Returns next cursor
type NextCursor struct {
	Cursor NullableString `json:"cursor"`
}

type _NextCursor NextCursor

// NewNextCursor instantiates a new NextCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNextCursor(cursor NullableString) *NextCursor {
	this := NextCursor{}
	this.Cursor = cursor
	return &this
}

// NewNextCursorWithDefaults instantiates a new NextCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNextCursorWithDefaults() *NextCursor {
	this := NextCursor{}
	return &this
}

// GetCursor returns the Cursor field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NextCursor) GetCursor() string {
	if o == nil || o.Cursor.Get() == nil {
		var ret string
		return ret
	}

	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NextCursor) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// SetCursor sets field value
func (o *NextCursor) SetCursor(v string) {
	o.Cursor.Set(&v)
}

func (o NextCursor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NextCursor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor.Get()
	return toSerialize, nil
}

func (o *NextCursor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cursor",
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

	varNextCursor := _NextCursor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNextCursor)

	if err != nil {
		return err
	}

	*o = NextCursor(varNextCursor)

	return err
}

type NullableNextCursor struct {
	value *NextCursor
	isSet bool
}

func (v NullableNextCursor) Get() *NextCursor {
	return v.value
}

func (v *NullableNextCursor) Set(val *NextCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableNextCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableNextCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNextCursor(val *NextCursor) *NullableNextCursor {
	return &NullableNextCursor{value: val, isSet: true}
}

func (v NullableNextCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNextCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


