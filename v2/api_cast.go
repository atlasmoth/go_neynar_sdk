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


// CastAPIService CastAPI service
type CastAPIService service

type ApiCastRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	identifier *string
	type_ *CastParamType
	viewerFid *int32
}

// API key required for authentication.
func (r ApiCastRequest) ApiKey(apiKey string) ApiCastRequest {
	r.apiKey = &apiKey
	return r
}

// Cast identifier (Its either a url or a hash)
func (r ApiCastRequest) Identifier(identifier string) ApiCastRequest {
	r.identifier = &identifier
	return r
}

func (r ApiCastRequest) Type_(type_ CastParamType) ApiCastRequest {
	r.type_ = &type_
	return r
}

// adds viewer_context to cast object to show whether viewer has liked or recasted the cast.
func (r ApiCastRequest) ViewerFid(viewerFid int32) ApiCastRequest {
	r.viewerFid = &viewerFid
	return r
}

func (r ApiCastRequest) Execute() (*CastResponse, *http.Response, error) {
	return r.ApiService.CastExecute(r)
}

/*
Cast Retrieve cast for a given hash or Warpcast URL

Gets information about an individual cast by passing in a Warpcast web URL or cast hash

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCastRequest
*/
func (a *CastAPIService) Cast(ctx context.Context) ApiCastRequest {
	return ApiCastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastResponse
func (a *CastAPIService) CastExecute(r ApiCastRequest) (*CastResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.Cast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.identifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifier", r.identifier, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
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

type ApiCastConversationRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	identifier *string
	type_ *CastParamType
	replyDepth *int32
	includeChronologicalParentCasts *bool
	viewerFid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiCastConversationRequest) ApiKey(apiKey string) ApiCastConversationRequest {
	r.apiKey = &apiKey
	return r
}

// Cast identifier (Its either a url or a hash)
func (r ApiCastConversationRequest) Identifier(identifier string) ApiCastConversationRequest {
	r.identifier = &identifier
	return r
}

func (r ApiCastConversationRequest) Type_(type_ CastParamType) ApiCastConversationRequest {
	r.type_ = &type_
	return r
}

// The depth of replies in the conversation that will be returned (default 2)
func (r ApiCastConversationRequest) ReplyDepth(replyDepth int32) ApiCastConversationRequest {
	r.replyDepth = &replyDepth
	return r
}

// Include all parent casts in chronological order
func (r ApiCastConversationRequest) IncludeChronologicalParentCasts(includeChronologicalParentCasts bool) ApiCastConversationRequest {
	r.includeChronologicalParentCasts = &includeChronologicalParentCasts
	return r
}

func (r ApiCastConversationRequest) ViewerFid(viewerFid int32) ApiCastConversationRequest {
	r.viewerFid = &viewerFid
	return r
}

// Number of results to retrieve (default 20, max 50)
func (r ApiCastConversationRequest) Limit(limit int32) ApiCastConversationRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiCastConversationRequest) Cursor(cursor string) ApiCastConversationRequest {
	r.cursor = &cursor
	return r
}

func (r ApiCastConversationRequest) Execute() (*Conversation, *http.Response, error) {
	return r.ApiService.CastConversationExecute(r)
}

/*
CastConversation Retrieve the conversation for a given cast

Gets all casts related to a conversation surrounding a cast by passing in a cast hash or Warpcast URL. Includes all the ancestors of a cast up to the root parent in a chronological order. Includes all direct_replies to the cast up to the reply_depth specified in the query parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCastConversationRequest
*/
func (a *CastAPIService) CastConversation(ctx context.Context) ApiCastConversationRequest {
	return ApiCastConversationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Conversation
func (a *CastAPIService) CastConversationExecute(r ApiCastConversationRequest) (*Conversation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Conversation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.CastConversation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast/conversation"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.identifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identifier", r.identifier, "form", "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "form", "")
	}
	if r.replyDepth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reply_depth", r.replyDepth, "form", "")
	} else {
		var defaultValue int32 = 2
		r.replyDepth = &defaultValue
	}
	if r.includeChronologicalParentCasts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_chronological_parent_casts", r.includeChronologicalParentCasts, "form", "")
	} else {
		var defaultValue bool = false
		r.includeChronologicalParentCasts = &defaultValue
	}
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 20
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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
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

type ApiCastSearchRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	q *string
	authorFid *int32
	parentUrl *string
	channelId *string
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiCastSearchRequest) ApiKey(apiKey string) ApiCastSearchRequest {
	r.apiKey = &apiKey
	return r
}

// Query string to search for casts
func (r ApiCastSearchRequest) Q(q string) ApiCastSearchRequest {
	r.q = &q
	return r
}

// Fid of the user whose casts you want to search
func (r ApiCastSearchRequest) AuthorFid(authorFid int32) ApiCastSearchRequest {
	r.authorFid = &authorFid
	return r
}

// Parent URL of the casts you want to search
func (r ApiCastSearchRequest) ParentUrl(parentUrl string) ApiCastSearchRequest {
	r.parentUrl = &parentUrl
	return r
}

// Channel ID of the casts you want to search
func (r ApiCastSearchRequest) ChannelId(channelId string) ApiCastSearchRequest {
	r.channelId = &channelId
	return r
}

func (r ApiCastSearchRequest) Limit(limit int32) ApiCastSearchRequest {
	r.limit = &limit
	return r
}

// Pagination cursor
func (r ApiCastSearchRequest) Cursor(cursor string) ApiCastSearchRequest {
	r.cursor = &cursor
	return r
}

func (r ApiCastSearchRequest) Execute() (*CastsSearchResponse, *http.Response, error) {
	return r.ApiService.CastSearchExecute(r)
}

