# ChannelSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | Pointer to [**[]Channel**](Channel.md) |  | [optional] 
**Next** | Pointer to [**NextCursor**](NextCursor.md) |  | [optional] 

## Methods

### NewChannelSearchResponse

`func NewChannelSearchResponse() *ChannelSearchResponse`

NewChannelSearchResponse instantiates a new ChannelSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelSearchResponseWithDefaults

`func NewChannelSearchResponseWithDefaults() *ChannelSearchResponse`

NewChannelSearchResponseWithDefaults instantiates a new ChannelSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannels

`func (o *ChannelSearchResponse) GetChannels() []Channel`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *ChannelSearchResponse) GetChannelsOk() (*[]Channel, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *ChannelSearchResponse) SetChannels(v []Channel)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *ChannelSearchResponse) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetNext

`func (o *ChannelSearchResponse) GetNext() NextCursor`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ChannelSearchResponse) GetNextOk() (*NextCursor, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ChannelSearchResponse) SetNext(v NextCursor)`

SetNext sets Next field to given value.

### HasNext

`func (o *ChannelSearchResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


