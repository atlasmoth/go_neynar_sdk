# RecasterViewerContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Following** | Pointer to **bool** | Indicates if the viewer is following the recaster. | [optional] 
**FollowedBy** | Pointer to **bool** | Indicates if the recaster is followed by the viewer. | [optional] 

## Methods

### NewRecasterViewerContext

`func NewRecasterViewerContext() *RecasterViewerContext`

NewRecasterViewerContext instantiates a new RecasterViewerContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecasterViewerContextWithDefaults

`func NewRecasterViewerContextWithDefaults() *RecasterViewerContext`

NewRecasterViewerContextWithDefaults instantiates a new RecasterViewerContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowing

`func (o *RecasterViewerContext) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *RecasterViewerContext) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *RecasterViewerContext) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *RecasterViewerContext) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetFollowedBy

`func (o *RecasterViewerContext) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *RecasterViewerContext) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *RecasterViewerContext) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.

### HasFollowedBy

`func (o *RecasterViewerContext) HasFollowedBy() bool`

HasFollowedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


