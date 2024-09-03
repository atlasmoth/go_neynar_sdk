# ReactionsAndRecastsResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | [**[]ReactionsAndRecastsNotification**](ReactionsAndRecastsNotification.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewReactionsAndRecastsResponseResult

`func NewReactionsAndRecastsResponseResult(notifications []ReactionsAndRecastsNotification, next NextCursor, ) *ReactionsAndRecastsResponseResult`

NewReactionsAndRecastsResponseResult instantiates a new ReactionsAndRecastsResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionsAndRecastsResponseResultWithDefaults

`func NewReactionsAndRecastsResponseResultWithDefaults() *ReactionsAndRecastsResponseResult`

NewReactionsAndRecastsResponseResultWithDefaults instantiates a new ReactionsAndRecastsResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *ReactionsAndRecastsResponseResult) GetNotifications() []ReactionsAndRecastsNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ReactionsAndRecastsResponseResult) GetNotificationsOk() (*[]ReactionsAndRecastsNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ReactionsAndRecastsResponseResult) SetNotifications(v []ReactionsAndRecastsNotification)`

SetNotifications sets Notifications field to given value.


### GetNext

`func (o *ReactionsAndRecastsResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ReactionsAndRecastsResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ReactionsAndRecastsResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


