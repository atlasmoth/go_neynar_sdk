# \UserAPI

All URIs are relative to *https://api.neynar.com/v2*

| Method                                                                            | HTTP request                            | Description                                                        |
| --------------------------------------------------------------------------------- | --------------------------------------- | ------------------------------------------------------------------ |
| [**ActiveUsers**](UserAPI.md#ActiveUsers)                                         | **Get** /farcaster/user/active          | Fetch active users                                                 |
| [**FarcasterUserVerificationDelete**](UserAPI.md#FarcasterUserVerificationDelete) | **Delete** /farcaster/user/verification | Removes verification for an eth address for the user               |
| [**FarcasterUserVerificationPost**](UserAPI.md#FarcasterUserVerificationPost)     | **Post** /farcaster/user/verification   | Adds verification for an ethereum address or contract for the user |
| [**FollowUser**](UserAPI.md#FollowUser)                                           | **Post** /farcaster/user/follow         | Follow a user                                                      |
| [**GetFreshFid**](UserAPI.md#GetFreshFid)                                         | **Get** /farcaster/user/fid             | Fetches fid to assign it new user                                  |
| [**LookupUserByCustodyAddress**](UserAPI.md#LookupUserByCustodyAddress)           | **Get** /farcaster/user/custody-address | Lookup a user by custody-address                                   |
| [**PowerUsers**](UserAPI.md#PowerUsers)                                           | **Get** /farcaster/user/power           | Fetch power user objects                                           |
| [**RegisterUser**](UserAPI.md#RegisterUser)                                       | **Post** /farcaster/user                | Register account on farcaster                                      |
| [**UnfollowUser**](UserAPI.md#UnfollowUser)                                       | **Delete** /farcaster/user/follow       | Unfollow a user                                                    |
| [**UpdateUser**](UserAPI.md#UpdateUser)                                           | **Patch** /farcaster/user               | Update user profile                                                |
| [**UserBulk**](UserAPI.md#UserBulk)                                               | **Get** /farcaster/user/bulk            | Fetch users based on FIDs                                          |
| [**UserBulkByAddress**](UserAPI.md#UserBulkByAddress)                             | **Get** /farcaster/user/bulk-by-address | Fetches users based on Eth or Sol addresses                        |
| [**UserPowerLite**](UserAPI.md#UserPowerLite)                                     | **Get** /farcaster/user/power_lite      | Fetch power user FIDs                                              |
| [**UserSearch**](UserAPI.md#UserSearch)                                           | **Get** /farcaster/user/search          | Search for Usernames                                               |

## ActiveUsers

> UsersResponse ActiveUsers(ctx).ApiKey(apiKey).Limit(limit).Cursor(cursor).Execute()

Fetch active users

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
	limit := int32(10) // int32 |  (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.ActiveUsers(context.Background()).ApiKey(apiKey).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.ActiveUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActiveUsers`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.ActiveUsers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiActiveUsersRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **limit**  | **int32**  |                                      | [default to 25]                          |
| **cursor** | **string** | Pagination cursor.                   |

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

## FarcasterUserVerificationDelete

> OperationResponse FarcasterUserVerificationDelete(ctx).ApiKey(apiKey).RemoveVerificationReqBody(removeVerificationReqBody).Execute()

Removes verification for an eth address for the user

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
	removeVerificationReqBody := *openapiclient.NewRemoveVerificationReqBody() // RemoveVerificationReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FarcasterUserVerificationDelete(context.Background()).ApiKey(apiKey).RemoveVerificationReqBody(removeVerificationReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FarcasterUserVerificationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FarcasterUserVerificationDelete`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FarcasterUserVerificationDelete`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFarcasterUserVerificationDeleteRequest struct via the builder pattern

| Name                          | Type                                                          | Description                          | Notes                                    |
| ----------------------------- | ------------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                    | **string**                                                    | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **removeVerificationReqBody** | [**RemoveVerificationReqBody**](RemoveVerificationReqBody.md) |                                      |

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

## FarcasterUserVerificationPost

> OperationResponse FarcasterUserVerificationPost(ctx).ApiKey(apiKey).AddVerificationReqBody(addVerificationReqBody).Execute()

Adds verification for an ethereum address or contract for the user

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
	addVerificationReqBody := *openapiclient.NewAddVerificationReqBody() // AddVerificationReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FarcasterUserVerificationPost(context.Background()).ApiKey(apiKey).AddVerificationReqBody(addVerificationReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FarcasterUserVerificationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FarcasterUserVerificationPost`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FarcasterUserVerificationPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFarcasterUserVerificationPostRequest struct via the builder pattern

| Name                       | Type                                                    | Description                          | Notes                                    |
| -------------------------- | ------------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**                 | **string**                                              | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **addVerificationReqBody** | [**AddVerificationReqBody**](AddVerificationReqBody.md) |                                      |

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

## FollowUser

> BulkFollowResponse FollowUser(ctx).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()

Follow a user

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
	followReqBody := *openapiclient.NewFollowReqBody() // FollowReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.FollowUser(context.Background()).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.FollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FollowUser`: BulkFollowResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.FollowUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiFollowUserRequest struct via the builder pattern

| Name              | Type                                  | Description                          | Notes                                    |
| ----------------- | ------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**        | **string**                            | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **followReqBody** | [**FollowReqBody**](FollowReqBody.md) |                                      |

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFreshFid

> UserFIDResponse GetFreshFid(ctx).ApiKey(apiKey).Execute()

Fetches fid to assign it new user

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
	resp, r, err := apiClient.UserAPI.GetFreshFid(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetFreshFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFreshFid`: UserFIDResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetFreshFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetFreshFidRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |

### Return type

[**UserFIDResponse**](UserFIDResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## LookupUserByCustodyAddress

> UserResponse LookupUserByCustodyAddress(ctx).ApiKey(apiKey).CustodyAddress(custodyAddress).Execute()

Lookup a user by custody-address

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
	custodyAddress := "0xd1b702203b1b3b641a699997746bd4a12d157909" // string | Custody Address associated with mnemonic (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.LookupUserByCustodyAddress(context.Background()).ApiKey(apiKey).CustodyAddress(custodyAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.LookupUserByCustodyAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LookupUserByCustodyAddress`: UserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.LookupUserByCustodyAddress`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiLookupUserByCustodyAddressRequest struct via the builder pattern

| Name               | Type       | Description                              | Notes                                    |
| ------------------ | ---------- | ---------------------------------------- | ---------------------------------------- |
| **apiKey**         | **string** | API key required for authentication.     | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **custodyAddress** | **string** | Custody Address associated with mnemonic |

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

## PowerUsers

> UsersResponse PowerUsers(ctx).ApiKey(apiKey).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Fetch power user objects

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
	viewerFid := int32(3) // int32 |  (optional)
	limit := int32(10) // int32 | Number of power users to fetch, max 100 (optional) (default to 25)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.PowerUsers(context.Background()).ApiKey(apiKey).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.PowerUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PowerUsers`: UsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.PowerUsers`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPowerUsersRequest struct via the builder pattern

| Name          | Type       | Description                             | Notes                                    |
| ------------- | ---------- | --------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **viewerFid** | **int32**  |                                         |
| **limit**     | **int32**  | Number of power users to fetch, max 100 | [default to 25]                          |
| **cursor**    | **string** | Pagination cursor.                      |

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

## RegisterUser

> RegisterUserResponse RegisterUser(ctx).ApiKey(apiKey).RegisterUserReqBody(registerUserReqBody).Execute()

Register account on farcaster

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
	registerUserReqBody := *openapiclient.NewRegisterUserReqBody() // RegisterUserReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.RegisterUser(context.Background()).ApiKey(apiKey).RegisterUserReqBody(registerUserReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.RegisterUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterUser`: RegisterUserResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.RegisterUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern

| Name                    | Type                                              | Description                          | Notes                                    |
| ----------------------- | ------------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**              | **string**                                        | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **registerUserReqBody** | [**RegisterUserReqBody**](RegisterUserReqBody.md) |                                      |

### Return type

[**RegisterUserResponse**](RegisterUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UnfollowUser

> BulkFollowResponse UnfollowUser(ctx).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()

Unfollow a user

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
	followReqBody := *openapiclient.NewFollowReqBody() // FollowReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UnfollowUser(context.Background()).ApiKey(apiKey).FollowReqBody(followReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UnfollowUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnfollowUser`: BulkFollowResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UnfollowUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUnfollowUserRequest struct via the builder pattern

| Name              | Type                                  | Description                          | Notes                                    |
| ----------------- | ------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**        | **string**                            | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **followReqBody** | [**FollowReqBody**](FollowReqBody.md) |                                      |

### Return type

[**BulkFollowResponse**](BulkFollowResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUser

> OperationResponse UpdateUser(ctx).ApiKey(apiKey).UpdateUserReqBody(updateUserReqBody).Execute()

Update user profile

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
	updateUserReqBody := *openapiclient.NewUpdateUserReqBody() // UpdateUserReqBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UpdateUser(context.Background()).ApiKey(apiKey).UpdateUserReqBody(updateUserReqBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUser`: OperationResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UpdateUser`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern

| Name                  | Type                                          | Description                          | Notes                                    |
| --------------------- | --------------------------------------------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**            | **string**                                    | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **updateUserReqBody** | [**UpdateUserReqBody**](UpdateUserReqBody.md) |                                      |

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

## UserBulk

> BulkUsersResponse UserBulk(ctx).ApiKey(apiKey).Fids(fids).ViewerFid(viewerFid).Execute()

Fetch users based on FIDs

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
	fids := "194, 191, 6131" // string | Comma separated list of FIDs, up to 100 at a time (optional)
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserBulk(context.Background()).ApiKey(apiKey).Fids(fids).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserBulk`: BulkUsersResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserBulk`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserBulkRequest struct via the builder pattern

| Name          | Type       | Description                                       | Notes                                    |
| ------------- | ---------- | ------------------------------------------------- | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication.              | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fids**      | **string** | Comma separated list of FIDs, up to 100 at a time |
| **viewerFid** | **int32**  |                                                   |

### Return type

[**BulkUsersResponse**](BulkUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserBulkByAddress

> map[string][]User UserBulkByAddress(ctx).ApiKey(apiKey).Addresses(addresses).AddressTypes(addressTypes).ViewerFid(viewerFid).Execute()

Fetches users based on Eth or Sol addresses

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
	addresses := "0xa6a8736f18f383f1cc2d938576933e5ea7df01a1,0x7cac817861e5c3384753403fb6c0c556c204b1ce" // string | Comma separated list of Ethereum addresses, up to 350 at a time (optional)
	addressTypes := "custody_address,verified_address" // string | Customize which address types the request should search for. This is a comma-separated string that can include the following values: 'custody_address' and 'verified_address'. By default api returns both. To select multiple types, use a comma-separated list of these values.  (optional)
	viewerFid := int32(3) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserBulkByAddress(context.Background()).ApiKey(apiKey).Addresses(addresses).AddressTypes(addressTypes).ViewerFid(viewerFid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserBulkByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserBulkByAddress`: map[string][]User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserBulkByAddress`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserBulkByAddressRequest struct via the builder pattern

| Name             | Type       | Description                                                                                                                                                                                                                                                                                       | Notes                                    |
| ---------------- | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**       | **string** | API key required for authentication.                                                                                                                                                                                                                                                              | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **addresses**    | **string** | Comma separated list of Ethereum addresses, up to 350 at a time                                                                                                                                                                                                                                   |
| **addressTypes** | **string** | Customize which address types the request should search for. This is a comma-separated string that can include the following values: &#39;custody_address&#39; and &#39;verified_address&#39;. By default api returns both. To select multiple types, use a comma-separated list of these values. |
| **viewerFid**    | **int32**  |                                                                                                                                                                                                                                                                                                   |

### Return type

[**map[string][]User**](array.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserPowerLite

> UserPowerLiteResponse UserPowerLite(ctx).ApiKey(apiKey).Execute()

Fetch power user FIDs

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
	resp, r, err := apiClient.UserAPI.UserPowerLite(context.Background()).ApiKey(apiKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserPowerLite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserPowerLite`: UserPowerLiteResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserPowerLite`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserPowerLiteRequest struct via the builder pattern

| Name       | Type       | Description                          | Notes                                    |
| ---------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey** | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |

### Return type

[**UserPowerLiteResponse**](UserPowerLiteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserSearch

> UserSearchResponse UserSearch(ctx).ApiKey(apiKey).Q(q).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()

Search for Usernames

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
	q := "r" // string |  (optional)
	viewerFid := int32(3) // int32 |  (optional)
	limit := int32(10) // int32 |  (optional) (default to 5)
	cursor := "cursor_example" // string | Pagination cursor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.UserSearch(context.Background()).ApiKey(apiKey).Q(q).ViewerFid(viewerFid).Limit(limit).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UserSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserSearch`: UserSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.UserSearch`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserSearchRequest struct via the builder pattern

| Name          | Type       | Description                          | Notes                                    |
| ------------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**    | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **q**         | **string** |                                      |
| **viewerFid** | **int32**  |                                      |
| **limit**     | **int32**  |                                      | [default to 5]                           |
| **cursor**    | **string** | Pagination cursor.                   |

### Return type

[**UserSearchResponse**](UserSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
