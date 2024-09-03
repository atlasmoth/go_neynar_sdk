/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListVerificationsByFid200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVerificationsByFid200Response{}

// ListVerificationsByFid200Response struct for ListVerificationsByFid200Response
type ListVerificationsByFid200Response struct {
	Messages []Verification `json:"messages"`
	NextPageToken string `json:"nextPageToken" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
}

type _ListVerificationsByFid200Response ListVerificationsByFid200Response

// NewListVerificationsByFid200Response instantiates a new ListVerificationsByFid200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVerificationsByFid200Response(messages []Verification, nextPageToken string) *ListVerificationsByFid200Response {
	this := ListVerificationsByFid200Response{}
	this.Messages = messages
	this.NextPageToken = nextPageToken
	return &this
}

// NewListVerificationsByFid200ResponseWithDefaults instantiates a new ListVerificationsByFid200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVerificationsByFid200ResponseWithDefaults() *ListVerificationsByFid200Response {
	this := ListVerificationsByFid200Response{}
	return &this
}

// GetMessages returns the Messages field value
func (o *ListVerificationsByFid200Response) GetMessages() []Verification {
	if o == nil {
		var ret []Verification
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *ListVerificationsByFid200Response) GetMessagesOk() ([]Verification, bool) {
	if o == nil {
		return nil, false
	}
	return o.Messages, true
}

// SetMessages sets field value
func (o *ListVerificationsByFid200Response) SetMessages(v []Verification) {
	o.Messages = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *ListVerificationsByFid200Response) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *ListVerificationsByFid200Response) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *ListVerificationsByFid200Response) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o ListVerificationsByFid200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVerificationsByFid200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["messages"] = o.Messages
	toSerialize["nextPageToken"] = o.NextPageToken
	return toSerialize, nil
}

func (o *ListVerificationsByFid200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"messages",
		"nextPageToken",
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

	varListVerificationsByFid200Response := _ListVerificationsByFid200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListVerificationsByFid200Response)

	if err != nil {
		return err
	}

	*o = ListVerificationsByFid200Response(varListVerificationsByFid200Response)

	return err
}

type NullableListVerificationsByFid200Response struct {
	value *ListVerificationsByFid200Response
	isSet bool
}

func (v NullableListVerificationsByFid200Response) Get() *ListVerificationsByFid200Response {
	return v.value
}

func (v *NullableListVerificationsByFid200Response) Set(val *ListVerificationsByFid200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListVerificationsByFid200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListVerificationsByFid200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVerificationsByFid200Response(val *ListVerificationsByFid200Response) *NullableListVerificationsByFid200Response {
	return &NullableListVerificationsByFid200Response{value: val, isSet: true}
}

func (v NullableListVerificationsByFid200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVerificationsByFid200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


