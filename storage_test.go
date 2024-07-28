package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestFetchStorageAllocations_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/storage/allocations"
	expectedParams := url.Values{
		"fid": []string{"12345"},
	}
	mockResponse := StorageAllocationsResponse{
		TotalActiveUnits: 10,
		Allocations:      []StorageAllocation{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := FetchStorageAllocationsParams{
		Fid: 12345,
	}
	ctx := context.Background()
	result, err := client.Storage.FetchStorageAllocations(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestFetchStorageAllocations_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchStorageAllocationsParams{}
	ctx := context.Background()
	_, err := client.Storage.FetchStorageAllocations(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchStorageAllocations_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/storage/allocations"
	expectedParams := url.Values{
		"fid": []string{"12345"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchStorageAllocationsParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Storage.FetchStorageAllocations(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchStorageUsage_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/storage/usage"
	expectedParams := url.Values{
		"fid": []string{"12345"},
	}
	mockResponse := StorageUsageResponse{
		TotalActiveUnits: 5,
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := FetchStorageUsageParams{
		Fid: 12345,
	}
	ctx := context.Background()
	result, err := client.Storage.FetchStorageUsage(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestFetchStorageUsage_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchStorageUsageParams{}
	ctx := context.Background()
	_, err := client.Storage.FetchStorageUsage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchStorageUsage_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/storage/usage"
	expectedParams := url.Values{
		"fid": []string{"12345"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchStorageUsageParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Storage.FetchStorageUsage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBuyStorage_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/storage/buy"
	expectedParams := url.Values{}
	mockResponse := StorageAllocationsResponse{
		TotalActiveUnits: 20,
		Allocations:      []StorageAllocation{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := BuyStorageParams{
		Fid:   12345,
		Units: 10,
		Idem:  "uniqueIdemKey",
	}
	ctx := context.Background()
	result, err := client.Storage.BuyStorage(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestBuyStorage_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := BuyStorageParams{
		Units: 10,
		Idem:  "uniqueIdemKey",
	}
	ctx := context.Background()
	_, err := client.Storage.BuyStorage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBuyStorage_MissingUnits(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := BuyStorageParams{
		Fid:  12345,
		Idem: "uniqueIdemKey",
	}
	ctx := context.Background()
	_, err := client.Storage.BuyStorage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Units"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBuyStorage_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/storage/buy"
	expectedParams := url.Values{}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := BuyStorageParams{
		Fid:   12345,
		Units: 10,
		Idem:  "uniqueIdemKey",
	}
	ctx := context.Background()
	_, err := client.Storage.BuyStorage(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
