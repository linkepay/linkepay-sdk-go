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

	resp, _ := client.CreateDepositAddress(&types.CreateDepositAddressRequest{
		NetworkID: 1,
		UserUID:   "user_uid",
	})
	fmt.Println(resp)

	resp2, _ := client.GetDepositAddress(&types.GetDepositAddressRequest{
		NetworkID: 1,
		UserUID:   "user_uid",
	})
	fmt.Println(resp2)
}
