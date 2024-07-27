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
