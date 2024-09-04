# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscribedTo**](SubscribersApi.md#SubscribedTo) | **Get** /farcaster/user/subscribed_to | Fetch what a given fid is subscribed to
[**Subscribers**](SubscribersApi.md#Subscribers) | **Get** /farcaster/user/subscribers | Fetch subscribers for a given fid
[**SubscriptionsCreated**](SubscribersApi.md#SubscriptionsCreated) | **Get** /farcaster/user/subscriptions_created | Fetch created subscriptions for a given fid

# **SubscribedTo**
> SubscribedToResponse SubscribedTo(ctx, apiKey, fid, subscriptionProvider, optional)
Fetch what a given fid is subscribed to

Fetch what fids and contracts a fid is subscribed to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)|  | 
  **subscriptionProvider** | [**SubscriptionProvider**](.md)|  | 
 **optional** | ***SubscribersApiSubscribedToOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscribersApiSubscribedToOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**SubscribedToResponse**](SubscribedToResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Subscribers**
> SubscribersResponse Subscribers(ctx, apiKey, fid, subscriptionProvider, optional)
Fetch subscribers for a given fid

Fetch subscribers for a given fid's contracts. Doesn't return addresses that don't have an fid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)|  | 
  **subscriptionProvider** | [**SubscriptionProviders**](.md)|  | 
 **optional** | ***SubscribersApiSubscribersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubscribersApiSubscribersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**SubscribersResponse**](SubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubscriptionsCreated**
> SubscriptionsResponse SubscriptionsCreated(ctx, apiKey, fid, subscriptionProvider)
Fetch created subscriptions for a given fid

Fetch created subscriptions for a given fid's.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)|  | 
  **subscriptionProvider** | [**SubscriptionProvider**](.md)|  | 

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

