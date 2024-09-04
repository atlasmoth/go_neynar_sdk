# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** | User identifier (unsigned integer) | [optional] [default to 3]
**Username** | Pointer to **string** | The username of the user. | [optional] 
**CustodyAddress** | Pointer to **string** | Custody Address of the user. | [optional] 
**DisplayName** | Pointer to **string** | The display of the reactor. | [optional] 
**Pfp** | Pointer to [**UserPfp**](UserPfp.md) |  | [optional] 
**Profile** | Pointer to [**UserProfile**](UserProfile.md) |  | [optional] 
**FollowerCount** | Pointer to **int32** | The number of followers the user has. | [optional] 
**FollowingCount** | Pointer to **int32** | The number of users the user is following. | [optional] 
**Verifications** | Pointer to **[]string** |  | [optional] 
**ActiveStatus** | Pointer to [**ActiveStatus**](ActiveStatus.md) |  | [optional] 
**ViewerContext** | Pointer to [**ViewerContext**](ViewerContext.md) |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *User) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *User) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *User) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *User) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *User) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *User) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *User) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.

### HasCustodyAddress

`func (o *User) HasCustodyAddress() bool`

HasCustodyAddress returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfp

`func (o *User) GetPfp() UserPfp`

GetPfp returns the Pfp field if non-nil, zero value otherwise.

### GetPfpOk

`func (o *User) GetPfpOk() (*UserPfp, bool)`

GetPfpOk returns a tuple with the Pfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfp

`func (o *User) SetPfp(v UserPfp)`

SetPfp sets Pfp field to given value.

### HasPfp

`func (o *User) HasPfp() bool`

HasPfp returns a boolean if a field has been set.

### GetProfile

`func (o *User) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *User) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *User) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *User) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetFollowerCount

`func (o *User) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *User) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *User) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *User) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetFollowingCount

`func (o *User) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *User) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *User) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *User) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetVerifications

`func (o *User) GetVerifications() []string`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *User) GetVerificationsOk() (*[]string, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *User) SetVerifications(v []string)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *User) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.

### GetActiveStatus

`func (o *User) GetActiveStatus() ActiveStatus`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *User) GetActiveStatusOk() (*ActiveStatus, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *User) SetActiveStatus(v ActiveStatus)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *User) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetViewerContext

`func (o *User) GetViewerContext() ViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *User) GetViewerContextOk() (*ViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *User) SetViewerContext(v ViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *User) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


