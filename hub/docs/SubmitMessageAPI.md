# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubmitMessage**](SubmitMessageApi.md#SubmitMessage) | **Post** /v1/submitMessage | Submit a signed protobuf-serialized message to the Hub

# **SubmitMessage**
> Message SubmitMessage(ctx, body)
Submit a signed protobuf-serialized message to the Hub

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](Object.md)| * 
A Message is a delta operation on the Farcaster network. The message protobuf is an envelope 
that wraps a MessageData object and contains a hash and signature which can verify its authenticity. | 

### Return type

[**Message**](Message.md)

### Authorization

[usernamePassword](../README.md#usernamePassword)

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

