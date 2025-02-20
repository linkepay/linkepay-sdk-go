package operations

import (
	"encoding/json"
	"fmt"

	"github.com/linkepay/linkepay-sdk-go/types"
	"github.com/linkepay/linkepay-sdk-go/utils"
)

func RequestWithdrawal(client *types.Client, data types.RequestWithdrawalRequest) (map[string]interface{}, error) {
	projectUID := client.Config.ProjectID

	clientWithdrawRequest := map[string]interface{}{
		"uid":        data.UID, // withdrawal uid
		"asset_id":   data.AssetID,
		"amount":     data.Amount,
		"to_address": data.Destination,
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
	}

	body, err := utils.Request(reqConfig)
	if err != nil {
		return nil, fmt.Errorf("withdrawal request failed: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	result["hasRequestedPayVendor"] = true
	return result, nil
}
