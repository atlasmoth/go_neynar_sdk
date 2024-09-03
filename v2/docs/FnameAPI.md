# \FnameAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                 | HTTP request                          | Description                         |
| ------------------------------------------------------ | ------------------------------------- | ----------------------------------- |
| [**FnameAvailability**](FnameAPI.md#FnameAvailability) | **Get** /farcaster/fname/availability | Check if a given fname is available |

## FnameAvailability

> FnameAvailabilityResponse FnameAvailability(ctx).ApiKey(apiKey).Fname(fname).Execute()

Check if a given fname is available

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
	fname := "farcaster" // string |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FnameAPI.FnameAvailability(context.Background()).ApiKey(apiKey).Fname(fname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FnameAPI.FnameAvailability``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FnameAvailability`: FnameAvailabilityResponse
	fmt.Fprintf(os.Stdout, "Response from `FnameAPI.FnameAvailability`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFnameAvailabilityRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fname**  | **string** |                                      |

### Return type

[**FnameAvailabilityResponse**](FnameAvailabilityResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
