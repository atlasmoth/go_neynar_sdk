# \FrameAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                                             | HTTP request                                       | Description                                                 |
| ---------------------------------------------------------------------------------- | -------------------------------------------------- | ----------------------------------------------------------- |
| [**DeleteNeynarFrame**](FrameAPI.md#DeleteNeynarFrame)                             | **Delete** /farcaster/frame                        | Delete a frame                                              |
| [**FetchNeynarFrames**](FrameAPI.md#FetchNeynarFrames)                             | **Get** /farcaster/frame/list                      | Retrieve a list of frames                                   |
| [**FrameFromUrl**](FrameAPI.md#FrameFromUrl)                                       | **Get** /farcaster/frame/crawl                     | Fetches the frame meta tags from the URL                    |
| [**LookupNeynarFrame**](FrameAPI.md#LookupNeynarFrame)                             | **Get** /farcaster/frame                           | Retrieve a frame by UUID or URL                             |
| [**PostFrameAction**](FrameAPI.md#PostFrameAction)                                 | **Post** /farcaster/frame/action                   | Posts a frame action, cast action or a cast composer action |
| [**PostFrameDeveloperManagedAction**](FrameAPI.md#PostFrameDeveloperManagedAction) | **Post** /farcaster/frame/developer_managed/action | Posts a frame signature packet                              |
| [**PublishNeynarFrame**](FrameAPI.md#PublishNeynarFrame)                           | **Post** /farcaster/frame                          | Create a new frame                                          |
| [**UpdateNeynarFrame**](FrameAPI.md#UpdateNeynarFrame)                             | **Put** /farcaster/frame                           | Update an existing frame                                    |
| [**ValidateFrame**](FrameAPI.md#ValidateFrame)                                     | **Post** /farcaster/frame/validate                 | Validates a frame action against Farcaster Hub              |
| [**ValidateFrameAnalytics**](FrameAPI.md#ValidateFrameAnalytics)                   | **Get** /farcaster/frame/validate/analytics        | Retrieve analytics for the frame                            |
| [**ValidateFrameList**](FrameAPI.md#ValidateFrameList)                             | **Get** /farcaster/frame/validate/list             | Retrieve a list of all the frames validated by a user       |

## DeleteNeynarFrame

> DeleteFrameResponse DeleteNeynarFrame(ctx).ApiKey(apiKey).DeleteNeynarFrameRequest(deleteNeynarFrameRequest).Execute()

Delete a frame

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
	deleteNeynarFrameRequest := *openapiclient.NewDeleteNeynarFrameRequest() // DeleteNeynarFrameRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.DeleteNeynarFrame(context.Background()).ApiKey(apiKey).DeleteNeynarFrameRequest(deleteNeynarFrameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.DeleteNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNeynarFrame`: DeleteFrameResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.DeleteNeynarFrame`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNeynarFrameRequest struct via the builder pattern

| Name                         | Type                                                        | Description                          | Notes                                    |
| ---------------------------- | ----------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                   | **string**                                                  | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **deleteNeynarFrameRequest** | [**DeleteNeynarFrameRequest**](DeleteNeynarFrameRequest.md) |                                      |

### Return type

[**DeleteFrameResponse**](DeleteFrameResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FetchNeynarFrames

> []NeynarFrame FetchNeynarFrames(ctx).ApiKey(apiKey).Execute()

Retrieve a list of frames

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FetchNeynarFrames(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FetchNeynarFrames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FetchNeynarFrames`: []NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FetchNeynarFrames`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFetchNeynarFramesRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |

### Return type

[**[]NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FrameFromUrl

> FrameFromUrl200Response FrameFromUrl(ctx).ApiKey(apiKey).Url(url).Execute()

Fetches the frame meta tags from the URL

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
	url := "https://frames.neynar.com/f/862277df/ff7be6a4" // string | The frame URL to crawl

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.FrameFromUrl(context.Background()).ApiKey(apiKey).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.FrameFromUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FrameFromUrl`: FrameFromUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.FrameFromUrl`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFrameFromUrlRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **url**    | **string** | The frame URL to crawl               |

### Return type

[**FrameFromUrl200Response**](FrameFromUrl200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LookupNeynarFrame

> NeynarFrame LookupNeynarFrame(ctx).ApiKey(apiKey).Type*(type*).Uuid(uuid).Url(url).Execute()

Retrieve a frame by UUID or URL

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
	type_ := openapiclient.FrameType("uuid") // FrameType |
	uuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the frame to retrieve (optional)
	url := "url_example" // string | URL of the Neynar frame to retrieve (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.LookupNeynarFrame(context.Background()).ApiKey(apiKey).Type_(type_).Uuid(uuid).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.LookupNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupNeynarFrame`: NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.LookupNeynarFrame`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiLookupNeynarFrameRequest struct via the builder pattern

| Name       | Type                          | Description                          | Notes                                    |
| ---------- | ----------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string**                    | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **type\_** | [**FrameType**](FrameType.md) |                                      |
| **uuid**   | **string**                    | UUID of the frame to retrieve        |
| **url**    | **string**                    | URL of the Neynar frame to retrieve  |

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostFrameAction

> Frame PostFrameAction(ctx).ApiKey(apiKey).FrameActionReqBody(frameActionReqBody).Execute()

Posts a frame action, cast action or a cast composer action

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
	frameActionReqBody := *openapiclient.NewFrameActionReqBody("19d0c5fd-9b33-4a48-a0e2-bc7b0555baec", *openapiclient.NewFrameAction(*openapiclient.NewFrameActionButton(int32(123), openapiclient.FrameButtonActionType("post")), "FramesUrl_example", "PostUrl_example")) // FrameActionReqBody |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PostFrameAction(context.Background()).ApiKey(apiKey).FrameActionReqBody(frameActionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PostFrameAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameAction`: Frame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PostFrameAction`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameActionRequest struct via the builder pattern

| Name                   | Type                                            | Description                          | Notes                                    |
| ---------------------- | ----------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**             | **string**                                      | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **frameActionReqBody** | [**FrameActionReqBody**](FrameActionReqBody.md) |                                      |

### Return type

[**Frame**](Frame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostFrameDeveloperManagedAction

> Frame PostFrameDeveloperManagedAction(ctx).ApiKey(apiKey).FrameDeveloperManagedActionReqBody(frameDeveloperManagedActionReqBody).Execute()

Posts a frame signature packet

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
	frameDeveloperManagedActionReqBody := *openapiclient.NewFrameDeveloperManagedActionReqBody(*openapiclient.NewFrameAction(*openapiclient.NewFrameActionButton(int32(123), openapiclient.FrameButtonActionType("post")), "FramesUrl_example", "PostUrl_example"), *openapiclient.NewFrameSignaturePacket(*openapiclient.NewFrameSignaturePacketUntrustedData(), *openapiclient.NewFrameSignaturePacketTrustedData())) // FrameDeveloperManagedActionReqBody |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PostFrameDeveloperManagedAction(context.Background()).ApiKey(apiKey).FrameDeveloperManagedActionReqBody(frameDeveloperManagedActionReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PostFrameDeveloperManagedAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFrameDeveloperManagedAction`: Frame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PostFrameDeveloperManagedAction`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostFrameDeveloperManagedActionRequest struct via the builder pattern

| Name                                   | Type                                                                            | Description                          | Notes                                    |
| -------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                             | **string**                                                                      | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **frameDeveloperManagedActionReqBody** | [**FrameDeveloperManagedActionReqBody**](FrameDeveloperManagedActionReqBody.md) |                                      |

### Return type

[**Frame**](Frame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PublishNeynarFrame

> NeynarFrame PublishNeynarFrame(ctx).ApiKey(apiKey).NeynarFrameCreationRequest(neynarFrameCreationRequest).Execute()

Create a new frame

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
	neynarFrameCreationRequest := *openapiclient.NewNeynarFrameCreationRequest("Name_example", []openapiclient.NeynarFramePage{*openapiclient.NewNeynarFramePage("Uuid_example", "vNext", "Welcome to Neynar", *openapiclient.NewNeynarPageImage("https://i.imgur.com/qo2AzBf.jpeg", "AspectRatio_example"))}) // NeynarFrameCreationRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.PublishNeynarFrame(context.Background()).ApiKey(apiKey).NeynarFrameCreationRequest(neynarFrameCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.PublishNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishNeynarFrame`: NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.PublishNeynarFrame`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPublishNeynarFrameRequest struct via the builder pattern

| Name                           | Type                                                            | Description                          | Notes                                    |
| ------------------------------ | --------------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                     | **string**                                                      | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **neynarFrameCreationRequest** | [**NeynarFrameCreationRequest**](NeynarFrameCreationRequest.md) |                                      |

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateNeynarFrame

> NeynarFrame UpdateNeynarFrame(ctx).ApiKey(apiKey).NeynarFrameUpdateRequest(neynarFrameUpdateRequest).Execute()

Update an existing frame

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
	neynarFrameUpdateRequest := *openapiclient.NewNeynarFrameUpdateRequest("Uuid_example", []openapiclient.NeynarFramePage{*openapiclient.NewNeynarFramePage("Uuid_example", "vNext", "Welcome to Neynar", *openapiclient.NewNeynarPageImage("https://i.imgur.com/qo2AzBf.jpeg", "AspectRatio_example"))}) // NeynarFrameUpdateRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.UpdateNeynarFrame(context.Background()).ApiKey(apiKey).NeynarFrameUpdateRequest(neynarFrameUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.UpdateNeynarFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNeynarFrame`: NeynarFrame
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.UpdateNeynarFrame`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNeynarFrameRequest struct via the builder pattern

| Name                         | Type                                                        | Description                          | Notes                                    |
| ---------------------------- | ----------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                   | **string**                                                  | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **neynarFrameUpdateRequest** | [**NeynarFrameUpdateRequest**](NeynarFrameUpdateRequest.md) |                                      |

### Return type

[**NeynarFrame**](NeynarFrame.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ValidateFrame

> ValidateFrameActionResponse ValidateFrame(ctx).ApiKey(apiKey).ValidateFrameRequest(validateFrameRequest).Execute()

Validates a frame action against Farcaster Hub

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
	validateFrameRequest := *openapiclient.NewValidateFrameRequest("MessageBytesInHex_example") // ValidateFrameRequest |

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.ValidateFrame(context.Background()).ApiKey(apiKey).ValidateFrameRequest(validateFrameRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.ValidateFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFrame`: ValidateFrameActionResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.ValidateFrame`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiValidateFrameRequest struct via the builder pattern

| Name                     | Type                                                | Description                          | Notes                                    |
| ------------------------ | --------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**               | **string**                                          | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **validateFrameRequest** | [**ValidateFrameRequest**](ValidateFrameRequest.md) |                                      |

### Return type

[**ValidateFrameActionResponse**](ValidateFrameActionResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ValidateFrameAnalytics

> FrameValidateAnalyticsResponse ValidateFrameAnalytics(ctx).ApiKey(apiKey).FrameUrl(frameUrl).AnalyticsType(analyticsType).Start(start).Stop(stop).AggregateWindow(aggregateWindow).Execute()

Retrieve analytics for the frame

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/atlasmoth/go_neynar_sdk/v2"
)

func main() {
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	frameUrl := "https://shorturl.at/bDRY9" // string |
	analyticsType := openapiclient.ValidateFrameAnalyticsType("total-interactors") // ValidateFrameAnalyticsType |
	start := time.Now() // time.Time |  (default to "2024-04-06T06:44:56.811Z")
	stop := time.Now() // time.Time |  (default to "2024-04-08T06:44:56.811Z")
	aggregateWindow := "aggregateWindow_example" // string | Required for `analytics_type=interactions-per-cast` (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.ValidateFrameAnalytics(context.Background()).ApiKey(apiKey).FrameUrl(frameUrl).AnalyticsType(analyticsType).Start(start).Stop(stop).AggregateWindow(aggregateWindow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.ValidateFrameAnalytics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFrameAnalytics`: FrameValidateAnalyticsResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.ValidateFrameAnalytics`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiValidateFrameAnalyticsRequest struct via the builder pattern

| Name                | Type                                                            | Description                                                        | Notes                                             |
| ------------------- | --------------------------------------------------------------- | ------------------------------------------------------------------ | ------------------------------------------------- |
| **apiKey**          | **string**                                                      | API key required for authentication.                               | [default to &quot;NEYNAR_API_DOCS&quot;]          |
| **frameUrl**        | **string**                                                      |                                                                    |
| **analyticsType**   | [**ValidateFrameAnalyticsType**](ValidateFrameAnalyticsType.md) |                                                                    |
| **start**           | **time.Time**                                                   |                                                                    | [default to &quot;2024-04-06T06:44:56.811Z&quot;] |
| **stop**            | **time.Time**                                                   |                                                                    | [default to &quot;2024-04-08T06:44:56.811Z&quot;] |
| **aggregateWindow** | **string**                                                      | Required for &#x60;analytics_type&#x3D;interactions-per-cast&#x60; |

### Return type

[**FrameValidateAnalyticsResponse**](FrameValidateAnalyticsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ValidateFrameList

> FrameValidateListResponse ValidateFrameList(ctx).ApiKey(apiKey).Execute()

Retrieve a list of all the frames validated by a user

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FrameAPI.ValidateFrameList(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FrameAPI.ValidateFrameList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateFrameList`: FrameValidateListResponse
	fmt.Fprintf(os.Stdout, "Response from `FrameAPI.ValidateFrameList`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiValidateFrameListRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |

### Return type

[**FrameValidateListResponse**](FrameValidateListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
