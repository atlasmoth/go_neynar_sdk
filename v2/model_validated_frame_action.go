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

type ValidatedFrameAction struct {
	Object string `json:"object"`
	Url string `json:"url"`
	Interactor *User `json:"interactor"`
	TappedButton *ValidatedFrameActionTappedButton `json:"tapped_button"`
	Input *FrameInput `json:"input,omitempty"`
	State *FrameState `json:"state"`
	Cast *CastWithInteractions `json:"cast"`
	Timestamp *time.Time `json:"timestamp"`
	Signer *ValidatedFrameActionSigner `json:"signer,omitempty"`
	Transaction *FrameTransaction `json:"transaction,omitempty"`
	Address string `json:"address,omitempty"`
}
