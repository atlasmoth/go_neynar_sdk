# \FollowsAPI

All URIs are relative to *https://api.neynar.com/v1*

| Method                                   | HTTP request                 | Description                        |
| ---------------------------------------- | ---------------------------- | ---------------------------------- |
| [**Followers**](FollowsAPI.md#Followers) | **Get** /farcaster/followers | Gets all followers for a given FID |
| [**Following**](FollowsAPI.md#Following) | **Get** /farcaster/following | Gets all following users of a FID  |

## Followers

> FollowResponse Followers(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Gets all followers for a given FID

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
	fid := int32(6131) // int32 | FID of the user (default to 3)
	viewerFid := int32(194) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.Followers(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.Followers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Followers`: FollowResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.Followers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFollowersRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | FID of the user                                                              | [default to 3]                           |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**FollowResponse**](FollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Following

> FollowResponse Following(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Gets all following users of a FID

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
	fid := int32(6131) // int32 | FID of the user (default to 3)
	viewerFid := int32(194) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.Following(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.Following``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Following`: FollowResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.Following`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFollowingRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | FID of the user                                                              | [default to 3]                           |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**FollowResponse**](FollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
