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


// MuteAPIService MuteAPI service
type MuteAPIService service

type ApiAddMuteRequest struct {
	ctx context.Context
	ApiService *MuteAPIService
	apiKey *string
	muteReqBody *MuteReqBody
}

// API key required for authentication.
func (r ApiAddMuteRequest) ApiKey(apiKey string) ApiAddMuteRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiAddMuteRequest) MuteReqBody(muteReqBody MuteReqBody) ApiAddMuteRequest {
	r.muteReqBody = &muteReqBody
	return r
}

func (r ApiAddMuteRequest) Execute() (*MuteResponse, *http.Response, error) {
	return r.ApiService.AddMuteExecute(r)
}

/*
AddMute Adds a mute for a fid

Adds a mute for a given fid. This is a whitelisted API, reach out if you want access.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddMuteRequest
*/
func (a *MuteAPIService) AddMute(ctx context.Context) ApiAddMuteRequest {
	return ApiAddMuteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MuteResponse
func (a *MuteAPIService) AddMuteExecute(r ApiAddMuteRequest) (*MuteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MuteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MuteAPIService.AddMute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/mute"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.muteReqBody == nil {
		return localVarReturnValue, nil, reportError("muteReqBody is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "")
	// body params
	localVarPostBody = r.muteReqBody
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

type ApiDeleteMuteRequest struct {
	ctx context.Context
	ApiService *MuteAPIService
	apiKey *string
	muteReqBody *MuteReqBody
}

// API key required for authentication.
func (r ApiDeleteMuteRequest) ApiKey(apiKey string) ApiDeleteMuteRequest {
	r.apiKey = &apiKey
	return r
}

func (r ApiDeleteMuteRequest) MuteReqBody(muteReqBody MuteReqBody) ApiDeleteMuteRequest {
	r.muteReqBody = &muteReqBody
	return r
}

func (r ApiDeleteMuteRequest) Execute() (*MuteResponse, *http.Response, error) {
	return r.ApiService.DeleteMuteExecute(r)
}

/*
DeleteMute Deletes a mute for a fid

Deletes a mute for a given fid. This is a whitelisted API, reach out if you want access.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteMuteRequest
*/
func (a *MuteAPIService) DeleteMute(ctx context.Context) ApiDeleteMuteRequest {
	return ApiDeleteMuteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MuteResponse
func (a *MuteAPIService) DeleteMuteExecute(r ApiDeleteMuteRequest) (*MuteResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MuteResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MuteAPIService.DeleteMute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/mute"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.muteReqBody == nil {
		return localVarReturnValue, nil, reportError("muteReqBody is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "")
	// body params
	localVarPostBody = r.muteReqBody
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

type ApiMuteListRequest struct {
	ctx context.Context
	ApiService *MuteAPIService
	apiKey *string
	fid *int32
	limit *int32
	cursor *string
}

// API key required for authentication.
func (r ApiMuteListRequest) ApiKey(apiKey string) ApiMuteListRequest {
	r.apiKey = &apiKey
	return r
}

// The user&#39;s fid (identifier)
func (r ApiMuteListRequest) Fid(fid int32) ApiMuteListRequest {
	r.fid = &fid
	return r
}

// Number of results to retrieve (default 20, max 100).
func (r ApiMuteListRequest) Limit(limit int32) ApiMuteListRequest {
	r.limit = &limit
	return r
}

// Pagination cursor.
func (r ApiMuteListRequest) Cursor(cursor string) ApiMuteListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiMuteListRequest) Execute() (*MuteListResponse, *http.Response, error) {
	return r.ApiService.MuteListExecute(r)
}

/*
MuteList Get fids that a user has muted

Fetches all fids that a user has muted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMuteListRequest
*/
func (a *MuteAPIService) MuteList(ctx context.Context) ApiMuteListRequest {
	return ApiMuteListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MuteListResponse
func (a *MuteAPIService) MuteListExecute(r ApiMuteListRequest) (*MuteListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MuteListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MuteAPIService.MuteList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/farcaster/mute/list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "")
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 20
		r.limit = &defaultValue
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "api_key", r.apiKey, "")
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
