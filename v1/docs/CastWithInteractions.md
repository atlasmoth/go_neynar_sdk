# CastWithInteractions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**ParentHash** | Pointer to **NullableString** |  | [optional] 
**ParentUrl** | Pointer to **NullableString** |  | [optional] 
**ThreadHash** | Pointer to **string** |  | [optional] 
**ParentAuthor** | Pointer to [**CastParentAuthor**](CastParentAuthor.md) |  | [optional] 
**MentionedProfiles** | Pointer to [**[]User**](User.md) |  | [optional] 
**Author** | Pointer to [**CastAuthor**](CastAuthor.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Embeds** | Pointer to [**[]EmbedUrl**](EmbedUrl.md) |  | [optional] 
**Type** | Pointer to [**CastType**](CastType.md) |  | [optional] 
**Reactions** | Pointer to [**CastWithInteractionsReactions**](CastWithInteractionsReactions.md) |  | [optional] 
**Recasts** | Pointer to [**CastWithInteractionsRecasts**](CastWithInteractionsRecasts.md) |  | [optional] 
**Recasters** | Pointer to **[]string** |  | [optional] 
**ViewerContext** | Pointer to [**ViewerContext**](ViewerContext.md) |  | [optional] 
**Replies** | Pointer to [**CastWithInteractionsReplies**](CastWithInteractionsReplies.md) |  | [optional] 

## Methods

### NewCastWithInteractions

`func NewCastWithInteractions() *CastWithInteractions`

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

### HasHash

`func (o *CastWithInteractions) HasHash() bool`

HasHash returns a boolean if a field has been set.

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

### HasParentHash

`func (o *CastWithInteractions) HasParentHash() bool`

HasParentHash returns a boolean if a field has been set.

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

### HasParentUrl

`func (o *CastWithInteractions) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

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

### HasThreadHash

`func (o *CastWithInteractions) HasThreadHash() bool`

HasThreadHash returns a boolean if a field has been set.

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

### HasParentAuthor

`func (o *CastWithInteractions) HasParentAuthor() bool`

HasParentAuthor returns a boolean if a field has been set.

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

### HasMentionedProfiles

`func (o *CastWithInteractions) HasMentionedProfiles() bool`

HasMentionedProfiles returns a boolean if a field has been set.

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

### HasAuthor

`func (o *CastWithInteractions) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

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

### HasText

`func (o *CastWithInteractions) HasText() bool`

HasText returns a boolean if a field has been set.

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

### HasTimestamp

`func (o *CastWithInteractions) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

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

### HasEmbeds

`func (o *CastWithInteractions) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

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

### HasReactions

`func (o *CastWithInteractions) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

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

### HasRecasts

`func (o *CastWithInteractions) HasRecasts() bool`

HasRecasts returns a boolean if a field has been set.

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

### HasRecasters

`func (o *CastWithInteractions) HasRecasters() bool`

HasRecasters returns a boolean if a field has been set.

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

### HasReplies

`func (o *CastWithInteractions) HasReplies() bool`

HasReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


