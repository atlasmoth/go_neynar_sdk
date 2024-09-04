/*
 * Raw Farcaster Hub API
 *
 * Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Adds a Verification of ownership of an Ethereum Address
type VerificationAddEthAddressBody struct {
	Address string `json:"address"`
	EthSignature string `json:"ethSignature"`
	BlockHash string `json:"blockHash"`
}
