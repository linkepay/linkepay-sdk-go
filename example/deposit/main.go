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

	response, err := client.CreateMultipleDepositAddress(&types.CreateMultipleDepositAddressRequest{
		NetworkID: 4,
		Count:     2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
