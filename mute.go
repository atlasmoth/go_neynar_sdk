package neynarsdk

import (
	"context"
	"net/http"
)

type MuteService struct {
	client *Client
}

type FetchMuteListParams struct {
	Fid    int32
	Limit  int
	Cursor string
}

type AddMuteParams struct {
	Fid      int32
	MutedFid int32
}

type DeleteMuteParams struct {
	Fid      int32
	MutedFid int32
}

type MuteListResponse struct {
	Mutes []MuteList `json:"mutes"`
	Next  string     `json:"next"`
	ErrorResponse
}

type MuteResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	ErrorResponse
}

func (s *MuteService) FetchMuteList(ctx context.Context, params FetchMuteListParams) (MuteListResponse, error) {
	var result MuteListResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/mute/list"
	values := map[string]any{
		"fid": params.Fid,
	}

	if params.Limit > 0 {
		values["limit"] = params.Limit
	}
	if params.Cursor != "" {
		values["cursor"] = params.Cursor
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

func (s *MuteService) AddMute(ctx context.Context, params AddMuteParams) (MuteResponse, error) {
	var result MuteResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.MutedFid == 0 {
		return result, &RequiredFieldError{Field: "MutedFid"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/mute"
	body := map[string]any{
		"fid":       params.Fid,
		"muted_fid": params.MutedFid,
	}

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

func (s *MuteService) DeleteMute(ctx context.Context, params DeleteMuteParams) (MuteResponse, error) {
	var result MuteResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.MutedFid == 0 {
		return result, &RequiredFieldError{Field: "MutedFid"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/mute"
	body := map[string]any{
		"fid":       params.Fid,
		"muted_fid": params.MutedFid,
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
