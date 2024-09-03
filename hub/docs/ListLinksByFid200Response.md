# ListLinksByFid200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]LinkAdd**](LinkAdd.md) |  | 
**NextPageToken** | **string** |  | 

## Methods

### NewListLinksByFid200Response

`func NewListLinksByFid200Response(messages []LinkAdd, nextPageToken string, ) *ListLinksByFid200Response`

NewListLinksByFid200Response instantiates a new ListLinksByFid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLinksByFid200ResponseWithDefaults

`func NewListLinksByFid200ResponseWithDefaults() *ListLinksByFid200Response`

NewListLinksByFid200ResponseWithDefaults instantiates a new ListLinksByFid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListLinksByFid200Response) GetMessages() []LinkAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListLinksByFid200Response) GetMessagesOk() (*[]LinkAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListLinksByFid200Response) SetMessages(v []LinkAdd)`

SetMessages sets Messages field to given value.


### GetNextPageToken

`func (o *ListLinksByFid200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListLinksByFid200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListLinksByFid200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


