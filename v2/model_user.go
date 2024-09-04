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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Object *string `json:"object,omitempty"`
	// User identifier (unsigned integer)
	Fid *int32 `json:"fid,omitempty"`
	Username *string `json:"username,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	// Ethereum address
	CustodyAddress *string `json:"custody_address,omitempty" validate:"regexp=^0x[a-fA-F0-9]{40}$"`
	// The URL of the user's profile picture
	PfpUrl *string `json:"pfp_url,omitempty"`
	Profile *UserProfile `json:"profile,omitempty"`
	// The number of followers the user has.
	FollowerCount *int32 `json:"follower_count,omitempty"`
	// The number of users the user is following.
	FollowingCount *int32 `json:"following_count,omitempty"`
	Verifications []string `json:"verifications,omitempty"`
	VerifiedAddresses *UserVerifiedAddresses `json:"verified_addresses,omitempty"`
	ActiveStatus *ActiveStatus `json:"active_status,omitempty"`
	PowerBadge *bool `json:"power_badge,omitempty"`
	ViewerContext *UserViewerContext `json:"viewer_context,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *User) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *User) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *User) SetObject(v string) {
	o.Object = &v
}

// GetFid returns the Fid field value if set, zero value otherwise.
func (o *User) GetFid() int32 {
	if o == nil || IsNil(o.Fid) {
		var ret int32
		return ret
	}
	return *o.Fid
}

// GetFidOk returns a tuple with the Fid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFidOk() (*int32, bool) {
	if o == nil || IsNil(o.Fid) {
		return nil, false
	}
	return o.Fid, true
}

// HasFid returns a boolean if a field has been set.
func (o *User) HasFid() bool {
	if o != nil && !IsNil(o.Fid) {
		return true
	}

	return false
}

// SetFid gets a reference to the given int32 and assigns it to the Fid field.
func (o *User) SetFid(v int32) {
	o.Fid = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *User) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *User) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *User) SetUsername(v string) {
	o.Username = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *User) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *User) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *User) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCustodyAddress returns the CustodyAddress field value if set, zero value otherwise.
func (o *User) GetCustodyAddress() string {
	if o == nil || IsNil(o.CustodyAddress) {
		var ret string
		return ret
	}
	return *o.CustodyAddress
}

// GetCustodyAddressOk returns a tuple with the CustodyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCustodyAddressOk() (*string, bool) {
	if o == nil || IsNil(o.CustodyAddress) {
		return nil, false
	}
	return o.CustodyAddress, true
}

// HasCustodyAddress returns a boolean if a field has been set.
func (o *User) HasCustodyAddress() bool {
	if o != nil && !IsNil(o.CustodyAddress) {
		return true
	}

	return false
}

// SetCustodyAddress gets a reference to the given string and assigns it to the CustodyAddress field.
func (o *User) SetCustodyAddress(v string) {
	o.CustodyAddress = &v
}

// GetPfpUrl returns the PfpUrl field value if set, zero value otherwise.
func (o *User) GetPfpUrl() string {
	if o == nil || IsNil(o.PfpUrl) {
		var ret string
		return ret
	}
	return *o.PfpUrl
}

// GetPfpUrlOk returns a tuple with the PfpUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPfpUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PfpUrl) {
		return nil, false
	}
	return o.PfpUrl, true
}

// HasPfpUrl returns a boolean if a field has been set.
func (o *User) HasPfpUrl() bool {
	if o != nil && !IsNil(o.PfpUrl) {
		return true
	}

	return false
}

// SetPfpUrl gets a reference to the given string and assigns it to the PfpUrl field.
func (o *User) SetPfpUrl(v string) {
	o.PfpUrl = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *User) GetProfile() UserProfile {
	if o == nil || IsNil(o.Profile) {
		var ret UserProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProfileOk() (*UserProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *User) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given UserProfile and assigns it to the Profile field.
func (o *User) SetProfile(v UserProfile) {
	o.Profile = &v
}

// GetFollowerCount returns the FollowerCount field value if set, zero value otherwise.
func (o *User) GetFollowerCount() int32 {
	if o == nil || IsNil(o.FollowerCount) {
		var ret int32
		return ret
	}
	return *o.FollowerCount
}

// GetFollowerCountOk returns a tuple with the FollowerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFollowerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FollowerCount) {
		return nil, false
	}
	return o.FollowerCount, true
}

// HasFollowerCount returns a boolean if a field has been set.
func (o *User) HasFollowerCount() bool {
	if o != nil && !IsNil(o.FollowerCount) {
		return true
	}

	return false
}

// SetFollowerCount gets a reference to the given int32 and assigns it to the FollowerCount field.
func (o *User) SetFollowerCount(v int32) {
	o.FollowerCount = &v
}

// GetFollowingCount returns the FollowingCount field value if set, zero value otherwise.
func (o *User) GetFollowingCount() int32 {
	if o == nil || IsNil(o.FollowingCount) {
		var ret int32
		return ret
	}
	return *o.FollowingCount
}

// GetFollowingCountOk returns a tuple with the FollowingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFollowingCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FollowingCount) {
		return nil, false
	}
	return o.FollowingCount, true
}

// HasFollowingCount returns a boolean if a field has been set.
func (o *User) HasFollowingCount() bool {
	if o != nil && !IsNil(o.FollowingCount) {
		return true
	}

	return false
}

// SetFollowingCount gets a reference to the given int32 and assigns it to the FollowingCount field.
func (o *User) SetFollowingCount(v int32) {
	o.FollowingCount = &v
}

