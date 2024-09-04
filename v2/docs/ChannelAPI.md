# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActiveChannels**](ChannelApi.md#ActiveChannels) | **Get** /farcaster/channel/user | Get channels that a user is active in
[**ChannelDetails**](ChannelApi.md#ChannelDetails) | **Get** /farcaster/channel | Retrieve channel details by id or parent_url
[**ChannelDetailsBulk**](ChannelApi.md#ChannelDetailsBulk) | **Get** /farcaster/channel/bulk | (Bulk) Retrieve channels by id or parent_url
[**ChannelFollowers**](ChannelApi.md#ChannelFollowers) | **Get** /farcaster/channel/followers | Retrieve followers for a given channel
[**ChannelUsers**](ChannelApi.md#ChannelUsers) | **Get** /farcaster/channel/users | Retrieve users who are active in a channel
[**ListAllChannels**](ChannelApi.md#ListAllChannels) | **Get** /farcaster/channel/list | Retrieve all channels with their details
[**SearchChannels**](ChannelApi.md#SearchChannels) | **Get** /farcaster/channel/search | Search for channels based on id or name
[**TrendingChannels**](ChannelApi.md#TrendingChannels) | **Get** /farcaster/channel/trending | Retrieve trending channels based on activity
[**UserChannels**](ChannelApi.md#UserChannels) | **Get** /farcaster/user/channels | Retrieve all channels that a given fid follows

# **ActiveChannels**
> UsersActiveChannelsResponse ActiveChannels(ctx, apiKey, fid, optional)
Get channels that a user is active in

Fetches all channels that a user has casted in, in reverse chronological order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| The user&#x27;s fid (identifier) | 
 **optional** | ***ChannelApiActiveChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiActiveChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Number of results to retrieve (default 20, max 100). | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**UsersActiveChannelsResponse**](UsersActiveChannelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelDetails**
> ChannelResponse ChannelDetails(ctx, apiKey, id, optional)
Retrieve channel details by id or parent_url

Returns details of a channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **id** | **string**| Channel ID for the channel being queried | 
 **optional** | ***ChannelApiChannelDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiChannelDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**optional.Interface of ChannelType**](.md)| Type of identifier being used to query the channel. Defaults to id. | 
 **viewerFid** | [**optional.Interface of int32**](.md)| FID of the user viewing the channel. | 

### Return type

[**ChannelResponse**](ChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelDetailsBulk**
> ChannelResponseBulk ChannelDetailsBulk(ctx, apiKey, ids, optional)
(Bulk) Retrieve channels by id or parent_url

Returns details of multiple channels

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **ids** | **string**| Comma separated list of channel IDs or parent_urls, up to 100 at a time | 
 **optional** | ***ChannelApiChannelDetailsBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiChannelDetailsBulkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **type_** | [**optional.Interface of ChannelType**](.md)| Type of identifier being used to query the channels. Defaults to id. | 
 **viewerFid** | [**optional.Interface of int32**](.md)| FID of the user viewing the channels. | 

### Return type

[**ChannelResponseBulk**](ChannelResponseBulk.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelFollowers**
> UsersResponse ChannelFollowers(ctx, apiKey, id, optional)
Retrieve followers for a given channel

Returns a list of followers for a specific channel. Max limit is 1000. Use cursor for pagination.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **id** | **string**| Channel ID for the channel being queried | 
 **optional** | ***ChannelApiChannelFollowersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiChannelFollowersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **optional.String**| Pagination cursor. | 
 **limit** | **optional.Int32**| Number of followers to retrieve (default 25, max 1000) | [default to 25]

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelUsers**
> UsersResponse ChannelUsers(ctx, apiKey, id, hasRootCastAuthors, optional)
Retrieve users who are active in a channel

Returns a list of users who are active in a given channel, ordered by ascending FIDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **id** | **string**| Channel ID for the channel being queried | 
  **hasRootCastAuthors** | **bool**| Include users who posted the root cast in the channel | 
 **optional** | ***ChannelApiChannelUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiChannelUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **hasCastLikers** | **optional.Bool**| Include users who liked a cast in the channel | 
 **hasCastRecasters** | **optional.Bool**| Include users who recasted a cast in the channel | 
 **hasReplyAuthors** | **optional.Bool**| Include users who replied to a cast in the channel | 
 **cursor** | **optional.String**| Pagination cursor. | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllChannels**
> ChannelListResponse ListAllChannels(ctx, apiKey, optional)
Retrieve all channels with their details

Returns a list of all channels with their details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***ChannelApiListAllChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiListAllChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to retrieve | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChannels**
> ChannelSearchResponse SearchChannels(ctx, apiKey, q, optional)
Search for channels based on id or name

Returns a list of channels based on id or name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **q** | **string**| Channel ID or name for the channel being queried | 
 **optional** | ***ChannelApiSearchChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiSearchChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Number of results to retrieve | [default to 20]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**ChannelSearchResponse**](ChannelSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrendingChannels**
> TrendingChannelResponse TrendingChannels(ctx, apiKey, optional)
Retrieve trending channels based on activity

Returns a list of trending channels based on activity

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***ChannelApiTrendingChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiTrendingChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeWindow** | **optional.String**|  | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 10, max 25) | [default to 10]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**TrendingChannelResponse**](TrendingChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserChannels**
> ChannelListResponse UserChannels(ctx, apiKey, fid, optional)
Retrieve all channels that a given fid follows

Returns a list of all channels with their details that a fid follows.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| The fid of the user. | 
 **optional** | ***ChannelApiUserChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelApiUserChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

