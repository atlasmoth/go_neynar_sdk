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

// checks if the WebhookPutReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookPutReqBody{}

// WebhookPutReqBody struct for WebhookPutReqBody
type WebhookPutReqBody struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
	Subscription *WebhookSubscriptionFilters `json:"subscription,omitempty"`
	WebhookId *string `json:"webhook_id,omitempty"`
}

// NewWebhookPutReqBody instantiates a new WebhookPutReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookPutReqBody() *WebhookPutReqBody {
	this := WebhookPutReqBody{}
	return &this
}

// NewWebhookPutReqBodyWithDefaults instantiates a new WebhookPutReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookPutReqBodyWithDefaults() *WebhookPutReqBody {
	this := WebhookPutReqBody{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WebhookPutReqBody) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPutReqBody) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WebhookPutReqBody) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WebhookPutReqBody) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *WebhookPutReqBody) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPutReqBody) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookPutReqBody) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *WebhookPutReqBody) SetUrl(v string) {
	o.Url = &v
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *WebhookPutReqBody) GetSubscription() WebhookSubscriptionFilters {
	if o == nil || IsNil(o.Subscription) {
		var ret WebhookSubscriptionFilters
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPutReqBody) GetSubscriptionOk() (*WebhookSubscriptionFilters, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *WebhookPutReqBody) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given WebhookSubscriptionFilters and assigns it to the Subscription field.
func (o *WebhookPutReqBody) SetSubscription(v WebhookSubscriptionFilters) {
	o.Subscription = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *WebhookPutReqBody) GetWebhookId() string {
	if o == nil || IsNil(o.WebhookId) {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookPutReqBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookId) {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *WebhookPutReqBody) HasWebhookId() bool {
	if o != nil && !IsNil(o.WebhookId) {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *WebhookPutReqBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

func (o WebhookPutReqBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookPutReqBody) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.WebhookId) {
		toSerialize["webhook_id"] = o.WebhookId
	}
	return toSerialize, nil
}

type NullableWebhookPutReqBody struct {
	value *WebhookPutReqBody
	isSet bool
}

func (v NullableWebhookPutReqBody) Get() *WebhookPutReqBody {
	return v.value
}

func (v *NullableWebhookPutReqBody) Set(val *WebhookPutReqBody) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookPutReqBody) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookPutReqBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookPutReqBody(val *WebhookPutReqBody) *NullableWebhookPutReqBody {
	return &NullableWebhookPutReqBody{value: val, isSet: true}
}

func (v NullableWebhookPutReqBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookPutReqBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


