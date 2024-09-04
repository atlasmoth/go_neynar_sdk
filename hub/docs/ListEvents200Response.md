# ListEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageEventId** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to [**[]HubEvent**](HubEvent.md) |  | [optional] 

## Methods

### NewListEvents200Response

`func NewListEvents200Response() *ListEvents200Response`

NewListEvents200Response instantiates a new ListEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEvents200ResponseWithDefaults

`func NewListEvents200ResponseWithDefaults() *ListEvents200Response`

NewListEvents200ResponseWithDefaults instantiates a new ListEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageEventId

`func (o *ListEvents200Response) GetNextPageEventId() int32`

GetNextPageEventId returns the NextPageEventId field if non-nil, zero value otherwise.

### GetNextPageEventIdOk

`func (o *ListEvents200Response) GetNextPageEventIdOk() (*int32, bool)`

GetNextPageEventIdOk returns a tuple with the NextPageEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageEventId

`func (o *ListEvents200Response) SetNextPageEventId(v int32)`

SetNextPageEventId sets NextPageEventId field to given value.

### HasNextPageEventId

`func (o *ListEvents200Response) HasNextPageEventId() bool`

HasNextPageEventId returns a boolean if a field has been set.

### GetEvents

`func (o *ListEvents200Response) GetEvents() []HubEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListEvents200Response) GetEventsOk() (*[]HubEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListEvents200Response) SetEvents(v []HubEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListEvents200Response) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


