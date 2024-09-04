# UsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]User**](User.md) |  | [optional] 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewUsersResponse

`func NewUsersResponse() *UsersResponse`

NewUsersResponse instantiates a new UsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersResponseWithDefaults

`func NewUsersResponseWithDefaults() *UsersResponse`

NewUsersResponseWithDefaults instantiates a new UsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *UsersResponse) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *UsersResponse) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *UsersResponse) SetUsers(v []User)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *UsersResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetNext

`func (o *UsersResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *UsersResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *UsersResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *UsersResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


