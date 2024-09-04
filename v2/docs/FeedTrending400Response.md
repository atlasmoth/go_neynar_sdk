# FeedTrending400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]ZodErrorErrorsInner**](ZodErrorErrorsInner.md) |  | [optional] 
**Property** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewFeedTrending400Response

`func NewFeedTrending400Response() *FeedTrending400Response`

NewFeedTrending400Response instantiates a new FeedTrending400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedTrending400ResponseWithDefaults

`func NewFeedTrending400ResponseWithDefaults() *FeedTrending400Response`

NewFeedTrending400ResponseWithDefaults instantiates a new FeedTrending400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *FeedTrending400Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FeedTrending400Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FeedTrending400Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FeedTrending400Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *FeedTrending400Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *FeedTrending400Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *FeedTrending400Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *FeedTrending400Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetErrors

`func (o *FeedTrending400Response) GetErrors() []ZodErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *FeedTrending400Response) GetErrorsOk() (*[]ZodErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *FeedTrending400Response) SetErrors(v []ZodErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *FeedTrending400Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetProperty

`func (o *FeedTrending400Response) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *FeedTrending400Response) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *FeedTrending400Response) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *FeedTrending400Response) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetStatus

`func (o *FeedTrending400Response) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FeedTrending400Response) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FeedTrending400Response) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FeedTrending400Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


