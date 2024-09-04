/*
 * Raw Farcaster Hub API
 *
 * Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type VerificationRemove struct {
	Data *interface{} `json:"data"`
	Hash string `json:"hash"`
	HashScheme *HashScheme `json:"hashScheme"`
	Signature string `json:"signature"`
	SignatureScheme *SignatureScheme `json:"signatureScheme"`
	Signer string `json:"signer"`
}
