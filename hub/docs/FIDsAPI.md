# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListFids**](FIDsApi.md#ListFids) | **Get** /v1/fids | Get a list of all the FIDs

# **ListFids**
> FidsResponse ListFids(ctx, apiKey, optional)
Get a list of all the FIDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***FIDsApiListFidsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FIDsApiListFidsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**FidsResponse**](FidsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

