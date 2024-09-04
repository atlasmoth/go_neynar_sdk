# {{classname}}

All URIs are relative to *https://api.neynar.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllCastsInThread**](CastApi.md#AllCastsInThread) | **Get** /farcaster/all-casts-in-thread | DEPRECATED - Retrieve all casts in a given thread hash
[**Cast**](CastApi.md#Cast) | **Get** /farcaster/cast | DEPRECATED - Retrieve cast for a given hash
[**Casts**](CastApi.md#Casts) | **Get** /farcaster/casts | DEPRECATED - Retrieve casts for a given user
[**RecentCasts**](CastApi.md#RecentCasts) | **Get** /farcaster/recent-casts | Get Recent Casts

# **AllCastsInThread**
> AllCastsInThreadResponse AllCastsInThread(ctx, apiKey, threadHash, optional)
DEPRECATED - Retrieve all casts in a given thread hash

Now deprecated, use [v2/cast/conversation](https://docs.neynar.com/reference/cast-conversation). Gets all casts, including root cast and all replies for a given thread hash. No limit the depth of replies.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **threadHash** | [**string**](.md)| The hash of the thread to retrieve casts from. | 
 **optional** | ***CastApiAllCastsInThreadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiAllCastsInThreadOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 

### Return type

[**AllCastsInThreadResponse**](AllCastsInThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Cast**
> CastResponse Cast(ctx, apiKey, hash, optional)
DEPRECATED - Retrieve cast for a given hash

Now deprecated, use [v2/cast](https://docs.neynar.com/reference/cast). Gets information about an individual cast

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **hash** | [**string**](.md)| Cast hash | 
 **optional** | ***CastApiCastOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiCastOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 

### Return type

[**CastResponse**](CastResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Casts**
> CastsResponse Casts(ctx, apiKey, fid, optional)
DEPRECATED - Retrieve casts for a given user

Now deprecated, use [/v2/farcaster/feed/user/casts](https://docs.neynar.com/reference/feed-user-casts) instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of a user | 
 **optional** | ***CastApiCastsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiCastsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **parentUrl** | **optional.String**| A cast can be part of a certain channel. The channel is identified by &#x60;parent_url&#x60;. All casts in the channel ladder up to the same parent_url. | 
 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**CastsResponse**](CastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecentCasts**
> RecentCastsResponse RecentCasts(ctx, apiKey, optional)
Get Recent Casts

Get a list of casts from the protocol in reverse chronological order based on timestamp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***CastApiRecentCastsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiRecentCastsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**RecentCastsResponse**](RecentCastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

