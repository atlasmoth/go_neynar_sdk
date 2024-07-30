package neynarsdk

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type UserService struct {
	client *Client
}

type SearchUserParams struct {
	Query     string
	ViewerFid int32
	Limit     int
	Cursor    string
}

type BulkUserParams struct {
	Fids      []int32
	ViewerFid int32
}

type PowerUserParams struct {
	ViewerFid int32
	Limit     int
	Cursor    string
}

type BulkUserByAddressParams struct {
	Addresses    []string
	AddressTypes []string
	ViewerFid    int32
}

type LookupUserByCustodyAddressParams struct {
	CustodyAddress string
}

type FollowUserParams struct {
	SignerUUID string
	TargetFids []int32
}

type UpdateUserParams struct {
	SignerUUID  string
	Bio         string
	PfpURL      string
	Username    string
	DisplayName string
}

type RegisterUserParams struct {
	Signature                   string
	Fid                         int32
	RequestedUserCustodyAddress string
	Deadline                    int64
	Fname                       string
}

type AddVerificationParams struct {
	SignerUUID   string
	Address      string
	BlockHash    string
	EthSignature string
}

type RemoveVerificationParams struct {
	SignerUUID string
	Address    string
}
type UserSearchResponse struct {
	Result struct {
		Users []SearchedUser `json:"users"`
		Next  NextCursor     `json:"next,omitempty"`
	} `json:"result"`
	ErrorResponse
}

type BulkUsersResponse struct {
	Users []User `json:"users"`
	ErrorResponse
}

type BulkUsersByAddressResponse map[string][]User

type UserResponse struct {
	User User `json:"user"`
	ErrorResponse
}

type BulkFollowResponse struct {
	Success bool             `json:"success"`
	Details []FollowResponse `json:"details"`
	ErrorResponse
}
type FollowResponse struct {
	Success   bool   `json:"success"`
	TargetFID int    `json:"target_fid"`
	Hash      string `json:"hash"`
}

type UserFIDResponse struct {
	FID int `json:"fid"`
	ErrorResponse
}

type OperationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	ErrorResponse
}

type RegisterUserResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Signer  Signer `json:"signer"`
	ErrorResponse
}

