package neynarsdk

import (
	"context"
	"net/http"
)

type NotificationService struct {
	client *Client
}

type RetrieveNotificationsForUserParams struct {
	Cursor        string
	Fid           int32
	IsNotPriority bool
}

type RetrieveNotificationsForUserResult struct {
	Notifications []Notification `json:"notifications"`
	ErrorResponse
	Next Next `json:"next"`
}

func (n *NotificationService) RetrieveNotificationsForUser(ctx context.Context, params RetrieveNotificationsForUserParams) (RetrieveNotificationsForUserResult, error) {
	var result RetrieveNotificationsForUserResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}

	baseURL := n.client.BaseURL.String() + "v2/farcaster/notifications"

	values := map[string]any{"fid": params.Fid, "cursor": params.Cursor, "is_priority" : !params.IsNotPriority}

	q := GetUrlValues(values)

	rawQuery := q.Encode()
	resp, err := n.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {

		err = n.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = n.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}

}
