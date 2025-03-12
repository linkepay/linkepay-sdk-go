package operations

import (
	"encoding/json"
	"fmt"

	"github.com/linkepay/linkepay-sdk-go/types"
	"github.com/linkepay/linkepay-sdk-go/utils"
)

func GetDepositAddress(client *types.Client, req *types.GetDepositAddressRequest) (*types.GetDepositAddressResponse, error) {

	// Get project UID from config
	projectUID := client.Config.ProjectID

	if projectUID == "" {
		return nil, fmt.Errorf("projectUID is required")
	}

	if req.UserUID == "" {
		return nil, fmt.Errorf("userUID is required")
	}

	if req.NetworkID == 0 {
		return nil, fmt.Errorf("networkID is required")
	}

	// Prepare request data
	// reqData := map[string]interface{}{
	// 	"user_uid":   userUID,
	// 	"network_id": networkID,
	// }

	// Sign request
	km := utils.NewKeyManager()
	km.LoadKeys(&client.Config)

	signedData, err := km.SignRequest(fmt.Sprintf("/api/v1/client/project/%s/user/%s/deposit-address/%d", projectUID, req.UserUID, req.NetworkID), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %v", err)
	}

	// Make request
	reqConfig := utils.RequestConfig{
		Method:  "GET",
		BaseURL: client.Config.BaseURL,
		Path:    fmt.Sprintf("/api/v1/client/project/%s/user/%s/deposit-address/%d", projectUID, req.UserUID, req.NetworkID),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Signature":  signedData,
		},
	}

	body, err := utils.Request(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	// Parse response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Get address from response
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("GenerateDepositAddress invalid response format: %v", response)
	}

	address, ok := data["address"].(string)
	if !ok || address == "" {
		return nil, fmt.Errorf("address not returned from server: %v", response)
	}
	return &types.GetDepositAddressResponse{
		Message: "Address fetched successfully",
		Address: address,
		UID:     req.UserUID,
	}, nil
}

func CreateDepositAddress(client *types.Client, req *types.CreateDepositAddressRequest) (*types.CreateDepositAddressResponse, error) {

	// Get project UID from config
	projectUID := client.Config.ProjectID

	if projectUID == "" {
		return nil, fmt.Errorf("projectUID is required")
	}

	if req.UserUID == "" {
		return nil, fmt.Errorf("userUID is required")
	}

	if req.NetworkID == 0 {
		return nil, fmt.Errorf("networkID is required")
	}

	// Prepare request data
	reqData := map[string]interface{}{
		"user_uid":   req.UserUID,
		"network_id": req.NetworkID,
	}

	// Sign request
	km := utils.NewKeyManager()
	km.LoadKeys(&client.Config)

	signedData, err := km.SignRequest(fmt.Sprintf("/api/v1/client/project/%s/user/generate-deposit-address", projectUID), reqData)

	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %v", err)
	}

	// Make request
	reqConfig := utils.RequestConfig{
		Method:  "POST",
		BaseURL: client.Config.BaseURL,
		Path:    fmt.Sprintf("/api/v1/client/project/%s/user/generate-deposit-address", projectUID),
		Body:    reqData,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Signature":  signedData,
		},
	}

	body, err := utils.Request(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	// Parse response
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// Get address from response
	data, ok := response["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("GenerateDepositAddress invalid response format: %v", response)
	}

	address, ok := data["address"].(string)
	if !ok || address == "" {
		return nil, fmt.Errorf("address not returned from server: %v", response)
	}
	return &types.CreateDepositAddressResponse{
		Message: "Address created successfully",
		Address: address,
		UID:     req.UserUID,
	}, nil
}

func CreateMultipleDepositAddress(client *types.Client, req *types.CreateMultipleDepositAddressRequest) (*types.CreateMultipleDepositAddressResponse, error) {

	// Get project UID from config
	projectUID := client.Config.ProjectID

	if projectUID == "" {
		return nil, fmt.Errorf("projectUID is required")
	}

	if req.NetworkID == 0 {
		return nil, fmt.Errorf("networkID is required")
	}

	if req.Count <= 0 {
		return nil, fmt.Errorf("count is required and should be greater than 0")
	}

	if req.Count > 100 {
		return nil, fmt.Errorf("count must be less than 100")
	}

	// Prepare request data
	reqData := map[string]interface{}{
		"network_id": req.NetworkID,
		"count":      req.Count,
	}

	// Sign request
	km := utils.NewKeyManager()
	km.LoadKeys(&client.Config)

	signedData, err := km.SignRequest(fmt.Sprintf("/api/v1/client/project/%s/user/generate-deposit-addresses", projectUID), reqData)

	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %v", err)
	}

	// Make request
	reqConfig := utils.RequestConfig{
		Method:  "POST",
		BaseURL: client.Config.BaseURL,
		Path:    fmt.Sprintf("/api/v1/client/project/%s/user/generate-deposit-address", projectUID),
		Body:    reqData,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Signature":  signedData,
		},
	}

	body, err := utils.Request(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}

	// Parse response
	var response types.CreateMultipleDepositAddressResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &response, nil
}
