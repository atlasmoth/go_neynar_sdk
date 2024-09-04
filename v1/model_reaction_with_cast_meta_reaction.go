/*
 * Farcaster API V1
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type ReactionWithCastMetaReaction struct {
	ReactorFid int32 `json:"reactor_fid"`
	ReactionType *ReactionType `json:"reaction_type"`
	ReactionHash string `json:"reaction_hash"`
	ReactionTargetHash string `json:"reaction_target_hash"`
	ReactionTimestamp time.Time `json:"reaction_timestamp"`
}
