# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockList**](BlockApi.md#BlockList) | **Get** /farcaster/block/list | Get fids that a user has blocked or has been blocked by

# **BlockList**
> BlockListResponse BlockList(ctx, apiKey, optional)
Get fids that a user has blocked or has been blocked by

Fetches all fids that a user has blocked or has been blocked by

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***BlockApiBlockListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BlockApiBlockListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blockerFid** | [**optional.Interface of int32**](.md)| Providing this will return the users that this user has blocked | 
 **blockedFid** | [**optional.Interface of int32**](.md)| Providing this will return the users that have blocked this user | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 20, max 100). | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**BlockListResponse**](BlockListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

