# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOnChainIdRegistrationByAddress**](OnChainEventsApi.md#GetOnChainIdRegistrationByAddress) | **Get** /v1/onChainIdRegistryEventByAddress | Get an on chain ID Registry Event for a given Address
[**ListOnChainEventsByFid**](OnChainEventsApi.md#ListOnChainEventsByFid) | **Get** /v1/onChainEventsByFid | Get a list of on-chain events provided by an FID
[**ListOnChainSignersByFid**](OnChainEventsApi.md#ListOnChainSignersByFid) | **Get** /v1/onChainSignersByFid | Get a list of signers provided by an FID

# **GetOnChainIdRegistrationByAddress**
> OnChainEventIdRegister GetOnChainIdRegistrationByAddress(ctx, apiKey, address)
Get an on chain ID Registry Event for a given Address

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **address** | **string**| The ETH address being requested | 

### Return type

[**OnChainEventIdRegister**](OnChainEventIdRegister.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOnChainEventsByFid**
> InlineResponse2005 ListOnChainEventsByFid(ctx, apiKey, fid, eventType)
Get a list of on-chain events provided by an FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID being requested | 
  **eventType** | [**OnChainEventType**](.md)| The numeric of string value of the event type being requested. | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOnChainSignersByFid**
> InlineResponse2006 ListOnChainSignersByFid(ctx, apiKey, fid, optional)
Get a list of signers provided by an FID

**Note:** one of two different response schemas is returned  based on whether the caller provides the `signer` parameter. If included, a single `OnChainEventSigner` message is returned (or a `not_found` error). If omitted, a  non-paginated list of `OnChainEventSigner` messages is returned instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID being requested | 
 **optional** | ***OnChainEventsApiListOnChainSignersByFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OnChainEventsApiListOnChainSignersByFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **signer** | **optional.String**| The optional key of signer | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

