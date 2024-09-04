# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveUsers**](UserApi.md#ActiveUsers) | **Get** /farcaster/user/active | Fetch active users
[**FarcasterUserVerificationDelete**](UserApi.md#FarcasterUserVerificationDelete) | **Delete** /farcaster/user/verification | Removes verification for an eth address for the user
[**FarcasterUserVerificationPost**](UserApi.md#FarcasterUserVerificationPost) | **Post** /farcaster/user/verification | Adds verification for an ethereum address or contract for the user
[**FollowUser**](UserApi.md#FollowUser) | **Post** /farcaster/user/follow | Follow a user
[**GetFreshFid**](UserApi.md#GetFreshFid) | **Get** /farcaster/user/fid | Fetches fid to assign it new user
[**LookupUserByCustodyAddress**](UserApi.md#LookupUserByCustodyAddress) | **Get** /farcaster/user/custody-address | Lookup a user by custody-address
[**PowerUsers**](UserApi.md#PowerUsers) | **Get** /farcaster/user/power | Fetch power user objects
[**RegisterUser**](UserApi.md#RegisterUser) | **Post** /farcaster/user | Register account on farcaster
[**UnfollowUser**](UserApi.md#UnfollowUser) | **Delete** /farcaster/user/follow | Unfollow a user
[**UpdateUser**](UserApi.md#UpdateUser) | **Patch** /farcaster/user | Update user profile
[**UserBulk**](UserApi.md#UserBulk) | **Get** /farcaster/user/bulk | Fetch users based on FIDs
[**UserBulkByAddress**](UserApi.md#UserBulkByAddress) | **Get** /farcaster/user/bulk-by-address | Fetches users based on Eth or Sol addresses
[**UserPowerLite**](UserApi.md#UserPowerLite) | **Get** /farcaster/user/power_lite | Fetch power user FIDs
[**UserSearch**](UserApi.md#UserSearch) | **Get** /farcaster/user/search | Search for Usernames

# **ActiveUsers**
> UsersResponse ActiveUsers(ctx, apiKey, optional)
Fetch active users

Warpcast has deprecated the active badge. Use user/power endpoint instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***UserApiActiveUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiActiveUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**|  | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FarcasterUserVerificationDelete**
> OperationResponse FarcasterUserVerificationDelete(ctx, body, apiKey)
Removes verification for an eth address for the user

Removes verification for an eth address for the user \\ (In order to delete verification `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RemoveVerificationReqBody**](RemoveVerificationReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FarcasterUserVerificationPost**
> OperationResponse FarcasterUserVerificationPost(ctx, body, apiKey)
Adds verification for an ethereum address or contract for the user

Adds verification for an eth address or contract for the user \\ (In order to add verification `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddVerificationReqBody**](AddVerificationReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FollowUser**
> BulkFollowResponse FollowUser(ctx, body, apiKey)
Follow a user

Follow a user \\ (In order to follow a user `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FollowReqBody**](FollowReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFreshFid**
> UserFidResponse GetFreshFid(ctx, apiKey)
Fetches fid to assign it new user

Fetches fid to assign it new user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**UserFidResponse**](UserFIDResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupUserByCustodyAddress**
> UserResponse LookupUserByCustodyAddress(ctx, apiKey, custodyAddress)
Lookup a user by custody-address

Lookup a user by custody-address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **custodyAddress** | **string**| Custody Address associated with mnemonic | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PowerUsers**
> UsersResponse PowerUsers(ctx, apiKey, optional)
Fetch power user objects

Fetches power users based on Warpcast power badges. Information is updated once a day.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***UserApiPowerUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiPowerUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **limit** | **optional.Int32**| Number of power users to fetch, max 100 | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterUser**
> RegisterUserResponse RegisterUser(ctx, body, apiKey)
Register account on farcaster

Register account on farcaster.  **Note:** This API must be called within 10 minutes of the fetch fid API call (i.e., /v2/farcaster/user/fid). Otherwise, Neynar will assign this fid to another available user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterUserReqBody**](RegisterUserReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**RegisterUserResponse**](RegisterUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnfollowUser**
> BulkFollowResponse UnfollowUser(ctx, body, apiKey)
Unfollow a user

Unfollow a user \\ (In order to unfollow a user `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FollowReqBody**](FollowReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUser**
> OperationResponse UpdateUser(ctx, body, apiKey)
Update user profile

Update user profile \\ (In order to update user's profile `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserReqBody**](UpdateUserReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserBulk**
> BulkUsersResponse UserBulk(ctx, apiKey, fids, optional)
Fetch users based on FIDs

Fetches information about multiple users based on FIDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fids** | **string**| Comma separated list of FIDs, up to 100 at a time | 
 **optional** | ***UserApiUserBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserBulkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**BulkUsersResponse**](BulkUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserBulkByAddress**
> map[string][]User UserBulkByAddress(ctx, apiKey, addresses, optional)
Fetches users based on Eth or Sol addresses

Fetches all users based on multiple Ethereum or Solana addresses.  Each farcaster user has a custody Ethereum address and optionally verified Ethereum or Solana addresses. This endpoint returns all users that have any of the given addresses as their custody or verified Ethereum or Solana addresses.  A custody address can be associated with only 1 farcaster user at a time but a verified address can be associated with multiple users. You can pass in Ethereum and Solana addresses, comma separated, in the same request. The response will contain users associated with the given addresses.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **addresses** | **string**| Comma separated list of Ethereum addresses, up to 350 at a time | 
 **optional** | ***UserApiUserBulkByAddressOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserBulkByAddressOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addressTypes** | **optional.String**| Customize which address types the request should search for. This is a comma-separated string that can include the following values: &#x27;custody_address&#x27; and &#x27;verified_address&#x27;. By default api returns both. To select multiple types, use a comma-separated list of these values.  | 
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**map[string][]User**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPowerLite**
> UserPowerLiteResponse UserPowerLite(ctx, apiKey)
Fetch power user FIDs

Fetches power users and respond in a backwards compatible format to Warpcast's deprecated power badge endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**UserPowerLiteResponse**](UserPowerLiteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSearch**
> UserSearchResponse UserSearch(ctx, apiKey, q, optional)
Search for Usernames

Search for Usernames

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **q** | **string**|  | 
 **optional** | ***UserApiUserSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserApiUserSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **limit** | **optional.Int32**|  | [default to 5]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

