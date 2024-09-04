# Signer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignerUuid** | Pointer to **string** | UUID of the signer | [optional] 
**PublicKey** | Pointer to **string** | Ed25519 public key | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SignerApprovalUrl** | Pointer to **string** |  | [optional] 
**Fid** | Pointer to **int32** | User identifier (unsigned integer) | [optional] 

## Methods

### NewSigner

`func NewSigner() *Signer`

NewSigner instantiates a new Signer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignerWithDefaults

`func NewSignerWithDefaults() *Signer`

NewSignerWithDefaults instantiates a new Signer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignerUuid

`func (o *Signer) GetSignerUuid() string`

GetSignerUuid returns the SignerUuid field if non-nil, zero value otherwise.

### GetSignerUuidOk

`func (o *Signer) GetSignerUuidOk() (*string, bool)`

GetSignerUuidOk returns a tuple with the SignerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerUuid

`func (o *Signer) SetSignerUuid(v string)`

SetSignerUuid sets SignerUuid field to given value.

### HasSignerUuid

`func (o *Signer) HasSignerUuid() bool`

HasSignerUuid returns a boolean if a field has been set.

### GetPublicKey

`func (o *Signer) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *Signer) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *Signer) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *Signer) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetStatus

`func (o *Signer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Signer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Signer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Signer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSignerApprovalUrl

`func (o *Signer) GetSignerApprovalUrl() string`

GetSignerApprovalUrl returns the SignerApprovalUrl field if non-nil, zero value otherwise.

### GetSignerApprovalUrlOk

`func (o *Signer) GetSignerApprovalUrlOk() (*string, bool)`

GetSignerApprovalUrlOk returns a tuple with the SignerApprovalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignerApprovalUrl

`func (o *Signer) SetSignerApprovalUrl(v string)`

SetSignerApprovalUrl sets SignerApprovalUrl field to given value.

### HasSignerApprovalUrl

`func (o *Signer) HasSignerApprovalUrl() bool`

HasSignerApprovalUrl returns a boolean if a field has been set.

### GetFid

`func (o *Signer) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *Signer) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *Signer) SetFid(v int32)`

SetFid sets Fid field to given value.

### HasFid

`func (o *Signer) HasFid() bool`

HasFid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


