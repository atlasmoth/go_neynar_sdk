package neynarsdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var mockDeveloperManagedSignerResponse = DeveloperManagedSignerResponse{
	PublicKey:         "publicKey123",
	Status:            "active",
	SignerApprovalURL: "https://example.com/approval",
	Fid:               12345,
}

var mockSignerResponse = SignerResponse{
	SignerUUID:        "uuid123",
	PublicKey:         "publicKey123",
	Status:            "active",
	SignerApprovalURL: "https://example.com/approval",
	Fid:               12345,
}

func TestFetchDeveloperManagedSigner_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/signer/developer_managed"
	expectedParams := url.Values{
		"public_key": []string{"publicKey123"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockDeveloperManagedSignerResponse, http.StatusOK))
	defer server.Close()

	params := FetchDeveloperManagedSignerParams{
		PublicKey: "publicKey123",
	}
	ctx := context.Background()
	result, err := client.Signer.FetchDeveloperManagedSigner(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockDeveloperManagedSignerResponse) {
		t.Errorf("Expected result %v, got %v", mockDeveloperManagedSignerResponse, result)
	}
}

func TestFetchDeveloperManagedSigner_MissingPublicKey(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchDeveloperManagedSignerParams{}
	ctx := context.Background()
	_, err := client.Signer.FetchDeveloperManagedSigner(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: PublicKey"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchDeveloperManagedSigner_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/signer/developer_managed"
	expectedParams := url.Values{
		"public_key": []string{"publicKey123"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchDeveloperManagedSignerParams{
		PublicKey: "publicKey123",
	}
	ctx := context.Background()
	_, err := client.Signer.FetchDeveloperManagedSigner(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRegisterDeveloperManagedSignedKey_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/signer/developer_managed/signed_key"

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockDeveloperManagedSignerResponse, http.StatusOK))
	defer server.Close()

	params := RegisterDeveloperManagedSignedKeyParams{
		AppFid:    12345,
		Deadline:  1234567890,
		PublicKey: "publicKey123",
		Signature: "signature123",
	}
	ctx := context.Background()
	result, err := client.Signer.RegisterDeveloperManagedSignedKey(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockDeveloperManagedSignerResponse) {
		t.Errorf("Expected result %v, got %v", mockDeveloperManagedSignerResponse, result)
	}
}

func TestRegisterDeveloperManagedSignedKey_MissingFields(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	testCases := []struct {
		name          string
		params        RegisterDeveloperManagedSignedKeyParams
		expectedError string
	}{
		{"Missing AppFid", RegisterDeveloperManagedSignedKeyParams{Deadline: 1, PublicKey: "pk", Signature: "sig"}, "The following field is required: AppFid"},
		{"Missing Deadline", RegisterDeveloperManagedSignedKeyParams{AppFid: 1, PublicKey: "pk", Signature: "sig"}, "The following field is required: Deadline"},
		{"Missing PublicKey", RegisterDeveloperManagedSignedKeyParams{AppFid: 1, Deadline: 1, Signature: "sig"}, "The following field is required: PublicKey"},
		{"Missing Signature", RegisterDeveloperManagedSignedKeyParams{AppFid: 1, Deadline: 1, PublicKey: "pk"}, "The following field is required: Signature"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := client.Signer.RegisterDeveloperManagedSignedKey(ctx, tc.params)
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if err.Error() != tc.expectedError {
				t.Errorf("Expected error message %v, got %v", tc.expectedError, err.Error())
			}
		})
	}
}

func TestRegisterDeveloperManagedSignedKey_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/signer/developer_managed/signed_key"
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RegisterDeveloperManagedSignedKeyParams{
		AppFid:    12345,
		Deadline:  1234567890,
		PublicKey: "publicKey123",
		Signature: "signature123",
	}
	ctx := context.Background()
	_, err := client.Signer.RegisterDeveloperManagedSignedKey(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchSigner_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/signer"
	expectedParams := url.Values{
		"signer_uuid": []string{"uuid123"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockSignerResponse, http.StatusOK))
	defer server.Close()

	params := FetchSignerParams{
		SignerUUID: "uuid123",
	}
	ctx := context.Background()
	result, err := client.Signer.FetchSigner(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockSignerResponse) {
		t.Errorf("Expected result %v, got %v", mockSignerResponse, result)
	}
}

func TestFetchSigner_MissingSignerUUID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchSignerParams{}
	ctx := context.Background()
	_, err := client.Signer.FetchSigner(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SignerUUID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchSigner_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/signer"
	expectedParams := url.Values{
		"signer_uuid": []string{"uuid123"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchSignerParams{
		SignerUUID: "uuid123",
	}
	ctx := context.Background()
	_, err := client.Signer.FetchSigner(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCreateSigner_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/signer"

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockSignerResponse, http.StatusOK))
	defer server.Close()

	ctx := context.Background()
	result, err := client.Signer.CreateSigner(ctx)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockSignerResponse) {
		t.Errorf("Expected result %v, got %v", mockSignerResponse, result)
	}
}

func TestCreateSigner_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/signer"

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	ctx := context.Background()
	_, err := client.Signer.CreateSigner(ctx)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRegisterSignedKey_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/signer/signed_key"

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockSignerResponse, http.StatusOK))
	defer server.Close()

	params := RegisterSignedKeyParams{
		SignerUUID: "uuid123",
		AppFid:     12345,
		Deadline:   1234567890,
		Signature:  "signature123",
	}
	ctx := context.Background()
	result, err := client.Signer.RegisterSignedKey(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockSignerResponse) {
		t.Errorf("Expected result %v, got %v", mockSignerResponse, result)
	}
}

