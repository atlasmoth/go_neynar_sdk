package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestFetchMuteList_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/mute/list"
	expectedParams := url.Values{
		"fid":    []string{"12345"},
		"limit":  []string{"10"},
		"cursor": []string{"someCursor"},
	}
	mockResponse := MuteListResponse{
		Mutes: []MuteList{},
		Next:  "nextCursor",
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := FetchMuteListParams{
		Fid:    12345,
		Limit:  10,
		Cursor: "someCursor",
	}
	ctx := context.Background()
	result, err := client.Mute.FetchMuteList(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestFetchMuteList_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchMuteListParams{}
	ctx := context.Background()
	_, err := client.Mute.FetchMuteList(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchMuteList_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/mute/list"
	expectedParams := url.Values{
		"fid": []string{"12345"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchMuteListParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Mute.FetchMuteList(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestAddMute_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/mute"
	
	mockResponse := MuteResponse{
		Success: true,
		Message: "Mute added successfully",
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := AddMuteParams{
		Fid:      12345,
		MutedFid: 67890,
	}
	ctx := context.Background()
	result, err := client.Mute.AddMute(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestAddMute_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := AddMuteParams{
		MutedFid: 67890,
	}
	ctx := context.Background()
	_, err := client.Mute.AddMute(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestAddMute_MissingMutedFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := AddMuteParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Mute.AddMute(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: MutedFid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestAddMute_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/mute"
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := AddMuteParams{
		Fid:      12345,
		MutedFid: 67890,
	}
	ctx := context.Background()
	_, err := client.Mute.AddMute(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteMute_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/mute"
	mockResponse := MuteResponse{
		Success: true,
		Message: "Mute deleted successfully",
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := DeleteMuteParams{
		Fid:      12345,
		MutedFid: 67890,
	}
	ctx := context.Background()
	result, err := client.Mute.DeleteMute(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestDeleteMute_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := DeleteMuteParams{
		MutedFid: 67890,
	}
	ctx := context.Background()
	_, err := client.Mute.DeleteMute(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteMute_MissingMutedFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := DeleteMuteParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Mute.DeleteMute(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: MutedFid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteMute_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/mute"
	
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := DeleteMuteParams{
		Fid:      12345,
		MutedFid: 67890,
	}
	ctx := context.Background()
	_, err := client.Mute.DeleteMute(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}