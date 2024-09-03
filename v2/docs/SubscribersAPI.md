# \SubscribersAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                             | HTTP request                                  | Description                                 |
| ------------------------------------------------------------------ | --------------------------------------------- | ------------------------------------------- |
| [**SubscribedTo**](SubscribersAPI.md#SubscribedTo)                 | **Get** /farcaster/user/subscribed_to         | Fetch what a given fid is subscribed to     |
| [**Subscribers**](SubscribersAPI.md#Subscribers)                   | **Get** /farcaster/user/subscribers           | Fetch subscribers for a given fid           |
| [**SubscriptionsCreated**](SubscribersAPI.md#SubscriptionsCreated) | **Get** /farcaster/user/subscriptions_created | Fetch created subscriptions for a given fid |

## SubscribedTo

> SubscribedToResponse SubscribedTo(ctx).ApiKey(apiKey).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()

Fetch what a given fid is subscribed to

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
	fid := int32(3206) // int32 |
	subscriptionProvider := openapiclient.SubscriptionProvider("fabric_stp") // SubscriptionProvider |
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.SubscribedTo(context.Background()).ApiKey(apiKey).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.SubscribedTo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribedTo`: SubscribedToResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.SubscribedTo`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribedToRequest struct via the builder pattern

| Name                     | Type                                                | Description                          | Notes                                    |
| ------------------------ | --------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**               | **string**                                          | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**                  | **int32**                                           |                                      |
| **subscriptionProvider** | [**SubscriptionProvider**](SubscriptionProvider.md) |                                      |
| **viewerFid**            | **int32**                                           |                                      |

### Return type

[**SubscribedToResponse**](SubscribedToResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Subscribers

> SubscribersResponse Subscribers(ctx).ApiKey(apiKey).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()

Fetch subscribers for a given fid

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
	fid := int32(3206) // int32 |
	subscriptionProvider := openapiclient.SubscriptionProviders("fabric_stp") // SubscriptionProviders |
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.Subscribers(context.Background()).ApiKey(apiKey).Fid(fid).SubscriptionProvider(subscriptionProvider).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.Subscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Subscribers`: SubscribersResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.Subscribers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribersRequest struct via the builder pattern

| Name                     | Type                                                  | Description                          | Notes                                    |
| ------------------------ | ----------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**               | **string**                                            | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**                  | **int32**                                             |                                      |
| **subscriptionProvider** | [**SubscriptionProviders**](SubscriptionProviders.md) |                                      |
| **viewerFid**            | **int32**                                             |                                      |

### Return type

[**SubscribersResponse**](SubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SubscriptionsCreated

> SubscriptionsResponse SubscriptionsCreated(ctx).ApiKey(apiKey).Fid(fid).SubscriptionProvider(subscriptionProvider).Execute()

Fetch created subscriptions for a given fid

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
	fid := int32(528) // int32 |
	subscriptionProvider := openapiclient.SubscriptionProvider("fabric_stp") // SubscriptionProvider |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscribersAPI.SubscriptionsCreated(context.Background()).ApiKey(apiKey).Fid(fid).SubscriptionProvider(subscriptionProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscribersAPI.SubscriptionsCreated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionsCreated`: SubscriptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubscribersAPI.SubscriptionsCreated`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsCreatedRequest struct via the builder pattern

| Name                     | Type                                                | Description                          | Notes                                    |
| ------------------------ | --------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**               | **string**                                          | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**                  | **int32**                                           |                                      |
| **subscriptionProvider** | [**SubscriptionProvider**](SubscriptionProvider.md) |                                      |

### Return type

[**SubscriptionsResponse**](SubscriptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
