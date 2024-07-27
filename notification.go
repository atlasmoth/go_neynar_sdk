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
type RetrieveNotificationsForChannelsParams struct {
	RetrieveNotificationsForUserParams
	ChannelIds        string
}

func (n *NotificationService) RetrieveNotificationsForChannels(ctx context.Context, params RetrieveNotificationsForChannelsParams) (RetrieveNotificationsForUserResult, error) {
	var result RetrieveNotificationsForUserResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.ChannelIds == ""{
		return result, &RequiredFieldError{Field: "ChannelIds"}
	}

	baseURL := n.client.BaseURL.String() + "v2/farcaster/notifications/channel"

	values := map[string]any{"fid": params.Fid, "cursor": params.Cursor, "is_priority" : !params.IsNotPriority,"channel_ids" : params.ChannelIds}

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

type RetrieveNotificationsParentUrlParams struct {
	RetrieveNotificationsForUserParams
	ParentUrls        string
}

func (n *NotificationService) RetrieveNotificationsParentUrl(ctx context.Context, params RetrieveNotificationsParentUrlParams) (RetrieveNotificationsForUserResult, error) {
	var result RetrieveNotificationsForUserResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.ParentUrls == ""{
		return result, &RequiredFieldError{Field: "ParentUrls"}
	}

	baseURL := n.client.BaseURL.String() + "v2/farcaster/notifications/parent_url"

	values := map[string]any{"fid": params.Fid, "cursor": params.Cursor, "is_priority" : !params.IsNotPriority,"parent_urls" : params.ParentUrls}

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
