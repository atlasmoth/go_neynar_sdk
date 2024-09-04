# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **float32** | Epoch timestamp in seconds. | [optional] 
**FollowerCount** | Pointer to **float32** | Number of followers the channel has. | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**ParentUrl** | Pointer to **string** |  | [optional] 
**Lead** | Pointer to [**User**](User.md) |  | [optional] 
**Moderator** | Pointer to [**User**](User.md) |  | [optional] 
**Hosts** | Pointer to [**[]User**](User.md) |  | [optional] 
**ViewerContext** | Pointer to [**ChannelViewerContext**](ChannelViewerContext.md) |  | [optional] 

## Methods

### NewChannel

`func NewChannel() *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Channel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Channel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Channel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Channel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Channel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Channel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Channel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Channel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Channel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Channel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Channel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Channel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Channel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObject

`func (o *Channel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Channel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Channel) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Channel) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Channel) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Channel) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Channel) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Channel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFollowerCount

`func (o *Channel) GetFollowerCount() float32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *Channel) GetFollowerCountOk() (*float32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *Channel) SetFollowerCount(v float32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *Channel) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetImageUrl

`func (o *Channel) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Channel) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Channel) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Channel) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetParentUrl

`func (o *Channel) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *Channel) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *Channel) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *Channel) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetLead

`func (o *Channel) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *Channel) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *Channel) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *Channel) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetModerator

`func (o *Channel) GetModerator() User`

GetModerator returns the Moderator field if non-nil, zero value otherwise.

### GetModeratorOk

`func (o *Channel) GetModeratorOk() (*User, bool)`

GetModeratorOk returns a tuple with the Moderator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerator

`func (o *Channel) SetModerator(v User)`

SetModerator sets Moderator field to given value.

### HasModerator

`func (o *Channel) HasModerator() bool`

HasModerator returns a boolean if a field has been set.

### GetHosts

`func (o *Channel) GetHosts() []User`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *Channel) GetHostsOk() (*[]User, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *Channel) SetHosts(v []User)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *Channel) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetViewerContext

`func (o *Channel) GetViewerContext() ChannelViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *Channel) GetViewerContextOk() (*ChannelViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *Channel) SetViewerContext(v ChannelViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *Channel) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


