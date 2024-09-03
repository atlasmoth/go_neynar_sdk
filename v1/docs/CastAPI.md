# \CastAPI

All URIs are relative to *https://api.neynar.com/v1*

| Method                                              | HTTP request                           | Description                                            |
| --------------------------------------------------- | -------------------------------------- | ------------------------------------------------------ |
| [**AllCastsInThread**](CastAPI.md#AllCastsInThread) | **Get** /farcaster/all-casts-in-thread | DEPRECATED - Retrieve all casts in a given thread hash |
| [**Cast**](CastAPI.md#Cast)                         | **Get** /farcaster/cast                | DEPRECATED - Retrieve cast for a given hash            |
| [**Casts**](CastAPI.md#Casts)                       | **Get** /farcaster/casts               | DEPRECATED - Retrieve casts for a given user           |
| [**RecentCasts**](CastAPI.md#RecentCasts)           | **Get** /farcaster/recent-casts        | Get Recent Casts                                       |

## AllCastsInThread

> AllCastsInThreadResponse AllCastsInThread(ctx).ApiKey(apiKey).ThreadHash(threadHash).ViewerFid(viewerFid).Execute()

DEPRECATED - Retrieve all casts in a given thread hash

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v1"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	threadHash := "threadHash_example" // string | The hash of the thread to retrieve casts from. (default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.AllCastsInThread(context.Background()).ApiKey(apiKey).ThreadHash(threadHash).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.AllCastsInThread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AllCastsInThread`: AllCastsInThreadResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.AllCastsInThread`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiAllCastsInThreadRequest struct via the builder pattern

| Name           | Type       | Description                                                                  | Notes                                                               |
| -------------- | ---------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| **apiKey**     | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;]                            |
| **threadHash** | **string** | The hash of the thread to retrieve casts from.                               | [default to &quot;0xfe90f9de682273e05b201629ad2338bdcd89b6be&quot;] |
| **viewerFid**  | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                                                      |

### Return type

[**AllCastsInThreadResponse**](AllCastsInThreadResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Cast

> CastResponse Cast(ctx).ApiKey(apiKey).Hash(hash).ViewerFid(viewerFid).Execute()

DEPRECATED - Retrieve cast for a given hash

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v1"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	hash := "hash_example" // string | Cast hash (default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.Cast(context.Background()).ApiKey(apiKey).Hash(hash).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.Cast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cast`: CastResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.Cast`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                                               |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;]                            |
| **hash**      | **string** | Cast hash                                                                    | [default to &quot;0xfe90f9de682273e05b201629ad2338bdcd89b6be&quot;] |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                                                      |

### Return type

[**CastResponse**](CastResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Casts

> CastsResponse Casts(ctx).ApiKey(apiKey).Fid(fid).ParentUrl(parentUrl).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

DEPRECATED - Retrieve casts for a given user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v1"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(56) // int32 | fid of a user (default to 3)
	parentUrl := "https://ethereum.org" // string | A cast can be part of a certain channel. The channel is identified by `parent_url`. All casts in the channel ladder up to the same parent_url. (optional)
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.Casts(context.Background()).ApiKey(apiKey).Fid(fid).ParentUrl(parentUrl).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.Casts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Casts`: CastsResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.Casts`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastsRequest struct via the builder pattern

| Name          | Type       | Description                                                                                                                                              | Notes                                    |
| ------------- | ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                                                                                     | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | fid of a user                                                                                                                                            | [default to 3]                           |
| **parentUrl** | **string** | A cast can be part of a certain channel. The channel is identified by &#x60;parent_url&#x60;. All casts in the channel ladder up to the same parent_url. |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information.                                                                             | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                                                                                                      | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                                                                                                                                       |

### Return type

[**CastsResponse**](CastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RecentCasts

> RecentCastsResponse RecentCasts(ctx).ApiKey(apiKey).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Get Recent Casts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v1"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.RecentCasts(context.Background()).ApiKey(apiKey).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.RecentCasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecentCasts`: RecentCastsResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.RecentCasts`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRecentCastsRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 100)                          | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**RecentCastsResponse**](RecentCastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
