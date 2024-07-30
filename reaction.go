package neynarsdk

import (
	"context"
	"net/http"
)

type ReactionService struct {
	client *Client
}

type ReactionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	ErrorResponse
}

type FetchUserReactionsParams struct {
	Fid       uint32
	ViewerFid uint32
	Type      string
	Limit     uint
	Cursor    string
}

type ReactionsListResponse struct {
	Reactions []ReactionWithCastInfo `json:"reactions"`
	Next      Next                   `json:"next"`
	ErrorResponse
}

func (s *ReactionService) FetchUserReactions(ctx context.Context, params FetchUserReactionsParams) (ReactionsListResponse, error) {
	var result ReactionsListResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.Type == "" {
		return result, &RequiredFieldError{Field: "Type"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/reactions/user"
	values := map[string]any{
		"fid":  params.Fid,
		"type": params.Type,
	}

	if params.ViewerFid != 0 {
		values["viewer_fid"] = params.ViewerFid
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

type FetchCastReactionsParams struct {
	Hash      string
	Types     string
	ViewerFid uint32
	Limit     int
	Cursor    string
}

type PostReactionParams struct {
	SignerUUID      string
	Target          string
	ReactionType    string
	TargetAuthorFid uint32
	Idem            string
}

func (s *ReactionService) FetchCastReactions(ctx context.Context, params FetchCastReactionsParams) (ReactionsListResponse, error) {
	var result ReactionsListResponse
	if params.Hash == "" {
		return result, &RequiredFieldError{Field: "Hash"}
	}
	if params.Types == "" {
		return result, &RequiredFieldError{Field: "Types"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/reactions/cast"
	values := map[string]any{
		"hash":  params.Hash,
		"types": params.Types,
	}

	if params.ViewerFid != 0 {
		values["viewer_fid"] = params.ViewerFid
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

func (s *ReactionService) PostReaction(ctx context.Context, params PostReactionParams) (ReactionResponse, error) {
	var result ReactionResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if params.Target == "" {
		return result, &RequiredFieldError{Field: "Target"}
	}
	if params.ReactionType == "" {
		return result, &RequiredFieldError{Field: "ReactionType"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/reaction"
	body := map[string]any{
		"signer_uuid":   params.SignerUUID,
		"target":        params.Target,
		"reaction_type": params.ReactionType,
	}

	if params.TargetAuthorFid != 0 {
		body["target_author_fid"] = params.TargetAuthorFid
	}
	if params.Idem != "" {
		body["idem"] = params.Idem
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

func (s *ReactionService) DeleteReaction(ctx context.Context, params PostReactionParams) (ReactionResponse, error) {
	var result ReactionResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if params.Target == "" {
		return result, &RequiredFieldError{Field: "Target"}
	}
	if params.ReactionType == "" {
		return result, &RequiredFieldError{Field: "ReactionType"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/reaction"
	body := map[string]any{
		"signer_uuid":   params.SignerUUID,
		"target":        params.Target,
		"reaction_type": params.ReactionType,
	}

	if params.TargetAuthorFid != 0 {
		body["target_author_fid"] = params.TargetAuthorFid
	}
	if params.Idem != "" {
		body["idem"] = params.Idem
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
