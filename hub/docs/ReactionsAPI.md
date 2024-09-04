# \ReactionsAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                             | HTTP request                  | Description                                        |
| ------------------------------------------------------------------ | ----------------------------- | -------------------------------------------------- |
| [**GetReactionById**](ReactionsAPI.md#GetReactionById)             | **Get** /v1/reactionById      | Get a reaction by its created FID and target Cast. |
| [**ListReactionsByCast**](ReactionsAPI.md#ListReactionsByCast)     | **Get** /v1/reactionsByCast   | Get all reactions to a cast                        |
| [**ListReactionsByFid**](ReactionsAPI.md#ListReactionsByFid)       | **Get** /v1/reactionsByFid    | Get all reactions by an FID                        |
| [**ListReactionsByTarget**](ReactionsAPI.md#ListReactionsByTarget) | **Get** /v1/reactionsByTarget | Get all reactions to a target URL                  |

## GetReactionById

> Reaction GetReactionById(ctx).ApiKey(apiKey).Fid(fid).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).Execute()

Get a reaction by its created FID and target Cast.

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
	fid := int32(56) // int32 | The FID of the reaction's creator (optional)
	targetFid := int32(56) // int32 | The FID of the cast's creator (optional)
	targetHash := "targetHash_example" // string | The cast's hash (optional)
	reactionType := openapiclient.ReactionType("REACTION_TYPE_LIKE") // ReactionType | The type of reaction, either as a numerical enum value or string representation (optional) (default to "REACTION_TYPE_LIKE")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.GetReactionById(context.Background()).ApiKey(apiKey).Fid(fid).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.GetReactionById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReactionById`: Reaction
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.GetReactionById`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetReactionByIdRequest struct via the builder pattern

| Name             | Type                                | Description                                                                     | Notes                                       |
| ---------------- | ----------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------- |
| **apiKey**       | **string**                          | API key required for authentication.                                            | [default to &quot;NEYNAR_API_DOCS&quot;]    |
| **fid**          | **int32**                           | The FID of the reaction&#39;s creator                                           |
| **targetFid**    | **int32**                           | The FID of the cast&#39;s creator                                               |
| **targetHash**   | **string**                          | The cast&#39;s hash                                                             |
| **reactionType** | [**ReactionType**](ReactionType.md) | The type of reaction, either as a numerical enum value or string representation | [default to &quot;REACTION_TYPE_LIKE&quot;] |

### Return type

[**Reaction**](Reaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListReactionsByCast

> ListReactionsByCast200Response ListReactionsByCast(ctx).ApiKey(apiKey).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get all reactions to a cast

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
	targetFid := int32(56) // int32 | The FID of the cast's creator (optional)
	targetHash := "targetHash_example" // string | The hash of the cast (optional)
	reactionType := openapiclient.ReactionType("REACTION_TYPE_LIKE") // ReactionType | The type of reaction, either as a numerical enum value or string representation (optional) (default to "REACTION_TYPE_LIKE")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.ListReactionsByCast(context.Background()).ApiKey(apiKey).TargetFid(targetFid).TargetHash(targetHash).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.ListReactionsByCast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReactionsByCast`: ListReactionsByCast200Response
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.ListReactionsByCast`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListReactionsByCastRequest struct via the builder pattern

| Name             | Type                                | Description                                                                                                             | Notes                                       |
| ---------------- | ----------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| **apiKey**       | **string**                          | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;]    |
| **targetFid**    | **int32**                           | The FID of the cast&#39;s creator                                                                                       |
| **targetHash**   | **string**                          | The hash of the cast                                                                                                    |
| **reactionType** | [**ReactionType**](ReactionType.md) | The type of reaction, either as a numerical enum value or string representation                                         | [default to &quot;REACTION_TYPE_LIKE&quot;] |
| **pageSize**     | **int32**                           | Maximum number of messages to return in a single response                                                               |
| **reverse**      | **bool**                            | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken**    | **string**                          | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListReactionsByCast200Response**](ListReactionsByCast200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListReactionsByFid

> ListReactionsByCast200Response ListReactionsByFid(ctx).ApiKey(apiKey).Fid(fid).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get all reactions by an FID

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
	fid := int32(56) // int32 | The FID of the reaction's creator (optional)
	reactionType := openapiclient.ReactionType("REACTION_TYPE_LIKE") // ReactionType | The type of reaction, either as a numerical enum value or string representation (optional) (default to "REACTION_TYPE_LIKE")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.ListReactionsByFid(context.Background()).ApiKey(apiKey).Fid(fid).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.ListReactionsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReactionsByFid`: ListReactionsByCast200Response
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.ListReactionsByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListReactionsByFidRequest struct via the builder pattern

| Name             | Type                                | Description                                                                                                             | Notes                                       |
| ---------------- | ----------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| **apiKey**       | **string**                          | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;]    |
| **fid**          | **int32**                           | The FID of the reaction&#39;s creator                                                                                   |
| **reactionType** | [**ReactionType**](ReactionType.md) | The type of reaction, either as a numerical enum value or string representation                                         | [default to &quot;REACTION_TYPE_LIKE&quot;] |
| **pageSize**     | **int32**                           | Maximum number of messages to return in a single response                                                               |
| **reverse**      | **bool**                            | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken**    | **string**                          | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListReactionsByCast200Response**](ListReactionsByCast200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListReactionsByTarget

> ListReactionsByCast200Response ListReactionsByTarget(ctx).ApiKey(apiKey).Url(url).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get all reactions to a target URL

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
	url := "chain://eip155:1/erc721:0x39d89b649ffa044383333d297e325d42d31329b2" // string | The URL of the parent cast (optional)
	reactionType := openapiclient.ReactionType("REACTION_TYPE_LIKE") // ReactionType | The type of reaction, either as a numerical enum value or string representation (optional) (default to "REACTION_TYPE_LIKE")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReactionsAPI.ListReactionsByTarget(context.Background()).ApiKey(apiKey).Url(url).ReactionType(reactionType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReactionsAPI.ListReactionsByTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReactionsByTarget`: ListReactionsByCast200Response
	fmt.Fprintf(os.Stdout, "Response from `ReactionsAPI.ListReactionsByTarget`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListReactionsByTargetRequest struct via the builder pattern

| Name             | Type                                | Description                                                                                                             | Notes                                       |
| ---------------- | ----------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| **apiKey**       | **string**                          | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;]    |
| **url**          | **string**                          | The URL of the parent cast                                                                                              |
| **reactionType** | [**ReactionType**](ReactionType.md) | The type of reaction, either as a numerical enum value or string representation                                         | [default to &quot;REACTION_TYPE_LIKE&quot;] |
| **pageSize**     | **int32**                           | Maximum number of messages to return in a single response                                                               |
| **reverse**      | **bool**                            | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken**    | **string**                          | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListReactionsByCast200Response**](ListReactionsByCast200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
