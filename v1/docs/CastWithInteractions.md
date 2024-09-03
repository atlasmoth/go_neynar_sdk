# CastWithInteractions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**ParentHash** | **NullableString** |  | 
**ParentUrl** | **NullableString** |  | 
**ThreadHash** | **string** |  | 
**ParentAuthor** | [**CastParentAuthor**](CastParentAuthor.md) |  | 
**MentionedProfiles** | [**[]User**](User.md) |  | 
**Author** | [**CastAuthor**](CastAuthor.md) |  | 
**Text** | **string** |  | 
**Timestamp** | **time.Time** |  | 
**Embeds** | [**[]EmbedUrl**](EmbedUrl.md) |  | 
**Type** | Pointer to [**CastType**](CastType.md) |  | [optional] 
**Reactions** | [**CastWithInteractionsReactions**](CastWithInteractionsReactions.md) |  | 
**Recasts** | [**CastWithInteractionsRecasts**](CastWithInteractionsRecasts.md) |  | 
**Recasters** | **[]string** |  | 
**ViewerContext** | Pointer to [**ViewerContext**](ViewerContext.md) |  | [optional] 
**Replies** | [**CastWithInteractionsReplies**](CastWithInteractionsReplies.md) |  | 

## Methods

### NewCastWithInteractions

`func NewCastWithInteractions(hash string, parentHash NullableString, parentUrl NullableString, threadHash string, parentAuthor CastParentAuthor, mentionedProfiles []User, author CastAuthor, text string, timestamp time.Time, embeds []EmbedUrl, reactions CastWithInteractionsReactions, recasts CastWithInteractionsRecasts, recasters []string, replies CastWithInteractionsReplies, ) *CastWithInteractions`

NewCastWithInteractions instantiates a new CastWithInteractions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithInteractionsWithDefaults

`func NewCastWithInteractionsWithDefaults() *CastWithInteractions`

NewCastWithInteractionsWithDefaults instantiates a new CastWithInteractions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *CastWithInteractions) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastWithInteractions) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastWithInteractions) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *CastWithInteractions) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *CastWithInteractions) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *CastWithInteractions) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *CastWithInteractions) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *CastWithInteractions) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *CastWithInteractions) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastWithInteractions) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastWithInteractions) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *CastWithInteractions) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *CastWithInteractions) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetThreadHash

`func (o *CastWithInteractions) GetThreadHash() string`

GetThreadHash returns the ThreadHash field if non-nil, zero value otherwise.

### GetThreadHashOk

`func (o *CastWithInteractions) GetThreadHashOk() (*string, bool)`

GetThreadHashOk returns a tuple with the ThreadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadHash

`func (o *CastWithInteractions) SetThreadHash(v string)`

SetThreadHash sets ThreadHash field to given value.


### GetParentAuthor

`func (o *CastWithInteractions) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *CastWithInteractions) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *CastWithInteractions) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetMentionedProfiles

`func (o *CastWithInteractions) GetMentionedProfiles() []User`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *CastWithInteractions) GetMentionedProfilesOk() (*[]User, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *CastWithInteractions) SetMentionedProfiles(v []User)`

SetMentionedProfiles sets MentionedProfiles field to given value.


### GetAuthor

`func (o *CastWithInteractions) GetAuthor() CastAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastWithInteractions) GetAuthorOk() (*CastAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastWithInteractions) SetAuthor(v CastAuthor)`

SetAuthor sets Author field to given value.


### GetText

`func (o *CastWithInteractions) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastWithInteractions) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastWithInteractions) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *CastWithInteractions) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastWithInteractions) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastWithInteractions) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *CastWithInteractions) GetEmbeds() []EmbedUrl`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastWithInteractions) GetEmbedsOk() (*[]EmbedUrl, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastWithInteractions) SetEmbeds(v []EmbedUrl)`

SetEmbeds sets Embeds field to given value.


### GetType

`func (o *CastWithInteractions) GetType() CastType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastWithInteractions) GetTypeOk() (*CastType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastWithInteractions) SetType(v CastType)`

SetType sets Type field to given value.

### HasType

`func (o *CastWithInteractions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReactions

`func (o *CastWithInteractions) GetReactions() CastWithInteractionsReactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CastWithInteractions) GetReactionsOk() (*CastWithInteractionsReactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CastWithInteractions) SetReactions(v CastWithInteractionsReactions)`

SetReactions sets Reactions field to given value.


### GetRecasts

`func (o *CastWithInteractions) GetRecasts() CastWithInteractionsRecasts`

GetRecasts returns the Recasts field if non-nil, zero value otherwise.

### GetRecastsOk

`func (o *CastWithInteractions) GetRecastsOk() (*CastWithInteractionsRecasts, bool)`

GetRecastsOk returns a tuple with the Recasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecasts

`func (o *CastWithInteractions) SetRecasts(v CastWithInteractionsRecasts)`

SetRecasts sets Recasts field to given value.


### GetRecasters

`func (o *CastWithInteractions) GetRecasters() []string`

GetRecasters returns the Recasters field if non-nil, zero value otherwise.

### GetRecastersOk

`func (o *CastWithInteractions) GetRecastersOk() (*[]string, bool)`

GetRecastersOk returns a tuple with the Recasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecasters

`func (o *CastWithInteractions) SetRecasters(v []string)`

SetRecasters sets Recasters field to given value.


### GetViewerContext

`func (o *CastWithInteractions) GetViewerContext() ViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastWithInteractions) GetViewerContextOk() (*ViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastWithInteractions) SetViewerContext(v ViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastWithInteractions) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetReplies

`func (o *CastWithInteractions) GetReplies() CastWithInteractionsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CastWithInteractions) GetRepliesOk() (*CastWithInteractionsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CastWithInteractions) SetReplies(v CastWithInteractionsReplies)`

SetReplies sets Replies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


