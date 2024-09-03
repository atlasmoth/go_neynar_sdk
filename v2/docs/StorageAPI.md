# \StorageAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                     | HTTP request                           | Description                                  |
| ---------------------------------------------------------- | -------------------------------------- | -------------------------------------------- |
| [**BuyStorage**](StorageAPI.md#BuyStorage)                 | **Post** /farcaster/storage/buy        | Buy storage for an fid                       |
| [**StorageAllocations**](StorageAPI.md#StorageAllocations) | **Get** /farcaster/storage/allocations | Fetches storage allocations for a given user |
| [**StorageUsage**](StorageAPI.md#StorageUsage)             | **Get** /farcaster/storage/usage       | Fetches storage usage for a given user       |

## BuyStorage

> StorageAllocationsResponse BuyStorage(ctx).ApiKey(apiKey).BuyStorageReqBody(buyStorageReqBody).Execute()

Buy storage for an fid

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	buyStorageReqBody := *openapiclient.NewBuyStorageReqBody(int32(1)) // BuyStorageReqBody |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.BuyStorage(context.Background()).ApiKey(apiKey).BuyStorageReqBody(buyStorageReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.BuyStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuyStorage`: StorageAllocationsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.BuyStorage`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiBuyStorageRequest struct via the builder pattern

| Name                  | Type                                          | Description                          | Notes                                    |
| --------------------- | --------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**            | **string**                                    | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **buyStorageReqBody** | [**BuyStorageReqBody**](BuyStorageReqBody.md) |                                      |

### Return type

[**StorageAllocationsResponse**](StorageAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## StorageAllocations

> StorageAllocationsResponse StorageAllocations(ctx).ApiKey(apiKey).Fid(fid).Execute()

Fetches storage allocations for a given user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(3) // int32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageAllocations(context.Background()).ApiKey(apiKey).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageAllocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageAllocations`: StorageAllocationsResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageAllocations`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiStorageAllocationsRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  |                                      |

### Return type

[**StorageAllocationsResponse**](StorageAllocationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## StorageUsage

> StorageUsageResponse StorageUsage(ctx).ApiKey(apiKey).Fid(fid).Execute()

Fetches storage usage for a given user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(3) // int32 |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StorageAPI.StorageUsage(context.Background()).ApiKey(apiKey).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.StorageUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorageUsage`: StorageUsageResponse
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.StorageUsage`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiStorageUsageRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  |                                      |

### Return type

[**StorageUsageResponse**](StorageUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
