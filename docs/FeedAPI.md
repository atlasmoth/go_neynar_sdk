# \FeedAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                          | HTTP request                                     | Description                                                     |
| --------------------------------------------------------------- | ------------------------------------------------ | --------------------------------------------------------------- |
| [**Feed**](FeedAPI.md#Feed)                                     | **Get** /farcaster/feed                          | Retrieve casts based on filters                                 |
| [**FeedChannels**](FeedAPI.md#FeedChannels)                     | **Get** /farcaster/feed/channels                 | Retrieve feed based on channel ids                              |
| [**FeedFollowing**](FeedAPI.md#FeedFollowing)                   | **Get** /farcaster/feed/following                | Retrieve feed based on who a user is following                  |
| [**FeedForYou**](FeedAPI.md#FeedForYou)                         | **Get** /farcaster/feed/for_you                  | Retrieve a personalized For You feed for a user                 |
| [**FeedFrames**](FeedAPI.md#FeedFrames)                         | **Get** /farcaster/feed/frames                   | Retrieve feed of casts with Frames, reverse chronological order |
| [**FeedTrending**](FeedAPI.md#FeedTrending)                     | **Get** /farcaster/feed/trending                 | Retrieve trending casts                                         |
| [**FeedUserPopular**](FeedAPI.md#FeedUserPopular)               | **Get** /farcaster/feed/user/popular             | Retrieve 10 most popular casts for a user                       |
| [**FeedUserRepliesRecasts**](FeedAPI.md#FeedUserRepliesRecasts) | **Get** /farcaster/feed/user/replies_and_recasts | Retrieve recent replies and recasts for a user                  |

## Feed

> FeedResponse Feed(ctx).ApiKey(apiKey).FeedType(feedType).FilterType(filterType).Fid(fid).Fids(fids).ParentUrl(parentUrl).ChannelId(channelId).EmbedUrl(embedUrl).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).Execute()

Retrieve casts based on filters

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	feedType := openapiclient.FeedType("following") // FeedType | Defaults to following (requires fid or address). If set to filter (requires filter_type)
	filterType := openapiclient.FilterType("fids") // FilterType | Used when feed_type=filter. Can be set to fids (requires fids) or parent_url (requires parent_url) or channel_id (requires channel_id) (optional)
	fid := int32(56) // int32 | (Optional) fid of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type (optional)
	fids := "3,2,194" // string | Used when filter_type=fids . Create a feed based on a list of fids. Max array size is 250. Requires feed_type and filter_type. (optional)
	parentUrl := "parentUrl_example" // string | Used when filter_type=parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type (optional)
	channelId := "channelId_example" // string | Used when filter_type=channel_id can be used to fetch all casts under a channel. Requires feed_type and filter_type (optional)
	embedUrl := "embedUrl_example" // string | Used when filter_type=embed_url can be used to fetch all casts with an embed url that contains embed_url. Requires feed_type and filter_type (optional)
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.Feed(context.Background()).ApiKey(apiKey).FeedType(feedType).FilterType(filterType).Fid(fid).Fids(fids).ParentUrl(parentUrl).ChannelId(channelId).EmbedUrl(embedUrl).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.Feed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Feed`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.Feed`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedRequest struct via the builder pattern

| Name            | Type                            | Description                                                                                                                                                 | Notes                                    |
| --------------- | ------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**      | **string**                      | API key required for authentication.                                                                                                                        | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **feedType**    | [**FeedType**](FeedType.md)     | Defaults to following (requires fid or address). If set to filter (requires filter_type)                                                                    |
| **filterType**  | [**FilterType**](FilterType.md) | Used when feed_type&#x3D;filter. Can be set to fids (requires fids) or parent_url (requires parent_url) or channel_id (requires channel_id)                 |
| **fid**         | **int32**                       | (Optional) fid of user whose feed you want to create. By default, the API expects this field, except if you pass a filter_type                              |
| **fids**        | **string**                      | Used when filter_type&#x3D;fids . Create a feed based on a list of fids. Max array size is 250. Requires feed_type and filter_type.                         |
| **parentUrl**   | **string**                      | Used when filter_type&#x3D;parent_url can be used to fetch content under any parent url e.g. FIP-2 channels on Warpcast. Requires feed_type and filter_type |
| **channelId**   | **string**                      | Used when filter_type&#x3D;channel_id can be used to fetch all casts under a channel. Requires feed_type and filter_type                                    |
| **embedUrl**    | **string**                      | Used when filter_type&#x3D;embed_url can be used to fetch all casts with an embed url that contains embed_url. Requires feed_type and filter_type           |
| **withRecasts** | **bool**                        | Include recasts in the response, true by default                                                                                                            | [default to true]                        |
| **limit**       | **int32**                       | Number of results to retrieve (default 25, max 100)                                                                                                         | [default to 25]                          |
| **cursor**      | **string**                      | Pagination cursor.                                                                                                                                          |
| **viewerFid**   | **int32**                       |                                                                                                                                                             |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedChannels

> FeedResponse FeedChannels(ctx).ApiKey(apiKey).ChannelIds(channelIds).WithRecasts(withRecasts).ViewerFid(viewerFid).WithReplies(withReplies).Limit(limit).Cursor(cursor).ShouldModerate(shouldModerate).Execute()

Retrieve feed based on channel ids

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	channelIds := "neynar,farcaster" // string | comma separated list of channel ids e.g. neynar,farcaster
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	viewerFid := int32(3) // int32 |  (optional)
	withReplies := true // bool | Include replies in the response, false by default (optional) (default to false)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	shouldModerate := true // bool | If true, only casts that have been liked by the moderator (if one exists) will be returned. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedChannels(context.Background()).ApiKey(apiKey).ChannelIds(channelIds).WithRecasts(withRecasts).ViewerFid(viewerFid).WithReplies(withReplies).Limit(limit).Cursor(cursor).ShouldModerate(shouldModerate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedChannels`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedChannels`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedChannelsRequest struct via the builder pattern

| Name               | Type       | Description                                                                                 | Notes                                    |
| ------------------ | ---------- | ------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**         | **string** | API key required for authentication.                                                        | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **channelIds**     | **string** | comma separated list of channel ids e.g. neynar,farcaster                                   |
| **withRecasts**    | **bool**   | Include recasts in the response, true by default                                            | [default to true]                        |
| **viewerFid**      | **int32**  |                                                                                             |
| **withReplies**    | **bool**   | Include replies in the response, false by default                                           | [default to false]                       |
| **limit**          | **int32**  | Number of results to retrieve (default 25, max 100)                                         | [default to 25]                          |
| **cursor**         | **string** | Pagination cursor.                                                                          |
| **shouldModerate** | **bool**   | If true, only casts that have been liked by the moderator (if one exists) will be returned. | [default to false]                       |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedFollowing

> FeedResponse FeedFollowing(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).Execute()

Retrieve feed based on who a user is following

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(3) // int32 | fid of user whose feed you want to create
	viewerFid := int32(3) // int32 |  (optional)
	withRecasts := true // bool | Include recasts in the response, true by default (optional) (default to true)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedFollowing(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).WithRecasts(withRecasts).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedFollowing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedFollowing`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedFollowing`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedFollowingRequest struct via the builder pattern

| Name            | Type       | Description                                         | Notes                                    |
| --------------- | ---------- | --------------------------------------------------- | ---------------------------------------- |
| **apiKey**      | **string** | API key required for authentication.                | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**         | **int32**  | fid of user whose feed you want to create           |
| **viewerFid**   | **int32**  |                                                     |
| **withRecasts** | **bool**   | Include recasts in the response, true by default    | [default to true]                        |
| **limit**       | **int32**  | Number of results to retrieve (default 25, max 100) | [default to 25]                          |
| **cursor**      | **string** | Pagination cursor.                                  |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedForYou

> FeedResponse FeedForYou(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Provider(provider).Limit(limit).Cursor(cursor).Execute()

Retrieve a personalized For You feed for a user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(194) // int32 | fid of user whose feed you want to create
	viewerFid := int32(3) // int32 |  (optional)
	provider := openapiclient.ForYouProvider("karma3") // ForYouProvider |  (optional)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 50) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedForYou(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Provider(provider).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedForYou``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedForYou`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedForYou`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedForYouRequest struct via the builder pattern

| Name          | Type                                    | Description                                        | Notes                                    |
| ------------- | --------------------------------------- | -------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                              | API key required for authentication.               | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**                               | fid of user whose feed you want to create          |
| **viewerFid** | **int32**                               |                                                    |
| **provider**  | [**ForYouProvider**](ForYouProvider.md) |                                                    |
| **limit**     | **int32**                               | Number of results to retrieve (default 25, max 50) | [default to 25]                          |
| **cursor**    | **string**                              | Pagination cursor.                                 |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedFrames

> FeedResponse FeedFrames(ctx).ApiKey(apiKey).Limit(limit).ViewerFid(viewerFid).Cursor(cursor).Execute()

Retrieve feed of casts with Frames, reverse chronological order

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	viewerFid := int32(3) // int32 |  (optional)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedFrames(context.Background()).ApiKey(apiKey).Limit(limit).ViewerFid(viewerFid).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedFrames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedFrames`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedFrames`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedFramesRequest struct via the builder pattern

| Name          | Type       | Description                                         | Notes                                    |
| ------------- | ---------- | --------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 100) | [default to 25]                          |
| **viewerFid** | **int32**  |                                                     |
| **cursor**    | **string** | Pagination cursor.                                  |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedTrending

> FeedResponse FeedTrending(ctx).ApiKey(apiKey).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).TimeWindow(timeWindow).ChannelId(channelId).Execute()

Retrieve trending casts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	limit := int32(56) // int32 | Number of results to retrieve (max 10) (optional) (default to 10)
	cursor := "cursor_example" // string | Pagination cursor (optional)
	viewerFid := int32(3) // int32 |  (optional)
	timeWindow := "timeWindow_example" // string | Time window for trending casts (7d window for channel feeds only) (optional) (default to "24h")
	channelId := "neynar" // string | Channel ID to filter trending casts. Less active channels might have no casts in the time window selected. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedTrending(context.Background()).ApiKey(apiKey).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).TimeWindow(timeWindow).ChannelId(channelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedTrending``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedTrending`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedTrending`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedTrendingRequest struct via the builder pattern

| Name           | Type       | Description                                                                                                | Notes                                    |
| -------------- | ---------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**     | **string** | API key required for authentication.                                                                       | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **limit**      | **int32**  | Number of results to retrieve (max 10)                                                                     | [default to 10]                          |
| **cursor**     | **string** | Pagination cursor                                                                                          |
| **viewerFid**  | **int32**  |                                                                                                            |
| **timeWindow** | **string** | Time window for trending casts (7d window for channel feeds only)                                          | [default to &quot;24h&quot;]             |
| **channelId**  | **string** | Channel ID to filter trending casts. Less active channels might have no casts in the time window selected. |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedUserPopular

> BulkCastsResponse FeedUserPopular(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Execute()

Retrieve 10 most popular casts for a user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(194) // int32 | fid of user whose feed you want to create
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedUserPopular(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedUserPopular``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedUserPopular`: BulkCastsResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedUserPopular`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedUserPopularRequest struct via the builder pattern

| Name          | Type       | Description                               | Notes                                    |
| ------------- | ---------- | ----------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.      | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | fid of user whose feed you want to create |
| **viewerFid** | **int32**  |                                           |

### Return type

[**BulkCastsResponse**](BulkCastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FeedUserRepliesRecasts

> FeedResponse FeedUserRepliesRecasts(ctx).ApiKey(apiKey).Fid(fid).Filter(filter).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).Execute()

Retrieve recent replies and recasts for a user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(194) // int32 | fid of user whose replies and recasts you want to fetch
	filter := "replies" // string | filter to fetch only replies or recasts (optional) (default to "all")
	limit := int32(25) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeedAPI.FeedUserRepliesRecasts(context.Background()).ApiKey(apiKey).Fid(fid).Filter(filter).Limit(limit).Cursor(cursor).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeedAPI.FeedUserRepliesRecasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FeedUserRepliesRecasts`: FeedResponse
	fmt.Fprintf(os.Stdout, "Response from `FeedAPI.FeedUserRepliesRecasts`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFeedUserRepliesRecastsRequest struct via the builder pattern

| Name          | Type       | Description                                             | Notes                                    |
| ------------- | ---------- | ------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | fid of user whose replies and recasts you want to fetch |
| **filter**    | **string** | filter to fetch only replies or recasts                 | [default to &quot;all&quot;]             |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 100)     | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                                      |
| **viewerFid** | **int32**  |                                                         |

### Return type

[**FeedResponse**](FeedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
