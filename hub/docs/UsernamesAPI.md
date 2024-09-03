# \UsernamesAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                                 | HTTP request                    | Description                                           |
| ---------------------------------------------------------------------- | ------------------------------- | ----------------------------------------------------- |
| [**GetUsernameProof**](UsernamesAPI.md#GetUsernameProof)               | **Get** /v1/userNameProofByName | Get an proof for a username by the Farcaster username |
| [**ListUsernameProofsByFid**](UsernamesAPI.md#ListUsernameProofsByFid) | **Get** /v1/userNameProofsByFid | Get a list of proofs provided by an FID               |

## GetUsernameProof

> UserNameProof GetUsernameProof(ctx).ApiKey(apiKey).Name(name).Execute()

Get an proof for a username by the Farcaster username

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
	name := "gavi" // string | The Farcaster username or ENS address

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsernamesAPI.GetUsernameProof(context.Background()).ApiKey(apiKey).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsernamesAPI.GetUsernameProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsernameProof`: UserNameProof
	fmt.Fprintf(os.Stdout, "Response from `UsernamesAPI.GetUsernameProof`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsernameProofRequest struct via the builder pattern

| Name       | Type       | Description                           | Notes                                    |
| ---------- | ---------- | ------------------------------------- | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication.  | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **name**   | **string** | The Farcaster username or ENS address |

### Return type

[**UserNameProof**](UserNameProof.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUsernameProofsByFid

> UsernameProofsResponse ListUsernameProofsByFid(ctx).ApiKey(apiKey).Fid(fid).Execute()

Get a list of proofs provided by an FID

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
	fid := int32(56) // int32 | The FID being requested

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsernamesAPI.ListUsernameProofsByFid(context.Background()).ApiKey(apiKey).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsernamesAPI.ListUsernameProofsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsernameProofsByFid`: UsernameProofsResponse
	fmt.Fprintf(os.Stdout, "Response from `UsernamesAPI.ListUsernameProofsByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListUsernameProofsByFidRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | The FID being requested              |

### Return type

[**UsernameProofsResponse**](UsernameProofsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
