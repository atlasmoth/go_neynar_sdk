# \CastAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                              | HTTP request                                  | Description                                    |
| --------------------------------------------------- | --------------------------------------------- | ---------------------------------------------- |
| [**Cast**](CastAPI.md#Cast)                         | **Get** /farcaster/cast                       | Retrieve cast for a given hash or Warpcast URL |
| [**CastConversation**](CastAPI.md#CastConversation) | **Get** /farcaster/cast/conversation          | Retrieve the conversation for a given cast     |
| [**CastSearch**](CastAPI.md#CastSearch)             | **Get** /farcaster/cast/search                | Search for casts                               |
| [**Casts**](CastAPI.md#Casts)                       | **Get** /farcaster/casts                      | Gets information about an array of casts       |
| [**ComposerList**](CastAPI.md#ComposerList)         | **Get** /farcaster/cast/composer_actions/list | Fetches all composer actions on Warpcast       |
| [**DeleteCast**](CastAPI.md#DeleteCast)             | **Delete** /farcaster/cast                    | Delete a cast                                  |
| [**PostCast**](CastAPI.md#PostCast)                 | **Post** /farcaster/cast                      | Posts a cast                                   |

## Cast

> CastResponse Cast(ctx).ApiKey(apiKey).Identifier(identifier).Type*(type*).ViewerFid(viewerFid).Execute()

Retrieve cast for a given hash or Warpcast URL

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
	identifier := "https://warpcast.com/rish/0x9288c1" // string | Cast identifier (Its either a url or a hash) (optional)
	type_ := openapiclient.CastParamType("url") // CastParamType |  (optional)
	viewerFid := int32(3) // int32 | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.Cast(context.Background()).ApiKey(apiKey).Identifier(identifier).Type_(type_).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.Cast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cast`: CastResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.Cast`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastRequest struct via the builder pattern

| Name           | Type                                  | Description                                                                               | Notes                                    |
| -------------- | ------------------------------------- | ----------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**     | **string**                            | API key required for authentication.                                                      | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **identifier** | **string**                            | Cast identifier (Its either a url or a hash)                                              |
| **type\_**     | [**CastParamType**](CastParamType.md) |                                                                                           |
| **viewerFid**  | **int32**                             | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. |

### Return type

[**CastResponse**](CastResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CastConversation

> Conversation CastConversation(ctx).ApiKey(apiKey).Identifier(identifier).Type*(type*).ReplyDepth(replyDepth).IncludeChronologicalParentCasts(includeChronologicalParentCasts).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Retrieve the conversation for a given cast

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
	identifier := "https://warpcast.com/rish/0x9288c1" // string | Cast identifier (Its either a url or a hash) (optional)
	type_ := openapiclient.CastParamType("url") // CastParamType |  (optional)
	replyDepth := int32(56) // int32 | The depth of replies in the conversation that will be returned (default 2) (optional) (default to 2)
	includeChronologicalParentCasts := true // bool | Include all parent casts in chronological order (optional) (default to false)
	viewerFid := int32(3) // int32 |  (optional)
	limit := int32(56) // int32 | Number of results to retrieve (default 20, max 50) (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.CastConversation(context.Background()).ApiKey(apiKey).Identifier(identifier).Type_(type_).ReplyDepth(replyDepth).IncludeChronologicalParentCasts(includeChronologicalParentCasts).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.CastConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastConversation`: Conversation
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.CastConversation`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastConversationRequest struct via the builder pattern

| Name                                | Type                                  | Description                                                                | Notes                                    |
| ----------------------------------- | ------------------------------------- | -------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**                          | **string**                            | API key required for authentication.                                       | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **identifier**                      | **string**                            | Cast identifier (Its either a url or a hash)                               |
| **type\_**                          | [**CastParamType**](CastParamType.md) |                                                                            |
| **replyDepth**                      | **int32**                             | The depth of replies in the conversation that will be returned (default 2) | [default to 2]                           |
| **includeChronologicalParentCasts** | **bool**                              | Include all parent casts in chronological order                            | [default to false]                       |
| **viewerFid**                       | **int32**                             |                                                                            |
| **limit**                           | **int32**                             | Number of results to retrieve (default 20, max 50)                         | [default to 20]                          |
| **cursor**                          | **string**                            | Pagination cursor.                                                         |

### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CastSearch

> CastsSearchResponse CastSearch(ctx).ApiKey(apiKey).Q(q).AuthorFid(authorFid).ParentUrl(parentUrl).ChannelId(channelId).Limit(limit).Cursor(cursor).Execute()

Search for casts

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
	q := "star wars" // string | Query string to search for casts (optional)
	authorFid := int32(194) // int32 | Fid of the user whose casts you want to search (optional)
	parentUrl := "parentUrl_example" // string | Parent URL of the casts you want to search (optional)
	channelId := "channelId_example" // string | Channel ID of the casts you want to search (optional)
	limit := int32(25) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.CastSearch(context.Background()).ApiKey(apiKey).Q(q).AuthorFid(authorFid).ParentUrl(parentUrl).ChannelId(channelId).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.CastSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CastSearch`: CastsSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.CastSearch`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastSearchRequest struct via the builder pattern

| Name          | Type       | Description                                    | Notes                                    |
| ------------- | ---------- | ---------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.           | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **q**         | **string** | Query string to search for casts               |
| **authorFid** | **int32**  | Fid of the user whose casts you want to search |
| **parentUrl** | **string** | Parent URL of the casts you want to search     |
| **channelId** | **string** | Channel ID of the casts you want to search     |
| **limit**     | **int32**  |                                                | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor                              |

### Return type

[**CastsSearchResponse**](CastsSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## Casts

> CastsResponse Casts(ctx).ApiKey(apiKey).Casts(casts).ViewerFid(viewerFid).SortType(sortType).Execute()

Gets information about an array of casts

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
	casts := "0xa896906a5e397b4fec247c3ee0e9e4d4990b8004,0x27ff810f7f718afd8c40be236411f017982e0994" // string | Hashes of the cast to be retrived (Comma separated) (optional)
	viewerFid := int32(3) // int32 | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. (optional)
	sortType := "sortType_example" // string | Optional parameter to sort the casts based on different criteria (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.Casts(context.Background()).ApiKey(apiKey).Casts(casts).ViewerFid(viewerFid).SortType(sortType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.Casts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Casts`: CastsResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.Casts`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCastsRequest struct via the builder pattern

| Name          | Type       | Description                                                                               | Notes                                    |
| ------------- | ---------- | ----------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                      | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **casts**     | **string** | Hashes of the cast to be retrived (Comma separated)                                       |
| **viewerFid** | **int32**  | adds viewer_context to cast object to show whether viewer has liked or recasted the cast. |
| **sortType**  | **string** | Optional parameter to sort the casts based on different criteria                          |

### Return type

[**CastsResponse**](CastsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ComposerList

> CastComposerActionsListResponse ComposerList(ctx).ApiKey(apiKey).List(list).Limit(limit).Cursor(cursor).Execute()

Fetches all composer actions on Warpcast

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
	list := openapiclient.CastComposerType("top") // CastComposerType | Type of list to fetch. (optional)
	limit := int32(25) // int32 | Number of results to retrieve (default 25, max 25). (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.ComposerList(context.Background()).ApiKey(apiKey).List(list).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.ComposerList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComposerList`: CastComposerActionsListResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.ComposerList`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiComposerListRequest struct via the builder pattern

| Name       | Type                                        | Description                                         | Notes                                    |
| ---------- | ------------------------------------------- | --------------------------------------------------- | ---------------------------------------- |
| **apiKey** | **string**                                  | API key required for authentication.                | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **list**   | [**CastComposerType**](CastComposerType.md) | Type of list to fetch.                              |
| **limit**  | **int32**                                   | Number of results to retrieve (default 25, max 25). | [default to 25]                          |
| **cursor** | **string**                                  | Pagination cursor.                                  |

### Return type

[**CastComposerActionsListResponse**](CastComposerActionsListResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteCast

> OperationResponse DeleteCast(ctx).ApiKey(apiKey).DeleteCastReqBody(deleteCastReqBody).Execute()

Delete a cast

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
	deleteCastReqBody := *openapiclient.NewDeleteCastReqBody() // DeleteCastReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.DeleteCast(context.Background()).ApiKey(apiKey).DeleteCastReqBody(deleteCastReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.DeleteCast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCast`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.DeleteCast`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCastRequest struct via the builder pattern

| Name                  | Type                                          | Description                          | Notes                                    |
| --------------------- | --------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**            | **string**                                    | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **deleteCastReqBody** | [**DeleteCastReqBody**](DeleteCastReqBody.md) |                                      |

### Return type

[**OperationResponse**](OperationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostCast

> PostCastResponse PostCast(ctx).ApiKey(apiKey).PostCastReqBody(postCastReqBody).Execute()

Posts a cast

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
	postCastReqBody := *openapiclient.NewPostCastReqBody() // PostCastReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastAPI.PostCast(context.Background()).ApiKey(apiKey).PostCastReqBody(postCastReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastAPI.PostCast``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCast`: PostCastResponse
	fmt.Fprintf(os.Stdout, "Response from `CastAPI.PostCast`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostCastRequest struct via the builder pattern

| Name                | Type                                      | Description                          | Notes                                    |
| ------------------- | ----------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**          | **string**                                | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **postCastReqBody** | [**PostCastReqBody**](PostCastReqBody.md) |                                      |

### Return type

[**PostCastResponse**](PostCastResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
