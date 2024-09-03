# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | User identifier (unsigned integer) | [default to 3]
**Username** | **string** | The username of the user. | 
**CustodyAddress** | **string** | Custody Address of the user. | 
**DisplayName** | **string** | The display of the reactor. | 
**Pfp** | [**UserPfp**](UserPfp.md) |  | 
**Profile** | [**UserProfile**](UserProfile.md) |  | 
**FollowerCount** | **int32** | The number of followers the user has. | 
**FollowingCount** | **int32** | The number of users the user is following. | 
**Verifications** | **[]string** |  | 
**ActiveStatus** | [**ActiveStatus**](ActiveStatus.md) |  | 
**ViewerContext** | Pointer to [**ViewerContext**](ViewerContext.md) |  | [optional] 

## Methods

### NewUser

`func NewUser(fid int32, username string, custodyAddress string, displayName string, pfp UserPfp, profile UserProfile, followerCount int32, followingCount int32, verifications []string, activeStatus ActiveStatus, ) *User`

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


