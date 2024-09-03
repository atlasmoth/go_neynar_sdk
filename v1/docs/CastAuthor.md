# CastAuthor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **string** |  | 
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

### NewCastAuthor

`func NewCastAuthor(fid string, username string, custodyAddress string, displayName string, pfp UserPfp, profile UserProfile, followerCount int32, followingCount int32, verifications []string, activeStatus ActiveStatus, ) *CastAuthor`

NewCastAuthor instantiates a new CastAuthor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastAuthorWithDefaults

`func NewCastAuthorWithDefaults() *CastAuthor`

NewCastAuthorWithDefaults instantiates a new CastAuthor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *CastAuthor) GetFid() string`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CastAuthor) GetFidOk() (*string, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CastAuthor) SetFid(v string)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *CastAuthor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CastAuthor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CastAuthor) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetCustodyAddress

`func (o *CastAuthor) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *CastAuthor) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *CastAuthor) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.


### GetDisplayName

`func (o *CastAuthor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CastAuthor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CastAuthor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPfp

`func (o *CastAuthor) GetPfp() UserPfp`

GetPfp returns the Pfp field if non-nil, zero value otherwise.

### GetPfpOk

`func (o *CastAuthor) GetPfpOk() (*UserPfp, bool)`

GetPfpOk returns a tuple with the Pfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfp

`func (o *CastAuthor) SetPfp(v UserPfp)`

SetPfp sets Pfp field to given value.


### GetProfile

`func (o *CastAuthor) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CastAuthor) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CastAuthor) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.


### GetFollowerCount

`func (o *CastAuthor) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *CastAuthor) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *CastAuthor) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.


### GetFollowingCount

`func (o *CastAuthor) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *CastAuthor) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *CastAuthor) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.


### GetVerifications

`func (o *CastAuthor) GetVerifications() []string`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *CastAuthor) GetVerificationsOk() (*[]string, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *CastAuthor) SetVerifications(v []string)`

SetVerifications sets Verifications field to given value.


### GetActiveStatus

`func (o *CastAuthor) GetActiveStatus() ActiveStatus`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *CastAuthor) GetActiveStatusOk() (*ActiveStatus, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *CastAuthor) SetActiveStatus(v ActiveStatus)`

SetActiveStatus sets ActiveStatus field to given value.


### GetViewerContext

`func (o *CastAuthor) GetViewerContext() ViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastAuthor) GetViewerContextOk() (*ViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastAuthor) SetViewerContext(v ViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastAuthor) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


