# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStorageLimitsByFid**](StorageApi.md#GetStorageLimitsByFid) | **Get** /v1/storageLimitsByFid | Get an FID&#x27;s storage limits.

# **GetStorageLimitsByFid**
> StorageLimitsResponse GetStorageLimitsByFid(ctx, apiKey, fid)
Get an FID's storage limits.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**|  | 

### Return type

[**StorageLimitsResponse**](StorageLimitsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

