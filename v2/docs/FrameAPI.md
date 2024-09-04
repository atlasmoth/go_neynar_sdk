# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNeynarFrame**](FrameApi.md#DeleteNeynarFrame) | **Delete** /farcaster/frame | Delete a frame
[**FetchNeynarFrames**](FrameApi.md#FetchNeynarFrames) | **Get** /farcaster/frame/list | Retrieve a list of frames
[**FrameFromUrl**](FrameApi.md#FrameFromUrl) | **Get** /farcaster/frame/crawl | Fetches the frame meta tags from the URL
[**LookupNeynarFrame**](FrameApi.md#LookupNeynarFrame) | **Get** /farcaster/frame | Retrieve a frame by UUID or URL
[**PostFrameAction**](FrameApi.md#PostFrameAction) | **Post** /farcaster/frame/action | Posts a frame action, cast action or a cast composer action
[**PostFrameDeveloperManagedAction**](FrameApi.md#PostFrameDeveloperManagedAction) | **Post** /farcaster/frame/developer_managed/action | Posts a frame signature packet
[**PublishNeynarFrame**](FrameApi.md#PublishNeynarFrame) | **Post** /farcaster/frame | Create a new frame
[**UpdateNeynarFrame**](FrameApi.md#UpdateNeynarFrame) | **Put** /farcaster/frame | Update an existing frame
[**ValidateFrame**](FrameApi.md#ValidateFrame) | **Post** /farcaster/frame/validate | Validates a frame action against Farcaster Hub
[**ValidateFrameAnalytics**](FrameApi.md#ValidateFrameAnalytics) | **Get** /farcaster/frame/validate/analytics | Retrieve analytics for the frame
[**ValidateFrameList**](FrameApi.md#ValidateFrameList) | **Get** /farcaster/frame/validate/list | Retrieve a list of all the frames validated by a user

# **DeleteNeynarFrame**
> DeleteFrameResponse DeleteNeynarFrame(ctx, body, apiKey)
Delete a frame

Delete an existing frame, if it was made by the developer (identified by API key)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FarcasterFrameBody**](FarcasterFrameBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**DeleteFrameResponse**](DeleteFrameResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchNeynarFrames**
> []NeynarFrame FetchNeynarFrames(ctx, apiKey)
Retrieve a list of frames

Retrieve a list of frames made by the developer (identified by API key)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**[]NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FrameFromUrl**
> InlineResponse200 FrameFromUrl(ctx, apiKey, url)
Fetches the frame meta tags from the URL

Fetches the frame meta tags from the URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **url** | **string**| The frame URL to crawl | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LookupNeynarFrame**
> NeynarFrame LookupNeynarFrame(ctx, apiKey, type_, optional)
Retrieve a frame by UUID or URL

Retrieve a frame either by UUID or Neynar URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **type_** | [**FrameType**](.md)|  | 
 **optional** | ***FrameApiLookupNeynarFrameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FrameApiLookupNeynarFrameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uuid** | [**optional.Interface of string**](.md)| UUID of the frame to retrieve | 
 **url** | **optional.String**| URL of the Neynar frame to retrieve | 

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFrameAction**
> Frame PostFrameAction(ctx, body, apiKey)
Posts a frame action, cast action or a cast composer action

Post frame actions, cast actions or cast composer actions to the server  \\ (In order to post any of these actions, you need to have an approved `signer_uuid`)  The POST request to the post_url has a timeout of 5 seconds for frames. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FrameActionReqBody**](FrameActionReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**Frame**](Frame.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFrameDeveloperManagedAction**
> Frame PostFrameDeveloperManagedAction(ctx, body, apiKey)
Posts a frame signature packet

Post a frame action that has been signed with a developer managed signer  The POST request to the post_url has a timeout of 5 seconds. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FrameDeveloperManagedActionReqBody**](FrameDeveloperManagedActionReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**Frame**](Frame.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishNeynarFrame**
> NeynarFrame PublishNeynarFrame(ctx, body, apiKey)
Create a new frame

Create a new frame with a list of pages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NeynarFrameCreationRequest**](NeynarFrameCreationRequest.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNeynarFrame**
> NeynarFrame UpdateNeynarFrame(ctx, body, apiKey)
Update an existing frame

Update an existing frame with a list of pages, if it was made by the developer (identified by API key)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NeynarFrameUpdateRequest**](NeynarFrameUpdateRequest.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateFrame**
> ValidateFrameActionResponse ValidateFrame(ctx, body, apiKey)
Validates a frame action against Farcaster Hub

Validates a frame against by an interacting user against a Farcaster Hub \\ (In order to validate a frame, message bytes from Frame Action must be provided in hex) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FrameValidateBody**](FrameValidateBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**ValidateFrameActionResponse**](ValidateFrameActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateFrameAnalytics**
> FrameValidateAnalyticsResponse ValidateFrameAnalytics(ctx, apiKey, frameUrl, analyticsType, start, stop, optional)
Retrieve analytics for the frame

Retrieve analytics for total-interactors, interactors, nteractions-per-cast and input-text.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **frameUrl** | **string**|  | 
  **analyticsType** | [**ValidateFrameAnalyticsType**](.md)|  | 
  **start** | **time.Time**|  | [default to 2024-04-06T06:44:56.811Z]
  **stop** | **time.Time**|  | [default to 2024-04-08T06:44:56.811Z]
 **optional** | ***FrameApiValidateFrameAnalyticsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FrameApiValidateFrameAnalyticsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **aggregateWindow** | **optional.String**| Required for &#x60;analytics_type&#x3D;interactions-per-cast&#x60; | 

### Return type

[**FrameValidateAnalyticsResponse**](FrameValidateAnalyticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateFrameList**
> FrameValidateListResponse ValidateFrameList(ctx, apiKey)
Retrieve a list of all the frames validated by a user

Retrieve a list of all the frames validated by a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**FrameValidateListResponse**](FrameValidateListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

