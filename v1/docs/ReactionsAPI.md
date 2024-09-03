# \ReactionsAPI

All URIs are relative to *https://api.neynar.com/v1*

| Method                                             | HTTP request                      | Description                                             |
| -------------------------------------------------- | --------------------------------- | ------------------------------------------------------- |
| [**CastLikes**](ReactionsAPI.md#CastLikes)         | **Get** /farcaster/cast-likes     | DEPRECATED - Get all like reactions for a specific cast |
| [**CastReactions**](ReactionsAPI.md#CastReactions) | **Get** /farcaster/cast-reactions | DEPRECATED - Get all reactions for a specific cast      |
| [**CastRecasters**](ReactionsAPI.md#CastRecasters) | **Get** /farcaster/cast-recasters | DEPRECATED - Get all recasters for a specific cast      |

## CastLikes

> CastLikesResponse CastLikes(ctx).ApiKey(apiKey).CastHash(castHash).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

DEPRECATED - Get all like reactions for a specific cast

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
	castHash := "castHash_example" // string | Cast hash (default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.CastLikes(context.Background()).ApiKey(apiKey).CastHash(castHash).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.CastLikes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastLikes`: CastLikesResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.CastLikes`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastLikesRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                                               |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;]                            |
| **castHash**  | **string** | Cast hash                                                                    | [default to &quot;0xfe90f9de682273e05b201629ad2338bdcd89b6be&quot;] |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                                                      |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                                                     |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**CastLikesResponse**](CastLikesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CastReactions

> CastReactionsResponse CastReactions(ctx).ApiKey(apiKey).CastHash(castHash).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

DEPRECATED - Get all reactions for a specific cast

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
	castHash := "castHash_example" // string | Cast hash (default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.CastReactions(context.Background()).ApiKey(apiKey).CastHash(castHash).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.CastReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastReactions`: CastReactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.CastReactions`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastReactionsRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                                               |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;]                            |
| **castHash**  | **string** | Cast hash                                                                    | [default to &quot;0xfe90f9de682273e05b201629ad2338bdcd89b6be&quot;] |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                                                      |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                                                     |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**CastReactionsResponse**](CastReactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CastRecasters

> CastRecasterResponse CastRecasters(ctx).ApiKey(apiKey).CastHash(castHash).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

DEPRECATED - Get all recasters for a specific cast

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
	castHash := "castHash_example" // string | Cast hash (default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.CastRecasters(context.Background()).ApiKey(apiKey).CastHash(castHash).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.CastRecasters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastRecasters`: CastRecasterResponse
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.CastRecasters`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastRecastersRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                                               |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;]                            |
| **castHash**  | **string** | Cast hash                                                                    | [default to &quot;0xfe90f9de682273e05b201629ad2338bdcd89b6be&quot;] |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                                                      |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                                                     |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**CastRecasterResponse**](CastRecasterResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
