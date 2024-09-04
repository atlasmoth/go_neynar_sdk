# {{classname}}

All URIs are relative to *https://api.neynar.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserByVerification**](VerificationApi.md#UserByVerification) | **Get** /farcaster/user-by-verification | DEPRECATED - Retrieve user for a given ethereum address
[**Verifications**](VerificationApi.md#Verifications) | **Get** /farcaster/verifications | DEPRECATED - Retrieve verifications for a given FID

# **UserByVerification**
> UserResponse UserByVerification(ctx, apiKey, address)
DEPRECATED - Retrieve user for a given ethereum address

Now deprecated. Use [v2/user/bulk-by-address](https://docs.neynar.com/reference/user-bulk-by-address). Checks if a given Ethereum address has a Farcaster user associated with it. Note: if an address is associated with multiple users, the API will return the user who most recently published a verification with the address (based on when Warpcast received the proof, not a self-reported timestamp).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **address** | [**string**](.md)|  | 

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Verifications**
> VerificationResponse Verifications(ctx, apiKey, fid)
DEPRECATED - Retrieve verifications for a given FID

Now deprecated, use [v2/user/bulk](https://docs.neynar.com/reference/user-bulk), verifications are in the user object. Get all known verifications of a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of the user | 

### Return type

[**VerificationResponse**](VerificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

