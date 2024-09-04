# FollowResponseUser

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
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFollowResponseUser

`func NewFollowResponseUser() *FollowResponseUser`

NewFollowResponseUser instantiates a new FollowResponseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowResponseUserWithDefaults

`func NewFollowResponseUserWithDefaults() *FollowResponseUser`

NewFollowResponseUserWithDefaults instantiates a new FollowResponseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *FollowResponseUser) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *FollowResponseUser) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *FollowResponseUser) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *FollowResponseUser) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetUsername

`func (o *FollowResponseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *FollowResponseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *FollowResponseUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *FollowResponseUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetCustodyAddress

`func (o *FollowResponseUser) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *FollowResponseUser) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *FollowResponseUser) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.

### HasCustodyAddress

`func (o *FollowResponseUser) HasCustodyAddress() bool`

HasCustodyAddress returns a boolean if a field has been set.

### GetDisplayName

`func (o *FollowResponseUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FollowResponseUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FollowResponseUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FollowResponseUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfp

`func (o *FollowResponseUser) GetPfp() UserPfp`

GetPfp returns the Pfp field if non-nil, zero value otherwise.

### GetPfpOk

`func (o *FollowResponseUser) GetPfpOk() (*UserPfp, bool)`

GetPfpOk returns a tuple with the Pfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfp

`func (o *FollowResponseUser) SetPfp(v UserPfp)`

SetPfp sets Pfp field to given value.

### HasPfp

`func (o *FollowResponseUser) HasPfp() bool`

HasPfp returns a boolean if a field has been set.

### GetProfile

`func (o *FollowResponseUser) GetProfile() UserProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FollowResponseUser) GetProfileOk() (*UserProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FollowResponseUser) SetProfile(v UserProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FollowResponseUser) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetFollowerCount

`func (o *FollowResponseUser) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *FollowResponseUser) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *FollowResponseUser) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *FollowResponseUser) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetFollowingCount

`func (o *FollowResponseUser) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *FollowResponseUser) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *FollowResponseUser) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *FollowResponseUser) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetVerifications

`func (o *FollowResponseUser) GetVerifications() []string`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *FollowResponseUser) GetVerificationsOk() (*[]string, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *FollowResponseUser) SetVerifications(v []string)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *FollowResponseUser) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.

### GetActiveStatus

`func (o *FollowResponseUser) GetActiveStatus() ActiveStatus`

GetActiveStatus returns the ActiveStatus field if non-nil, zero value otherwise.

### GetActiveStatusOk

`func (o *FollowResponseUser) GetActiveStatusOk() (*ActiveStatus, bool)`

GetActiveStatusOk returns a tuple with the ActiveStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveStatus

`func (o *FollowResponseUser) SetActiveStatus(v ActiveStatus)`

SetActiveStatus sets ActiveStatus field to given value.

### HasActiveStatus

`func (o *FollowResponseUser) HasActiveStatus() bool`

HasActiveStatus returns a boolean if a field has been set.

### GetViewerContext

`func (o *FollowResponseUser) GetViewerContext() ViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *FollowResponseUser) GetViewerContextOk() (*ViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *FollowResponseUser) SetViewerContext(v ViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *FollowResponseUser) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetTimestamp

`func (o *FollowResponseUser) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *FollowResponseUser) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *FollowResponseUser) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *FollowResponseUser) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


