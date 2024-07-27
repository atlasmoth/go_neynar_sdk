package neynarsdk

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func NewTestClient(handler http.Handler) (*Client, *httptest.Server) {
	server := httptest.NewServer(handler)
	baseURL, _ := url.Parse(server.URL + "/")
	apiKey := "testApiKey"
	client,_ := NewClient(server.Client(),apiKey)
	client.BaseURL = baseURL
	client.Feed = FeedService{client: client}
	client.Cast = CastService{client: client}
	client.Notification = NotificationService{client: client}
	return client, server
}

func TestNewClient(t *testing.T) {
	apiKey := "testApiKey"
	client,err := NewClient(nil, apiKey)
	if err != nil{
		t.Errorf("Unable to create client")
	}
	if client.HTTPClient == nil {
		t.Errorf("Expected default HTTP client, got nil")
	}
	if client.ApiKey == nil || *client.ApiKey != apiKey {
		t.Errorf("Expected apiKey to be %s, got %s", apiKey, *client.ApiKey)
	}
	expectedBaseURL, _ := url.Parse(defaultBaseURL)
	if client.BaseURL.String() != expectedBaseURL.String() {
		t.Errorf("Expected baseURL to be %s, got %s", expectedBaseURL, client.BaseURL)
	}
}

func TestHandleJsonRequest(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("api_key") != "testApiKey" {
			t.Errorf("Expected api_key to be testApiKey, got %s", r.Header.Get("api_key"))
		}
		w.WriteHeader(http.StatusOK)
	})
	client, server := NewTestClient(handler)
	defer server.Close()

	ctx := context.Background()
	url := client.BaseURL.String()
	resp, err := client.HandleJsonRequest(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestHandleJsonResponse(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"key": "value"})
	})
	client, server := NewTestClient(handler)
	defer server.Close()

	ctx := context.Background()
	url := client.BaseURL.String()
	resp, err := client.HandleJsonRequest(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	var payload map[string]string
	err = client.HandleJsonResponse(resp, &payload)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedPayload := map[string]string{"key": "value"}
	if !reflect.DeepEqual(payload, expectedPayload) {
		t.Errorf("Expected payload %v, got %v", expectedPayload, payload)
	}
}

func TestGetUrlValues(t *testing.T) {
	params := map[string]any{
		"param1": "value1",
		"param2": 123,
		"param3": true,
	}
	expectedValues := url.Values{
		"param1": []string{"value1"},
		"param2": []string{"123"},
		"param3": []string{"true"},
	}
	values := GetUrlValues(params)
	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Expected %v, got %v", expectedValues, values)
	}
}

func TestErrorResponse(t *testing.T) {
	errResp := ErrorResponse{
		Code:     "123",
		Message:  "Test error",
		Property: "test_property",
		Status:   400,
	}
	expectedError := "json error code: 123 message: Test error property: test_property status: 400"
	if errResp.Error() != expectedError {
		t.Errorf("Expected %v, got %v", expectedError, errResp.Error())
	}
}

func TestRequiredFieldError(t *testing.T) {
	errResp := RequiredFieldError{
		Field: "test_field",
	}
	expectedError := "The following field is required: test_field"
	if errResp.Error() != expectedError {
		t.Errorf("Expected %v, got %v", expectedError, errResp.Error())
	}
}

func TestHandleJsonRequest_WithQueryParams(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "param1=value1&param2=value2" {
			t.Errorf("Expected query params param1=value1&param2=value2, got %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
	})
	client, server := NewTestClient(handler)
	defer server.Close()

	ctx := context.Background()
	url := client.BaseURL.String()
	queryParams := "param1=value1&param2=value2"
	resp, err := client.HandleJsonRequest(ctx, http.MethodGet, url, nil, &queryParams)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}

func TestHandleJsonResponse_Error(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{
			Code:     "400",
			Message:  "Bad Request",
			Property: "test_property",
			Status:   400,
		})
	})
	client, server := NewTestClient(handler)
	defer server.Close()

	ctx := context.Background()
	url := client.BaseURL.String()
	resp, err := client.HandleJsonRequest(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	var errResp ErrorResponse
	err = client.HandleJsonResponse(resp, &errResp)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	expectedError := "json error code: 400 message: Bad Request property: test_property status: 400"
	if errResp.Error() != expectedError {
		t.Errorf("Expected %v, got %v", expectedError, errResp.Error())
	}
}
