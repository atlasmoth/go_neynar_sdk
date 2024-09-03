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

// checks if the WebhookResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResponse{}

// WebhookResponse struct for WebhookResponse
type WebhookResponse struct {
	Message *string `json:"message,omitempty"`
	Success *bool `json:"success,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

// NewWebhookResponse instantiates a new WebhookResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponse() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// NewWebhookResponseWithDefaults instantiates a new WebhookResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseWithDefaults() *WebhookResponse {
	this := WebhookResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *WebhookResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *WebhookResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *WebhookResponse) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *WebhookResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *WebhookResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *WebhookResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *WebhookResponse) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponse) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *WebhookResponse) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *WebhookResponse) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o WebhookResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

type NullableWebhookResponse struct {
	value *WebhookResponse
	isSet bool
}

func (v NullableWebhookResponse) Get() *WebhookResponse {
	return v.value
}

func (v *NullableWebhookResponse) Set(val *WebhookResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponse(val *WebhookResponse) *NullableWebhookResponse {
	return &NullableWebhookResponse{value: val, isSet: true}
}

func (v NullableWebhookResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


