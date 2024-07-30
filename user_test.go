package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestSearchUsers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/search"
	expectedParams := url.Values{
		"q":          []string{"testQuery"},
		"viewer_fid": []string{"12345"},
		"limit":      []string{"10"},
		"cursor":     []string{"someCursor"},
	}
	mockResponse := UserSearchResponse{
		Result: struct {
			Users []SearchedUser `json:"users"`
			Next  NextCursor     `json:"next,omitempty"`
		}{
			Users: []SearchedUser{{User: User{}}},
			Next:  NextCursor{Cursor: "nextCursor"},
		},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := SearchUserParams{
		Query:     "testQuery",
		ViewerFid: 12345,
		Limit:     10,
		Cursor:    "someCursor",
	}
	ctx := context.Background()
	result, err := client.User.SearchUsers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestSearchUsers_MissingQuery(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := SearchUserParams{}
	ctx := context.Background()
	_, err := client.User.SearchUsers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Query"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestSearchUsers_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/search"
	expectedParams := url.Values{
		"q": []string{"testQuery"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := SearchUserParams{
		Query: "testQuery",
	}
	ctx := context.Background()
	_, err := client.User.SearchUsers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBulkUsers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/bulk"
	expectedParams := url.Values{
		"fids":       []string{"12345,67890"},
		"viewer_fid": []string{"12345"},
	}
	mockResponse := BulkUsersResponse{
		Users: []User{{Fid: 12345, Username: "testUser1"}, {Fid: 67890, Username: "testUser2"}},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := BulkUserParams{
		Fids:      "12345,67890",
		ViewerFid: 12345,
	}
	ctx := context.Background()
	result, err := client.User.BulkUsers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestBulkUsers_MissingFids(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := BulkUserParams{}
	ctx := context.Background()
	_, err := client.User.BulkUsers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fids"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBulkUsers_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/bulk"
	expectedParams := url.Values{
		"fids": []string{"12345,67890"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := BulkUserParams{
		Fids: "12345,67890",
	}
	ctx := context.Background()
	_, err := client.User.BulkUsers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPowerUsers_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/power"
	expectedParams := url.Values{
		"viewer_fid": []string{"12345"},
		"limit":      []string{"10"},
		"cursor":     []string{"someCursor"},
	}
	mockResponse := UsersResponse{
		Users: []User{{Fid: 12345, Username: "testUser"}},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := PowerUserParams{
		ViewerFid: 12345,
		Limit:     10,
		Cursor:    "someCursor",
	}
	ctx := context.Background()
	result, err := client.User.PowerUsers(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestPowerUsers_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/power"
	expectedParams := url.Values{
		"viewer_fid": []string{"12345"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := PowerUserParams{
		ViewerFid: 12345,
	}
	ctx := context.Background()
	_, err := client.User.PowerUsers(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBulkUsersByAddress_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/user/bulk-by-address"
	expectedParams := url.Values{
		"addresses":     []string{"addr1,addr2"},
		"address_types": []string{"type1,type2"},
		"viewer_fid":    []string{"12345"},
	}
	mockResponse := BulkUsersByAddressResponse{
		BulkUsersByAddress: map[string][]User{
			"addr1": {{Fid: 12345, Username: "user1"}},
			"addr2": {{Fid: 67890, Username: "user2"}},
		},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := BulkUserByAddressParams{
		Addresses:    []string{"addr1", "addr2"},
		AddressTypes: []string{"type1", "type2"},
		ViewerFid:    12345,
	}
	ctx := context.Background()
	result, err := client.User.BulkUsersByAddress(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestBulkUsersByAddress_MissingAddresses(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := BulkUserByAddressParams{}
	ctx := context.Background()
	_, err := client.User.BulkUsersByAddress(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Addresses"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestBulkUsersByAddress_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/user/bulk-by-address"
	expectedParams := url.Values{
		"addresses": []string{"addr1"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := BulkUserByAddressParams{
		Addresses: []string{"addr1"},
	}
	ctx := context.Background()
	_, err := client.User.BulkUsersByAddress(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchAuthorizationUrl_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/login/authorize"
	expectedParams := url.Values{
		"client_id":     []string{"testClientID"},
		"response_type": []string{"code"},
	}
	mockResponse := AuthorizationUrlResponse{
		AuthorizationUrl: "https://example.com/authorize?client_id=testClientID&response_type=code",
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := FetchAuthorizationUrlParams{
		ClientID:     "testClientID",
		ResponseType: "code",
	}
	ctx := context.Background()
	result, err := client.User.FetchAuthorizationUrl(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestFetchAuthorizationUrl_MissingClientID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchAuthorizationUrlParams{}
	ctx := context.Background()
	_, err := client.User.FetchAuthorizationUrl(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: ClientID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchAuthorizationUrl_DefaultResponseType(t *testing.T) {
	expectedPath := "/v2/farcaster/login/authorize"
	expectedParams := url.Values{
		"client_id":     []string{"testClientID"},
		"response_type": []string{"code"},
	}
	mockResponse := AuthorizationUrlResponse{
		AuthorizationUrl: "https://example.com/authorize?client_id=testClientID&response_type=code",
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := FetchAuthorizationUrlParams{
		ClientID: "testClientID",
	}
	ctx := context.Background()
	result, err := client.User.FetchAuthorizationUrl(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestFetchAuthorizationUrl_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/login/authorize"
	expectedParams := url.Values{
		"client_id":     []string{"testClientID"},
		"response_type": []string{"code"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchAuthorizationUrlParams{
		ClientID:     "testClientID",
		ResponseType: "code",
	}
	ctx := context.Background()
	_, err := client.User.FetchAuthorizationUrl(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
