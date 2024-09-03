/*
Farcaster API V2

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SubscribersAPIService SubscribersAPI service
type SubscribersAPIService service

type ApiSubscribedToRequest struct {
	ctx context.Context
	ApiService *SubscribersAPIService
	apiKey *string
	fid *int32
	subscriptionProvider *SubscriptionProvider
	viewerFid *int32
}

// API key required for authentication.
func (r ApiSubscribedToRequest) ApiKey(apiKey string) ApiSubscribedToRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiSubscribedToRequest) Fid(fid int32) ApiSubscribedToRequest {
	r.fid = &fid
	return r
}

func (r ApiSubscribedToRequest) SubscriptionProvider(subscriptionProvider SubscriptionProvider) ApiSubscribedToRequest {
	r.subscriptionProvider = &subscriptionProvider
	return r
}

func (r ApiSubscribedToRequest) ViewerFid(viewerFid int32) ApiSubscribedToRequest {
	r.viewerFid = &viewerFid
	return r
}

func (r ApiSubscribedToRequest) Execute() (*SubscribedToResponse, *http.Response, error) {
	return r.ApiService.SubscribedToExecute(r)
}

/*
SubscribedTo Fetch what a given fid is subscribed to

Fetch what fids and contracts a fid is subscribed to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscribedToRequest
*/
func (a *SubscribersAPIService) SubscribedTo(ctx context.Context) ApiSubscribedToRequest {
	return ApiSubscribedToRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscribedToResponse
func (a *SubscribersAPIService) SubscribedToExecute(r ApiSubscribedToRequest) (*SubscribedToResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscribedToResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscribersAPIService.SubscribedTo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/user/subscribed_to"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.subscriptionProvider == nil {
		return localVarReturnValue, nil, reportError("subscriptionProvider is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "subscription_provider", r.subscriptionProvider, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSubscribersRequest struct {
	ctx context.Context
	ApiService *SubscribersAPIService
	apiKey *string
	fid *int32
	subscriptionProvider *SubscriptionProviders
	viewerFid *int32
}

// API key required for authentication.
func (r ApiSubscribersRequest) ApiKey(apiKey string) ApiSubscribersRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiSubscribersRequest) Fid(fid int32) ApiSubscribersRequest {
	r.fid = &fid
	return r
}

func (r ApiSubscribersRequest) SubscriptionProvider(subscriptionProvider SubscriptionProviders) ApiSubscribersRequest {
	r.subscriptionProvider = &subscriptionProvider
	return r
}

func (r ApiSubscribersRequest) ViewerFid(viewerFid int32) ApiSubscribersRequest {
	r.viewerFid = &viewerFid
	return r
}

func (r ApiSubscribersRequest) Execute() (*SubscribersResponse, *http.Response, error) {
	return r.ApiService.SubscribersExecute(r)
}

/*
Subscribers Fetch subscribers for a given fid

Fetch subscribers for a given fid's contracts. Doesn't return addresses that don't have an fid.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscribersRequest
*/
func (a *SubscribersAPIService) Subscribers(ctx context.Context) ApiSubscribersRequest {
	return ApiSubscribersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscribersResponse
func (a *SubscribersAPIService) SubscribersExecute(r ApiSubscribersRequest) (*SubscribersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscribersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscribersAPIService.Subscribers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/user/subscribers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.subscriptionProvider == nil {
		return localVarReturnValue, nil, reportError("subscriptionProvider is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "subscription_provider", r.subscriptionProvider, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSubscriptionsCreatedRequest struct {
	ctx context.Context
	ApiService *SubscribersAPIService
	apiKey *string
	fid *int32
	subscriptionProvider *SubscriptionProvider
}

// API key required for authentication.
func (r ApiSubscriptionsCreatedRequest) ApiKey(apiKey string) ApiSubscriptionsCreatedRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiSubscriptionsCreatedRequest) Fid(fid int32) ApiSubscriptionsCreatedRequest {
	r.fid = &fid
	return r
}

func (r ApiSubscriptionsCreatedRequest) SubscriptionProvider(subscriptionProvider SubscriptionProvider) ApiSubscriptionsCreatedRequest {
	r.subscriptionProvider = &subscriptionProvider
	return r
}

func (r ApiSubscriptionsCreatedRequest) Execute() (*SubscriptionsResponse, *http.Response, error) {
	return r.ApiService.SubscriptionsCreatedExecute(r)
}

/*
SubscriptionsCreated Fetch created subscriptions for a given fid

Fetch created subscriptions for a given fid's.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSubscriptionsCreatedRequest
*/
func (a *SubscribersAPIService) SubscriptionsCreated(ctx context.Context) ApiSubscriptionsCreatedRequest {
	return ApiSubscriptionsCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SubscriptionsResponse
func (a *SubscribersAPIService) SubscriptionsCreatedExecute(r ApiSubscriptionsCreatedRequest) (*SubscriptionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SubscriptionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscribersAPIService.SubscriptionsCreated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/user/subscriptions_created"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.subscriptionProvider == nil {
		return localVarReturnValue, nil, reportError("subscriptionProvider is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "subscription_provider", r.subscriptionProvider, "form", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
