# {{classname}}

All URIs are relative to *https://api.neynar.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustodyAddress**](UserApi.md#CustodyAddress) | **Get** /farcaster/custody-address | DEPRECATED - Get the custody address for a given FID
[**RecentUsers**](UserApi.md#RecentUsers) | **Get** /farcaster/recent-users | Get Recent Users
[**User**](UserApi.md#User) | **Get** /farcaster/user | DEPRECATED - Get User Information by FID
[**UserByUsername**](UserApi.md#UserByUsername) | **Get** /farcaster/user-by-username | Get User Information by username
[**UserCastLikes**](UserApi.md#UserCastLikes) | **Get** /farcaster/user-cast-likes | DEPRECATED -- Get User Cast Likes

# **CustodyAddress**
> CustodyAddressResponse CustodyAddress(ctx, apiKey, fid)
DEPRECATED - Get the custody address for a given FID

Now deprecated, use [v2/user/bulk](https://docs.neynar.com/reference/user-bulk), find custody address in user obj. Returns the custody address for a given FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of a user | 

### Return type

[**CustodyAddressResponse**](CustodyAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecentUsers**
> RecentUsersResponse RecentUsers(ctx, apiKey, optional)
Get Recent Users

Get a list of casts from the protocol in reverse chronological order based on timestamp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***UserApiRecentUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiRecentUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 100, max 1000) | [default to 100]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**RecentUsersResponse**](RecentUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **User**
> UserResponse User(ctx, apiKey, fid, optional)
DEPRECATED - Get User Information by FID

Now deprecated, use [v2/user/bulk](https://docs.neynar.com/reference/user-bulk). Returns metadata about a specific user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of a user | 
 **optional** | ***UserApiUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserByUsername**
> UserResponse UserByUsername(ctx, apiKey, username, optional)
Get User Information by username

Returns metadata about a specific user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **username** | **string**| Username of the user | [default to shreyas-chorge]
 **optional** | ***UserApiUserByUsernameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserByUsernameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserCastLikes**
> UserCastLikeResponse UserCastLikes(ctx, apiKey, fid, optional)
DEPRECATED -- Get User Cast Likes

Now deprecated. use [/v2/reactions/user](https://docs.neynar.com/reference/reactions-user). Fetch all the liked cast of a User

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user | 
 **optional** | ***UserApiUserCastLikesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserCastLikesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| fid of the user viewing this information, needed for contextual information. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 150) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor | 

### Return type

[**UserCastLikeResponse**](UserCastLikeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

