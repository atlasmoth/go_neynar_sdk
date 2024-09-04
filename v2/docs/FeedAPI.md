# {{classname}}

All URIs are relative to *https://api.neynar.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Feed**](FeedApi.md#Feed) | **Get** /farcaster/feed | Retrieve casts based on filters
[**FeedChannels**](FeedApi.md#FeedChannels) | **Get** /farcaster/feed/channels | Retrieve feed based on channel ids
[**FeedFollowing**](FeedApi.md#FeedFollowing) | **Get** /farcaster/feed/following | Retrieve feed based on who a user is following
[**FeedForYou**](FeedApi.md#FeedForYou) | **Get** /farcaster/feed/for_you | Retrieve a personalized For You feed for a user
[**FeedFrames**](FeedApi.md#FeedFrames) | **Get** /farcaster/feed/frames | Retrieve feed of casts with Frames, reverse chronological order
[**FeedParentUrls**](FeedApi.md#FeedParentUrls) | **Get** /farcaster/feed/parent_urls | Retrieve feed based on parent urls
[**FeedTrending**](FeedApi.md#FeedTrending) | **Get** /farcaster/feed/trending | Retrieve trending casts
[**FeedUserCasts**](FeedApi.md#FeedUserCasts) | **Get** /farcaster/feed/user/casts | Retrieve casts for a user
[**FeedUserPopular**](FeedApi.md#FeedUserPopular) | **Get** /farcaster/feed/user/popular | Retrieve 10 most popular casts for a user
[**FeedUserRepliesRecasts**](FeedApi.md#FeedUserRepliesRecasts) | **Get** /farcaster/feed/user/replies_and_recasts | Retrieve recent replies and recasts for a user

# **Feed**
> FeedResponse Feed(ctx, apiKey, feedType, optional)
Retrieve casts based on filters

Retrieve casts based on filters. Ensure setting the correct parameters based on the feed_type and filter_type.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **feedType** | [**FeedType**](.md)| Defaults to following (requires fid or address). If set to filter (requires filter_type) | 
 **optional** | ***FeedApiFeedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filterType** | [**optional.Interface of FilterType**](.md)| Used when feed_type&#x3D;filter. Can be set to fids (requires fids) or parent_url (requires parent_url) or channel_id (requires channel_id) | 
 **fid** | [**optional.Interface of int32**](.md)| (Optional) fid of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type | 
 **fids** | **optional.String**| Used when filter_type&#x3D;fids . Create a feed based on a list of fids. Max array size is 250. Requires feed_type and filter_type. | 
 **parentUrl** | **optional.String**| Used when filter_type&#x3D;parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type | 
 **channelId** | **optional.String**| Used when filter_type&#x3D;channel_id can be used to fetch all casts under a channel. Requires feed_type and filter_type | 
 **embedUrl** | **optional.String**| Used when filter_type&#x3D;embed_url can be used to fetch all casts with an embed url that contains embed_url. Requires feed_type and filter_type | 
 **embedTypes** | [**optional.Interface of []EmbedType**](EmbedType.md)| Used when filter_type&#x3D;embed_types can be used to fetch all casts with matching content types. Requires feed_type and filter_type | 
 **withRecasts** | **optional.Bool**| Include recasts in the response, true by default | [default to true]
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedChannels**
> FeedResponse FeedChannels(ctx, apiKey, channelIds, optional)
Retrieve feed based on channel ids

Retrieve feed based on channel ids

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **channelIds** | **string**| Comma separated list of channel ids e.g. neynar,farcaster | 
 **optional** | ***FeedApiFeedChannelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedChannelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withRecasts** | **optional.Bool**| Include recasts in the response, true by default | [default to true]
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **withReplies** | **optional.Bool**| Include replies in the response, false by default | [default to false]
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 
 **shouldModerate** | **optional.Bool**| If true, only casts that have been liked by the moderator (if one exists) will be returned. | [default to false]

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedFollowing**
> FeedResponse FeedFollowing(ctx, apiKey, fid, optional)
Retrieve feed based on who a user is following

Retrieve feed based on who a user is following

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of user whose feed you want to create | 
 **optional** | ***FeedApiFeedFollowingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedFollowingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **withRecasts** | **optional.Bool**| Include recasts in the response, true by default | [default to true]
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedForYou**
> FeedResponse FeedForYou(ctx, apiKey, fid, optional)
Retrieve a personalized For You feed for a user

Retrieve a personalized For You feed for a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of user whose feed you want to create | 
 **optional** | ***FeedApiFeedForYouOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedForYouOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **provider** | [**optional.Interface of ForYouProvider**](.md)|  | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 50) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedFrames**
> FeedResponse FeedFrames(ctx, apiKey, optional)
Retrieve feed of casts with Frames, reverse chronological order

Retrieve feed of casts with Frames, reverse chronological order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***FeedApiFeedFramesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedFramesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedParentUrls**
> FeedResponse FeedParentUrls(ctx, apiKey, parentUrls, optional)
Retrieve feed based on parent urls

Retrieve feed based on parent urls

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **parentUrls** | **string**| Comma separated list of parent_urls | 
 **optional** | ***FeedApiFeedParentUrlsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedParentUrlsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withRecasts** | **optional.Bool**| Include recasts in the response, true by default | [default to true]
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **withReplies** | **optional.Bool**| Include replies in the response, false by default | [default to false]
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedTrending**
> FeedResponse FeedTrending(ctx, apiKey, optional)
Retrieve trending casts

Retrieve trending casts or on the global feed or channels feeds. 7d time window available for channel feeds only.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
 **optional** | ***FeedApiFeedTrendingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedTrendingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| Number of results to retrieve (max 10) | [default to 10]
 **cursor** | **optional.String**| Pagination cursor | 
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 
 **timeWindow** | **optional.String**| Time window for trending casts (7d window for channel feeds only) | [default to 24h]
 **channelId** | **optional.String**| Channel ID to filter trending casts. Less active channels might have no casts in the time window selected. | 
 **provider** | [**optional.Interface of FeedTrendingProvider**](.md)| The provider of the trending casts feed. | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedUserCasts**
> FeedResponse FeedUserCasts(ctx, apiKey, fid, optional)
Retrieve casts for a user

Retrieve casts for a given user FID in reverse chronological order. Also allows filtering by parent_url and channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| FID of user whose recent casts you want to fetch | 
 **optional** | ***FeedApiFeedUserCastsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedUserCastsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)| FID of the user viewing the feed | 
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 50) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor | 
 **includeReplies** | **optional.Bool**| Include reply casts by the author in the response, true by default | [default to true]
 **parentUrl** | **optional.String**| Parent URL to filter the feed; mutually exclusive with channel_id | 
 **channelId** | **optional.String**| Channel ID to filter the feed; mutually exclusive with parent_url | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedUserPopular**
> BulkCastsResponse FeedUserPopular(ctx, apiKey, fid, optional)
Retrieve 10 most popular casts for a user

Retrieve 10 most popular casts for a given user FID; popularity based on replies, likes and recasts; sorted by most popular first

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of user whose feed you want to create | 
 **optional** | ***FeedApiFeedUserPopularOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedUserPopularOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**BulkCastsResponse**](BulkCastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FeedUserRepliesRecasts**
> FeedResponse FeedUserRepliesRecasts(ctx, apiKey, fid, optional)
Retrieve recent replies and recasts for a user

Retrieve recent replies and recasts for a given user FID; sorted by most recent first

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **apiKey** | **string**| API key required for authentication. | [default to NEYNAR_API_DOCS]
  **fid** | [**int32**](.md)| fid of user whose replies and recasts you want to fetch | 
 **optional** | ***FeedApiFeedUserRepliesRecastsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FeedApiFeedUserRepliesRecastsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| filter to fetch only replies or recasts | [default to all]
 **limit** | **optional.Int32**| Number of results to retrieve (default 25, max 100) | [default to 25]
 **cursor** | **optional.String**| Pagination cursor. | 
 **viewerFid** | [**optional.Interface of int32**](.md)|  | 

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

