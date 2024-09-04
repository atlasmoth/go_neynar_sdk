# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListVerificationsByFid**](VerificationsApi.md#ListVerificationsByFid) | **Get** /v1/verificationsByFid | Get a list of verifications provided by an FID

# **ListVerificationsByFid**
> InlineResponse2004 ListVerificationsByFid(ctx, apiKey, fid, optional)
Get a list of verifications provided by an FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID being requested | 
 **optional** | ***VerificationsApiListVerificationsByFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerificationsApiListVerificationsByFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **address** | **optional.String**| The optional ETH address to filter by | 
 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

