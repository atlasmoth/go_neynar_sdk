package neynarsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
}

func (f *FeedService) RetrieveCastsFromFilters(ctx context.Context, params RetrieveCastsFromFiltersParams) (RetrieveCastsFromFiltersResult, error) {
	var result RetrieveCastsFromFiltersResult
    if params.FeedType == "" {
        return result, errors.New("feed type must be set")
    }
    
    baseURL, err := url.JoinPath(f.client.BaseURL.String(), "v2", "farcaster", "feed")
	
    if err != nil {
        return result, err
    }
	u, err := url.Parse(baseURL)
	
    if err != nil {
        return result, err
    }
    
    q := url.Values{}
    q.Set("feed_type", string(params.FeedType))

    q.Set("with_recasts", fmt.Sprintf("%v", !params.OmitRecasts))

    if params.FilterType != "" {
        q.Set("filter_type", string(params.FilterType))
    }
    if params.Fid > 0 {
        q.Set("fid", strconv.Itoa(int(params.Fid)))
    }
    if params.Fids != "" {
        q.Set("fids", params.Fids)
    }
    if params.ParentUrl != "" {
        q.Set("parent_url", params.ParentUrl)
    }
    if params.ChannelId != "" {
        q.Set("channel_id", params.ChannelId)
    }
    if params.EmbedUrl != "" {
        q.Set("embed_url", params.EmbedUrl)
    }
    
    if params.Limit > 0 {
        q.Set("limit", strconv.Itoa(params.Limit))
    }
    if params.Cursor != "" {
        q.Set("cursor", params.Cursor)
    }
    if params.ViewerFid > 0 {
        q.Set("viewer_fid", strconv.Itoa(int(params.ViewerFid)))
    }
    
    u.RawQuery = q.Encode()
	fmt.Println(u.String())
	resp, err := f.client.HandleJsonRequest(ctx, "GET",u.String(), nil)
	
	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		
		err = f.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}else{
		var errorResponse ErrorResponse
		err = f.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
	
	
	
    
    
}