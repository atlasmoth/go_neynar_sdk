# StorageLimitsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**[]StorageLimit**](StorageLimit.md) |  | [optional] 

## Methods

### NewStorageLimitsResponse

`func NewStorageLimitsResponse() *StorageLimitsResponse`

NewStorageLimitsResponse instantiates a new StorageLimitsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageLimitsResponseWithDefaults

`func NewStorageLimitsResponseWithDefaults() *StorageLimitsResponse`

NewStorageLimitsResponseWithDefaults instantiates a new StorageLimitsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *StorageLimitsResponse) GetLimits() []StorageLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *StorageLimitsResponse) GetLimitsOk() (*[]StorageLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *StorageLimitsResponse) SetLimits(v []StorageLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *StorageLimitsResponse) HasLimits() bool`

HasLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


