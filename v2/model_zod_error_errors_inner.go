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

// checks if the ZodErrorErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ZodErrorErrorsInner{}

// ZodErrorErrorsInner struct for ZodErrorErrorsInner
type ZodErrorErrorsInner struct {
	Code *string `json:"code,omitempty"`
	Expected *string `json:"expected,omitempty"`
	Received *string `json:"received,omitempty"`
	Path []string `json:"path,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewZodErrorErrorsInner instantiates a new ZodErrorErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZodErrorErrorsInner() *ZodErrorErrorsInner {
	this := ZodErrorErrorsInner{}
	return &this
}

// NewZodErrorErrorsInnerWithDefaults instantiates a new ZodErrorErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZodErrorErrorsInnerWithDefaults() *ZodErrorErrorsInner {
	this := ZodErrorErrorsInner{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ZodErrorErrorsInner) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZodErrorErrorsInner) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ZodErrorErrorsInner) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ZodErrorErrorsInner) SetCode(v string) {
	o.Code = &v
}

// GetExpected returns the Expected field value if set, zero value otherwise.
func (o *ZodErrorErrorsInner) GetExpected() string {
	if o == nil || IsNil(o.Expected) {
		var ret string
		return ret
	}
	return *o.Expected
}

// GetExpectedOk returns a tuple with the Expected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZodErrorErrorsInner) GetExpectedOk() (*string, bool) {
	if o == nil || IsNil(o.Expected) {
		return nil, false
	}
	return o.Expected, true
}

// HasExpected returns a boolean if a field has been set.
func (o *ZodErrorErrorsInner) HasExpected() bool {
	if o != nil && !IsNil(o.Expected) {
		return true
	}

	return false
}

// SetExpected gets a reference to the given string and assigns it to the Expected field.
func (o *ZodErrorErrorsInner) SetExpected(v string) {
	o.Expected = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *ZodErrorErrorsInner) GetReceived() string {
	if o == nil || IsNil(o.Received) {
		var ret string
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZodErrorErrorsInner) GetReceivedOk() (*string, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *ZodErrorErrorsInner) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given string and assigns it to the Received field.
func (o *ZodErrorErrorsInner) SetReceived(v string) {
	o.Received = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ZodErrorErrorsInner) GetPath() []string {
	if o == nil || IsNil(o.Path) {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZodErrorErrorsInner) GetPathOk() ([]string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ZodErrorErrorsInner) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *ZodErrorErrorsInner) SetPath(v []string) {
	o.Path = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ZodErrorErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ZodErrorErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ZodErrorErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ZodErrorErrorsInner) SetMessage(v string) {
	o.Message = &v
}

func (o ZodErrorErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ZodErrorErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Expected) {
		toSerialize["expected"] = o.Expected
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableZodErrorErrorsInner struct {
	value *ZodErrorErrorsInner
	isSet bool
}

func (v NullableZodErrorErrorsInner) Get() *ZodErrorErrorsInner {
	return v.value
}

func (v *NullableZodErrorErrorsInner) Set(val *ZodErrorErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableZodErrorErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableZodErrorErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZodErrorErrorsInner(val *ZodErrorErrorsInner) *NullableZodErrorErrorsInner {
	return &NullableZodErrorErrorsInner{value: val, isSet: true}
}

func (v NullableZodErrorErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZodErrorErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


