# GetUserDataByFid200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]UserDataAdd**](UserDataAdd.md) |  | 
**NextPageToken** | **string** |  | 

## Methods

### NewGetUserDataByFid200ResponseOneOf

`func NewGetUserDataByFid200ResponseOneOf(messages []UserDataAdd, nextPageToken string, ) *GetUserDataByFid200ResponseOneOf`

NewGetUserDataByFid200ResponseOneOf instantiates a new GetUserDataByFid200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDataByFid200ResponseOneOfWithDefaults

`func NewGetUserDataByFid200ResponseOneOfWithDefaults() *GetUserDataByFid200ResponseOneOf`

NewGetUserDataByFid200ResponseOneOfWithDefaults instantiates a new GetUserDataByFid200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetUserDataByFid200ResponseOneOf) GetMessages() []UserDataAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetUserDataByFid200ResponseOneOf) GetMessagesOk() (*[]UserDataAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetUserDataByFid200ResponseOneOf) SetMessages(v []UserDataAdd)`

SetMessages sets Messages field to given value.


### GetNextPageToken

`func (o *GetUserDataByFid200ResponseOneOf) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GetUserDataByFid200ResponseOneOf) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GetUserDataByFid200ResponseOneOf) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


