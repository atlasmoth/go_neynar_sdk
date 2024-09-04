# {{classname}}

All URIs are relative to *https://api.neynar.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Followers**](FollowsApi.md#Followers) | **Get** /farcaster/followers | Gets all followers for a given FID
[**Following**](FollowsApi.md#Following) | **Get** /farcaster/following | Gets all following users of a FID

# **Followers**
> FollowResponse Followers(ctx, apiKey, fid, optional)
Gets all followers for a given FID

Gets a list of users who follow a given user in reverse chronological order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user | 
 **optional** | ***FollowsApiFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FollowsApiFollowersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**FollowResponse**](FollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Following**
> FollowResponse Following(ctx, apiKey, fid, optional)
Gets all following users of a FID

Gets a list of users who is following a given user in reverse chronological order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user | 
 **optional** | ***FollowsApiFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FollowsApiFollowingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**FollowResponse**](FollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

