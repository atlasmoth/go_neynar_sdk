# ValidateFrameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageBytesInHex** | Pointer to **string** | Hexadecimal string of message bytes. | [optional] 
**CastReactionContext** | Pointer to **bool** | Adds viewer_context inside the cast object to indicate whether the interactor reacted to the cast housing the frame. | [optional] [default to true]
**FollowContext** | Pointer to **bool** | Adds viewer_context inside the user (interactor) object to indicate whether the interactor follows or is followed by the cast author. | [optional] [default to false]
**SignerContext** | Pointer to **bool** | Adds context about the app used by the user inside &#x60;frame.action&#x60;. | [optional] [default to false]
**ChannelFollowContext** | Pointer to **bool** | Adds context about the channel that the cast belongs to inside of the cast object. | [optional] [default to false]

## Methods

### NewValidateFrameRequest

`func NewValidateFrameRequest() *ValidateFrameRequest`

NewValidateFrameRequest instantiates a new ValidateFrameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateFrameRequestWithDefaults

`func NewValidateFrameRequestWithDefaults() *ValidateFrameRequest`

NewValidateFrameRequestWithDefaults instantiates a new ValidateFrameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageBytesInHex

`func (o *ValidateFrameRequest) GetMessageBytesInHex() string`

GetMessageBytesInHex returns the MessageBytesInHex field if non-nil, zero value otherwise.

### GetMessageBytesInHexOk

`func (o *ValidateFrameRequest) GetMessageBytesInHexOk() (*string, bool)`

GetMessageBytesInHexOk returns a tuple with the MessageBytesInHex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBytesInHex

`func (o *ValidateFrameRequest) SetMessageBytesInHex(v string)`

SetMessageBytesInHex sets MessageBytesInHex field to given value.

### HasMessageBytesInHex

`func (o *ValidateFrameRequest) HasMessageBytesInHex() bool`

HasMessageBytesInHex returns a boolean if a field has been set.

### GetCastReactionContext

`func (o *ValidateFrameRequest) GetCastReactionContext() bool`

GetCastReactionContext returns the CastReactionContext field if non-nil, zero value otherwise.

### GetCastReactionContextOk

`func (o *ValidateFrameRequest) GetCastReactionContextOk() (*bool, bool)`

GetCastReactionContextOk returns a tuple with the CastReactionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastReactionContext

`func (o *ValidateFrameRequest) SetCastReactionContext(v bool)`

SetCastReactionContext sets CastReactionContext field to given value.

### HasCastReactionContext

`func (o *ValidateFrameRequest) HasCastReactionContext() bool`

HasCastReactionContext returns a boolean if a field has been set.

### GetFollowContext

`func (o *ValidateFrameRequest) GetFollowContext() bool`

GetFollowContext returns the FollowContext field if non-nil, zero value otherwise.

### GetFollowContextOk

`func (o *ValidateFrameRequest) GetFollowContextOk() (*bool, bool)`

GetFollowContextOk returns a tuple with the FollowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowContext

`func (o *ValidateFrameRequest) SetFollowContext(v bool)`

SetFollowContext sets FollowContext field to given value.

### HasFollowContext

`func (o *ValidateFrameRequest) HasFollowContext() bool`

HasFollowContext returns a boolean if a field has been set.

### GetSignerContext

`func (o *ValidateFrameRequest) GetSignerContext() bool`

GetSignerContext returns the SignerContext field if non-nil, zero value otherwise.

### GetSignerContextOk

`func (o *ValidateFrameRequest) GetSignerContextOk() (*bool, bool)`

GetSignerContextOk returns a tuple with the SignerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerContext

`func (o *ValidateFrameRequest) SetSignerContext(v bool)`

SetSignerContext sets SignerContext field to given value.

### HasSignerContext

`func (o *ValidateFrameRequest) HasSignerContext() bool`

HasSignerContext returns a boolean if a field has been set.

### GetChannelFollowContext

`func (o *ValidateFrameRequest) GetChannelFollowContext() bool`

GetChannelFollowContext returns the ChannelFollowContext field if non-nil, zero value otherwise.

### GetChannelFollowContextOk

`func (o *ValidateFrameRequest) GetChannelFollowContextOk() (*bool, bool)`

GetChannelFollowContextOk returns a tuple with the ChannelFollowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelFollowContext

`func (o *ValidateFrameRequest) SetChannelFollowContext(v bool)`

SetChannelFollowContext sets ChannelFollowContext field to given value.

### HasChannelFollowContext

`func (o *ValidateFrameRequest) HasChannelFollowContext() bool`

HasChannelFollowContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


