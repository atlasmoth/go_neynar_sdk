# \FollowsAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                   | HTTP request                          | Description                                  |
| -------------------------------------------------------- | ------------------------------------- | -------------------------------------------- |
| [**FollowersV2**](FollowsAPI.md#FollowersV2)             | **Get** /farcaster/followers          | Retrieve followers for a given user          |
| [**FollowingV2**](FollowsAPI.md#FollowingV2)             | **Get** /farcaster/following          | Retrieve a list of users followed by a user  |
| [**RelevantFollowers**](FollowsAPI.md#RelevantFollowers) | **Get** /farcaster/followers/relevant | Retrieve relevant followers for a given user |

## FollowersV2

> UsersResponse FollowersV2(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).Execute()

Retrieve followers for a given user

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
	fid := int32(56) // int32 | User who's profile you are looking at (optional)
	viewerFid := int32(56) // int32 | Viewer who's looking at the profile. (optional)
	sortType := openapiclient.FollowSortType("desc_chron") // FollowSortType | Sort type for retrieve followers. Default is `desc_chron` (optional)
	limit := int32(56) // int32 | Number of results to retrieve (default 20, max 100) (optional) (default to 20)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.FollowersV2(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.FollowersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowersV2`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.FollowersV2`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFollowersV2Request struct via the builder pattern

| Name          | Type                                    | Description                                                         | Notes                                    |
| ------------- | --------------------------------------- | ------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                              | API key required for authentication.                                | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**                               | User who&#39;s profile you are looking at                           |
| **viewerFid** | **int32**                               | Viewer who&#39;s looking at the profile.                            |
| **sortType**  | [**FollowSortType**](FollowSortType.md) | Sort type for retrieve followers. Default is &#x60;desc_chron&#x60; |
| **limit**     | **int32**                               | Number of results to retrieve (default 20, max 100)                 | [default to 20]                          |
| **cursor**    | **string**                              | Pagination cursor.                                                  |

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## FollowingV2

> UsersResponse FollowingV2(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).Execute()

Retrieve a list of users followed by a user

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
	fid := int32(2) // int32 | FID of the user whose following you want to fetch. (optional)
	viewerFid := int32(3) // int32 | FID of the user viewing the user. (optional)
	sortType := openapiclient.FollowSortType("desc_chron") // FollowSortType | Optional parameter to sort the users based on different criteria. (optional)
	limit := int32(25) // int32 | Number of results to retrieve (default 25, max 100) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.FollowingV2(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).SortType(sortType).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.FollowingV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowingV2`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.FollowingV2`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFollowingV2Request struct via the builder pattern

| Name          | Type                                    | Description                                                       | Notes                                    |
| ------------- | --------------------------------------- | ----------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string**                              | API key required for authentication.                              | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**                               | FID of the user whose following you want to fetch.                |
| **viewerFid** | **int32**                               | FID of the user viewing the user.                                 |
| **sortType**  | [**FollowSortType**](FollowSortType.md) | Optional parameter to sort the users based on different criteria. |
| **limit**     | **int32**                               | Number of results to retrieve (default 25, max 100)               | [default to 25]                          |
| **cursor**    | **string**                              | Pagination cursor.                                                |

### Return type

[**UsersResponse**](UsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RelevantFollowers

> RelevantFollowersResponse RelevantFollowers(ctx).ApiKey(apiKey).TargetFid(targetFid).ViewerFid(viewerFid).Execute()

Retrieve relevant followers for a given user

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
	targetFid := int32(56) // int32 | User who's profile you are looking at (optional)
	viewerFid := int32(56) // int32 | Viewer who's looking at the profile (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FollowsAPI.RelevantFollowers(context.Background()).ApiKey(apiKey).TargetFid(targetFid).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FollowsAPI.RelevantFollowers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RelevantFollowers`: RelevantFollowersResponse
	fmt.Fprintf(os.Stdout, "Response from `FollowsAPI.RelevantFollowers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRelevantFollowersRequest struct via the builder pattern

| Name          | Type       | Description                               | Notes                                    |
| ------------- | ---------- | ----------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.      | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **targetFid** | **int32**  | User who&#39;s profile you are looking at |
| **viewerFid** | **int32**  | Viewer who&#39;s looking at the profile   |

### Return type

[**RelevantFollowersResponse**](RelevantFollowersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
