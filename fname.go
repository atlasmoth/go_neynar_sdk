package neynarsdk

import (
	"context"
	"net/http"
)

type FnameService struct {
	client *Client
}

type CheckFnameParams struct{
	Fname string
}

type CheckFnameResult struct{
	Available bool `json:"available"`
	ErrorResponse
}

func (f *FnameService) CheckFname(ctx context.Context, params CheckFnameParams) (CheckFnameResult, error) {
	var result CheckFnameResult
	if params.Fname == "" {
		return result, &RequiredFieldError{Field: "Fname"}
	}

	baseURL := f.client.BaseURL.String() + "v2/farcaster/fname/availability"
	values := map[string]any{"fname": params.Fname}

	q := GetUrlValues(values)
	rawQuery := q.Encode()
	resp, err := f.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = f.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = f.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}