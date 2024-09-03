## Documentation for API Endpoints

All URIs are relative to *https://hub-api.neynar.com*

| Class                | Method                                                                                              | HTTP request                                | Description                                                                  |
| -------------------- | --------------------------------------------------------------------------------------------------- | ------------------------------------------- | ---------------------------------------------------------------------------- |
| _CastsAPI_           | [**GetCastById**](docs/CastsAPI.md#getcastbyid)                                                     | **Get** /v1/castById                        | Get a cast by its FID and Hash.                                              |
| _CastsAPI_           | [**ListCastsByFid**](docs/CastsAPI.md#listcastsbyfid)                                               | **Get** /v1/castsByFid                      | Fetch all casts authored by an FID.                                          |
| _CastsAPI_           | [**ListCastsByMention**](docs/CastsAPI.md#listcastsbymention)                                       | **Get** /v1/castsByMention                  | Fetch all casts that mention an FID                                          |
| _CastsAPI_           | [**ListCastsByParent**](docs/CastsAPI.md#listcastsbyparent)                                         | **Get** /v1/castsByParent                   | Fetch all casts by parent cast&#39;s FID and Hash OR by the parent&#39;s URL |
| _FIDsAPI_            | [**ListFids**](docs/FIDsAPI.md#listfids)                                                            | **Get** /v1/fids                            | Get a list of all the FIDs                                                   |
| _HubEventsAPI_       | [**GetEventById**](docs/HubEventsAPI.md#geteventbyid)                                               | **Get** /v1/eventById                       | Get an event by its ID                                                       |
| _HubEventsAPI_       | [**ListEvents**](docs/HubEventsAPI.md#listevents)                                                   | **Get** /v1/events                          | Get a page of Hub events                                                     |
| _InfoAPI_            | [**GetInfo**](docs/InfoAPI.md#getinfo)                                                              | **Get** /v1/info                            | Sync Methods                                                                 |
| _LinksAPI_           | [**GetLinkById**](docs/LinksAPI.md#getlinkbyid)                                                     | **Get** /v1/linkById                        | Get a link by its FID and target FID.                                        |
| _LinksAPI_           | [**ListLinksByFid**](docs/LinksAPI.md#listlinksbyfid)                                               | **Get** /v1/linksByFid                      | Get all links from a source FID                                              |
| _LinksAPI_           | [**ListLinksByTargetFid**](docs/LinksAPI.md#listlinksbytargetfid)                                   | **Get** /v1/linksByTargetFid                | Get all links to a target FID                                                |
| _OnChainEventsAPI_   | [**GetOnChainIdRegistrationByAddress**](docs/OnChainEventsAPI.md#getonchainidregistrationbyaddress) | **Get** /v1/onChainIdRegistryEventByAddress | Get an on chain ID Registry Event for a given Address                        |
| _OnChainEventsAPI_   | [**ListOnChainEventsByFid**](docs/OnChainEventsAPI.md#listonchaineventsbyfid)                       | **Get** /v1/onChainEventsByFid              | Get a list of on-chain events provided by an FID                             |
| _OnChainEventsAPI_   | [**ListOnChainSignersByFid**](docs/OnChainEventsAPI.md#listonchainsignersbyfid)                     | **Get** /v1/onChainSignersByFid             | Get a list of signers provided by an FID                                     |
| _ReactionsAPI_       | [**GetReactionById**](docs/ReactionsAPI.md#getreactionbyid)                                         | **Get** /v1/reactionById                    | Get a reaction by its created FID and target Cast.                           |
| _ReactionsAPI_       | [**ListReactionsByCast**](docs/ReactionsAPI.md#listreactionsbycast)                                 | **Get** /v1/reactionsByCast                 | Get all reactions to a cast                                                  |
| _ReactionsAPI_       | [**ListReactionsByFid**](docs/ReactionsAPI.md#listreactionsbyfid)                                   | **Get** /v1/reactionsByFid                  | Get all reactions by an FID                                                  |
| _ReactionsAPI_       | [**ListReactionsByTarget**](docs/ReactionsAPI.md#listreactionsbytarget)                             | **Get** /v1/reactionsByTarget               | Get all reactions to a target URL                                            |
| _StorageAPI_         | [**GetStorageLimitsByFid**](docs/StorageAPI.md#getstoragelimitsbyfid)                               | **Get** /v1/storageLimitsByFid              | Get an FID&#39;s storage limits.                                             |
| _SubmitMessageAPI_   | [**SubmitMessage**](docs/SubmitMessageAPI.md#submitmessage)                                         | **Post** /v1/submitMessage                  | Submit a signed protobuf-serialized message to the Hub                       |
| _UserDataAPI_        | [**GetUserDataByFid**](docs/UserDataAPI.md#getuserdatabyfid)                                        | **Get** /v1/userDataByFid                   | Get UserData for a FID.                                                      |
| _UsernamesAPI_       | [**GetUsernameProof**](docs/UsernamesAPI.md#getusernameproof)                                       | **Get** /v1/userNameProofByName             | Get an proof for a username by the Farcaster username                        |
| _UsernamesAPI_       | [**ListUsernameProofsByFid**](docs/UsernamesAPI.md#listusernameproofsbyfid)                         | **Get** /v1/userNameProofsByFid             | Get a list of proofs provided by an FID                                      |
| _ValidateMessageAPI_ | [**ValidateMessage**](docs/ValidateMessageAPI.md#validatemessage)                                   | **Post** /v1/validateMessage                | Validate a signed protobuf-serialized message with the Hub                   |
| _VerificationsAPI_   | [**ListVerificationsByFid**](docs/VerificationsAPI.md#listverificationsbyfid)                       | **Get** /v1/verificationsByFid              | Get a list of verifications provided by an FID                               |

## Documentation For Models

- [CastAdd](docs/CastAdd.md)
- [CastAddAllOfData](docs/CastAddAllOfData.md)
- [CastAddBody](docs/CastAddBody.md)
- [CastEmbed](docs/CastEmbed.md)
- [CastId](docs/CastId.md)
- [CastRemove](docs/CastRemove.md)
- [CastRemoveAllOfData](docs/CastRemoveAllOfData.md)
- [CastRemoveBody](docs/CastRemoveBody.md)
- [DbStats](docs/DbStats.md)
- [Embed](docs/Embed.md)
- [ErrorResponse](docs/ErrorResponse.md)
- [ErrorResponseMetadata](docs/ErrorResponseMetadata.md)
- [FarcasterNetwork](docs/FarcasterNetwork.md)
- [FidsResponse](docs/FidsResponse.md)
- [FrameActionBody](docs/FrameActionBody.md)
- [GetUserDataByFid200Response](docs/GetUserDataByFid200Response.md)
- [GetUserDataByFid200ResponseOneOf](docs/GetUserDataByFid200ResponseOneOf.md)
- [HashScheme](docs/HashScheme.md)
- [HubEvent](docs/HubEvent.md)
- [HubEventMergeMessage](docs/HubEventMergeMessage.md)
- [HubEventMergeOnChainEvent](docs/HubEventMergeOnChainEvent.md)
- [HubEventMergeUsernameProof](docs/HubEventMergeUsernameProof.md)
- [HubEventPruneMessage](docs/HubEventPruneMessage.md)
- [HubEventRevokeMessage](docs/HubEventRevokeMessage.md)
- [HubInfoResponse](docs/HubInfoResponse.md)
- [IdRegisterEventBody](docs/IdRegisterEventBody.md)
- [IdRegisterEventType](docs/IdRegisterEventType.md)
- [LinkAdd](docs/LinkAdd.md)
- [LinkAddAllOfData](docs/LinkAddAllOfData.md)
- [LinkBody](docs/LinkBody.md)
- [LinkRemove](docs/LinkRemove.md)
- [LinkRemoveAllOfData](docs/LinkRemoveAllOfData.md)
- [LinkType](docs/LinkType.md)
- [ListCastsByFid200Response](docs/ListCastsByFid200Response.md)
- [ListEvents200Response](docs/ListEvents200Response.md)
- [ListLinksByFid200Response](docs/ListLinksByFid200Response.md)
- [ListOnChainEventsByFid200Response](docs/ListOnChainEventsByFid200Response.md)
- [ListOnChainSignersByFid200Response](docs/ListOnChainSignersByFid200Response.md)
- [ListOnChainSignersByFid200ResponseOneOf](docs/ListOnChainSignersByFid200ResponseOneOf.md)
- [ListReactionsByCast200Response](docs/ListReactionsByCast200Response.md)
- [ListVerificationsByFid200Response](docs/ListVerificationsByFid200Response.md)
- [MergeMessageBody](docs/MergeMessageBody.md)
- [MergeOnChainEventBody](docs/MergeOnChainEventBody.md)
- [MergeUserNameProofBody](docs/MergeUserNameProofBody.md)
- [Message](docs/Message.md)
- [MessageAllOfData](docs/MessageAllOfData.md)
- [MessageCommon](docs/MessageCommon.md)
- [MessageDataCastAdd](docs/MessageDataCastAdd.md)
- [MessageDataCastRemove](docs/MessageDataCastRemove.md)
- [MessageDataCommon](docs/MessageDataCommon.md)
- [MessageDataFrameAction](docs/MessageDataFrameAction.md)
- [MessageDataLink](docs/MessageDataLink.md)
- [MessageDataReaction](docs/MessageDataReaction.md)
- [MessageDataUserDataAdd](docs/MessageDataUserDataAdd.md)
- [MessageDataUsernameProof](docs/MessageDataUsernameProof.md)
- [MessageDataVerificationAdd](docs/MessageDataVerificationAdd.md)
- [MessageDataVerificationRemove](docs/MessageDataVerificationRemove.md)
- [MessageType](docs/MessageType.md)
- [OnChainEvent](docs/OnChainEvent.md)
- [OnChainEventCommon](docs/OnChainEventCommon.md)
- [OnChainEventIdRegister](docs/OnChainEventIdRegister.md)
- [OnChainEventSigner](docs/OnChainEventSigner.md)
- [OnChainEventSignerMigrated](docs/OnChainEventSignerMigrated.md)
- [OnChainEventStorageRent](docs/OnChainEventStorageRent.md)
- [OnChainEventType](docs/OnChainEventType.md)
- [PruneMessageBody](docs/PruneMessageBody.md)
- [Reaction](docs/Reaction.md)
- [ReactionAllOfData](docs/ReactionAllOfData.md)
- [ReactionBody](docs/ReactionBody.md)
- [ReactionRemove](docs/ReactionRemove.md)
- [ReactionType](docs/ReactionType.md)
- [RevokeMessageBody](docs/RevokeMessageBody.md)
- [SignatureScheme](docs/SignatureScheme.md)
- [SignerEventBody](docs/SignerEventBody.md)
- [SignerEventType](docs/SignerEventType.md)
- [SignerMigratedEventBody](docs/SignerMigratedEventBody.md)
- [StorageLimit](docs/StorageLimit.md)
- [StorageLimitsResponse](docs/StorageLimitsResponse.md)
- [StorageRentEventBody](docs/StorageRentEventBody.md)
- [StoreType](docs/StoreType.md)
- [UrlEmbed](docs/UrlEmbed.md)
- [UserDataAdd](docs/UserDataAdd.md)
- [UserDataAddAllOfData](docs/UserDataAddAllOfData.md)
- [UserDataBody](docs/UserDataBody.md)
- [UserDataType](docs/UserDataType.md)
- [UserNameProof](docs/UserNameProof.md)
- [UserNameType](docs/UserNameType.md)
- [UsernameProofsResponse](docs/UsernameProofsResponse.md)
- [ValidateMessageResponse](docs/ValidateMessageResponse.md)
- [Verification](docs/Verification.md)
- [VerificationAddEthAddressBody](docs/VerificationAddEthAddressBody.md)
- [VerificationAllOfData](docs/VerificationAllOfData.md)
- [VerificationRemove](docs/VerificationRemove.md)
- [VerificationRemoveAllOfData](docs/VerificationRemoveAllOfData.md)
- [VerificationRemoveBody](docs/VerificationRemoveBody.md)

## Documentation For Authorization

Authentication schemes defined for the API:

### usernamePassword

- **Type**: HTTP basic authentication

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextBasicAuth, openapi.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

- `PtrBool`
- `PtrInt`
- `PtrInt32`
- `PtrInt64`
- `PtrFloat`
- `PtrFloat32`
- `PtrFloat64`
- `PtrString`
- `PtrTime`
