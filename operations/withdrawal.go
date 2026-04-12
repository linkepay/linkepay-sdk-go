package operations

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/linkepay/linkepay-sdk-go/types"
	"github.com/linkepay/linkepay-sdk-go/utils"
)

func GetWithdrawals(client *types.Client, req *types.GetWithdrawalsRequest) (*types.GetWithdrawalsResponse, error) {
	projectUID := client.Config.ProjectID
	if projectUID == "" {
		return nil, fmt.Errorf("projectUID is required")
	}

	if client.Config.ApiKey == "" {
		return nil, fmt.Errorf("apiKey is required for GetWithdrawals")
	}

	params := url.Values{}
	if req != nil {
		if req.Page > 0 {
			params.Set("page", fmt.Sprintf("%d", req.Page))
		}
		if req.PageSize > 0 {
			params.Set("page_size", fmt.Sprintf("%d", req.PageSize))
		}
	}

	reqConfig := utils.RequestConfig{
		Method:  "GET",
		BaseURL: client.Config.BaseURL,
		Path:    fmt.Sprintf("/api/v1/client/project/%s/user/withdrawals", projectUID),
		Params:  params,
		Headers: map[string]string{
			"X-API-Key": client.Config.ApiKey,
		},
		Timeout: client.Config.Timeout,
	}

	body, err := utils.Request(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	var response types.GetWithdrawalsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &response, nil
}

func RequestWithdrawal(client *types.Client, data types.RequestWithdrawalRequest) (*types.RequestWithdrawalResponse, error) {
	projectUID := client.Config.ProjectID

	clientWithdrawRequest := map[string]interface{}{
		"uid":        data.UID, // withdrawal uid
		"asset_id":   data.AssetID,
		"amount":     strconv.FormatFloat(data.Amount, 'f', -1, 64),
		"to_address": data.ToAddress,
		"network_id": data.NetworkID,
	}

	// Sign request
	km := utils.NewKeyManager()
	km.LoadKeys(&client.Config)

	signedData, err := km.SignRequest(fmt.Sprintf("/api/v1/client/project/%s/withdraw", projectUID), clientWithdrawRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	// Make request
	reqConfig := utils.RequestConfig{
		Method:  "POST",
		BaseURL: client.Config.BaseURL,
		Path:    fmt.Sprintf("/api/v1/client/project/%s/withdraw", projectUID),
		Body:    clientWithdrawRequest,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Signature":  signedData,
		},
		Timeout: client.Config.Timeout,
	}

	body, err := utils.Request(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("withdrawal request failed: %w", err)
	}

	var result types.RequestWithdrawalResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response, body: %s, error: %w", string(body), err)
	}

	return &result, nil
}
