# \HubEventsAPI

All URIs are relative to *https://hub-api.neynar.com*

| Method                                           | HTTP request          | Description              |
| ------------------------------------------------ | --------------------- | ------------------------ |
| [**GetEventById**](HubEventsAPI.md#GetEventById) | **Get** /v1/eventById | Get an event by its ID   |
| [**ListEvents**](HubEventsAPI.md#ListEvents)     | **Get** /v1/events    | Get a page of Hub events |

## GetEventById

> HubEvent GetEventById(ctx).ApiKey(apiKey).EventId(eventId).Execute()

Get an event by its ID

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	eventId := int32(56) // int32 | The Hub Id of the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubEventsAPI.GetEventById(context.Background()).ApiKey(apiKey).EventId(eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubEventsAPI.GetEventById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventById`: HubEvent
	fmt.Fprintf(os.Stdout, "Response from `HubEventsAPI.GetEventById`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventByIdRequest struct via the builder pattern

| Name        | Type       | Description                          | Notes                                    |
| ----------- | ---------- | ------------------------------------ | ---------------------------------------- |
| **apiKey**  | **string** | API key required for authentication. | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **eventId** | **int32**  | The Hub Id of the event              |

### Return type

[**HubEvent**](HubEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEvents

> ListEvents200Response ListEvents(ctx).ApiKey(apiKey).FromEventId(fromEventId).Execute()

Get a page of Hub events

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
	apiKey := "apiKey_example" // string | API key required for authentication. (default to "NEYNAR_API_DOCS")
	fromEventId := int32(56) // int32 | An optional Hub Id to start getting events from. This is also returned from the API as nextPageEventId, which can be used to page through all the Hub events. Set it to 0 to start from the first event (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HubEventsAPI.ListEvents(context.Background()).ApiKey(apiKey).FromEventId(fromEventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HubEventsAPI.ListEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvents`: ListEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `HubEventsAPI.ListEvents`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern

| Name            | Type       | Description                                                                                                                                                                                             | Notes                                    |
| --------------- | ---------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------- |
| **apiKey**      | **string** | API key required for authentication.                                                                                                                                                                    | [default to &quot;NEYNAR_API_DOCS&quot;] |
| **fromEventId** | **int32**  | An optional Hub Id to start getting events from. This is also returned from the API as nextPageEventId, which can be used to page through all the Hub events. Set it to 0 to start from the first event |

### Return type

[**ListEvents200Response**](ListEvents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
