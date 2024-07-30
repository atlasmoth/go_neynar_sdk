package neynarsdk

import (
	"context"
	"encoding/json"
	"net/http"
)

type SignerService struct {
	client *Client
}

type FetchDeveloperManagedSignerParams struct {
	PublicKey string
}

type RegisterDeveloperManagedSignedKeyParams struct {
	AppFid    uint32
	Deadline  int64
	PublicKey string
	Signature string
}

type FetchSignerParams struct {
	SignerUUID string
}

type RegisterSignedKeyParams struct {
	SignerUUID string
	AppFid     uint32
	Deadline   int64
	Signature  string
}

type DeveloperManagedSignerResponse struct {
	PublicKey         string `json:"public_key"`
	Status            string `json:"status"`
	SignerApprovalURL string `json:"signer_approval_url,omitempty"`
	Fid               uint32 `json:"fid,omitempty"`
	ErrorResponse
}

type SignerResponse struct {
	SignerUUID        string `json:"signer_uuid"`
	PublicKey         string `json:"public_key"`
	Status            string `json:"status"`
	SignerApprovalURL string `json:"signer_approval_url,omitempty"`
	Fid               uint32 `json:"fid,omitempty"`
	ErrorResponse
}

func (s *SignerService) FetchDeveloperManagedSigner(ctx context.Context, params FetchDeveloperManagedSignerParams) (DeveloperManagedSignerResponse, error) {
	var result DeveloperManagedSignerResponse
	if params.PublicKey == "" {
		return result, &RequiredFieldError{Field: "PublicKey"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/signer/developer_managed"
	values := map[string]any{
		"public_key": params.PublicKey,
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

func (s *SignerService) RegisterDeveloperManagedSignedKey(ctx context.Context, params RegisterDeveloperManagedSignedKeyParams) (DeveloperManagedSignerResponse, error) {
	var result DeveloperManagedSignerResponse
	if params.AppFid == 0 {
		return result, &RequiredFieldError{Field: "AppFid"}
	}
	if params.Deadline == 0 {
		return result, &RequiredFieldError{Field: "Deadline"}
	}
	if params.PublicKey == "" {
		return result, &RequiredFieldError{Field: "PublicKey"}
	}
	if params.Signature == "" {
		return result, &RequiredFieldError{Field: "Signature"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/signer/developer_managed/signed_key"
	body := map[string]any{
		"app_fid":    params.AppFid,
		"deadline":   params.Deadline,
		"public_key": params.PublicKey,
		"signature":  params.Signature,
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

func (s *SignerService) FetchSigner(ctx context.Context, params FetchSignerParams) (SignerResponse, error) {
	var result SignerResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/signer"
	values := map[string]any{
		"signer_uuid": params.SignerUUID,
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

func (s *SignerService) CreateSigner(ctx context.Context) (SignerResponse, error) {
	var result SignerResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/signer"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, nil, nil)

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

func (s *SignerService) RegisterSignedKey(ctx context.Context, params RegisterSignedKeyParams) (SignerResponse, error) {
	var result SignerResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if params.AppFid == 0 {
		return result, &RequiredFieldError{Field: "AppFid"}
	}
	if params.Deadline == 0 {
		return result, &RequiredFieldError{Field: "Deadline"}
	}
	if params.Signature == "" {
		return result, &RequiredFieldError{Field: "Signature"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/signer/signed_key"
	body := map[string]any{
		"signer_uuid": params.SignerUUID,
		"app_fid":     params.AppFid,
		"deadline":    params.Deadline,
		"signature":   params.Signature,
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

type PublishMessageParams struct {
	Message json.RawMessage
}

type PublishMessageResponse struct {
	ErrorResponse
}

func (s *SignerService) PublishMessage(ctx context.Context, params PublishMessageParams) (PublishMessageResponse, error) {
	var result PublishMessageResponse

	if len(params.Message) == 0 {

		return result, &RequiredFieldError{Field: "Message"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/message"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, params.Message, nil)

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
