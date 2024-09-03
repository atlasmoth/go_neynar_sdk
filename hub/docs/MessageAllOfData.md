# MessageAllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fid** | **int32** |  | 
**Timestamp** | **int64** |  | 
**Network** | [**FarcasterNetwork**](FarcasterNetwork.md) |  | [default to FARCASTERNETWORK_MAINNET]
**CastAddBody** | [**CastAddBody**](CastAddBody.md) |  | 
**CastRemoveBody** | [**CastRemoveBody**](CastRemoveBody.md) |  | 
**ReactionBody** | [**ReactionBody**](ReactionBody.md) |  | 
**LinkBody** | [**LinkBody**](LinkBody.md) |  | 
**VerificationAddEthAddressBody** | [**VerificationAddEthAddressBody**](VerificationAddEthAddressBody.md) |  | 
**VerificationRemoveBody** | [**VerificationRemoveBody**](VerificationRemoveBody.md) |  | 
**UserDataBody** | [**UserDataBody**](UserDataBody.md) |  | 
**UsernameProofBody** | [**UserNameProof**](UserNameProof.md) |  | 
**FrameActionBody** | [**FrameActionBody**](FrameActionBody.md) |  | 

## Methods

### NewMessageAllOfData

`func NewMessageAllOfData(fid int32, timestamp int64, network FarcasterNetwork, castAddBody CastAddBody, castRemoveBody CastRemoveBody, reactionBody ReactionBody, linkBody LinkBody, verificationAddEthAddressBody VerificationAddEthAddressBody, verificationRemoveBody VerificationRemoveBody, userDataBody UserDataBody, usernameProofBody UserNameProof, frameActionBody FrameActionBody, ) *MessageAllOfData`

NewMessageAllOfData instantiates a new MessageAllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageAllOfDataWithDefaults

`func NewMessageAllOfDataWithDefaults() *MessageAllOfData`

NewMessageAllOfDataWithDefaults instantiates a new MessageAllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFid

`func (o *MessageAllOfData) GetFid() int32`

GetFid returns the Fid field if non-nil, zero value otherwise.

### GetFidOk

`func (o *MessageAllOfData) GetFidOk() (*int32, bool)`

GetFidOk returns a tuple with the Fid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFid

`func (o *MessageAllOfData) SetFid(v int32)`

SetFid sets Fid field to given value.


### GetTimestamp

