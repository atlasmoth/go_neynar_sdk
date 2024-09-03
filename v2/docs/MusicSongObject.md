# MusicSongObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disc** | Pointer to **string** |  | [optional] 
**Track** | Pointer to **int32** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewMusicSongObject

`func NewMusicSongObject(url string, ) *MusicSongObject`

NewMusicSongObject instantiates a new MusicSongObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMusicSongObjectWithDefaults

`func NewMusicSongObjectWithDefaults() *MusicSongObject`

NewMusicSongObjectWithDefaults instantiates a new MusicSongObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisc

`func (o *MusicSongObject) GetDisc() string`

GetDisc returns the Disc field if non-nil, zero value otherwise.

### GetDiscOk

`func (o *MusicSongObject) GetDiscOk() (*string, bool)`

GetDiscOk returns a tuple with the Disc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisc

`func (o *MusicSongObject) SetDisc(v string)`

SetDisc sets Disc field to given value.

### HasDisc

`func (o *MusicSongObject) HasDisc() bool`

HasDisc returns a boolean if a field has been set.

### GetTrack

`func (o *MusicSongObject) GetTrack() int32`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *MusicSongObject) GetTrackOk() (*int32, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *MusicSongObject) SetTrack(v int32)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *MusicSongObject) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetUrl

`func (o *MusicSongObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MusicSongObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MusicSongObject) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


