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

	mockResponse := PostCastResult{}
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

	params := PostCastParams{}
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

func TestDeleteCast_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"

	mockResponse := DeleteCastResult{
		// fill in the fields as necessary
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	params := DeleteCastParams{
		SignerUUID: "test-signer-uuid",
		TargetHash: "test-target-hash",
	}
	ctx := context.Background()
	result, err := client.Cast.DeleteCast(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestDeleteCast_MissingSignerUUID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := DeleteCastParams{
		TargetHash: "test-target-hash",
	}
	ctx := context.Background()
	_, err := client.Cast.DeleteCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SignerUUID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteCast_MissingTargetHash(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := DeleteCastParams{
		SignerUUID: "test-signer-uuid",
	}
	ctx := context.Background()
	_, err := client.Cast.DeleteCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: TargetHash"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteCast_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"

	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := DeleteCastParams{
		SignerUUID: "test-signer-uuid",
		TargetHash: "test-target-hash",
	}
	ctx := context.Background()
	_, err := client.Cast.DeleteCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteCast_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/cast"

	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := DeleteCastParams{
		SignerUUID: "test-signer-uuid",
		TargetHash: "test-target-hash",
	}
	ctx := context.Background()
	_, err := client.Cast.DeleteCast(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveCastsFromArray_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/casts"

	expectedParams := url.Values{
		"casts":      []string{"test-casts"},
		"viewer_fid": []string{"12345"},
		"sort_type":  []string{"date"},
	}
	mockResponse := RetrieveCastsFromArrayResult{
		Casts: []Cast{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveCastsFromArrayParams{
		Casts:     "test-casts",
		ViewerFid: 12345,
		SortType:  "date",
	}
	ctx := context.Background()
	result, err := client.Cast.RetrieveCastsFromArray(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveCastsFromArray_MissingCasts(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveCastsFromArrayParams{
		ViewerFid: 12345,
		SortType:  "date",
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastsFromArray(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Casts"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveCastsFromArray_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/casts"
	expectedParams := url.Values{
		"casts":      []string{"test-casts"},
		"viewer_fid": []string{"12345"},
		"sort_type":  []string{"date"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveCastsFromArrayParams{
		Casts:     "test-casts",
		ViewerFid: 12345,
		SortType:  "date",
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastsFromArray(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveCastsFromArray_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/casts"
	expectedParams := url.Values{
		"casts":      []string{"test-casts"},
		"viewer_fid": []string{"12345"},
		"sort_type":  []string{"date"},
	}
	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := RetrieveCastsFromArrayParams{
		Casts:     "test-casts",
		ViewerFid: 12345,
		SortType:  "date",
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveCastsFromArray(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveConversation_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/cast/conversation"
	expectedParams := url.Values{
		"type":                               []string{"test-type"},
		"identifier":                         []string{"test-identifier"},
		"reply_depth":                        []string{"2"},
		"include_chronological_parent_casts": []string{"true"},
		"limit":                              []string{"10"},
		"viewer_fid":                         []string{"12345"},
	}
	mockResponse := RetrieveConversationResult{
		Conversation: Conversation{},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusOK))
	defer server.Close()

	params := RetrieveConversationParams{
		Type:                            "test-type",
		Identifier:                      "test-identifier",
		ReplyDepth:                      2,
		IncludeChronologicalParentCasts: true,
		Limit:                           10,
		Cursor:                          "",
		ViewerFid:                       12345,
	}
	ctx := context.Background()
	result, err := client.Cast.RetrieveConversation(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestRetrieveConversation_MissingType(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveConversationParams{
		Identifier: "test-identifier",
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveConversation(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveConversation_MissingIdentifier(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveConversationParams{
		Type: "test-type",
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveConversation(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Identifier"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveConversation_InvalidReplyDepth(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := RetrieveConversationParams{
		Type:       "test-type",
		Identifier: "test-identifier",
		ReplyDepth: 10,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveConversation(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := InvalidFieldError{Field: "ReplyDepth", Values: "0-5"}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveConversation_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/cast/conversation"
	expectedParams := url.Values{
		"type":                               []string{"test-type"},
		"identifier":                         []string{"test-identifier"},
		"reply_depth":                        []string{"2"},
		"include_chronological_parent_casts": []string{"true"},
		"limit":                              []string{"10"},
		"viewer_fid":                         []string{"12345"},
	}
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	params := RetrieveConversationParams{
		Type:                            "test-type",
		Identifier:                      "test-identifier",
		ReplyDepth:                      2,
		IncludeChronologicalParentCasts: true,
		Limit:                           10,
		ViewerFid:                       12345,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveConversation(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestRetrieveConversation_BadRequest(t *testing.T) {
	expectedPath := "/v2/farcaster/cast/conversation"
	expectedParams := url.Values{
		"type":                               []string{"test-type"},
		"identifier":                         []string{"test-identifier"},
		"reply_depth":                        []string{"2"},
		"include_chronological_parent_casts": []string{"true"},
		"limit":                              []string{"10"},
		"viewer_fid":                         []string{"12345"},
	}
	mockResponse := ErrorResponse{
		Code:    "400",
		Message: "Bad Request",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockResponse, http.StatusBadRequest))
	defer server.Close()

	params := RetrieveConversationParams{
		Type:                            "test-type",
		Identifier:                      "test-identifier",
		ReplyDepth:                      2,
		IncludeChronologicalParentCasts: true,
		Limit:                           10,
		ViewerFid:                       12345,
	}
	ctx := context.Background()
	_, err := client.Cast.RetrieveConversation(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
