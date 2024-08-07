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

// checks if the WebhookPostReqBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookPostReqBody{}

// WebhookPostReqBody struct for WebhookPostReqBody
type WebhookPostReqBody struct {
	Name         string                      `json:"name"`
	Url          string                      `json:"url"`
	Subscription *WebhookSubscriptionFilters `json:"subscription,omitempty"`
}

type _WebhookPostReqBody WebhookPostReqBody

// NewWebhookPostReqBody instantiates a new WebhookPostReqBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookPostReqBody(name string, url string) *WebhookPostReqBody {
	this := WebhookPostReqBody{}
	this.Name = name
	this.Url = url
	return &this
}

// NewWebhookPostReqBodyWithDefaults instantiates a new WebhookPostReqBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookPostReqBodyWithDefaults() *WebhookPostReqBody {
	this := WebhookPostReqBody{}
	return &this
}

// GetName returns the Name field value
func (o *WebhookPostReqBody) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookPostReqBody) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookPostReqBody) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *WebhookPostReqBody) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WebhookPostReqBody) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WebhookPostReqBody) SetUrl(v string) {
	o.Url = v
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
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookPostReqBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	return toSerialize, nil
}

func (o *WebhookPostReqBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"url",
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

	varWebhookPostReqBody := _WebhookPostReqBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookPostReqBody)

	if err != nil {
		return err
	}

	*o = WebhookPostReqBody(varWebhookPostReqBody)

	return err
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
