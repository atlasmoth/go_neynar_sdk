# go_neynar_sdk

go_neynar_sdk is a Go client library for accessing the [Neynar API's](https://docs.neynar.com/).

## Install

```
go get github.com/atlasmoth/go_neynar_sdk@vX.Y.Z
```

where X.Y.Z is the version you need.

## Usage

```
import "github.com/atlasmoth/go_neynar_sdk"
```

Create a new Neynar client, then use the exposed services to access different parts of the Neynar API.

## Authentication

You need to provide an API key to create a new client:

```
package main

import (
    "github.com/atlasmoth/go_neynar_sdk"
)

func main() {
    client, err := neynarsdk.NewClient(nil, "your-neynar-api-key")
    if err != nil {
        // Handle error
    }
    // Use the client
}
```

## Examples

To use the Feed service:

```
ctx := context.Background()
feedParams := &neynarsdk.FeedParams{
    // Set your feed parameters here
}
feed, err := client.Feed.GetFeed(ctx, feedParams)
if err != nil {
    // Handle error
}
// Use the feed data
```

## Services

The SDK provides the following services:

- Feed
- Cast
- Notification
- Follow
- Webhook
- Subscription
- Storage
- Mute
- Fname
- Reaction
- Signer
- Frame
- Channel
- User

Each service corresponds to a group of API endpoints. Refer to the Neynar API documentation for more details on available endpoints and their usage.

## Versioning

Each version of the client is tagged and the version is updated accordingly. The current version is 0.0.0.

To see the list of past versions, `run git tag`.

## Documentation

For a comprehensive list of examples, check out the API documentation.

For details on all the functionality in this library, see the GoDoc documentation.
