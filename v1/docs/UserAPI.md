# \UserAPI

All URIs are relative to *https://api.neynar.com/v1*

| Method                                          | HTTP request                        | Description                                          |
| ----------------------------------------------- | ----------------------------------- | ---------------------------------------------------- |
| [**CustodyAddress**](UserAPI.md#CustodyAddress) | **Get** /farcaster/custody-address  | DEPRECATED - Get the custody address for a given FID |
| [**RecentUsers**](UserAPI.md#RecentUsers)       | **Get** /farcaster/recent-users     | Get Recent Users                                     |
| [**User**](UserAPI.md#User)                     | **Get** /farcaster/user             | DEPRECATED - Get User Information by FID             |
| [**UserByUsername**](UserAPI.md#UserByUsername) | **Get** /farcaster/user-by-username | Get User Information by username                     |
| [**UserCastLikes**](UserAPI.md#UserCastLikes)   | **Get** /farcaster/user-cast-likes  | DEPRECATED -- Get User Cast Likes                    |

## CustodyAddress

> CustodyAddressResponse CustodyAddress(ctx).ApiKey(apiKey).Fid(fid).Execute()

DEPRECATED - Get the custody address for a given FID

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(56) // int32 | fid of a user (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.CustodyAddress(context.Background()).ApiKey(apiKey).Fid(fid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.CustodyAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustodyAddress`: CustodyAddressResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.CustodyAddress`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCustodyAddressRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**    | **int32**  | fid of a user                        | [default to 3]                           |

### Return type

[**CustodyAddressResponse**](CustodyAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RecentUsers

> RecentUsersResponse RecentUsers(ctx).ApiKey(apiKey).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Get Recent Users

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 100, max 1000) (optional) (default to 100)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RecentUsers(context.Background()).ApiKey(apiKey).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RecentUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecentUsers`: RecentUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RecentUsers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRecentUsersRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 100, max 1000)                        | [default to 100]                         |
| **cursor**    | **string** | Pagination cursor.                                                           |

### Return type

[**RecentUsersResponse**](RecentUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## User

> UserResponse User(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Execute()

DEPRECATED - Get User Information by FID

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(56) // int32 | fid of a user (default to 3)
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.User(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.User``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `User`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.User`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | fid of a user                                                                | [default to 3]                           |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |

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

## UserByUsername

> UserResponse UserByUsername(ctx).ApiKey(apiKey).Username(username).ViewerFid(viewerFid).Execute()

Get User Information by username

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	username := "username_example" // string | Username of the user (default to "shreyas-chorge")
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserByUsername(context.Background()).ApiKey(apiKey).Username(username).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserByUsername``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserByUsername`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserByUsername`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserByUsernameRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **username**  | **string** | Username of the user                                                         | [default to &quot;shreyas-chorge&quot;]  |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |

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

## UserCastLikes

> UserCastLikeResponse UserCastLikes(ctx).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

DEPRECATED -- Get User Cast Likes

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fid := int32(56) // int32 | FID of the user (default to 3)
	viewerFid := int32(56) // int32 | fid of the user viewing this information, needed for contextual information. (optional) (default to 3)
	limit := int32(56) // int32 | Number of results to retrieve (default 25, max 150) (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserCastLikes(context.Background()).ApiKey(apiKey).Fid(fid).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserCastLikes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserCastLikes`: UserCastLikeResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserCastLikes`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserCastLikesRequest struct via the builder pattern

| Name          | Type       | Description                                                                  | Notes                                    |
| ------------- | ---------- | ---------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.                                         | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fid**       | **int32**  | FID of the user                                                              | [default to 3]                           |
| **viewerFid** | **int32**  | fid of the user viewing this information, needed for contextual information. | [default to 3]                           |
| **limit**     | **int32**  | Number of results to retrieve (default 25, max 150)                          | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor                                                            |

### Return type

[**UserCastLikeResponse**](UserCastLikeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
