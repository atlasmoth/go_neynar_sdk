# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCastById**](CastsApi.md#GetCastById) | **Get** /v1/castById | Get a cast by its FID and Hash.
[**ListCastsByFid**](CastsApi.md#ListCastsByFid) | **Get** /v1/castsByFid | Fetch all casts authored by an FID.
[**ListCastsByMention**](CastsApi.md#ListCastsByMention) | **Get** /v1/castsByMention | Fetch all casts that mention an FID
[**ListCastsByParent**](CastsApi.md#ListCastsByParent) | **Get** /v1/castsByParent | Fetch all casts by parent cast&#x27;s FID and Hash OR by the parent&#x27;s URL

# **GetCastById**
> CastAdd GetCastById(ctx, apiKey, fid, hash)
Get a cast by its FID and Hash.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID of the cast&#x27;s creator | 
  **hash** | **string**| The cast&#x27;s hash | 

### Return type

[**CastAdd**](CastAdd.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCastsByFid**
> InlineResponse200 ListCastsByFid(ctx, apiKey, fid, optional)
Fetch all casts authored by an FID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID of the casts&#x27; creator | 
 **optional** | ***CastsApiListCastsByFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastsApiListCastsByFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCastsByMention**
> InlineResponse200 ListCastsByMention(ctx, apiKey, fid, optional)
Fetch all casts that mention an FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID that is mentioned in a cast | 
 **optional** | ***CastsApiListCastsByMentionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastsApiListCastsByMentionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCastsByParent**
> InlineResponse200 ListCastsByParent(ctx, apiKey, optional)
Fetch all casts by parent cast's FID and Hash OR by the parent's URL

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***CastsApiListCastsByParentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastsApiListCastsByParentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fid** | **optional.Int32**| The FID of the parent cast | 
 **hash** | **optional.String**| The parent cast&#x27;s hash | 
 **url** | **optional.String**|  | 
 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

