# FidsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fids** | Pointer to **[]int32** |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewFidsResponse

`func NewFidsResponse() *FidsResponse`

NewFidsResponse instantiates a new FidsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFidsResponseWithDefaults

`func NewFidsResponseWithDefaults() *FidsResponse`

NewFidsResponseWithDefaults instantiates a new FidsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFids

`func (o *FidsResponse) GetFids() []int32`

GetFids returns the Fids field if non-nil, zero value otherwise.

### GetFidsOk

`func (o *FidsResponse) GetFidsOk() (*[]int32, bool)`

GetFidsOk returns a tuple with the Fids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFids

`func (o *FidsResponse) SetFids(v []int32)`

SetFids sets Fids field to given value.

### HasFids

`func (o *FidsResponse) HasFids() bool`

HasFids returns a boolean if a field has been set.

### GetNextPageToken

`func (o *FidsResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *FidsResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *FidsResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *FidsResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


