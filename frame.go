package neynarsdk

import (
	"context"
	"net/http"
	"time"
)

type FrameService struct {
	client *Client
}



type PostFrameActionResponse struct{
	Frame
	ErrorResponse
}

type FrameActionReqBody struct {
    SignerUUID SignerUUID `json:"signer_uuid"`
    CastHash   CastHash   `json:"cast_hash"`
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
	MessageBytesInHex    string `json:"message_bytes_in_hex"`
	OmitCastReactionContext  bool   `json:"cast_reaction_context,omitempty"`
	FollowContext        bool   `json:"follow_context,omitempty"`
	SignerContext        bool   `json:"signer_context,omitempty"`
	ChannelFollowContext bool   `json:"channel_follow_context,omitempty"`
}


type ValidateFrameActionResponse struct {
	Valid  bool                `json:"valid"`
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
    FrameURL         string    `json:"frame_url"`
    AnalyticsType    string    `json:"analytics_type"`
    Start            time.Time `json:"start"`
    Stop             time.Time `json:"stop"`
    AggregateWindow  string    `json:"aggregate_window,omitempty"`
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
        return result, &InvalidFieldError{Field: "AnalyticsType", Values:params.AnalyticsType}
    }

    
    if params.AnalyticsType == "interactions-per-cast" {
        validAggregateWindows := []string{"10s", "1m", "2m", "5m", "10m", "20m", "30m", "2h", "12h", "1d", "7d"}
        if params.AggregateWindow == "" || !Contains(validAggregateWindows, params.AggregateWindow) {
            return result, &InvalidFieldError{Field: "AggregateWindow", Values: params.AggregateWindow}
        }
    }

    baseURL := s.client.BaseURL.String() + "v2/farcaster/frame/validate/analytics"

	values := map[string]any{
		"frame_url": params.FrameURL,
		"analytics_type" : params.AnalyticsType,
		"start" : params.Start.Format(time.RFC3339),
		"stop" : params.Stop.Format(time.RFC3339),
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


