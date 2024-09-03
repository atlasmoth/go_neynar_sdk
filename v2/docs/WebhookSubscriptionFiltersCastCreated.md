# WebhookSubscriptionFiltersCastCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeAuthorFids** | Pointer to **[]int32** | Exclude casts that matches these authors. **Note:** This is applied as an AND operation against rest of the filters. Rest of the filters are bundled as an OR operation.  | [optional] 
**AuthorFids** | Pointer to **[]int32** |  | [optional] 
**MentionedFids** | Pointer to **[]int32** |  | [optional] 
**ParentUrls** | Pointer to **[]string** |  | [optional] 
**RootParentUrls** | Pointer to **[]string** |  | [optional] 
**ParentHashes** | Pointer to **[]string** |  | [optional] 
**ParentAuthorFids** | Pointer to **[]int32** |  | [optional] 
**Text** | Pointer to **string** | Regex pattern to match the text key of the cast. **Note:**  1) Regex must be parsed by Go&#39;s RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: (?i)\\\\$degen should be written as (?i)\\\\\\\\$degen  | [optional] 
**Embeds** | Pointer to **string** | Regex pattern to match the embeded_url (key embeds) of the cast. **Note:**  1) Regex must be parsed by Go&#39;s RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: \\\\b(farcaster|neynar)\\\\b should be written as \\\\\\\\b(farcaster|neynar)\\\\\\\\b  | [optional] 

## Methods

### NewWebhookSubscriptionFiltersCastCreated

`func NewWebhookSubscriptionFiltersCastCreated() *WebhookSubscriptionFiltersCastCreated`

NewWebhookSubscriptionFiltersCastCreated instantiates a new WebhookSubscriptionFiltersCastCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionFiltersCastCreatedWithDefaults

`func NewWebhookSubscriptionFiltersCastCreatedWithDefaults() *WebhookSubscriptionFiltersCastCreated`

NewWebhookSubscriptionFiltersCastCreatedWithDefaults instantiates a new WebhookSubscriptionFiltersCastCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) GetExcludeAuthorFids() []int32`

GetExcludeAuthorFids returns the ExcludeAuthorFids field if non-nil, zero value otherwise.

### GetExcludeAuthorFidsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetExcludeAuthorFidsOk() (*[]int32, bool)`

GetExcludeAuthorFidsOk returns a tuple with the ExcludeAuthorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) SetExcludeAuthorFids(v []int32)`

SetExcludeAuthorFids sets ExcludeAuthorFids field to given value.

### HasExcludeAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) HasExcludeAuthorFids() bool`

HasExcludeAuthorFids returns a boolean if a field has been set.

### GetAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) GetAuthorFids() []int32`

GetAuthorFids returns the AuthorFids field if non-nil, zero value otherwise.

### GetAuthorFidsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetAuthorFidsOk() (*[]int32, bool)`

GetAuthorFidsOk returns a tuple with the AuthorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) SetAuthorFids(v []int32)`

SetAuthorFids sets AuthorFids field to given value.

### HasAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) HasAuthorFids() bool`

HasAuthorFids returns a boolean if a field has been set.

### GetMentionedFids

`func (o *WebhookSubscriptionFiltersCastCreated) GetMentionedFids() []int32`

GetMentionedFids returns the MentionedFids field if non-nil, zero value otherwise.

### GetMentionedFidsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetMentionedFidsOk() (*[]int32, bool)`

GetMentionedFidsOk returns a tuple with the MentionedFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedFids

`func (o *WebhookSubscriptionFiltersCastCreated) SetMentionedFids(v []int32)`

SetMentionedFids sets MentionedFids field to given value.

### HasMentionedFids

`func (o *WebhookSubscriptionFiltersCastCreated) HasMentionedFids() bool`

HasMentionedFids returns a boolean if a field has been set.

### GetParentUrls

`func (o *WebhookSubscriptionFiltersCastCreated) GetParentUrls() []string`

GetParentUrls returns the ParentUrls field if non-nil, zero value otherwise.

### GetParentUrlsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetParentUrlsOk() (*[]string, bool)`

GetParentUrlsOk returns a tuple with the ParentUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrls

`func (o *WebhookSubscriptionFiltersCastCreated) SetParentUrls(v []string)`

SetParentUrls sets ParentUrls field to given value.

### HasParentUrls

`func (o *WebhookSubscriptionFiltersCastCreated) HasParentUrls() bool`

HasParentUrls returns a boolean if a field has been set.

### GetRootParentUrls

`func (o *WebhookSubscriptionFiltersCastCreated) GetRootParentUrls() []string`

GetRootParentUrls returns the RootParentUrls field if non-nil, zero value otherwise.

### GetRootParentUrlsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetRootParentUrlsOk() (*[]string, bool)`

GetRootParentUrlsOk returns a tuple with the RootParentUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrls

`func (o *WebhookSubscriptionFiltersCastCreated) SetRootParentUrls(v []string)`

SetRootParentUrls sets RootParentUrls field to given value.

### HasRootParentUrls

`func (o *WebhookSubscriptionFiltersCastCreated) HasRootParentUrls() bool`

HasRootParentUrls returns a boolean if a field has been set.

### GetParentHashes

`func (o *WebhookSubscriptionFiltersCastCreated) GetParentHashes() []string`

GetParentHashes returns the ParentHashes field if non-nil, zero value otherwise.

### GetParentHashesOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetParentHashesOk() (*[]string, bool)`

GetParentHashesOk returns a tuple with the ParentHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHashes

`func (o *WebhookSubscriptionFiltersCastCreated) SetParentHashes(v []string)`

SetParentHashes sets ParentHashes field to given value.

### HasParentHashes

`func (o *WebhookSubscriptionFiltersCastCreated) HasParentHashes() bool`

HasParentHashes returns a boolean if a field has been set.

### GetParentAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) GetParentAuthorFids() []int32`

GetParentAuthorFids returns the ParentAuthorFids field if non-nil, zero value otherwise.

### GetParentAuthorFidsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetParentAuthorFidsOk() (*[]int32, bool)`

GetParentAuthorFidsOk returns a tuple with the ParentAuthorFids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) SetParentAuthorFids(v []int32)`

SetParentAuthorFids sets ParentAuthorFids field to given value.

### HasParentAuthorFids

`func (o *WebhookSubscriptionFiltersCastCreated) HasParentAuthorFids() bool`

HasParentAuthorFids returns a boolean if a field has been set.

### GetText

`func (o *WebhookSubscriptionFiltersCastCreated) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WebhookSubscriptionFiltersCastCreated) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WebhookSubscriptionFiltersCastCreated) HasText() bool`

HasText returns a boolean if a field has been set.

### GetEmbeds

`func (o *WebhookSubscriptionFiltersCastCreated) GetEmbeds() string`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *WebhookSubscriptionFiltersCastCreated) GetEmbedsOk() (*string, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *WebhookSubscriptionFiltersCastCreated) SetEmbeds(v string)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *WebhookSubscriptionFiltersCastCreated) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


