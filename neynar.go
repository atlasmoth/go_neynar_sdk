package neynarsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type FilterType string
type FeedType string

const (
	libraryVersion = "0.0.1"
	defaultBaseURL = "https://api.neynar.com/"
	mediaType      = "application/json"
)

const (
	Fid            FilterType = "fids"
	ParentUrl      FilterType = "parent_url"
	ChannelId      FilterType = "channel_id"
	EmbedURL       FilterType = "embed_url"
	GlobalTrending FilterType = "global_trending"
	Following      FeedType   = "following"
	Filter         FeedType   = "filter"
)

type ErrorResponse struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Property string `json:"property"`
	Status   int    `json:"status"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("json error code: %v message: %v property: %v status: %v", e.Code, e.Message, e.Property, e.Status)
}

type RequiredFieldError struct {
	Field string
}

func (e *RequiredFieldError) Error() string {
	return fmt.Sprintf("The following field is required: %v", e.Field)
}

type Client struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
	ApiKey     *string
	Feed       FeedService
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{HTTPClient: httpClient, BaseURL: baseURL, ApiKey: &apiKey}
	c.Feed = FeedService{client: c}
	return c
}

func (c *Client) HandleJsonRequest(ctx context.Context, method string, url string, body io.Reader, queryParams *string) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	if queryParams != nil {
		req.URL.RawQuery = *queryParams
	}
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("api_key", *c.ApiKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil

}

func (c *Client) HandleJsonResponse(r *http.Response, payload any) error {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &payload)
	if err != nil {
		return err
	}
	return nil

}

func GetUrlValues(obj map[string]any) url.Values {
	q := url.Values{}
	for i, v := range obj {
		toString := fmt.Sprintf("%v", v)
		if toString != "" {
			q.Set(i, toString)
		}
	}
	return q
}
