# IdRegisterEventBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** |  | [optional] 
**EventType** | Pointer to [**IdRegisterEventType**](IdRegisterEventType.md) |  | [optional] [default to IDREGISTEREVENTTYPE_REGISTER]
**From** | Pointer to **string** |  | [optional] 
**RecoveryAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewIdRegisterEventBody

`func NewIdRegisterEventBody() *IdRegisterEventBody`

NewIdRegisterEventBody instantiates a new IdRegisterEventBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdRegisterEventBodyWithDefaults

`func NewIdRegisterEventBodyWithDefaults() *IdRegisterEventBody`

NewIdRegisterEventBodyWithDefaults instantiates a new IdRegisterEventBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *IdRegisterEventBody) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *IdRegisterEventBody) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *IdRegisterEventBody) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *IdRegisterEventBody) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetEventType

`func (o *IdRegisterEventBody) GetEventType() IdRegisterEventType`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IdRegisterEventBody) GetEventTypeOk() (*IdRegisterEventType, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IdRegisterEventBody) SetEventType(v IdRegisterEventType)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *IdRegisterEventBody) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetFrom

`func (o *IdRegisterEventBody) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *IdRegisterEventBody) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *IdRegisterEventBody) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *IdRegisterEventBody) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetRecoveryAddress

`func (o *IdRegisterEventBody) GetRecoveryAddress() string`

GetRecoveryAddress returns the RecoveryAddress field if non-nil, zero value otherwise.

### GetRecoveryAddressOk

`func (o *IdRegisterEventBody) GetRecoveryAddressOk() (*string, bool)`

GetRecoveryAddressOk returns a tuple with the RecoveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAddress

`func (o *IdRegisterEventBody) SetRecoveryAddress(v string)`

SetRecoveryAddress sets RecoveryAddress field to given value.

### HasRecoveryAddress

`func (o *IdRegisterEventBody) HasRecoveryAddress() bool`

HasRecoveryAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


