# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FnameAvailability**](FnameApi.md#FnameAvailability) | **Get** /farcaster/fname/availability | Check if a given fname is available

# **FnameAvailability**
> FnameAvailabilityResponse FnameAvailability(ctx, apiKey, fname)
Check if a given fname is available

Check if a given fname is available

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fname** | **string**|  | 

### Return type

[**FnameAvailabilityResponse**](FnameAvailabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

