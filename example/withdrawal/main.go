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
	withdrawal, err := client.RequestWithdrawal(types.RequestWithdrawalRequest{
		Amount:      100.00,
		Destination: "destination_address",
		NetworkID:   1,
		UserUID:     "user_uid",
		UID:         "withdrawal_uid",
		AssetID:     2,
	})
	fmt.Println(withdrawal)
	if err != nil {
		panic(err)
	}
}
