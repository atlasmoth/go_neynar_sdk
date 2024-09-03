/*
Raw Farcaster Hub API

Perform basic queries of Farcaster state via the REST API of a Farcaster hub. See the [Farcaster docs](https://www.thehubble.xyz/docs/httpapi/httpapi.html) for more details. 

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


// ReactionsAPIService ReactionsAPI service
type ReactionsAPIService service

type ApiGetReactionByIdRequest struct {
	ctx context.Context
	ApiService *ReactionsAPIService
	apiKey *string
	fid *int32
	targetFid *int32
	targetHash *string
	reactionType *ReactionType
}

// API key required for authentication.
func (r ApiGetReactionByIdRequest) ApiKey(apiKey string) ApiGetReactionByIdRequest {
	r.apiKey = &apiKey
	return r
}

// The FID of the reaction&#39;s creator
func (r ApiGetReactionByIdRequest) Fid(fid int32) ApiGetReactionByIdRequest {
	r.fid = &fid
	return r
}

// The FID of the cast&#39;s creator
func (r ApiGetReactionByIdRequest) TargetFid(targetFid int32) ApiGetReactionByIdRequest {
	r.targetFid = &targetFid
	return r
}

// The cast&#39;s hash
func (r ApiGetReactionByIdRequest) TargetHash(targetHash string) ApiGetReactionByIdRequest {
	r.targetHash = &targetHash
	return r
}

// The type of reaction, either as a numerical enum value or string representation
func (r ApiGetReactionByIdRequest) ReactionType(reactionType ReactionType) ApiGetReactionByIdRequest {
	r.reactionType = &reactionType
	return r
}

func (r ApiGetReactionByIdRequest) Execute() (*Reaction, *http.Response, error) {
	return r.ApiService.GetReactionByIdExecute(r)
}

/*
GetReactionById Get a reaction by its created FID and target Cast.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetReactionByIdRequest
*/
func (a *ReactionsAPIService) GetReactionById(ctx context.Context) ApiGetReactionByIdRequest {
	return ApiGetReactionByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Reaction
func (a *ReactionsAPIService) GetReactionByIdExecute(r ApiGetReactionByIdRequest) (*Reaction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Reaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionsAPIService.GetReactionById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reactionById"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.targetFid == nil {
		return localVarReturnValue, nil, reportError("targetFid is required and must be specified")
	}
	if r.targetHash == nil {
		return localVarReturnValue, nil, reportError("targetHash is required and must be specified")
	}
	if r.reactionType == nil {
		return localVarReturnValue, nil, reportError("reactionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "target_fid", r.targetFid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "target_hash", r.targetHash, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "reaction_type", r.reactionType, "form", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiListReactionsByCastRequest struct {
	ctx context.Context
	ApiService *ReactionsAPIService
	apiKey *string
	targetFid *int32
	targetHash *string
	reactionType *ReactionType
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListReactionsByCastRequest) ApiKey(apiKey string) ApiListReactionsByCastRequest {
	r.apiKey = &apiKey
	return r
}

// The FID of the cast&#39;s creator
func (r ApiListReactionsByCastRequest) TargetFid(targetFid int32) ApiListReactionsByCastRequest {
	r.targetFid = &targetFid
	return r
}

// The hash of the cast
func (r ApiListReactionsByCastRequest) TargetHash(targetHash string) ApiListReactionsByCastRequest {
	r.targetHash = &targetHash
	return r
}

// The type of reaction, either as a numerical enum value or string representation
func (r ApiListReactionsByCastRequest) ReactionType(reactionType ReactionType) ApiListReactionsByCastRequest {
	r.reactionType = &reactionType
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListReactionsByCastRequest) PageSize(pageSize int32) ApiListReactionsByCastRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListReactionsByCastRequest) Reverse(reverse bool) ApiListReactionsByCastRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListReactionsByCastRequest) PageToken(pageToken string) ApiListReactionsByCastRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListReactionsByCastRequest) Execute() (*ListReactionsByCast200Response, *http.Response, error) {
	return r.ApiService.ListReactionsByCastExecute(r)
}

