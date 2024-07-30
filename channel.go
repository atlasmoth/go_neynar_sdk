package neynarsdk

import (
	"context"
	"net/http"
	"strings"
)

type ChannelService struct {
	client *Client
}

type ListChannelsParams struct {
	Limit  int
	Cursor string
}

type ChannelListResponse struct {
	Channels []Channel  `json:"channels"`
	Next     NextCursor `json:"next,omitempty"`
	ErrorResponse
}


func (c *ChannelService) ListChannels(ctx context.Context, params ListChannelsParams) (ChannelListResponse, error) {
	var result ChannelListResponse

	if params.Limit < 1 || params.Limit > 200 {
		return result, &InvalidFieldError{Field: "Limit", Values: "Limit"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/list"

	values := map[string]interface{}{
		"limit":  params.Limit,
		"cursor": params.Cursor,
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

type SearchChannelsParams struct {
	Query string
}

type ChannelSearchResponse struct {
	Channels []Channel `json:"channels"`
	ErrorResponse
}



func (c *ChannelService) SearchChannels(ctx context.Context, params SearchChannelsParams) (ChannelSearchResponse, error) {
	var result ChannelSearchResponse

	if params.Query == "" {
		return result, &RequiredFieldError{Field: "Query"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/search"

	values := map[string]interface{}{
		"q": params.Query,
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

type ChannelBulkParams struct {
	IDs       []string
	Type      string
	ViewerFid uint32
}

func (c *ChannelService) ChannelBulk(ctx context.Context, params ChannelBulkParams) (ChannelResponseBulk, error) {
	var result ChannelResponseBulk

	if len(params.IDs) == 0 {
		return result, &RequiredFieldError{Field: "IDs"}
	}
	if len(params.IDs) > 100 {
		return result, &InvalidFieldError{Field: "IDs", Values: "IDs"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/bulk"

	values := map[string]interface{}{
		"ids":  strings.Join(params.IDs, ","),
		"type": params.Type,
	}
	if values["type"] == ""{
		values["type"] = "id"
	}

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

type ChannelParams struct {
	ID        string
	Type      string
	ViewerFid uint32
}

type ChannelResponse struct {
	Channel Channel `json:"channel"`
	ErrorResponse
}

func (c *ChannelService) ChannelDetails(ctx context.Context, params ChannelParams) (ChannelResponse, error) {
	var result ChannelResponse

	if params.ID == "" {
		return result, &RequiredFieldError{Field: "ID"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel"

	values := map[string]interface{}{
		"id":   params.ID,
		"type": params.Type,
	}

	if values["type"] == ""{
		values["type"] = "id"
	}

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

type ChannelFollowersParams struct {
	ID     string
	Cursor string
	Limit  int
}

type UsersResponse struct {
	Users []User     `json:"users"`
	Next  NextCursor `json:"next"`
	ErrorResponse
}

func (c *ChannelService) ChannelFollowers(ctx context.Context, params ChannelFollowersParams) (UsersResponse, error) {
	var result UsersResponse

	if params.ID == "" {
		return result, &RequiredFieldError{Field: "ID"}
	}
	if params.Limit < 15 || params.Limit > 1000 {
		return result, &InvalidFieldError{Field: "Limit", Values: "must be between 15 and 1000"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/followers"

	values := map[string]interface{}{
		"id":     params.ID,
		"cursor": params.Cursor,
		"limit":  params.Limit,
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

type ActiveChannelsParams struct {
	FID    uint32
	Limit  int
	Cursor string
}

type UsersActiveChannelsResponse struct {
	Channels []Channel  `json:"channels"`
	Next     NextCursor `json:"next,omitempty"`
	ErrorResponse
}

func (c *ChannelService) ActiveChannels(ctx context.Context, params ActiveChannelsParams) (UsersActiveChannelsResponse, error) {
	var result UsersActiveChannelsResponse

	if params.FID == 0 {
		return result, &RequiredFieldError{Field: "FID"}
	}
	if params.Limit < 1 || params.Limit > 100 {
		return result, &InvalidFieldError{Field: "Limit", Values: "must be between 1 and 100"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/user"

	values := map[string]interface{}{
		"fid":    params.FID,
		"limit":  params.Limit,
		"cursor": params.Cursor,
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

type ChannelUsersParams struct {
	ID                 string
	HasRootCastAuthors bool
	HasCastLikers      bool
	HasCastRecasters   bool
	HasReplyAuthors    bool
	Cursor             string
	Limit              int
}


func (c *ChannelService) ChannelUsers(ctx context.Context, params ChannelUsersParams) (UsersResponse, error) {
	var result UsersResponse

	if params.ID == "" {
		return result, &RequiredFieldError{Field: "ID"}
	}
	if params.Limit < 1 || params.Limit > 100 {
		return result, &InvalidFieldError{Field: "Limit", Values: "must be between 1 and 100"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/users"

	values := map[string]interface{}{
		"id":                    params.ID,
		"has_root_cast_authors": params.HasRootCastAuthors,
		"has_cast_likers":       params.HasCastLikers,
		"has_cast_recasters":    params.HasCastRecasters,
		"has_reply_authors":     params.HasReplyAuthors,
		"cursor":                params.Cursor,
		"limit":                 params.Limit,
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

type TrendingChannelsParams struct {
	TimeWindow string
	Limit      int
	Cursor     string
}

type TrendingChannelResponse struct {
	Channels []ChannelActivity `json:"channels"`
	Next     NextCursor        `json:"next"`
	ErrorResponse
}

func (c *ChannelService) TrendingChannels(ctx context.Context, params TrendingChannelsParams) (TrendingChannelResponse, error) {
	var result TrendingChannelResponse

	if params.TimeWindow != "" && params.TimeWindow != "1d" && params.TimeWindow != "7d" && params.TimeWindow != "30d" {
		return result, &InvalidFieldError{Field: "TimeWindow", Values: "must be '1d', '7d', or '30d'"}
	}
	if params.Limit < 1 || params.Limit > 25 {
		return result, &InvalidFieldError{Field: "Limit", Values: "must be between 1 and 25"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/channel/trending"

	values := map[string]interface{}{
		"time_window": params.TimeWindow,
		"limit":       params.Limit,
		"cursor":      params.Cursor,
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

type UserChannelsParams struct {
	FID    uint32
	Limit  int
	Cursor string
}




func (c *ChannelService) UserChannels(ctx context.Context, params UserChannelsParams) (ChannelListResponse, error) {
	var result ChannelListResponse

	if params.FID == 0 {
		return result, &RequiredFieldError{Field: "FID"}
	}
	if params.Limit < 1 || params.Limit > 100 {
		return result, &InvalidFieldError{Field: "Limit", Values: "must be between 1 and 100"}
	}

	baseURL := c.client.BaseURL.String() + "v2/farcaster/user/channels"

	values := map[string]interface{}{
		"fid":    params.FID,
		"limit":  params.Limit,
		"cursor": params.Cursor,
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



type ChannelResponseBulk struct {
	Channels []Channel `json:"channels"`
	ErrorResponse
}






