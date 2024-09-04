/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Signer struct {
	SignerUuid string `json:"signer_uuid"`
	PublicKey string `json:"public_key"`
	Status string `json:"status"`
	SignerApprovalUrl string `json:"signer_approval_url,omitempty"`
	Fid int32 `json:"fid,omitempty"`
}
