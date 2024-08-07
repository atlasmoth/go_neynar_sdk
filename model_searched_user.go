/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SearchedUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchedUser{}

// SearchedUser struct for SearchedUser
type SearchedUser struct {
	Object string `json:"object"`
	// User identifier (unsigned integer)
	Fid int32 `json:"fid"`
	Username string `json:"username"`
	DisplayName *string `json:"display_name,omitempty"`
	// Ethereum address
	CustodyAddress string `json:"custody_address" validate:"regexp=^0x[a-fA-F0-9]{40}$"`
	// The URL of the user's profile picture
	PfpUrl *string `json:"pfp_url,omitempty"`
	Profile UserProfile `json:"profile"`
	// The number of followers the user has.
	FollowerCount int32 `json:"follower_count"`
	// The number of users the user is following.
	FollowingCount int32 `json:"following_count"`
	Verifications []string `json:"verifications"`
	VerifiedAddresses UserVerifiedAddresses `json:"verified_addresses"`
	ActiveStatus ActiveStatus `json:"active_status"`
	PowerBadge bool `json:"power_badge"`
	ViewerContext *UserViewerContext `json:"viewer_context,omitempty"`
	Pfp ProfileUrlPfp `json:"pfp"`
}

type _SearchedUser SearchedUser

// NewSearchedUser instantiates a new SearchedUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchedUser(object string, fid int32, username string, custodyAddress string, profile UserProfile, followerCount int32, followingCount int32, verifications []string, verifiedAddresses UserVerifiedAddresses, activeStatus ActiveStatus, powerBadge bool, pfp ProfileUrlPfp) *SearchedUser {
	this := SearchedUser{}
	this.Object = object
	this.Fid = fid
	this.Username = username
	this.CustodyAddress = custodyAddress
	this.Profile = profile
	this.FollowerCount = followerCount
	this.FollowingCount = followingCount
	this.Verifications = verifications
	this.VerifiedAddresses = verifiedAddresses
	this.ActiveStatus = activeStatus
	this.PowerBadge = powerBadge
	this.Pfp = pfp
	return &this
}

// NewSearchedUserWithDefaults instantiates a new SearchedUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchedUserWithDefaults() *SearchedUser {
	this := SearchedUser{}
	return &this
}

// GetObject returns the Object field value
func (o *SearchedUser) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *SearchedUser) SetObject(v string) {
	o.Object = v
}

// GetFid returns the Fid field value
func (o *SearchedUser) GetFid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fid
}

// GetFidOk returns a tuple with the Fid field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetFidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fid, true
}

// SetFid sets field value
func (o *SearchedUser) SetFid(v int32) {
	o.Fid = v
}

// GetUsername returns the Username field value
func (o *SearchedUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *SearchedUser) SetUsername(v string) {
	o.Username = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SearchedUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SearchedUser) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SearchedUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCustodyAddress returns the CustodyAddress field value
func (o *SearchedUser) GetCustodyAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustodyAddress
}

// GetCustodyAddressOk returns a tuple with the CustodyAddress field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetCustodyAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustodyAddress, true
}

// SetCustodyAddress sets field value
func (o *SearchedUser) SetCustodyAddress(v string) {
	o.CustodyAddress = v
}

// GetPfpUrl returns the PfpUrl field value if set, zero value otherwise.
func (o *SearchedUser) GetPfpUrl() string {
	if o == nil || IsNil(o.PfpUrl) {
		var ret string
		return ret
	}
	return *o.PfpUrl
}

// GetPfpUrlOk returns a tuple with the PfpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetPfpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PfpUrl) {
		return nil, false
	}
	return o.PfpUrl, true
}

// HasPfpUrl returns a boolean if a field has been set.
func (o *SearchedUser) HasPfpUrl() bool {
	if o != nil && !IsNil(o.PfpUrl) {
		return true
	}

	return false
}

// SetPfpUrl gets a reference to the given string and assigns it to the PfpUrl field.
func (o *SearchedUser) SetPfpUrl(v string) {
	o.PfpUrl = &v
}

// GetProfile returns the Profile field value
func (o *SearchedUser) GetProfile() UserProfile {
	if o == nil {
		var ret UserProfile
		return ret
	}

	return o.Profile
}

// GetProfileOk returns a tuple with the Profile field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetProfileOk() (*UserProfile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Profile, true
}

// SetProfile sets field value
func (o *SearchedUser) SetProfile(v UserProfile) {
	o.Profile = v
}

// GetFollowerCount returns the FollowerCount field value
func (o *SearchedUser) GetFollowerCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetFollowerCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowerCount, true
}

// SetFollowerCount sets field value
func (o *SearchedUser) SetFollowerCount(v int32) {
	o.FollowerCount = v
}

