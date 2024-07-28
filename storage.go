package neynarsdk

import (
	"context"
	"net/http"
)

type StorageService struct {
	client *Client
}

type FetchStorageAllocationsParams struct {
	Fid int32
}

type FetchStorageUsageParams struct {
	Fid int32
}

type BuyStorageParams struct {
	Fid   int32
	Units int32
	Idem  string
}

type StorageAllocationsResponse struct {
	TotalActiveUnits int              `json:"total_active_units"`
	Allocations      []StorageAllocation `json:"allocations"`
	ErrorResponse
}



func (s *StorageService) FetchStorageAllocations(ctx context.Context, params FetchStorageAllocationsParams) (StorageAllocationsResponse, error) {
	var result StorageAllocationsResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/storage/allocations"
	values := map[string]any{"fid": params.Fid}

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

func (s *StorageService) FetchStorageUsage(ctx context.Context, params FetchStorageUsageParams) (StorageUsageResponse, error) {
	var result StorageUsageResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/storage/usage"
	values := map[string]any{"fid": params.Fid}

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

func (s *StorageService) BuyStorage(ctx context.Context, params BuyStorageParams) (StorageAllocationsResponse, error) {
	var result StorageAllocationsResponse
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.Units == 0 {
		return result, &RequiredFieldError{Field: "Units"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/storage/buy"
	body := map[string]any{"fid": params.Fid, "units": params.Units, "idem": params.Idem}

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
