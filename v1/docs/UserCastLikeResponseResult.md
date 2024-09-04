# UserCastLikeResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reactor** | Pointer to [**User**](User.md) |  | [optional] 
**Likes** | Pointer to [**[]ReactionWithCastMeta**](ReactionWithCastMeta.md) |  | [optional] 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewUserCastLikeResponseResult

`func NewUserCastLikeResponseResult() *UserCastLikeResponseResult`

NewUserCastLikeResponseResult instantiates a new UserCastLikeResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCastLikeResponseResultWithDefaults

`func NewUserCastLikeResponseResultWithDefaults() *UserCastLikeResponseResult`

NewUserCastLikeResponseResultWithDefaults instantiates a new UserCastLikeResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactor

`func (o *UserCastLikeResponseResult) GetReactor() User`

GetReactor returns the Reactor field if non-nil, zero value otherwise.

### GetReactorOk

`func (o *UserCastLikeResponseResult) GetReactorOk() (*User, bool)`

GetReactorOk returns a tuple with the Reactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactor

`func (o *UserCastLikeResponseResult) SetReactor(v User)`

SetReactor sets Reactor field to given value.

### HasReactor

`func (o *UserCastLikeResponseResult) HasReactor() bool`

HasReactor returns a boolean if a field has been set.

### GetLikes

`func (o *UserCastLikeResponseResult) GetLikes() []ReactionWithCastMeta`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *UserCastLikeResponseResult) GetLikesOk() (*[]ReactionWithCastMeta, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *UserCastLikeResponseResult) SetLikes(v []ReactionWithCastMeta)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *UserCastLikeResponseResult) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetNext

`func (o *UserCastLikeResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *UserCastLikeResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *UserCastLikeResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *UserCastLikeResponseResult) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


