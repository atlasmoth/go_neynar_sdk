# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsernameProof**](UsernamesApi.md#GetUsernameProof) | **Get** /v1/userNameProofByName | Get an proof for a username by the Farcaster username
[**ListUsernameProofsByFid**](UsernamesApi.md#ListUsernameProofsByFid) | **Get** /v1/userNameProofsByFid | Get a list of proofs provided by an FID

# **GetUsernameProof**
> UserNameProof GetUsernameProof(ctx, apiKey, name)
Get an proof for a username by the Farcaster username

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **name** | **string**| The Farcaster username or ENS address | 

### Return type

[**UserNameProof**](UserNameProof.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsernameProofsByFid**
> UsernameProofsResponse ListUsernameProofsByFid(ctx, apiKey, fid)
Get a list of proofs provided by an FID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID being requested | 

### Return type

[**UsernameProofsResponse**](UsernameProofsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

