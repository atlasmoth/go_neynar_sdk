# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMute**](MuteApi.md#AddMute) | **Post** /farcaster/mute | Adds a mute for a fid
[**DeleteMute**](MuteApi.md#DeleteMute) | **Delete** /farcaster/mute | Deletes a mute for a fid
[**MuteList**](MuteApi.md#MuteList) | **Get** /farcaster/mute/list | Get fids that a user has muted

# **AddMute**
> MuteResponse AddMute(ctx, body, apiKey)
Adds a mute for a fid

Adds a mute for a given fid. This is a whitelisted API, reach out if you want access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MuteReqBody**](MuteReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**MuteResponse**](MuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMute**
> MuteResponse DeleteMute(ctx, body, apiKey)
Deletes a mute for a fid

Deletes a mute for a given fid. This is a whitelisted API, reach out if you want access.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MuteReqBody**](MuteReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**MuteResponse**](MuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteList**
> MuteListResponse MuteList(ctx, apiKey, fid, optional)
Get fids that a user has muted

Fetches all fids that a user has muted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| The user&#x27;s fid (identifier) | 
 **optional** | ***MuteApiMuteListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuteApiMuteListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Number of results to retrieve (default 20, max 100). | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**MuteListResponse**](MuteListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

