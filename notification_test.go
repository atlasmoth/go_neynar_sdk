package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestRetrieveNotificationsForUser_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/notifications"
	expectedParams := url.Values{
		"fid":        []string{"12345"},
		"cursor":     []string{"test-cursor"},
		"is_priority": []string{"true"},
	}
	mockResponse := RetrieveNotificationsForUserResult{
		Notifications: []Notification{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveNotificationsForUserParams{
		Fid:           12345,
		Cursor:        "test-cursor",
		IsNotPriority: false,
	}
	ctx := context.Background()
	result, err := client.Notification.RetrieveNotificationsForUser(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveNotificationsForUser_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveNotificationsForUserParams{}
	ctx := context.Background()
	_, err := client.Notification.RetrieveNotificationsForUser(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveNotificationsForUser_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/notifications"
	expectedParams := url.Values{
		"fid":        []string{"12345"},
		"cursor":     []string{"test-cursor"},
		"is_priority": []string{"true"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveNotificationsForUserParams{
		Fid:           12345,
		Cursor:        "test-cursor",
		IsNotPriority: false,
	}
	ctx := context.Background()
	_, err := client.Notification.RetrieveNotificationsForUser(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveNotificationsForUser_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/notifications"
	expectedParams := url.Values{
		"fid":        []string{"12345"},
		"cursor":     []string{"test-cursor"},
		"is_priority": []string{"true"},
	}
	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := RetrieveNotificationsForUserParams{
		Fid:           12345,
		Cursor:        "test-cursor",
		IsNotPriority: false,
	}
	ctx := context.Background()
	_, err := client.Notification.RetrieveNotificationsForUser(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveNotificationsForUser_NonPriority(t *testing.T) {
	expectedPath := "/v2/farcaster/notifications"
	expectedParams := url.Values{
		"fid":        []string{"12345"},
		"cursor":     []string{"test-cursor"},
		"is_priority": []string{"false"},
	}
	mockResponse := RetrieveNotificationsForUserResult{
		Notifications: []Notification{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()


	params := RetrieveNotificationsForUserParams{
		Fid:           12345,
		Cursor:        "test-cursor",
		IsNotPriority: true,
	}
	ctx := context.Background()
	result, err := client.Notification.RetrieveNotificationsForUser(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}
