package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestRetrieveFollowers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/followers"
	expectedParams := url.Values{
		"fid":       []string{"12345"},
		"sort_type": []string{"newest"},
		"limit":     []string{"10"},
		"cursor":    []string{"test-cursor"},
	}
	mockResponse := RetrieveFollowersResult{
		Users: []User{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveFollowersParams{
		Fid:       12345,
		SortType:  "newest",
		Limit:     10,
		Cursor:    "test-cursor",
	}
	ctx := context.Background()
	result, err := client.Follow.RetrieveFollowers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveFollowers_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveFollowersParams{}
	ctx := context.Background()
	_, err := client.Follow.RetrieveFollowers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveFollowers_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/followers"
	expectedParams := url.Values{
		"fid":       []string{"12345"},
		"sort_type": []string{"newest"},
		"limit":     []string{"10"},
		"cursor":    []string{"test-cursor"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveFollowersParams{
		Fid:       12345,
		SortType:  "newest",
		Limit:     10,
		Cursor:    "test-cursor",
	}
	ctx := context.Background()
	_, err := client.Follow.RetrieveFollowers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveRelevantFollowers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/followers/relevant"
	expectedParams := url.Values{
		"target_fid": []string{"12345"},
		"viewer_fid": []string{"67890"},
	}
	mockResponse := RetrieveRelevantFollowersResult{
		TopRelevantFollowersHydrated: []User{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveRelevantFollowersParams{
		TargetFid: 12345,
		ViewerFid: 67890,
	}
	ctx := context.Background()
	result, err := client.Follow.RetrieveRelevantFollowers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result.AllRelevantFollowersDehydrated, mockResponse.AllRelevantFollowersDehydrated) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveRelevantFollowers_MissingTargetFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveRelevantFollowersParams{
		ViewerFid: 67890,
	}
	ctx := context.Background()
	_, err := client.Follow.RetrieveRelevantFollowers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: TargetFid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveRelevantFollowers_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/followers/relevant"
	expectedParams := url.Values{
		"target_fid": []string{"12345"},
		"viewer_fid": []string{"67890"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveRelevantFollowersParams{
		TargetFid: 12345,
		ViewerFid: 67890,
	}
	ctx := context.Background()
	_, err := client.Follow.RetrieveRelevantFollowers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveFollowing_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/following"
	expectedParams := url.Values{
		"fid":       []string{"12345"},
		"sort_type": []string{"newest"},
		"limit":     []string{"10"},
		"cursor":    []string{"test-cursor"},
	}
	mockResponse := RetrieveFollowingResult{
		Users: []User{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveFollowingParams{
		Fid:       12345,
		SortType:  "newest",
		Limit:     10,
		Cursor:    "test-cursor",
	}
	ctx := context.Background()
	result, err := client.Follow.RetrieveFollowing(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveFollowing_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveFollowingParams{}
	ctx := context.Background()
	_, err := client.Follow.RetrieveFollowing(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveFollowing_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/following"
	expectedParams := url.Values{
		"fid":       []string{"12345"},
		"sort_type": []string{"newest"},
		"limit":     []string{"10"},
		"cursor":    []string{"test-cursor"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveFollowingParams{
		Fid:       12345,
		SortType:  "newest",
		Limit:     10,
		Cursor:    "test-cursor",
	}
	ctx := context.Background()
	_, err := client.Follow.RetrieveFollowing(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
