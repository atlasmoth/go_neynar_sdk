/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NeynarFrameCreationRequest struct {
	// The name of the frame.
	Name string `json:"name"`
	Pages []NeynarFramePage `json:"pages"`
}
