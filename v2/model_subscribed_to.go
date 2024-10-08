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

type SubscribedTo struct {
	ExpiresAt time.Time `json:"expires_at"`
	SubscribedAt time.Time `json:"subscribed_at"`
	Tier *SubscriptionTier `json:"tier"`
	Creator *User `json:"creator"`
	Object string `json:"object"`
	ProviderName string `json:"provider_name,omitempty"`
	ContractAddress string `json:"contract_address"`
	Chain int32 `json:"chain"`
	Metadata *SubscriptionMetadata `json:"metadata"`
	OwnerAddress string `json:"owner_address"`
	Price *SubscriptionPrice `json:"price"`
	Tiers []SubscriptionTier `json:"tiers,omitempty"`
	ProtocolVersion int32 `json:"protocol_version"`
	Token *SubscriptionToken `json:"token"`
}