// GetVerifications returns the Verifications field value if set, zero value otherwise.
func (o *User) GetVerifications() []string {
	if o == nil || IsNil(o.Verifications) {
		var ret []string
		return ret
	}
	return o.Verifications
}

// GetVerificationsOk returns a tuple with the Verifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerificationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Verifications) {
		return nil, false
	}
	return o.Verifications, true
}

// HasVerifications returns a boolean if a field has been set.
func (o *User) HasVerifications() bool {
	if o != nil && !IsNil(o.Verifications) {
		return true
	}

	return false
}

// SetVerifications gets a reference to the given []string and assigns it to the Verifications field.
func (o *User) SetVerifications(v []string) {
	o.Verifications = v
}

// GetVerifiedAddresses returns the VerifiedAddresses field value if set, zero value otherwise.
func (o *User) GetVerifiedAddresses() UserVerifiedAddresses {
	if o == nil || IsNil(o.VerifiedAddresses) {
		var ret UserVerifiedAddresses
		return ret
	}
	return *o.VerifiedAddresses
}

// GetVerifiedAddressesOk returns a tuple with the VerifiedAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetVerifiedAddressesOk() (*UserVerifiedAddresses, bool) {
	if o == nil || IsNil(o.VerifiedAddresses) {
		return nil, false
	}
	return o.VerifiedAddresses, true
}

// HasVerifiedAddresses returns a boolean if a field has been set.
func (o *User) HasVerifiedAddresses() bool {
	if o != nil && !IsNil(o.VerifiedAddresses) {
		return true
	}

	return false
}

// SetVerifiedAddresses gets a reference to the given UserVerifiedAddresses and assigns it to the VerifiedAddresses field.
func (o *User) SetVerifiedAddresses(v UserVerifiedAddresses) {
	o.VerifiedAddresses = &v
}

// GetActiveStatus returns the ActiveStatus field value if set, zero value otherwise.
func (o *User) GetActiveStatus() ActiveStatus {
	if o == nil || IsNil(o.ActiveStatus) {
		var ret ActiveStatus
		return ret
	}
	return *o.ActiveStatus
}

// GetActiveStatusOk returns a tuple with the ActiveStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetActiveStatusOk() (*ActiveStatus, bool) {
	if o == nil || IsNil(o.ActiveStatus) {
		return nil, false
	}
	return o.ActiveStatus, true
}

// HasActiveStatus returns a boolean if a field has been set.
func (o *User) HasActiveStatus() bool {
	if o != nil && !IsNil(o.ActiveStatus) {
		return true
	}

	return false
}

// SetActiveStatus gets a reference to the given ActiveStatus and assigns it to the ActiveStatus field.
func (o *User) SetActiveStatus(v ActiveStatus) {
	o.ActiveStatus = &v
}

// GetPowerBadge returns the PowerBadge field value if set, zero value otherwise.
func (o *User) GetPowerBadge() bool {
	if o == nil || IsNil(o.PowerBadge) {
		var ret bool
		return ret
	}
	return *o.PowerBadge
}

// GetPowerBadgeOk returns a tuple with the PowerBadge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPowerBadgeOk() (*bool, bool) {
	if o == nil || IsNil(o.PowerBadge) {
		return nil, false
	}
	return o.PowerBadge, true
}

// HasPowerBadge returns a boolean if a field has been set.
func (o *User) HasPowerBadge() bool {
	if o != nil && !IsNil(o.PowerBadge) {
		return true
	}

	return false
}

// SetPowerBadge gets a reference to the given bool and assigns it to the PowerBadge field.
func (o *User) SetPowerBadge(v bool) {
	o.PowerBadge = &v
}

// GetViewerContext returns the ViewerContext field value if set, zero value otherwise.
func (o *User) GetViewerContext() UserViewerContext {
	if o == nil || IsNil(o.ViewerContext) {
		var ret UserViewerContext
		return ret
	}
	return *o.ViewerContext
}

// GetViewerContextOk returns a tuple with the ViewerContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetViewerContextOk() (*UserViewerContext, bool) {
	if o == nil || IsNil(o.ViewerContext) {
		return nil, false
	}
	return o.ViewerContext, true
}

// HasViewerContext returns a boolean if a field has been set.
func (o *User) HasViewerContext() bool {
	if o != nil && !IsNil(o.ViewerContext) {
		return true
	}

	return false
}

// SetViewerContext gets a reference to the given UserViewerContext and assigns it to the ViewerContext field.
func (o *User) SetViewerContext(v UserViewerContext) {
	o.ViewerContext = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Fid) {
		toSerialize["fid"] = o.Fid
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.CustodyAddress) {
		toSerialize["custody_address"] = o.CustodyAddress
	}
	if !IsNil(o.PfpUrl) {
		toSerialize["pfp_url"] = o.PfpUrl
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.FollowerCount) {
		toSerialize["follower_count"] = o.FollowerCount
	}
	if !IsNil(o.FollowingCount) {
		toSerialize["following_count"] = o.FollowingCount
	}
	if !IsNil(o.Verifications) {
		toSerialize["verifications"] = o.Verifications
	}
	if !IsNil(o.VerifiedAddresses) {
		toSerialize["verified_addresses"] = o.VerifiedAddresses
	}
	if !IsNil(o.ActiveStatus) {
		toSerialize["active_status"] = o.ActiveStatus
	}
	if !IsNil(o.PowerBadge) {
		toSerialize["power_badge"] = o.PowerBadge
	}
	if !IsNil(o.ViewerContext) {
		toSerialize["viewer_context"] = o.ViewerContext
	}
	return toSerialize, nil
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


