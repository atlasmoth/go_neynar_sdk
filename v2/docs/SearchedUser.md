# SearchedUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**Fid** | Pointer to **int32** | User identifier (unsigned integer) | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**CustodyAddress** | Pointer to **string** | Ethereum address | [optional] 
**PfpUrl** | Pointer to **string** | The URL of the user&#39;s profile picture | [optional] 
**Profile** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 
**FollowerCount** | Pointer to **int32** | The number of followers the user has. | [optional] 
**FollowingCount** | Pointer to **int32** | The number of users the user is following. | [optional] 
**Verifications** | Pointer to **[]string** |  | [optional] 
**VerifiedAddresses** | Pointer to [**UserVerifiedAddresses**](UserVerifiedAddresses.md) |  | [optional] 
**ActiveStatus** | Pointer to [**ActiveStatus**](ActiveStatus.md) |  | [optional] 
**PowerBadge** | Pointer to **bool** |  | [optional] 
**ViewerContext** | Pointer to [**UserViewerContext**](UserViewerContext.md) |  | [optional] 
**Pfp** | Pointer to [**ProfileUrlPfp**](ProfileUrlPfp.md) |  | [optional] 

## Methods

### NewSearchedUser

`func NewSearchedUser() *SearchedUser`

NewSearchedUser instantiates a new SearchedUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchedUserWithDefaults

`func NewSearchedUserWithDefaults() *SearchedUser`

NewSearchedUserWithDefaults instantiates a new SearchedUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *SearchedUser) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SearchedUser) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SearchedUser) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SearchedUser) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetFid

`func (o *SearchedUser) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *SearchedUser) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *SearchedUser) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *SearchedUser) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetUsername

`func (o *SearchedUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SearchedUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SearchedUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SearchedUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *SearchedUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SearchedUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SearchedUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SearchedUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *SearchedUser) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *SearchedUser) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *SearchedUser) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.

### HasCustodyAddress

`func (o *SearchedUser) HasCustodyAddress() bool`

HasCustodyAddress returns a boolean if a field has been set.

### GetPfpUrl

`func (o *SearchedUser) GetPfpUrl() string`

GetPfpUrl returns the PfpUrl field if non-nil, zero value otherwise.

### GetPfpUrlOk

`func (o *SearchedUser) GetPfpUrlOk() (*string, bool)`

GetPfpUrlOk returns a tuple with the PfpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfpUrl

`func (o *SearchedUser) SetPfpUrl(v string)`

SetPfpUrl sets PfpUrl field to given value.

### HasPfpUrl

`func (o *SearchedUser) HasPfpUrl() bool`

HasPfpUrl returns a boolean if a field has been set.

### GetProfile

`func (o *SearchedUser) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SearchedUser) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SearchedUser) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SearchedUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetFollowerCount

`func (o *SearchedUser) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *SearchedUser) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *SearchedUser) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *SearchedUser) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetFollowingCount

`func (o *SearchedUser) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *SearchedUser) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *SearchedUser) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *SearchedUser) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetVerifications

`func (o *SearchedUser) GetVerifications() []string`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *SearchedUser) GetVerificationsOk() (*[]string, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *SearchedUser) SetVerifications(v []string)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *SearchedUser) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.

### GetVerifiedAddresses

`func (o *SearchedUser) GetVerifiedAddresses() UserVerifiedAddresses`

GetVerifiedAddresses returns the VerifiedAddresses field if non-nil, zero value otherwise.

### GetVerifiedAddressesOk

`func (o *SearchedUser) GetVerifiedAddressesOk() (*UserVerifiedAddresses, bool)`

GetVerifiedAddressesOk returns a tuple with the VerifiedAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAddresses

`func (o *SearchedUser) SetVerifiedAddresses(v UserVerifiedAddresses)`

SetVerifiedAddresses sets VerifiedAddresses field to given value.

### HasVerifiedAddresses

`func (o *SearchedUser) HasVerifiedAddresses() bool`

HasVerifiedAddresses returns a boolean if a field has been set.

### GetActiveStatus

`func (o *SearchedUser) GetActiveStatus() ActiveStatus`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *SearchedUser) GetActiveStatusOk() (*ActiveStatus, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *SearchedUser) SetActiveStatus(v ActiveStatus)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *SearchedUser) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetPowerBadge

`func (o *SearchedUser) GetPowerBadge() bool`

GetPowerBadge returns the PowerBadge field if non-nil, zero value otherwise.

### GetPowerBadgeOk

`func (o *SearchedUser) GetPowerBadgeOk() (*bool, bool)`

GetPowerBadgeOk returns a tuple with the PowerBadge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerBadge

`func (o *SearchedUser) SetPowerBadge(v bool)`

SetPowerBadge sets PowerBadge field to given value.

### HasPowerBadge

`func (o *SearchedUser) HasPowerBadge() bool`

HasPowerBadge returns a boolean if a field has been set.

### GetViewerContext

`func (o *SearchedUser) GetViewerContext() UserViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *SearchedUser) GetViewerContextOk() (*UserViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *SearchedUser) SetViewerContext(v UserViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *SearchedUser) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetPfp

`func (o *SearchedUser) GetPfp() ProfileUrlPfp`

GetPfp returns the Pfp field if non-nil, zero value otherwise.

### GetPfpOk

`func (o *SearchedUser) GetPfpOk() (*ProfileUrlPfp, bool)`

GetPfpOk returns a tuple with the Pfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfp

`func (o *SearchedUser) SetPfp(v ProfileUrlPfp)`

SetPfp sets Pfp field to given value.

### HasPfp

`func (o *SearchedUser) HasPfp() bool`

HasPfp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


