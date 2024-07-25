package neynarsdk

import (
	"context"
	"net/http"
)

type FeedService struct {
	client *Client
}

type RetrieveCastsFromFiltersParams struct {
	FeedType    FeedType
	FilterType  FilterType
	Fid         int32
	Fids        string
	ParentUrl   string
	ChannelId   string
	EmbedUrl    string
	OmitRecasts bool
	Limit       int
	Cursor      string
	ViewerFid   int32
}

type RetrieveCastsFromFiltersResult struct {
	Casts []Cast `json:"casts"`
	Next  Next   `json:"next"`
	ErrorResponse
}

func (f *FeedService) RetrieveCastsFromFilters(ctx context.Context, params RetrieveCastsFromFiltersParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
	if params.FeedType == "" {
		return result, &RequiredFieldError{Field: "FeedType"}
	}

	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed"

	values := map[string]any{"feed_type": params.FeedType, "with_recasts": !params.OmitRecasts, "filter_type": params.FilterType, "cursor": params.Cursor, "embed_url": params.EmbedUrl, "channel_id": params.ChannelId, "parent_url": params.ParentUrl, "fids": params.Fids}

	if params.Fid > 0 {
		values["fid"] = params.Fid
	}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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

type RetrieveFeedFromFollowingParams struct {
	Fid         int32
	OmitRecasts bool
	Limit       int
	Cursor      string
	ViewerFid   int32
}

func (f *FeedService) RetrieveFeedFromFollowing(ctx context.Context, params RetrieveCastsFromFiltersParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
	if params.Fid < 1 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/following"

	values := map[string]any{"with_recasts": !params.OmitRecasts, "cursor": params.Cursor, "fid": params.Fid}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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

type RetrieveFeedForYouParams struct {
	Fid       int32
	Provider  string
	Limit     int
	Cursor    string
	ViewerFid int32
}

func (f *FeedService) RetrieveFeedForYou(ctx context.Context, params RetrieveFeedForYouParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
	if params.Fid < 1 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/for_you"
	if params.Provider == "" {
		params.Provider = "karma3"
	}

	values := map[string]any{"provider": params.Provider, "cursor": params.Cursor, "fid": params.Fid}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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

type RetrieveFeedFromChannelIdsParams struct {
	ChannelIds     string
	Limit          int
	Cursor         string
	ViewerFid      int32
	OmitRecasts    bool
	WithReplies    bool
	ShouldModerate bool
}

func (f *FeedService) RetrieveFeedFromChannelIds(ctx context.Context, params RetrieveFeedFromChannelIdsParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
	if params.ChannelIds == "" {
		return result, &RequiredFieldError{Field: "ChannelIds"}
	}
	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/channels"

	values := map[string]any{"with_recasts": !params.OmitRecasts, "with_replies": params.WithReplies, "cursor": params.Cursor, "channel_ids": params.ChannelIds, "should_moderate": params.ShouldModerate}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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

type RetrieveFeedWithFramesParams struct {
	Limit     int
	Cursor    string
	ViewerFid int32
}

func (f *FeedService) RetrieveFeedWithFrames(ctx context.Context, params RetrieveFeedWithFramesParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult

	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/frames"

	values := map[string]any{"cursor": params.Cursor}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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

type RetrieveRecentRepliesAndRecastsParams struct {
	Fid       int32
	Limit     int
	Cursor    string
	ViewerFid int32
}

func (f *FeedService) RetrieveRecentRepliesAndRecasts(ctx context.Context, params RetrieveRecentRepliesAndRecastsParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
	if params.Fid < 1 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/user/replies_and_recasts"

	values := map[string]any{"cursor": params.Cursor, "fid": params.Fid}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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

type Retrieve10PopularCastsParams struct {
	Fid       int32
	ViewerFid int32
}

func (f *FeedService) Retrieve10PopularCasts(ctx context.Context, params Retrieve10PopularCastsParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
	if params.Fid < 1 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/user/popular"

	values := map[string]any{"fid": params.Fid}

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

type RetrieveTrendingCastsParams struct {
	ChannelId  string
	Limit      int
	Cursor     string
	ViewerFid  int32
	TimeWindow string
}

func (f *FeedService) RetrieveTrendingCasts(ctx context.Context, params RetrieveTrendingCastsParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult

	baseURL := f.client.BaseURL.String() + "v2/farcaster/feed/trending"

	values := map[string]any{"cursor": params.Cursor, "channel_id": params.ChannelId, "time_window": params.TimeWindow}

	if params.Limit > 0 {
		values["limit"] = params.Limit
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
