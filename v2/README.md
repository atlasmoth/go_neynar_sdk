## Documentation for API Endpoints

All URIs are relative to *https://api.neynar.com/v2*

| Class              | Method                                                                                                         | HTTP request                                            | Description                                                        |
| ------------------ | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------------------ |
| _CastApi_          | [**Cast**](docs/CastApi.md#cast)                                                                               | **Get** /farcaster/cast                                 | Retrieve cast for a given hash or Warpcast URL                     |
| _CastApi_          | [**CastConversation**](docs/CastApi.md#castconversation)                                                       | **Get** /farcaster/cast/conversation                    | Retrieve the conversation for a given cast                         |
| _CastApi_          | [**CastSearch**](docs/CastApi.md#castsearch)                                                                   | **Get** /farcaster/cast/search                          | Search for casts                                                   |
| _CastApi_          | [**Casts**](docs/CastApi.md#casts)                                                                             | **Get** /farcaster/casts                                | Gets information about an array of casts                           |
| _CastApi_          | [**ComposerList**](docs/CastApi.md#composerlist)                                                               | **Get** /farcaster/cast/composer_actions/list           | Fetches all composer actions on Warpcast                           |
| _CastApi_          | [**DeleteCast**](docs/CastApi.md#deletecast)                                                                   | **Delete** /farcaster/cast                              | Delete a cast                                                      |
| _CastApi_          | [**PostCast**](docs/CastApi.md#postcast)                                                                       | **Post** /farcaster/cast                                | Posts a cast                                                       |
| _ChannelApi_       | [**ActiveChannels**](docs/ChannelApi.md#activechannels)                                                        | **Get** /farcaster/channel/user                         | Get channels that a user is active in                              |
| _ChannelApi_       | [**ChannelDetails**](docs/ChannelApi.md#channeldetails)                                                        | **Get** /farcaster/channel                              | Retrieve channel details by id or parent_url                       |
| _ChannelApi_       | [**ChannelDetailsBulk**](docs/ChannelApi.md#channeldetailsbulk)                                                | **Get** /farcaster/channel/bulk                         | (Bulk) Retrieve channels by id or parent_url                       |
| _ChannelApi_       | [**ChannelFollowers**](docs/ChannelApi.md#channelfollowers)                                                    | **Get** /farcaster/channel/followers                    | Retrieve followers for a given channel                             |
| _ChannelApi_       | [**ChannelUsers**](docs/ChannelApi.md#channelusers)                                                            | **Get** /farcaster/channel/users                        | Retrieve users who are active in a channel                         |
| _ChannelApi_       | [**ListAllChannels**](docs/ChannelApi.md#listallchannels)                                                      | **Get** /farcaster/channel/list                         | Retrieve all channels with their details                           |
| _ChannelApi_       | [**SearchChannels**](docs/ChannelApi.md#searchchannels)                                                        | **Get** /farcaster/channel/search                       | Search for channels based on id or name                            |
| _ChannelApi_       | [**TrendingChannels**](docs/ChannelApi.md#trendingchannels)                                                    | **Get** /farcaster/channel/trending                     | Retrieve trending channels based on activity                       |
| _ChannelApi_       | [**UserChannels**](docs/ChannelApi.md#userchannels)                                                            | **Get** /farcaster/user/channels                        | Retrieve all channels that a given fid follows                     |
| _FeedApi_          | [**Feed**](docs/FeedApi.md#feed)                                                                               | **Get** /farcaster/feed                                 | Retrieve casts based on filters                                    |
| _FeedApi_          | [**FeedChannels**](docs/FeedApi.md#feedchannels)                                                               | **Get** /farcaster/feed/channels                        | Retrieve feed based on channel ids                                 |
| _FeedApi_          | [**FeedFollowing**](docs/FeedApi.md#feedfollowing)                                                             | **Get** /farcaster/feed/following                       | Retrieve feed based on who a user is following                     |
| _FeedApi_          | [**FeedForYou**](docs/FeedApi.md#feedforyou)                                                                   | **Get** /farcaster/feed/for_you                         | Retrieve a personalized For You feed for a user                    |
| _FeedApi_          | [**FeedFrames**](docs/FeedApi.md#feedframes)                                                                   | **Get** /farcaster/feed/frames                          | Retrieve feed of casts with Frames, reverse chronological order    |
| _FeedApi_          | [**FeedParentUrls**](docs/FeedApi.md#feedparenturls)                                                           | **Get** /farcaster/feed/parent_urls                     | Retrieve feed based on parent urls                                 |
| _FeedApi_          | [**FeedTrending**](docs/FeedApi.md#feedtrending)                                                               | **Get** /farcaster/feed/trending                        | Retrieve trending casts                                            |
| _FeedApi_          | [**FeedUserCasts**](docs/FeedApi.md#feedusercasts)                                                             | **Get** /farcaster/feed/user/casts                      | Retrieve casts for a user                                          |
| _FeedApi_          | [**FeedUserPopular**](docs/FeedApi.md#feeduserpopular)                                                         | **Get** /farcaster/feed/user/popular                    | Retrieve 10 most popular casts for a user                          |
| _FeedApi_          | [**FeedUserRepliesRecasts**](docs/FeedApi.md#feeduserrepliesrecasts)                                           | **Get** /farcaster/feed/user/replies_and_recasts        | Retrieve recent replies and recasts for a user                     |
| _FnameApi_         | [**FnameAvailability**](docs/FnameApi.md#fnameavailability)                                                    | **Get** /farcaster/fname/availability                   | Check if a given fname is available                                |
| _FollowsApi_       | [**FollowersV2**](docs/FollowsApi.md#followersv2)                                                              | **Get** /farcaster/followers                            | Retrieve followers for a given user                                |
| _FollowsApi_       | [**FollowingV2**](docs/FollowsApi.md#followingv2)                                                              | **Get** /farcaster/following                            | Retrieve a list of users followed by a user                        |
| _FollowsApi_       | [**RelevantFollowers**](docs/FollowsApi.md#relevantfollowers)                                                  | **Get** /farcaster/followers/relevant                   | Retrieve relevant followers for a given user                       |
| _FrameApi_         | [**DeleteNeynarFrame**](docs/FrameApi.md#deleteneynarframe)                                                    | **Delete** /farcaster/frame                             | Delete a frame                                                     |
| _FrameApi_         | [**FetchNeynarFrames**](docs/FrameApi.md#fetchneynarframes)                                                    | **Get** /farcaster/frame/list                           | Retrieve a list of frames                                          |
| _FrameApi_         | [**FrameFromUrl**](docs/FrameApi.md#framefromurl)                                                              | **Get** /farcaster/frame/crawl                          | Fetches the frame meta tags from the URL                           |
| _FrameApi_         | [**LookupNeynarFrame**](docs/FrameApi.md#lookupneynarframe)                                                    | **Get** /farcaster/frame                                | Retrieve a frame by UUID or URL                                    |
| _FrameApi_         | [**PostFrameAction**](docs/FrameApi.md#postframeaction)                                                        | **Post** /farcaster/frame/action                        | Posts a frame action, cast action or a cast composer action        |
| _FrameApi_         | [**PostFrameDeveloperManagedAction**](docs/FrameApi.md#postframedevelopermanagedaction)                        | **Post** /farcaster/frame/developer_managed/action      | Posts a frame signature packet                                     |
| _FrameApi_         | [**PublishNeynarFrame**](docs/FrameApi.md#publishneynarframe)                                                  | **Post** /farcaster/frame                               | Create a new frame                                                 |
| _FrameApi_         | [**UpdateNeynarFrame**](docs/FrameApi.md#updateneynarframe)                                                    | **Put** /farcaster/frame                                | Update an existing frame                                           |
| _FrameApi_         | [**ValidateFrame**](docs/FrameApi.md#validateframe)                                                            | **Post** /farcaster/frame/validate                      | Validates a frame action against Farcaster Hub                     |
| _FrameApi_         | [**ValidateFrameAnalytics**](docs/FrameApi.md#validateframeanalytics)                                          | **Get** /farcaster/frame/validate/analytics             | Retrieve analytics for the frame                                   |
| _FrameApi_         | [**ValidateFrameList**](docs/FrameApi.md#validateframelist)                                                    | **Get** /farcaster/frame/validate/list                  | Retrieve a list of all the frames validated by a user              |
| _MuteApi_          | [**AddMute**](docs/MuteApi.md#addmute)                                                                         | **Post** /farcaster/mute                                | Adds a mute for a fid                                              |
| _MuteApi_          | [**DeleteMute**](docs/MuteApi.md#deletemute)                                                                   | **Delete** /farcaster/mute                              | Deletes a mute for a fid                                           |
| _MuteApi_          | [**MuteList**](docs/MuteApi.md#mutelist)                                                                       | **Get** /farcaster/mute/list                            | Get fids that a user has muted                                     |
| _NotificationsApi_ | [**MarkNotificationsAsSeen**](docs/NotificationsApi.md#marknotificationsasseen)                                | **Post** /farcaster/notifications/seen                  | Mark notifications as seen                                         |
| _NotificationsApi_ | [**Notifications**](docs/NotificationsApi.md#notifications)                                                    | **Get** /farcaster/notifications                        | Retrieve notifications for a given user                            |
| _NotificationsApi_ | [**NotificationsChannel**](docs/NotificationsApi.md#notificationschannel)                                      | **Get** /farcaster/notifications/channel                | Retrieve notifications for a user in given channels                |
| _NotificationsApi_ | [**NotificationsParentUrl**](docs/NotificationsApi.md#notificationsparenturl)                                  | **Get** /farcaster/notifications/parent_url             | Retrieve notifications for a user in given parent_urls             |
| _ReactionApi_      | [**DeleteReaction**](docs/ReactionApi.md#deletereaction)                                                       | **Delete** /farcaster/reaction                          | Delete a reaction                                                  |
| _ReactionApi_      | [**PostReaction**](docs/ReactionApi.md#postreaction)                                                           | **Post** /farcaster/reaction                            | Posts a reaction                                                   |
| _ReactionApi_      | [**ReactionsCast**](docs/ReactionApi.md#reactionscast)                                                         | **Get** /farcaster/reactions/cast                       | Fetches reactions for a given cast                                 |
| _ReactionApi_      | [**ReactionsUser**](docs/ReactionApi.md#reactionsuser)                                                         | **Get** /farcaster/reactions/user                       | Fetches reactions for a given user                                 |
| _SignerApi_        | [**CreateSigner**](docs/SignerApi.md#createsigner)                                                             | **Post** /farcaster/signer                              | Creates a signer and returns the signer status                     |
| _SignerApi_        | [**DeveloperManagedSigner**](docs/SignerApi.md#developermanagedsigner)                                         | **Get** /farcaster/signer/developer_managed             | Fetches the status of a signer by public key                       |
| _SignerApi_        | [**FetchAuthorizationUrl**](docs/SignerApi.md#fetchauthorizationurl)                                           | **Get** /farcaster/login/authorize                      | Fetch authorization url                                            |
| _SignerApi_        | [**PublishMessage**](docs/SignerApi.md#publishmessage)                                                         | **Post** /farcaster/message                             | Publish a message to farcaster                                     |
| _SignerApi_        | [**RegisterSignedKey**](docs/SignerApi.md#registersignedkey)                                                   | **Post** /farcaster/signer/signed_key                   | Register Signed Key                                                |
| _SignerApi_        | [**RegisterSignedKeyForDeveloperManagedSigner**](docs/SignerApi.md#registersignedkeyfordevelopermanagedsigner) | **Post** /farcaster/signer/developer_managed/signed_key | Registers Signed Key                                               |
| _SignerApi_        | [**Signer**](docs/SignerApi.md#signer)                                                                         | **Get** /farcaster/signer                               | Fetches the status of a signer                                     |
| _StorageApi_       | [**BuyStorage**](docs/StorageApi.md#buystorage)                                                                | **Post** /farcaster/storage/buy                         | Buy storage for an fid                                             |
| _StorageApi_       | [**StorageAllocations**](docs/StorageApi.md#storageallocations)                                                | **Get** /farcaster/storage/allocations                  | Fetches storage allocations for a given user                       |
| _StorageApi_       | [**StorageUsage**](docs/StorageApi.md#storageusage)                                                            | **Get** /farcaster/storage/usage                        | Fetches storage usage for a given user                             |
| _SubscribersApi_   | [**SubscribedTo**](docs/SubscribersApi.md#subscribedto)                                                        | **Get** /farcaster/user/subscribed_to                   | Fetch what a given fid is subscribed to                            |
| _SubscribersApi_   | [**Subscribers**](docs/SubscribersApi.md#subscribers)                                                          | **Get** /farcaster/user/subscribers                     | Fetch subscribers for a given fid                                  |
| _SubscribersApi_   | [**SubscriptionsCreated**](docs/SubscribersApi.md#subscriptionscreated)                                        | **Get** /farcaster/user/subscriptions_created           | Fetch created subscriptions for a given fid                        |
| _UserApi_          | [**ActiveUsers**](docs/UserApi.md#activeusers)                                                                 | **Get** /farcaster/user/active                          | Fetch active users                                                 |
| _UserApi_          | [**FarcasterUserVerificationDelete**](docs/UserApi.md#farcasteruserverificationdelete)                         | **Delete** /farcaster/user/verification                 | Removes verification for an eth address for the user               |
| _UserApi_          | [**FarcasterUserVerificationPost**](docs/UserApi.md#farcasteruserverificationpost)                             | **Post** /farcaster/user/verification                   | Adds verification for an ethereum address or contract for the user |
| _UserApi_          | [**FollowUser**](docs/UserApi.md#followuser)                                                                   | **Post** /farcaster/user/follow                         | Follow a user                                                      |
| _UserApi_          | [**GetFreshFid**](docs/UserApi.md#getfreshfid)                                                                 | **Get** /farcaster/user/fid                             | Fetches fid to assign it new user                                  |
| _UserApi_          | [**LookupUserByCustodyAddress**](docs/UserApi.md#lookupuserbycustodyaddress)                                   | **Get** /farcaster/user/custody-address                 | Lookup a user by custody-address                                   |
| _UserApi_          | [**PowerUsers**](docs/UserApi.md#powerusers)                                                                   | **Get** /farcaster/user/power                           | Fetch power user objects                                           |
| _UserApi_          | [**RegisterUser**](docs/UserApi.md#registeruser)                                                               | **Post** /farcaster/user                                | Register account on farcaster                                      |
| _UserApi_          | [**UnfollowUser**](docs/UserApi.md#unfollowuser)                                                               | **Delete** /farcaster/user/follow                       | Unfollow a user                                                    |
| _UserApi_          | [**UpdateUser**](docs/UserApi.md#updateuser)                                                                   | **Patch** /farcaster/user                               | Update user profile                                                |
| _UserApi_          | [**UserBulk**](docs/UserApi.md#userbulk)                                                                       | **Get** /farcaster/user/bulk                            | Fetch users based on FIDs                                          |
| _UserApi_          | [**UserBulkByAddress**](docs/UserApi.md#userbulkbyaddress)                                                     | **Get** /farcaster/user/bulk-by-address                 | Fetches users based on Eth or Sol addresses                        |
| _UserApi_          | [**UserPowerLite**](docs/UserApi.md#userpowerlite)                                                             | **Get** /farcaster/user/power_lite                      | Fetch power user FIDs                                              |
| _UserApi_          | [**UserSearch**](docs/UserApi.md#usersearch)                                                                   | **Get** /farcaster/user/search                          | Search for Usernames                                               |
| _WebhookApi_       | [**DeleteWebhook**](docs/WebhookApi.md#deletewebhook)                                                          | **Delete** /farcaster/webhook                           | Delete a webhook                                                   |
| _WebhookApi_       | [**FetchWebhooks**](docs/WebhookApi.md#fetchwebhooks)                                                          | **Get** /farcaster/webhook/list                         | Fetch a list of webhooks associated to a user                      |
| _WebhookApi_       | [**LookupWebhook**](docs/WebhookApi.md#lookupwebhook)                                                          | **Get** /farcaster/webhook                              | Fetch a webhook                                                    |
| _WebhookApi_       | [**PublishWebhook**](docs/WebhookApi.md#publishwebhook)                                                        | **Post** /farcaster/webhook                             | Create a webhook                                                   |
| _WebhookApi_       | [**UpdateWebhook**](docs/WebhookApi.md#updatewebhook)                                                          | **Put** /farcaster/webhook                              | Update a webhook                                                   |
| _WebhookApi_       | [**UpdateWebhookActiveStatus**](docs/WebhookApi.md#updatewebhookactivestatus)                                  | **Patch** /farcaster/webhook                            | Update webhook active status                                       |

## Documentation For Models

- [ActiveStatus](docs/ActiveStatus.md)
- [AddVerificationReqBody](docs/AddVerificationReqBody.md)
- [AllOfCastParentAuthor](docs/AllOfCastParentAuthor.md)
- [AuthorizationUrlResponse](docs/AuthorizationUrlResponse.md)
- [AuthorizationUrlResponseType](docs/AuthorizationUrlResponseType.md)
- [BulkCastsResponse](docs/BulkCastsResponse.md)
- [BulkFollowResponse](docs/BulkFollowResponse.md)
- [BulkUsersResponse](docs/BulkUsersResponse.md)
- [BuyStorageReqBody](docs/BuyStorageReqBody.md)
- [Cast](docs/Cast.md)
- [CastComposerActionsListResponse](docs/CastComposerActionsListResponse.md)
- [CastComposerActionsListResponseAction](docs/CastComposerActionsListResponseAction.md)
- [CastComposerActionsListResponseActions](docs/CastComposerActionsListResponseActions.md)
- [CastComposerType](docs/CastComposerType.md)
- [CastDehydrated](docs/CastDehydrated.md)
- [CastId](docs/CastId.md)
- [CastNotificationType](docs/CastNotificationType.md)
- [CastParamType](docs/CastParamType.md)
- [CastResponse](docs/CastResponse.md)
- [CastViewerContext](docs/CastViewerContext.md)
- [CastWithInteractions](docs/CastWithInteractions.md)
- [CastWithInteractionsAndConversations](docs/CastWithInteractionsAndConversations.md)
- [CastWithInteractionsReactions](docs/CastWithInteractionsReactions.md)
- [CastWithInteractionsReplies](docs/CastWithInteractionsReplies.md)
- [CastsResponse](docs/CastsResponse.md)
- [CastsResponseResult](docs/CastsResponseResult.md)
- [CastsSearchResponse](docs/CastsSearchResponse.md)
- [CastsSearchResponseResult](docs/CastsSearchResponseResult.md)
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
- [DeveloperManagedSigner](docs/DeveloperManagedSigner.md)
- [EmbedCastId](docs/EmbedCastId.md)
- [EmbedType](docs/EmbedType.md)
- [EmbedUrl](docs/EmbedUrl.md)
- [EmbedUrlMetadata](docs/EmbedUrlMetadata.md)
- [EmbedUrlMetadataImage](docs/EmbedUrlMetadataImage.md)
- [EmbedUrlMetadataVideo](docs/EmbedUrlMetadataVideo.md)
- [EmbedUrlMetadataVideoStream](docs/EmbedUrlMetadataVideoStream.md)
- [EmbeddedCast](docs/EmbeddedCast.md)
- [ErrorRes](docs/ErrorRes.md)
- [FarcasterFrameBody](docs/FarcasterFrameBody.md)
- [FeedResponse](docs/FeedResponse.md)
- [FeedTrendingProvider](docs/FeedTrendingProvider.md)
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
- [FrameValidateAnalyticsInputTextInputTexts](docs/FrameValidateAnalyticsInputTextInputTexts.md)
- [FrameValidateAnalyticsInteractionsPerCast](docs/FrameValidateAnalyticsInteractionsPerCast.md)
- [FrameValidateAnalyticsInteractionsPerCastInteractionsPerCast](docs/FrameValidateAnalyticsInteractionsPerCastInteractionsPerCast.md)
- [FrameValidateAnalyticsInteractors](docs/FrameValidateAnalyticsInteractors.md)
- [FrameValidateAnalyticsInteractorsInteractors](docs/FrameValidateAnalyticsInteractorsInteractors.md)
- [FrameValidateAnalyticsResponse](docs/FrameValidateAnalyticsResponse.md)
- [FrameValidateAnalyticsTotalInteractors](docs/FrameValidateAnalyticsTotalInteractors.md)
- [FrameValidateBody](docs/FrameValidateBody.md)
- [FrameValidateListResponse](docs/FrameValidateListResponse.md)
- [GetCastsReqBody](docs/GetCastsReqBody.md)
- [HydratedFollower](docs/HydratedFollower.md)
- [ImageObject](docs/ImageObject.md)
- [IndividualHashObj](docs/IndividualHashObj.md)
- [InlineResponse200](docs/InlineResponse200.md)
- [InlineResponse400](docs/InlineResponse400.md)
- [MarkNotificationsAsSeenReqBody](docs/MarkNotificationsAsSeenReqBody.md)
- [MusicSongObject](docs/MusicSongObject.md)
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
- [NeynarPageImage](docs/NeynarPageImage.md)
- [NeynarPageInput](docs/NeynarPageInput.md)
- [NeynarPageInputText](docs/NeynarPageInputText.md)
- [Notification](docs/Notification.md)
- [NotificationType](docs/NotificationType.md)
- [NotificationsResponse](docs/NotificationsResponse.md)
- [OgObject](docs/OgObject.md)
- [OneOfNeynarPageButtonNextPage](docs/OneOfNeynarPageButtonNextPage.md)
- [OperationResponse](docs/OperationResponse.md)
- [PostCastReqBody](docs/PostCastReqBody.md)
- [PostCastResponse](docs/PostCastResponse.md)
- [PostCastResponseCast](docs/PostCastResponseCast.md)
- [PostCastResponseCastAuthor](docs/PostCastResponseCastAuthor.md)
- [ProfileUrl](docs/ProfileUrl.md)
- [ProfileUrlPfp](docs/ProfileUrlPfp.md)
- [PublishMessageReqBody](docs/PublishMessageReqBody.md)
- [PublishMessageResponse](docs/PublishMessageResponse.md)
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
- [RegisterSignerKeyReqBodySponsor](docs/RegisterSignerKeyReqBodySponsor.md)
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
- [SubscriptionStatus](docs/SubscriptionStatus.md)
- [SubscriptionTier](docs/SubscriptionTier.md)
- [SubscriptionTierPrice](docs/SubscriptionTierPrice.md)
- [SubscriptionToken](docs/SubscriptionToken.md)
- [Subscriptions](docs/Subscriptions.md)
- [SubscriptionsResponse](docs/SubscriptionsResponse.md)
- [TrendingChannelResponse](docs/TrendingChannelResponse.md)
- [TwitterImageObject](docs/TwitterImageObject.md)
- [TwitterPlayerObject](docs/TwitterPlayerObject.md)
- [UpdateUserReqBody](docs/UpdateUserReqBody.md)
- [User](docs/User.md)
- [UserDehydrated](docs/UserDehydrated.md)
- [UserFidResponse](docs/UserFidResponse.md)
- [UserPowerLiteResponse](docs/UserPowerLiteResponse.md)
- [UserPowerLiteResponseResult](docs/UserPowerLiteResponseResult.md)
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
- [ValidatedFrameAction](docs/ValidatedFrameAction.md)
- [ValidatedFrameActionSigner](docs/ValidatedFrameActionSigner.md)
- [ValidatedFrameActionTappedButton](docs/ValidatedFrameActionTappedButton.md)
- [VerificationChainId](docs/VerificationChainId.md)
- [VerificationType](docs/VerificationType.md)
- [VideoObject](docs/VideoObject.md)
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
- [ZodErrorErrors](docs/ZodErrorErrors.md)

## Documentation For Authorization

Endpoints do not require authorization.

## Author