/*
CastSearch Search for casts

Search for casts based on a query string, with optional AND filters

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCastSearchRequest
*/
func (a *CastAPIService) CastSearch(ctx context.Context) ApiCastSearchRequest {
	return ApiCastSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastsSearchResponse
func (a *CastAPIService) CastSearchExecute(r ApiCastSearchRequest) (*CastsSearchResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastsSearchResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.CastSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
	}
	if r.authorFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "author_fid", r.authorFid, "form", "")
	}
	if r.parentUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "parent_url", r.parentUrl, "form", "")
	}
	if r.channelId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "channel_id", r.channelId, "form", "")
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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
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

type ApiCastsRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	casts *string
	viewerFid *int32
	sortType *string
}

// API key required for authentication.
func (r ApiCastsRequest) ApiKey(apiKey string) ApiCastsRequest {
	r.apiKey = &apiKey
	return r
}

// Hashes of the cast to be retrived (Comma separated)
func (r ApiCastsRequest) Casts(casts string) ApiCastsRequest {
	r.casts = &casts
	return r
}

// adds viewer_context to cast object to show whether viewer has liked or recasted the cast.
func (r ApiCastsRequest) ViewerFid(viewerFid int32) ApiCastsRequest {
	r.viewerFid = &viewerFid
	return r
}

// Optional parameter to sort the casts based on different criteria
func (r ApiCastsRequest) SortType(sortType string) ApiCastsRequest {
	r.sortType = &sortType
	return r
}

func (r ApiCastsRequest) Execute() (*CastsResponse, *http.Response, error) {
	return r.ApiService.CastsExecute(r)
}

/*
Casts Gets information about an array of casts

Retrieve multiple casts using their respective hashes.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCastsRequest
*/
func (a *CastAPIService) Casts(ctx context.Context) ApiCastsRequest {
	return ApiCastsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastsResponse
func (a *CastAPIService) CastsExecute(r ApiCastsRequest) (*CastsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.Casts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/casts"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.casts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "casts", r.casts, "form", "")
	}
	if r.viewerFid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewer_fid", r.viewerFid, "form", "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_type", r.sortType, "form", "")
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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
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

type ApiComposerListRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	list *CastComposerType
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiComposerListRequest) ApiKey(apiKey string) ApiComposerListRequest {
	r.apiKey = &apiKey
	return r
}

// Type of list to fetch.
func (r ApiComposerListRequest) List(list CastComposerType) ApiComposerListRequest {
	r.list = &list
	return r
}

// Number of results to retrieve (default 25, max 25).
func (r ApiComposerListRequest) Limit(limit int32) ApiComposerListRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiComposerListRequest) Cursor(cursor string) ApiComposerListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiComposerListRequest) Execute() (*CastComposerActionsListResponse, *http.Response, error) {
	return r.ApiService.ComposerListExecute(r)
}

/*
ComposerList Fetches all composer actions on Warpcast

Fetches all composer actions on Warpcast. You can filter by top or featured.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiComposerListRequest
*/
func (a *CastAPIService) ComposerList(ctx context.Context) ApiComposerListRequest {
	return ApiComposerListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastComposerActionsListResponse
func (a *CastAPIService) ComposerListExecute(r ApiComposerListRequest) (*CastComposerActionsListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastComposerActionsListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.ComposerList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast/composer_actions/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.list != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "list", r.list, "form", "")
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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiDeleteCastRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	deleteCastReqBody *DeleteCastReqBody
}

// API key required for authentication.
func (r ApiDeleteCastRequest) ApiKey(apiKey string) ApiDeleteCastRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiDeleteCastRequest) DeleteCastReqBody(deleteCastReqBody DeleteCastReqBody) ApiDeleteCastRequest {
	r.deleteCastReqBody = &deleteCastReqBody
	return r
}

func (r ApiDeleteCastRequest) Execute() (*OperationResponse, *http.Response, error) {
	return r.ApiService.DeleteCastExecute(r)
}

/*
DeleteCast Delete a cast

Delete an existing cast. \
(In order to delete a cast `signer_uuid` must be approved)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteCastRequest
*/
func (a *CastAPIService) DeleteCast(ctx context.Context) ApiDeleteCastRequest {
	return ApiDeleteCastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OperationResponse
func (a *CastAPIService) DeleteCastExecute(r ApiDeleteCastRequest) (*OperationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OperationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.DeleteCast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
	// body params
	localVarPostBody = r.deleteCastReqBody
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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

type ApiPostCastRequest struct {
	ctx context.Context
	ApiService *CastAPIService
	apiKey *string
	postCastReqBody *PostCastReqBody
}

// API key required for authentication.
func (r ApiPostCastRequest) ApiKey(apiKey string) ApiPostCastRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiPostCastRequest) PostCastReqBody(postCastReqBody PostCastReqBody) ApiPostCastRequest {
	r.postCastReqBody = &postCastReqBody
	return r
}

func (r ApiPostCastRequest) Execute() (*PostCastResponse, *http.Response, error) {
	return r.ApiService.PostCastExecute(r)
}

/*
PostCast Posts a cast

Posts a cast or cast reply. Works with mentions and embeds.  
(In order to post a cast `signer_uuid` must be approved)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostCastRequest
*/
func (a *CastAPIService) PostCast(ctx context.Context) ApiPostCastRequest {
	return ApiPostCastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PostCastResponse
func (a *CastAPIService) PostCastExecute(r ApiPostCastRequest) (*PostCastResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostCastResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastAPIService.PostCast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/cast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.apiKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "simple", "")
	}
	// body params
	localVarPostBody = r.postCastReqBody
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorRes
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
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
