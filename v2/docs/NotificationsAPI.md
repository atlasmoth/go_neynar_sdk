# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkNotificationsAsSeen**](NotificationsApi.md#MarkNotificationsAsSeen) | **Post** /farcaster/notifications/seen | Mark notifications as seen
[**Notifications**](NotificationsApi.md#Notifications) | **Get** /farcaster/notifications | Retrieve notifications for a given user
[**NotificationsChannel**](NotificationsApi.md#NotificationsChannel) | **Get** /farcaster/notifications/channel | Retrieve notifications for a user in given channels
[**NotificationsParentUrl**](NotificationsApi.md#NotificationsParentUrl) | **Get** /farcaster/notifications/parent_url | Retrieve notifications for a user in given parent_urls

# **MarkNotificationsAsSeen**
> OperationResponse MarkNotificationsAsSeen(ctx, body, apiKey)
Mark notifications as seen

Mark notifications as seen

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MarkNotificationsAsSeenReqBody**](MarkNotificationsAsSeenReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Notifications**
> NotificationsResponse Notifications(ctx, apiKey, fid, optional)
Retrieve notifications for a given user

Returns a list of notifications for a specific FID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user you you want to fetch notifications for | 
 **optional** | ***NotificationsApiNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiNotificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**optional.Interface of NotificationType**](.md)| Notification type to fetch. | 
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationsChannel**
> NotificationsResponse NotificationsChannel(ctx, apiKey, fid, channelIds, optional)
Retrieve notifications for a user in given channels

Returns a list of notifications for a user in specific channels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user you you want to fetch notifications for | 
  **channelIds** | **string**| Comma separated channel_ids (find list of all channels here - https://docs.neynar.com/reference/list-all-channels) | 
 **optional** | ***NotificationsApiNotificationsChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiNotificationsChannelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotificationsParentUrl**
> NotificationsResponse NotificationsParentUrl(ctx, apiKey, fid, parentUrls, optional)
Retrieve notifications for a user in given parent_urls

Returns a list of notifications for a user in specific parent_urls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user you you want to fetch notifications for | 
  **parentUrls** | **string**| Comma separated parent_urls | 
 **optional** | ***NotificationsApiNotificationsParentUrlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NotificationsApiNotificationsParentUrlOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**NotificationsResponse**](NotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

