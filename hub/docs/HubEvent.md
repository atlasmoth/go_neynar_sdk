# HubEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**MergeMessageBody** | Pointer to [**MergeMessageBody**](MergeMessageBody.md) |  | [optional] 
**PruneMessageBody** | Pointer to [**PruneMessageBody**](PruneMessageBody.md) |  | [optional] 
**RevokeMessageBody** | Pointer to [**RevokeMessageBody**](RevokeMessageBody.md) |  | [optional] 
**MergeUsernameProofBody** | Pointer to [**MergeUserNameProofBody**](MergeUserNameProofBody.md) |  | [optional] 
**MergeOnChainEventBody** | Pointer to [**MergeOnChainEventBody**](MergeOnChainEventBody.md) |  | [optional] 

## Methods

### NewHubEvent

`func NewHubEvent() *HubEvent`

NewHubEvent instantiates a new HubEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubEventWithDefaults

`func NewHubEventWithDefaults() *HubEvent`

NewHubEventWithDefaults instantiates a new HubEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *HubEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HubEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HubEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HubEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *HubEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HubEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HubEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *HubEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMergeMessageBody

`func (o *HubEvent) GetMergeMessageBody() MergeMessageBody`

GetMergeMessageBody returns the MergeMessageBody field if non-nil, zero value otherwise.

### GetMergeMessageBodyOk

`func (o *HubEvent) GetMergeMessageBodyOk() (*MergeMessageBody, bool)`

GetMergeMessageBodyOk returns a tuple with the MergeMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMessageBody

`func (o *HubEvent) SetMergeMessageBody(v MergeMessageBody)`

SetMergeMessageBody sets MergeMessageBody field to given value.

### HasMergeMessageBody

`func (o *HubEvent) HasMergeMessageBody() bool`

HasMergeMessageBody returns a boolean if a field has been set.

### GetPruneMessageBody

`func (o *HubEvent) GetPruneMessageBody() PruneMessageBody`

GetPruneMessageBody returns the PruneMessageBody field if non-nil, zero value otherwise.

### GetPruneMessageBodyOk

`func (o *HubEvent) GetPruneMessageBodyOk() (*PruneMessageBody, bool)`

GetPruneMessageBodyOk returns a tuple with the PruneMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruneMessageBody

`func (o *HubEvent) SetPruneMessageBody(v PruneMessageBody)`

SetPruneMessageBody sets PruneMessageBody field to given value.

### HasPruneMessageBody

`func (o *HubEvent) HasPruneMessageBody() bool`

HasPruneMessageBody returns a boolean if a field has been set.

### GetRevokeMessageBody

`func (o *HubEvent) GetRevokeMessageBody() RevokeMessageBody`

GetRevokeMessageBody returns the RevokeMessageBody field if non-nil, zero value otherwise.

### GetRevokeMessageBodyOk

`func (o *HubEvent) GetRevokeMessageBodyOk() (*RevokeMessageBody, bool)`

GetRevokeMessageBodyOk returns a tuple with the RevokeMessageBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeMessageBody

`func (o *HubEvent) SetRevokeMessageBody(v RevokeMessageBody)`

SetRevokeMessageBody sets RevokeMessageBody field to given value.

### HasRevokeMessageBody

`func (o *HubEvent) HasRevokeMessageBody() bool`

HasRevokeMessageBody returns a boolean if a field has been set.

### GetMergeUsernameProofBody

`func (o *HubEvent) GetMergeUsernameProofBody() MergeUserNameProofBody`

GetMergeUsernameProofBody returns the MergeUsernameProofBody field if non-nil, zero value otherwise.

### GetMergeUsernameProofBodyOk

`func (o *HubEvent) GetMergeUsernameProofBodyOk() (*MergeUserNameProofBody, bool)`

GetMergeUsernameProofBodyOk returns a tuple with the MergeUsernameProofBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeUsernameProofBody

`func (o *HubEvent) SetMergeUsernameProofBody(v MergeUserNameProofBody)`

SetMergeUsernameProofBody sets MergeUsernameProofBody field to given value.

### HasMergeUsernameProofBody

`func (o *HubEvent) HasMergeUsernameProofBody() bool`

HasMergeUsernameProofBody returns a boolean if a field has been set.

### GetMergeOnChainEventBody

`func (o *HubEvent) GetMergeOnChainEventBody() MergeOnChainEventBody`

GetMergeOnChainEventBody returns the MergeOnChainEventBody field if non-nil, zero value otherwise.

### GetMergeOnChainEventBodyOk

`func (o *HubEvent) GetMergeOnChainEventBodyOk() (*MergeOnChainEventBody, bool)`

GetMergeOnChainEventBodyOk returns a tuple with the MergeOnChainEventBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeOnChainEventBody

`func (o *HubEvent) SetMergeOnChainEventBody(v MergeOnChainEventBody)`

SetMergeOnChainEventBody sets MergeOnChainEventBody field to given value.

### HasMergeOnChainEventBody

`func (o *HubEvent) HasMergeOnChainEventBody() bool`

HasMergeOnChainEventBody returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


