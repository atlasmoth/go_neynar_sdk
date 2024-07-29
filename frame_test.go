package neynarsdk

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestPostFrameAction_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/action"
	mockResponse := PostFrameActionResponse{
		Frame: Frame{},
		ErrorResponse: ErrorResponse{
			Code:    "",
			Message: "",
		},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	reqBody := FrameActionReqBody{
		SignerUUID: "test-signer-uuid",
		CastHash:   "test-cast-hash",
		Action:     FrameAction{ /* Initialize as required */ },
	}
	ctx := context.Background()
	result, err := client.Frame.PostFrameAction(ctx, reqBody)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestPostFrameAction_MissingFields(t *testing.T) {
	tests := []struct {
		name        string
		reqBody     FrameActionReqBody
		expectedErr string
	}{
		{"Missing SignerUUID", FrameActionReqBody{CastHash: "test-cast-hash", Action: FrameAction{}}, "The following field is required: SignerUUID"},
		{"Missing CastHash", FrameActionReqBody{SignerUUID: "test-signer-uuid", Action: FrameAction{}}, "The following field is required: CastHash"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := NewTestClient(nil)
			defer server.Close()

			ctx := context.Background()
			_, err := client.Frame.PostFrameAction(ctx, tt.reqBody)
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if err.Error() != tt.expectedErr {
				t.Errorf("Expected error message %v, got %v", tt.expectedErr, err.Error())
			}
		})
	}
}

func TestPostFrameAction_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/action"
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	reqBody := FrameActionReqBody{
		SignerUUID: "test-signer-uuid",
		CastHash:   "test-cast-hash",
		Action:     FrameAction{ /* Initialize as required */ },
	}
	ctx := context.Background()
	_, err := client.Frame.PostFrameAction(ctx, reqBody)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestValidateFrame_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/validate"
	mockResponse := ValidateFrameActionResponse{
		Valid: true,
		Action: ValidatedFrameAction{ /* Initialize as required */ },
		ErrorResponse: ErrorResponse{
			Code:    "",
			Message: "",
		},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	reqBody := ValidateFrameActionRequest{
		MessageBytesInHex: "test-message-hex",
	}
	ctx := context.Background()
	result, err := client.Frame.ValidateFrame(ctx, reqBody)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestValidateFrame_MissingMessageBytesInHex(t *testing.T) {
	client, server := NewTestClient(nil)
	defer server.Close()

	reqBody := ValidateFrameActionRequest{}
	ctx := context.Background()
	_, err := client.Frame.ValidateFrame(ctx, reqBody)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := "The following field is required: message_bytes_in_hex"
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestValidateFrame_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/validate"
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	reqBody := ValidateFrameActionRequest{
		MessageBytesInHex: "test-message-hex",
	}
	ctx := context.Background()
	_, err := client.Frame.ValidateFrame(ctx, reqBody)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestGetValidatedFrames_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/validate/list"
	mockResponse := FrameValidateListResponse{
		Frames: []string{"frame1", "frame2"},
		ErrorResponse: ErrorResponse{
			Code:    "",
			Message: "",
		},
	}

	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusOK))
	defer server.Close()

	ctx := context.Background()
	result, err := client.Frame.GetValidatedFrames(ctx)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestGetValidatedFrames_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/validate/list"
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	client, server := NewTestClient(mockHandler(t, expectedPath, nil, mockResponse, http.StatusInternalServerError))
	defer server.Close()

	ctx := context.Background()
	_, err := client.Frame.GetValidatedFrames(ctx)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}

func TestGetFrameAnalytics_Success(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/validate/analytics"
	mockResponse := FrameAnalyticsResponse{
		FrameValidateAnalyticsInteractors: FrameValidateAnalyticsInteractors{},
		FrameValidateAnalyticsTotalInteractors: FrameValidateAnalyticsTotalInteractors{},
		FrameValidateAnalyticsInteractionsPerCast: FrameValidateAnalyticsInteractionsPerCast{},
		FrameValidateAnalyticsInputText: FrameValidateAnalyticsInputText{},
		ErrorResponse: ErrorResponse{
			Code:    "",
			Message: "",
		},
	}

	params := FrameAnalyticsRequestParams{
		FrameURL:      "http://example.com/frame",
		AnalyticsType: "total-interactors",
		Start:         time.Now().Add(-24 * time.Hour),
		Stop:          time.Now(),
	}
	values := map[string]any{
		"frame_url":         params.FrameURL,
		"analytics_type":    params.AnalyticsType,
		"start":             params.Start.Format(time.RFC3339),
		"stop":              params.Stop.Format(time.RFC3339),
	}
	rawQuery := GetUrlValues(values)
	client, server := NewTestClient(mockHandler(t, expectedPath, rawQuery, mockResponse, http.StatusOK))
	defer server.Close()

	ctx := context.Background()
	result, err := client.Frame.GetFrameAnalytics(ctx, params)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, mockResponse) {
		t.Errorf("Expected result %v, got %v", mockResponse, result)
	}
}

func TestGetFrameAnalytics_MissingFields(t *testing.T) {
	tests := []struct {
		name        string
		params      FrameAnalyticsRequestParams
		expectedErr string
	}{
		{"Missing FrameURL", FrameAnalyticsRequestParams{AnalyticsType: "total-interactors", Start: time.Now(), Stop: time.Now()}, "The following field is required: FrameURL"},
		{"Missing AnalyticsType", FrameAnalyticsRequestParams{FrameURL: "http://example.com/frame", Start: time.Now(), Stop: time.Now()}, "The following field is required: AnalyticsType"},
		{"Missing Start", FrameAnalyticsRequestParams{FrameURL: "http://example.com/frame", AnalyticsType: "total-interactors", Stop: time.Now()}, "The following field is required: Start"},
		{"Missing Stop", FrameAnalyticsRequestParams{FrameURL: "http://example.com/frame", AnalyticsType: "total-interactors", Start: time.Now()}, "The following field is required: Stop"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, server := NewTestClient(nil)
			defer server.Close()

			ctx := context.Background()
			_, err := client.Frame.GetFrameAnalytics(ctx, tt.params)
			if err == nil {
				t.Errorf("Expected error, got nil")
			}
			if err.Error() != tt.expectedErr {
				t.Errorf("Expected error message %v, got %v", tt.expectedErr, err.Error())
			}
		})
	}
}

func TestGetFrameAnalytics_ServerError(t *testing.T) {
	expectedPath := "/v2/farcaster/frame/validate/analytics"
	mockResponse := ErrorResponse{
		Code:    "500",
		Message: "Internal Server Error",
	}
	params := FrameAnalyticsRequestParams{
		FrameURL:      "http://example.com/frame",
		AnalyticsType: "total-interactors",
		Start:         time.Now().Add(-24 * time.Hour),
		Stop:          time.Now(),
	}
	values := map[string]any{
		"frame_url":         params.FrameURL,
		"analytics_type":    params.AnalyticsType,
		"start":             params.Start.Format(time.RFC3339),
		"stop":              params.Stop.Format(time.RFC3339),
	}
	rawQuery := GetUrlValues(values)
	client, server := NewTestClient(mockHandler(t, expectedPath,rawQuery , mockResponse, http.StatusInternalServerError))
	defer server.Close()

	ctx := context.Background()
	_, err := client.Frame.GetFrameAnalytics(ctx, params)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	expectedError := mockResponse.Error()
	if err.Error() != expectedError {
		t.Errorf("Expected error message %v, got %v", expectedError, err.Error())
	}
}
