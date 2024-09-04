# {{classname}}

All URIs are relative to *https://api.neynar.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MentionsAndReplies**](NotificationsApi.md#MentionsAndReplies) | **Get** /farcaster/mentions-and-replies | Get mentions and replies
[**ReactionsAndRecasts**](NotificationsApi.md#ReactionsAndRecasts) | **Get** /farcaster/reactions-and-recasts | Get reactions and recasts

# **MentionsAndReplies**
> MentionsAndRepliesResponse MentionsAndReplies(ctx, apiKey, fid, optional)
Get mentions and replies

Gets a list of 15 mentions and replies to the user’s casts in reverse chronological order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of a user | 
 **optional** | ***NotificationsApiMentionsAndRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiMentionsAndRepliesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**MentionsAndRepliesResponse**](MentionsAndRepliesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsAndRecasts**
> ReactionsAndRecastsResponse ReactionsAndRecasts(ctx, apiKey, fid, optional)
Get reactions and recasts

Get a list of reactions and recasts to the users’s casts in reverse chronological order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of a user | 
 **optional** | ***NotificationsApiReactionsAndRecastsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiReactionsAndRecastsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**ReactionsAndRecastsResponse**](ReactionsAndRecastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

