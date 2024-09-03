# ReactionsAndRecastsNotification

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
**Reactors** | Pointer to [**[]User**](User.md) |  | [optional] 
**ReactionType** | Pointer to [**ReactionType**](ReactionType.md) |  | [optional] 

## Methods

### NewReactionsAndRecastsNotification

`func NewReactionsAndRecastsNotification(hash string, parentHash NullableString, parentUrl NullableString, threadHash string, parentAuthor CastParentAuthor, mentionedProfiles []User, author CastAuthor, text string, timestamp time.Time, embeds []EmbedUrl, ) *ReactionsAndRecastsNotification`

NewReactionsAndRecastsNotification instantiates a new ReactionsAndRecastsNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionsAndRecastsNotificationWithDefaults

`func NewReactionsAndRecastsNotificationWithDefaults() *ReactionsAndRecastsNotification`

NewReactionsAndRecastsNotificationWithDefaults instantiates a new ReactionsAndRecastsNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *ReactionsAndRecastsNotification) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ReactionsAndRecastsNotification) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ReactionsAndRecastsNotification) SetHash(v string)`

SetHash sets Hash field to given value.


### GetParentHash

`func (o *ReactionsAndRecastsNotification) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *ReactionsAndRecastsNotification) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *ReactionsAndRecastsNotification) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.


### SetParentHashNil

`func (o *ReactionsAndRecastsNotification) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *ReactionsAndRecastsNotification) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *ReactionsAndRecastsNotification) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *ReactionsAndRecastsNotification) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *ReactionsAndRecastsNotification) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.


### SetParentUrlNil

`func (o *ReactionsAndRecastsNotification) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *ReactionsAndRecastsNotification) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetThreadHash

`func (o *ReactionsAndRecastsNotification) GetThreadHash() string`

GetThreadHash returns the ThreadHash field if non-nil, zero value otherwise.

### GetThreadHashOk

`func (o *ReactionsAndRecastsNotification) GetThreadHashOk() (*string, bool)`

GetThreadHashOk returns a tuple with the ThreadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadHash

`func (o *ReactionsAndRecastsNotification) SetThreadHash(v string)`

SetThreadHash sets ThreadHash field to given value.


### GetParentAuthor

`func (o *ReactionsAndRecastsNotification) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *ReactionsAndRecastsNotification) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *ReactionsAndRecastsNotification) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.


### GetMentionedProfiles

`func (o *ReactionsAndRecastsNotification) GetMentionedProfiles() []User`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *ReactionsAndRecastsNotification) GetMentionedProfilesOk() (*[]User, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *ReactionsAndRecastsNotification) SetMentionedProfiles(v []User)`

SetMentionedProfiles sets MentionedProfiles field to given value.


### GetAuthor

`func (o *ReactionsAndRecastsNotification) GetAuthor() CastAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ReactionsAndRecastsNotification) GetAuthorOk() (*CastAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ReactionsAndRecastsNotification) SetAuthor(v CastAuthor)`

SetAuthor sets Author field to given value.


### GetText

`func (o *ReactionsAndRecastsNotification) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReactionsAndRecastsNotification) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReactionsAndRecastsNotification) SetText(v string)`

SetText sets Text field to given value.


### GetTimestamp

`func (o *ReactionsAndRecastsNotification) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ReactionsAndRecastsNotification) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ReactionsAndRecastsNotification) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetEmbeds

`func (o *ReactionsAndRecastsNotification) GetEmbeds() []EmbedUrl`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *ReactionsAndRecastsNotification) GetEmbedsOk() (*[]EmbedUrl, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *ReactionsAndRecastsNotification) SetEmbeds(v []EmbedUrl)`

SetEmbeds sets Embeds field to given value.


### GetType

`func (o *ReactionsAndRecastsNotification) GetType() CastType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReactionsAndRecastsNotification) GetTypeOk() (*CastType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReactionsAndRecastsNotification) SetType(v CastType)`

SetType sets Type field to given value.

### HasType

`func (o *ReactionsAndRecastsNotification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetReactors

`func (o *ReactionsAndRecastsNotification) GetReactors() []User`

GetReactors returns the Reactors field if non-nil, zero value otherwise.

### GetReactorsOk

`func (o *ReactionsAndRecastsNotification) GetReactorsOk() (*[]User, bool)`

GetReactorsOk returns a tuple with the Reactors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactors

`func (o *ReactionsAndRecastsNotification) SetReactors(v []User)`

SetReactors sets Reactors field to given value.

### HasReactors

`func (o *ReactionsAndRecastsNotification) HasReactors() bool`

HasReactors returns a boolean if a field has been set.

### GetReactionType

`func (o *ReactionsAndRecastsNotification) GetReactionType() ReactionType`

GetReactionType returns the ReactionType field if non-nil, zero value otherwise.

### GetReactionTypeOk

`func (o *ReactionsAndRecastsNotification) GetReactionTypeOk() (*ReactionType, bool)`

GetReactionTypeOk returns a tuple with the ReactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionType

`func (o *ReactionsAndRecastsNotification) SetReactionType(v ReactionType)`

SetReactionType sets ReactionType field to given value.

### HasReactionType

`func (o *ReactionsAndRecastsNotification) HasReactionType() bool`

HasReactionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


