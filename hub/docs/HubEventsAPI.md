# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventById**](HubEventsApi.md#GetEventById) | **Get** /v1/eventById | Get an event by its ID
[**ListEvents**](HubEventsApi.md#ListEvents) | **Get** /v1/events | Get a page of Hub events

# **GetEventById**
> HubEvent GetEventById(ctx, apiKey, eventId)
Get an event by its ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **eventId** | **int32**| The Hub Id of the event | 

### Return type

[**HubEvent**](HubEvent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEvents**
> InlineResponse2007 ListEvents(ctx, apiKey, optional)
Get a page of Hub events

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***HubEventsApiListEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HubEventsApiListEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromEventId** | **optional.Int32**| An optional Hub Id to start getting events from.  This is also returned from the API as nextPageEventId, which  can be used to page through all the Hub events. Set it to 0  to start from the first event | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

