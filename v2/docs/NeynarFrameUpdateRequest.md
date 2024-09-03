# NeynarFrameUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The UUID of the frame to update. | 
**Name** | Pointer to **string** | The name of the frame. | [optional] 
**Pages** | [**[]NeynarFramePage**](NeynarFramePage.md) |  | 

## Methods

### NewNeynarFrameUpdateRequest

`func NewNeynarFrameUpdateRequest(uuid string, pages []NeynarFramePage, ) *NeynarFrameUpdateRequest`

NewNeynarFrameUpdateRequest instantiates a new NeynarFrameUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeynarFrameUpdateRequestWithDefaults

`func NewNeynarFrameUpdateRequestWithDefaults() *NeynarFrameUpdateRequest`

NewNeynarFrameUpdateRequestWithDefaults instantiates a new NeynarFrameUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NeynarFrameUpdateRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NeynarFrameUpdateRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NeynarFrameUpdateRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *NeynarFrameUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NeynarFrameUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NeynarFrameUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NeynarFrameUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPages

`func (o *NeynarFrameUpdateRequest) GetPages() []NeynarFramePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NeynarFrameUpdateRequest) GetPagesOk() (*[]NeynarFramePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NeynarFrameUpdateRequest) SetPages(v []NeynarFramePage)`

SetPages sets Pages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


