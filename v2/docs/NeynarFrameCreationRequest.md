# NeynarFrameCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the frame. | [optional] 
**Pages** | Pointer to [**[]NeynarFramePage**](NeynarFramePage.md) |  | [optional] 

## Methods

### NewNeynarFrameCreationRequest

`func NewNeynarFrameCreationRequest() *NeynarFrameCreationRequest`

NewNeynarFrameCreationRequest instantiates a new NeynarFrameCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeynarFrameCreationRequestWithDefaults

`func NewNeynarFrameCreationRequestWithDefaults() *NeynarFrameCreationRequest`

NewNeynarFrameCreationRequestWithDefaults instantiates a new NeynarFrameCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NeynarFrameCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NeynarFrameCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NeynarFrameCreationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NeynarFrameCreationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPages

`func (o *NeynarFrameCreationRequest) GetPages() []NeynarFramePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NeynarFrameCreationRequest) GetPagesOk() (*[]NeynarFramePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NeynarFrameCreationRequest) SetPages(v []NeynarFramePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *NeynarFrameCreationRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


