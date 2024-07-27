package neynarsdk

import (
	"context"
	"net/http"
)

type CastService struct {
	client *Client
}

type RetrieveCastWithHashOrUrlParams struct {
	Identifier string
	Type       string
	ViewerFid  int32
}

type RetrieveCastWithHashOrUrlResult struct {
	Cast Cast `json:"cast"`
	ErrorResponse
}

func (c *CastService) RetrieveCastWithHashOrUrl(ctx context.Context, params RetrieveCastWithHashOrUrlParams) (RetrieveCastWithHashOrUrlResult, error) {
	var result RetrieveCastWithHashOrUrlResult
	if params.Identifier == "" {
		return result, &RequiredFieldError{Field: "Identifier"}
	}
	if params.Type != "hash" && params.Type != "url" {
		return result, &InvalidFieldError{Field: "Identifier", Values: "hash, url"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/cast"

	values := map[string]any{"identifier": params.Identifier, "type": params.Type}

	if params.ViewerFid > 0 {
		values["viewer_fid"] = params.ViewerFid
	}

	q := GetUrlValues(values)

	rawQuery := q.Encode()
	resp, err := c.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {

		err = c.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = c.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}

}

type PostCastResult struct {
	Cast Cast `json:"cast"`
	ErrorResponse
	Success bool `json:"success"`
}

type PostCastMetadata struct {
	ContentType   string `json:"content_type"`
	ContentLength uint32 `json:"content_length"`
}

type PostCastCastId struct {
	FID  uint32 `json:"fid"`
	Hash string `json:"hash"`
}

type PostCastEmbed struct {
	Metadata PostCastMetadata `json:"metadata,omitempty"`
	URL      string           `json:"url,omitempty"`
	CastId   PostCastCastId   `json:"cast_id,omitempty"`
}

type PostCastParams struct {
	Embeds          []PostCastEmbed `json:"embeds"`
	SignerUUID      string          `json:"signer_uuid"`
	Text            string          `json:"text"`
	Parent          string          `json:"parent,omitempty"`
	ChannelID       string          `json:"channel_id,omitempty"`
	Idem            string          `json:"idem,omitempty"`
	ParentAuthorFID int             `json:"parent_author_fid,omitempty"`
}

func (c *CastService) PostCast(ctx context.Context, params PostCastParams) (PostCastResult, error) {
	var result PostCastResult
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/cast"

	resp, err := c.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, params, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {

		err = c.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = c.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}

}

type DeleteCastResult struct {
	Message string `json:"message"`
	ErrorResponse
	Success bool `json:"success"`
}

type DeleteCastParams struct {
	SignerUUID string `json:"signer_uuid"`
	TargetHash string `json:"target_hash"`
}

func (c *CastService) DeleteCast(ctx context.Context, params DeleteCastParams) (DeleteCastResult, error) {
	var result DeleteCastResult
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}

	if params.TargetHash == "" {
		return result, &RequiredFieldError{Field: "TargetHash"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/cast"

	resp, err := c.client.HandleJsonRequest(ctx, http.MethodDelete, baseURL, params, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {

		err = c.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = c.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}

}