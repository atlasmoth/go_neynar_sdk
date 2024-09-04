# PostCastResponseCast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** | Ethereum address | [optional] 
**Author** | Pointer to [**PostCastResponseCastAuthor**](PostCastResponseCastAuthor.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 

## Methods

### NewPostCastResponseCast

`func NewPostCastResponseCast() *PostCastResponseCast`

NewPostCastResponseCast instantiates a new PostCastResponseCast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCastResponseCastWithDefaults

`func NewPostCastResponseCastWithDefaults() *PostCastResponseCast`

NewPostCastResponseCastWithDefaults instantiates a new PostCastResponseCast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *PostCastResponseCast) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *PostCastResponseCast) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *PostCastResponseCast) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *PostCastResponseCast) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetAuthor

`func (o *PostCastResponseCast) GetAuthor() PostCastResponseCastAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PostCastResponseCast) GetAuthorOk() (*PostCastResponseCastAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PostCastResponseCast) SetAuthor(v PostCastResponseCastAuthor)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *PostCastResponseCast) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetText

`func (o *PostCastResponseCast) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostCastResponseCast) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostCastResponseCast) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostCastResponseCast) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


