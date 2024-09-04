# \InfoAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                            | HTTP request     | Description  |
| --------------------------------- | ---------------- | ------------ |
| [**GetInfo**](InfoAPI.md#GetInfo) | **Get** /v1/info | Sync Methods |

## GetInfo

> HubInfoResponse GetInfo(ctx).ApiKey(apiKey).Dbstats(dbstats).Execute()

Sync Methods

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
	dbstats := true // bool | Whether to return DB stats (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InfoAPI.GetInfo(context.Background()).ApiKey(apiKey).Dbstats(dbstats).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.GetInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInfo`: HubInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.GetInfo`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoRequest struct via the builder pattern

| Name        | Type       | Description                          | Notes                                    |
| ----------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**  | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **dbstats** | **bool**   | Whether to return DB stats           |

### Return type

[**HubInfoResponse**](HubInfoResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
