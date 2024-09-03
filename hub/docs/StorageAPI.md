# \StorageAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                           | HTTP request                   | Description                      |
| ---------------------------------------------------------------- | ------------------------------ | -------------------------------- |
| [**GetStorageLimitsByFid**](StorageAPI.md#GetStorageLimitsByFid) | **Get** /v1/storageLimitsByFid | Get an FID&#39;s storage limits. |

## GetStorageLimitsByFid

> StorageLimitsResponse GetStorageLimitsByFid(ctx).ApiKey(apiKey).Fid(fid).Execute()

Get an FID's storage limits.

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
	fid := int32(56) // int32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.GetStorageLimitsByFid(context.Background()).ApiKey(apiKey).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageLimitsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageLimitsByFid`: StorageLimitsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageLimitsByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageLimitsByFidRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  |                                      |

### Return type

[**StorageLimitsResponse**](StorageLimitsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
