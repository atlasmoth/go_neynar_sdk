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


// CastsAPIService CastsAPI service
type CastsAPIService service

type ApiGetCastByIdRequest struct {
	ctx context.Context
	ApiService *CastsAPIService
	apiKey *string
	fid *int32
	hash *string
}

// API key required for authentication.
func (r ApiGetCastByIdRequest) ApiKey(apiKey string) ApiGetCastByIdRequest {
	r.apiKey = &apiKey
	return r
}

// The FID of the cast&#39;s creator
func (r ApiGetCastByIdRequest) Fid(fid int32) ApiGetCastByIdRequest {
	r.fid = &fid
	return r
}

// The cast&#39;s hash
func (r ApiGetCastByIdRequest) Hash(hash string) ApiGetCastByIdRequest {
	r.hash = &hash
	return r
}

func (r ApiGetCastByIdRequest) Execute() (*CastAdd, *http.Response, error) {
	return r.ApiService.GetCastByIdExecute(r)
}

/*
GetCastById Get a cast by its FID and Hash.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCastByIdRequest
*/
func (a *CastsAPIService) GetCastById(ctx context.Context) ApiGetCastByIdRequest {
	return ApiGetCastByIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CastAdd
func (a *CastsAPIService) GetCastByIdExecute(r ApiGetCastByIdRequest) (*CastAdd, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CastAdd
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastsAPIService.GetCastById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/castById"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.hash == nil {
		return localVarReturnValue, nil, reportError("hash is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "form", "")
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

type ApiListCastsByFidRequest struct {
	ctx context.Context
	ApiService *CastsAPIService
	apiKey *string
	fid *int32
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListCastsByFidRequest) ApiKey(apiKey string) ApiListCastsByFidRequest {
	r.apiKey = &apiKey
	return r
}

// The FID of the casts&#39; creator
func (r ApiListCastsByFidRequest) Fid(fid int32) ApiListCastsByFidRequest {
	r.fid = &fid
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListCastsByFidRequest) PageSize(pageSize int32) ApiListCastsByFidRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListCastsByFidRequest) Reverse(reverse bool) ApiListCastsByFidRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListCastsByFidRequest) PageToken(pageToken string) ApiListCastsByFidRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListCastsByFidRequest) Execute() (*ListCastsByFid200Response, *http.Response, error) {
	return r.ApiService.ListCastsByFidExecute(r)
}

/*
ListCastsByFid Fetch all casts authored by an FID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCastsByFidRequest
*/
func (a *CastsAPIService) ListCastsByFid(ctx context.Context) ApiListCastsByFidRequest {
	return ApiListCastsByFidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCastsByFid200Response
func (a *CastsAPIService) ListCastsByFidExecute(r ApiListCastsByFidRequest) (*ListCastsByFid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCastsByFid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastsAPIService.ListCastsByFid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/castsByFid"

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

type ApiListCastsByMentionRequest struct {
	ctx context.Context
	ApiService *CastsAPIService
	apiKey *string
	fid *int32
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListCastsByMentionRequest) ApiKey(apiKey string) ApiListCastsByMentionRequest {
	r.apiKey = &apiKey
	return r
}

// The FID that is mentioned in a cast
func (r ApiListCastsByMentionRequest) Fid(fid int32) ApiListCastsByMentionRequest {
	r.fid = &fid
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListCastsByMentionRequest) PageSize(pageSize int32) ApiListCastsByMentionRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListCastsByMentionRequest) Reverse(reverse bool) ApiListCastsByMentionRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListCastsByMentionRequest) PageToken(pageToken string) ApiListCastsByMentionRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListCastsByMentionRequest) Execute() (*ListCastsByFid200Response, *http.Response, error) {
	return r.ApiService.ListCastsByMentionExecute(r)
}

/*
ListCastsByMention Fetch all casts that mention an FID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCastsByMentionRequest
*/
func (a *CastsAPIService) ListCastsByMention(ctx context.Context) ApiListCastsByMentionRequest {
	return ApiListCastsByMentionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCastsByFid200Response
func (a *CastsAPIService) ListCastsByMentionExecute(r ApiListCastsByMentionRequest) (*ListCastsByFid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCastsByFid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastsAPIService.ListCastsByMention")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/castsByMention"

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

type ApiListCastsByParentRequest struct {
	ctx context.Context
	ApiService *CastsAPIService
	apiKey *string
	fid *int32
	hash *string
	url *string
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListCastsByParentRequest) ApiKey(apiKey string) ApiListCastsByParentRequest {
	r.apiKey = &apiKey
	return r
}

// The FID of the parent cast
func (r ApiListCastsByParentRequest) Fid(fid int32) ApiListCastsByParentRequest {
	r.fid = &fid
	return r
}

// The parent cast&#39;s hash
func (r ApiListCastsByParentRequest) Hash(hash string) ApiListCastsByParentRequest {
	r.hash = &hash
	return r
}

func (r ApiListCastsByParentRequest) Url(url string) ApiListCastsByParentRequest {
	r.url = &url
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListCastsByParentRequest) PageSize(pageSize int32) ApiListCastsByParentRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListCastsByParentRequest) Reverse(reverse bool) ApiListCastsByParentRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListCastsByParentRequest) PageToken(pageToken string) ApiListCastsByParentRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListCastsByParentRequest) Execute() (*ListCastsByFid200Response, *http.Response, error) {
	return r.ApiService.ListCastsByParentExecute(r)
}

/*
ListCastsByParent Fetch all casts by parent cast's FID and Hash OR by the parent's URL

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListCastsByParentRequest
*/
func (a *CastsAPIService) ListCastsByParent(ctx context.Context) ApiListCastsByParentRequest {
	return ApiListCastsByParentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListCastsByFid200Response
func (a *CastsAPIService) ListCastsByParentExecute(r ApiListCastsByParentRequest) (*ListCastsByFid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListCastsByFid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CastsAPIService.ListCastsByParent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/castsByParent"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}

	if r.fid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	}
	if r.hash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "form", "")
	}
	if r.url != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "url", r.url, "form", "")
	}
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