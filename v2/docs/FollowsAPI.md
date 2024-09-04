# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FollowersV2**](FollowsApi.md#FollowersV2) | **Get** /farcaster/followers | Retrieve followers for a given user
[**FollowingV2**](FollowsApi.md#FollowingV2) | **Get** /farcaster/following | Retrieve a list of users followed by a user
[**RelevantFollowers**](FollowsApi.md#RelevantFollowers) | **Get** /farcaster/followers/relevant | Retrieve relevant followers for a given user

# **FollowersV2**
> UsersResponse FollowersV2(ctx, apiKey, fid, optional)
Retrieve followers for a given user

Returns a list of followers for a specific FID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| User who&#x27;s profile you are looking at | 
 **optional** | ***FollowsApiFollowersV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FollowsApiFollowersV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| Viewer who&#x27;s looking at the profile. | 
 **sortType** | [**optional.Interface of FollowSortType**](.md)| Sort type for retrieve followers. Default is &#x60;desc_chron&#x60; | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 20, max 100) | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FollowingV2**
> UsersResponse FollowingV2(ctx, apiKey, fid, optional)
Retrieve a list of users followed by a user

Retrieve a list of users followed by a user. Can optionally include a viewer_fid and sort_type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user whose following you want to fetch. | 
 **optional** | ***FollowsApiFollowingV2Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FollowsApiFollowingV2Opts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| FID of the user viewing the user. | 
 **sortType** | [**optional.Interface of FollowSortType**](.md)| Optional parameter to sort the users based on different criteria. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RelevantFollowers**
> RelevantFollowersResponse RelevantFollowers(ctx, apiKey, targetFid, viewerFid)
Retrieve relevant followers for a given user

Returns a list of relevant followers for a specific FID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **targetFid** | [**int32**](.md)| User who&#x27;s profile you are looking at | 
  **viewerFid** | [**int32**](.md)| Viewer who&#x27;s looking at the profile | 

### Return type

[**RelevantFollowersResponse**](RelevantFollowersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

