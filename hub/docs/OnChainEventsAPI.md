# \OnChainEventsAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                                                         | HTTP request                                | Description                                           |
| ---------------------------------------------------------------------------------------------- | ------------------------------------------- | ----------------------------------------------------- |
| [**GetOnChainIdRegistrationByAddress**](OnChainEventsAPI.md#GetOnChainIdRegistrationByAddress) | **Get** /v1/onChainIdRegistryEventByAddress | Get an on chain ID Registry Event for a given Address |
| [**ListOnChainEventsByFid**](OnChainEventsAPI.md#ListOnChainEventsByFid)                       | **Get** /v1/onChainEventsByFid              | Get a list of on-chain events provided by an FID      |
| [**ListOnChainSignersByFid**](OnChainEventsAPI.md#ListOnChainSignersByFid)                     | **Get** /v1/onChainSignersByFid             | Get a list of signers provided by an FID              |

## GetOnChainIdRegistrationByAddress

> OnChainEventIdRegister GetOnChainIdRegistrationByAddress(ctx).ApiKey(apiKey).Address(address).Execute()

Get an on chain ID Registry Event for a given Address

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
	address := "address_example" // string | The ETH address being requested (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEventsAPI.GetOnChainIdRegistrationByAddress(context.Background()).ApiKey(apiKey).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEventsAPI.GetOnChainIdRegistrationByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOnChainIdRegistrationByAddress`: OnChainEventIdRegister
	fmt.Fprintf(os.Stdout, "Response from `OnChainEventsAPI.GetOnChainIdRegistrationByAddress`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetOnChainIdRegistrationByAddressRequest struct via the builder pattern

| Name        | Type       | Description                          | Notes                                    |
| ----------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**  | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **address** | **string** | The ETH address being requested      |

### Return type

[**OnChainEventIdRegister**](OnChainEventIdRegister.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOnChainEventsByFid

> ListOnChainEventsByFid200Response ListOnChainEventsByFid(ctx).ApiKey(apiKey).Fid(fid).EventType(eventType).Execute()

Get a list of on-chain events provided by an FID

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
	eventType := openapiclient.OnChainEventType("EVENT_TYPE_SIGNER") // OnChainEventType | The numeric of string value of the event type being requested. (optional) (default to "EVENT_TYPE_SIGNER")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEventsAPI.ListOnChainEventsByFid(context.Background()).ApiKey(apiKey).Fid(fid).EventType(eventType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEventsAPI.ListOnChainEventsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOnChainEventsByFid`: ListOnChainEventsByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `OnChainEventsAPI.ListOnChainEventsByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListOnChainEventsByFidRequest struct via the builder pattern

| Name          | Type                                        | Description                                                    | Notes                                      |
| ------------- | ------------------------------------------- | -------------------------------------------------------------- | ------------------------------------------ |
| **apiKey**    | **string**                                  | API key required for authentication.                           | [default to &quot;NEYNAR_API_DOCS&quot;]   |
| **fid**       | **int32**                                   | The FID being requested                                        |
| **eventType** | [**OnChainEventType**](OnChainEventType.md) | The numeric of string value of the event type being requested. | [default to &quot;EVENT_TYPE_SIGNER&quot;] |

### Return type

[**ListOnChainEventsByFid200Response**](ListOnChainEventsByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListOnChainSignersByFid

> ListOnChainSignersByFid200Response ListOnChainSignersByFid(ctx).ApiKey(apiKey).Fid(fid).Signer(signer).Execute()

Get a list of signers provided by an FID

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
	signer := "0x0852c07b5695ff94138b025e3f9b4788e06133f04e254f0ea0eb85a06e999cdd" // string | The optional key of signer (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OnChainEventsAPI.ListOnChainSignersByFid(context.Background()).ApiKey(apiKey).Fid(fid).Signer(signer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OnChainEventsAPI.ListOnChainSignersByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOnChainSignersByFid`: ListOnChainSignersByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `OnChainEventsAPI.ListOnChainSignersByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListOnChainSignersByFidRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | The FID being requested              |
| **signer** | **string** | The optional key of signer           |

### Return type

[**ListOnChainSignersByFid200Response**](ListOnChainSignersByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
