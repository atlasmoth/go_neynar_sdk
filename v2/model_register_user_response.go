/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RegisterUserResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Signer *Signer `json:"signer"`
}
