# \MuteAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                  | HTTP request                 | Description                    |
| --------------------------------------- | ---------------------------- | ------------------------------ |
| [**AddMute**](MuteAPI.md#AddMute)       | **Post** /farcaster/mute     | Adds a mute for a fid          |
| [**DeleteMute**](MuteAPI.md#DeleteMute) | **Delete** /farcaster/mute   | Deletes a mute for a fid       |
| [**MuteList**](MuteAPI.md#MuteList)     | **Get** /farcaster/mute/list | Get fids that a user has muted |

## AddMute

> MuteResponse AddMute(ctx).ApiKey(apiKey).MuteReqBody(muteReqBody).Execute()

Adds a mute for a fid

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
	muteReqBody := *openapiclient.NewMuteReqBody(int32(123), int32(123)) // MuteReqBody |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteAPI.AddMute(context.Background()).ApiKey(apiKey).MuteReqBody(muteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MuteAPI.AddMute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMute`: MuteResponse
	fmt.Fprintf(os.Stdout, "Response from `MuteAPI.AddMute`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiAddMuteRequest struct via the builder pattern

| Name            | Type                              | Description                          | Notes                                    |
| --------------- | --------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**      | **string**                        | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **muteReqBody** | [**MuteReqBody**](MuteReqBody.md) |                                      |

### Return type

[**MuteResponse**](MuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteMute

> MuteResponse DeleteMute(ctx).ApiKey(apiKey).MuteReqBody(muteReqBody).Execute()

Deletes a mute for a fid

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
	muteReqBody := *openapiclient.NewMuteReqBody(int32(123), int32(123)) // MuteReqBody |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteAPI.DeleteMute(context.Background()).ApiKey(apiKey).MuteReqBody(muteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MuteAPI.DeleteMute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMute`: MuteResponse
	fmt.Fprintf(os.Stdout, "Response from `MuteAPI.DeleteMute`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMuteRequest struct via the builder pattern

| Name            | Type                              | Description                          | Notes                                    |
| --------------- | --------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**      | **string**                        | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **muteReqBody** | [**MuteReqBody**](MuteReqBody.md) |                                      |

### Return type

[**MuteResponse**](MuteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MuteList

> MuteListResponse MuteList(ctx).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()

Get fids that a user has muted

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
	fid := int32(194) // int32 | The user's fid (identifier)
	limit := int32(20) // int32 | Number of results to retrieve (default 20, max 100). (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MuteAPI.MuteList(context.Background()).ApiKey(apiKey).Fid(fid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MuteAPI.MuteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteList`: MuteListResponse
	fmt.Fprintf(os.Stdout, "Response from `MuteAPI.MuteList`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiMuteListRequest struct via the builder pattern

| Name       | Type       | Description                                          | Notes                                    |
| ---------- | ---------- | ---------------------------------------------------- | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication.                 | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | The user&#39;s fid (identifier)                      |
| **limit**  | **int32**  | Number of results to retrieve (default 20, max 100). | [default to 20]                          |
| **cursor** | **string** | Pagination cursor.                                   |

### Return type

[**MuteListResponse**](MuteListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
