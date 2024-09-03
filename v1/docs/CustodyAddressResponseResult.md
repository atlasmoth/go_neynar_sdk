# CustodyAddressResponseResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** | User identifier (unsigned integer) | [default to 3]
**CustodyAddress** | **NullableString** |  | 

## Methods

### NewCustodyAddressResponseResult

`func NewCustodyAddressResponseResult(fid int32, custodyAddress NullableString, ) *CustodyAddressResponseResult`

NewCustodyAddressResponseResult instantiates a new CustodyAddressResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustodyAddressResponseResultWithDefaults

`func NewCustodyAddressResponseResultWithDefaults() *CustodyAddressResponseResult`

NewCustodyAddressResponseResultWithDefaults instantiates a new CustodyAddressResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *CustodyAddressResponseResult) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *CustodyAddressResponseResult) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *CustodyAddressResponseResult) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetCustodyAddress

`func (o *CustodyAddressResponseResult) GetCustodyAddress() string`

GetCustodyAddress returns the CustodyAddress field if non-nil, zero value otherwise.

### GetCustodyAddressOk

`func (o *CustodyAddressResponseResult) GetCustodyAddressOk() (*string, bool)`

GetCustodyAddressOk returns a tuple with the CustodyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustodyAddress

`func (o *CustodyAddressResponseResult) SetCustodyAddress(v string)`

SetCustodyAddress sets CustodyAddress field to given value.


### SetCustodyAddressNil

`func (o *CustodyAddressResponseResult) SetCustodyAddressNil(b bool)`

 SetCustodyAddressNil sets the value for CustodyAddress to be an explicit nil

### UnsetCustodyAddress
`func (o *CustodyAddressResponseResult) UnsetCustodyAddress()`

UnsetCustodyAddress ensures that no value is present for CustodyAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


