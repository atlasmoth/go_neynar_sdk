package neynarsdk

import (
	"context"
	"fmt"
	"net/http"
)

type WebhookService struct {
	client *Client
}

type WebhookResponse struct {
	Message string   `json:"message,omitempty"`
	Success bool     `json:"success,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
	ErrorResponse
}

type GetWebhookParams struct {
	WebhookID string
}

func (s *WebhookService) GetWebhook(ctx context.Context, params GetWebhookParams) (WebhookResponse, error) {
	var result WebhookResponse

	if params.WebhookID == "" {
		return result, &RequiredFieldError{Field: "WebhookID"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/webhook"

	values := map[string]any{
		"webhook_id": params.WebhookID,
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

type WebhookParams struct {
	WebhookID    string                      `json:"webhook_id,omitempty"`
	Name         string                      `json:"name,omitempty"`
	URL          string                      `json:"url,omitempty"`
	Active       bool                        `json:"active,omitempty"`
	Subscription *WebhookSubscriptionFilters `json:"subscription,omitempty"`
}

func (s *WebhookService) CreateWebhook(ctx context.Context, params WebhookParams) (WebhookResponse, error) {
	var result WebhookResponse

	if params.Name == "" {
		return result, &RequiredFieldError{Field: "name"}
	}
	if params.URL == "" {
		return result, &RequiredFieldError{Field: "url"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/webhook"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, params, nil)
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

func (s *WebhookService) UpdateWebhookActiveStatus(ctx context.Context, params WebhookParams) (WebhookResponse, error) {
	var result WebhookResponse

	if params.WebhookID == "" {
		return result, &RequiredFieldError{Field: "webhook_id"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/webhook"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPatch, baseURL, params, nil)
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

func (s *WebhookService) UpdateWebhook(ctx context.Context, params WebhookParams) (WebhookResponse, error) {
	var result WebhookResponse

	if params.WebhookID == "" {
		return result, &RequiredFieldError{Field: "WebhookID"}
	}
	if params.Name == "" {
		return result, &RequiredFieldError{Field: "Name"}
	}
	if params.URL == "" {
		return result, &RequiredFieldError{Field: "URL"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/webhook"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPut, baseURL, params, nil)
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

type WebhookListResponse struct {
	Webhooks []Webhook `json:"webhooks"`
}

type WebhookSubscriptionFilters struct {
	CastCreated     CastCreatedFilter                  `json:"cast.created,omitempty"`
	UserCreated     interface{}                        `json:"user.created,omitempty"`
	UserUpdated     UserUpdatedFilter                  `json:"user.updated,omitempty"`
	FollowCreated   WebhookSubscriptionFiltersFollow   `json:"follow.created,omitempty"`
	FollowDeleted   WebhookSubscriptionFiltersFollow   `json:"follow.deleted,omitempty"`
	ReactionCreated WebhookSubscriptionFiltersReaction `json:"reaction.created,omitempty"`
	ReactionDeleted WebhookSubscriptionFiltersReaction `json:"reaction.deleted,omitempty"`
}

type WebhookSubscriptionFiltersFollow struct {
	FIDs       []int `json:"fids,omitempty"`
	TargetFIDs []int `json:"target_fids,omitempty"`
}

type WebhookSubscriptionFiltersReaction struct {
	FIDs       []int `json:"fids,omitempty"`
	TargetFIDs []int `json:"target_fids,omitempty"`
}

type CastCreatedFilter struct {
	ExcludeAuthorFIDs []int    `json:"exclude_author_fids,omitempty"`
	AuthorFIDs        []int    `json:"author_fids,omitempty"`
	MentionedFIDs     []int    `json:"mentioned_fids,omitempty"`
	ParentURLs        []string `json:"parent_urls,omitempty"`
	RootParentURLs    []string `json:"root_parent_urls,omitempty"`
	ParentAuthorFIDs  []int    `json:"parent_author_fids,omitempty"`
	Text              string   `json:"text,omitempty"`
	Embeds            string   `json:"embeds,omitempty"`
}

type UserUpdatedFilter struct {
	FIDs []int `json:"fids,omitempty"`
}

func (s *WebhookService) DeleteWebhook(ctx context.Context, webhookID string) (WebhookResponse, error) {
	var result WebhookResponse

	if webhookID == "" {
		return result, &RequiredFieldError{Field: "webhookID"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/webhook"

	params := WebhookParams{
		WebhookID: webhookID,
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodDelete, baseURL, params, nil)
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

func (s *WebhookService) ListWebhooks(ctx context.Context) (WebhookListResponse, error) {
	var result WebhookListResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/webhook/list"

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodGet, baseURL, nil, nil)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	}

	return result, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}
