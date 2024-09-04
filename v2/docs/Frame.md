# Frame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version of the frame | [optional] 
**Image** | Pointer to **string** | URL of the image | [optional] 
**Buttons** | Pointer to [**[]FrameActionButton**](FrameActionButton.md) |  | [optional] 
**PostUrl** | Pointer to **string** | Post URL to take an action on this frame | [optional] 
**FramesUrl** | Pointer to **string** | URL of the frames | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**ImageAspectRatio** | Pointer to **string** |  | [optional] 
**Input** | Pointer to [**FrameInput**](FrameInput.md) |  | [optional] 
**State** | Pointer to [**FrameState**](FrameState.md) |  | [optional] 

## Methods

### NewFrame

`func NewFrame() *Frame`

NewFrame instantiates a new Frame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameWithDefaults

`func NewFrameWithDefaults() *Frame`

NewFrameWithDefaults instantiates a new Frame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *Frame) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Frame) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Frame) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Frame) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetImage

`func (o *Frame) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Frame) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Frame) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Frame) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetButtons

`func (o *Frame) GetButtons() []FrameActionButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *Frame) GetButtonsOk() (*[]FrameActionButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *Frame) SetButtons(v []FrameActionButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *Frame) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetPostUrl

`func (o *Frame) GetPostUrl() string`

GetPostUrl returns the PostUrl field if non-nil, zero value otherwise.

### GetPostUrlOk

`func (o *Frame) GetPostUrlOk() (*string, bool)`

GetPostUrlOk returns a tuple with the PostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrl

`func (o *Frame) SetPostUrl(v string)`

SetPostUrl sets PostUrl field to given value.

### HasPostUrl

`func (o *Frame) HasPostUrl() bool`

HasPostUrl returns a boolean if a field has been set.

### GetFramesUrl

`func (o *Frame) GetFramesUrl() string`

GetFramesUrl returns the FramesUrl field if non-nil, zero value otherwise.

### GetFramesUrlOk

`func (o *Frame) GetFramesUrlOk() (*string, bool)`

GetFramesUrlOk returns a tuple with the FramesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramesUrl

`func (o *Frame) SetFramesUrl(v string)`

SetFramesUrl sets FramesUrl field to given value.

### HasFramesUrl

`func (o *Frame) HasFramesUrl() bool`

HasFramesUrl returns a boolean if a field has been set.

### GetTitle

`func (o *Frame) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Frame) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Frame) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Frame) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetImageAspectRatio

`func (o *Frame) GetImageAspectRatio() string`

GetImageAspectRatio returns the ImageAspectRatio field if non-nil, zero value otherwise.

### GetImageAspectRatioOk

`func (o *Frame) GetImageAspectRatioOk() (*string, bool)`

GetImageAspectRatioOk returns a tuple with the ImageAspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageAspectRatio

`func (o *Frame) SetImageAspectRatio(v string)`

SetImageAspectRatio sets ImageAspectRatio field to given value.

### HasImageAspectRatio

`func (o *Frame) HasImageAspectRatio() bool`

HasImageAspectRatio returns a boolean if a field has been set.

### GetInput

`func (o *Frame) GetInput() FrameInput`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Frame) GetInputOk() (*FrameInput, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Frame) SetInput(v FrameInput)`

SetInput sets Input field to given value.

### HasInput

`func (o *Frame) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetState

`func (o *Frame) GetState() FrameState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Frame) GetStateOk() (*FrameState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Frame) SetState(v FrameState)`

SetState sets State field to given value.

### HasState

`func (o *Frame) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


