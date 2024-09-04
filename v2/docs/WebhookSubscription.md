# WebhookSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] 
**SubscriptionId** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**WebhookSubscriptionFilters**](WebhookSubscriptionFilters.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWebhookSubscription

`func NewWebhookSubscription() *WebhookSubscription`

NewWebhookSubscription instantiates a new WebhookSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSubscriptionWithDefaults

`func NewWebhookSubscriptionWithDefaults() *WebhookSubscription`

NewWebhookSubscriptionWithDefaults instantiates a new WebhookSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *WebhookSubscription) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *WebhookSubscription) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *WebhookSubscription) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *WebhookSubscription) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *WebhookSubscription) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *WebhookSubscription) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *WebhookSubscription) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *WebhookSubscription) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetFilters

`func (o *WebhookSubscription) GetFilters() WebhookSubscriptionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *WebhookSubscription) GetFiltersOk() (*WebhookSubscriptionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *WebhookSubscription) SetFilters(v WebhookSubscriptionFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *WebhookSubscription) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WebhookSubscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookSubscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookSubscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WebhookSubscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WebhookSubscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WebhookSubscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WebhookSubscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WebhookSubscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


