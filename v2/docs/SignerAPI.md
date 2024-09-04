# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSigner**](SignerApi.md#CreateSigner) | **Post** /farcaster/signer | Creates a signer and returns the signer status
[**DeveloperManagedSigner**](SignerApi.md#DeveloperManagedSigner) | **Get** /farcaster/signer/developer_managed | Fetches the status of a signer by public key
[**FetchAuthorizationUrl**](SignerApi.md#FetchAuthorizationUrl) | **Get** /farcaster/login/authorize | Fetch authorization url
[**PublishMessage**](SignerApi.md#PublishMessage) | **Post** /farcaster/message | Publish a message to farcaster
[**RegisterSignedKey**](SignerApi.md#RegisterSignedKey) | **Post** /farcaster/signer/signed_key | Register Signed Key
[**RegisterSignedKeyForDeveloperManagedSigner**](SignerApi.md#RegisterSignedKeyForDeveloperManagedSigner) | **Post** /farcaster/signer/developer_managed/signed_key | Registers Signed Key
[**Signer**](SignerApi.md#Signer) | **Get** /farcaster/signer | Fetches the status of a signer

# **CreateSigner**
> Signer CreateSigner(ctx, apiKey)
Creates a signer and returns the signer status

Creates a signer and returns the signer status. \\ **Note**: While tesing please reuse the signer, it costs money to approve a signer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeveloperManagedSigner**
> DeveloperManagedSigner DeveloperManagedSigner(ctx, apiKey, publicKey)
Fetches the status of a signer by public key

Fetches the status of a developer managed signer by public key

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **publicKey** | [**string**](.md)|  | 

### Return type

[**DeveloperManagedSigner**](DeveloperManagedSigner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FetchAuthorizationUrl**
> AuthorizationUrlResponse FetchAuthorizationUrl(ctx, apiKey, clientId, responseType)
Fetch authorization url

Fetch authorization url (Fetched authorized url useful for SIWN login operation)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **clientId** | [**string**](.md)|  | 
  **responseType** | [**AuthorizationUrlResponseType**](.md)|  | 

### Return type

[**AuthorizationUrlResponse**](AuthorizationUrlResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PublishMessage**
> PublishMessageResponse PublishMessage(ctx, body, apiKey)
Publish a message to farcaster

Publish a message to farcaster. The message must be signed by a signer managed by the developer. Use the @farcaster/core library to construct and sign the message. Use the Message.toJSON method on the signed message and pass the JSON in the body of this POST request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PublishMessageReqBody**](PublishMessageReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**PublishMessageResponse**](PublishMessageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterSignedKey**
> Signer RegisterSignedKey(ctx, body, apiKey)
Register Signed Key

Registers an app fid, deadline and a signature. Returns the signer status with an approval url.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterSignerKeyReqBody**](RegisterSignerKeyReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterSignedKeyForDeveloperManagedSigner**
> DeveloperManagedSigner RegisterSignedKeyForDeveloperManagedSigner(ctx, body, apiKey)
Registers Signed Key

Registers an signed key and returns the developer managed signer status with an approval url.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegisterDeveloperManagedSignedKeyReqBody**](RegisterDeveloperManagedSignedKeyReqBody.md)|  | 
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]

### Return type

[**DeveloperManagedSigner**](DeveloperManagedSigner.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Signer**
> Signer Signer(ctx, apiKey, signerUuid)
Fetches the status of a signer

Gets information status of a signer by passing in a signer_uuid (Use post API to generate a signer)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **signerUuid** | [**string**](.md)|  | 

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

