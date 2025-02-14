package main

import (
	"fmt"

	"github.com/linkepay/linkepay-sdk-go"
	"github.com/linkepay/linkepay-sdk-go/types"
)

func main() {
	linkepay.SetClient(&types.Config{
		BaseURL:              "",
		ProjectID:            "",
		PublicKey:            "",
		PrivateKey:           "",
		PayPlatformPublicKey: "",
	})
	deposit, err := linkepay.GetDepositAddress(&types.GetDepositAddressRequest{
		NetworkID: 1,
		UserUID:   "user_uid",
	})
	fmt.Println(deposit)
	if err != nil {
		panic(err)
	}
}
