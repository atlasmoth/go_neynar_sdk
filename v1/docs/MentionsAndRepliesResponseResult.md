# MentionsAndRepliesResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notifications** | [**[]CastWithInteractions**](CastWithInteractions.md) |  | 
**Next** | [**NextCursor**](NextCursor.md) |  | 

## Methods

### NewMentionsAndRepliesResponseResult

`func NewMentionsAndRepliesResponseResult(notifications []CastWithInteractions, next NextCursor, ) *MentionsAndRepliesResponseResult`

NewMentionsAndRepliesResponseResult instantiates a new MentionsAndRepliesResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMentionsAndRepliesResponseResultWithDefaults

`func NewMentionsAndRepliesResponseResultWithDefaults() *MentionsAndRepliesResponseResult`

NewMentionsAndRepliesResponseResultWithDefaults instantiates a new MentionsAndRepliesResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotifications

`func (o *MentionsAndRepliesResponseResult) GetNotifications() []CastWithInteractions`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MentionsAndRepliesResponseResult) GetNotificationsOk() (*[]CastWithInteractions, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MentionsAndRepliesResponseResult) SetNotifications(v []CastWithInteractions)`

SetNotifications sets Notifications field to given value.


### GetNext

`func (o *MentionsAndRepliesResponseResult) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *MentionsAndRepliesResponseResult) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *MentionsAndRepliesResponseResult) SetNext(v NextCursor)`

SetNext sets Next field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