func (s *UserService) SearchUsers(ctx context.Context, params SearchUserParams) (UserSearchResponse, error) {
	var result UserSearchResponse
	if params.Query == "" {
		return result, &RequiredFieldError{Field: "Query"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/search"
	values := map[string]any{
		"q": params.Query,
	}

	if params.ViewerFid != 0 {
		values["viewer_fid"] = params.ViewerFid
	}
	if params.Limit > 0 {
		values["limit"] = params.Limit
	}
	if params.Cursor != "" {
		values["cursor"] = params.Cursor
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

func (s *UserService) BulkUsers(ctx context.Context, params BulkUserParams) (BulkUsersResponse, error) {
	var result BulkUsersResponse
	if len(params.Fids) == 0 {
		return result, &RequiredFieldError{Field: "Fids"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/bulk"
	values := map[string]any{
		"fids": strings.Trim(strings.Join(strings.Fields(fmt.Sprint(params.Fids)), ","), "[]"),
	}

	if params.ViewerFid != 0 {
		values["viewer_fid"] = params.ViewerFid
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

func (s *UserService) PowerUsers(ctx context.Context, params PowerUserParams) (UsersResponse, error) {
	var result UsersResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/power"
	values := make(map[string]any)

	if params.ViewerFid != 0 {
		values["viewer_fid"] = params.ViewerFid
	}
	if params.Limit > 0 {
		values["limit"] = params.Limit
	}
	if params.Cursor != "" {
		values["cursor"] = params.Cursor
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

func (s *UserService) BulkUsersByAddress(ctx context.Context, params BulkUserByAddressParams) (BulkUsersByAddressResponse, error) {
	var result BulkUsersByAddressResponse
	if len(params.Addresses) == 0 {
		return result, &RequiredFieldError{Field: "Addresses"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/bulk-by-address"
	values := map[string]any{
		"addresses": strings.Join(params.Addresses, ","),
	}

	if len(params.AddressTypes) > 0 {
		values["address_types"] = strings.Join(params.AddressTypes, ",")
	}
	if params.ViewerFid != 0 {
		values["viewer_fid"] = params.ViewerFid
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

func (s *UserService) LookupUserByCustodyAddress(ctx context.Context, params LookupUserByCustodyAddressParams) (UserResponse, error) {
	var result UserResponse
	if params.CustodyAddress == "" {
		return result, &RequiredFieldError{Field: "CustodyAddress"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/custody-address"
	values := map[string]any{
		"custody_address": params.CustodyAddress,
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

func (s *UserService) FollowUser(ctx context.Context, params FollowUserParams) (BulkFollowResponse, error) {
	var result BulkFollowResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if len(params.TargetFids) == 0 {
		return result, &RequiredFieldError{Field: "TargetFids"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/follow"
	body := map[string]any{
		"signer_uuid": params.SignerUUID,
		"target_fids": params.TargetFids,
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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

func (s *UserService) UnfollowUser(ctx context.Context, params FollowUserParams) (BulkFollowResponse, error) {
	var result BulkFollowResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if len(params.TargetFids) == 0 {
		return result, &RequiredFieldError{Field: "TargetFids"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/follow"
	body := map[string]any{
		"signer_uuid": params.SignerUUID,
		"target_fids": params.TargetFids,
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodDelete, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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

func (s *UserService) GetFreshFid(ctx context.Context) (UserFIDResponse, error) {
	var result UserFIDResponse

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/fid"
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
	} else {
		var errorResponse ErrorResponse
		err = s.client.HandleJsonResponse(resp, &errorResponse)
		if err != nil {
			return result, err
		}
		return result, &errorResponse
	}
}

func (s *UserService) UpdateUser(ctx context.Context, params UpdateUserParams) (OperationResponse, error) {
	var result OperationResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user"
	body := map[string]any{
		"signer_uuid": params.SignerUUID,
	}

	if params.Bio != "" {
		body["bio"] = params.Bio
	}
	if params.PfpURL != "" {
		body["pfp_url"] = params.PfpURL
	}
	if params.Username != "" {
		body["username"] = params.Username
	}
	if params.DisplayName != "" {
		body["display_name"] = params.DisplayName
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPatch, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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

func (s *UserService) RegisterUser(ctx context.Context, params RegisterUserParams) (RegisterUserResponse, error) {
	var result RegisterUserResponse
	if params.Signature == "" {
		return result, &RequiredFieldError{Field: "Signature"}
	}
	if params.Fid == 0 {
		return result, &RequiredFieldError{Field: "Fid"}
	}
	if params.RequestedUserCustodyAddress == "" {
		return result, &RequiredFieldError{Field: "RequestedUserCustodyAddress"}
	}
	if params.Deadline == 0 {
		return result, &RequiredFieldError{Field: "Deadline"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user"
	body := map[string]any{
		"signature":                      params.Signature,
		"fid":                            params.Fid,
		"requested_user_custody_address": params.RequestedUserCustodyAddress,
		"deadline":                       params.Deadline,
	}

	if params.Fname != "" {
		body["fname"] = params.Fname
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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

func (s *UserService) AddVerification(ctx context.Context, params AddVerificationParams) (OperationResponse, error) {
	var result OperationResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if params.Address == "" {
		return result, &RequiredFieldError{Field: "Address"}
	}
	if params.BlockHash == "" {
		return result, &RequiredFieldError{Field: "BlockHash"}
	}
	if params.EthSignature == "" {
		return result, &RequiredFieldError{Field: "EthSignature"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/verification"
	body := map[string]any{
		"signer_uuid":   params.SignerUUID,
		"address":       params.Address,
		"block_hash":    params.BlockHash,
		"eth_signature": params.EthSignature,
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodPost, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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

func (s *UserService) RemoveVerification(ctx context.Context, params RemoveVerificationParams) (OperationResponse, error) {
	var result OperationResponse
	if params.SignerUUID == "" {
		return result, &RequiredFieldError{Field: "SignerUUID"}
	}
	if params.Address == "" {
		return result, &RequiredFieldError{Field: "Address"}
	}

	baseURL := s.client.BaseURL.String() + "v2/farcaster/user/verification"
	body := map[string]any{
		"signer_uuid": params.SignerUUID,
		"address":     params.Address,
	}

	resp, err := s.client.HandleJsonRequest(ctx, http.MethodDelete, baseURL, body, nil)

	if err != nil {
		return result, err
	}
	if resp.StatusCode == http.StatusOK {
		err = s.client.HandleJsonResponse(resp, &result)
		if err != nil {
			return result, err
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
