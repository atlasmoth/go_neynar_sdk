# CastWithInteractionsAndConversations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**ParentHash** | Pointer to **NullableString** |  | [optional] 
**ParentUrl** | Pointer to **NullableString** |  | [optional] 
**RootParentUrl** | Pointer to **NullableString** |  | [optional] 
**ParentAuthor** | Pointer to [**CastParentAuthor**](CastParentAuthor.md) |  | [optional] 
**Author** | Pointer to [**User**](User.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Embeds** | Pointer to [**[]EmbeddedCast**](EmbeddedCast.md) |  | [optional] 
**Type** | Pointer to [**CastNotificationType**](CastNotificationType.md) |  | [optional] 
**Frames** | Pointer to [**[]Frame**](Frame.md) |  | [optional] 
**Reactions** | Pointer to [**CastWithInteractionsReactions**](CastWithInteractionsReactions.md) |  | [optional] 
**Replies** | Pointer to [**CastWithInteractionsReplies**](CastWithInteractionsReplies.md) |  | [optional] 
**ThreadHash** | Pointer to **NullableString** |  | [optional] 
**MentionedProfiles** | Pointer to [**[]User**](User.md) |  | [optional] 
**Channel** | Pointer to [**ChannelOrDehydratedChannel**](ChannelOrDehydratedChannel.md) |  | [optional] 
**ViewerContext** | Pointer to [**CastViewerContext**](CastViewerContext.md) |  | [optional] 
**DirectReplies** | Pointer to [**[]CastWithInteractionsAndConversations**](CastWithInteractionsAndConversations.md) | note: This is recursive. It contains the direct replies to the cast and their direct replies up to n reply_depth. | [optional] 

## Methods

### NewCastWithInteractionsAndConversations

`func NewCastWithInteractionsAndConversations() *CastWithInteractionsAndConversations`

NewCastWithInteractionsAndConversations instantiates a new CastWithInteractionsAndConversations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithInteractionsAndConversationsWithDefaults

`func NewCastWithInteractionsAndConversationsWithDefaults() *CastWithInteractionsAndConversations`

NewCastWithInteractionsAndConversationsWithDefaults instantiates a new CastWithInteractionsAndConversations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *CastWithInteractionsAndConversations) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastWithInteractionsAndConversations) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastWithInteractionsAndConversations) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CastWithInteractionsAndConversations) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetParentHash

`func (o *CastWithInteractionsAndConversations) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *CastWithInteractionsAndConversations) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *CastWithInteractionsAndConversations) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.

### HasParentHash

`func (o *CastWithInteractionsAndConversations) HasParentHash() bool`

HasParentHash returns a boolean if a field has been set.

### SetParentHashNil

`func (o *CastWithInteractionsAndConversations) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *CastWithInteractionsAndConversations) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *CastWithInteractionsAndConversations) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *CastWithInteractionsAndConversations) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *CastWithInteractionsAndConversations) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *CastWithInteractionsAndConversations) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### SetParentUrlNil

`func (o *CastWithInteractionsAndConversations) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *CastWithInteractionsAndConversations) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *CastWithInteractionsAndConversations) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *CastWithInteractionsAndConversations) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *CastWithInteractionsAndConversations) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.

### HasRootParentUrl

`func (o *CastWithInteractionsAndConversations) HasRootParentUrl() bool`

HasRootParentUrl returns a boolean if a field has been set.

### SetRootParentUrlNil

`func (o *CastWithInteractionsAndConversations) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *CastWithInteractionsAndConversations) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *CastWithInteractionsAndConversations) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *CastWithInteractionsAndConversations) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *CastWithInteractionsAndConversations) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.

### HasParentAuthor

`func (o *CastWithInteractionsAndConversations) HasParentAuthor() bool`

HasParentAuthor returns a boolean if a field has been set.

### GetAuthor

`func (o *CastWithInteractionsAndConversations) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CastWithInteractionsAndConversations) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CastWithInteractionsAndConversations) SetAuthor(v User)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CastWithInteractionsAndConversations) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetText

`func (o *CastWithInteractionsAndConversations) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CastWithInteractionsAndConversations) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CastWithInteractionsAndConversations) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CastWithInteractionsAndConversations) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *CastWithInteractionsAndConversations) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CastWithInteractionsAndConversations) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CastWithInteractionsAndConversations) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CastWithInteractionsAndConversations) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEmbeds

