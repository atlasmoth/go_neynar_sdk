# ReactionWithCastMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reaction** | [**ReactionWithCastMetaReaction**](ReactionWithCastMetaReaction.md) |  | 
**Cast** | Pointer to [**ReactionWithCastMetaCast**](ReactionWithCastMetaCast.md) |  | [optional] 
**CastAuthor** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewReactionWithCastMeta

`func NewReactionWithCastMeta(reaction ReactionWithCastMetaReaction, ) *ReactionWithCastMeta`

NewReactionWithCastMeta instantiates a new ReactionWithCastMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithCastMetaWithDefaults

`func NewReactionWithCastMetaWithDefaults() *ReactionWithCastMeta`

NewReactionWithCastMetaWithDefaults instantiates a new ReactionWithCastMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReaction

`func (o *ReactionWithCastMeta) GetReaction() ReactionWithCastMetaReaction`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *ReactionWithCastMeta) GetReactionOk() (*ReactionWithCastMetaReaction, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *ReactionWithCastMeta) SetReaction(v ReactionWithCastMetaReaction)`

SetReaction sets Reaction field to given value.


### GetCast

`func (o *ReactionWithCastMeta) GetCast() ReactionWithCastMetaCast`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *ReactionWithCastMeta) GetCastOk() (*ReactionWithCastMetaCast, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *ReactionWithCastMeta) SetCast(v ReactionWithCastMetaCast)`

SetCast sets Cast field to given value.

### HasCast

`func (o *ReactionWithCastMeta) HasCast() bool`

HasCast returns a boolean if a field has been set.

### GetCastAuthor

`func (o *ReactionWithCastMeta) GetCastAuthor() User`

GetCastAuthor returns the CastAuthor field if non-nil, zero value otherwise.

### GetCastAuthorOk

`func (o *ReactionWithCastMeta) GetCastAuthorOk() (*User, bool)`

GetCastAuthorOk returns a tuple with the CastAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastAuthor

`func (o *ReactionWithCastMeta) SetCastAuthor(v User)`

SetCastAuthor sets CastAuthor field to given value.

### HasCastAuthor

`func (o *ReactionWithCastMeta) HasCastAuthor() bool`

HasCastAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


