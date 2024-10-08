/*
 * Farcaster API V2
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Frame struct {
	// Version of the frame
	Version string `json:"version"`
	// URL of the image
	Image string `json:"image"`
	Buttons []FrameActionButton `json:"buttons,omitempty"`
	// Post URL to take an action on this frame
	PostUrl string `json:"post_url,omitempty"`
	// URL of the frames
	FramesUrl string `json:"frames_url"`
	Title string `json:"title,omitempty"`
	ImageAspectRatio string `json:"image_aspect_ratio,omitempty"`
	Input *FrameInput `json:"input,omitempty"`
	State *FrameState `json:"state,omitempty"`
}
