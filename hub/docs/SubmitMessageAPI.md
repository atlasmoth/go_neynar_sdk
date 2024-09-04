# \SubmitMessageAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                 | HTTP request               | Description                                            |
| ------------------------------------------------------ | -------------------------- | ------------------------------------------------------ |
| [**SubmitMessage**](SubmitMessageAPI.md#SubmitMessage) | **Post** /v1/submitMessage | Submit a signed protobuf-serialized message to the Hub |

## SubmitMessage

> Message SubmitMessage(ctx).Body(body).Execute()

Submit a signed protobuf-serialized message to the Hub

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
	body := os.NewFile(1234, "some_file") // *os.File | *  A Message is a delta operation on the Farcaster network. The message protobuf is an envelope  that wraps a MessageData object and contains a hash and signature which can verify its authenticity. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubmitMessageAPI.SubmitMessage(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubmitMessageAPI.SubmitMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitMessage`: Message
	fmt.Fprintf(os.Stdout, "Response from `SubmitMessageAPI.SubmitMessage`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMessageRequest struct via the builder pattern

| Name     | Type          | Description                                                                                                                                                                                          | Notes |
| -------- | ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **body** | **\*os.File** | \* A Message is a delta operation on the Farcaster network. The message protobuf is an envelope that wraps a MessageData object and contains a hash and signature which can verify its authenticity. |

### Return type

[**Message**](Message.md)

### Authorization

[usernamePassword](../README.md#usernamePassword)

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
