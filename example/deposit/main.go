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
	deposit, err := client.GetDepositAddress(&types.GetDepositAddressRequest{
		NetworkID: 1,
		UserUID:   "user_uid",
	})
	fmt.Println(deposit)
	if err != nil {
		panic(err)
	}
}
