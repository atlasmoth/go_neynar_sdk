package neynarsdk

import (
	"context"
	"net/http"
	"time"
)

type FrameService struct {
	client *Client
}

type PostFrameActionResponse struct {
	Frame
	ErrorResponse
}

type FrameActionReqBody struct {
	SignerUUID SignerUUID  `json:"signer_uuid"`
	CastHash   CastHash    `json:"cast_hash"`
	Action     FrameAction `json:"action"`
}

func (s *FrameService) PostFrameAction(ctx context.Context, reqBody FrameActionReqBody) (PostFrameActionResponse, error) {
	var result PostFrameActionResponse

	if reqBody.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if reqBody.CastHash == "" {
		return result, &RequiredFieldError{Field: "CastHash"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/action"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, reqBody, nil)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type ValidateFrameActionRequest struct {
	MessageBytesInHex       string `json:"message_bytes_in_hex"`
	OmitCastReactionContext bool   `json:"cast_reaction_context,omitempty"`
	FollowContext           bool   `json:"follow_context,omitempty"`
	SignerContext           bool   `json:"signer_context,omitempty"`
	ChannelFollowContext    bool   `json:"channel_follow_context,omitempty"`
}

type ValidateFrameActionResponse struct {
	Valid  bool                 `json:"valid"`
	Action ValidatedFrameAction `json:"action"`
	ErrorResponse
}

func (s *FrameService) ValidateFrame(ctx context.Context, reqBody ValidateFrameActionRequest) (ValidateFrameActionResponse, error) {
	var result ValidateFrameActionResponse
	reqBody.OmitCastReactionContext = !reqBody.OmitCastReactionContext

	if reqBody.MessageBytesInHex == "" {
		return result, &RequiredFieldError{Field: "message_bytes_in_hex"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/validate"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, reqBody, nil)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type FrameValidateListResponse struct {
	Frames []string `json:"frames"`
	ErrorResponse
}

func (s *FrameService) GetValidatedFrames(ctx context.Context) (FrameValidateListResponse, error) {
	var result FrameValidateListResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/validate/list"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, nil)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type FrameAnalyticsResponse struct {
	FrameValidateAnalyticsInteractors
	FrameValidateAnalyticsTotalInteractors
	FrameValidateAnalyticsInteractionsPerCast
	FrameValidateAnalyticsInputText
	ErrorResponse
}

type FrameAnalyticsRequestParams struct {
	FrameURL        string    `json:"frame_url"`
	AnalyticsType   string    `json:"analytics_type"`
	Start           time.Time `json:"start"`
	Stop            time.Time `json:"stop"`
	AggregateWindow string    `json:"aggregate_window,omitempty"`
}

func (s *FrameService) GetFrameAnalytics(ctx context.Context, params FrameAnalyticsRequestParams) (FrameAnalyticsResponse, error) {
	var result FrameAnalyticsResponse

	if params.FrameURL == "" {
		return result, &RequiredFieldError{Field: "FrameURL"}
	}
	if params.AnalyticsType == "" {
		return result, &RequiredFieldError{Field: "AnalyticsType"}
	}
	if params.Start.IsZero() {
		return result, &RequiredFieldError{Field: "Start"}
	}
	if params.Stop.IsZero() {
		return result, &RequiredFieldError{Field: "Stop"}
	}

	validAnalyticsTypes := []string{"total-interactors", "interactors", "interactions-per-cast", "input-text"}
	if !Contains(validAnalyticsTypes, params.AnalyticsType) {
		return result, &InvalidFieldError{Field: "AnalyticsType", Values: params.AnalyticsType}
	}

	if params.AnalyticsType == "interactions-per-cast" {
		validAggregateWindows := []string{"10s", "1m", "2m", "5m", "10m", "20m", "30m", "2h", "12h", "1d", "7d"}
		if params.AggregateWindow == "" || !Contains(validAggregateWindows, params.AggregateWindow) {
			return result, &InvalidFieldError{Field: "AggregateWindow", Values: params.AggregateWindow}
		}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/validate/analytics"

	values := map[string]any{
		"frame_url":      params.FrameURL,
		"analytics_type": params.AnalyticsType,
		"start":          params.Start.Format(time.RFC3339),
		"stop":           params.Stop.Format(time.RFC3339),
	}

	if params.AggregateWindow != "" {
		values["aggregate_window"] = params.AggregateWindow
	}

	q := GetUrlValues(values)

	rawQuery := q.Encode()
	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type FrameResponse struct {
	NeynarFrame
	ErrorResponse
}

func (s *FrameService) FetchFrame(ctx context.Context, params FetchFrameParams) (FrameResponse, error) {
	var result FrameResponse
	if params.Type == "" {
		return result, &RequiredFieldError{Field: "Type"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
	values := map[string]any{
		"type": params.Type,
		"uuid": params.UUID,
		"url":  params.URL,
	}

	q := GetUrlValues(values)
	rawQuery := q.Encode()
	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

func (s *FrameService) CreateFrame(ctx context.Context, params CreateFrameParams) (FrameResponse, error) {
	var result FrameResponse

	if params.Frame.Name == "" {
		return result, &RequiredFieldError{Field: "Frame.Name"}
	}
	if len(params.Frame.Pages) == 0 {
		return result, &RequiredFieldError{Field: "Frame.Pages"}
	}

	for _, page := range params.Frame.Pages {
		if page.UUID == "" {
			return result, &RequiredFieldError{Field: "Page.UUID"}
		}
		if page.Title == "" {
			return result, &RequiredFieldError{Field: "Page.Title"}
		}
		if page.Image.URL == "" {
			return result, &RequiredFieldError{Field: "Page.Image.URL"}
		}
		if page.Image.AspectRatio == "" {
			return result, &RequiredFieldError{Field: "Page.Image.AspectRatio"}
		}
		if page.Version == "" {
			page.Version = "vNext" // Set default value if not provided
		}

		// Validate Buttons
		for _, button := range page.Buttons {
			if button.Title == "" {
				return result, &RequiredFieldError{Field: "Page.Buttons.Title"}
			}
			if button.Index <= 0 {
				return result, &InvalidFieldError{Field: "Page.Buttons.Index", Values: "Index must be greater than 0"}
			}
			if button.ActionType == "" {
				return result, &RequiredFieldError{Field: "Page.Buttons.ActionType"}
			}
			if button.ActionType != "post" && button.ActionType != "post_redirect" && button.ActionType != "mint" && button.ActionType != "link" {
				return result, &InvalidFieldError{Field: "Page.Buttons.ActionType", Values: "Invalid ActionType"}
			}
			if button.NextPage != nil {
				// Validate NextPage
				if button.NextPage.UUID != "" && button.NextPage.RedirectURL != "" && button.NextPage.MintURL != "" {
					return result, &InvalidFieldError{Field: "Page.Buttons.NextPage", Values: "Only one of UUID, RedirectURL, or MintURL should be set"}
				}
				if button.NextPage.UUID == "" && button.NextPage.RedirectURL == "" && button.NextPage.MintURL == "" {
					return result, &RequiredFieldError{Field: "Page.Buttons.NextPage"}
				}
			}
		}

		// Validate Input
		if page.Input.Text.Enabled && page.Input.Text.Placeholder == "" {
			return result, &RequiredFieldError{Field: "Page.Input.Text.Placeholder"}
		}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
	body := params.Frame

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

func (s *FrameService) UpdateFrame(ctx context.Context, params UpdateFrameParams) (FrameResponse, error) {
	var result FrameResponse

	// Validate Frame
	if params.Frame.UUID == "" {
		return result, &RequiredFieldError{Field: "Frame.UUID"}
	}
	if len(params.Frame.Pages) == 0 {
		return result, &RequiredFieldError{Field: "Frame.Pages"}
	}

	// Validate Pages
	for _, page := range params.Frame.Pages {
		if page.UUID == "" {
			return result, &RequiredFieldError{Field: "Page.UUID"}
		}
		if page.Title == "" {
			return result, &RequiredFieldError{Field: "Page.Title"}
		}
		if page.Image.URL == "" {
			return result, &RequiredFieldError{Field: "Page.Image.URL"}
		}
		if page.Image.AspectRatio == "" {
			return result, &RequiredFieldError{Field: "Page.Image.AspectRatio"}
		}
		if page.Version == "" {
			page.Version = "vNext" // Set default value if not provided
		}

		// Validate Buttons
		for _, button := range page.Buttons {
			if button.Title == "" {
				return result, &RequiredFieldError{Field: "Page.Buttons.Title"}
			}
			if button.Index <= 0 {
				return result, &InvalidFieldError{Field: "Page.Buttons.Index", Values: "Index must be greater than 0"}
			}
			if button.ActionType == "" {
				return result, &RequiredFieldError{Field: "Page.Buttons.ActionType"}
			}
			if button.ActionType != "post" && button.ActionType != "post_redirect" && button.ActionType != "mint" && button.ActionType != "link" {
				return result, &InvalidFieldError{Field: "Page.Buttons.ActionType", Values: "Invalid ActionType"}
			}
			if button.NextPage != nil {
				// Validate NextPage
				if button.NextPage.UUID != "" && button.NextPage.RedirectURL != "" && button.NextPage.MintURL != "" {
					return result, &InvalidFieldError{Field: "Page.Buttons.NextPage", Values: "Only one of UUID, RedirectURL, or MintURL should be set"}
				}
				if button.NextPage.UUID == "" && button.NextPage.RedirectURL == "" && button.NextPage.MintURL == "" {
					return result, &RequiredFieldError{Field: "Page.Buttons.NextPage"}
				}
			}
		}

		// Validate Input
		if page.Input.Text.Enabled && page.Input.Text.Placeholder == "" {
			return result, &RequiredFieldError{Field: "Page.Input.Text.Placeholder"}
		}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
	body := params.Frame

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPut, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type DeleteFrameRequest struct {
	UUID string `json:"uuid"`
}

func (s *FrameService) DeleteFrame(ctx context.Context, params DeleteFrameParams) (DeleteFrameResponse, error) {
	var result DeleteFrameResponse
	if params.UUID == "" {
		return result, &RequiredFieldError{Field: "UUID"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodDelete, baseURL, params, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type Frames = []NeynarFrame
type FrameResponseList struct {
	Frames
	ErrorResponse
}

func (s *FrameService) FetchFrameList(ctx context.Context) (FrameResponseList, error) {
	var result FrameResponseList

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/list"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

type FrameDeveloperManagedActionReqBody struct {
	CastHash        string               `json:"cast_hash"`
	Action          FrameAction          `json:"action"`
	SignaturePacket FrameSignaturePacket `json:"signature_packet"`
}

func (s *FrameService) PostFrameDeveloperManagedAction(ctx context.Context, params FrameDeveloperManagedActionReqBody) (FrameResponse, error) {
	var result FrameResponse
	// !important to revisit this damn logic!!!!!

	if params.CastHash == "" {
		return result, &RequiredFieldError{Field: "RequestBody.CastHash"}
	}
	if params.Action.Button.Title == "" {
		return result, &RequiredFieldError{Field: "RequestBody.Action.Button.Title"}
	}
	if params.Action.Button.Index == 0 {
		return result, &RequiredFieldError{Field: "RequestBody.Action.Button.Index"}
	}
	if params.Action.Button.ActionType == "" {
		return result, &RequiredFieldError{Field: "RequestBody.Action.Button.ActionType"}
	}
	if params.Action.FramesURL == "" {
		return result, &RequiredFieldError{Field: "RequestBody.Action.FramesURL"}
	}
	if params.Action.PostURL == "" {
		return result, &RequiredFieldError{Field: "RequestBody.Action.PostURL"}
	}

	if params.SignaturePacket.UntrustedData.FID == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.FID"}
	}
	if params.SignaturePacket.UntrustedData.URL == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.URL"}
	}
	if params.SignaturePacket.UntrustedData.MessageHash == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.MessageHash"}
	}
	if params.SignaturePacket.UntrustedData.Timestamp == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.Timestamp"}
	}
	if params.SignaturePacket.UntrustedData.Network == 0 {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.Network"}
	}
	if params.SignaturePacket.UntrustedData.ButtonIndex == 0 {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.ButtonIndex"}
	}
	if params.SignaturePacket.UntrustedData.Address == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.Address"}
	}
	if params.SignaturePacket.UntrustedData.CastID.FID == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.CastID.FID"}
	}
	if params.SignaturePacket.UntrustedData.CastID.Hash == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.UntrustedData.CastID.Hash"}
	}
	if params.SignaturePacket.TrustedData.MessageBytes == "" {
		return result, &RequiredFieldError{Field: "RequestBody.SignaturePacket.TrustedData.MessageBytes"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/developer_managed/action"
	body := params

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}
