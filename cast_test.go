package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)



func TestRetrieveCastWithHashOrUrl_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"
	expectedParams := url.Values{
		"identifier": []string{"testhash"},
		"type":       []string{"hash"},
		"viewer_fid": []string{"1234"},
	}

	mockResponse := RetrieveCastWithHashOrUrlResult{
		Cast: Cast{Text: "test content"},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveCastWithHashOrUrlParams{
		Identifier: "testhash",
		Type:       "hash",
		ViewerFid:  1234,
	}
	ctx := context.Background()
	result, err := client.Cast.RetrieveCastWithHashOrUrl(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveCastWithHashOrUrl_MissingIdentifier(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveCastWithHashOrUrlParams{
		Type:      "hash",
		ViewerFid: 1234,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastWithHashOrUrl(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Identifier"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveCastWithHashOrUrl_InvalidType(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveCastWithHashOrUrlParams{
		Identifier: "testhash",
		Type:       "invalid",
		ViewerFid:  1234,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastWithHashOrUrl(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The field - Identifier, can only be one of the following :hash, url"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveCastWithHashOrUrl_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"
	expectedParams := url.Values{
		"identifier": []string{"testhash"},
		"type":       []string{"hash"},
		"viewer_fid": []string{"1234"},
	}

	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveCastWithHashOrUrlParams{
		Identifier: "testhash",
		Type:       "hash",
		ViewerFid:  1234,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastWithHashOrUrl(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveCastWithHashOrUrl_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"
	expectedParams := url.Values{
		"identifier": []string{"testhash"},
		"type":       []string{"hash"},
		"viewer_fid": []string{"1234"},
	}

	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := RetrieveCastWithHashOrUrlParams{
		Identifier: "testhash",
		Type:       "hash",
		ViewerFid:  1234,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastWithHashOrUrl(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPostCast_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"

	mockResponse := PostCastResult{
		
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := PostCastParams{
		SignerUUID: "test-signer-uuid",
		
	}
	ctx := context.Background()
	result, err := client.Cast.PostCast(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestPostCast_MissingSignerUUID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostCastParams{
	}
	ctx := context.Background()
	_, err := client.Cast.PostCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SignerUUID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPostCast_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"

	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := PostCastParams{
		SignerUUID: "test-signer-uuid",
		
	}
	ctx := context.Background()
	_, err := client.Cast.PostCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPostCast_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"

	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := PostCastParams{
		SignerUUID: "test-signer-uuid",
		
	}
	ctx := context.Background()
	_, err := client.Cast.PostCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}