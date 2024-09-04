# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrCode** | Pointer to **string** |  | [optional] 
**Presentable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ErrorResponseMetadata**](ErrorResponseMetadata.md) |  | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse() *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrCode

`func (o *ErrorResponse) GetErrCode() string`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *ErrorResponse) GetErrCodeOk() (*string, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *ErrorResponse) SetErrCode(v string)`

SetErrCode sets ErrCode field to given value.

### HasErrCode

`func (o *ErrorResponse) HasErrCode() bool`

HasErrCode returns a boolean if a field has been set.

### GetPresentable

`func (o *ErrorResponse) GetPresentable() bool`

GetPresentable returns the Presentable field if non-nil, zero value otherwise.

### GetPresentableOk

`func (o *ErrorResponse) GetPresentableOk() (*bool, bool)`

GetPresentableOk returns a tuple with the Presentable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentable

`func (o *ErrorResponse) SetPresentable(v bool)`

SetPresentable sets Presentable field to given value.

### HasPresentable

`func (o *ErrorResponse) HasPresentable() bool`

HasPresentable returns a boolean if a field has been set.

### GetName

`func (o *ErrorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ErrorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ErrorResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ErrorResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ErrorResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorResponse) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponse) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponse) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMetadata

`func (o *ErrorResponse) GetMetadata() ErrorResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ErrorResponse) GetMetadataOk() (*ErrorResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ErrorResponse) SetMetadata(v ErrorResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ErrorResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


