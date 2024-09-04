# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteReaction**](ReactionApi.md#DeleteReaction) | **Delete** /farcaster/reaction | Delete a reaction
[**PostReaction**](ReactionApi.md#PostReaction) | **Post** /farcaster/reaction | Posts a reaction
[**ReactionsCast**](ReactionApi.md#ReactionsCast) | **Get** /farcaster/reactions/cast | Fetches reactions for a given cast
[**ReactionsUser**](ReactionApi.md#ReactionsUser) | **Get** /farcaster/reactions/user | Fetches reactions for a given user

# **DeleteReaction**
> OperationResponse DeleteReaction(ctx, body, apiKey)
Delete a reaction

Delete a reaction (like or recast) to a given cast \\ (In order to delete a reaction `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReactionReqBody**](ReactionReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReaction**
> OperationResponse PostReaction(ctx, body, apiKey)
Posts a reaction

Post a reaction (like or recast) to a given cast \\ (In order to post a reaction `signer_uuid` must be approved) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ReactionReqBody**](ReactionReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsCast**
> ReactionsCastResponse ReactionsCast(ctx, apiKey, hash, types, optional)
Fetches reactions for a given cast

Fetches reactions for a given cast

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **hash** | [**string**](.md)|  | 
  **types** | **string**| Customize which reaction types the request should search for. This is a comma-separated string that can include the following values: &#x27;likes&#x27; and &#x27;recasts&#x27;. By default api returns both. To select multiple types, use a comma-separated list of these values.  | 
 **optional** | ***ReactionApiReactionsCastOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionApiReactionsCastOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**ReactionsCastResponse**](ReactionsCastResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsUser**
> ReactionsResponse ReactionsUser(ctx, apiKey, fid, type_, optional)
Fetches reactions for a given user

Fetches reactions for a given user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)|  | 
  **type_** | [**ReactionsType**](.md)| Type of reaction to fetch (likes or recasts or all) | 
 **optional** | ***ReactionApiReactionsUserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionApiReactionsUserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**ReactionsResponse**](ReactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

