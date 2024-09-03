# ReactorViewerContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Following** | **bool** | Indicates if the viewer is following the reactor. | 
**FollowedBy** | **bool** | Indicates if the reactor is followed by the viewer. | 

## Methods

### NewReactorViewerContext

`func NewReactorViewerContext(following bool, followedBy bool, ) *ReactorViewerContext`

NewReactorViewerContext instantiates a new ReactorViewerContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactorViewerContextWithDefaults

`func NewReactorViewerContextWithDefaults() *ReactorViewerContext`

NewReactorViewerContextWithDefaults instantiates a new ReactorViewerContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowing

`func (o *ReactorViewerContext) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *ReactorViewerContext) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *ReactorViewerContext) SetFollowing(v bool)`

SetFollowing sets Following field to given value.


### GetFollowedBy

`func (o *ReactorViewerContext) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *ReactorViewerContext) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *ReactorViewerContext) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


