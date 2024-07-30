package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestCheckFname_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/fname/availability"
	expectedParams := url.Values{
		"fname": []string{"testname"},
	}
	mockResponse := CheckFnameResult{
		Available: true,
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := CheckFnameParams{
		Fname: "testname",
	}
	ctx := context.Background()
	result, err := client.Fname.CheckFname(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

// TestCheckFname_MissingFname tests the CheckFname method when the fname is missing.
func TestCheckFname_MissingFname(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := CheckFnameParams{
		Fname: "",
	}
	ctx := context.Background()
	_, err := client.Fname.CheckFname(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fname"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCheckFname_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/fname/availability"
	expectedParams := url.Values{
		"fname": []string{"testname"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := CheckFnameParams{
		Fname: "testname",
	}
	ctx := context.Background()
	_, err := client.Fname.CheckFname(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCheckFname_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/fname/availability"
	expectedParams := url.Values{
		"fname": []string{"testname"},
	}
	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := CheckFnameParams{
		Fname: "testname",
	}
	ctx := context.Background()
	_, err := client.Fname.CheckFname(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCheckFname_InvalidJsonResponse(t *testing.T) {
	expectedPath := "/v2/farcaster/fname/availability"
	expectedParams := url.Values{
		"fname": []string{"testname"},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, `{"invalid_json"`, http.StatusOK))
	defer server.Close()

	params := CheckFnameParams{
		Fname: "testname",
	}
	ctx := context.Background()
	_, err := client.Fname.CheckFname(ctx, params)
	if err == nil {
		t.Errorf("Expected JSON parsing error, got nil")
	}
}

func TestCheckFname_InvalidUrl(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	client.BaseURL = &url.URL{Scheme: "http", Host: "invalid-url"}

	params := CheckFnameParams{
		Fname: "testname",
	}
	ctx := context.Background()
	_, err := client.Fname.CheckFname(ctx, params)
	if err == nil {
		t.Errorf("Expected error due to invalid URL, got nil")
	}
}