func TestRegisterSignedKey_MissingFields(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	testCases := []struct {
		name          string
		params        RegisterSignedKeyParams
		expectedError string
	}{
		{"Missing SignerUUID", RegisterSignedKeyParams{AppFid: 1, Deadline: 1, Signature: "sig"}, "The following field is required: SignerUUID"},
		{"Missing AppFid", RegisterSignedKeyParams{SignerUUID: "uuid", Deadline: 1, Signature: "sig"}, "The following field is required: AppFid"},
		{"Missing Deadline", RegisterSignedKeyParams{SignerUUID: "uuid", AppFid: 1, Signature: "sig"}, "The following field is required: Deadline"},
		{"Missing Signature", RegisterSignedKeyParams{SignerUUID: "uuid", AppFid: 1, Deadline: 1}, "The following field is required: Signature"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := client.Signer.RegisterSignedKey(ctx, tc.params)
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if err.Error() != tc.expectedError {
				t.Errorf("Expected error message %v, got %v", tc.expectedError, err.Error())
			}
		})
	}
}

func TestRegisterSignedKey_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/signer/signed_key"
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RegisterSignedKeyParams{
		SignerUUID: "uuid123",
		AppFid:     12345,
		Deadline:   1234567890,
		Signature:  "signature123",
	}
	ctx := context.Background()
	_, err := client.Signer.RegisterSignedKey(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPublishMessage_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/message"
	mockMessage := json.RawMessage(`{"key": "value"}`)
	mockResponse := PublishMessageResponse{}

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := PublishMessageParams{
		Message: mockMessage,
	}
	ctx := context.Background()
	result, err := client.Signer.PublishMessage(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestPublishMessage_MissingMessage(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PublishMessageParams{}
	ctx := context.Background()
	_, err := client.Signer.PublishMessage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Message"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPublishMessage_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/message"
	mockMessage := json.RawMessage(`{"key": "value"}`)

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := PublishMessageParams{
		Message: mockMessage,
	}
	ctx := context.Background()
	_, err := client.Signer.PublishMessage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPublishMessage_InvalidJSON(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	invalidJSON := json.RawMessage(`{"key": "value"`)
	params := PublishMessageParams{
		Message: invalidJSON,
	}
	ctx := context.Background()
	_, err := client.Signer.PublishMessage(ctx, params)
	if err == nil {
		t.Errorf("Expected error for invalid JSON, got nil")
	}
}

func TestPublishMessage_EmptyMessage(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PublishMessageParams{
		Message: json.RawMessage(`{}`),
	}
	ctx := context.Background()
	_, err := client.Signer.PublishMessage(ctx, params)
	if err == nil {
		t.Errorf("Expected error for empty JSON, got nil")
	}
}
