# ChannelOrDehydratedChannel

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

### NewChannelOrDehydratedChannel

`func NewChannelOrDehydratedChannel() *ChannelOrDehydratedChannel`

NewChannelOrDehydratedChannel instantiates a new ChannelOrDehydratedChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelOrDehydratedChannelWithDefaults

`func NewChannelOrDehydratedChannelWithDefaults() *ChannelOrDehydratedChannel`

NewChannelOrDehydratedChannelWithDefaults instantiates a new ChannelOrDehydratedChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChannelOrDehydratedChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChannelOrDehydratedChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChannelOrDehydratedChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChannelOrDehydratedChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ChannelOrDehydratedChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChannelOrDehydratedChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChannelOrDehydratedChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ChannelOrDehydratedChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *ChannelOrDehydratedChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChannelOrDehydratedChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChannelOrDehydratedChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChannelOrDehydratedChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ChannelOrDehydratedChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ChannelOrDehydratedChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ChannelOrDehydratedChannel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ChannelOrDehydratedChannel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObject

`func (o *ChannelOrDehydratedChannel) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChannelOrDehydratedChannel) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChannelOrDehydratedChannel) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ChannelOrDehydratedChannel) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChannelOrDehydratedChannel) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChannelOrDehydratedChannel) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChannelOrDehydratedChannel) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChannelOrDehydratedChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFollowerCount

`func (o *ChannelOrDehydratedChannel) GetFollowerCount() float32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *ChannelOrDehydratedChannel) GetFollowerCountOk() (*float32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *ChannelOrDehydratedChannel) SetFollowerCount(v float32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *ChannelOrDehydratedChannel) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### GetImageUrl

`func (o *ChannelOrDehydratedChannel) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ChannelOrDehydratedChannel) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ChannelOrDehydratedChannel) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ChannelOrDehydratedChannel) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetParentUrl

`func (o *ChannelOrDehydratedChannel) GetParentUrl() string`

GetParentUrl returns the ParentUrl field if non-nil, zero value otherwise.

### GetParentUrlOk

`func (o *ChannelOrDehydratedChannel) GetParentUrlOk() (*string, bool)`

GetParentUrlOk returns a tuple with the ParentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUrl

`func (o *ChannelOrDehydratedChannel) SetParentUrl(v string)`

SetParentUrl sets ParentUrl field to given value.

### HasParentUrl

`func (o *ChannelOrDehydratedChannel) HasParentUrl() bool`

HasParentUrl returns a boolean if a field has been set.

### GetLead

`func (o *ChannelOrDehydratedChannel) GetLead() User`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *ChannelOrDehydratedChannel) GetLeadOk() (*User, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *ChannelOrDehydratedChannel) SetLead(v User)`

SetLead sets Lead field to given value.

### HasLead

`func (o *ChannelOrDehydratedChannel) HasLead() bool`

HasLead returns a boolean if a field has been set.

### GetModerator

`func (o *ChannelOrDehydratedChannel) GetModerator() User`

GetModerator returns the Moderator field if non-nil, zero value otherwise.

### GetModeratorOk

`func (o *ChannelOrDehydratedChannel) GetModeratorOk() (*User, bool)`

GetModeratorOk returns a tuple with the Moderator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerator

`func (o *ChannelOrDehydratedChannel) SetModerator(v User)`

SetModerator sets Moderator field to given value.

### HasModerator

`func (o *ChannelOrDehydratedChannel) HasModerator() bool`

HasModerator returns a boolean if a field has been set.

### GetHosts

`func (o *ChannelOrDehydratedChannel) GetHosts() []User`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ChannelOrDehydratedChannel) GetHostsOk() (*[]User, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ChannelOrDehydratedChannel) SetHosts(v []User)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ChannelOrDehydratedChannel) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetViewerContext

`func (o *ChannelOrDehydratedChannel) GetViewerContext() ChannelViewerContext`

GetViewerContext returns the ViewerContext field if non-nil, zero value otherwise.

### GetViewerContextOk

`func (o *ChannelOrDehydratedChannel) GetViewerContextOk() (*ChannelViewerContext, bool)`

GetViewerContextOk returns a tuple with the ViewerContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewerContext

`func (o *ChannelOrDehydratedChannel) SetViewerContext(v ChannelViewerContext)`

SetViewerContext sets ViewerContext field to given value.

### HasViewerContext

`func (o *ChannelOrDehydratedChannel) HasViewerContext() bool`

HasViewerContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


