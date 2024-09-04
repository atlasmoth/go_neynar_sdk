/*
 * Raw Farcaster Hub API
 *
 * Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// HashScheme : Type of hashing scheme used to produce a digest of MessageData. - HASH_SCHEME_BLAKE3: Default scheme for hashing MessageData 
type HashScheme string

// List of HashScheme
const (
	HASH_SCHEME_BLAKE3_HashScheme HashScheme = "HASH_SCHEME_BLAKE3"
)
