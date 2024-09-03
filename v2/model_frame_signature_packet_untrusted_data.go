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

// checks if the FrameSignaturePacketUntrustedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameSignaturePacketUntrustedData{}

// FrameSignaturePacketUntrustedData Untrusted data from the user
type FrameSignaturePacketUntrustedData struct {
	// User identifier (unsigned integer)
	Fid *int32 `json:"fid,omitempty"`
	// URL of the frame
	Url *string `json:"url,omitempty"`
	// Message hash
	MessageHash *string `json:"messageHash,omitempty"`
	// Timestamp
	Timestamp *string `json:"timestamp,omitempty"`
	// Network
	Network *int32 `json:"network,omitempty"`
	// Index of the button
	ButtonIndex *int32 `json:"buttonIndex,omitempty"`
	// Input text
	InputText *string `json:"inputText,omitempty"`
	// State
	State *string `json:"state,omitempty"`
	// Transaction ID
	TransactionId *string `json:"transactionId,omitempty"`
	// Ethereum address
	Address *string `json:"address,omitempty" validate:"regexp=^0x[a-fA-F0-9]{40}$"`
	CastId *CastId `json:"castId,omitempty"`
}

// NewFrameSignaturePacketUntrustedData instantiates a new FrameSignaturePacketUntrustedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameSignaturePacketUntrustedData() *FrameSignaturePacketUntrustedData {
	this := FrameSignaturePacketUntrustedData{}
	return &this
}

// NewFrameSignaturePacketUntrustedDataWithDefaults instantiates a new FrameSignaturePacketUntrustedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameSignaturePacketUntrustedDataWithDefaults() *FrameSignaturePacketUntrustedData {
	this := FrameSignaturePacketUntrustedData{}
	return &this
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetFid() int32 {
	if o == nil || IsNil(o.Fid) {
		var ret int32
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetFidOk() (*int32, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given int32 and assigns it to the Fid field.
func (o *FrameSignaturePacketUntrustedData) SetFid(v int32) {
	o.Fid = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FrameSignaturePacketUntrustedData) SetUrl(v string) {
	o.Url = &v
}

// GetMessageHash returns the MessageHash field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetMessageHash() string {
	if o == nil || IsNil(o.MessageHash) {
		var ret string
		return ret
	}
	return *o.MessageHash
}

// GetMessageHashOk returns a tuple with the MessageHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetMessageHashOk() (*string, bool) {
	if o == nil || IsNil(o.MessageHash) {
		return nil, false
	}
	return o.MessageHash, true
}

// HasMessageHash returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasMessageHash() bool {
	if o != nil && !IsNil(o.MessageHash) {
		return true
	}

	return false
}

// SetMessageHash gets a reference to the given string and assigns it to the MessageHash field.
func (o *FrameSignaturePacketUntrustedData) SetMessageHash(v string) {
	o.MessageHash = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *FrameSignaturePacketUntrustedData) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetNetwork() int32 {
	if o == nil || IsNil(o.Network) {
		var ret int32
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetNetworkOk() (*int32, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given int32 and assigns it to the Network field.
func (o *FrameSignaturePacketUntrustedData) SetNetwork(v int32) {
	o.Network = &v
}

// GetButtonIndex returns the ButtonIndex field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetButtonIndex() int32 {
	if o == nil || IsNil(o.ButtonIndex) {
		var ret int32
		return ret
	}
	return *o.ButtonIndex
}

// GetButtonIndexOk returns a tuple with the ButtonIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetButtonIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.ButtonIndex) {
		return nil, false
	}
	return o.ButtonIndex, true
}

// HasButtonIndex returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasButtonIndex() bool {
	if o != nil && !IsNil(o.ButtonIndex) {
		return true
	}

	return false
}

// SetButtonIndex gets a reference to the given int32 and assigns it to the ButtonIndex field.
func (o *FrameSignaturePacketUntrustedData) SetButtonIndex(v int32) {
	o.ButtonIndex = &v
}

// GetInputText returns the InputText field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetInputText() string {
	if o == nil || IsNil(o.InputText) {
		var ret string
		return ret
	}
	return *o.InputText
}

// GetInputTextOk returns a tuple with the InputText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetInputTextOk() (*string, bool) {
	if o == nil || IsNil(o.InputText) {
		return nil, false
	}
	return o.InputText, true
}

// HasInputText returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasInputText() bool {
	if o != nil && !IsNil(o.InputText) {
		return true
	}

	return false
}

// SetInputText gets a reference to the given string and assigns it to the InputText field.
func (o *FrameSignaturePacketUntrustedData) SetInputText(v string) {
	o.InputText = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *FrameSignaturePacketUntrustedData) SetState(v string) {
	o.State = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *FrameSignaturePacketUntrustedData) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *FrameSignaturePacketUntrustedData) SetAddress(v string) {
	o.Address = &v
}

// GetCastId returns the CastId field value if set, zero value otherwise.
func (o *FrameSignaturePacketUntrustedData) GetCastId() CastId {
	if o == nil || IsNil(o.CastId) {
		var ret CastId
		return ret
	}
	return *o.CastId
}

// GetCastIdOk returns a tuple with the CastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameSignaturePacketUntrustedData) GetCastIdOk() (*CastId, bool) {
	if o == nil || IsNil(o.CastId) {
		return nil, false
	}
	return o.CastId, true
}

// HasCastId returns a boolean if a field has been set.
func (o *FrameSignaturePacketUntrustedData) HasCastId() bool {
	if o != nil && !IsNil(o.CastId) {
		return true
	}

	return false
}

// SetCastId gets a reference to the given CastId and assigns it to the CastId field.
func (o *FrameSignaturePacketUntrustedData) SetCastId(v CastId) {
	o.CastId = &v
}

func (o FrameSignaturePacketUntrustedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameSignaturePacketUntrustedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.MessageHash) {
		toSerialize["messageHash"] = o.MessageHash
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.ButtonIndex) {
		toSerialize["buttonIndex"] = o.ButtonIndex
	}
	if !IsNil(o.InputText) {
		toSerialize["inputText"] = o.InputText
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.CastId) {
		toSerialize["castId"] = o.CastId
	}
	return toSerialize, nil
}

type NullableFrameSignaturePacketUntrustedData struct {
	value *FrameSignaturePacketUntrustedData
	isSet bool
}

func (v NullableFrameSignaturePacketUntrustedData) Get() *FrameSignaturePacketUntrustedData {
	return v.value
}

func (v *NullableFrameSignaturePacketUntrustedData) Set(val *FrameSignaturePacketUntrustedData) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameSignaturePacketUntrustedData) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameSignaturePacketUntrustedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameSignaturePacketUntrustedData(val *FrameSignaturePacketUntrustedData) *NullableFrameSignaturePacketUntrustedData {
	return &NullableFrameSignaturePacketUntrustedData{value: val, isSet: true}
}

func (v NullableFrameSignaturePacketUntrustedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameSignaturePacketUntrustedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


