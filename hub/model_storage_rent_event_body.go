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

// checks if the StorageRentEventBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageRentEventBody{}

// StorageRentEventBody struct for StorageRentEventBody
type StorageRentEventBody struct {
	Payer string `json:"payer" validate:"regexp=^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=)?$"`
	Units int64 `json:"units"`
	Expiry int64 `json:"expiry"`
}

type _StorageRentEventBody StorageRentEventBody

// NewStorageRentEventBody instantiates a new StorageRentEventBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageRentEventBody(payer string, units int64, expiry int64) *StorageRentEventBody {
	this := StorageRentEventBody{}
	this.Payer = payer
	this.Units = units
	this.Expiry = expiry
	return &this
}

// NewStorageRentEventBodyWithDefaults instantiates a new StorageRentEventBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageRentEventBodyWithDefaults() *StorageRentEventBody {
	this := StorageRentEventBody{}
	return &this
}

// GetPayer returns the Payer field value
func (o *StorageRentEventBody) GetPayer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payer
}

// GetPayerOk returns a tuple with the Payer field value
// and a boolean to check if the value has been set.
func (o *StorageRentEventBody) GetPayerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payer, true
}

// SetPayer sets field value
func (o *StorageRentEventBody) SetPayer(v string) {
	o.Payer = v
}

// GetUnits returns the Units field value
func (o *StorageRentEventBody) GetUnits() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *StorageRentEventBody) GetUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *StorageRentEventBody) SetUnits(v int64) {
	o.Units = v
}

// GetExpiry returns the Expiry field value
func (o *StorageRentEventBody) GetExpiry() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *StorageRentEventBody) GetExpiryOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *StorageRentEventBody) SetExpiry(v int64) {
	o.Expiry = v
}

func (o StorageRentEventBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageRentEventBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payer"] = o.Payer
	toSerialize["units"] = o.Units
	toSerialize["expiry"] = o.Expiry
	return toSerialize, nil
}

func (o *StorageRentEventBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"payer",
		"units",
		"expiry",
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

	varStorageRentEventBody := _StorageRentEventBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageRentEventBody)

	if err != nil {
		return err
	}

	*o = StorageRentEventBody(varStorageRentEventBody)

	return err
}

type NullableStorageRentEventBody struct {
	value *StorageRentEventBody
	isSet bool
}

func (v NullableStorageRentEventBody) Get() *StorageRentEventBody {
	return v.value
}

func (v *NullableStorageRentEventBody) Set(val *StorageRentEventBody) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageRentEventBody) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageRentEventBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageRentEventBody(val *StorageRentEventBody) *NullableStorageRentEventBody {
	return &NullableStorageRentEventBody{value: val, isSet: true}
}

func (v NullableStorageRentEventBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageRentEventBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

