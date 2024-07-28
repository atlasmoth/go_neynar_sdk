package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestGetWebhook_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"
	expectedParams := url.Values{
		"webhook_id": []string{"test-webhook-id"},
	}

	mockResponse := WebhookResponse{
		Message: "Success",
		Success: true,
		Webhook: &Webhook{WebhookID: "test-webhook-id"},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := GetWebhookParams{
		WebhookID: "test-webhook-id",
	}
	ctx := context.Background()
	result, err := client.Webhook.GetWebhook(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestGetWebhook_MissingWebhookID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := GetWebhookParams{}
	ctx := context.Background()
	_, err := client.Webhook.GetWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: WebhookID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestGetWebhook_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"
	expectedParams := url.Values{
		"webhook_id": []string{"test-webhook-id"},
	}

	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := GetWebhookParams{
		WebhookID: "test-webhook-id",
	}
	ctx := context.Background()
	_, err := client.Webhook.GetWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestGetWebhook_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"
	expectedParams := url.Values{
		"webhook_id": []string{"test-webhook-id"},
	}

	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := GetWebhookParams{
		WebhookID: "test-webhook-id",
	}
	ctx := context.Background()
	_, err := client.Webhook.GetWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCreateWebhook_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"

	mockResponse := WebhookResponse{
		Message: "Webhook created successfully",
		Success: true,
		Webhook: &Webhook{WebhookID: "test-webhook-id"},
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := WebhookParams{
		Name: "test-webhook",
		URL:  "http://example.com",
	}
	ctx := context.Background()
	result, err := client.Webhook.CreateWebhook(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestCreateWebhook_MissingName(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := WebhookParams{
		URL: "http://example.com",
	}
	ctx := context.Background()
	_, err := client.Webhook.CreateWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: name"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCreateWebhook_MissingURL(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := WebhookParams{
		Name: "test-webhook",
	}
	ctx := context.Background()
	_, err := client.Webhook.CreateWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: url"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCreateWebhook_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"

	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := WebhookParams{
		Name: "test-webhook",
		URL:  "http://example.com",
	}
	ctx := context.Background()
	_, err := client.Webhook.CreateWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestCreateWebhook_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"

	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := WebhookParams{
		Name: "test-webhook",
		URL:  "http://example.com",
	}
	ctx := context.Background()
	_, err := client.Webhook.CreateWebhook(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestUpdateWebhookActiveStatus_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"

	mockResponse := WebhookResponse{
		Message: "Webhook updated successfully",
		Success: true,
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := WebhookParams{
		WebhookID: "test-webhook-id",
		Active:    true,
	}
	ctx := context.Background()
	result, err := client.Webhook.UpdateWebhookActiveStatus(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestUpdateWebhookActiveStatus_MissingWebhookID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := WebhookParams{
		Active: true,
	}
	ctx := context.Background()
	_, err := client.Webhook.UpdateWebhookActiveStatus(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: webhook_id"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestUpdateWebhookActiveStatus_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"

	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := WebhookParams{
		WebhookID: "test-webhook-id",
		Active:    true,
	}
	ctx := context.Background()
	_, err := client.Webhook.UpdateWebhookActiveStatus(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestUpdateWebhookActiveStatus_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/webhook"

	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := WebhookParams{
		WebhookID: "test-webhook-id",
		Active:    true,
	}
	ctx := context.Background()
	_, err := client.Webhook.UpdateWebhookActiveStatus(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
