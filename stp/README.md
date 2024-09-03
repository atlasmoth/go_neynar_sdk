## Documentation for API Endpoints

All URIs are relative to *https://api.neynar.com/v2*

| Class    | Method                                                    | HTTP request                    | Description                                               |
| -------- | --------------------------------------------------------- | ------------------------------- | --------------------------------------------------------- |
| _STPAPI_ | [**SubscriptionCheck**](docs/STPAPI.md#subscriptioncheck) | **Get** /stp/subscription_check | Check if a wallet address is subscribed to a STP contract |

## Documentation For Models

- [ConflictErrorRes](docs/ConflictErrorRes.md)
- [ErrorRes](docs/ErrorRes.md)
- [SubscriptionStatus](docs/SubscriptionStatus.md)
- [SubscriptionTier](docs/SubscriptionTier.md)
- [SubscriptionTierPrice](docs/SubscriptionTierPrice.md)

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
