# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuyStorage**](StorageApi.md#BuyStorage) | **Post** /farcaster/storage/buy | Buy storage for an fid
[**StorageAllocations**](StorageApi.md#StorageAllocations) | **Get** /farcaster/storage/allocations | Fetches storage allocations for a given user
[**StorageUsage**](StorageApi.md#StorageUsage) | **Get** /farcaster/storage/usage | Fetches storage usage for a given user

# **BuyStorage**
> StorageAllocationsResponse BuyStorage(ctx, body, apiKey)
Buy storage for an fid

This api will help you rent units of storage for an year for a specific fid. A storage unit lets you store 5000 casts, 2500 reactions and 2500 links. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BuyStorageReqBody**](BuyStorageReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**StorageAllocationsResponse**](StorageAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorageAllocations**
> StorageAllocationsResponse StorageAllocations(ctx, apiKey, fid)
Fetches storage allocations for a given user

Fetches storage allocations for a given user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)|  | 

### Return type

[**StorageAllocationsResponse**](StorageAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StorageUsage**
> StorageUsageResponse StorageUsage(ctx, apiKey, fid)
Fetches storage usage for a given user

Fetches storage usage for a given user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)|  | 

### Return type

[**StorageUsageResponse**](StorageUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

