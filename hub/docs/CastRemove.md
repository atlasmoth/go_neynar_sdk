# CastRemove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | Pointer to **string** |  | [optional] 
**HashScheme** | Pointer to [**HashScheme**](HashScheme.md) |  | [optional] [default to HASHSCHEME_HASH_SCHEME_BLAKE3]
**Signature** | Pointer to **string** |  | [optional] 
**SignatureScheme** | Pointer to [**SignatureScheme**](SignatureScheme.md) |  | [optional] [default to SIGNATURESCHEME_ED25519]
**Signer** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CastRemoveAllOfData**](CastRemoveAllOfData.md) |  | [optional] 

## Methods

### NewCastRemove

`func NewCastRemove() *CastRemove`

NewCastRemove instantiates a new CastRemove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCastRemoveWithDefaults

`func NewCastRemoveWithDefaults() *CastRemove`

NewCastRemoveWithDefaults instantiates a new CastRemove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *CastRemove) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *CastRemove) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *CastRemove) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *CastRemove) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetHashScheme

`func (o *CastRemove) GetHashScheme() HashScheme`

GetHashScheme returns the HashScheme field if non-nil, zero value otherwise.

### GetHashSchemeOk

`func (o *CastRemove) GetHashSchemeOk() (*HashScheme, bool)`

GetHashSchemeOk returns a tuple with the HashScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashScheme

`func (o *CastRemove) SetHashScheme(v HashScheme)`

SetHashScheme sets HashScheme field to given value.

### HasHashScheme

`func (o *CastRemove) HasHashScheme() bool`

HasHashScheme returns a boolean if a field has been set.

### GetSignature

`func (o *CastRemove) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CastRemove) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CastRemove) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CastRemove) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSignatureScheme

`func (o *CastRemove) GetSignatureScheme() SignatureScheme`

GetSignatureScheme returns the SignatureScheme field if non-nil, zero value otherwise.

### GetSignatureSchemeOk

`func (o *CastRemove) GetSignatureSchemeOk() (*SignatureScheme, bool)`

GetSignatureSchemeOk returns a tuple with the SignatureScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureScheme

`func (o *CastRemove) SetSignatureScheme(v SignatureScheme)`

SetSignatureScheme sets SignatureScheme field to given value.

### HasSignatureScheme

`func (o *CastRemove) HasSignatureScheme() bool`

HasSignatureScheme returns a boolean if a field has been set.

### GetSigner

`func (o *CastRemove) GetSigner() string`

GetSigner returns the Signer field if non-nil, zero value otherwise.

### GetSignerOk

`func (o *CastRemove) GetSignerOk() (*string, bool)`

GetSignerOk returns a tuple with the Signer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigner

`func (o *CastRemove) SetSigner(v string)`

SetSigner sets Signer field to given value.

### HasSigner

`func (o *CastRemove) HasSigner() bool`

HasSigner returns a boolean if a field has been set.

### GetData

`func (o *CastRemove) GetData() CastRemoveAllOfData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CastRemove) GetDataOk() (*CastRemoveAllOfData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CastRemove) SetData(v CastRemoveAllOfData)`

SetData sets Data field to given value.

### HasData

`func (o *CastRemove) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


