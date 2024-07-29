package neynarsdk

import (
	"context"
	"net/http"
)

type FrameService struct {
	client *Client
}

type FetchFrameParams struct {
	Type string
	UUID string
	URL  string
}





type FrameListResponse struct {
	Frames []NeynarFrame `json:"frames"`
	ErrorResponse
}

type FrameResponse struct {
	Frame NeynarFrame `json:"frame"`
	ErrorResponse
}

func (s *FrameService) FetchFrame(ctx context.Context, params FetchFrameParams) (FrameResponse, error) {
	var result FrameResponse
	if params.Type == "" {
		return result, &RequiredFieldError{Field: "Type"}
	}
	if params.UUID == "" && params.URL == "" {
		return result, &RequiredFieldError{Field: "UUID or URL"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
	values := map[string]any{
		"type": params.Type,
	}
	if params.UUID != "" {
		values["uuid"] = params.UUID
	}
	if params.URL != "" {
		values["url"] = params.URL
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

func (s *FrameService) CreateFrame(ctx context.Context, body FrameActionReqBody) (FrameResponse, error) {
	var result FrameResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
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

func (s *FrameService) UpdateFrame(ctx context.Context, body FrameActionReqBody) (FrameResponse, error) {
	var result FrameResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
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

func (s *FrameService) DeleteFrame(ctx context.Context, uuid string) (FrameResponse, error) {
	var result FrameResponse
	if uuid == "" {
		return result, &RequiredFieldError{Field: "UUID"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame"
	body := map[string]any{
		"uuid": uuid,
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodDelete, baseURL, body, nil)

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

func (s *FrameService) FetchFrameList(ctx context.Context) (FrameListResponse, error) {
	var result FrameListResponse

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






func (s *FrameService) PostFrameDeveloperManagedAction(ctx context.Context, body FrameDeveloperManagedActionReqBody) (Frame, error) {
	var result Frame
	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/developer_managed/action"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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


func (s *FrameService) ValidateFrame(ctx context.Context, body ValidateFrameReqBody) (ValidateFrameActionResponse, error) {
	var result ValidateFrameActionResponse
	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/validate"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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


func (s *FrameService) ValidateFrameList(ctx context.Context) (FrameValidateListResponse, error) {
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


func (s *FrameService) ValidateFrameAnalytics(ctx context.Context, params ValidateFrameAnalyticsParams) (FrameValidateAnalyticsResponse, error) {
	var result FrameValidateAnalyticsResponse
	if params.FrameURL == "" || params.AnalyticsType == "" || params.Start == "" || params.Stop == "" {
		return result, &RequiredFieldError{Field: "FrameURL, AnalyticsType, Start, Stop"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/validate/analytics"
	values := map[string]any{
		"frame_url":      params.FrameURL,
		"analytics_type": params.AnalyticsType,
		"start":          params.Start,
		"stop":           params.Stop,
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



type ValidateFrameReqBody struct {
	MessageBytesInHex     string `json:"message_bytes_in_hex"`
	CastReactionContext   bool   `json:"cast_reaction_context"`
	FollowContext         bool   `json:"follow_context"`
	SignerContext         bool   `json:"signer_context"`
	ChannelFollowContext  bool   `json:"channel_follow_context"`
}

type ValidateFrameAnalyticsParams struct {
	FrameURL        string `json:"frame_url"`
	AnalyticsType   string `json:"analytics_type"`
	Start           string `json:"start"`
	Stop            string `json:"stop"`
	AggregateWindow string `json:"aggregate_window,omitempty"`
}

type FrameValidateListResponse struct {
	Frames []Frame `json:"frames"`
	ErrorResponse
}

type FrameValidateAnalyticsResponse struct {
	Frames []string `json:"frames"`
	ErrorResponse
}

type ValidateFrameActionResponse struct {
	Valid  bool               `json:"valid"`
	Action ValidatedFrameAction `json:"action"`
	ErrorResponse
}








type Frame struct {
	Version           string            `json:"version"`
	Image             string            `json:"image"`
	Buttons           []FrameActionButton `json:"buttons,omitempty"`
	PostURL           string            `json:"post_url,omitempty"`
	FramesURL         string            `json:"frames_url"`
	Title             string            `json:"title,omitempty"`
	ImageAspectRatio  string            `json:"image_aspect_ratio,omitempty"`
	Input             *FrameInput       `json:"input,omitempty"`
	State             *FrameState       `json:"state,omitempty"`
}

type FrameInput struct {
	Text struct {
		Text string `json:"text"`
	} `json:"text,omitempty"`
}

type FrameState struct {
	Serialized string `json:"serialized"`
}

type FrameDeveloperManagedActionReqBody struct {
	CastHash       CastHash          `json:"cast_hash"`
	Action         FrameAction       `json:"action"`
	SignaturePacket FrameSignaturePacket `json:"signature_packet"`
}

type FrameActionReqBody struct {
	SignerUUID string      `json:"signer_uuid"`
	CastHash   CastHash    `json:"cast_hash"`
	Action     FrameAction `json:"action"`
}



type FrameType string

const (
	FrameTypeUUID FrameType = "uuid"
	FrameTypeURL  FrameType = "url"
)



type FrameValidateAnalyticsInteractors struct {
	Interactors []FrameValidateAnalyticsInteractor `json:"interactors"`
}

type FrameValidateAnalyticsInteractor struct {
	Fid              uint32    `json:"fid"`
	Username         string `json:"username"`
	InteractionCount int    `json:"interaction_count"`
}

type FrameValidateAnalyticsTotalInteractors struct {
	TotalInteractors int `json:"total_interactors"`
}

type FrameValidateAnalyticsInteractionsPerCast struct {
	InteractionsPerCast []FrameValidateAnalyticsInteraction `json:"interactions_per_cast"`
}

type FrameValidateAnalyticsInteraction struct {
	Start            string  `json:"start"`
	Stop             string  `json:"stop"`
	Time             string  `json:"time"`
	InteractionCount int     `json:"interaction_count"`
	CastURL          string  `json:"cast_url"`
}

type FrameValidateAnalyticsInputText struct {
	InputTexts []FrameValidateAnalyticsInputTextEntry `json:"input_texts"`
}

type FrameValidateAnalyticsInputTextEntry struct {
	Fid       uint32    `json:"fid"`
	Username  string `json:"username"`
	InputText string `json:"input_text"`
}



type ValidateFrameAnalyticsType string

const (
	ValidateFrameAnalyticsTypeTotalInteractors   ValidateFrameAnalyticsType = "total-interactors"
	ValidateFrameAnalyticsTypeInteractors        ValidateFrameAnalyticsType = "interactors"
	ValidateFrameAnalyticsTypeInteractionsPerCast ValidateFrameAnalyticsType = "interactions-per-cast"
	ValidateFrameAnalyticsTypeInputText          ValidateFrameAnalyticsType = "input-text"
)

type SignerUUID string

type DeleteFrameResponse struct {
	Success bool   `json:"success"`
	UUID    string `json:"uuid"`
}

type NeynarFrame struct {
	UUID   string           `json:"uuid"`
	Name   string           `json:"name"`
	Link   string           `json:"link"`
	Pages  []NeynarFramePage `json:"pages"`
	Valid  bool             `json:"valid"`
}

type NeynarFramePage struct {
	UUID    string             `json:"uuid"`
	Version string             `json:"version"`
	Title   string             `json:"title"`
	Image   NeynarPageImage    `json:"image"`
	Buttons []NeynarPageButton `json:"buttons,omitempty"`
	Input   NeynarPageInput    `json:"input,omitempty"`
}

type NeynarPageImage struct {
	URL         string `json:"url"`
	AspectRatio string `json:"aspect_ratio"`
}

type NeynarPageButton struct {
	Title      string                `json:"title"`
	Index      int                   `json:"index"`
	ActionType string                `json:"action_type"`
	NextPage   interface{} `json:"next_page,omitempty"`
}

type NeynarPageInput struct {
	Text struct {
		Enabled     bool   `json:"enabled"`
		Placeholder string `json:"placeholder,omitempty"`
	} `json:"text,omitempty"`
}

type NeynarNextFramePage struct {
	UUID string `json:"uuid"`
}

type NeynarNextFramePageRedirect struct {
	RedirectURL string `json:"redirect_url"`
}

type NeynarNextFramePageMintUrl struct {
	MintURL string `json:"mint_url"`
}

type NeynarFrameCreationRequest struct {
	Name  string           `json:"name"`
	Pages []NeynarFramePage `json:"pages"`
}

type NeynarFrameUpdateRequest struct {
	UUID  string           `json:"uuid"`
	Name  string           `json:"name,omitempty"`
	Pages []NeynarFramePage `json:"pages"`
}

type ValidatedFrameAction struct {
	Cast           CastWithInteractions `json:"cast"`
	Interactor     User                 `json:"interactor"`
	Timestamp      Timestamp            `json:"timestamp"`
	TappedButton   FrameTappedButton    `json:"tapped_button"`
	URL            string               `json:"url"`
	Object         string               `json:"object"`
	State          FrameState           `json:"state"`
	Input          FrameInput           `json:"input,omitempty"`
	Signer         ValidatedFrameSigner `json:"signer,omitempty"`
	Transaction    FrameTransaction     `json:"transaction,omitempty"`
	Address        FrameAddress         `json:"address,omitempty"`
}

type ValidatedFrameSigner struct {
	Client User `json:"client"`
}


type CastHash struct {}
type FrameAction struct {}
type FrameSignaturePacket struct {}
type CastWithInteractions struct {}
type Timestamp struct {}
type FrameTransaction struct {}
type FrameAddress struct {}
type FrameTappedButton struct {
	Index int `json:"index"`
}

type FrameActionButton struct {
	Title      string              `json:"title,omitempty"`
	Index      int                 `json:"index"`
	ActionType FrameButtonActionType `json:"action_type"`
	Target     string              `json:"target,omitempty"`
	PostURL    string              `json:"post_url,omitempty"`
}