# RegisterUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Signer** | Pointer to [**Signer**](Signer.md) |  | [optional] 

## Methods

### NewRegisterUserResponse

`func NewRegisterUserResponse() *RegisterUserResponse`

NewRegisterUserResponse instantiates a new RegisterUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserResponseWithDefaults

`func NewRegisterUserResponseWithDefaults() *RegisterUserResponse`

NewRegisterUserResponseWithDefaults instantiates a new RegisterUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RegisterUserResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RegisterUserResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RegisterUserResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RegisterUserResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *RegisterUserResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegisterUserResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegisterUserResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RegisterUserResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSigner

`func (o *RegisterUserResponse) GetSigner() Signer`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *RegisterUserResponse) GetSignerOk() (*Signer, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *RegisterUserResponse) SetSigner(v Signer)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *RegisterUserResponse) HasSigner() bool`

HasSigner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


