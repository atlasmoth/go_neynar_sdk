# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReactionById**](ReactionsApi.md#GetReactionById) | **Get** /v1/reactionById | Get a reaction by its created FID and target Cast.
[**ListReactionsByCast**](ReactionsApi.md#ListReactionsByCast) | **Get** /v1/reactionsByCast | Get all reactions to a cast
[**ListReactionsByFid**](ReactionsApi.md#ListReactionsByFid) | **Get** /v1/reactionsByFid | Get all reactions by an FID
[**ListReactionsByTarget**](ReactionsApi.md#ListReactionsByTarget) | **Get** /v1/reactionsByTarget | Get all reactions to a target URL

# **GetReactionById**
> Reaction GetReactionById(ctx, apiKey, fid, targetFid, targetHash, reactionType)
Get a reaction by its created FID and target Cast.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID of the reaction&#x27;s creator | 
  **targetFid** | **int32**| The FID of the cast&#x27;s creator | 
  **targetHash** | **string**| The cast&#x27;s hash | 
  **reactionType** | [**ReactionType**](.md)| The type of reaction, either as a numerical enum value or string representation | 

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReactionsByCast**
> InlineResponse2001 ListReactionsByCast(ctx, apiKey, targetFid, targetHash, reactionType, optional)
Get all reactions to a cast

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **targetFid** | **int32**| The FID of the cast&#x27;s creator | 
  **targetHash** | **string**| The hash of the cast | 
  **reactionType** | [**ReactionType**](.md)| The type of reaction, either as a numerical enum value or string representation | 
 **optional** | ***ReactionsApiListReactionsByCastOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsApiListReactionsByCastOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReactionsByFid**
> InlineResponse2001 ListReactionsByFid(ctx, apiKey, fid, reactionType, optional)
Get all reactions by an FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID of the reaction&#x27;s creator | 
  **reactionType** | [**ReactionType**](.md)| The type of reaction, either as a numerical enum value or string representation | 
 **optional** | ***ReactionsApiListReactionsByFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsApiListReactionsByFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListReactionsByTarget**
> InlineResponse2001 ListReactionsByTarget(ctx, apiKey, url, reactionType, optional)
Get all reactions to a target URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **url** | **string**| The URL of the parent cast | 
  **reactionType** | [**ReactionType**](.md)| The type of reaction, either as a numerical enum value or string representation | 
 **optional** | ***ReactionsApiListReactionsByTargetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsApiListReactionsByTargetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

