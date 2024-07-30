package neynarsdk

import (
	"context"
	"net/http"
)

type FollowService struct {
	client *Client
}

type RetrieveFollowersParams struct {
	Fid       uint32
	ViewerFid int32
	SortType  string
	Limit     int
	Cursor    string
}

type Follower struct {
	TopRelevantFollowersHydrated   []Follow `json:"top_relevant_followers_hydrated"`
	AllRelevantFollowersDehydrated []Follow `json:"all_relevant_followers_dehydrated"`
}

type RetrieveFollowersResult struct {
	Users []User `json:"users"`
	ErrorResponse
	Next Next `json:"next"`
}

func (f *FollowService) RetrieveFollowers(ctx context.Context, params RetrieveFollowersParams) (RetrieveFollowersResult, error) {
	var result RetrieveFollowersResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}

	baseURL := f.client.BaseURL.String() + "v2/farcaster/followers"

	values := map[string]interface{}{
		"fid":       params.Fid,
		"sort_type": params.SortType,
		"limit":     params.Limit,
		"cursor":    params.Cursor,
	}

	if params.ViewerFid > 0 {
		values["viewer_fid"] = params.ViewerFid
	}
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

type RetrieveRelevantFollowersParams struct {
	TargetFid uint32
	ViewerFid uint32
}

type RetrieveRelevantFollowersResult struct {
	TopRelevantFollowersHydrated   []Follower `json:"top_relevant_followers_hydrated"`
	AllRelevantFollowersDehydrated []Follower `json:"all_relevant_followers_dehydrated"`
	ErrorResponse
}

func (f *FollowService) RetrieveRelevantFollowers(ctx context.Context, params RetrieveRelevantFollowersParams) (RetrieveRelevantFollowersResult, error) {
	var result RetrieveRelevantFollowersResult
	if params.TargetFid == 0 {
		return result, &RequiredFieldError{Field: "TargetFid"}
	}
	if params.ViewerFid == 0 {
		return result, &RequiredFieldError{Field: "ViewerFid"}
	}

	baseURL := f.client.BaseURL.String() + "v2/farcaster/followers/relevant"

	values := map[string]interface{}{
		"target_fid": params.TargetFid,
		"viewer_fid": params.ViewerFid,
	}

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

type RetrieveFollowingParams struct {
	Fid       int32
	ViewerFid int32
	SortType  string
	Limit     int
	Cursor    string
}

type RetrieveFollowingResult struct {
	Users []User `json:"users"`
	ErrorResponse
	Next Next `json:"next"`
}

func (f *FollowService) RetrieveFollowing(ctx context.Context, params RetrieveFollowingParams) (RetrieveFollowingResult, error) {
	var result RetrieveFollowingResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}

	baseURL := f.client.BaseURL.String() + "v2/farcaster/following"

	values := map[string]interface{}{
		"fid":       params.Fid,
		"sort_type": params.SortType,
		"limit":     params.Limit,
		"cursor":    params.Cursor,
	}

	if params.ViewerFid > 0 {
		values["viewer_fid"] = params.ViewerFid
	}
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
