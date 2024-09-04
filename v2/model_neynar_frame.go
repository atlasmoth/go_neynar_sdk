/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NeynarFrame struct {
	// Unique identifier for the frame.
	Uuid string `json:"uuid"`
	// Name of the frame.
	Name string `json:"name"`
	// Generated link for the frame's first page.
	Link string `json:"link"`
	Pages []NeynarFramePage `json:"pages"`
	// Indicates if the frame is valid.
	Valid bool `json:"valid,omitempty"`
}
