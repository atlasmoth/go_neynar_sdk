# \SignerAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                                                                    | HTTP request                                            | Description                                    |
| --------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- | ---------------------------------------------- |
| [**CreateSigner**](SignerAPI.md#CreateSigner)                                                             | **Post** /farcaster/signer                              | Creates a signer and returns the signer status |
| [**DeveloperManagedSigner**](SignerAPI.md#DeveloperManagedSigner)                                         | **Get** /farcaster/signer/developer_managed             | Fetches the status of a signer by public key   |
| [**FetchAuthorizationUrl**](SignerAPI.md#FetchAuthorizationUrl)                                           | **Get** /farcaster/login/authorize                      | Fetch authorization url                        |
| [**PublishMessage**](SignerAPI.md#PublishMessage)                                                         | **Post** /farcaster/message                             | Publish a message to farcaster                 |
| [**RegisterSignedKey**](SignerAPI.md#RegisterSignedKey)                                                   | **Post** /farcaster/signer/signed_key                   | Register Signed Key                            |
| [**RegisterSignedKeyForDeveloperManagedSigner**](SignerAPI.md#RegisterSignedKeyForDeveloperManagedSigner) | **Post** /farcaster/signer/developer_managed/signed_key | Registers Signed Key                           |
| [**Signer**](SignerAPI.md#Signer)                                                                         | **Get** /farcaster/signer                               | Fetches the status of a signer                 |

## CreateSigner

> Signer CreateSigner(ctx).ApiKey(apiKey).Execute()

Creates a signer and returns the signer status

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.CreateSigner(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.CreateSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSigner`: Signer
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.CreateSigner`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSignerRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeveloperManagedSigner

> DeveloperManagedSigner DeveloperManagedSigner(ctx).ApiKey(apiKey).PublicKey(publicKey).Execute()

Fetches the status of a signer by public key

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	publicKey := "publicKey_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.DeveloperManagedSigner(context.Background()).ApiKey(apiKey).PublicKey(publicKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.DeveloperManagedSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeveloperManagedSigner`: DeveloperManagedSigner
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.DeveloperManagedSigner`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeveloperManagedSignerRequest struct via the builder pattern

| Name          | Type       | Description                          | Notes                                    |
| ------------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **publicKey** | **string** |                                      |

### Return type

[**DeveloperManagedSigner**](DeveloperManagedSigner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FetchAuthorizationUrl

> AuthorizationUrlResponse FetchAuthorizationUrl(ctx).ApiKey(apiKey).ClientId(clientId).ResponseType(responseType).Execute()

Fetch authorization url

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	clientId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
	responseType := openapiclient.AuthorizationUrlResponseType("code") // AuthorizationUrlResponseType |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.FetchAuthorizationUrl(context.Background()).ApiKey(apiKey).ClientId(clientId).ResponseType(responseType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.FetchAuthorizationUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchAuthorizationUrl`: AuthorizationUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.FetchAuthorizationUrl`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFetchAuthorizationUrlRequest struct via the builder pattern

| Name             | Type                                                                | Description                          | Notes                                    |
| ---------------- | ------------------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**       | **string**                                                          | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **clientId**     | **string**                                                          |                                      |
| **responseType** | [**AuthorizationUrlResponseType**](AuthorizationUrlResponseType.md) |                                      |

### Return type

[**AuthorizationUrlResponse**](AuthorizationUrlResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PublishMessage

> map[string]interface{} PublishMessage(ctx).ApiKey(apiKey).Body(body).Execute()

Publish a message to farcaster

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.PublishMessage(context.Background()).ApiKey(apiKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.PublishMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishMessage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.PublishMessage`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPublishMessageRequest struct via the builder pattern

| Name       | Type                       | Description                          | Notes                                    |
| ---------- | -------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string**                 | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **body**   | **map[string]interface{}** |                                      |

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RegisterSignedKey

> Signer RegisterSignedKey(ctx).ApiKey(apiKey).RegisterSignerKeyReqBody(registerSignerKeyReqBody).Execute()

Register Signed Key

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	registerSignerKeyReqBody := *openapiclient.NewRegisterSignerKeyReqBody() // RegisterSignerKeyReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.RegisterSignedKey(context.Background()).ApiKey(apiKey).RegisterSignerKeyReqBody(registerSignerKeyReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.RegisterSignedKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterSignedKey`: Signer
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.RegisterSignedKey`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSignedKeyRequest struct via the builder pattern

| Name                         | Type                                                        | Description                          | Notes                                    |
| ---------------------------- | ----------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                   | **string**                                                  | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **registerSignerKeyReqBody** | [**RegisterSignerKeyReqBody**](RegisterSignerKeyReqBody.md) |                                      |

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RegisterSignedKeyForDeveloperManagedSigner

> DeveloperManagedSigner RegisterSignedKeyForDeveloperManagedSigner(ctx).ApiKey(apiKey).RegisterDeveloperManagedSignedKeyReqBody(registerDeveloperManagedSignedKeyReqBody).Execute()

Registers Signed Key

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	registerDeveloperManagedSignedKeyReqBody := *openapiclient.NewRegisterDeveloperManagedSignedKeyReqBody() // RegisterDeveloperManagedSignedKeyReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.RegisterSignedKeyForDeveloperManagedSigner(context.Background()).ApiKey(apiKey).RegisterDeveloperManagedSignedKeyReqBody(registerDeveloperManagedSignedKeyReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.RegisterSignedKeyForDeveloperManagedSigner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterSignedKeyForDeveloperManagedSigner`: DeveloperManagedSigner
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.RegisterSignedKeyForDeveloperManagedSigner`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSignedKeyForDeveloperManagedSignerRequest struct via the builder pattern

| Name                                         | Type                                                                                        | Description                          | Notes                                    |
| -------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                                   | **string**                                                                                  | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **registerDeveloperManagedSignedKeyReqBody** | [**RegisterDeveloperManagedSignedKeyReqBody**](RegisterDeveloperManagedSignedKeyReqBody.md) |                                      |

### Return type

[**DeveloperManagedSigner**](DeveloperManagedSigner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Signer

> Signer Signer(ctx).ApiKey(apiKey).SignerUuid(signerUuid).Execute()

Fetches the status of a signer

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
	apiKey := "apiKey_example" // string | API key required for authentication. (optional) (default to "NEYNAR_API_DOCS")
	signerUuid := "19d0c5fd-9b33-4a48-a0e2-bc7b0555baec" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SignerAPI.Signer(context.Background()).ApiKey(apiKey).SignerUuid(signerUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SignerAPI.Signer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Signer`: Signer
	fmt.Fprintf(os.Stdout, "Response from `SignerAPI.Signer`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSignerRequest struct via the builder pattern

| Name           | Type       | Description                          | Notes                                    |
| -------------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**     | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **signerUuid** | **string** |                                      |

### Return type

[**Signer**](Signer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
