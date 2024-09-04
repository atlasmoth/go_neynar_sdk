# WebhookSubscriptionFiltersCastCreated

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeAuthorFids** | **[]int32** | Exclude casts that matches these authors. **Note:** This is applied as an AND operation against rest of the filters. Rest of the filters are bundled as an OR operation.  | [optional] [default to null]
**AuthorFids** | **[]int32** |  | [optional] [default to null]
**MentionedFids** | **[]int32** |  | [optional] [default to null]
**ParentUrls** | **[]string** |  | [optional] [default to null]
**RootParentUrls** | **[]string** |  | [optional] [default to null]
**ParentHashes** | **[]string** |  | [optional] [default to null]
**ParentAuthorFids** | **[]int32** |  | [optional] [default to null]
**Text** | **string** | Regex pattern to match the text key of the cast. **Note:**  1) Regex must be parsed by Go&#x27;s RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: (?i)\\\\$degen should be written as (?i)\\\\\\\\$degen  | [optional] [default to null]
**Embeds** | **string** | Regex pattern to match the embeded_url (key embeds) of the cast. **Note:**  1) Regex must be parsed by Go&#x27;s RE2 engine (Test your expression here: https://www.lddgo.net/en/string/golangregex) 2) Use backslashes to escape special characters. For example: \\\\b(farcaster|neynar)\\\\b should be written as \\\\\\\\b(farcaster|neynar)\\\\\\\\b  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

