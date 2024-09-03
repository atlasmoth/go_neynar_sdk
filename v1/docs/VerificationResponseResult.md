# VerificationResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **string** |  | 
**Username** | **string** |  | 
**DisplayName** | **string** |  | 
**Verifications** | **[]string** |  | 

## Methods

### NewVerificationResponseResult

`func NewVerificationResponseResult(fid string, username string, displayName string, verifications []string, ) *VerificationResponseResult`

NewVerificationResponseResult instantiates a new VerificationResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationResponseResultWithDefaults

`func NewVerificationResponseResultWithDefaults() *VerificationResponseResult`

NewVerificationResponseResultWithDefaults instantiates a new VerificationResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *VerificationResponseResult) GetFid() string`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *VerificationResponseResult) GetFidOk() (*string, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *VerificationResponseResult) SetFid(v string)`

SetFid sets Fid field to given value.


### GetUsername

`func (o *VerificationResponseResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VerificationResponseResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VerificationResponseResult) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *VerificationResponseResult) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VerificationResponseResult) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VerificationResponseResult) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetVerifications

`func (o *VerificationResponseResult) GetVerifications() []string`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *VerificationResponseResult) GetVerificationsOk() (*[]string, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *VerificationResponseResult) SetVerifications(v []string)`

SetVerifications sets Verifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


