# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cast**](CastApi.md#Cast) | **Get** /farcaster/cast | Retrieve cast for a given hash or Warpcast URL
[**CastConversation**](CastApi.md#CastConversation) | **Get** /farcaster/cast/conversation | Retrieve the conversation for a given cast
[**CastSearch**](CastApi.md#CastSearch) | **Get** /farcaster/cast/search | Search for casts
[**Casts**](CastApi.md#Casts) | **Get** /farcaster/casts | Gets information about an array of casts
[**ComposerList**](CastApi.md#ComposerList) | **Get** /farcaster/cast/composer_actions/list | Fetches all composer actions on Warpcast
[**DeleteCast**](CastApi.md#DeleteCast) | **Delete** /farcaster/cast | Delete a cast
[**PostCast**](CastApi.md#PostCast) | **Post** /farcaster/cast | Posts a cast

# **Cast**
> CastResponse Cast(ctx, apiKey, identifier, type_, optional)
Retrieve cast for a given hash or Warpcast URL

Gets information about an individual cast by passing in a Warpcast web URL or cast hash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **identifier** | **string**| Cast identifier (Its either a url or a hash) | 
  **type_** | [**CastParamType**](.md)|  | 
 **optional** | ***CastApiCastOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiCastOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewerFid** | [**optional.Interface of int32**](.md)| adds viewer_context to cast object to show whether viewer has liked or recasted the cast. | 

### Return type

[**CastResponse**](CastResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CastConversation**
> Conversation CastConversation(ctx, apiKey, identifier, type_, optional)
Retrieve the conversation for a given cast

Gets all casts related to a conversation surrounding a cast by passing in a cast hash or Warpcast URL. Includes all the ancestors of a cast up to the root parent in a chronological order. Includes all direct_replies to the cast up to the reply_depth specified in the query parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **identifier** | **string**| Cast identifier (Its either a url or a hash) | 
  **type_** | [**CastParamType**](.md)|  | 
 **optional** | ***CastApiCastConversationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiCastConversationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **replyDepth** | [**optional.Interface of int32**](.md)| The depth of replies in the conversation that will be returned (default 2) | 
 **includeChronologicalParentCasts** | **optional.Bool**| Include all parent casts in chronological order | [default to false]
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 20, max 50) | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CastSearch**
> CastsSearchResponse CastSearch(ctx, apiKey, q, optional)
Search for casts

Search for casts based on a query string, with optional AND filters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **q** | **string**| Query string to search for casts | 
 **optional** | ***CastApiCastSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiCastSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorFid** | [**optional.Interface of int32**](.md)| Fid of the user whose casts you want to search | 
 **viewerFid** | [**optional.Interface of int32**](.md)| Fid of the viewer of the casts, used to show viewer_context | 
 **parentUrl** | **optional.String**| Parent URL of the casts you want to search | 
 **channelId** | **optional.String**| Channel ID of the casts you want to search | 
 **limit** | **optional.Int32**|  | [default to 25]
 **cursor** | **optional.String**| Pagination cursor | 

### Return type

[**CastsSearchResponse**](CastsSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Casts**
> CastsResponse Casts(ctx, apiKey, casts, optional)
Gets information about an array of casts

Retrieve multiple casts using their respective hashes.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **casts** | **string**| Hashes of the cast to be retrived (Comma separated) | 
 **optional** | ***CastApiCastsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiCastsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| adds viewer_context to cast object to show whether viewer has liked or recasted the cast. | 
 **sortType** | **optional.String**| Optional parameter to sort the casts based on different criteria | 

### Return type

[**CastsResponse**](CastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComposerList**
> CastComposerActionsListResponse ComposerList(ctx, apiKey, list, optional)
Fetches all composer actions on Warpcast

Fetches all composer actions on Warpcast. You can filter by top or featured.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **list** | [**CastComposerType**](.md)| Type of list to fetch. | 
 **optional** | ***CastApiComposerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CastApiComposerListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 25). | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**CastComposerActionsListResponse**](CastComposerActionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCast**
> OperationResponse DeleteCast(ctx, body, apiKey)
Delete a cast

Delete an existing cast. \\ (In order to delete a cast `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DeleteCastReqBody**](DeleteCastReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostCast**
> PostCastResponse PostCast(ctx, body, apiKey)
Posts a cast

Posts a cast or cast reply. Works with mentions and embeds.   (In order to post a cast `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PostCastReqBody**](PostCastReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**PostCastResponse**](PostCastResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

