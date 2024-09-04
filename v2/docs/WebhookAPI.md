# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhook**](WebhookApi.md#DeleteWebhook) | **Delete** /farcaster/webhook | Delete a webhook
[**FetchWebhooks**](WebhookApi.md#FetchWebhooks) | **Get** /farcaster/webhook/list | Fetch a list of webhooks associated to a user
[**LookupWebhook**](WebhookApi.md#LookupWebhook) | **Get** /farcaster/webhook | Fetch a webhook
[**PublishWebhook**](WebhookApi.md#PublishWebhook) | **Post** /farcaster/webhook | Create a webhook
[**UpdateWebhook**](WebhookApi.md#UpdateWebhook) | **Put** /farcaster/webhook | Update a webhook
[**UpdateWebhookActiveStatus**](WebhookApi.md#UpdateWebhookActiveStatus) | **Patch** /farcaster/webhook | Update webhook active status

# **DeleteWebhook**
> WebhookResponse DeleteWebhook(ctx, body, apiKey)
Delete a webhook

Delete a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookDeleteReqBody**](WebhookDeleteReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchWebhooks**
> WebhookListResponse FetchWebhooks(ctx, apiKey)
Fetch a list of webhooks associated to a user

Fetch a list of webhooks associated to a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**WebhookListResponse**](WebhookListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupWebhook**
> WebhookResponse LookupWebhook(ctx, apiKey, webhookId)
Fetch a webhook

Fetch a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **webhookId** | **string**|  | 

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishWebhook**
> WebhookResponse PublishWebhook(ctx, body, apiKey)
Create a webhook

Create a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookPostReqBody**](WebhookPostReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhook**
> WebhookResponse UpdateWebhook(ctx, body, apiKey)
Update a webhook

Update a webhook

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookPutReqBody**](WebhookPutReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWebhookActiveStatus**
> WebhookResponse UpdateWebhookActiveStatus(ctx, body, apiKey)
Update webhook active status

Update webhook active status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookPatchReqBody**](WebhookPatchReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