/*
ListReactionsByCast Get all reactions to a cast

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListReactionsByCastRequest
*/
func (a *ReactionsAPIService) ListReactionsByCast(ctx context.Context) ApiListReactionsByCastRequest {
	return ApiListReactionsByCastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListReactionsByCast200Response
func (a *ReactionsAPIService) ListReactionsByCastExecute(r ApiListReactionsByCastRequest) (*ListReactionsByCast200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListReactionsByCast200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionsAPIService.ListReactionsByCast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reactionsByCast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.targetFid == nil {
		return localVarReturnValue, nil, reportError("targetFid is required and must be specified")
	}
	if r.targetHash == nil {
		return localVarReturnValue, nil, reportError("targetHash is required and must be specified")
	}
	if r.reactionType == nil {
		return localVarReturnValue, nil, reportError("reactionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "target_fid", r.targetFid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "target_hash", r.targetHash, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "reaction_type", r.reactionType, "form", "")
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.reverse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reverse", r.reverse, "form", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiListReactionsByFidRequest struct {
	ctx context.Context
	ApiService *ReactionsAPIService
	apiKey *string
	fid *int32
	reactionType *ReactionType
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListReactionsByFidRequest) ApiKey(apiKey string) ApiListReactionsByFidRequest {
	r.apiKey = &apiKey
	return r
}

// The FID of the reaction&#39;s creator
func (r ApiListReactionsByFidRequest) Fid(fid int32) ApiListReactionsByFidRequest {
	r.fid = &fid
	return r
}

// The type of reaction, either as a numerical enum value or string representation
func (r ApiListReactionsByFidRequest) ReactionType(reactionType ReactionType) ApiListReactionsByFidRequest {
	r.reactionType = &reactionType
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListReactionsByFidRequest) PageSize(pageSize int32) ApiListReactionsByFidRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListReactionsByFidRequest) Reverse(reverse bool) ApiListReactionsByFidRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListReactionsByFidRequest) PageToken(pageToken string) ApiListReactionsByFidRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListReactionsByFidRequest) Execute() (*ListReactionsByCast200Response, *http.Response, error) {
	return r.ApiService.ListReactionsByFidExecute(r)
}

/*
ListReactionsByFid Get all reactions by an FID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListReactionsByFidRequest
*/
func (a *ReactionsAPIService) ListReactionsByFid(ctx context.Context) ApiListReactionsByFidRequest {
	return ApiListReactionsByFidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListReactionsByCast200Response
func (a *ReactionsAPIService) ListReactionsByFidExecute(r ApiListReactionsByFidRequest) (*ListReactionsByCast200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListReactionsByCast200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionsAPIService.ListReactionsByFid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reactionsByFid"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.reactionType == nil {
		return localVarReturnValue, nil, reportError("reactionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "reaction_type", r.reactionType, "form", "")
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.reverse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reverse", r.reverse, "form", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiListReactionsByTargetRequest struct {
	ctx context.Context
	ApiService *ReactionsAPIService
	apiKey *string
	url *string
	reactionType *ReactionType
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListReactionsByTargetRequest) ApiKey(apiKey string) ApiListReactionsByTargetRequest {
	r.apiKey = &apiKey
	return r
}

// The URL of the parent cast
func (r ApiListReactionsByTargetRequest) Url(url string) ApiListReactionsByTargetRequest {
	r.url = &url
	return r
}

// The type of reaction, either as a numerical enum value or string representation
func (r ApiListReactionsByTargetRequest) ReactionType(reactionType ReactionType) ApiListReactionsByTargetRequest {
	r.reactionType = &reactionType
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListReactionsByTargetRequest) PageSize(pageSize int32) ApiListReactionsByTargetRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListReactionsByTargetRequest) Reverse(reverse bool) ApiListReactionsByTargetRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListReactionsByTargetRequest) PageToken(pageToken string) ApiListReactionsByTargetRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListReactionsByTargetRequest) Execute() (*ListReactionsByCast200Response, *http.Response, error) {
	return r.ApiService.ListReactionsByTargetExecute(r)
}

/*
ListReactionsByTarget Get all reactions to a target URL

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListReactionsByTargetRequest
*/
func (a *ReactionsAPIService) ListReactionsByTarget(ctx context.Context) ApiListReactionsByTargetRequest {
	return ApiListReactionsByTargetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListReactionsByCast200Response
func (a *ReactionsAPIService) ListReactionsByTargetExecute(r ApiListReactionsByTargetRequest) (*ListReactionsByCast200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListReactionsByCast200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReactionsAPIService.ListReactionsByTarget")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/reactionsByTarget"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.url == nil {
		return localVarReturnValue, nil, reportError("url is required and must be specified")
	}
	if r.reactionType == nil {
		return localVarReturnValue, nil, reportError("reactionType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "reaction_type", r.reactionType, "form", "")
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.reverse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reverse", r.reverse, "form", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
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
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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