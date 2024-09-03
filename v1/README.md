## Documentation for API Endpoints

All URIs are relative to *https://api.neynar.com/v1*

| Class              | Method                                                                  | HTTP request                             | Description                                             |
| ------------------ | ----------------------------------------------------------------------- | ---------------------------------------- | ------------------------------------------------------- |
| _CastAPI_          | [**AllCastsInThread**](docs/CastAPI.md#allcastsinthread)                | **Get** /farcaster/all-casts-in-thread   | DEPRECATED - Retrieve all casts in a given thread hash  |
| _CastAPI_          | [**Cast**](docs/CastAPI.md#cast)                                        | **Get** /farcaster/cast                  | DEPRECATED - Retrieve cast for a given hash             |
| _CastAPI_          | [**Casts**](docs/CastAPI.md#casts)                                      | **Get** /farcaster/casts                 | DEPRECATED - Retrieve casts for a given user            |
| _CastAPI_          | [**RecentCasts**](docs/CastAPI.md#recentcasts)                          | **Get** /farcaster/recent-casts          | Get Recent Casts                                        |
| _FollowsAPI_       | [**Followers**](docs/FollowsAPI.md#followers)                           | **Get** /farcaster/followers             | Gets all followers for a given FID                      |
| _FollowsAPI_       | [**Following**](docs/FollowsAPI.md#following)                           | **Get** /farcaster/following             | Gets all following users of a FID                       |
| _NotificationsAPI_ | [**MentionsAndReplies**](docs/NotificationsAPI.md#mentionsandreplies)   | **Get** /farcaster/mentions-and-replies  | Get mentions and replies                                |
| _NotificationsAPI_ | [**ReactionsAndRecasts**](docs/NotificationsAPI.md#reactionsandrecasts) | **Get** /farcaster/reactions-and-recasts | Get reactions and recasts                               |
| _ReactionsAPI_     | [**CastLikes**](docs/ReactionsAPI.md#castlikes)                         | **Get** /farcaster/cast-likes            | DEPRECATED - Get all like reactions for a specific cast |
| _ReactionsAPI_     | [**CastReactions**](docs/ReactionsAPI.md#castreactions)                 | **Get** /farcaster/cast-reactions        | DEPRECATED - Get all reactions for a specific cast      |
| _ReactionsAPI_     | [**CastRecasters**](docs/ReactionsAPI.md#castrecasters)                 | **Get** /farcaster/cast-recasters        | DEPRECATED - Get all recasters for a specific cast      |
| _UserAPI_          | [**CustodyAddress**](docs/UserAPI.md#custodyaddress)                    | **Get** /farcaster/custody-address       | DEPRECATED - Get the custody address for a given FID    |
| _UserAPI_          | [**RecentUsers**](docs/UserAPI.md#recentusers)                          | **Get** /farcaster/recent-users          | Get Recent Users                                        |
| _UserAPI_          | [**User**](docs/UserAPI.md#user)                                        | **Get** /farcaster/user                  | DEPRECATED - Get User Information by FID                |
| _UserAPI_          | [**UserByUsername**](docs/UserAPI.md#userbyusername)                    | **Get** /farcaster/user-by-username      | Get User Information by username                        |
| _UserAPI_          | [**UserCastLikes**](docs/UserAPI.md#usercastlikes)                      | **Get** /farcaster/user-cast-likes       | DEPRECATED -- Get User Cast Likes                       |
| _VerificationAPI_  | [**UserByVerification**](docs/VerificationAPI.md#userbyverification)    | **Get** /farcaster/user-by-verification  | DEPRECATED - Retrieve user for a given ethereum address |
| _VerificationAPI_  | [**Verifications**](docs/VerificationAPI.md#verifications)              | **Get** /farcaster/verifications         | DEPRECATED - Retrieve verifications for a given FID     |

## Documentation For Models

- [ActiveStatus](docs/ActiveStatus.md)
- [AllCastsInThreadResponse](docs/AllCastsInThreadResponse.md)
- [AllCastsInThreadResponseResult](docs/AllCastsInThreadResponseResult.md)
- [Cast](docs/Cast.md)
- [CastAuthor](docs/CastAuthor.md)
- [CastAuthorOneOf](docs/CastAuthorOneOf.md)
- [CastLikesResponse](docs/CastLikesResponse.md)
- [CastLikesResponseResult](docs/CastLikesResponseResult.md)
- [CastParentAuthor](docs/CastParentAuthor.md)
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
