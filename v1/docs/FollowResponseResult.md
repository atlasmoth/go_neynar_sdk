# FollowResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]FollowResponseUser**](FollowResponseUser.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewFollowResponseResult

`func NewFollowResponseResult(users []FollowResponseUser, next NextCursor, ) *FollowResponseResult`

NewFollowResponseResult instantiates a new FollowResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowResponseResultWithDefaults

`func NewFollowResponseResultWithDefaults() *FollowResponseResult`

NewFollowResponseResultWithDefaults instantiates a new FollowResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *FollowResponseResult) GetUsers() []FollowResponseUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FollowResponseResult) GetUsersOk() (*[]FollowResponseUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FollowResponseResult) SetUsers(v []FollowResponseUser)`

SetUsers sets Users field to given value.


### GetNext

`func (o *FollowResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FollowResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FollowResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


