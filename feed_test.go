package neynarsdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)




func mockHandler(t *testing.T, expectedPath string, expectedParams url.Values, responseBody any, responseCode int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("Expected path %s, got %s", expectedPath, r.URL.Path)
		}
		if r.URL.RawQuery != expectedParams.Encode() {
			t.Errorf("Expected query params %s, got %s", expectedParams.Encode(), r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(responseCode)
		json.NewEncoder(w).Encode(responseBody)
	}
}


func TestRetrieveCastsFromFilters(t *testing.T) {
	expectedPath := "/v2/farcaster/feed"
	expectedParams := url.Values{
		"feed_type":  []string{"filter"},
		"with_recasts": []string{"true"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()
	
	params := RetrieveCastsFromFiltersParams{
		FeedType:    Filter,
		OmitRecasts: false,
	}
	
	ctx := context.Background()
	
	result, err := client.Feed.RetrieveCastsFromFilters(ctx, params)
	
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveCastsFromFilters_MissingFeedType(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveCastsFromFiltersParams{}
	ctx := context.Background()
	_, err := client.Feed.RetrieveCastsFromFilters(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: FeedType"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}


func TestRetrieveFeedFromFollowing(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/following"
	expectedParams := url.Values{
		"with_recasts": []string{"true"},
		"fid":          []string{"1"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveCastsFromFiltersParams{
		Fid:         1,
		OmitRecasts: false,
	}
	ctx := context.Background()
	result, err := client.Feed.RetrieveFeedFromFollowing(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

// Test RetrieveFeedFromFollowing with missing Fid
func TestRetrieveFeedFromFollowing_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveCastsFromFiltersParams{}
	ctx := context.Background()
	_, err := client.Feed.RetrieveFeedFromFollowing(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}


func TestRetrieveFeedForYou(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/for_you"
	expectedParams := url.Values{
		"provider": []string{"karma3"},
		"fid":      []string{"1"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveFeedForYouParams{
		Fid: 1,
	}
	ctx := context.Background()
	result, err := client.Feed.RetrieveFeedForYou(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}


func TestRetrieveFeedForYou_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveFeedForYouParams{}
	ctx := context.Background()
	_, err := client.Feed.RetrieveFeedForYou(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}


func TestRetrieveFeedFromChannelIds(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/channels"
	expectedParams := url.Values{
		"channel_ids": []string{"1"},
		"with_recasts": []string{"true"},
		"with_replies": []string{"true"},
		"should_moderate": []string{"false"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveFeedFromChannelIdsParams{
		ChannelIds:  "1",
		OmitRecasts: false,
		WithReplies: true,
	}
	ctx := context.Background()
	result, err := client.Feed.RetrieveFeedFromChannelIds(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}


func TestRetrieveFeedFromChannelIds_MissingChannelIds(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveFeedFromChannelIdsParams{}
	ctx := context.Background()
	_, err := client.Feed.RetrieveFeedFromChannelIds(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: ChannelIds"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}


func TestRetrieveFeedWithFrames(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/frames"
	expectedParams := url.Values{}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveFeedWithFramesParams{}
	ctx := context.Background()
	result, err := client.Feed.RetrieveFeedWithFrames(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

// Test RetrieveRecentRepliesAndRecasts method
func TestRetrieveRecentRepliesAndRecasts(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/user/replies_and_recasts"
	expectedParams := url.Values{
		"fid": []string{"1"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveRecentRepliesAndRecastsParams{
		Fid: 1,
	}
	ctx := context.Background()
	result, err := client.Feed.RetrieveRecentRepliesAndRecasts(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}


func TestRetrieveRecentRepliesAndRecasts_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveRecentRepliesAndRecastsParams{}
	ctx := context.Background()
	_, err := client.Feed.RetrieveRecentRepliesAndRecasts(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}


func TestRetrieve10PopularCasts(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/user/popular"
	expectedParams := url.Values{
		"fid": []string{"1"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := Retrieve10PopularCastsParams{
		Fid: 1,
	}
	ctx := context.Background()
	result, err := client.Feed.Retrieve10PopularCasts(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}


func TestRetrieve10PopularCasts_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := Retrieve10PopularCastsParams{}
	ctx := context.Background()
	_, err := client.Feed.Retrieve10PopularCasts(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}


func TestRetrieveTrendingCasts(t *testing.T) {
	expectedPath := "/v2/farcaster/feed/trending"
	expectedParams := url.Values{
		"time_window": []string{"24h"},
	}

	mockResponse := RetrieveCastsFromFiltersResult{
		Casts: []Cast{{Text: "test content" }},
		Next:  Next{Cursor: "next_cursor"},
		ErrorResponse: ErrorResponse{},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveTrendingCastsParams{
		TimeWindow: "24h",
	}
	ctx := context.Background()
	result, err := client.Feed.RetrieveTrendingCasts(ctx, params)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

