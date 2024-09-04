# Cast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**ParentHash** | Pointer to **NullableString** |  | [optional] 
**ParentUrl** | Pointer to **NullableString** |  | [optional] 
**RootParentUrl** | Pointer to **NullableString** |  | [optional] 
**ParentAuthor** | Pointer to [**CastParentAuthor**](CastParentAuthor.md) |  | [optional] 
**Author** | Pointer to [**User**](User.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Embeds** | Pointer to [**[]EmbeddedCast**](EmbeddedCast.md) |  | [optional] 
**Type** | Pointer to [**CastNotificationType**](CastNotificationType.md) |  | [optional] 

## Methods

### NewCast

`func NewCast() *Cast`

NewCast instantiates a new Cast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastWithDefaults

`func NewCastWithDefaults() *Cast`

NewCastWithDefaults instantiates a new Cast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *Cast) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Cast) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Cast) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Cast) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetParentHash

`func (o *Cast) GetParentHash() string`

GetParentHash returns the ParentHash field if non-nil, zero value otherwise.

### GetParentHashOk

`func (o *Cast) GetParentHashOk() (*string, bool)`

GetParentHashOk returns a tuple with the ParentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentHash

`func (o *Cast) SetParentHash(v string)`

SetParentHash sets ParentHash field to given value.

### HasParentHash

`func (o *Cast) HasParentHash() bool`

HasParentHash returns a boolean if a field has been set.

### SetParentHashNil

`func (o *Cast) SetParentHashNil(b bool)`

 SetParentHashNil sets the value for ParentHash to be an explicit nil

### UnsetParentHash
`func (o *Cast) UnsetParentHash()`

UnsetParentHash ensures that no value is present for ParentHash, not even an explicit nil
### GetParentUrl

`func (o *Cast) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *Cast) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *Cast) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *Cast) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### SetParentUrlNil

`func (o *Cast) SetParentUrlNil(b bool)`

 SetParentUrlNil sets the value for ParentUrl to be an explicit nil

### UnsetParentUrl
`func (o *Cast) UnsetParentUrl()`

UnsetParentUrl ensures that no value is present for ParentUrl, not even an explicit nil
### GetRootParentUrl

`func (o *Cast) GetRootParentUrl() string`

GetRootParentUrl returns the RootParentUrl field if non-nil, zero value otherwise.

### GetRootParentUrlOk

`func (o *Cast) GetRootParentUrlOk() (*string, bool)`

GetRootParentUrlOk returns a tuple with the RootParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootParentUrl

`func (o *Cast) SetRootParentUrl(v string)`

SetRootParentUrl sets RootParentUrl field to given value.

### HasRootParentUrl

`func (o *Cast) HasRootParentUrl() bool`

HasRootParentUrl returns a boolean if a field has been set.

### SetRootParentUrlNil

`func (o *Cast) SetRootParentUrlNil(b bool)`

 SetRootParentUrlNil sets the value for RootParentUrl to be an explicit nil

### UnsetRootParentUrl
`func (o *Cast) UnsetRootParentUrl()`

UnsetRootParentUrl ensures that no value is present for RootParentUrl, not even an explicit nil
### GetParentAuthor

`func (o *Cast) GetParentAuthor() CastParentAuthor`

GetParentAuthor returns the ParentAuthor field if non-nil, zero value otherwise.

### GetParentAuthorOk

`func (o *Cast) GetParentAuthorOk() (*CastParentAuthor, bool)`

GetParentAuthorOk returns a tuple with the ParentAuthor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAuthor

`func (o *Cast) SetParentAuthor(v CastParentAuthor)`

SetParentAuthor sets ParentAuthor field to given value.

### HasParentAuthor

`func (o *Cast) HasParentAuthor() bool`

HasParentAuthor returns a boolean if a field has been set.

### GetAuthor

`func (o *Cast) GetAuthor() User`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Cast) GetAuthorOk() (*User, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Cast) SetAuthor(v User)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Cast) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetText

`func (o *Cast) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Cast) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Cast) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Cast) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTimestamp

`func (o *Cast) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Cast) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Cast) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Cast) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEmbeds

`func (o *Cast) GetEmbeds() []EmbeddedCast`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *Cast) GetEmbedsOk() (*[]EmbeddedCast, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *Cast) SetEmbeds(v []EmbeddedCast)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *Cast) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetType

`func (o *Cast) GetType() CastNotificationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cast) GetTypeOk() (*CastNotificationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cast) SetType(v CastNotificationType)`

SetType sets Type field to given value.

### HasType

`func (o *Cast) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


