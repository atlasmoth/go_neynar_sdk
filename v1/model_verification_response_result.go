/*
 * Farcaster API V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerificationResponseResult struct {
	Fid string `json:"fid"`
	Username string `json:"username"`
	DisplayName string `json:"display_name"`
	Verifications []string `json:"verifications"`
}
