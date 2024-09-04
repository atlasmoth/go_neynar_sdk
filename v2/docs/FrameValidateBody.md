# FrameValidateBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageBytesInHex** | **string** | Hexadecimal string of message bytes. | [default to null]
**CastReactionContext** | **bool** | Adds viewer_context inside the cast object to indicate whether the interactor reacted to the cast housing the frame. | [optional] [default to true]
**FollowContext** | **bool** | Adds viewer_context inside the user (interactor) object to indicate whether the interactor follows or is followed by the cast author. | [optional] [default to false]
**SignerContext** | **bool** | Adds context about the app used by the user inside &#x60;frame.action&#x60;. | [optional] [default to false]
**ChannelFollowContext** | **bool** | Adds context about the channel that the cast belongs to inside of the cast object. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