`func (o *CastWithInteractionsAndConversations) GetEmbeds() []EmbeddedCast`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *CastWithInteractionsAndConversations) GetEmbedsOk() (*[]EmbeddedCast, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *CastWithInteractionsAndConversations) SetEmbeds(v []EmbeddedCast)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *CastWithInteractionsAndConversations) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetType

`func (o *CastWithInteractionsAndConversations) GetType() CastNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CastWithInteractionsAndConversations) GetTypeOk() (*CastNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CastWithInteractionsAndConversations) SetType(v CastNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *CastWithInteractionsAndConversations) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrames

`func (o *CastWithInteractionsAndConversations) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *CastWithInteractionsAndConversations) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *CastWithInteractionsAndConversations) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *CastWithInteractionsAndConversations) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetReactions

`func (o *CastWithInteractionsAndConversations) GetReactions() CastWithInteractionsReactions`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *CastWithInteractionsAndConversations) GetReactionsOk() (*CastWithInteractionsReactions, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *CastWithInteractionsAndConversations) SetReactions(v CastWithInteractionsReactions)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *CastWithInteractionsAndConversations) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetReplies

`func (o *CastWithInteractionsAndConversations) GetReplies() CastWithInteractionsReplies`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *CastWithInteractionsAndConversations) GetRepliesOk() (*CastWithInteractionsReplies, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *CastWithInteractionsAndConversations) SetReplies(v CastWithInteractionsReplies)`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *CastWithInteractionsAndConversations) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### GetThreadHash

`func (o *CastWithInteractionsAndConversations) GetThreadHash() string`

GetThreadHash returns the ThreadHash field if non-nil, zero value otherwise.

### GetThreadHashOk

`func (o *CastWithInteractionsAndConversations) GetThreadHashOk() (*string, bool)`

GetThreadHashOk returns a tuple with the ThreadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadHash

`func (o *CastWithInteractionsAndConversations) SetThreadHash(v string)`

SetThreadHash sets ThreadHash field to given value.

### HasThreadHash

`func (o *CastWithInteractionsAndConversations) HasThreadHash() bool`

HasThreadHash returns a boolean if a field has been set.

### SetThreadHashNil

`func (o *CastWithInteractionsAndConversations) SetThreadHashNil(b bool)`

 SetThreadHashNil sets the value for ThreadHash to be an explicit nil

### UnsetThreadHash
`func (o *CastWithInteractionsAndConversations) UnsetThreadHash()`

UnsetThreadHash ensures that no value is present for ThreadHash, not even an explicit nil
### GetMentionedProfiles

`func (o *CastWithInteractionsAndConversations) GetMentionedProfiles() []User`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *CastWithInteractionsAndConversations) GetMentionedProfilesOk() (*[]User, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *CastWithInteractionsAndConversations) SetMentionedProfiles(v []User)`

SetMentionedProfiles sets MentionedProfiles field to given value.

### HasMentionedProfiles

`func (o *CastWithInteractionsAndConversations) HasMentionedProfiles() bool`

HasMentionedProfiles returns a boolean if a field has been set.

### GetChannel

`func (o *CastWithInteractionsAndConversations) GetChannel() ChannelOrDehydratedChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CastWithInteractionsAndConversations) GetChannelOk() (*ChannelOrDehydratedChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CastWithInteractionsAndConversations) SetChannel(v ChannelOrDehydratedChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CastWithInteractionsAndConversations) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetViewerContext

`func (o *CastWithInteractionsAndConversations) GetViewerContext() CastViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *CastWithInteractionsAndConversations) GetViewerContextOk() (*CastViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *CastWithInteractionsAndConversations) SetViewerContext(v CastViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *CastWithInteractionsAndConversations) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.

### GetDirectReplies

`func (o *CastWithInteractionsAndConversations) GetDirectReplies() []CastWithInteractionsAndConversations`

GetDirectReplies returns the DirectReplies field if non-nil, zero value otherwise.

### GetDirectRepliesOk

`func (o *CastWithInteractionsAndConversations) GetDirectRepliesOk() (*[]CastWithInteractionsAndConversations, bool)`

GetDirectRepliesOk returns a tuple with the DirectReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectReplies

`func (o *CastWithInteractionsAndConversations) SetDirectReplies(v []CastWithInteractionsAndConversations)`

SetDirectReplies sets DirectReplies field to given value.

### HasDirectReplies

`func (o *CastWithInteractionsAndConversations) HasDirectReplies() bool`

HasDirectReplies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