// GetFollowingCount returns the FollowingCount field value
func (o *SearchedUser) GetFollowingCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetFollowingCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingCount, true
}

// SetFollowingCount sets field value
func (o *SearchedUser) SetFollowingCount(v int32) {
	o.FollowingCount = v
}

// GetVerifications returns the Verifications field value
func (o *SearchedUser) GetVerifications() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetVerificationsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Verifications, true
}

// SetVerifications sets field value
func (o *SearchedUser) SetVerifications(v []string) {
	o.Verifications = v
}

// GetVerifiedAddresses returns the VerifiedAddresses field value
func (o *SearchedUser) GetVerifiedAddresses() UserVerifiedAddresses {
	if o == nil {
		var ret UserVerifiedAddresses
		return ret
	}

	return o.VerifiedAddresses
}

// GetVerifiedAddressesOk returns a tuple with the VerifiedAddresses field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetVerifiedAddressesOk() (*UserVerifiedAddresses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifiedAddresses, true
}

// SetVerifiedAddresses sets field value
func (o *SearchedUser) SetVerifiedAddresses(v UserVerifiedAddresses) {
	o.VerifiedAddresses = v
}

// GetActiveStatus returns the ActiveStatus field value
func (o *SearchedUser) GetActiveStatus() ActiveStatus {
	if o == nil {
		var ret ActiveStatus
		return ret
	}

	return o.ActiveStatus
}

// GetActiveStatusOk returns a tuple with the ActiveStatus field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetActiveStatusOk() (*ActiveStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveStatus, true
}

// SetActiveStatus sets field value
func (o *SearchedUser) SetActiveStatus(v ActiveStatus) {
	o.ActiveStatus = v
}

// GetPowerBadge returns the PowerBadge field value
func (o *SearchedUser) GetPowerBadge() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PowerBadge
}

// GetPowerBadgeOk returns a tuple with the PowerBadge field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetPowerBadgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerBadge, true
}

// SetPowerBadge sets field value
func (o *SearchedUser) SetPowerBadge(v bool) {
	o.PowerBadge = v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *SearchedUser) GetViewerContext() UserViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret UserViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetViewerContextOk() (*UserViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *SearchedUser) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given UserViewerContext and assigns it to the ViewerContext field.
func (o *SearchedUser) SetViewerContext(v UserViewerContext) {
	o.ViewerContext = &v
}

// GetPfp returns the Pfp field value
func (o *SearchedUser) GetPfp() ProfileUrlPfp {
	if o == nil {
		var ret ProfileUrlPfp
		return ret
	}

	return o.Pfp
}

// GetPfpOk returns a tuple with the Pfp field value
// and a boolean to check if the value has been set.
func (o *SearchedUser) GetPfpOk() (*ProfileUrlPfp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pfp, true
}

// SetPfp sets field value
func (o *SearchedUser) SetPfp(v ProfileUrlPfp) {
	o.Pfp = v
}

func (o SearchedUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchedUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["fid"] = o.Fid
	toSerialize["username"] = o.Username
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	toSerialize["custody_address"] = o.CustodyAddress
	if !IsNil(o.PfpUrl) {
		toSerialize["pfp_url"] = o.PfpUrl
	}
	toSerialize["profile"] = o.Profile
	toSerialize["follower_count"] = o.FollowerCount
	toSerialize["following_count"] = o.FollowingCount
	toSerialize["verifications"] = o.Verifications
	toSerialize["verified_addresses"] = o.VerifiedAddresses
	toSerialize["active_status"] = o.ActiveStatus
	toSerialize["power_badge"] = o.PowerBadge
	if !IsNil(o.ViewerContext) {
		toSerialize["viewer_context"] = o.ViewerContext
	}
	toSerialize["pfp"] = o.Pfp
	return toSerialize, nil
}

func (o *SearchedUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"object",
		"fid",
		"username",
		"custody_address",
		"profile",
		"follower_count",
		"following_count",
		"verifications",
		"verified_addresses",
		"active_status",
		"power_badge",
		"pfp",
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

	varSearchedUser := _SearchedUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchedUser)

	if err != nil {
		return err
	}

	*o = SearchedUser(varSearchedUser)

	return err
}

type NullableSearchedUser struct {
	value *SearchedUser
	isSet bool
}

func (v NullableSearchedUser) Get() *SearchedUser {
	return v.value
}

func (v *NullableSearchedUser) Set(val *SearchedUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchedUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchedUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchedUser(val *SearchedUser) *NullableSearchedUser {
	return &NullableSearchedUser{value: val, isSet: true}
}

func (v NullableSearchedUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchedUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


