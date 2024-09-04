/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type Cast struct {
	Hash string `json:"hash"`
	ParentHash string `json:"parent_hash"`
	ParentUrl string `json:"parent_url"`
	RootParentUrl string `json:"root_parent_url"`
	ParentAuthor *AllOfCastParentAuthor `json:"parent_author"`
	Author *User `json:"author"`
	Text string `json:"text"`
	Timestamp *time.Time `json:"timestamp"`
	Embeds []EmbeddedCast `json:"embeds"`
	Type_ *CastNotificationType `json:"type,omitempty"`
}
