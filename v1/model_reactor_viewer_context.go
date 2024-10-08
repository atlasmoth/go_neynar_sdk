/*
 * Farcaster API V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ReactorViewerContext struct {
	// Indicates if the viewer is following the reactor.
	Following bool `json:"following"`
	// Indicates if the reactor is followed by the viewer.
	FollowedBy bool `json:"followedBy"`
}
