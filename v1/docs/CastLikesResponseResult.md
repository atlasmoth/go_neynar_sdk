# CastLikesResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Likes** | [**[]Reaction**](Reaction.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewCastLikesResponseResult

`func NewCastLikesResponseResult(likes []Reaction, next NextCursor, ) *CastLikesResponseResult`

NewCastLikesResponseResult instantiates a new CastLikesResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastLikesResponseResultWithDefaults

`func NewCastLikesResponseResultWithDefaults() *CastLikesResponseResult`

NewCastLikesResponseResultWithDefaults instantiates a new CastLikesResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLikes

`func (o *CastLikesResponseResult) GetLikes() []Reaction`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *CastLikesResponseResult) GetLikesOk() (*[]Reaction, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *CastLikesResponseResult) SetLikes(v []Reaction)`

SetLikes sets Likes field to given value.


### GetNext

`func (o *CastLikesResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CastLikesResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CastLikesResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


