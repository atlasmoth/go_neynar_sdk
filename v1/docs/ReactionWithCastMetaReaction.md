# ReactionWithCastMetaReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReactorFid** | **int32** | User identifier (unsigned integer) | [default to 3]
**ReactionType** | [**ReactionType**](ReactionType.md) |  | 
**ReactionHash** | **string** |  | 
**ReactionTargetHash** | **string** |  | 
**ReactionTimestamp** | **time.Time** |  | 

## Methods

### NewReactionWithCastMetaReaction

`func NewReactionWithCastMetaReaction(reactorFid int32, reactionType ReactionType, reactionHash string, reactionTargetHash string, reactionTimestamp time.Time, ) *ReactionWithCastMetaReaction`

NewReactionWithCastMetaReaction instantiates a new ReactionWithCastMetaReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithCastMetaReactionWithDefaults

`func NewReactionWithCastMetaReactionWithDefaults() *ReactionWithCastMetaReaction`

NewReactionWithCastMetaReactionWithDefaults instantiates a new ReactionWithCastMetaReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReactorFid

`func (o *ReactionWithCastMetaReaction) GetReactorFid() int32`

GetReactorFid returns the ReactorFid field if non-nil, zero value otherwise.

### GetReactorFidOk

`func (o *ReactionWithCastMetaReaction) GetReactorFidOk() (*int32, bool)`

GetReactorFidOk returns a tuple with the ReactorFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactorFid

`func (o *ReactionWithCastMetaReaction) SetReactorFid(v int32)`

SetReactorFid sets ReactorFid field to given value.


### GetReactionType

`func (o *ReactionWithCastMetaReaction) GetReactionType() ReactionType`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionWithCastMetaReaction) GetReactionTypeOk() (*ReactionType, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionWithCastMetaReaction) SetReactionType(v ReactionType)`

SetReactionType sets ReactionType field to given value.


### GetReactionHash

`func (o *ReactionWithCastMetaReaction) GetReactionHash() string`

GetReactionHash returns the ReactionHash field if non-nil, zero value otherwise.

### GetReactionHashOk

`func (o *ReactionWithCastMetaReaction) GetReactionHashOk() (*string, bool)`

GetReactionHashOk returns a tuple with the ReactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionHash

`func (o *ReactionWithCastMetaReaction) SetReactionHash(v string)`

SetReactionHash sets ReactionHash field to given value.


### GetReactionTargetHash

`func (o *ReactionWithCastMetaReaction) GetReactionTargetHash() string`

GetReactionTargetHash returns the ReactionTargetHash field if non-nil, zero value otherwise.

### GetReactionTargetHashOk

`func (o *ReactionWithCastMetaReaction) GetReactionTargetHashOk() (*string, bool)`

GetReactionTargetHashOk returns a tuple with the ReactionTargetHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionTargetHash

`func (o *ReactionWithCastMetaReaction) SetReactionTargetHash(v string)`

SetReactionTargetHash sets ReactionTargetHash field to given value.


### GetReactionTimestamp

`func (o *ReactionWithCastMetaReaction) GetReactionTimestamp() time.Time`

GetReactionTimestamp returns the ReactionTimestamp field if non-nil, zero value otherwise.

### GetReactionTimestampOk

`func (o *ReactionWithCastMetaReaction) GetReactionTimestampOk() (*time.Time, bool)`

GetReactionTimestampOk returns a tuple with the ReactionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionTimestamp

`func (o *ReactionWithCastMetaReaction) SetReactionTimestamp(v time.Time)`

SetReactionTimestamp sets ReactionTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


