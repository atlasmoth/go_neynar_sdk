/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type EmbedUrlMetadata struct {
	Status string `json:"_status"`
	ContentType string `json:"content_type,omitempty"`
	ContentLength int32 `json:"content_length,omitempty"`
	Image *EmbedUrlMetadataImage `json:"image,omitempty"`
	Video *EmbedUrlMetadataVideo `json:"video,omitempty"`
	Html *OgObject `json:"html,omitempty"`
}
