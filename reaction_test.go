package neynarsdk

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

var mockFetchUserReactionsResponse = ReactionsListResponse{
	Reactions: []ReactionWithCastInfo{},
	Next:      Next{Cursor: "nextCursor"},
}

var mockFetchCastReactionsResponse = ReactionsListResponse{
	Reactions: []ReactionWithCastInfo{},
	Next:      Next{Cursor: "nextCursor"},
}

var mockPostReactionResponse = ReactionResponse{
	Success: true,
	Message: "Reaction posted successfully",
}

var mockDeleteReactionResponse = ReactionResponse{
	Success: true,
	Message: "Reaction deleted successfully",
}

var mockErrorResponse = ErrorResponse{
	Code:    "500",
	Message: "Internal Server Error",
}

// FetchUserReactions tests
func TestFetchUserReactions_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/reactions/user"
	expectedParams := url.Values{
		"fid":    []string{"12345"},
		"type":   []string{"like"},
		"limit":  []string{"10"},
		"cursor": []string{"someCursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockFetchUserReactionsResponse, http.StatusOK))
	defer server.Close()

	params := FetchUserReactionsParams{
		Fid:    12345,
		Type:   "like",
		Limit:  10,
		Cursor: "someCursor",
	}
	ctx := context.Background()
	result, err := client.Reaction.FetchUserReactions(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockFetchUserReactionsResponse) {
		t.Errorf("Expected result %v, got %v", mockFetchUserReactionsResponse, result)
	}
}

func TestFetchUserReactions_MissingFid(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchUserReactionsParams{
		Type: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.FetchUserReactions(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Fid"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchUserReactions_MissingType(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchUserReactionsParams{
		Fid: 12345,
	}
	ctx := context.Background()
	_, err := client.Reaction.FetchUserReactions(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Type"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchUserReactions_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/reactions/user"
	expectedParams := url.Values{
		"fid":  []string{"12345"},
		"type": []string{"like"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchUserReactionsParams{
		Fid:  12345,
		Type: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.FetchUserReactions(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

// FetchCastReactions tests
func TestFetchCastReactions_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/reactions/cast"
	expectedParams := url.Values{
		"hash":   []string{"hash123"},
		"types":  []string{"like"},
		"limit":  []string{"10"},
		"cursor": []string{"someCursor"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockFetchCastReactionsResponse, http.StatusOK))
	defer server.Close()

	params := FetchCastReactionsParams{
		Hash:   "hash123",
		Types:  "like",
		Limit:  10,
		Cursor: "someCursor",
	}
	ctx := context.Background()
	result, err := client.Reaction.FetchCastReactions(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockFetchCastReactionsResponse) {
		t.Errorf("Expected result %v, got %v", mockFetchCastReactionsResponse, result)
	}
}

func TestFetchCastReactions_MissingHash(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchCastReactionsParams{
		Types: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.FetchCastReactions(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Hash"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchCastReactions_MissingTypes(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := FetchCastReactionsParams{
		Hash: "hash123",
	}
	ctx := context.Background()
	_, err := client.Reaction.FetchCastReactions(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Types"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestFetchCastReactions_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/reactions/cast"
	expectedParams := url.Values{
		"hash":  []string{"hash123"},
		"types": []string{"like"},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, expectedParams, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := FetchCastReactionsParams{
		Hash:  "hash123",
		Types: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.FetchCastReactions(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

// PostReaction tests
func TestPostReaction_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/reaction"
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockPostReactionResponse, http.StatusOK))
	defer server.Close()

	params := PostReactionParams{
		SignerUUID:   "uuid123",
		Target:       "target123",
		ReactionType: "like",
	}
	ctx := context.Background()
	result, err := client.Reaction.PostReaction(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockPostReactionResponse) {
		t.Errorf("Expected result %v, got %v", mockPostReactionResponse, result)
	}
}

func TestPostReaction_MissingSignerUUID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostReactionParams{
		Target:       "target123",
		ReactionType: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.PostReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SignerUUID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPostReaction_MissingTarget(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostReactionParams{
		SignerUUID:   "uuid123",
		ReactionType: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.PostReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Target"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPostReaction_MissingReactionType(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostReactionParams{
		SignerUUID: "uuid123",
		Target:     "target123",
	}
	ctx := context.Background()
	_, err := client.Reaction.PostReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: ReactionType"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestPostReaction_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/reaction"
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := PostReactionParams{
		SignerUUID:   "uuid123",
		Target:       "target123",
		ReactionType: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.PostReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

// DeleteReaction tests
func TestDeleteReaction_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/reaction"
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockDeleteReactionResponse, http.StatusOK))
	defer server.Close()

	params := PostReactionParams{
		SignerUUID:   "uuid123",
		Target:       "target123",
		ReactionType: "like",
	}
	ctx := context.Background()
	result, err := client.Reaction.DeleteReaction(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockDeleteReactionResponse) {
		t.Errorf("Expected result %v, got %v", mockDeleteReactionResponse, result)
	}
}

func TestDeleteReaction_MissingSignerUUID(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostReactionParams{
		Target:       "target123",
		ReactionType: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.DeleteReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: SignerUUID"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteReaction_MissingTarget(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostReactionParams{
		SignerUUID:   "uuid123",
		ReactionType: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.DeleteReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: Target"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteReaction_MissingReactionType(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	params := PostReactionParams{
		SignerUUID: "uuid123",
		Target:     "target123",
	}
	ctx := context.Background()
	_, err := client.Reaction.DeleteReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: ReactionType"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestDeleteReaction_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/reaction"
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockErrorResponse, http.StatusInternalServerError))
	defer server.Close()

	params := PostReactionParams{
		SignerUUID:   "uuid123",
		Target:       "target123",
		ReactionType: "like",
	}
	ctx := context.Background()
	_, err := client.Reaction.DeleteReaction(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockErrorResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
