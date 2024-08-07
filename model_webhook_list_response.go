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

// checks if the WebhookListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookListResponse{}

// WebhookListResponse struct for WebhookListResponse
type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type _WebhookListResponse WebhookListResponse

// NewWebhookListResponse instantiates a new WebhookListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookListResponse(webhooks []Webhook) *WebhookListResponse {
	this := WebhookListResponse{}
	this.Webhooks = webhooks
	return &this
}

// NewWebhookListResponseWithDefaults instantiates a new WebhookListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookListResponseWithDefaults() *WebhookListResponse {
	this := WebhookListResponse{}
	return &this
}

// GetWebhooks returns the Webhooks field value
func (o *WebhookListResponse) GetWebhooks() []Webhook {
	if o == nil {
		var ret []Webhook
		return ret
	}

	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value
// and a boolean to check if the value has been set.
func (o *WebhookListResponse) GetWebhooksOk() ([]Webhook, bool) {
	if o == nil {
		return nil, false
	}
	return o.Webhooks, true
}

// SetWebhooks sets field value
func (o *WebhookListResponse) SetWebhooks(v []Webhook) {
	o.Webhooks = v
}

func (o WebhookListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["webhooks"] = o.Webhooks
	return toSerialize, nil
}

func (o *WebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"webhooks",
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

	varWebhookListResponse := _WebhookListResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookListResponse)

	if err != nil {
		return err
	}

	*o = WebhookListResponse(varWebhookListResponse)

	return err
}

type NullableWebhookListResponse struct {
	value *WebhookListResponse
	isSet bool
}

func (v NullableWebhookListResponse) Get() *WebhookListResponse {
	return v.value
}

func (v *NullableWebhookListResponse) Set(val *WebhookListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookListResponse(val *WebhookListResponse) *NullableWebhookListResponse {
	return &NullableWebhookListResponse{value: val, isSet: true}
}

func (v NullableWebhookListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

