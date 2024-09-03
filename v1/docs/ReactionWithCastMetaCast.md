# ReactionWithCastMetaCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CastFid** | **int32** | User identifier (unsigned integer) | [default to 3]
**CastHash** | **string** |  | 
**CastText** | **string** |  | 
**CastEmbeds** | [**[]EmbedUrl**](EmbedUrl.md) |  | 
**CastTimestamp** | **time.Time** |  | 

## Methods

### NewReactionWithCastMetaCast

`func NewReactionWithCastMetaCast(castFid int32, castHash string, castText string, castEmbeds []EmbedUrl, castTimestamp time.Time, ) *ReactionWithCastMetaCast`

NewReactionWithCastMetaCast instantiates a new ReactionWithCastMetaCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithCastMetaCastWithDefaults

`func NewReactionWithCastMetaCastWithDefaults() *ReactionWithCastMetaCast`

NewReactionWithCastMetaCastWithDefaults instantiates a new ReactionWithCastMetaCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCastFid

`func (o *ReactionWithCastMetaCast) GetCastFid() int32`

GetCastFid returns the CastFid field if non-nil, zero value otherwise.

### GetCastFidOk

`func (o *ReactionWithCastMetaCast) GetCastFidOk() (*int32, bool)`

GetCastFidOk returns a tuple with the CastFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastFid

`func (o *ReactionWithCastMetaCast) SetCastFid(v int32)`

SetCastFid sets CastFid field to given value.


### GetCastHash

`func (o *ReactionWithCastMetaCast) GetCastHash() string`

GetCastHash returns the CastHash field if non-nil, zero value otherwise.

### GetCastHashOk

`func (o *ReactionWithCastMetaCast) GetCastHashOk() (*string, bool)`

GetCastHashOk returns a tuple with the CastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastHash

`func (o *ReactionWithCastMetaCast) SetCastHash(v string)`

SetCastHash sets CastHash field to given value.


### GetCastText

`func (o *ReactionWithCastMetaCast) GetCastText() string`

GetCastText returns the CastText field if non-nil, zero value otherwise.

### GetCastTextOk

`func (o *ReactionWithCastMetaCast) GetCastTextOk() (*string, bool)`

GetCastTextOk returns a tuple with the CastText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastText

`func (o *ReactionWithCastMetaCast) SetCastText(v string)`

SetCastText sets CastText field to given value.


### GetCastEmbeds

`func (o *ReactionWithCastMetaCast) GetCastEmbeds() []EmbedUrl`

GetCastEmbeds returns the CastEmbeds field if non-nil, zero value otherwise.

### GetCastEmbedsOk

`func (o *ReactionWithCastMetaCast) GetCastEmbedsOk() (*[]EmbedUrl, bool)`

GetCastEmbedsOk returns a tuple with the CastEmbeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastEmbeds

`func (o *ReactionWithCastMetaCast) SetCastEmbeds(v []EmbedUrl)`

SetCastEmbeds sets CastEmbeds field to given value.


### GetCastTimestamp

`func (o *ReactionWithCastMetaCast) GetCastTimestamp() time.Time`

GetCastTimestamp returns the CastTimestamp field if non-nil, zero value otherwise.

### GetCastTimestampOk

`func (o *ReactionWithCastMetaCast) GetCastTimestampOk() (*time.Time, bool)`

GetCastTimestampOk returns a tuple with the CastTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastTimestamp

`func (o *ReactionWithCastMetaCast) SetCastTimestamp(v time.Time)`

SetCastTimestamp sets CastTimestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


