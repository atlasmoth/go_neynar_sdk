# {{classname}}

All URIs are relative to *https://api.neynar.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CastLikes**](ReactionsApi.md#CastLikes) | **Get** /farcaster/cast-likes | DEPRECATED - Get all like reactions for a specific cast
[**CastReactions**](ReactionsApi.md#CastReactions) | **Get** /farcaster/cast-reactions | DEPRECATED - Get all reactions for a specific cast
[**CastRecasters**](ReactionsApi.md#CastRecasters) | **Get** /farcaster/cast-recasters | DEPRECATED - Get all recasters for a specific cast

# **CastLikes**
> CastLikesResponse CastLikes(ctx, apiKey, castHash, optional)
DEPRECATED - Get all like reactions for a specific cast

Now deprecated, use [/v2/farcaster/reactions/cast](https://docs.neynar.com/reference/reactions-cast) - Get all like reactions for a specific cast in reverse chronological order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **castHash** | [**string**](.md)| Cast hash | 
 **optional** | ***ReactionsApiCastLikesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsApiCastLikesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**CastLikesResponse**](CastLikesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CastReactions**
> CastReactionsResponse CastReactions(ctx, apiKey, castHash, optional)
DEPRECATED - Get all reactions for a specific cast

Now deprecated, use [/v2/farcaster/reactions/cast](https://docs.neynar.com/reference/reactions-cast) - Get all reactions (likes and recasts) for a specific cast.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **castHash** | [**string**](.md)| Cast hash | 
 **optional** | ***ReactionsApiCastReactionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsApiCastReactionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**CastReactionsResponse**](CastReactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CastRecasters**
> CastRecasterResponse CastRecasters(ctx, apiKey, castHash, optional)
DEPRECATED - Get all recasters for a specific cast

Now deprecated, use [/v2/farcaster/reactions/cast](https://docs.neynar.com/reference/reactions-cast) - Get all recasters for a specific cast.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **castHash** | [**string**](.md)| Cast hash | 
 **optional** | ***ReactionsApiCastRecastersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsApiCastRecastersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**CastRecasterResponse**](CastRecasterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

