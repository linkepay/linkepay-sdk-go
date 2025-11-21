package main

import (
	"encoding/json"
	"fmt"

	"github.com/linkepay/linkepay-sdk-go"
	"github.com/linkepay/linkepay-sdk-go/types"
)

func main() {
	client := linkepay.NewClient(&types.Config{
		BaseURL:              "",
		ProjectID:            "",
		PublicKey:            "",
		PrivateKey:           "",
		PayPlatformPublicKey: "",
	})

	// Raw JSON callback data
	rawJSON := "{\n  \"verify_data\": {\n    \"project_uid\": \"ROhTAi9AEd\",\n    \"from_address\": \"0x8d36a471cb5cfc32d3e08a8dce748c64be9e7308\",\n    \"to_address\": \"0xc5af7534b80a35f2e820d9e1e8dbce789600889f\",\n    \"address_uid\": \"nT6Evp\",\n    \"amount\": \"1000000000000000000\",\n    \"float_amount\": \"1\",\n    \"decimal\": 18,\n    \"tx_hash\": \"0x1bf1fff17f9071ea10f60a6a5f5a4f1e855a2056aa13e832101fa30cf6ca62d4\",\n    \"status\": \"confirmed\",\n    \"asset_id\": 2,\n    \"asset_name\": \"USDT\",\n    \"network_id\": 56,\n    \"network_name\": \"BSC Mainnet\",\n    \"confirmed\": true,\n    \"confirmed_at\": \"2025-11-21T03:07:09.15083Z\",\n    \"uid\": \"0x1bf1fff17f9071ea10f60a6a5f5a4f1e855a2056aa13e832101fa30cf6ca62d4\",\n    \"type\": \"deposit\"\n  },\n  \"sig\": \"bdd75214e4f1c2c31eeca324b7a0b76f04f906e8bf7f9a001d21492c033c89ba16b9ca10a752e936327497bd8d2c1498fd08e688fb03c46c47dd518b63e0744f00\"\n}"

	var callbackData types.CallbackRequestDataWithSig
	if err := json.Unmarshal([]byte(rawJSON), &callbackData); err != nil {
		panic(fmt.Sprintf("failed to parse JSON: %v", err))
	}

	// Debug: print the data being verified
	fmt.Printf("Data: %+v\n", callbackData.Data)
	fmt.Printf("Sig: %s\n", callbackData.Sig)

	// Verify signature directly
	ok, err := client.VerifyPlatformSignature(client.GetPlatformPublicKey(), callbackData.Data, callbackData.Sig)
	if err != nil {
		fmt.Printf("Verification error: %v\n", err)
	}
	fmt.Printf("Signature valid: %v\n", ok)

	// Also try ParseCallbackData
	response, err := client.ParseCallbackData(callbackData)
	fmt.Printf("Response: %+v\n", response)
	if err != nil {
		fmt.Printf("ParseCallbackData error: %v\n", err)
	}
}
