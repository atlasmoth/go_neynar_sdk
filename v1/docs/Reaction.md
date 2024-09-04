# Reaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ReactionType**](ReactionType.md) |  | [optional] 
**Hash** | Pointer to **string** | Ethereum address | [optional] [default to "0x5A927Ac639636E534b678e81768CA19e2C6280B7"]
**Reactor** | Pointer to [**Reactor**](Reactor.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**CastHash** | Pointer to **string** | Cast Hash | [optional] [default to "0xfe90f9de682273e05b201629ad2338bdcd89b6be"]

## Methods

### NewReaction

`func NewReaction() *Reaction`

NewReaction instantiates a new Reaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReactionWithDefaults

`func NewReactionWithDefaults() *Reaction`

NewReactionWithDefaults instantiates a new Reaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Reaction) GetType() ReactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Reaction) GetTypeOk() (*ReactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Reaction) SetType(v ReactionType)`

SetType sets Type field to given value.

### HasType

`func (o *Reaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHash

`func (o *Reaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Reaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Reaction) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Reaction) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetReactor

`func (o *Reaction) GetReactor() Reactor`

GetReactor returns the Reactor field if non-nil, zero value otherwise.

### GetReactorOk

`func (o *Reaction) GetReactorOk() (*Reactor, bool)`

GetReactorOk returns a tuple with the Reactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactor

`func (o *Reaction) SetReactor(v Reactor)`

SetReactor sets Reactor field to given value.

### HasReactor

`func (o *Reaction) HasReactor() bool`

HasReactor returns a boolean if a field has been set.

### GetTimestamp

`func (o *Reaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Reaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Reaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Reaction) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCastHash

`func (o *Reaction) GetCastHash() string`

GetCastHash returns the CastHash field if non-nil, zero value otherwise.

### GetCastHashOk

`func (o *Reaction) GetCastHashOk() (*string, bool)`

GetCastHashOk returns a tuple with the CastHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastHash

`func (o *Reaction) SetCastHash(v string)`

SetCastHash sets CastHash field to given value.

### HasCastHash

`func (o *Reaction) HasCastHash() bool`

HasCastHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


