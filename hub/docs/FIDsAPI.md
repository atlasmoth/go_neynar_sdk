# \FIDsAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                              | HTTP request     | Description                |
| ----------------------------------- | ---------------- | -------------------------- |
| [**ListFids**](FIDsAPI.md#ListFids) | **Get** /v1/fids | Get a list of all the FIDs |

## ListFids

> FidsResponse ListFids(ctx).ApiKey(apiKey).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get a list of all the FIDs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/hub"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FIDsAPI.ListFids(context.Background()).ApiKey(apiKey).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FIDsAPI.ListFids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFids`: FidsResponse
	fmt.Fprintf(os.Stdout, "Response from `FIDsAPI.ListFids`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListFidsRequest struct via the builder pattern

| Name          | Type       | Description                                                                                                             | Notes                                    |
| ------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **pageSize**  | **int32**  | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**   | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**FidsResponse**](FidsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
