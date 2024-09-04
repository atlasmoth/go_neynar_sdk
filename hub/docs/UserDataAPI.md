# {{classname}}

All URIs are relative to *https://hub-api.neynar.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserDataByFid**](UserDataApi.md#GetUserDataByFid) | **Get** /v1/userDataByFid | Get UserData for a FID.

# **GetUserDataByFid**
> InlineResponse2003 GetUserDataByFid(ctx, apiKey, fid, optional)
Get UserData for a FID.

**Note:** one of two different response schemas is returned  based on whether the caller provides the `user_data_type` parameter. If included, a single `UserDataAdd` message is returned (or a `not_found` error). If omitted, a paginated list of `UserDataAdd` messages is returned instead

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | **int32**| The FID that&#x27;s being requested | 
 **optional** | ***UserDataApiGetUserDataByFidOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserDataApiGetUserDataByFidOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userDataType** | [**optional.Interface of UserDataType**](.md)| The type of user data, either as a numerical value or type string. If this is omitted, all user data for the FID is returned | 
 **pageSize** | **optional.Int32**| Maximum number of messages to return in a single response | 
 **reverse** | **optional.Bool**| Reverse the sort order, returning latest messages first | 
 **pageToken** | **optional.String**| The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

