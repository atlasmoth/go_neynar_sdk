/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type FrameValidateBody struct {
	// Hexadecimal string of message bytes.
	MessageBytesInHex string `json:"message_bytes_in_hex"`
	// Adds viewer_context inside the cast object to indicate whether the interactor reacted to the cast housing the frame.
	CastReactionContext bool `json:"cast_reaction_context,omitempty"`
	// Adds viewer_context inside the user (interactor) object to indicate whether the interactor follows or is followed by the cast author.
	FollowContext bool `json:"follow_context,omitempty"`
	// Adds context about the app used by the user inside `frame.action`.
	SignerContext bool `json:"signer_context,omitempty"`
	// Adds context about the channel that the cast belongs to inside of the cast object.
	ChannelFollowContext bool `json:"channel_follow_context,omitempty"`
}
