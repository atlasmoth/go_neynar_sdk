/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the SignerEventBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignerEventBody{}

// SignerEventBody struct for SignerEventBody
type SignerEventBody struct {
	Key string `json:"key" validate:"regexp=^0x[a-fA-F0-9]{64}$"`
	KeyType int64 `json:"keyType"`
	EventType SignerEventType `json:"eventType"`
	Metadata string `json:"metadata" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	MetadataType int64 `json:"metadataType"`
	AdditionalProperties map[string]interface{}
}

type _SignerEventBody SignerEventBody

// NewSignerEventBody instantiates a new SignerEventBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignerEventBody(key string, keyType int64, eventType SignerEventType, metadata string, metadataType int64) *SignerEventBody {
	this := SignerEventBody{}
	this.Key = key
	this.KeyType = keyType
	this.EventType = eventType
	this.Metadata = metadata
	this.MetadataType = metadataType
	return &this
}

// NewSignerEventBodyWithDefaults instantiates a new SignerEventBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignerEventBodyWithDefaults() *SignerEventBody {
	this := SignerEventBody{}
	var eventType SignerEventType = SIGNEREVENTTYPE_ADD
	this.EventType = eventType
	return &this
}

// GetKey returns the Key field value
func (o *SignerEventBody) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *SignerEventBody) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *SignerEventBody) SetKey(v string) {
	o.Key = v
}

// GetKeyType returns the KeyType field value
func (o *SignerEventBody) GetKeyType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value
// and a boolean to check if the value has been set.
func (o *SignerEventBody) GetKeyTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyType, true
}

// SetKeyType sets field value
func (o *SignerEventBody) SetKeyType(v int64) {
	o.KeyType = v
}

// GetEventType returns the EventType field value
func (o *SignerEventBody) GetEventType() SignerEventType {
	if o == nil {
		var ret SignerEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SignerEventBody) GetEventTypeOk() (*SignerEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SignerEventBody) SetEventType(v SignerEventType) {
	o.EventType = v
}

// GetMetadata returns the Metadata field value
func (o *SignerEventBody) GetMetadata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SignerEventBody) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SignerEventBody) SetMetadata(v string) {
	o.Metadata = v
}

// GetMetadataType returns the MetadataType field value
func (o *SignerEventBody) GetMetadataType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MetadataType
}

// GetMetadataTypeOk returns a tuple with the MetadataType field value
// and a boolean to check if the value has been set.
func (o *SignerEventBody) GetMetadataTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataType, true
}

// SetMetadataType sets field value
func (o *SignerEventBody) SetMetadataType(v int64) {
	o.MetadataType = v
}

func (o SignerEventBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignerEventBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["keyType"] = o.KeyType
	toSerialize["eventType"] = o.EventType
	toSerialize["metadata"] = o.Metadata
	toSerialize["metadataType"] = o.MetadataType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SignerEventBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"keyType",
		"eventType",
		"metadata",
		"metadataType",
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

	varSignerEventBody := _SignerEventBody{}

	err = json.Unmarshal(data, &varSignerEventBody)

	if err != nil {
		return err
	}

	*o = SignerEventBody(varSignerEventBody)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "keyType")
		delete(additionalProperties, "eventType")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "metadataType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSignerEventBody struct {
	value *SignerEventBody
	isSet bool
}

func (v NullableSignerEventBody) Get() *SignerEventBody {
	return v.value
}

func (v *NullableSignerEventBody) Set(val *SignerEventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerEventBody(val *SignerEventBody) *NullableSignerEventBody {
	return &NullableSignerEventBody{value: val, isSet: true}
}

func (v NullableSignerEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


