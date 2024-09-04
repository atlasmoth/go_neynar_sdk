## Documentation for API Endpoints

All URIs are relative to *https://hub-api.neynar.com*

| Class                | Method                                                                                              | HTTP request                                | Description                                                                    |
| -------------------- | --------------------------------------------------------------------------------------------------- | ------------------------------------------- | ------------------------------------------------------------------------------ |
| _CastsApi_           | [**GetCastById**](docs/CastsApi.md#getcastbyid)                                                     | **Get** /v1/castById                        | Get a cast by its FID and Hash.                                                |
| _CastsApi_           | [**ListCastsByFid**](docs/CastsApi.md#listcastsbyfid)                                               | **Get** /v1/castsByFid                      | Fetch all casts authored by an FID.                                            |
| _CastsApi_           | [**ListCastsByMention**](docs/CastsApi.md#listcastsbymention)                                       | **Get** /v1/castsByMention                  | Fetch all casts that mention an FID                                            |
| _CastsApi_           | [**ListCastsByParent**](docs/CastsApi.md#listcastsbyparent)                                         | **Get** /v1/castsByParent                   | Fetch all casts by parent cast&#x27;s FID and Hash OR by the parent&#x27;s URL |
| _FIDsApi_            | [**ListFids**](docs/FIDsApi.md#listfids)                                                            | **Get** /v1/fids                            | Get a list of all the FIDs                                                     |
| _HubEventsApi_       | [**GetEventById**](docs/HubEventsApi.md#geteventbyid)                                               | **Get** /v1/eventById                       | Get an event by its ID                                                         |
| _HubEventsApi_       | [**ListEvents**](docs/HubEventsApi.md#listevents)                                                   | **Get** /v1/events                          | Get a page of Hub events                                                       |
| _InfoApi_            | [**GetInfo**](docs/InfoApi.md#getinfo)                                                              | **Get** /v1/info                            | Sync Methods                                                                   |
| _LinksApi_           | [**GetLinkById**](docs/LinksApi.md#getlinkbyid)                                                     | **Get** /v1/linkById                        | Get a link by its FID and target FID.                                          |
| _LinksApi_           | [**ListLinksByFid**](docs/LinksApi.md#listlinksbyfid)                                               | **Get** /v1/linksByFid                      | Get all links from a source FID                                                |
| _LinksApi_           | [**ListLinksByTargetFid**](docs/LinksApi.md#listlinksbytargetfid)                                   | **Get** /v1/linksByTargetFid                | Get all links to a target FID                                                  |
| _OnChainEventsApi_   | [**GetOnChainIdRegistrationByAddress**](docs/OnChainEventsApi.md#getonchainidregistrationbyaddress) | **Get** /v1/onChainIdRegistryEventByAddress | Get an on chain ID Registry Event for a given Address                          |
| _OnChainEventsApi_   | [**ListOnChainEventsByFid**](docs/OnChainEventsApi.md#listonchaineventsbyfid)                       | **Get** /v1/onChainEventsByFid              | Get a list of on-chain events provided by an FID                               |
| _OnChainEventsApi_   | [**ListOnChainSignersByFid**](docs/OnChainEventsApi.md#listonchainsignersbyfid)                     | **Get** /v1/onChainSignersByFid             | Get a list of signers provided by an FID                                       |
| _ReactionsApi_       | [**GetReactionById**](docs/ReactionsApi.md#getreactionbyid)                                         | **Get** /v1/reactionById                    | Get a reaction by its created FID and target Cast.                             |
| _ReactionsApi_       | [**ListReactionsByCast**](docs/ReactionsApi.md#listreactionsbycast)                                 | **Get** /v1/reactionsByCast                 | Get all reactions to a cast                                                    |
| _ReactionsApi_       | [**ListReactionsByFid**](docs/ReactionsApi.md#listreactionsbyfid)                                   | **Get** /v1/reactionsByFid                  | Get all reactions by an FID                                                    |
| _ReactionsApi_       | [**ListReactionsByTarget**](docs/ReactionsApi.md#listreactionsbytarget)                             | **Get** /v1/reactionsByTarget               | Get all reactions to a target URL                                              |
| _StorageApi_         | [**GetStorageLimitsByFid**](docs/StorageApi.md#getstoragelimitsbyfid)                               | **Get** /v1/storageLimitsByFid              | Get an FID&#x27;s storage limits.                                              |
| _SubmitMessageApi_   | [**SubmitMessage**](docs/SubmitMessageApi.md#submitmessage)                                         | **Post** /v1/submitMessage                  | Submit a signed protobuf-serialized message to the Hub                         |
| _UserDataApi_        | [**GetUserDataByFid**](docs/UserDataApi.md#getuserdatabyfid)                                        | **Get** /v1/userDataByFid                   | Get UserData for a FID.                                                        |
| _UsernamesApi_       | [**GetUsernameProof**](docs/UsernamesApi.md#getusernameproof)                                       | **Get** /v1/userNameProofByName             | Get an proof for a username by the Farcaster username                          |
| _UsernamesApi_       | [**ListUsernameProofsByFid**](docs/UsernamesApi.md#listusernameproofsbyfid)                         | **Get** /v1/userNameProofsByFid             | Get a list of proofs provided by an FID                                        |
| _ValidateMessageApi_ | [**ValidateMessage**](docs/ValidateMessageApi.md#validatemessage)                                   | **Post** /v1/validateMessage                | Validate a signed protobuf-serialized message with the Hub                     |
| _VerificationsApi_   | [**ListVerificationsByFid**](docs/VerificationsApi.md#listverificationsbyfid)                       | **Get** /v1/verificationsByFid              | Get a list of verifications provided by an FID                                 |

## Documentation For Models

- [CastAdd](docs/CastAdd.md)
- [CastAddBody](docs/CastAddBody.md)
- [CastEmbed](docs/CastEmbed.md)
- [CastId](docs/CastId.md)
- [CastRemove](docs/CastRemove.md)
- [CastRemoveBody](docs/CastRemoveBody.md)
- [DbStats](docs/DbStats.md)
- [Embed](docs/Embed.md)
- [ErrorResponse](docs/ErrorResponse.md)
- [ErrorResponseMetadata](docs/ErrorResponseMetadata.md)
- [FarcasterNetwork](docs/FarcasterNetwork.md)
- [FidsResponse](docs/FidsResponse.md)
- [FrameActionBody](docs/FrameActionBody.md)
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
- [InlineResponse200](docs/InlineResponse200.md)
- [InlineResponse2001](docs/InlineResponse2001.md)
- [InlineResponse2002](docs/InlineResponse2002.md)
- [InlineResponse2003](docs/InlineResponse2003.md)
- [InlineResponse2004](docs/InlineResponse2004.md)
- [InlineResponse2005](docs/InlineResponse2005.md)
- [InlineResponse2006](docs/InlineResponse2006.md)
- [InlineResponse2007](docs/InlineResponse2007.md)
- [LinkAdd](docs/LinkAdd.md)
- [LinkBody](docs/LinkBody.md)
- [LinkRemove](docs/LinkRemove.md)
- [LinkType](docs/LinkType.md)
- [MergeMessageBody](docs/MergeMessageBody.md)
- [MergeOnChainEventBody](docs/MergeOnChainEventBody.md)
- [MergeUserNameProofBody](docs/MergeUserNameProofBody.md)
- [Message](docs/Message.md)
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
- [UserDataBody](docs/UserDataBody.md)
- [UserDataType](docs/UserDataType.md)
- [UserNameProof](docs/UserNameProof.md)
- [UserNameType](docs/UserNameType.md)
- [UsernameProofsResponse](docs/UsernameProofsResponse.md)
- [ValidateMessageResponse](docs/ValidateMessageResponse.md)
- [Verification](docs/Verification.md)
- [VerificationAddEthAddressBody](docs/VerificationAddEthAddressBody.md)
- [VerificationRemove](docs/VerificationRemove.md)
- [VerificationRemoveBody](docs/VerificationRemoveBody.md)

## Documentation For Authorization

## usernamePassword

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author
