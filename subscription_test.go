package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestRetrieveSubscriptionsCreated_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/subscriptions_created"
	expectedParams := url.Values{
		"fid":                  []string{"12345"},
		"subscription_provider": []string{"provider"},
	}
	mockResponse := RetrieveSubscriptionsCreatedResult{
		SubscriptionsCreated: []SubscriptionsCreated{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveSubscriptionsCreatedParams{
		Fid:                  12345,
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	result, err := client.Subscription.RetrieveSubscriptionsCreated(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveSubscriptionsCreated_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveSubscriptionsCreatedParams{
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscriptionsCreated(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscriptionsCreated_MissingProvider(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveSubscriptionsCreatedParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscriptionsCreated(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SubscriptionProvider"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscriptionsCreated_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/subscriptions_created"
	expectedParams := url.Values{
		"fid":                  []string{"12345"},
		"subscription_provider": []string{"provider"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveSubscriptionsCreatedParams{
		Fid:                  12345,
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscriptionsCreated(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscribedTo_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/subscribed_to"
	expectedParams := url.Values{
		"fid":                  []string{"12345"},
		"subscription_provider": []string{"provider"},
		"viewer_fid":           []string{"54321"},
	}
	mockResponse := RetrieveSubscribedToResult{
		SubscribedTo: []Subscription{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveSubscribedToParams{
		Fid:                  12345,
		ViewerFid:            54321,
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	result, err := client.Subscription.RetrieveSubscribedTo(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveSubscribedTo_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveSubscribedToParams{
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscribedTo(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscribedTo_MissingProvider(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveSubscribedToParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscribedTo(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SubscriptionProvider"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscribedTo_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/subscribed_to"
	expectedParams := url.Values{
		"fid":                  []string{"12345"},
		"subscription_provider": []string{"provider"},
		"viewer_fid":           []string{"54321"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveSubscribedToParams{
		Fid:                  12345,
		ViewerFid:            54321,
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscribedTo(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscribers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/subscribers"
	expectedParams := url.Values{
		"fid":                  []string{"12345"},
		"subscription_provider": []string{"provider"},
		"viewer_fid":           []string{"54321"},
	}
	mockResponse := RetrieveSubscribersResult{
		Subscribers: []Subscriber{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveSubscribersParams{
		Fid:                  12345,
		ViewerFid:            54321,
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	result, err := client.Subscription.RetrieveSubscribers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveSubscribers_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveSubscribersParams{
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscribers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscribers_MissingProvider(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveSubscribersParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscribers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SubscriptionProvider"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveSubscribers_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/subscribers"
	expectedParams := url.Values{
		"fid":                  []string{"12345"},
		"subscription_provider": []string{"provider"},
		"viewer_fid":           []string{"54321"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveSubscribersParams{
		Fid:                  12345,
		ViewerFid:            54321,
		SubscriptionProvider: "provider",
	}
	ctx := context.Background()
	_, err := client.Subscription.RetrieveSubscribers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
