# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**MostRecentTimestamp** | **time.Time** |  | 
**Type** | **string** |  | 
**Follows** | Pointer to [**[]Follow**](Follow.md) |  | [optional] 
**Cast** | Pointer to [**CastWithInteractions**](CastWithInteractions.md) |  | [optional] 
**Reactions** | Pointer to [**[]ReactionWithUserInfo**](ReactionWithUserInfo.md) |  | [optional] 

## Methods

### NewNotification

`func NewNotification(object string, mostRecentTimestamp time.Time, type_ string, ) *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Notification) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Notification) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Notification) SetObject(v string)`

SetObject sets Object field to given value.


### GetMostRecentTimestamp

`func (o *Notification) GetMostRecentTimestamp() time.Time`

GetMostRecentTimestamp returns the MostRecentTimestamp field if non-nil, zero value otherwise.

### GetMostRecentTimestampOk

`func (o *Notification) GetMostRecentTimestampOk() (*time.Time, bool)`

GetMostRecentTimestampOk returns a tuple with the MostRecentTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentTimestamp

`func (o *Notification) SetMostRecentTimestamp(v time.Time)`

SetMostRecentTimestamp sets MostRecentTimestamp field to given value.


### GetType

`func (o *Notification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v string)`

SetType sets Type field to given value.


### GetFollows

`func (o *Notification) GetFollows() []Follow`

GetFollows returns the Follows field if non-nil, zero value otherwise.

### GetFollowsOk

`func (o *Notification) GetFollowsOk() (*[]Follow, bool)`

GetFollowsOk returns a tuple with the Follows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollows

`func (o *Notification) SetFollows(v []Follow)`

SetFollows sets Follows field to given value.

### HasFollows

`func (o *Notification) HasFollows() bool`

HasFollows returns a boolean if a field has been set.

### GetCast

`func (o *Notification) GetCast() CastWithInteractions`

GetCast returns the Cast field if non-nil, zero value otherwise.

### GetCastOk

`func (o *Notification) GetCastOk() (*CastWithInteractions, bool)`

GetCastOk returns a tuple with the Cast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCast

`func (o *Notification) SetCast(v CastWithInteractions)`

SetCast sets Cast field to given value.

### HasCast

`func (o *Notification) HasCast() bool`

HasCast returns a boolean if a field has been set.

### GetReactions

`func (o *Notification) GetReactions() []ReactionWithUserInfo`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Notification) GetReactionsOk() (*[]ReactionWithUserInfo, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Notification) SetReactions(v []ReactionWithUserInfo)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *Notification) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


