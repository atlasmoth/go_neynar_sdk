/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the ValidatedFrameAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidatedFrameAction{}

// ValidatedFrameAction struct for ValidatedFrameAction
type ValidatedFrameAction struct {
	Object string `json:"object"`
	Url string `json:"url"`
	Interactor User `json:"interactor"`
	TappedButton ValidatedFrameActionTappedButton `json:"tapped_button"`
	Input *FrameInput `json:"input,omitempty"`
	State FrameState `json:"state"`
	Cast CastWithInteractions `json:"cast"`
	Timestamp time.Time `json:"timestamp"`
	Signer *ValidatedFrameActionSigner `json:"signer,omitempty"`
	Transaction *FrameTransaction `json:"transaction,omitempty"`
	// The connected wallet address of the interacting user.
	Address *string `json:"address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidatedFrameAction ValidatedFrameAction

// NewValidatedFrameAction instantiates a new ValidatedFrameAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatedFrameAction(object string, url string, interactor User, tappedButton ValidatedFrameActionTappedButton, state FrameState, cast CastWithInteractions, timestamp time.Time) *ValidatedFrameAction {
	this := ValidatedFrameAction{}
	this.Object = object
	this.Url = url
	this.Interactor = interactor
	this.TappedButton = tappedButton
	this.State = state
	this.Cast = cast
	this.Timestamp = timestamp
	return &this
}

// NewValidatedFrameActionWithDefaults instantiates a new ValidatedFrameAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatedFrameActionWithDefaults() *ValidatedFrameAction {
	this := ValidatedFrameAction{}
	return &this
}

// GetObject returns the Object field value
func (o *ValidatedFrameAction) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ValidatedFrameAction) SetObject(v string) {
	o.Object = v
}

// GetUrl returns the Url field value
func (o *ValidatedFrameAction) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ValidatedFrameAction) SetUrl(v string) {
	o.Url = v
}

// GetInteractor returns the Interactor field value
func (o *ValidatedFrameAction) GetInteractor() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.Interactor
}

// GetInteractorOk returns a tuple with the Interactor field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetInteractorOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interactor, true
}

// SetInteractor sets field value
func (o *ValidatedFrameAction) SetInteractor(v User) {
	o.Interactor = v
}

// GetTappedButton returns the TappedButton field value
func (o *ValidatedFrameAction) GetTappedButton() ValidatedFrameActionTappedButton {
	if o == nil {
		var ret ValidatedFrameActionTappedButton
		return ret
	}

	return o.TappedButton
}

// GetTappedButtonOk returns a tuple with the TappedButton field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetTappedButtonOk() (*ValidatedFrameActionTappedButton, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TappedButton, true
}

// SetTappedButton sets field value
func (o *ValidatedFrameAction) SetTappedButton(v ValidatedFrameActionTappedButton) {
	o.TappedButton = v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *ValidatedFrameAction) GetInput() FrameInput {
	if o == nil || IsNil(o.Input) {
		var ret FrameInput
		return ret
	}
	return *o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetInputOk() (*FrameInput, bool) {
	if o == nil || IsNil(o.Input) {
		return nil, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *ValidatedFrameAction) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given FrameInput and assigns it to the Input field.
func (o *ValidatedFrameAction) SetInput(v FrameInput) {
	o.Input = &v
}

// GetState returns the State field value
func (o *ValidatedFrameAction) GetState() FrameState {
	if o == nil {
		var ret FrameState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetStateOk() (*FrameState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ValidatedFrameAction) SetState(v FrameState) {
	o.State = v
}

// GetCast returns the Cast field value
func (o *ValidatedFrameAction) GetCast() CastWithInteractions {
	if o == nil {
		var ret CastWithInteractions
		return ret
	}

	return o.Cast
}

// GetCastOk returns a tuple with the Cast field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetCastOk() (*CastWithInteractions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cast, true
}

// SetCast sets field value
func (o *ValidatedFrameAction) SetCast(v CastWithInteractions) {
	o.Cast = v
}

// GetTimestamp returns the Timestamp field value
func (o *ValidatedFrameAction) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ValidatedFrameAction) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *ValidatedFrameAction) GetSigner() ValidatedFrameActionSigner {
	if o == nil || IsNil(o.Signer) {
		var ret ValidatedFrameActionSigner
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetSignerOk() (*ValidatedFrameActionSigner, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *ValidatedFrameAction) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given ValidatedFrameActionSigner and assigns it to the Signer field.
func (o *ValidatedFrameAction) SetSigner(v ValidatedFrameActionSigner) {
	o.Signer = &v
}

// GetTransaction returns the Transaction field value if set, zero value otherwise.
func (o *ValidatedFrameAction) GetTransaction() FrameTransaction {
	if o == nil || IsNil(o.Transaction) {
		var ret FrameTransaction
		return ret
	}
	return *o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetTransactionOk() (*FrameTransaction, bool) {
	if o == nil || IsNil(o.Transaction) {
		return nil, false
	}
	return o.Transaction, true
}

// HasTransaction returns a boolean if a field has been set.
func (o *ValidatedFrameAction) HasTransaction() bool {
	if o != nil && !IsNil(o.Transaction) {
		return true
	}

	return false
}

// SetTransaction gets a reference to the given FrameTransaction and assigns it to the Transaction field.
func (o *ValidatedFrameAction) SetTransaction(v FrameTransaction) {
	o.Transaction = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ValidatedFrameAction) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedFrameAction) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ValidatedFrameAction) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ValidatedFrameAction) SetAddress(v string) {
	o.Address = &v
}

func (o ValidatedFrameAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidatedFrameAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["url"] = o.Url
	toSerialize["interactor"] = o.Interactor
	toSerialize["tapped_button"] = o.TappedButton
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	toSerialize["state"] = o.State
	toSerialize["cast"] = o.Cast
	toSerialize["timestamp"] = o.Timestamp
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.Transaction) {
		toSerialize["transaction"] = o.Transaction
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ValidatedFrameAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"url",
		"interactor",
		"tapped_button",
		"state",
		"cast",
		"timestamp",
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

	varValidatedFrameAction := _ValidatedFrameAction{}

	err = json.Unmarshal(data, &varValidatedFrameAction)

	if err != nil {
		return err
	}

	*o = ValidatedFrameAction(varValidatedFrameAction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "object")
		delete(additionalProperties, "url")
		delete(additionalProperties, "interactor")
		delete(additionalProperties, "tapped_button")
		delete(additionalProperties, "input")
		delete(additionalProperties, "state")
		delete(additionalProperties, "cast")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "signer")
		delete(additionalProperties, "transaction")
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableValidatedFrameAction struct {
	value *ValidatedFrameAction
	isSet bool
}

func (v NullableValidatedFrameAction) Get() *ValidatedFrameAction {
	return v.value
}

func (v *NullableValidatedFrameAction) Set(val *ValidatedFrameAction) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatedFrameAction) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatedFrameAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatedFrameAction(val *ValidatedFrameAction) *NullableValidatedFrameAction {
	return &NullableValidatedFrameAction{value: val, isSet: true}
}

func (v NullableValidatedFrameAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatedFrameAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


