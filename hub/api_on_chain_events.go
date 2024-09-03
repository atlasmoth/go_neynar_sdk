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


// OnChainEventsAPIService OnChainEventsAPI service
type OnChainEventsAPIService service

type ApiGetOnChainIdRegistrationByAddressRequest struct {
	ctx context.Context
	ApiService *OnChainEventsAPIService
	apiKey *string
	address *string
}

// API key required for authentication.
func (r ApiGetOnChainIdRegistrationByAddressRequest) ApiKey(apiKey string) ApiGetOnChainIdRegistrationByAddressRequest {
	r.apiKey = &apiKey
	return r
}

// The ETH address being requested
func (r ApiGetOnChainIdRegistrationByAddressRequest) Address(address string) ApiGetOnChainIdRegistrationByAddressRequest {
	r.address = &address
	return r
}

func (r ApiGetOnChainIdRegistrationByAddressRequest) Execute() (*OnChainEventIdRegister, *http.Response, error) {
	return r.ApiService.GetOnChainIdRegistrationByAddressExecute(r)
}

/*
GetOnChainIdRegistrationByAddress Get an on chain ID Registry Event for a given Address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOnChainIdRegistrationByAddressRequest
*/
func (a *OnChainEventsAPIService) GetOnChainIdRegistrationByAddress(ctx context.Context) ApiGetOnChainIdRegistrationByAddressRequest {
	return ApiGetOnChainIdRegistrationByAddressRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OnChainEventIdRegister
func (a *OnChainEventsAPIService) GetOnChainIdRegistrationByAddressExecute(r ApiGetOnChainIdRegistrationByAddressRequest) (*OnChainEventIdRegister, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OnChainEventIdRegister
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEventsAPIService.GetOnChainIdRegistrationByAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/onChainIdRegistryEventByAddress"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.address == nil {
		return localVarReturnValue, nil, reportError("address is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "form", "")
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

type ApiListOnChainEventsByFidRequest struct {
	ctx context.Context
	ApiService *OnChainEventsAPIService
	apiKey *string
	fid *int32
	eventType *OnChainEventType
}

// API key required for authentication.
func (r ApiListOnChainEventsByFidRequest) ApiKey(apiKey string) ApiListOnChainEventsByFidRequest {
	r.apiKey = &apiKey
	return r
}

// The FID being requested
func (r ApiListOnChainEventsByFidRequest) Fid(fid int32) ApiListOnChainEventsByFidRequest {
	r.fid = &fid
	return r
}

// The numeric of string value of the event type being requested.
func (r ApiListOnChainEventsByFidRequest) EventType(eventType OnChainEventType) ApiListOnChainEventsByFidRequest {
	r.eventType = &eventType
	return r
}

func (r ApiListOnChainEventsByFidRequest) Execute() (*ListOnChainEventsByFid200Response, *http.Response, error) {
	return r.ApiService.ListOnChainEventsByFidExecute(r)
}

/*
ListOnChainEventsByFid Get a list of on-chain events provided by an FID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOnChainEventsByFidRequest
*/
func (a *OnChainEventsAPIService) ListOnChainEventsByFid(ctx context.Context) ApiListOnChainEventsByFidRequest {
	return ApiListOnChainEventsByFidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOnChainEventsByFid200Response
func (a *OnChainEventsAPIService) ListOnChainEventsByFidExecute(r ApiListOnChainEventsByFidRequest) (*ListOnChainEventsByFid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOnChainEventsByFid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEventsAPIService.ListOnChainEventsByFid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/onChainEventsByFid"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiKey == nil {
		return localVarReturnValue, nil, reportError("apiKey is required and must be specified")
	}
	if r.fid == nil {
		return localVarReturnValue, nil, reportError("fid is required and must be specified")
	}
	if r.eventType == nil {
		return localVarReturnValue, nil, reportError("eventType is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fid", r.fid, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "event_type", r.eventType, "form", "")
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

type ApiListOnChainSignersByFidRequest struct {
	ctx context.Context
	ApiService *OnChainEventsAPIService
	apiKey *string
	fid *int32
	signer *string
}

// API key required for authentication.
func (r ApiListOnChainSignersByFidRequest) ApiKey(apiKey string) ApiListOnChainSignersByFidRequest {
	r.apiKey = &apiKey
	return r
}

// The FID being requested
func (r ApiListOnChainSignersByFidRequest) Fid(fid int32) ApiListOnChainSignersByFidRequest {
	r.fid = &fid
	return r
}

// The optional key of signer
func (r ApiListOnChainSignersByFidRequest) Signer(signer string) ApiListOnChainSignersByFidRequest {
	r.signer = &signer
	return r
}

func (r ApiListOnChainSignersByFidRequest) Execute() (*ListOnChainSignersByFid200Response, *http.Response, error) {
	return r.ApiService.ListOnChainSignersByFidExecute(r)
}

/*
ListOnChainSignersByFid Get a list of signers provided by an FID

**Note:** one of two different response schemas is returned based on whether the caller provides the `signer` parameter. If included, a single `OnChainEventSigner` message is returned (or a `not_found` error). If omitted, a non-paginated list of `OnChainEventSigner` messages is returned instead

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOnChainSignersByFidRequest
*/
func (a *OnChainEventsAPIService) ListOnChainSignersByFid(ctx context.Context) ApiListOnChainSignersByFidRequest {
	return ApiListOnChainSignersByFidRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOnChainSignersByFid200Response
func (a *OnChainEventsAPIService) ListOnChainSignersByFidExecute(r ApiListOnChainSignersByFidRequest) (*ListOnChainSignersByFid200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOnChainSignersByFid200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OnChainEventsAPIService.ListOnChainSignersByFid")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/onChainSignersByFid"

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
	if r.signer != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "signer", r.signer, "form", "")
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