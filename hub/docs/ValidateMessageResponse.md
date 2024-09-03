# ValidateMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**Message** | [**Message**](Message.md) |  | 

## Methods

### NewValidateMessageResponse

`func NewValidateMessageResponse(valid bool, message Message, ) *ValidateMessageResponse`

NewValidateMessageResponse instantiates a new ValidateMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateMessageResponseWithDefaults

`func NewValidateMessageResponseWithDefaults() *ValidateMessageResponse`

NewValidateMessageResponseWithDefaults instantiates a new ValidateMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ValidateMessageResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidateMessageResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidateMessageResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetMessage

`func (o *ValidateMessageResponse) GetMessage() Message`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidateMessageResponse) GetMessageOk() (*Message, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidateMessageResponse) SetMessage(v Message)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


