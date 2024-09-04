# UserDataBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**UserDataType**](UserDataType.md) |  | [optional] [default to USERDATATYPE_PFP]
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewUserDataBody

`func NewUserDataBody() *UserDataBody`

NewUserDataBody instantiates a new UserDataBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataBodyWithDefaults

`func NewUserDataBodyWithDefaults() *UserDataBody`

NewUserDataBodyWithDefaults instantiates a new UserDataBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserDataBody) GetType() UserDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDataBody) GetTypeOk() (*UserDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDataBody) SetType(v UserDataType)`

SetType sets Type field to given value.

### HasType

`func (o *UserDataBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *UserDataBody) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UserDataBody) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UserDataBody) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UserDataBody) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


