# NeynarFrame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique identifier for the frame. | [optional] 
**Name** | Pointer to **string** | Name of the frame. | [optional] 
**Link** | Pointer to **string** | Generated link for the frame&#39;s first page. | [optional] 
**Pages** | Pointer to [**[]NeynarFramePage**](NeynarFramePage.md) |  | [optional] 
**Valid** | Pointer to **bool** | Indicates if the frame is valid. | [optional] 

## Methods

### NewNeynarFrame

`func NewNeynarFrame() *NeynarFrame`

NewNeynarFrame instantiates a new NeynarFrame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNeynarFrameWithDefaults

`func NewNeynarFrameWithDefaults() *NeynarFrame`

NewNeynarFrameWithDefaults instantiates a new NeynarFrame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *NeynarFrame) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NeynarFrame) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NeynarFrame) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NeynarFrame) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *NeynarFrame) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NeynarFrame) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NeynarFrame) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NeynarFrame) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLink

`func (o *NeynarFrame) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *NeynarFrame) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *NeynarFrame) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *NeynarFrame) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetPages

`func (o *NeynarFrame) GetPages() []NeynarFramePage`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *NeynarFrame) GetPagesOk() (*[]NeynarFramePage, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *NeynarFrame) SetPages(v []NeynarFramePage)`

SetPages sets Pages field to given value.

### HasPages

`func (o *NeynarFrame) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetValid

`func (o *NeynarFrame) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *NeynarFrame) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *NeynarFrame) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *NeynarFrame) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


