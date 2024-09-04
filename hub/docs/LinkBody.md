# LinkBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LinkType**](LinkType.md) |  | [optional] [default to LINKTYPE_FOLLOW]
**DisplayTimestamp** | Pointer to **int64** |  | [optional] 
**TargetFid** | Pointer to **int32** |  | [optional] 

## Methods

### NewLinkBody

`func NewLinkBody() *LinkBody`

NewLinkBody instantiates a new LinkBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkBodyWithDefaults

`func NewLinkBodyWithDefaults() *LinkBody`

NewLinkBodyWithDefaults instantiates a new LinkBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LinkBody) GetType() LinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LinkBody) GetTypeOk() (*LinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LinkBody) SetType(v LinkType)`

SetType sets Type field to given value.

### HasType

`func (o *LinkBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDisplayTimestamp

`func (o *LinkBody) GetDisplayTimestamp() int64`

GetDisplayTimestamp returns the DisplayTimestamp field if non-nil, zero value otherwise.

### GetDisplayTimestampOk

`func (o *LinkBody) GetDisplayTimestampOk() (*int64, bool)`

GetDisplayTimestampOk returns a tuple with the DisplayTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayTimestamp

`func (o *LinkBody) SetDisplayTimestamp(v int64)`

SetDisplayTimestamp sets DisplayTimestamp field to given value.

### HasDisplayTimestamp

`func (o *LinkBody) HasDisplayTimestamp() bool`

HasDisplayTimestamp returns a boolean if a field has been set.

### GetTargetFid

`func (o *LinkBody) GetTargetFid() int32`

GetTargetFid returns the TargetFid field if non-nil, zero value otherwise.

### GetTargetFidOk

`func (o *LinkBody) GetTargetFidOk() (*int32, bool)`

GetTargetFidOk returns a tuple with the TargetFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFid

`func (o *LinkBody) SetTargetFid(v int32)`

SetTargetFid sets TargetFid field to given value.

### HasTargetFid

`func (o *LinkBody) HasTargetFid() bool`

HasTargetFid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


