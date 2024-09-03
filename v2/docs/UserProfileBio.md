# UserProfileBio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**MentionedProfiles** | **[]string** |  | [default to []]

## Methods

### NewUserProfileBio

`func NewUserProfileBio(text string, mentionedProfiles []string, ) *UserProfileBio`

NewUserProfileBio instantiates a new UserProfileBio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserProfileBioWithDefaults

`func NewUserProfileBioWithDefaults() *UserProfileBio`

NewUserProfileBioWithDefaults instantiates a new UserProfileBio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *UserProfileBio) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *UserProfileBio) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *UserProfileBio) SetText(v string)`

SetText sets Text field to given value.


### GetMentionedProfiles

`func (o *UserProfileBio) GetMentionedProfiles() []string`

GetMentionedProfiles returns the MentionedProfiles field if non-nil, zero value otherwise.

### GetMentionedProfilesOk

`func (o *UserProfileBio) GetMentionedProfilesOk() (*[]string, bool)`

GetMentionedProfilesOk returns a tuple with the MentionedProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionedProfiles

`func (o *UserProfileBio) SetMentionedProfiles(v []string)`

SetMentionedProfiles sets MentionedProfiles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


