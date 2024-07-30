package neynarsdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestListChannels_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/list"
	expectedParams := url.Values{
		"limit":  []string{"50"},
		"cursor": []string{"test-cursor"},
	}
	mockResponse := ChannelListResponse{
		Channels: []Channel{{ID: "1"}, {ID: "2"}},
		Next:     NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ListChannelsParams{
		Limit:  50,
		Cursor: "test-cursor",
	}
	ctx := context.Background()
	result, err := client.Channel.ListChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestListChannels_InvalidLimit(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := ListChannelsParams{Limit: 0}
	ctx := context.Background()
	_, err := client.Channel.ListChannels(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &InvalidFieldError{Field: "Limit", Values: "Limit"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestSearchChannels_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/search"
	expectedParams := url.Values{
		"q": []string{"test query"},
	}
	mockResponse := ChannelSearchResponse{
		Channels: []Channel{{ID: "1"}, {ID: "2"}},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := SearchChannelsParams{Query: "test query"}
	ctx := context.Background()
	result, err := client.Channel.SearchChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestSearchChannels_EmptyQuery(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := SearchChannelsParams{}
	ctx := context.Background()
	_, err := client.Channel.SearchChannels(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &RequiredFieldError{Field: "Query"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestChannelBulk_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/bulk"
	expectedParams := url.Values{
		"ids":        []string{"1,2,3"},
		"type":       []string{"id"},
		"viewer_fid": []string{"12345"},
	}
	mockResponse := ChannelResponseBulk{
		Channels: []Channel{{ID: "1"}, {ID: "2"}, {ID: "3"}},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ChannelBulkParams{
		IDs:       []string{"1", "2", "3"},
		Type:      "id",
		ViewerFid: 12345,
	}
	ctx := context.Background()
	result, err := client.Channel.ChannelBulk(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestChannelBulk_EmptyIDs(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := ChannelBulkParams{}
	ctx := context.Background()
	_, err := client.Channel.ChannelBulk(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &RequiredFieldError{Field: "IDs"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestChannelDetails_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel"
	expectedParams := url.Values{
		"id":         []string{"123"},
		"type":       []string{"id"},
		"viewer_fid": []string{"12345"},
	}
	mockResponse := ChannelResponse{
		Channel: Channel{ID: "123"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ChannelParams{
		ID:        "123",
		Type:      "id",
		ViewerFid: 12345,
	}
	ctx := context.Background()
	result, err := client.Channel.ChannelDetails(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestChannelDetails_EmptyID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := ChannelParams{}
	ctx := context.Background()
	_, err := client.Channel.ChannelDetails(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &RequiredFieldError{Field: "ID"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestChannelFollowers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/followers"
	expectedParams := url.Values{
		"id":     []string{"123"},
		"cursor": []string{"test-cursor"},
		"limit":  []string{"50"},
	}
	mockResponse := UsersResponse{
		Users: []User{{Fid: 1}, {Fid: 2}},
		Next:  NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ChannelFollowersParams{
		ID:     "123",
		Cursor: "test-cursor",
		Limit:  50,
	}
	ctx := context.Background()
	result, err := client.Channel.ChannelFollowers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestChannelFollowers_InvalidLimit(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := ChannelFollowersParams{
		ID:    "123",
		Limit: 10, // Less than minimum 15
	}
	ctx := context.Background()
	_, err := client.Channel.ChannelFollowers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &InvalidFieldError{Field: "Limit", Values: "must be between 15 and 1000"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestActiveChannels_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/user"
	expectedParams := url.Values{
		"fid":    []string{"12345"},
		"limit":  []string{"50"},
		"cursor": []string{"test-cursor"},
	}
	mockResponse := UsersActiveChannelsResponse{
		Channels: []Channel{{ID: "1"}, {ID: "2"}},
		Next:     NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ActiveChannelsParams{
		FID:    12345,
		Limit:  50,
		Cursor: "test-cursor",
	}
	ctx := context.Background()
	result, err := client.Channel.ActiveChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestActiveChannels_MissingFID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := ActiveChannelsParams{
		Limit: 50,
	}
	ctx := context.Background()
	_, err := client.Channel.ActiveChannels(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &RequiredFieldError{Field: "FID"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestChannelUsers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/users"
	expectedParams := url.Values{
		"id":                    []string{"123"},
		"has_root_cast_authors": []string{"true"},
		"has_cast_likers":       []string{"true"},
		"has_cast_recasters":    []string{"false"},
		"has_reply_authors":     []string{"false"},
		"cursor":                []string{"test-cursor"},
		"limit":                 []string{"50"},
	}
	mockResponse := UsersResponse{
		Users: []User{{Fid: 1}, {Fid: 2}},
		Next:  NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ChannelUsersParams{
		ID:                 "123",
		HasRootCastAuthors: true,
		HasCastLikers:      true,
		HasCastRecasters:   false,
		HasReplyAuthors:    false,
		Cursor:             "test-cursor",
		Limit:              50,
	}
	ctx := context.Background()
	result, err := client.Channel.ChannelUsers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestChannelUsers_MissingID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := ChannelUsersParams{
		Limit: 50,
	}
	ctx := context.Background()
	_, err := client.Channel.ChannelUsers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &RequiredFieldError{Field: "ID"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestTrendingChannels_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/trending"
	expectedParams := url.Values{
		"time_window": []string{"7d"},
		"limit":       []string{"20"},
		"cursor":      []string{"test-cursor"},
	}
	mockResponse := TrendingChannelResponse{
		Channels: []ChannelActivity{{Channel: Channel{ID: "1"}}, {Channel: Channel{ID: "2"}}},
		Next:     NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := TrendingChannelsParams{
		TimeWindow: "7d",
		Limit:      20,
		Cursor:     "test-cursor",
	}
	ctx := context.Background()
	result, err := client.Channel.TrendingChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestTrendingChannels_InvalidTimeWindow(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := TrendingChannelsParams{
		TimeWindow: "2d",
		Limit:      20,
	}
	ctx := context.Background()
	_, err := client.Channel.TrendingChannels(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &InvalidFieldError{Field: "TimeWindow", Values: "must be '1d', '7d', or '30d'"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestUserChannels_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/channels"
	expectedParams := url.Values{
		"fid":    []string{"12345"},
		"limit":  []string{"50"},
		"cursor": []string{"test-cursor"},
	}
	mockResponse := ChannelListResponse{
		Channels: []Channel{{ID: "1"}, {ID: "2"}},
		Next:     NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := UserChannelsParams{
		FID:    12345,
		Limit:  50,
		Cursor: "test-cursor",
	}
	ctx := context.Background()
	result, err := client.Channel.UserChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestUserChannels_MissingFID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := UserChannelsParams{
		Limit: 50,
	}
	ctx := context.Background()
	_, err := client.Channel.UserChannels(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &RequiredFieldError{Field: "FID"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestUserChannels_InvalidLimit(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := UserChannelsParams{
		FID:   12345,
		Limit: 101, // Greater than maximum 100
	}
	ctx := context.Background()
	_, err := client.Channel.UserChannels(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &InvalidFieldError{Field: "Limit", Values: "must be between 1 and 100"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestSearchChannels_EmptyResponse(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/search"
	expectedParams := url.Values{
		"q": []string{"nonexistent channel"},
	}
	mockResponse := ChannelSearchResponse{
		Channels: []Channel{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := SearchChannelsParams{Query: "nonexistent channel"}
	ctx := context.Background()
	result, err := client.Channel.SearchChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(result.Channels) != 0 {
		t.Errorf("Expected empty channels slice, got %v", result.Channels)
	}
}

func TestChannelBulk_MaxIDs(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	ids := make([]string, 101) // Exceeds maximum of 100
	for i := range ids {
		ids[i] = fmt.Sprintf("id%d", i)
	}

	params := ChannelBulkParams{
		IDs: ids,
	}
	ctx := context.Background()
	_, err := client.Channel.ChannelBulk(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := &InvalidFieldError{Field: "IDs", Values: "IDs"}
	if !reflect.DeepEqual(err, expectedError) {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestChannelDetails_DefaultType(t *testing.T) {
	expectedPath := "/v2/farcaster/channel"
	expectedParams := url.Values{
		"id":   []string{"123"},
		"type": []string{"id"}, // Default type should be "id"
	}
	mockResponse := ChannelResponse{
		Channel: Channel{ID: "123"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ChannelParams{
		ID: "123",
		// Type is not set, should default to "id"
	}
	ctx := context.Background()
	result, err := client.Channel.ChannelDetails(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestChannelFollowers_MinimumLimit(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/followers"
	expectedParams := url.Values{
		"id":    []string{"123"},
		"limit": []string{"15"}, // Minimum limit
	}
	mockResponse := UsersResponse{
		Users: []User{{Fid: 1}, {Fid: 2}},
		Next:  NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := ChannelFollowersParams{
		ID:    "123",
		Limit: 15,
	}
	ctx := context.Background()
	result, err := client.Channel.ChannelFollowers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestTrendingChannels_EmptyTimeWindow(t *testing.T) {
	expectedPath := "/v2/farcaster/channel/trending"
	expectedParams := url.Values{
		"limit": []string{"20"},
	}
	mockResponse := TrendingChannelResponse{
		Channels: []ChannelActivity{{Channel: Channel{ID: "1"}}, {Channel: Channel{ID: "2"}}},
		Next:     NextCursor{Cursor: "next-cursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := TrendingChannelsParams{
		Limit: 20,
		// TimeWindow is not set
	}
	ctx := context.Background()
	result, err := client.Channel.TrendingChannels(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}
