# \CastsAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                   | HTTP request               | Description                                                                  |
| -------------------------------------------------------- | -------------------------- | ---------------------------------------------------------------------------- |
| [**GetCastById**](CastsAPI.md#GetCastById)               | **Get** /v1/castById       | Get a cast by its FID and Hash.                                              |
| [**ListCastsByFid**](CastsAPI.md#ListCastsByFid)         | **Get** /v1/castsByFid     | Fetch all casts authored by an FID.                                          |
| [**ListCastsByMention**](CastsAPI.md#ListCastsByMention) | **Get** /v1/castsByMention | Fetch all casts that mention an FID                                          |
| [**ListCastsByParent**](CastsAPI.md#ListCastsByParent)   | **Get** /v1/castsByParent  | Fetch all casts by parent cast&#39;s FID and Hash OR by the parent&#39;s URL |

## GetCastById

> CastAdd GetCastById(ctx).ApiKey(apiKey).Fid(fid).Hash(hash).Execute()

Get a cast by its FID and Hash.

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
	fid := int32(3) // int32 | The FID of the cast's creator (optional)
	hash := "0x03aff391a6eb1772b20b4ead9a89f732be75fe27" // string | The cast's hash (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.GetCastById(context.Background()).ApiKey(apiKey).Fid(fid).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.GetCastById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCastById`: CastAdd
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.GetCastById`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetCastByIdRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | The FID of the cast&#39;s creator    |
| **hash**   | **string** | The cast&#39;s hash                  |

### Return type

[**CastAdd**](CastAdd.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCastsByFid

> ListCastsByFid200Response ListCastsByFid(ctx).ApiKey(apiKey).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Fetch all casts authored by an FID.

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
	fid := int32(3) // int32 | The FID of the casts' creator (optional)
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.ListCastsByFid(context.Background()).ApiKey(apiKey).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.ListCastsByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCastsByFid`: ListCastsByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.ListCastsByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCastsByFidRequest struct via the builder pattern

| Name          | Type       | Description                                                                                                             | Notes                                    |
| ------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | The FID of the casts&#39; creator                                                                                       |
| **pageSize**  | **int32**  | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**   | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListCastsByFid200Response**](ListCastsByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCastsByMention

> ListCastsByFid200Response ListCastsByMention(ctx).ApiKey(apiKey).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Fetch all casts that mention an FID

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
	fid := int32(3) // int32 | The FID that is mentioned in a cast (optional)
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.ListCastsByMention(context.Background()).ApiKey(apiKey).Fid(fid).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.ListCastsByMention``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCastsByMention`: ListCastsByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.ListCastsByMention`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCastsByMentionRequest struct via the builder pattern

| Name          | Type       | Description                                                                                                             | Notes                                    |
| ------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | The FID that is mentioned in a cast                                                                                     |
| **pageSize**  | **int32**  | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**   | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListCastsByFid200Response**](ListCastsByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCastsByParent

> ListCastsByFid200Response ListCastsByParent(ctx).ApiKey(apiKey).Fid(fid).Hash(hash).Url(url).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Fetch all casts by parent cast's FID and Hash OR by the parent's URL

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
	fid := int32(226) // int32 | The FID of the parent cast (optional)
	hash := "0x03aff391a6eb1772b20b4ead9a89f732be75fe27" // string | The parent cast's hash (optional)
	url := "chain://eip155:1/erc721:0x39d89b649ffa044383333d297e325d42d31329b2" // string |  (optional)
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CastsAPI.ListCastsByParent(context.Background()).ApiKey(apiKey).Fid(fid).Hash(hash).Url(url).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CastsAPI.ListCastsByParent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCastsByParent`: ListCastsByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `CastsAPI.ListCastsByParent`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCastsByParentRequest struct via the builder pattern

| Name          | Type       | Description                                                                                                             | Notes                                    |
| ------------- | ---------- | ----------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | The FID of the parent cast                                                                                              |
| **hash**      | **string** | The parent cast&#39;s hash                                                                                              |
| **url**       | **string** |                                                                                                                         |
| **pageSize**  | **int32**  | Maximum number of messages to return in a single response                                                               |
| **reverse**   | **bool**   | Reverse the sort order, returning latest messages first                                                                 |
| **pageToken** | **string** | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page |

### Return type

[**ListCastsByFid200Response**](ListCastsByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
