package main

import (
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
	callbackData := map[string]interface{}{
		"deposit": map[string]interface{}{
			// deposit data
		},
		"withdrawal": map[string]interface{}{
			// withdrawal data
		},
	}
	callbackDataSignature := "123"
	ok, err := client.VerifyPlatformSignature(client.GetPlatformPublicKey(), callbackData, callbackDataSignature)
	fmt.Println(ok)
	if err != nil {
		panic(err)
	}
}
