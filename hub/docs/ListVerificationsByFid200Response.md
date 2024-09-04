# ListVerificationsByFid200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]Verification**](Verification.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListVerificationsByFid200Response

`func NewListVerificationsByFid200Response() *ListVerificationsByFid200Response`

NewListVerificationsByFid200Response instantiates a new ListVerificationsByFid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVerificationsByFid200ResponseWithDefaults

`func NewListVerificationsByFid200ResponseWithDefaults() *ListVerificationsByFid200Response`

NewListVerificationsByFid200ResponseWithDefaults instantiates a new ListVerificationsByFid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *ListVerificationsByFid200Response) GetMessages() []Verification`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ListVerificationsByFid200Response) GetMessagesOk() (*[]Verification, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ListVerificationsByFid200Response) SetMessages(v []Verification)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ListVerificationsByFid200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNextPageToken

`func (o *ListVerificationsByFid200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ListVerificationsByFid200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ListVerificationsByFid200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *ListVerificationsByFid200Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


