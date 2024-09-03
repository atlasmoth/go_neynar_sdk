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

// checks if the StorageLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageLimit{}

// StorageLimit struct for StorageLimit
type StorageLimit struct {
	StoreType StoreType `json:"storeType"`
	Limit int32 `json:"limit"`
}

type _StorageLimit StorageLimit

// NewStorageLimit instantiates a new StorageLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageLimit(storeType StoreType, limit int32) *StorageLimit {
	this := StorageLimit{}
	this.StoreType = storeType
	this.Limit = limit
	return &this
}

// NewStorageLimitWithDefaults instantiates a new StorageLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageLimitWithDefaults() *StorageLimit {
	this := StorageLimit{}
	var storeType StoreType = STORETYPE_CASTS
	this.StoreType = storeType
	return &this
}

// GetStoreType returns the StoreType field value
func (o *StorageLimit) GetStoreType() StoreType {
	if o == nil {
		var ret StoreType
		return ret
	}

	return o.StoreType
}

// GetStoreTypeOk returns a tuple with the StoreType field value
// and a boolean to check if the value has been set.
func (o *StorageLimit) GetStoreTypeOk() (*StoreType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoreType, true
}

// SetStoreType sets field value
func (o *StorageLimit) SetStoreType(v StoreType) {
	o.StoreType = v
}

// GetLimit returns the Limit field value
func (o *StorageLimit) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *StorageLimit) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *StorageLimit) SetLimit(v int32) {
	o.Limit = v
}

func (o StorageLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["storeType"] = o.StoreType
	toSerialize["limit"] = o.Limit
	return toSerialize, nil
}

func (o *StorageLimit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"storeType",
		"limit",
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

	varStorageLimit := _StorageLimit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageLimit)

	if err != nil {
		return err
	}

	*o = StorageLimit(varStorageLimit)

	return err
}

type NullableStorageLimit struct {
	value *StorageLimit
	isSet bool
}

func (v NullableStorageLimit) Get() *StorageLimit {
	return v.value
}

func (v *NullableStorageLimit) Set(val *StorageLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageLimit(val *StorageLimit) *NullableStorageLimit {
	return &NullableStorageLimit{value: val, isSet: true}
}

func (v NullableStorageLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


