# \LinksAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                       | HTTP request                 | Description                           |
| ------------------------------------------------------------ | ---------------------------- | ------------------------------------- |
| [**GetLinkById**](LinksAPI.md#GetLinkById)                   | **Get** /v1/linkById         | Get a link by its FID and target FID. |
| [**ListLinksByFid**](LinksAPI.md#ListLinksByFid)             | **Get** /v1/linksByFid       | Get all links from a source FID       |
| [**ListLinksByTargetFid**](LinksAPI.md#ListLinksByTargetFid) | **Get** /v1/linksByTargetFid | Get all links to a target FID         |

## GetLinkById

> LinkAdd GetLinkById(ctx).ApiKey(apiKey).Fid(fid).TargetFid(targetFid).LinkType(linkType).Execute()

Get a link by its FID and target FID.

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
	fid := int32(56) // int32 | The FID of the link's originator (optional)
	targetFid := int32(56) // int32 | The FID of the target of the link (optional)
	linkType := openapiclient.LinkType("follow") // LinkType | The type of link, as a string value (optional) (default to "follow")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.GetLinkById(context.Background()).ApiKey(apiKey).Fid(fid).TargetFid(targetFid).LinkType(linkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.GetLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkById`: LinkAdd
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.GetLinkById`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkByIdRequest struct via the builder pattern

| Name          | Type                        | Description                          | Notes                                    |
| ------------- | --------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**    | **string**                  | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**                   | The FID of the link&#39;s originator |
| **targetFid** | **int32**                   | The FID of the target of the link    |
| **linkType**  | [**LinkType**](LinkType.md) | The type of link, as a string value  | [default to &quot;follow&quot;]          |

### Return type

[**LinkAdd**](LinkAdd.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLinksByFid

> ListLinksByFid200Response ListLinksByFid(ctx).ApiKey(apiKey).Fid(fid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get all links from a source FID

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
	fid := int32(56) // int32 | The FID of the link's originator (optional)
	linkType := openapiclient.LinkType("follow") // LinkType | The type of link, as a string value (optional) (default to "follow")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.ListLinksByFid(context.Background()).ApiKey(apiKey).Fid(fid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.ListLinksByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLinksByFid`: ListLinksByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.ListLinksByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListLinksByFidRequest struct via the builder pattern

| Name          | Type                        | Description                                                                                                             | Notes                                    |
| ------------- | --------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                  | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**                   | The FID of the link&#39;s originator                                                                                    |
| **linkType**  | [**LinkType**](LinkType.md) | The type of link, as a string value                                                                                     | [default to &quot;follow&quot;]          |
| **pageSize**  | **int32**                   | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**                    | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string**                  | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListLinksByFid200Response**](ListLinksByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLinksByTargetFid

> ListLinksByFid200Response ListLinksByTargetFid(ctx).ApiKey(apiKey).TargetFid(targetFid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get all links to a target FID

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
	targetFid := int32(56) // int32 | The FID of the target of the link (optional)
	linkType := openapiclient.LinkType("follow") // LinkType | The type of link, as a string value (optional) (default to "follow")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinksAPI.ListLinksByTargetFid(context.Background()).ApiKey(apiKey).TargetFid(targetFid).LinkType(linkType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinksAPI.ListLinksByTargetFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLinksByTargetFid`: ListLinksByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `LinksAPI.ListLinksByTargetFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListLinksByTargetFidRequest struct via the builder pattern

| Name          | Type                        | Description                                                                                                             | Notes                                    |
| ------------- | --------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                  | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **targetFid** | **int32**                   | The FID of the target of the link                                                                                       |
| **linkType**  | [**LinkType**](LinkType.md) | The type of link, as a string value                                                                                     | [default to &quot;follow&quot;]          |
| **pageSize**  | **int32**                   | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**                    | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string**                  | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListLinksByFid200Response**](ListLinksByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
