# RegisterDeveloperManagedSignedKeyReqBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** | Ed25519 public key | 
**Signature** | **string** | Signature generated by the custody address of the app. Signed data includes app_fid, deadline, signer’s public key | 
**AppFid** | **int32** | User identifier (unsigned integer) | 
**Deadline** | **int32** | unix timestamp in seconds that controls how long the signed key request is valid for. (24 hours from now is recommended) | 

## Methods

### NewRegisterDeveloperManagedSignedKeyReqBody

`func NewRegisterDeveloperManagedSignedKeyReqBody(publicKey string, signature string, appFid int32, deadline int32, ) *RegisterDeveloperManagedSignedKeyReqBody`

NewRegisterDeveloperManagedSignedKeyReqBody instantiates a new RegisterDeveloperManagedSignedKeyReqBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults

`func NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults() *RegisterDeveloperManagedSignedKeyReqBody`

NewRegisterDeveloperManagedSignedKeyReqBodyWithDefaults instantiates a new RegisterDeveloperManagedSignedKeyReqBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *RegisterDeveloperManagedSignedKeyReqBody) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetSignature

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *RegisterDeveloperManagedSignedKeyReqBody) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetAppFid

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetAppFid() int32`

GetAppFid returns the AppFid field if non-nil, zero value otherwise.

### GetAppFidOk

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetAppFidOk() (*int32, bool)`

GetAppFidOk returns a tuple with the AppFid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppFid

`func (o *RegisterDeveloperManagedSignedKeyReqBody) SetAppFid(v int32)`

SetAppFid sets AppFid field to given value.


### GetDeadline

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetDeadline() int32`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *RegisterDeveloperManagedSignedKeyReqBody) GetDeadlineOk() (*int32, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *RegisterDeveloperManagedSignedKeyReqBody) SetDeadline(v int32)`

SetDeadline sets Deadline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


