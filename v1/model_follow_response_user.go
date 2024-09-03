/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the FollowResponseUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FollowResponseUser{}

// FollowResponseUser struct for FollowResponseUser
type FollowResponseUser struct {
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	// The username of the user.
	Username string `json:"username"`
	// Custody Address of the user.
	CustodyAddress string `json:"custodyAddress"`
	// The display of the reactor.
	DisplayName string `json:"displayName"`
	Pfp UserPfp `json:"pfp"`
	Profile UserProfile `json:"profile"`
	// The number of followers the user has.
	FollowerCount int32 `json:"followerCount"`
	// The number of users the user is following.
	FollowingCount int32 `json:"followingCount"`
	Verifications []string `json:"verifications"`
	ActiveStatus ActiveStatus `json:"activeStatus"`
	ViewerContext *ViewerContext `json:"viewerContext,omitempty"`
	Timestamp time.Time `json:"timestamp"`
	AdditionalProperties map[string]interface{}
}

type _FollowResponseUser FollowResponseUser

// NewFollowResponseUser instantiates a new FollowResponseUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFollowResponseUser(fid int32, username string, custodyAddress string, displayName string, pfp UserPfp, profile UserProfile, followerCount int32, followingCount int32, verifications []string, activeStatus ActiveStatus, timestamp time.Time) *FollowResponseUser {
	this := FollowResponseUser{}
	this.Fid = fid
	this.Username = username
	this.CustodyAddress = custodyAddress
	this.DisplayName = displayName
	this.Pfp = pfp
	this.Profile = profile
	this.FollowerCount = followerCount
	this.FollowingCount = followingCount
	this.Verifications = verifications
	this.ActiveStatus = activeStatus
	this.Timestamp = timestamp
	return &this
}

// NewFollowResponseUserWithDefaults instantiates a new FollowResponseUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFollowResponseUserWithDefaults() *FollowResponseUser {
	this := FollowResponseUser{}
	var fid int32 = 3
	this.Fid = fid
	return &this
}

// GetFid returns the Fid field value
func (o *FollowResponseUser) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *FollowResponseUser) SetFid(v int32) {
	o.Fid = v
}

// GetUsername returns the Username field value
func (o *FollowResponseUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *FollowResponseUser) SetUsername(v string) {
	o.Username = v
}

// GetCustodyAddress returns the CustodyAddress field value
func (o *FollowResponseUser) GetCustodyAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustodyAddress
}

// GetCustodyAddressOk returns a tuple with the CustodyAddress field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetCustodyAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustodyAddress, true
}

// SetCustodyAddress sets field value
func (o *FollowResponseUser) SetCustodyAddress(v string) {
	o.CustodyAddress = v
}

// GetDisplayName returns the DisplayName field value
func (o *FollowResponseUser) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *FollowResponseUser) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetPfp returns the Pfp field value
func (o *FollowResponseUser) GetPfp() UserPfp {
	if o == nil {
		var ret UserPfp
		return ret
	}

	return o.Pfp
}

// GetPfpOk returns a tuple with the Pfp field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetPfpOk() (*UserPfp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pfp, true
}

// SetPfp sets field value
func (o *FollowResponseUser) SetPfp(v UserPfp) {
	o.Pfp = v
}

// GetProfile returns the Profile field value
func (o *FollowResponseUser) GetProfile() UserProfile {
	if o == nil {
		var ret UserProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetProfileOk() (*UserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *FollowResponseUser) SetProfile(v UserProfile) {
	o.Profile = v
}

// GetFollowerCount returns the FollowerCount field value
func (o *FollowResponseUser) GetFollowerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetFollowerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowerCount, true
}

// SetFollowerCount sets field value
func (o *FollowResponseUser) SetFollowerCount(v int32) {
	o.FollowerCount = v
}

// GetFollowingCount returns the FollowingCount field value
func (o *FollowResponseUser) GetFollowingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetFollowingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingCount, true
}

// SetFollowingCount sets field value
func (o *FollowResponseUser) SetFollowingCount(v int32) {
	o.FollowingCount = v
}

// GetVerifications returns the Verifications field value
func (o *FollowResponseUser) GetVerifications() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetVerificationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *FollowResponseUser) SetVerifications(v []string) {
	o.Verifications = v
}

// GetActiveStatus returns the ActiveStatus field value
func (o *FollowResponseUser) GetActiveStatus() ActiveStatus {
	if o == nil {
		var ret ActiveStatus
		return ret
	}

	return o.ActiveStatus
}

// GetActiveStatusOk returns a tuple with the ActiveStatus field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetActiveStatusOk() (*ActiveStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveStatus, true
}

// SetActiveStatus sets field value
func (o *FollowResponseUser) SetActiveStatus(v ActiveStatus) {
	o.ActiveStatus = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *FollowResponseUser) GetViewerContext() ViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret ViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetViewerContextOk() (*ViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *FollowResponseUser) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given ViewerContext and assigns it to the ViewerContext field.
func (o *FollowResponseUser) SetViewerContext(v ViewerContext) {
	o.ViewerContext = &v
}

// GetTimestamp returns the Timestamp field value
func (o *FollowResponseUser) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *FollowResponseUser) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *FollowResponseUser) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o FollowResponseUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FollowResponseUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	toSerialize["custodyAddress"] = o.CustodyAddress
	toSerialize["displayName"] = o.DisplayName
	toSerialize["pfp"] = o.Pfp
	toSerialize["profile"] = o.Profile
	toSerialize["followerCount"] = o.FollowerCount
	toSerialize["followingCount"] = o.FollowingCount
	toSerialize["verifications"] = o.Verifications
	toSerialize["activeStatus"] = o.ActiveStatus
	if !IsNil(o.ViewerContext) {
		toSerialize["viewerContext"] = o.ViewerContext
	}
	toSerialize["timestamp"] = o.Timestamp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FollowResponseUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fid",
		"username",
		"custodyAddress",
		"displayName",
		"pfp",
		"profile",
		"followerCount",
		"followingCount",
		"verifications",
		"activeStatus",
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

	varFollowResponseUser := _FollowResponseUser{}

	err = json.Unmarshal(data, &varFollowResponseUser)

	if err != nil {
		return err
	}

	*o = FollowResponseUser(varFollowResponseUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "fid")
		delete(additionalProperties, "username")
		delete(additionalProperties, "custodyAddress")
		delete(additionalProperties, "displayName")
		delete(additionalProperties, "pfp")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "followerCount")
		delete(additionalProperties, "followingCount")
		delete(additionalProperties, "verifications")
		delete(additionalProperties, "activeStatus")
		delete(additionalProperties, "viewerContext")
		delete(additionalProperties, "timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFollowResponseUser struct {
	value *FollowResponseUser
	isSet bool
}

func (v NullableFollowResponseUser) Get() *FollowResponseUser {
	return v.value
}

func (v *NullableFollowResponseUser) Set(val *FollowResponseUser) {
	v.value = val
	v.isSet = true
}

func (v NullableFollowResponseUser) IsSet() bool {
	return v.isSet
}

func (v *NullableFollowResponseUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFollowResponseUser(val *FollowResponseUser) *NullableFollowResponseUser {
	return &NullableFollowResponseUser{value: val, isSet: true}
}

func (v NullableFollowResponseUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFollowResponseUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


