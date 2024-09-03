# Subscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**SubscriptionsCreated** | [**[]Subscription**](Subscription.md) |  | 

## Methods

### NewSubscriptions

`func NewSubscriptions(object string, subscriptionsCreated []Subscription, ) *Subscriptions`

NewSubscriptions instantiates a new Subscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsWithDefaults

`func NewSubscriptionsWithDefaults() *Subscriptions`

NewSubscriptionsWithDefaults instantiates a new Subscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Subscriptions) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Subscriptions) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Subscriptions) SetObject(v string)`

SetObject sets Object field to given value.


### GetSubscriptionsCreated

`func (o *Subscriptions) GetSubscriptionsCreated() []Subscription`

GetSubscriptionsCreated returns the SubscriptionsCreated field if non-nil, zero value otherwise.

### GetSubscriptionsCreatedOk

`func (o *Subscriptions) GetSubscriptionsCreatedOk() (*[]Subscription, bool)`

GetSubscriptionsCreatedOk returns a tuple with the SubscriptionsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsCreated

`func (o *Subscriptions) SetSubscriptionsCreated(v []Subscription)`

SetSubscriptionsCreated sets SubscriptionsCreated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


