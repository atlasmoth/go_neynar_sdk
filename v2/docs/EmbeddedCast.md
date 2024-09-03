# EmbeddedCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Metadata** | Pointer to [**EmbedUrlMetadata**](EmbedUrlMetadata.md) |  | [optional] 
**CastId** | [**CastId**](CastId.md) |  | 

## Methods

### NewEmbeddedCast

`func NewEmbeddedCast(url string, castId CastId, ) *EmbeddedCast`

NewEmbeddedCast instantiates a new EmbeddedCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddedCastWithDefaults

`func NewEmbeddedCastWithDefaults() *EmbeddedCast`

NewEmbeddedCastWithDefaults instantiates a new EmbeddedCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *EmbeddedCast) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EmbeddedCast) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EmbeddedCast) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMetadata

`func (o *EmbeddedCast) GetMetadata() EmbedUrlMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EmbeddedCast) GetMetadataOk() (*EmbedUrlMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EmbeddedCast) SetMetadata(v EmbedUrlMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EmbeddedCast) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCastId

`func (o *EmbeddedCast) GetCastId() CastId`

GetCastId returns the CastId field if non-nil, zero value otherwise.

### GetCastIdOk

`func (o *EmbeddedCast) GetCastIdOk() (*CastId, bool)`

GetCastIdOk returns a tuple with the CastId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastId

`func (o *EmbeddedCast) SetCastId(v CastId)`

SetCastId sets CastId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


