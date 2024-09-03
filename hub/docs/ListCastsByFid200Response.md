# ListCastsByFid200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]CastAdd**](CastAdd.md) |  | 
**NextPageToken** | **string** |  | 

## Methods

### NewListCastsByFid200Response

`func NewListCastsByFid200Response(messages []CastAdd, nextPageToken string, ) *ListCastsByFid200Response`

NewListCastsByFid200Response instantiates a new ListCastsByFid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCastsByFid200ResponseWithDefaults

`func NewListCastsByFid200ResponseWithDefaults() *ListCastsByFid200Response`

NewListCastsByFid200ResponseWithDefaults instantiates a new ListCastsByFid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListCastsByFid200Response) GetMessages() []CastAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListCastsByFid200Response) GetMessagesOk() (*[]CastAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListCastsByFid200Response) SetMessages(v []CastAdd)`

SetMessages sets Messages field to given value.


### GetNextPageToken

`func (o *ListCastsByFid200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListCastsByFid200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListCastsByFid200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


