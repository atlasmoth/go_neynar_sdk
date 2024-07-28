package neynarsdk

import (
	"context"
	"net/http"
)


type SubscriptionService struct {
	client *Client
}



type RetrieveSubscriptionsCreatedParams struct {
	Fid                  int32  `json:"fid"`
	SubscriptionProvider string `json:"subscription_provider"`
}


type RetrieveSubscriptionsCreatedResult struct {
	SubscriptionsCreated []Subscription `json:"subscriptions_created"`
	ErrorResponse
}


type RetrieveSubscribedToParams struct {
	Fid                  int32  `json:"fid"`
	ViewerFid            int32  `json:"viewer_fid,omitempty"`
	SubscriptionProvider string `json:"subscription_provider"`
}


type RetrieveSubscribedToResult struct {
	SubscribedTo []SubscribedTo `json:"subscribed_to"`
	ErrorResponse
}


type RetrieveSubscribersParams struct {
	Fid                  int32  `json:"fid"`
	ViewerFid            int32  `json:"viewer_fid,omitempty"`
	SubscriptionProvider string `json:"subscription_provider"`
}


type RetrieveSubscribersResult struct {
	Subscribers []Subscriber `json:"subscribers"`
	ErrorResponse
}


func (s *SubscriptionService) RetrieveSubscriptionsCreated(ctx context.Context, params RetrieveSubscriptionsCreatedParams) (RetrieveSubscriptionsCreatedResult, error) {
	var result RetrieveSubscriptionsCreatedResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.SubscriptionProvider == "" {
		return result, &RequiredFieldError{Field: "SubscriptionProvider"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/subscriptions_created"
	values := map[string]any{"fid": params.Fid, "subscription_provider": params.SubscriptionProvider}

	q := GetUrlValues(values)
	rawQuery := q.Encode()
	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}


func (s *SubscriptionService) RetrieveSubscribedTo(ctx context.Context, params RetrieveSubscribedToParams) (RetrieveSubscribedToResult, error) {
	var result RetrieveSubscribedToResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.SubscriptionProvider == "" {
		return result, &RequiredFieldError{Field: "SubscriptionProvider"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/subscribed_to"
	values := map[string]any{"fid": params.Fid, "subscription_provider": params.SubscriptionProvider}
	if params.ViewerFid > 0 {
		values["viewer_fid"] = params.ViewerFid
	}

	q := GetUrlValues(values)
	rawQuery := q.Encode()
	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}


func (s *SubscriptionService) RetrieveSubscribers(ctx context.Context, params RetrieveSubscribersParams) (RetrieveSubscribersResult, error) {
	var result RetrieveSubscribersResult
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.SubscriptionProvider == "" {
		return result, &RequiredFieldError{Field: "SubscriptionProvider"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/subscribers"
	values := map[string]any{"fid": params.Fid, "viewer_fid": params.ViewerFid, "subscription_provider": params.SubscriptionProvider}

	q := GetUrlValues(values)
	rawQuery := q.Encode()
	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, &rawQuery)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		if result.Code != "" {
			return result, &result.ErrorResponse
		}
		return result, nil
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}
