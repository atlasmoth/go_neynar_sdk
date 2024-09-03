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

// checks if the FidsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FidsResponse{}

// FidsResponse struct for FidsResponse
type FidsResponse struct {
	Fids []int32 `json:"fids"`
	NextPageToken string `json:"nextPageToken" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
}

type _FidsResponse FidsResponse

// NewFidsResponse instantiates a new FidsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFidsResponse(fids []int32, nextPageToken string) *FidsResponse {
	this := FidsResponse{}
	this.Fids = fids
	this.NextPageToken = nextPageToken
	return &this
}

// NewFidsResponseWithDefaults instantiates a new FidsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFidsResponseWithDefaults() *FidsResponse {
	this := FidsResponse{}
	return &this
}

// GetFids returns the Fids field value
func (o *FidsResponse) GetFids() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Fids
}

// GetFidsOk returns a tuple with the Fids field value
// and a boolean to check if the value has been set.
func (o *FidsResponse) GetFidsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fids, true
}

// SetFids sets field value
func (o *FidsResponse) SetFids(v []int32) {
	o.Fids = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *FidsResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *FidsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *FidsResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o FidsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FidsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fids"] = o.Fids
	toSerialize["nextPageToken"] = o.NextPageToken
	return toSerialize, nil
}

func (o *FidsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fids",
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

	varFidsResponse := _FidsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFidsResponse)

	if err != nil {
		return err
	}

	*o = FidsResponse(varFidsResponse)

	return err
}

type NullableFidsResponse struct {
	value *FidsResponse
	isSet bool
}

func (v NullableFidsResponse) Get() *FidsResponse {
	return v.value
}

func (v *NullableFidsResponse) Set(val *FidsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFidsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFidsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFidsResponse(val *FidsResponse) *NullableFidsResponse {
	return &NullableFidsResponse{value: val, isSet: true}
}

func (v NullableFidsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFidsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

