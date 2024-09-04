# \ChannelAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                     | HTTP request                         | Description                                    |
| ---------------------------------------------------------- | ------------------------------------ | ---------------------------------------------- |
| [**ActiveChannels**](ChannelAPI.md#ActiveChannels)         | **Get** /farcaster/channel/user      | Get channels that a user is active in          |
| [**ChannelDetails**](ChannelAPI.md#ChannelDetails)         | **Get** /farcaster/channel           | Retrieve channel details by id or parent_url   |
| [**ChannelDetailsBulk**](ChannelAPI.md#ChannelDetailsBulk) | **Get** /farcaster/channel/bulk      | (Bulk) Retrieve channels by id or parent_url   |
| [**ChannelFollowers**](ChannelAPI.md#ChannelFollowers)     | **Get** /farcaster/channel/followers | Retrieve followers for a given channel         |
| [**ChannelUsers**](ChannelAPI.md#ChannelUsers)             | **Get** /farcaster/channel/users     | Retrieve users who are active in a channel     |
| [**ListAllChannels**](ChannelAPI.md#ListAllChannels)       | **Get** /farcaster/channel/list      | Retrieve all channels with their details       |
| [**SearchChannels**](ChannelAPI.md#SearchChannels)         | **Get** /farcaster/channel/search    | Search for channels based on id or name        |
| [**TrendingChannels**](ChannelAPI.md#TrendingChannels)     | **Get** /farcaster/channel/trending  | Retrieve trending channels based on activity   |
| [**UserChannels**](ChannelAPI.md#UserChannels)             | **Get** /farcaster/user/channels     | Retrieve all channels that a given fid follows |

## ActiveChannels

> UsersActiveChannelsResponse ActiveChannels(ctx).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Get channels that a user is active in

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	fid := int32(194) // int32 | The user's fid (identifier) (optional)
	limit := int32(20) // int32 | Number of results to retrieve (default 20, max 100). (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.ActiveChannels(context.Background()).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ActiveChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActiveChannels`: UsersActiveChannelsResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ActiveChannels`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiActiveChannelsRequest struct via the builder pattern

| Name       | Type       | Description                                          | Notes                                    |
| ---------- | ---------- | ---------------------------------------------------- | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication.                 | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | The user&#39;s fid (identifier)                      |
| **limit**  | **int32**  | Number of results to retrieve (default 20, max 100). | [default to 20]                          |
| **cursor** | **string** | Pagination cursor.                                   |

### Return type

[**UsersActiveChannelsResponse**](UsersActiveChannelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ChannelDetails

> ChannelResponse ChannelDetails(ctx).ApiKey(apiKey).Id(id).Type*(type*).ViewerFid(viewerFid).Execute()

Retrieve channel details by id or parent_url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	id := "neynar" // string | Channel ID for the channel being queried (optional)
	type_ := openapiclient.ChannelType("id") // ChannelType | Type of identifier being used to query the channel. Defaults to id. (optional)
	viewerFid := int32(194) // int32 | FID of the user viewing the channel. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.ChannelDetails(context.Background()).ApiKey(apiKey).Id(id).Type_(type_).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelDetails`: ChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ChannelDetails`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiChannelDetailsRequest struct via the builder pattern

| Name          | Type                              | Description                                                         | Notes                                    |
| ------------- | --------------------------------- | ------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                        | API key required for authentication.                                | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **id**        | **string**                        | Channel ID for the channel being queried                            |
| **type\_**    | [**ChannelType**](ChannelType.md) | Type of identifier being used to query the channel. Defaults to id. |
| **viewerFid** | **int32**                         | FID of the user viewing the channel.                                |

### Return type

[**ChannelResponse**](ChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ChannelDetailsBulk

> ChannelResponseBulk ChannelDetailsBulk(ctx).ApiKey(apiKey).Ids(ids).Type*(type*).ViewerFid(viewerFid).Execute()

(Bulk) Retrieve channels by id or parent_url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	ids := "neynar,warpcast" // string | Comma separated list of channel IDs or parent_urls, up to 100 at a time (optional)
	type_ := openapiclient.ChannelType("id") // ChannelType | Type of identifier being used to query the channels. Defaults to id. (optional)
	viewerFid := int32(194) // int32 | FID of the user viewing the channels. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.ChannelDetailsBulk(context.Background()).ApiKey(apiKey).Ids(ids).Type_(type_).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelDetailsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelDetailsBulk`: ChannelResponseBulk
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ChannelDetailsBulk`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiChannelDetailsBulkRequest struct via the builder pattern

| Name          | Type                              | Description                                                             | Notes                                    |
| ------------- | --------------------------------- | ----------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                        | API key required for authentication.                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **ids**       | **string**                        | Comma separated list of channel IDs or parent_urls, up to 100 at a time |
| **type\_**    | [**ChannelType**](ChannelType.md) | Type of identifier being used to query the channels. Defaults to id.    |
| **viewerFid** | **int32**                         | FID of the user viewing the channels.                                   |

### Return type

[**ChannelResponseBulk**](ChannelResponseBulk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ChannelFollowers

> UsersResponse ChannelFollowers(ctx).ApiKey(apiKey).Id(id).Cursor(cursor).Limit(limit).Execute()

Retrieve followers for a given channel

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	id := "founders" // string | Channel ID for the channel being queried (optional)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	limit := int32(56) // int32 | Number of followers to retrieve (default 25, max 1000) (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.ChannelFollowers(context.Background()).ApiKey(apiKey).Id(id).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelFollowers`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ChannelFollowers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiChannelFollowersRequest struct via the builder pattern

| Name       | Type       | Description                                            | Notes                                    |
| ---------- | ---------- | ------------------------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication.                   | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **id**     | **string** | Channel ID for the channel being queried               |
| **cursor** | **string** | Pagination cursor.                                     |
| **limit**  | **int32**  | Number of followers to retrieve (default 25, max 1000) | [default to 25]                          |

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ChannelUsers

> UsersResponse ChannelUsers(ctx).ApiKey(apiKey).Id(id).HasRootCastAuthors(hasRootCastAuthors).HasCastLikers(hasCastLikers).HasCastRecasters(hasCastRecasters).HasReplyAuthors(hasReplyAuthors).Cursor(cursor).Limit(limit).Execute()

Retrieve users who are active in a channel

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	id := "neynar" // string | Channel ID for the channel being queried (optional)
	hasRootCastAuthors := false // bool | Include users who posted the root cast in the channel (optional)
	hasCastLikers := false // bool | Include users who liked a cast in the channel (optional)
	hasCastRecasters := false // bool | Include users who recasted a cast in the channel (optional)
	hasReplyAuthors := false // bool | Include users who replied to a cast in the channel (optional)
	cursor := "cursor_example" // string | Pagination cursor. (optional)
	limit := int32(25) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.ChannelUsers(context.Background()).ApiKey(apiKey).Id(id).HasRootCastAuthors(hasRootCastAuthors).HasCastLikers(hasCastLikers).HasCastRecasters(hasCastRecasters).HasReplyAuthors(hasReplyAuthors).Cursor(cursor).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ChannelUsers`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ChannelUsers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiChannelUsersRequest struct via the builder pattern

| Name                   | Type       | Description                                           | Notes                                    |
| ---------------------- | ---------- | ----------------------------------------------------- | ---------------------------------------- |
| **apiKey**             | **string** | API key required for authentication.                  | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **id**                 | **string** | Channel ID for the channel being queried              |
| **hasRootCastAuthors** | **bool**   | Include users who posted the root cast in the channel |
| **hasCastLikers**      | **bool**   | Include users who liked a cast in the channel         |
| **hasCastRecasters**   | **bool**   | Include users who recasted a cast in the channel      |
| **hasReplyAuthors**    | **bool**   | Include users who replied to a cast in the channel    |
| **cursor**             | **string** | Pagination cursor.                                    |
| **limit**              | **int32**  | Number of results to retrieve (default 25, max 100)   | [default to 25]                          |

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAllChannels

> ChannelListResponse ListAllChannels(ctx).ApiKey(apiKey).Limit(limit).Cursor(cursor).Execute()

Retrieve all channels with their details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	limit := int32(56) // int32 | Number of results to retrieve (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.ListAllChannels(context.Background()).ApiKey(apiKey).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ListAllChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllChannels`: ChannelListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.ListAllChannels`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListAllChannelsRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **limit**  | **int32**  | Number of results to retrieve        | [default to 20]                          |
| **cursor** | **string** | Pagination cursor.                   |

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SearchChannels

> ChannelSearchResponse SearchChannels(ctx).ApiKey(apiKey).Q(q).Limit(limit).Cursor(cursor).Execute()

Search for channels based on id or name

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	q := "neynar" // string | Channel ID or name for the channel being queried (optional)
	limit := int32(56) // int32 | Number of results to retrieve (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.SearchChannels(context.Background()).ApiKey(apiKey).Q(q).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.SearchChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChannels`: ChannelSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.SearchChannels`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSearchChannelsRequest struct via the builder pattern

| Name       | Type       | Description                                      | Notes                                    |
| ---------- | ---------- | ------------------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication.             | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **q**      | **string** | Channel ID or name for the channel being queried |
| **limit**  | **int32**  | Number of results to retrieve                    | [default to 20]                          |
| **cursor** | **string** | Pagination cursor.                               |

### Return type

[**ChannelSearchResponse**](ChannelSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## TrendingChannels

> TrendingChannelResponse TrendingChannels(ctx).ApiKey(apiKey).TimeWindow(timeWindow).Limit(limit).Cursor(cursor).Execute()

Retrieve trending channels based on activity

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	timeWindow := "timeWindow_example" // string |  (optional)
	limit := int32(10) // int32 | Number of results to retrieve (default 10, max 25) (optional) (default to 10)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.TrendingChannels(context.Background()).ApiKey(apiKey).TimeWindow(timeWindow).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.TrendingChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrendingChannels`: TrendingChannelResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.TrendingChannels`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiTrendingChannelsRequest struct via the builder pattern

| Name           | Type       | Description                                        | Notes                                    |
| -------------- | ---------- | -------------------------------------------------- | ---------------------------------------- |
| **apiKey**     | **string** | API key required for authentication.               | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **timeWindow** | **string** |                                                    |
| **limit**      | **int32**  | Number of results to retrieve (default 10, max 25) | [default to 10]                          |
| **cursor**     | **string** | Pagination cursor.                                 |

### Return type

[**TrendingChannelResponse**](TrendingChannelResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserChannels

> ChannelListResponse UserChannels(ctx).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Retrieve all channels that a given fid follows

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	fid := int32(56) // int32 | The fid of the user. (optional)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ChannelAPI.UserChannels(context.Background()).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.UserChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserChannels`: ChannelListResponse
	fmt.Fprintf(os.Stdout, "Response from `ChannelAPI.UserChannels`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserChannelsRequest struct via the builder pattern

| Name       | Type       | Description                                         | Notes                                    |
| ---------- | ---------- | --------------------------------------------------- | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication.                | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | The fid of the user.                                |
| **limit**  | **int32**  | Number of results to retrieve (default 25, max 100) | [default to 25]                          |
| **cursor** | **string** | Pagination cursor.                                  |

### Return type

[**ChannelListResponse**](ChannelListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
