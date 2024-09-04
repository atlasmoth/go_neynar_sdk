# Subscriber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**SubscribedTo** | Pointer to [**SubscribedToObject**](SubscribedToObject.md) |  | [optional] 

## Methods

### NewSubscriber

`func NewSubscriber() *Subscriber`

NewSubscriber instantiates a new Subscriber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriberWithDefaults

`func NewSubscriberWithDefaults() *Subscriber`

NewSubscriberWithDefaults instantiates a new Subscriber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *Subscriber) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Subscriber) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Subscriber) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Subscriber) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetUser

`func (o *Subscriber) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Subscriber) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Subscriber) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Subscriber) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSubscribedTo

`func (o *Subscriber) GetSubscribedTo() SubscribedToObject`

GetSubscribedTo returns the SubscribedTo field if non-nil, zero value otherwise.

### GetSubscribedToOk

`func (o *Subscriber) GetSubscribedToOk() (*SubscribedToObject, bool)`

GetSubscribedToOk returns a tuple with the SubscribedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedTo

`func (o *Subscriber) SetSubscribedTo(v SubscribedToObject)`

SetSubscribedTo sets SubscribedTo field to given value.

### HasSubscribedTo

`func (o *Subscriber) HasSubscribedTo() bool`

HasSubscribedTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


