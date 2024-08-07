# UserProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bio** | [**UserProfileBio**](UserProfileBio.md) |  | 

## Methods

### NewUserProfile

`func NewUserProfile(bio UserProfileBio, ) *UserProfile`

NewUserProfile instantiates a new UserProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileWithDefaults

`func NewUserProfileWithDefaults() *UserProfile`

NewUserProfileWithDefaults instantiates a new UserProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBio

`func (o *UserProfile) GetBio() UserProfileBio`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UserProfile) GetBioOk() (*UserProfileBio, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UserProfile) SetBio(v UserProfileBio)`

SetBio sets Bio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


