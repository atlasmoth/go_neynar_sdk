# \UserDataAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                                  | HTTP request              | Description             |
| ------------------------------------------------------- | ------------------------- | ----------------------- |
| [**GetUserDataByFid**](UserDataAPI.md#GetUserDataByFid) | **Get** /v1/userDataByFid | Get UserData for a FID. |

## GetUserDataByFid

> GetUserDataByFid200Response GetUserDataByFid(ctx).ApiKey(apiKey).Fid(fid).UserDataType(userDataType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()

Get UserData for a FID.

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
	fid := int32(56) // int32 | The FID that's being requested (optional)
	userDataType := openapiclient.UserDataType("USER_DATA_TYPE_PFP") // UserDataType | The type of user data, either as a numerical value or type string. If this is omitted, all user data for the FID is returned (optional) (default to "USER_DATA_TYPE_PFP")
	pageSize := int32(56) // int32 | Maximum number of messages to return in a single response (optional)
	reverse := true // bool | Reverse the sort order, returning latest messages first (optional)
	pageToken := "pageToken_example" // string | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserDataAPI.GetUserDataByFid(context.Background()).ApiKey(apiKey).Fid(fid).UserDataType(userDataType).PageSize(pageSize).Reverse(reverse).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserDataAPI.GetUserDataByFid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDataByFid`: GetUserDataByFid200Response
	fmt.Fprintf(os.Stdout, "Response from `UserDataAPI.GetUserDataByFid`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDataByFidRequest struct via the builder pattern

| Name             | Type                                | Description                                                                                                                  | Notes                                       |
| ---------------- | ----------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------- |
| **apiKey**       | **string**                          | API key required for authentication.                                                                                         | [default to &quot;NEYNAR_API_DOCS&quot;]    |
| **fid**          | **int32**                           | The FID that&#39;s being requested                                                                                           |
| **userDataType** | [**UserDataType**](UserDataType.md) | The type of user data, either as a numerical value or type string. If this is omitted, all user data for the FID is returned | [default to &quot;USER_DATA_TYPE_PFP&quot;] |
| **pageSize**     | **int32**                           | Maximum number of messages to return in a single response                                                                    |
| **reverse**      | **bool**                            | Reverse the sort order, returning latest messages first                                                                      |
| **pageToken**    | **string**                          | The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page      |

### Return type

[**GetUserDataByFid200Response**](GetUserDataByFid200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
