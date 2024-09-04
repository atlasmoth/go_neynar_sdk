## Documentation for API Endpoints

All URIs are relative to *https://api.neynar.com/v1*

| Class              | Method                                                                  | HTTP request                             | Description                                             |
| ------------------ | ----------------------------------------------------------------------- | ---------------------------------------- | ------------------------------------------------------- |
| _CastApi_          | [**AllCastsInThread**](docs/CastApi.md#allcastsinthread)                | **Get** /farcaster/all-casts-in-thread   | DEPRECATED - Retrieve all casts in a given thread hash  |
| _CastApi_          | [**Cast**](docs/CastApi.md#cast)                                        | **Get** /farcaster/cast                  | DEPRECATED - Retrieve cast for a given hash             |
| _CastApi_          | [**Casts**](docs/CastApi.md#casts)                                      | **Get** /farcaster/casts                 | DEPRECATED - Retrieve casts for a given user            |
| _CastApi_          | [**RecentCasts**](docs/CastApi.md#recentcasts)                          | **Get** /farcaster/recent-casts          | Get Recent Casts                                        |
| _FollowsApi_       | [**Followers**](docs/FollowsApi.md#followers)                           | **Get** /farcaster/followers             | Gets all followers for a given FID                      |
| _FollowsApi_       | [**Following**](docs/FollowsApi.md#following)                           | **Get** /farcaster/following             | Gets all following users of a FID                       |
| _NotificationsApi_ | [**MentionsAndReplies**](docs/NotificationsApi.md#mentionsandreplies)   | **Get** /farcaster/mentions-and-replies  | Get mentions and replies                                |
| _NotificationsApi_ | [**ReactionsAndRecasts**](docs/NotificationsApi.md#reactionsandrecasts) | **Get** /farcaster/reactions-and-recasts | Get reactions and recasts                               |
| _ReactionsApi_     | [**CastLikes**](docs/ReactionsApi.md#castlikes)                         | **Get** /farcaster/cast-likes            | DEPRECATED - Get all like reactions for a specific cast |
| _ReactionsApi_     | [**CastReactions**](docs/ReactionsApi.md#castreactions)                 | **Get** /farcaster/cast-reactions        | DEPRECATED - Get all reactions for a specific cast      |
| _ReactionsApi_     | [**CastRecasters**](docs/ReactionsApi.md#castrecasters)                 | **Get** /farcaster/cast-recasters        | DEPRECATED - Get all recasters for a specific cast      |
| _UserApi_          | [**CustodyAddress**](docs/UserApi.md#custodyaddress)                    | **Get** /farcaster/custody-address       | DEPRECATED - Get the custody address for a given FID    |
| _UserApi_          | [**RecentUsers**](docs/UserApi.md#recentusers)                          | **Get** /farcaster/recent-users          | Get Recent Users                                        |
| _UserApi_          | [**User**](docs/UserApi.md#user)                                        | **Get** /farcaster/user                  | DEPRECATED - Get User Information by FID                |
| _UserApi_          | [**UserByUsername**](docs/UserApi.md#userbyusername)                    | **Get** /farcaster/user-by-username      | Get User Information by username                        |
| _UserApi_          | [**UserCastLikes**](docs/UserApi.md#usercastlikes)                      | **Get** /farcaster/user-cast-likes       | DEPRECATED -- Get User Cast Likes                       |
| _VerificationApi_  | [**UserByVerification**](docs/VerificationApi.md#userbyverification)    | **Get** /farcaster/user-by-verification  | DEPRECATED - Retrieve user for a given ethereum address |
| _VerificationApi_  | [**Verifications**](docs/VerificationApi.md#verifications)              | **Get** /farcaster/verifications         | DEPRECATED - Retrieve verifications for a given FID     |

## Documentation For Models

- [ActiveStatus](docs/ActiveStatus.md)
- [AllCastsInThreadResponse](docs/AllCastsInThreadResponse.md)
- [AllCastsInThreadResponseResult](docs/AllCastsInThreadResponseResult.md)
- [AllOfCastParentAuthor](docs/AllOfCastParentAuthor.md)
- [Cast](docs/Cast.md)
- [CastLikesResponse](docs/CastLikesResponse.md)
- [CastLikesResponseResult](docs/CastLikesResponseResult.md)
- [CastReactionsResponse](docs/CastReactionsResponse.md)
- [CastReactionsResponseResult](docs/CastReactionsResponseResult.md)
- [CastRecasterResponse](docs/CastRecasterResponse.md)
- [CastRecasterResponseResult](docs/CastRecasterResponseResult.md)
- [CastResponse](docs/CastResponse.md)
- [CastResponseResult](docs/CastResponseResult.md)
- [CastType](docs/CastType.md)
- [CastWithInteractions](docs/CastWithInteractions.md)
- [CastWithInteractionsReactions](docs/CastWithInteractionsReactions.md)
- [CastWithInteractionsRecasts](docs/CastWithInteractionsRecasts.md)
- [CastWithInteractionsReplies](docs/CastWithInteractionsReplies.md)
- [CastsResponse](docs/CastsResponse.md)
- [CastsResponseResult](docs/CastsResponseResult.md)
- [CustodyAddressResponse](docs/CustodyAddressResponse.md)
- [CustodyAddressResponseResult](docs/CustodyAddressResponseResult.md)
- [EmbedUrl](docs/EmbedUrl.md)
- [ErrorRes](docs/ErrorRes.md)
- [FollowResponse](docs/FollowResponse.md)
- [FollowResponseResult](docs/FollowResponseResult.md)
- [FollowResponseUser](docs/FollowResponseUser.md)
- [MentionsAndRepliesResponse](docs/MentionsAndRepliesResponse.md)
- [MentionsAndRepliesResponseResult](docs/MentionsAndRepliesResponseResult.md)
- [NextCursor](docs/NextCursor.md)
- [OneOfCastAuthor](docs/OneOfCastAuthor.md)
- [Reaction](docs/Reaction.md)
- [ReactionType](docs/ReactionType.md)
- [ReactionWithCastMeta](docs/ReactionWithCastMeta.md)
- [ReactionWithCastMetaCast](docs/ReactionWithCastMetaCast.md)
- [ReactionWithCastMetaReaction](docs/ReactionWithCastMetaReaction.md)
- [ReactionsAndRecastsNotification](docs/ReactionsAndRecastsNotification.md)
- [ReactionsAndRecastsResponse](docs/ReactionsAndRecastsResponse.md)
- [ReactionsAndRecastsResponseResult](docs/ReactionsAndRecastsResponseResult.md)
- [Reactor](docs/Reactor.md)
- [ReactorPfp](docs/ReactorPfp.md)
- [ReactorViewerContext](docs/ReactorViewerContext.md)
- [Recaster](docs/Recaster.md)
- [RecasterPfp](docs/RecasterPfp.md)
- [RecasterProfile](docs/RecasterProfile.md)
- [RecasterProfileBio](docs/RecasterProfileBio.md)
- [RecasterViewerContext](docs/RecasterViewerContext.md)
- [RecentCastsResponse](docs/RecentCastsResponse.md)
- [RecentUsersResponse](docs/RecentUsersResponse.md)
- [RecentUsersResponseResult](docs/RecentUsersResponseResult.md)
- [User](docs/User.md)
- [UserCastLikeResponse](docs/UserCastLikeResponse.md)
- [UserCastLikeResponseResult](docs/UserCastLikeResponseResult.md)
- [UserPfp](docs/UserPfp.md)
- [UserProfile](docs/UserProfile.md)
- [UserProfileBio](docs/UserProfileBio.md)
- [UserResponse](docs/UserResponse.md)
- [UserResponseResult](docs/UserResponseResult.md)
- [VerificationResponse](docs/VerificationResponse.md)
- [VerificationResponseResult](docs/VerificationResponseResult.md)
- [ViewerContext](docs/ViewerContext.md)

## Documentation For Authorization

Endpoints do not require authorization.

## Author
