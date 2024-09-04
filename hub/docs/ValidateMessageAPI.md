# \ValidateMessageAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                       | HTTP request                 | Description                                                |
| ------------------------------------------------------------ | ---------------------------- | ---------------------------------------------------------- |
| [**ValidateMessage**](ValidateMessageAPI.md#ValidateMessage) | **Post** /v1/validateMessage | Validate a signed protobuf-serialized message with the Hub |

## ValidateMessage

> ValidateMessageResponse ValidateMessage(ctx).ApiKey(apiKey).Body(body).Execute()

Validate a signed protobuf-serialized message with the Hub

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	body := os.NewFile(1234, "some_file") // *os.File | *  A Message is a delta operation on the Farcaster network. The message protobuf is an envelope  that wraps a MessageData object and contains a hash and signature which can verify its authenticity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidateMessageAPI.ValidateMessage(context.Background()).ApiKey(apiKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidateMessageAPI.ValidateMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateMessage`: ValidateMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ValidateMessageAPI.ValidateMessage`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMessageRequest struct via the builder pattern

| Name       | Type          | Description                                                                                                                                                                                          | Notes                                    |
| ---------- | ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey** | **string**    | API key required for authentication.                                                                                                                                                                 | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **body**   | **\*os.File** | \* A Message is a delta operation on the Farcaster network. The message protobuf is an envelope that wraps a MessageData object and contains a hash and signature which can verify its authenticity. |

### Return type

[**ValidateMessageResponse**](ValidateMessageResponse.md)

### Authorization

[usernamePassword](../README.md#usernamePassword)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
