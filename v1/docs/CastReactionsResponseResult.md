# CastReactionsResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Casts** | [**[]Reaction**](Reaction.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewCastReactionsResponseResult

`func NewCastReactionsResponseResult(casts []Reaction, next NextCursor, ) *CastReactionsResponseResult`

NewCastReactionsResponseResult instantiates a new CastReactionsResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastReactionsResponseResultWithDefaults

`func NewCastReactionsResponseResultWithDefaults() *CastReactionsResponseResult`

NewCastReactionsResponseResultWithDefaults instantiates a new CastReactionsResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCasts

`func (o *CastReactionsResponseResult) GetCasts() []Reaction`

GetCasts returns the Casts field if non-nil, zero value otherwise.

### GetCastsOk

`func (o *CastReactionsResponseResult) GetCastsOk() (*[]Reaction, bool)`

GetCastsOk returns a tuple with the Casts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCasts

`func (o *CastReactionsResponseResult) SetCasts(v []Reaction)`

SetCasts sets Casts field to given value.


### GetNext

`func (o *CastReactionsResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *CastReactionsResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *CastReactionsResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


