# \VerificationAPI

All URIs are relative to *https://api.neynar.com/v1*

| Method                                                          | HTTP request                            | Description                                             |
| --------------------------------------------------------------- | --------------------------------------- | ------------------------------------------------------- |
| [**UserByVerification**](VerificationAPI.md#UserByVerification) | **Get** /farcaster/user-by-verification | DEPRECATED - Retrieve user for a given ethereum address |
| [**Verifications**](VerificationAPI.md#Verifications)           | **Get** /farcaster/verifications        | DEPRECATED - Retrieve verifications for a given FID     |

## UserByVerification

> UserResponse UserByVerification(ctx).ApiKey(apiKey).Address(address).Execute()

DEPRECATED - Retrieve user for a given ethereum address

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v1"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	address := "address_example" // string |  (optional) (default to "0x5A927Ac639636E534b678e81768CA19e2C6280B7")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.UserByVerification(context.Background()).ApiKey(apiKey).Address(address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.UserByVerification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserByVerification`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.UserByVerification`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserByVerificationRequest struct via the builder pattern

| Name        | Type       | Description                          | Notes                                                               |
| ----------- | ---------- | ------------------------------------ | ------------------------------------------------------------------- |
| **apiKey**  | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;]                            |
| **address** | **string** |                                      | [default to &quot;0x5A927Ac639636E534b678e81768CA19e2C6280B7&quot;] |

### Return type

[**UserResponse**](UserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Verifications

> VerificationResponse Verifications(ctx).ApiKey(apiKey).Fid(fid).Execute()

DEPRECATED - Retrieve verifications for a given FID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v1"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	fid := int32(56) // int32 | FID of the user (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VerificationAPI.Verifications(context.Background()).ApiKey(apiKey).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VerificationAPI.Verifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Verifications`: VerificationResponse
	fmt.Fprintf(os.Stdout, "Response from `VerificationAPI.Verifications`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiVerificationsRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | FID of the user                      | [default to 3]                           |

### Return type

[**VerificationResponse**](VerificationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
