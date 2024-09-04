# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionCheck**](STPApi.md#SubscriptionCheck) | **Get** /stp/subscription_check | Check if a wallet address is subscribed to a STP contract

# **SubscriptionCheck**
> map[string]SubscriptionStatus SubscriptionCheck(ctx, apiKey, addresses, contractAddress, chainId)
Check if a wallet address is subscribed to a STP contract

Check if a wallet address is subscribed to a given STP contract.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **addresses** | **string**| Comma separated list of Ethereum addresses, up to 350 at a time | 
  **contractAddress** | **string**| Ethereum address of the STP contract | 
  **chainId** | **string**| Chain ID of the STP contract | 

### Return type

[**map[string]SubscriptionStatus**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

