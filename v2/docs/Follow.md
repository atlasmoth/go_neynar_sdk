# Follow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewFollow

`func NewFollow() *Follow`

NewFollow instantiates a new Follow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowWithDefaults

`func NewFollowWithDefaults() *Follow`

NewFollowWithDefaults instantiates a new Follow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Follow) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Follow) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Follow) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Follow) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUser

`func (o *Follow) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Follow) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Follow) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Follow) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


