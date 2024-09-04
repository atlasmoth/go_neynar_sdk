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

// checks if the WebhookPostReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookPostReqBody{}

// WebhookPostReqBody struct for WebhookPostReqBody
type WebhookPostReqBody struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Subscription *WebhookSubscriptionFilters `json:"subscription,omitempty"`
}

// NewWebhookPostReqBody instantiates a new WebhookPostReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookPostReqBody() *WebhookPostReqBody {
	this := WebhookPostReqBody{}
	return &this
}

// NewWebhookPostReqBodyWithDefaults instantiates a new WebhookPostReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookPostReqBodyWithDefaults() *WebhookPostReqBody {
	this := WebhookPostReqBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookPostReqBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPostReqBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookPostReqBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookPostReqBody) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhookPostReqBody) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPostReqBody) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookPostReqBody) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebhookPostReqBody) SetUrl(v string) {
	o.Url = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *WebhookPostReqBody) GetSubscription() WebhookSubscriptionFilters {
	if o == nil || IsNil(o.Subscription) {
		var ret WebhookSubscriptionFilters
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPostReqBody) GetSubscriptionOk() (*WebhookSubscriptionFilters, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *WebhookPostReqBody) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given WebhookSubscriptionFilters and assigns it to the Subscription field.
func (o *WebhookPostReqBody) SetSubscription(v WebhookSubscriptionFilters) {
	o.Subscription = &v
}

func (o WebhookPostReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookPostReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

type NullableWebhookPostReqBody struct {
	value *WebhookPostReqBody
	isSet bool
}

func (v NullableWebhookPostReqBody) Get() *WebhookPostReqBody {
	return v.value
}

func (v *NullableWebhookPostReqBody) Set(val *WebhookPostReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookPostReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookPostReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookPostReqBody(val *WebhookPostReqBody) *NullableWebhookPostReqBody {
	return &NullableWebhookPostReqBody{value: val, isSet: true}
}

func (v NullableWebhookPostReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookPostReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


