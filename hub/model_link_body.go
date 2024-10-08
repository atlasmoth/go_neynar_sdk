/*
 * Raw Farcaster Hub API
 *
 * Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Adds or removes a Link
type LinkBody struct {
	Type_ *LinkType `json:"type"`
	DisplayTimestamp int64 `json:"displayTimestamp,omitempty"`
	TargetFid int32 `json:"targetFid"`
}
