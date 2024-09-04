# ViewerContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Following** | Pointer to **bool** |  | [optional] 
**FollowedBy** | Pointer to **bool** |  | [optional] 
**Liked** | Pointer to **bool** |  | [optional] 
**Recasted** | Pointer to **bool** |  | [optional] 

## Methods

### NewViewerContext

`func NewViewerContext() *ViewerContext`

NewViewerContext instantiates a new ViewerContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewerContextWithDefaults

`func NewViewerContextWithDefaults() *ViewerContext`

NewViewerContextWithDefaults instantiates a new ViewerContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollowing

`func (o *ViewerContext) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *ViewerContext) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *ViewerContext) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *ViewerContext) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetFollowedBy

`func (o *ViewerContext) GetFollowedBy() bool`

GetFollowedBy returns the FollowedBy field if non-nil, zero value otherwise.

### GetFollowedByOk

`func (o *ViewerContext) GetFollowedByOk() (*bool, bool)`

GetFollowedByOk returns a tuple with the FollowedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowedBy

`func (o *ViewerContext) SetFollowedBy(v bool)`

SetFollowedBy sets FollowedBy field to given value.

### HasFollowedBy

`func (o *ViewerContext) HasFollowedBy() bool`

HasFollowedBy returns a boolean if a field has been set.

### GetLiked

`func (o *ViewerContext) GetLiked() bool`

GetLiked returns the Liked field if non-nil, zero value otherwise.

### GetLikedOk

`func (o *ViewerContext) GetLikedOk() (*bool, bool)`

GetLikedOk returns a tuple with the Liked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiked

`func (o *ViewerContext) SetLiked(v bool)`

SetLiked sets Liked field to given value.

### HasLiked

`func (o *ViewerContext) HasLiked() bool`

HasLiked returns a boolean if a field has been set.

### GetRecasted

`func (o *ViewerContext) GetRecasted() bool`

GetRecasted returns the Recasted field if non-nil, zero value otherwise.

### GetRecastedOk

`func (o *ViewerContext) GetRecastedOk() (*bool, bool)`

GetRecastedOk returns a tuple with the Recasted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecasted

`func (o *ViewerContext) SetRecasted(v bool)`

SetRecasted sets Recasted field to given value.

### HasRecasted

`func (o *ViewerContext) HasRecasted() bool`

HasRecasted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


