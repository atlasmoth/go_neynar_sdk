# Reactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | The unique identifier of the reactor. | 
**Username** | **string** | The username of the reactor. | 
**DisplayName** | **string** | The display name of the reactor. | 
**Pfp** | [**ReactorPfp**](ReactorPfp.md) |  | 
**FollowerCount** | **int32** | The number of followers the reactor has. | 
**FollowingCount** | **int32** | The number of users the reactor is following. | 
**ViewerContext** | Pointer to [**ReactorViewerContext**](ReactorViewerContext.md) |  | [optional] 

## Methods

### NewReactor

`func NewReactor(fid int32, username string, displayName string, pfp ReactorPfp, followerCount int32, followingCount int32, ) *Reactor`

NewReactor instantiates a new Reactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactorWithDefaults

`func NewReactorWithDefaults() *Reactor`

NewReactorWithDefaults instantiates a new Reactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *Reactor) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *Reactor) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *Reactor) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *Reactor) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Reactor) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Reactor) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *Reactor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Reactor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Reactor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetPfp

`func (o *Reactor) GetPfp() ReactorPfp`

GetPfp returns the Pfp field if non-nil, zero value otherwise.

### GetPfpOk

`func (o *Reactor) GetPfpOk() (*ReactorPfp, bool)`

GetPfpOk returns a tuple with the Pfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfp

`func (o *Reactor) SetPfp(v ReactorPfp)`

SetPfp sets Pfp field to given value.


### GetFollowerCount

`func (o *Reactor) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *Reactor) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *Reactor) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.


### GetFollowingCount

`func (o *Reactor) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *Reactor) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *Reactor) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.


### GetViewerContext

`func (o *Reactor) GetViewerContext() ReactorViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *Reactor) GetViewerContextOk() (*ReactorViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *Reactor) SetViewerContext(v ReactorViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *Reactor) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


