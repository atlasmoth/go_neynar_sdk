# RecentUsersResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]User**](User.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewRecentUsersResponseResult

`func NewRecentUsersResponseResult(users []User, next NextCursor, ) *RecentUsersResponseResult`

NewRecentUsersResponseResult instantiates a new RecentUsersResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecentUsersResponseResultWithDefaults

`func NewRecentUsersResponseResultWithDefaults() *RecentUsersResponseResult`

NewRecentUsersResponseResultWithDefaults instantiates a new RecentUsersResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *RecentUsersResponseResult) GetUsers() []User`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RecentUsersResponseResult) GetUsersOk() (*[]User, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RecentUsersResponseResult) SetUsers(v []User)`

SetUsers sets Users field to given value.


### GetNext

`func (o *RecentUsersResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *RecentUsersResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *RecentUsersResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


