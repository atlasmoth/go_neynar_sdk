# \STPAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                               | HTTP request                    | Description                                               |
| ---------------------------------------------------- | ------------------------------- | --------------------------------------------------------- |
| [**SubscriptionCheck**](STPAPI.md#SubscriptionCheck) | **Get** /stp/subscription_check | Check if a wallet address is subscribed to a STP contract |

## SubscriptionCheck

> map[string]SubscriptionStatus SubscriptionCheck(ctx).ApiKey(apiKey).Addresses(addresses).ContractAddress(contractAddress).ChainId(chainId).Execute()

Check if a wallet address is subscribed to a STP contract

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/stp"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	addresses := "0xedd3783e8c7c52b80cfbd026a63c207edc9cbee7,0x5a927ac639636e534b678e81768ca19e2c6280b7" // string | Comma separated list of Ethereum addresses, up to 350 at a time (optional)
	contractAddress := "0x76ad4cb9ac51c09f4d9c2cadcea75c9fa9074e5b" // string | Ethereum address of the STP contract (optional)
	chainId := "8453" // string | Chain ID of the STP contract (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.STPAPI.SubscriptionCheck(context.Background()).ApiKey(apiKey).Addresses(addresses).ContractAddress(contractAddress).ChainId(chainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `STPAPI.SubscriptionCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCheck`: map[string]SubscriptionStatus
	fmt.Fprintf(os.Stdout, "Response from `STPAPI.SubscriptionCheck`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCheckRequest struct via the builder pattern

| Name                | Type       | Description                                                     | Notes                                    |
| ------------------- | ---------- | --------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**          | **string** | API key required for authentication.                            | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **addresses**       | **string** | Comma separated list of Ethereum addresses, up to 350 at a time |
| **contractAddress** | **string** | Ethereum address of the STP contract                            |
| **chainId**         | **string** | Chain ID of the STP contract                                    |

### Return type

[**map[string]SubscriptionStatus**](SubscriptionStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
