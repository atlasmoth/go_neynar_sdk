# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLinkById**](LinksApi.md#GetLinkById) | **Get** /v1/linkById | Get a link by its FID and target FID.
[**ListLinksByFid**](LinksApi.md#ListLinksByFid) | **Get** /v1/linksByFid | Get all links from a source FID
[**ListLinksByTargetFid**](LinksApi.md#ListLinksByTargetFid) | **Get** /v1/linksByTargetFid | Get all links to a target FID

# **GetLinkById**
> LinkAdd GetLinkById(ctx, apiKey, fid, targetFid, linkType)
Get a link by its FID and target FID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID of the link&#x27;s originator | 
  **targetFid** | **int32**| The FID of the target of the link | 
  **linkType** | [**LinkType**](.md)| The type of link, as a string value | 

### Return type

[**LinkAdd**](LinkAdd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLinksByFid**
> InlineResponse2002 ListLinksByFid(ctx, apiKey, fid, optional)
Get all links from a source FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID of the link&#x27;s originator | 
 **optional** | ***LinksApiListLinksByFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinksApiListLinksByFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkType** | [**optional.Interface of LinkType**](.md)| The type of link, as a string value | 
 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLinksByTargetFid**
> InlineResponse2002 ListLinksByTargetFid(ctx, apiKey, targetFid, optional)
Get all links to a target FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **targetFid** | **int32**| The FID of the target of the link | 
 **optional** | ***LinksApiListLinksByTargetFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LinksApiListLinksByTargetFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linkType** | [**optional.Interface of LinkType**](.md)| The type of link, as a string value | 
 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

