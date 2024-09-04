# Recaster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | Pointer to **int32** | The unique identifier of the recaster. | [optional] 
**Username** | Pointer to **string** | The username of the recaster. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the recaster. | [optional] 
**Pfp** | Pointer to [**RecasterPfp**](RecasterPfp.md) |  | [optional] 
**Profile** | Pointer to [**RecasterProfile**](RecasterProfile.md) |  | [optional] 
**FollowerCount** | Pointer to **int32** | The number of followers the recaster has. | [optional] 
**FollowingCount** | Pointer to **int32** | The number of users the recaster is following. | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**ViewerContext** | Pointer to [**RecasterViewerContext**](RecasterViewerContext.md) |  | [optional] 

## Methods

### NewRecaster

`func NewRecaster() *Recaster`

NewRecaster instantiates a new Recaster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecasterWithDefaults

`func NewRecasterWithDefaults() *Recaster`

NewRecasterWithDefaults instantiates a new Recaster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *Recaster) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *Recaster) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *Recaster) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *Recaster) HasFid() bool`

HasFid returns a boolean if a field has been set.

### GetUsername

`func (o *Recaster) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Recaster) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Recaster) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Recaster) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *Recaster) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Recaster) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Recaster) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Recaster) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPfp

`func (o *Recaster) GetPfp() RecasterPfp`

GetPfp returns the Pfp field if non-nil, zero value otherwise.

### GetPfpOk

`func (o *Recaster) GetPfpOk() (*RecasterPfp, bool)`

GetPfpOk returns a tuple with the Pfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfp

`func (o *Recaster) SetPfp(v RecasterPfp)`

SetPfp sets Pfp field to given value.

### HasPfp

`func (o *Recaster) HasPfp() bool`

HasPfp returns a boolean if a field has been set.

### GetProfile

`func (o *Recaster) GetProfile() RecasterProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *Recaster) GetProfileOk() (*RecasterProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *Recaster) SetProfile(v RecasterProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *Recaster) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetFollowerCount

`func (o *Recaster) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *Recaster) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *Recaster) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *Recaster) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetFollowingCount

`func (o *Recaster) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *Recaster) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *Recaster) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *Recaster) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetTimestamp

`func (o *Recaster) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Recaster) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Recaster) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Recaster) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetViewerContext

`func (o *Recaster) GetViewerContext() RecasterViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *Recaster) GetViewerContextOk() (*RecasterViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *Recaster) SetViewerContext(v RecasterViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *Recaster) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


