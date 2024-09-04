# \VerificationsAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                                   | HTTP request                   | Description                                    |
| ------------------------------------------------------------------------ | ------------------------------ | ---------------------------------------------- |
| [**ListVerificationsByFid**](VerificationsAPI.md#ListVerificationsByFid) | **Get** /v1/verificationsByFid | Get a list of verifications provided by an FID |

## ListVerificationsByFid

> ListVerificationsByFid200Response ListVerificationsByFid(ctx).ApiKey(apiKey).Fid(fid).Address(address).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get a list of verifications provided by an FID

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
	fid := int32(56) // int32 | The FID being requested (optional)
	address := "address_example" // string | The optional ETH address to filter by (optional)
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationsAPI.ListVerificationsByFid(context.Background()).ApiKey(apiKey).Fid(fid).Address(address).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationsAPI.ListVerificationsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVerificationsByFid`: ListVerificationsByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `VerificationsAPI.ListVerificationsByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListVerificationsByFidRequest struct via the builder pattern

| Name          | Type       | Description                                                                                                             | Notes                                    |
| ------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | The FID being requested                                                                                                 |
| **address**   | **string** | The optional ETH address to filter by                                                                                   |
| **pageSize**  | **int32**  | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**   | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListVerificationsByFid200Response**](ListVerificationsByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
