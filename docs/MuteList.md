# MuteList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | **string** |  | 
**Muted** | [**User**](User.md) |  | 
**MutedAt** | **time.Time** |  | 

## Methods

### NewMuteList

`func NewMuteList(object string, muted User, mutedAt time.Time, ) *MuteList`

NewMuteList instantiates a new MuteList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMuteListWithDefaults

`func NewMuteListWithDefaults() *MuteList`

NewMuteListWithDefaults instantiates a new MuteList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *MuteList) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MuteList) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MuteList) SetObject(v string)`

SetObject sets Object field to given value.


### GetMuted

`func (o *MuteList) GetMuted() User`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *MuteList) GetMutedOk() (*User, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *MuteList) SetMuted(v User)`

SetMuted sets Muted field to given value.


### GetMutedAt

`func (o *MuteList) GetMutedAt() time.Time`

GetMutedAt returns the MutedAt field if non-nil, zero value otherwise.

### GetMutedAtOk

`func (o *MuteList) GetMutedAtOk() (*time.Time, bool)`

GetMutedAtOk returns a tuple with the MutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutedAt

`func (o *MuteList) SetMutedAt(v time.Time)`

SetMutedAt sets MutedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


