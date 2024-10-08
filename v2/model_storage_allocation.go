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

type StorageAllocation struct {
	Object string `json:"object,omitempty"`
	User *UserDehydrated `json:"user,omitempty"`
	Units int32 `json:"units,omitempty"`
	Expiry time.Time `json:"expiry,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
