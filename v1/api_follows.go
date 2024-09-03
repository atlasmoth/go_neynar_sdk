/*
Farcaster API V1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
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


// FollowsAPIService FollowsAPI service
type FollowsAPIService service

type ApiFollowersRequest struct {
	ctx context.Context
	ApiService *FollowsAPIService
	apiKey *string
	fid *int32
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiFollowersRequest) ApiKey(apiKey string) ApiFollowersRequest {
	r.apiKey = &apiKey
	return r
}

// FID of the user
func (r ApiFollowersRequest) Fid(fid int32) ApiFollowersRequest {
	r.fid = &fid
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiFollowersRequest) ViewerFid(viewerFid int32) ApiFollowersRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 25, max 150)
func (r ApiFollowersRequest) Limit(limit int32) ApiFollowersRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFollowersRequest) Cursor(cursor string) ApiFollowersRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFollowersRequest) Execute() (*FollowResponse, *http.Response, error) {
	return r.ApiService.FollowersExecute(r)
}

/*
Followers Gets all followers for a given FID

Gets a list of users who follow a given user in reverse chronological order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFollowersRequest
*/
func (a *FollowsAPIService) Followers(ctx context.Context) ApiFollowersRequest {
	return ApiFollowersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FollowResponse
func (a *FollowsAPIService) FollowersExecute(r ApiFollowersRequest) (*FollowResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FollowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FollowsAPIService.Followers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/followers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewerFid", r.viewerFid, "form", "")
	} else {
		var defaultValue int32 = 3
		r.viewerFid = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 25
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
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

type ApiFollowingRequest struct {
	ctx context.Context
	ApiService *FollowsAPIService
	apiKey *string
	fid *int32
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiFollowingRequest) ApiKey(apiKey string) ApiFollowingRequest {
	r.apiKey = &apiKey
	return r
}

// FID of the user
func (r ApiFollowingRequest) Fid(fid int32) ApiFollowingRequest {
	r.fid = &fid
	return r
}

// fid of the user viewing this information, needed for contextual information.
func (r ApiFollowingRequest) ViewerFid(viewerFid int32) ApiFollowingRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 25, max 150)
func (r ApiFollowingRequest) Limit(limit int32) ApiFollowingRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiFollowingRequest) Cursor(cursor string) ApiFollowingRequest {
	r.cursor = &cursor
	return r
}

func (r ApiFollowingRequest) Execute() (*FollowResponse, *http.Response, error) {
	return r.ApiService.FollowingExecute(r)
}

/*
Following Gets all following users of a FID

Gets a list of users who is following a given user in reverse chronological order.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFollowingRequest
*/
func (a *FollowsAPIService) Following(ctx context.Context) ApiFollowingRequest {
	return ApiFollowingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FollowResponse
func (a *FollowsAPIService) FollowingExecute(r ApiFollowingRequest) (*FollowResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FollowResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FollowsAPIService.Following")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/following"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewerFid", r.viewerFid, "form", "")
	} else {
		var defaultValue int32 = 3
		r.viewerFid = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 25
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
	}
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
