# GetUserDataByFid200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**HashScheme** | Pointer to [**HashScheme**](HashScheme.md) |  | [optional] [default to HASHSCHEME_HASH_SCHEME_BLAKE3]
**Signature** | Pointer to **string** |  | [optional] 
**SignatureScheme** | Pointer to [**SignatureScheme**](SignatureScheme.md) |  | [optional] [default to SIGNATURESCHEME_ED25519]
**Signer** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**UserDataAddAllOfData**](UserDataAddAllOfData.md) |  | [optional] 
**Messages** | Pointer to [**[]UserDataAdd**](UserDataAdd.md) |  | [optional] 
**NextPageToken** | Pointer to **string** |  | [optional] 

## Methods

### NewGetUserDataByFid200Response

`func NewGetUserDataByFid200Response() *GetUserDataByFid200Response`

NewGetUserDataByFid200Response instantiates a new GetUserDataByFid200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserDataByFid200ResponseWithDefaults

`func NewGetUserDataByFid200ResponseWithDefaults() *GetUserDataByFid200Response`

NewGetUserDataByFid200ResponseWithDefaults instantiates a new GetUserDataByFid200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *GetUserDataByFid200Response) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *GetUserDataByFid200Response) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *GetUserDataByFid200Response) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *GetUserDataByFid200Response) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHashScheme

`func (o *GetUserDataByFid200Response) GetHashScheme() HashScheme`

GetHashScheme returns the HashScheme field if non-nil, zero value otherwise.

### GetHashSchemeOk

`func (o *GetUserDataByFid200Response) GetHashSchemeOk() (*HashScheme, bool)`

GetHashSchemeOk returns a tuple with the HashScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashScheme

`func (o *GetUserDataByFid200Response) SetHashScheme(v HashScheme)`

SetHashScheme sets HashScheme field to given value.

### HasHashScheme

`func (o *GetUserDataByFid200Response) HasHashScheme() bool`

HasHashScheme returns a boolean if a field has been set.

### GetSignature

`func (o *GetUserDataByFid200Response) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GetUserDataByFid200Response) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GetUserDataByFid200Response) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *GetUserDataByFid200Response) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureScheme

`func (o *GetUserDataByFid200Response) GetSignatureScheme() SignatureScheme`

GetSignatureScheme returns the SignatureScheme field if non-nil, zero value otherwise.

### GetSignatureSchemeOk

`func (o *GetUserDataByFid200Response) GetSignatureSchemeOk() (*SignatureScheme, bool)`

GetSignatureSchemeOk returns a tuple with the SignatureScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScheme

`func (o *GetUserDataByFid200Response) SetSignatureScheme(v SignatureScheme)`

SetSignatureScheme sets SignatureScheme field to given value.

### HasSignatureScheme

`func (o *GetUserDataByFid200Response) HasSignatureScheme() bool`

HasSignatureScheme returns a boolean if a field has been set.

### GetSigner

`func (o *GetUserDataByFid200Response) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *GetUserDataByFid200Response) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *GetUserDataByFid200Response) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *GetUserDataByFid200Response) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetData

`func (o *GetUserDataByFid200Response) GetData() UserDataAddAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserDataByFid200Response) GetDataOk() (*UserDataAddAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserDataByFid200Response) SetData(v UserDataAddAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserDataByFid200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessages

`func (o *GetUserDataByFid200Response) GetMessages() []UserDataAdd`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetUserDataByFid200Response) GetMessagesOk() (*[]UserDataAdd, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetUserDataByFid200Response) SetMessages(v []UserDataAdd)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetUserDataByFid200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetNextPageToken

`func (o *GetUserDataByFid200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GetUserDataByFid200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GetUserDataByFid200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GetUserDataByFid200Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


