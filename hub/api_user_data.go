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


// UserDataAPIService UserDataAPI service
type UserDataAPIService service

type ApiGetUserDataByFidRequest struct {
	ctx context.Context
	ApiService *UserDataAPIService
	apiKey *string
	fid *int32
	userDataType *UserDataType
	pageSize *int32
	reverse *bool
	pageToken *string
}

// API key required for authentication.
func (r ApiGetUserDataByFidRequest) ApiKey(apiKey string) ApiGetUserDataByFidRequest {
	r.apiKey = &apiKey
	return r
}

// The FID that&#39;s being requested
func (r ApiGetUserDataByFidRequest) Fid(fid int32) ApiGetUserDataByFidRequest {
	r.fid = &fid
	return r
}

// The type of user data, either as a numerical value or type string. If this is omitted, all user data for the FID is returned
func (r ApiGetUserDataByFidRequest) UserDataType(userDataType UserDataType) ApiGetUserDataByFidRequest {
	r.userDataType = &userDataType
	return r
}

// Maximum number of messages to return in a single response
func (r ApiGetUserDataByFidRequest) PageSize(pageSize int32) ApiGetUserDataByFidRequest {
	r.pageSize = &pageSize
	return r
}

// Reverse the sort order, returning latest messages first
func (r ApiGetUserDataByFidRequest) Reverse(reverse bool) ApiGetUserDataByFidRequest {
	r.reverse = &reverse
	return r
}

// The page token returned by the previous query, to fetch the next page. If this parameter is empty, fetch the first page
func (r ApiGetUserDataByFidRequest) PageToken(pageToken string) ApiGetUserDataByFidRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiGetUserDataByFidRequest) Execute() (*GetUserDataByFid200Response, *http.Response, error) {
	return r.ApiService.GetUserDataByFidExecute(r)
}

/*
GetUserDataByFid Get UserData for a FID.

**Note:** one of two different response schemas is returned based on whether the caller provides the `user_data_type` parameter. If included, a single `UserDataAdd` message is returned (or a `not_found` error). If omitted, a paginated list of `UserDataAdd` messages is returned instead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserDataByFidRequest
*/
func (a *UserDataAPIService) GetUserDataByFid(ctx context.Context) ApiGetUserDataByFidRequest {
	return ApiGetUserDataByFidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetUserDataByFid200Response
func (a *UserDataAPIService) GetUserDataByFidExecute(r ApiGetUserDataByFidRequest) (*GetUserDataByFid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetUserDataByFid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserDataAPIService.GetUserDataByFid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/userDataByFid"

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
	if r.userDataType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_data_type", r.userDataType, "form", "")
	} else {
		var defaultValue UserDataType = "USER_DATA_TYPE_PFP"
		r.userDataType = &defaultValue
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
