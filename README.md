# go_neynar_sdk

go_neynar_sdk is a Go client library for accessing the [Neynar API's](https://docs.neynar.com/).

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
go get github.com/atlasmoth/go_neynar_sdk@vX.Y.Z
```

where X.Y.Z is the version you need.

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.neynar.com/v2*

| Class              | Method                                                                                                         | HTTP request                                            | Description                                                        |
| ------------------ | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------------------ |
| _CastAPI_          | [**Cast**](docs/CastAPI.md#cast)                                                                               | **Get** /farcaster/cast                                 | Retrieve cast for a given hash or Warpcast URL                     |
| _CastAPI_          | [**CastConversation**](docs/CastAPI.md#castconversation)                                                       | **Get** /farcaster/cast/conversation                    | Retrieve the conversation for a given cast                         |
| _CastAPI_          | [**Casts**](docs/CastAPI.md#casts)                                                                             | **Get** /farcaster/casts                                | Gets information about an array of casts                           |
| _CastAPI_          | [**DeleteCast**](docs/CastAPI.md#deletecast)                                                                   | **Delete** /farcaster/cast                              | Delete a cast                                                      |
| _CastAPI_          | [**PostCast**](docs/CastAPI.md#postcast)                                                                       | **Post** /farcaster/cast                                | Posts a cast                                                       |
| _ChannelAPI_       | [**ActiveChannels**](docs/ChannelAPI.md#activechannels)                                                        | **Get** /farcaster/channel/user                         | Get channels that a user is active in                              |
| _ChannelAPI_       | [**ChannelDetails**](docs/ChannelAPI.md#channeldetails)                                                        | **Get** /farcaster/channel                              | Retrieve channel details by id or parent_url                       |
| _ChannelAPI_       | [**ChannelDetailsBulk**](docs/ChannelAPI.md#channeldetailsbulk)                                                | **Get** /farcaster/channel/bulk                         | (Bulk) Retrieve channels by id or parent_url                       |
| _ChannelAPI_       | [**ChannelFollowers**](docs/ChannelAPI.md#channelfollowers)                                                    | **Get** /farcaster/channel/followers                    | Retrieve followers for a given channel                             |
| _ChannelAPI_       | [**ChannelUsers**](docs/ChannelAPI.md#channelusers)                                                            | **Get** /farcaster/channel/users                        | Retrieve users who are active in a channel                         |
| _ChannelAPI_       | [**ListAllChannels**](docs/ChannelAPI.md#listallchannels)                                                      | **Get** /farcaster/channel/list                         | Retrieve all channels with their details                           |
| _ChannelAPI_       | [**SearchChannels**](docs/ChannelAPI.md#searchchannels)                                                        | **Get** /farcaster/channel/search                       | Search for channels based on id or name                            |
| _ChannelAPI_       | [**TrendingChannels**](docs/ChannelAPI.md#trendingchannels)                                                    | **Get** /farcaster/channel/trending                     | Retrieve trending channels based on activity                       |
| _ChannelAPI_       | [**UserChannels**](docs/ChannelAPI.md#userchannels)                                                            | **Get** /farcaster/user/channels                        | Retrieve all channels that a given fid follows                     |
| _FeedAPI_          | [**Feed**](docs/FeedAPI.md#feed)                                                                               | **Get** /farcaster/feed                                 | Retrieve casts based on filters                                    |
| _FeedAPI_          | [**FeedChannels**](docs/FeedAPI.md#feedchannels)                                                               | **Get** /farcaster/feed/channels                        | Retrieve feed based on channel ids                                 |
| _FeedAPI_          | [**FeedFollowing**](docs/FeedAPI.md#feedfollowing)                                                             | **Get** /farcaster/feed/following                       | Retrieve feed based on who a user is following                     |
| _FeedAPI_          | [**FeedForYou**](docs/FeedAPI.md#feedforyou)                                                                   | **Get** /farcaster/feed/for_you                         | Retrieve a personalized For You feed for a user                    |
| _FeedAPI_          | [**FeedFrames**](docs/FeedAPI.md#feedframes)                                                                   | **Get** /farcaster/feed/frames                          | Retrieve feed of casts with Frames, reverse chronological order    |
| _FeedAPI_          | [**FeedTrending**](docs/FeedAPI.md#feedtrending)                                                               | **Get** /farcaster/feed/trending                        | Retrieve trending casts                                            |
| _FeedAPI_          | [**FeedUserPopular**](docs/FeedAPI.md#feeduserpopular)                                                         | **Get** /farcaster/feed/user/popular                    | Retrieve 10 most popular casts for a user                          |
| _FeedAPI_          | [**FeedUserRepliesRecasts**](docs/FeedAPI.md#feeduserrepliesrecasts)                                           | **Get** /farcaster/feed/user/replies_and_recasts        | Retrieve recent replies and recasts for a user                     |
| _FnameAPI_         | [**FnameAvailability**](docs/FnameAPI.md#fnameavailability)                                                    | **Get** /farcaster/fname/availability                   | Check if a given fname is available                                |
| _FollowsAPI_       | [**FollowersV2**](docs/FollowsAPI.md#followersv2)                                                              | **Get** /farcaster/followers                            | Retrieve followers for a given user                                |
| _FollowsAPI_       | [**FollowingV2**](docs/FollowsAPI.md#followingv2)                                                              | **Get** /farcaster/following                            | Retrieve a list of users followed by a user                        |
| _FollowsAPI_       | [**RelevantFollowers**](docs/FollowsAPI.md#relevantfollowers)                                                  | **Get** /farcaster/followers/relevant                   | Retrieve relevant followers for a given user                       |
| _FrameAPI_         | [**DeleteNeynarFrame**](docs/FrameAPI.md#deleteneynarframe)                                                    | **Delete** /farcaster/frame                             | Delete a frame                                                     |
| _FrameAPI_         | [**FetchNeynarFrames**](docs/FrameAPI.md#fetchneynarframes)                                                    | **Get** /farcaster/frame/list                           | Retrieve a list of frames                                          |
| _FrameAPI_         | [**LookupNeynarFrame**](docs/FrameAPI.md#lookupneynarframe)                                                    | **Get** /farcaster/frame                                | Retrieve a frame by UUID or URL                                    |
| _FrameAPI_         | [**PostFrameAction**](docs/FrameAPI.md#postframeaction)                                                        | **Post** /farcaster/frame/action                        | Posts a frame action, cast action or a cast composer action        |
| _FrameAPI_         | [**PostFrameDeveloperManagedAction**](docs/FrameAPI.md#postframedevelopermanagedaction)                        | **Post** /farcaster/frame/developer_managed/action      | Posts a frame signature packet                                     |
| _FrameAPI_         | [**PublishNeynarFrame**](docs/FrameAPI.md#publishneynarframe)                                                  | **Post** /farcaster/frame                               | Create a new frame                                                 |
| _FrameAPI_         | [**UpdateNeynarFrame**](docs/FrameAPI.md#updateneynarframe)                                                    | **Put** /farcaster/frame                                | Update an existing frame                                           |
| _FrameAPI_         | [**ValidateFrame**](docs/FrameAPI.md#validateframe)                                                            | **Post** /farcaster/frame/validate                      | Validates a frame action against Farcaster Hub                     |
| _FrameAPI_         | [**ValidateFrameAnalytics**](docs/FrameAPI.md#validateframeanalytics)                                          | **Get** /farcaster/frame/validate/analytics             | Retrieve analytics for the frame                                   |
| _FrameAPI_         | [**ValidateFrameList**](docs/FrameAPI.md#validateframelist)                                                    | **Get** /farcaster/frame/validate/list                  | Retrieve a list of all the frames validated by a user              |
| _MuteAPI_          | [**AddMute**](docs/MuteAPI.md#addmute)                                                                         | **Post** /farcaster/mute                                | Adds a mute for a fid                                              |
| _MuteAPI_          | [**DeleteMute**](docs/MuteAPI.md#deletemute)                                                                   | **Delete** /farcaster/mute                              | Deletes a mute for a fid                                           |
| _MuteAPI_          | [**MuteList**](docs/MuteAPI.md#mutelist)                                                                       | **Get** /farcaster/mute/list                            | Get fids that a user has muted                                     |
| _NotificationsAPI_ | [**Notifications**](docs/NotificationsAPI.md#notifications)                                                    | **Get** /farcaster/notifications                        | Retrieve notifications for a given user                            |
| _NotificationsAPI_ | [**NotificationsChannel**](docs/NotificationsAPI.md#notificationschannel)                                      | **Get** /farcaster/notifications/channel                | Retrieve notifications for a user in given channels                |
| _NotificationsAPI_ | [**NotificationsParentUrl**](docs/NotificationsAPI.md#notificationsparenturl)                                  | **Get** /farcaster/notifications/parent_url             | Retrieve notifications for a user in given parent_urls             |
| _ReactionAPI_      | [**DeleteReaction**](docs/ReactionAPI.md#deletereaction)                                                       | **Delete** /farcaster/reaction                          | Delete a reaction                                                  |
| _ReactionAPI_      | [**PostReaction**](docs/ReactionAPI.md#postreaction)                                                           | **Post** /farcaster/reaction                            | Posts a reaction                                                   |
| _ReactionAPI_      | [**ReactionsCast**](docs/ReactionAPI.md#reactionscast)                                                         | **Get** /farcaster/reactions/cast                       | Fetches reactions for a given cast                                 |
| _ReactionAPI_      | [**ReactionsUser**](docs/ReactionAPI.md#reactionsuser)                                                         | **Get** /farcaster/reactions/user                       | Fetches reactions for a given user                                 |
| _SignerAPI_        | [**CreateSigner**](docs/SignerAPI.md#createsigner)                                                             | **Post** /farcaster/signer                              | Creates a signer and returns the signer status                     |
| _SignerAPI_        | [**DeveloperManagedSigner**](docs/SignerAPI.md#developermanagedsigner)                                         | **Get** /farcaster/signer/developer_managed             | Fetches the status of a signer by public key                       |
| _SignerAPI_        | [**FetchAuthorizationUrl**](docs/SignerAPI.md#fetchauthorizationurl)                                           | **Get** /farcaster/login/authorize                      | Fetch authorization url                                            |
| _SignerAPI_        | [**PublishMessage**](docs/SignerAPI.md#publishmessage)                                                         | **Post** /farcaster/message                             | Publish a message to farcaster                                     |
| _SignerAPI_        | [**RegisterSignedKey**](docs/SignerAPI.md#registersignedkey)                                                   | **Post** /farcaster/signer/signed_key                   | Register Signed Key                                                |
| _SignerAPI_        | [**RegisterSignedKeyForDeveloperManagedSigner**](docs/SignerAPI.md#registersignedkeyfordevelopermanagedsigner) | **Post** /farcaster/signer/developer_managed/signed_key | Registers Signed Key                                               |
| _SignerAPI_        | [**Signer**](docs/SignerAPI.md#signer)                                                                         | **Get** /farcaster/signer                               | Fetches the status of a signer                                     |
| _StorageAPI_       | [**BuyStorage**](docs/StorageAPI.md#buystorage)                                                                | **Post** /farcaster/storage/buy                         | Buy storage for an fid                                             |
| _StorageAPI_       | [**StorageAllocations**](docs/StorageAPI.md#storageallocations)                                                | **Get** /farcaster/storage/allocations                  | Fetches storage allocations for a given user                       |
| _StorageAPI_       | [**StorageUsage**](docs/StorageAPI.md#storageusage)                                                            | **Get** /farcaster/storage/usage                        | Fetches storage usage for a given user                             |
| _SubscribersAPI_   | [**SubscribedTo**](docs/SubscribersAPI.md#subscribedto)                                                        | **Get** /farcaster/user/subscribed_to                   | Fetch what a given fid is subscribed to                            |
| _SubscribersAPI_   | [**Subscribers**](docs/SubscribersAPI.md#subscribers)                                                          | **Get** /farcaster/user/subscribers                     | Fetch subscribers for a given fid                                  |
| _SubscribersAPI_   | [**SubscriptionsCreated**](docs/SubscribersAPI.md#subscriptionscreated)                                        | **Get** /farcaster/user/subscriptions_created           | Fetch created subscriptions for a given fid                        |
| _UserAPI_          | [**ActiveUsers**](docs/UserAPI.md#activeusers)                                                                 | **Get** /farcaster/user/active                          | Fetch active users                                                 |
| _UserAPI_          | [**FarcasterUserVerificationDelete**](docs/UserAPI.md#farcasteruserverificationdelete)                         | **Delete** /farcaster/user/verification                 | Removes verification for an eth address for the user               |
| _UserAPI_          | [**FarcasterUserVerificationPost**](docs/UserAPI.md#farcasteruserverificationpost)                             | **Post** /farcaster/user/verification                   | Adds verification for an ethereum address or contract for the user |
| _UserAPI_          | [**FollowUser**](docs/UserAPI.md#followuser)                                                                   | **Post** /farcaster/user/follow                         | Follow a user                                                      |
| _UserAPI_          | [**GetFreshFid**](docs/UserAPI.md#getfreshfid)                                                                 | **Get** /farcaster/user/fid                             | Fetches fid to assign it new user                                  |
| _UserAPI_          | [**LookupUserByCustodyAddress**](docs/UserAPI.md#lookupuserbycustodyaddress)                                   | **Get** /farcaster/user/custody-address                 | Lookup a user by custody-address                                   |
| _UserAPI_          | [**PowerUsers**](docs/UserAPI.md#powerusers)                                                                   | **Get** /farcaster/user/power                           | Fetch power users                                                  |
| _UserAPI_          | [**RegisterUser**](docs/UserAPI.md#registeruser)                                                               | **Post** /farcaster/user                                | Register account on farcaster                                      |
| _UserAPI_          | [**UnfollowUser**](docs/UserAPI.md#unfollowuser)                                                               | **Delete** /farcaster/user/follow                       | Unfollow a user                                                    |
| _UserAPI_          | [**UpdateUser**](docs/UserAPI.md#updateuser)                                                                   | **Patch** /farcaster/user                               | Update user profile                                                |
| _UserAPI_          | [**UserBulk**](docs/UserAPI.md#userbulk)                                                                       | **Get** /farcaster/user/bulk                            | Fetch users based on FIDs                                          |
| _UserAPI_          | [**UserBulkByAddress**](docs/UserAPI.md#userbulkbyaddress)                                                     | **Get** /farcaster/user/bulk-by-address                 | Fetches users based on Eth or Sol addresses                        |
| _UserAPI_          | [**UserSearch**](docs/UserAPI.md#usersearch)                                                                   | **Get** /farcaster/user/search                          | Search for Usernames                                               |
| _WebhookAPI_       | [**DeleteWebhook**](docs/WebhookAPI.md#deletewebhook)                                                          | **Delete** /farcaster/webhook                           | Delete a webhook                                                   |
| _WebhookAPI_       | [**FetchWebhooks**](docs/WebhookAPI.md#fetchwebhooks)                                                          | **Get** /farcaster/webhook/list                         | Fetch a list of webhooks associated to a user                      |
| _WebhookAPI_       | [**LookupWebhook**](docs/WebhookAPI.md#lookupwebhook)                                                          | **Get** /farcaster/webhook                              | Fetch a webhook                                                    |
| _WebhookAPI_       | [**PublishWebhook**](docs/WebhookAPI.md#publishwebhook)                                                        | **Post** /farcaster/webhook                             | Create a webhook                                                   |
| _WebhookAPI_       | [**UpdateWebhook**](docs/WebhookAPI.md#updatewebhook)                                                          | **Put** /farcaster/webhook                              | Update a webhook                                                   |
| _WebhookAPI_       | [**UpdateWebhookActiveStatus**](docs/WebhookAPI.md#updatewebhookactivestatus)                                  | **Patch** /farcaster/webhook                            | Update webhook active status                                       |

## Documentation For Models

- [ActiveStatus](docs/ActiveStatus.md)
- [AddVerificationReqBody](docs/AddVerificationReqBody.md)
- [AuthorizationUrlResponse](docs/AuthorizationUrlResponse.md)
- [AuthorizationUrlResponseType](docs/AuthorizationUrlResponseType.md)
- [BulkCastsResponse](docs/BulkCastsResponse.md)
- [BulkFollowResponse](docs/BulkFollowResponse.md)
- [BulkUsersResponse](docs/BulkUsersResponse.md)
- [BuyStorageReqBody](docs/BuyStorageReqBody.md)
- [Cast](docs/Cast.md)
- [CastDehydrated](docs/CastDehydrated.md)
- [CastId](docs/CastId.md)
- [CastNotificationType](docs/CastNotificationType.md)
- [CastParamType](docs/CastParamType.md)
- [CastParentAuthor](docs/CastParentAuthor.md)
- [CastResponse](docs/CastResponse.md)
- [CastViewerContext](docs/CastViewerContext.md)
- [CastWithInteractions](docs/CastWithInteractions.md)
- [CastWithInteractionsAndConversations](docs/CastWithInteractionsAndConversations.md)
- [CastWithInteractionsReactions](docs/CastWithInteractionsReactions.md)
- [CastWithInteractionsReplies](docs/CastWithInteractionsReplies.md)
- [CastsResponse](docs/CastsResponse.md)
- [CastsResponseResult](docs/CastsResponseResult.md)
- [Channel](docs/Channel.md)
- [ChannelActivity](docs/ChannelActivity.md)
- [ChannelListResponse](docs/ChannelListResponse.md)
- [ChannelOrDehydratedChannel](docs/ChannelOrDehydratedChannel.md)
- [ChannelResponse](docs/ChannelResponse.md)
- [ChannelResponseBulk](docs/ChannelResponseBulk.md)
- [ChannelSearchResponse](docs/ChannelSearchResponse.md)
- [ChannelType](docs/ChannelType.md)
- [ChannelViewerContext](docs/ChannelViewerContext.md)
- [ConflictErrorRes](docs/ConflictErrorRes.md)
- [Conversation](docs/Conversation.md)
- [ConversationConversation](docs/ConversationConversation.md)
- [DehydratedChannel](docs/DehydratedChannel.md)
- [DehydratedFollower](docs/DehydratedFollower.md)
- [DeleteCastReqBody](docs/DeleteCastReqBody.md)
- [DeleteFrameResponse](docs/DeleteFrameResponse.md)
- [DeleteNeynarFrameRequest](docs/DeleteNeynarFrameRequest.md)
- [DeveloperManagedSigner](docs/DeveloperManagedSigner.md)
- [EmbedCastId](docs/EmbedCastId.md)
- [EmbedUrl](docs/EmbedUrl.md)
- [EmbedUrlMetadata](docs/EmbedUrlMetadata.md)
- [EmbeddedCast](docs/EmbeddedCast.md)
- [ErrorRes](docs/ErrorRes.md)
- [FeedResponse](docs/FeedResponse.md)
- [FeedType](docs/FeedType.md)
- [FilterType](docs/FilterType.md)
- [FnameAvailabilityResponse](docs/FnameAvailabilityResponse.md)
- [Follow](docs/Follow.md)
- [FollowReqBody](docs/FollowReqBody.md)
- [FollowResponse](docs/FollowResponse.md)
- [FollowSortType](docs/FollowSortType.md)
- [ForYouProvider](docs/ForYouProvider.md)
- [Frame](docs/Frame.md)
- [FrameAction](docs/FrameAction.md)
- [FrameActionButton](docs/FrameActionButton.md)
- [FrameActionReqBody](docs/FrameActionReqBody.md)
- [FrameButtonActionType](docs/FrameButtonActionType.md)
- [FrameDeveloperManagedActionReqBody](docs/FrameDeveloperManagedActionReqBody.md)
- [FrameInput](docs/FrameInput.md)
- [FrameSignaturePacket](docs/FrameSignaturePacket.md)
- [FrameSignaturePacketTrustedData](docs/FrameSignaturePacketTrustedData.md)
- [FrameSignaturePacketUntrustedData](docs/FrameSignaturePacketUntrustedData.md)
- [FrameState](docs/FrameState.md)
- [FrameTransaction](docs/FrameTransaction.md)
- [FrameType](docs/FrameType.md)
- [FrameValidateAnalyticsInputText](docs/FrameValidateAnalyticsInputText.md)
- [FrameValidateAnalyticsInputTextInputTextsInner](docs/FrameValidateAnalyticsInputTextInputTextsInner.md)
- [FrameValidateAnalyticsInteractionsPerCast](docs/FrameValidateAnalyticsInteractionsPerCast.md)
- [FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner](docs/FrameValidateAnalyticsInteractionsPerCastInteractionsPerCastInner.md)
- [FrameValidateAnalyticsInteractors](docs/FrameValidateAnalyticsInteractors.md)
- [FrameValidateAnalyticsInteractorsInteractorsInner](docs/FrameValidateAnalyticsInteractorsInteractorsInner.md)
- [FrameValidateAnalyticsResponse](docs/FrameValidateAnalyticsResponse.md)
- [FrameValidateAnalyticsTotalInteractors](docs/FrameValidateAnalyticsTotalInteractors.md)
- [FrameValidateListResponse](docs/FrameValidateListResponse.md)
- [GetCastsReqBody](docs/GetCastsReqBody.md)
- [HydratedFollower](docs/HydratedFollower.md)
- [IndividualHashObj](docs/IndividualHashObj.md)
- [MuteList](docs/MuteList.md)
- [MuteListResponse](docs/MuteListResponse.md)
- [MuteReqBody](docs/MuteReqBody.md)
- [MuteResponse](docs/MuteResponse.md)
- [NextCursor](docs/NextCursor.md)
- [NeynarFrame](docs/NeynarFrame.md)
- [NeynarFrameCreationRequest](docs/NeynarFrameCreationRequest.md)
- [NeynarFramePage](docs/NeynarFramePage.md)
- [NeynarFrameUpdateRequest](docs/NeynarFrameUpdateRequest.md)
- [NeynarNextFramePage](docs/NeynarNextFramePage.md)
- [NeynarNextFramePageMintUrl](docs/NeynarNextFramePageMintUrl.md)
- [NeynarNextFramePageRedirect](docs/NeynarNextFramePageRedirect.md)
- [NeynarPageButton](docs/NeynarPageButton.md)
- [NeynarPageButtonNextPage](docs/NeynarPageButtonNextPage.md)
- [NeynarPageImage](docs/NeynarPageImage.md)
- [NeynarPageInput](docs/NeynarPageInput.md)
- [NeynarPageInputText](docs/NeynarPageInputText.md)
- [Notification](docs/Notification.md)
- [NotificationsResponse](docs/NotificationsResponse.md)
- [OperationResponse](docs/OperationResponse.md)
- [PostCastReqBody](docs/PostCastReqBody.md)
- [PostCastResponse](docs/PostCastResponse.md)
- [PostCastResponseCast](docs/PostCastResponseCast.md)
- [PostCastResponseCastAuthor](docs/PostCastResponseCastAuthor.md)
- [ProfileUrl](docs/ProfileUrl.md)
- [ProfileUrlPfp](docs/ProfileUrlPfp.md)
- [ReactionForCast](docs/ReactionForCast.md)
- [ReactionLike](docs/ReactionLike.md)
- [ReactionRecast](docs/ReactionRecast.md)
- [ReactionReqBody](docs/ReactionReqBody.md)
- [ReactionType](docs/ReactionType.md)
- [ReactionWithCastInfo](docs/ReactionWithCastInfo.md)
- [ReactionWithUserInfo](docs/ReactionWithUserInfo.md)
- [ReactionsCastResponse](docs/ReactionsCastResponse.md)
- [ReactionsResponse](docs/ReactionsResponse.md)
- [ReactionsType](docs/ReactionsType.md)
- [RegisterDeveloperManagedSignedKeyReqBody](docs/RegisterDeveloperManagedSignedKeyReqBody.md)
- [RegisterSignerKeyReqBody](docs/RegisterSignerKeyReqBody.md)
- [RegisterUserReqBody](docs/RegisterUserReqBody.md)
- [RegisterUserResponse](docs/RegisterUserResponse.md)
- [RelevantFollowersResponse](docs/RelevantFollowersResponse.md)
- [RemoveVerificationReqBody](docs/RemoveVerificationReqBody.md)
- [SearchedUser](docs/SearchedUser.md)
- [Signer](docs/Signer.md)
- [StorageAllocation](docs/StorageAllocation.md)
- [StorageAllocationsResponse](docs/StorageAllocationsResponse.md)
- [StorageObject](docs/StorageObject.md)
- [StorageUsageResponse](docs/StorageUsageResponse.md)
- [SubscribedTo](docs/SubscribedTo.md)
- [SubscribedToObject](docs/SubscribedToObject.md)
- [SubscribedToResponse](docs/SubscribedToResponse.md)
- [Subscriber](docs/Subscriber.md)
- [SubscribersResponse](docs/SubscribersResponse.md)
- [Subscription](docs/Subscription.md)
- [SubscriptionMetadata](docs/SubscriptionMetadata.md)
- [SubscriptionPrice](docs/SubscriptionPrice.md)
- [SubscriptionProvider](docs/SubscriptionProvider.md)
- [SubscriptionProviders](docs/SubscriptionProviders.md)
- [SubscriptionTier](docs/SubscriptionTier.md)
- [SubscriptionTierPrice](docs/SubscriptionTierPrice.md)
- [SubscriptionToken](docs/SubscriptionToken.md)
- [Subscriptions](docs/Subscriptions.md)
- [SubscriptionsResponse](docs/SubscriptionsResponse.md)
- [TrendingChannelResponse](docs/TrendingChannelResponse.md)
- [UpdateUserReqBody](docs/UpdateUserReqBody.md)
- [User](docs/User.md)
- [UserDehydrated](docs/UserDehydrated.md)
- [UserFIDResponse](docs/UserFIDResponse.md)
- [UserProfile](docs/UserProfile.md)
- [UserProfileBio](docs/UserProfileBio.md)
- [UserResponse](docs/UserResponse.md)
- [UserSearchResponse](docs/UserSearchResponse.md)
- [UserSearchResponseResult](docs/UserSearchResponseResult.md)
- [UserVerifiedAddresses](docs/UserVerifiedAddresses.md)
- [UserViewerContext](docs/UserViewerContext.md)
- [UsersActiveChannelsResponse](docs/UsersActiveChannelsResponse.md)
- [UsersResponse](docs/UsersResponse.md)
- [ValidateFrameActionResponse](docs/ValidateFrameActionResponse.md)
- [ValidateFrameAnalyticsType](docs/ValidateFrameAnalyticsType.md)
- [ValidateFrameRequest](docs/ValidateFrameRequest.md)
- [ValidatedFrameAction](docs/ValidatedFrameAction.md)
- [ValidatedFrameActionSigner](docs/ValidatedFrameActionSigner.md)
- [ValidatedFrameActionTappedButton](docs/ValidatedFrameActionTappedButton.md)
- [VerificationChainId](docs/VerificationChainId.md)
- [VerificationType](docs/VerificationType.md)
- [Webhook](docs/Webhook.md)
- [WebhookDeleteReqBody](docs/WebhookDeleteReqBody.md)
- [WebhookListResponse](docs/WebhookListResponse.md)
- [WebhookPatchReqBody](docs/WebhookPatchReqBody.md)
- [WebhookPostReqBody](docs/WebhookPostReqBody.md)
- [WebhookPutReqBody](docs/WebhookPutReqBody.md)
- [WebhookResponse](docs/WebhookResponse.md)
- [WebhookSecret](docs/WebhookSecret.md)
- [WebhookSubscription](docs/WebhookSubscription.md)
- [WebhookSubscriptionFilters](docs/WebhookSubscriptionFilters.md)
- [WebhookSubscriptionFiltersCastCreated](docs/WebhookSubscriptionFiltersCastCreated.md)
- [WebhookSubscriptionFiltersFollow](docs/WebhookSubscriptionFiltersFollow.md)
- [WebhookSubscriptionFiltersReaction](docs/WebhookSubscriptionFiltersReaction.md)
- [WebhookSubscriptionFiltersUserUpdated](docs/WebhookSubscriptionFiltersUserUpdated.md)
- [ZodError](docs/ZodError.md)
- [ZodErrorErrorsInner](docs/ZodErrorErrorsInner.md)

## Documentation For Authorization

Endpoints do not require authorization.

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

## Author
