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

	resp, _ := client.CreateMultipleDepositAddress(&types.CreateMultipleDepositAddressRequest{
		NetworkID: 11155111,
		Count:     2,
	})
	fmt.Println(resp)

}