`func (o *MessageAllOfData) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MessageAllOfData) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MessageAllOfData) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetNetwork

`func (o *MessageAllOfData) GetNetwork() FarcasterNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MessageAllOfData) GetNetworkOk() (*FarcasterNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MessageAllOfData) SetNetwork(v FarcasterNetwork)`

SetNetwork sets Network field to given value.


### GetCastAddBody

`func (o *MessageAllOfData) GetCastAddBody() CastAddBody`

GetCastAddBody returns the CastAddBody field if non-nil, zero value otherwise.

### GetCastAddBodyOk

`func (o *MessageAllOfData) GetCastAddBodyOk() (*CastAddBody, bool)`

GetCastAddBodyOk returns a tuple with the CastAddBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastAddBody

`func (o *MessageAllOfData) SetCastAddBody(v CastAddBody)`

SetCastAddBody sets CastAddBody field to given value.


### GetCastRemoveBody

`func (o *MessageAllOfData) GetCastRemoveBody() CastRemoveBody`

GetCastRemoveBody returns the CastRemoveBody field if non-nil, zero value otherwise.

### GetCastRemoveBodyOk

`func (o *MessageAllOfData) GetCastRemoveBodyOk() (*CastRemoveBody, bool)`

GetCastRemoveBodyOk returns a tuple with the CastRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCastRemoveBody

`func (o *MessageAllOfData) SetCastRemoveBody(v CastRemoveBody)`

SetCastRemoveBody sets CastRemoveBody field to given value.


### GetReactionBody

`func (o *MessageAllOfData) GetReactionBody() ReactionBody`

GetReactionBody returns the ReactionBody field if non-nil, zero value otherwise.

### GetReactionBodyOk

`func (o *MessageAllOfData) GetReactionBodyOk() (*ReactionBody, bool)`

GetReactionBodyOk returns a tuple with the ReactionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactionBody

`func (o *MessageAllOfData) SetReactionBody(v ReactionBody)`

SetReactionBody sets ReactionBody field to given value.


### GetLinkBody

`func (o *MessageAllOfData) GetLinkBody() LinkBody`

GetLinkBody returns the LinkBody field if non-nil, zero value otherwise.

### GetLinkBodyOk

`func (o *MessageAllOfData) GetLinkBodyOk() (*LinkBody, bool)`

GetLinkBodyOk returns a tuple with the LinkBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBody

`func (o *MessageAllOfData) SetLinkBody(v LinkBody)`

SetLinkBody sets LinkBody field to given value.


### GetVerificationAddEthAddressBody

`func (o *MessageAllOfData) GetVerificationAddEthAddressBody() VerificationAddEthAddressBody`

GetVerificationAddEthAddressBody returns the VerificationAddEthAddressBody field if non-nil, zero value otherwise.

### GetVerificationAddEthAddressBodyOk

`func (o *MessageAllOfData) GetVerificationAddEthAddressBodyOk() (*VerificationAddEthAddressBody, bool)`

GetVerificationAddEthAddressBodyOk returns a tuple with the VerificationAddEthAddressBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationAddEthAddressBody

`func (o *MessageAllOfData) SetVerificationAddEthAddressBody(v VerificationAddEthAddressBody)`

SetVerificationAddEthAddressBody sets VerificationAddEthAddressBody field to given value.


### GetVerificationRemoveBody

`func (o *MessageAllOfData) GetVerificationRemoveBody() VerificationRemoveBody`

GetVerificationRemoveBody returns the VerificationRemoveBody field if non-nil, zero value otherwise.

### GetVerificationRemoveBodyOk

`func (o *MessageAllOfData) GetVerificationRemoveBodyOk() (*VerificationRemoveBody, bool)`

GetVerificationRemoveBodyOk returns a tuple with the VerificationRemoveBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRemoveBody

`func (o *MessageAllOfData) SetVerificationRemoveBody(v VerificationRemoveBody)`

SetVerificationRemoveBody sets VerificationRemoveBody field to given value.


### GetUserDataBody

`func (o *MessageAllOfData) GetUserDataBody() UserDataBody`

GetUserDataBody returns the UserDataBody field if non-nil, zero value otherwise.

### GetUserDataBodyOk

`func (o *MessageAllOfData) GetUserDataBodyOk() (*UserDataBody, bool)`

GetUserDataBodyOk returns a tuple with the UserDataBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataBody

`func (o *MessageAllOfData) SetUserDataBody(v UserDataBody)`

SetUserDataBody sets UserDataBody field to given value.


### GetUsernameProofBody

`func (o *MessageAllOfData) GetUsernameProofBody() UserNameProof`

GetUsernameProofBody returns the UsernameProofBody field if non-nil, zero value otherwise.

### GetUsernameProofBodyOk

`func (o *MessageAllOfData) GetUsernameProofBodyOk() (*UserNameProof, bool)`

GetUsernameProofBodyOk returns a tuple with the UsernameProofBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameProofBody

`func (o *MessageAllOfData) SetUsernameProofBody(v UserNameProof)`

SetUsernameProofBody sets UsernameProofBody field to given value.


### GetFrameActionBody

`func (o *MessageAllOfData) GetFrameActionBody() FrameActionBody`

GetFrameActionBody returns the FrameActionBody field if non-nil, zero value otherwise.

### GetFrameActionBodyOk

`func (o *MessageAllOfData) GetFrameActionBodyOk() (*FrameActionBody, bool)`

GetFrameActionBodyOk returns a tuple with the FrameActionBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameActionBody

`func (o *MessageAllOfData) SetFrameActionBody(v FrameActionBody)`

SetFrameActionBody sets FrameActionBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


