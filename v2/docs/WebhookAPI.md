# \WebhookAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                                   | HTTP request                    | Description                                   |
| ------------------------------------------------------------------------ | ------------------------------- | --------------------------------------------- |
| [**DeleteWebhook**](WebhookAPI.md#DeleteWebhook)                         | **Delete** /farcaster/webhook   | Delete a webhook                              |
| [**FetchWebhooks**](WebhookAPI.md#FetchWebhooks)                         | **Get** /farcaster/webhook/list | Fetch a list of webhooks associated to a user |
| [**LookupWebhook**](WebhookAPI.md#LookupWebhook)                         | **Get** /farcaster/webhook      | Fetch a webhook                               |
| [**PublishWebhook**](WebhookAPI.md#PublishWebhook)                       | **Post** /farcaster/webhook     | Create a webhook                              |
| [**UpdateWebhook**](WebhookAPI.md#UpdateWebhook)                         | **Put** /farcaster/webhook      | Update a webhook                              |
| [**UpdateWebhookActiveStatus**](WebhookAPI.md#UpdateWebhookActiveStatus) | **Patch** /farcaster/webhook    | Update webhook active status                  |

## DeleteWebhook

> WebhookResponse DeleteWebhook(ctx).ApiKey(apiKey).WebhookDeleteReqBody(webhookDeleteReqBody).Execute()

Delete a webhook

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
	webhookDeleteReqBody := *openapiclient.NewWebhookDeleteReqBody() // WebhookDeleteReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.DeleteWebhook(context.Background()).ApiKey(apiKey).WebhookDeleteReqBody(webhookDeleteReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern

| Name                     | Type                                                | Description                          | Notes                                    |
| ------------------------ | --------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**               | **string**                                          | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **webhookDeleteReqBody** | [**WebhookDeleteReqBody**](WebhookDeleteReqBody.md) |                                      |

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FetchWebhooks

> WebhookListResponse FetchWebhooks(ctx).ApiKey(apiKey).Execute()

Fetch a list of webhooks associated to a user

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
	resp, r, err := apiClient.WebhookAPI.FetchWebhooks(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.FetchWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchWebhooks`: WebhookListResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.FetchWebhooks`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFetchWebhooksRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |

### Return type

[**WebhookListResponse**](WebhookListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LookupWebhook

> WebhookResponse LookupWebhook(ctx).ApiKey(apiKey).WebhookId(webhookId).Execute()

Fetch a webhook

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
	webhookId := "webhookId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.LookupWebhook(context.Background()).ApiKey(apiKey).WebhookId(webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.LookupWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.LookupWebhook`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiLookupWebhookRequest struct via the builder pattern

| Name          | Type       | Description                          | Notes                                    |
| ------------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **webhookId** | **string** |                                      |

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PublishWebhook

> WebhookResponse PublishWebhook(ctx).ApiKey(apiKey).WebhookPostReqBody(webhookPostReqBody).Execute()

Create a webhook

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
	webhookPostReqBody := *openapiclient.NewWebhookPostReqBody() // WebhookPostReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.PublishWebhook(context.Background()).ApiKey(apiKey).WebhookPostReqBody(webhookPostReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.PublishWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.PublishWebhook`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPublishWebhookRequest struct via the builder pattern

| Name                   | Type                                            | Description                          | Notes                                    |
| ---------------------- | ----------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**             | **string**                                      | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **webhookPostReqBody** | [**WebhookPostReqBody**](WebhookPostReqBody.md) |                                      |

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateWebhook

> WebhookResponse UpdateWebhook(ctx).ApiKey(apiKey).WebhookPutReqBody(webhookPutReqBody).Execute()

Update a webhook

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
	webhookPutReqBody := *openapiclient.NewWebhookPutReqBody() // WebhookPutReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdateWebhook(context.Background()).ApiKey(apiKey).WebhookPutReqBody(webhookPutReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhook`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdateWebhook`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookRequest struct via the builder pattern

| Name                  | Type                                          | Description                          | Notes                                    |
| --------------------- | --------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**            | **string**                                    | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **webhookPutReqBody** | [**WebhookPutReqBody**](WebhookPutReqBody.md) |                                      |

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateWebhookActiveStatus

> WebhookResponse UpdateWebhookActiveStatus(ctx).ApiKey(apiKey).WebhookPatchReqBody(webhookPatchReqBody).Execute()

Update webhook active status

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
	webhookPatchReqBody := *openapiclient.NewWebhookPatchReqBody() // WebhookPatchReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhookAPI.UpdateWebhookActiveStatus(context.Background()).ApiKey(apiKey).WebhookPatchReqBody(webhookPatchReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhookAPI.UpdateWebhookActiveStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWebhookActiveStatus`: WebhookResponse
	fmt.Fprintf(os.Stdout, "Response from `WebhookAPI.UpdateWebhookActiveStatus`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebhookActiveStatusRequest struct via the builder pattern

| Name                    | Type                                              | Description                          | Notes                                    |
| ----------------------- | ------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**              | **string**                                        | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **webhookPatchReqBody** | [**WebhookPatchReqBody**](WebhookPatchReqBody.md) |                                      |

### Return type

[**WebhookResponse**](WebhookResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
