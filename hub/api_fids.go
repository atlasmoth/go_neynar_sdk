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


// FIDsAPIService FIDsAPI service
type FIDsAPIService service

type ApiListFidsRequest struct {
	ctx context.Context
	ApiService *FIDsAPIService
	apiKey *string
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiListFidsRequest) ApiKey(apiKey string) ApiListFidsRequest {
	r.apiKey = &apiKey
	return r
}

// Maximum number of messages to return in a single response
func (r ApiListFidsRequest) PageSize(pageSize int32) ApiListFidsRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiListFidsRequest) Reverse(reverse bool) ApiListFidsRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiListFidsRequest) PageToken(pageToken string) ApiListFidsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListFidsRequest) Execute() (*FidsResponse, *http.Response, error) {
	return r.ApiService.ListFidsExecute(r)
}

/*
ListFids Get a list of all the FIDs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListFidsRequest
*/
func (a *FIDsAPIService) ListFids(ctx context.Context) ApiListFidsRequest {
	return ApiListFidsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FidsResponse
func (a *FIDsAPIService) ListFidsExecute(r ApiListFidsRequest) (*FidsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FidsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FIDsAPIService.ListFids")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/fids"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
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
