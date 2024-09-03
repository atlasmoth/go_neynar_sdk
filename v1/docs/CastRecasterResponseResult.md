# CastRecasterResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]Recaster**](Recaster.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewCastRecasterResponseResult

`func NewCastRecasterResponseResult(users []Recaster, next NextCursor, ) *CastRecasterResponseResult`

NewCastRecasterResponseResult instantiates a new CastRecasterResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastRecasterResponseResultWithDefaults

`func NewCastRecasterResponseResultWithDefaults() *CastRecasterResponseResult`

NewCastRecasterResponseResultWithDefaults instantiates a new CastRecasterResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *CastRecasterResponseResult) GetUsers() []Recaster`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *CastRecasterResponseResult) GetUsersOk() (*[]Recaster, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *CastRecasterResponseResult) SetUsers(v []Recaster)`

SetUsers sets Users field to given value.


### GetNext

`func (o *CastRecasterResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CastRecasterResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CastRecasterResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


