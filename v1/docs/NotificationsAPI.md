# \NotificationsAPI

All URIs are relative to *https://api.neynar.com/v1*

| Method                                                             | HTTP request                             | Description               |
| ------------------------------------------------------------------ | ---------------------------------------- | ------------------------- |
| [**MentionsAndReplies**](NotificationsAPI.md#MentionsAndReplies)   | **Get** /farcaster/mentions-and-replies  | Get mentions and replies  |
| [**ReactionsAndRecasts**](NotificationsAPI.md#ReactionsAndRecasts) | **Get** /farcaster/reactions-and-recasts | Get reactions and recasts |

## MentionsAndReplies

> MentionsAndRepliesResponse MentionsAndReplies(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Cursor(cursor).Execute()

Get mentions and replies

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
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.MentionsAndReplies(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.MentionsAndReplies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MentionsAndReplies`: MentionsAndRepliesResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.MentionsAndReplies`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMentionsAndRepliesRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | fid of a user                                                                | [default to 3]                           |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**MentionsAndRepliesResponse**](MentionsAndRepliesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ReactionsAndRecasts

> ReactionsAndRecastsResponse ReactionsAndRecasts(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Get reactions and recasts

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
	fid := int32(6131) // int32 | fid of a user (default to 3)
	viewerFid := int32(3) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsAPI.ReactionsAndRecasts(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ReactionsAndRecasts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReactionsAndRecasts`: ReactionsAndRecastsResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ReactionsAndRecasts`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiReactionsAndRecastsRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | fid of a user                                                                | [default to 3]                           |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**ReactionsAndRecastsResponse**](ReactionsAndRecastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
