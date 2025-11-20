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

	res, e := client.RequestWithdrawal(types.RequestWithdrawalRequest{
		Amount:    100.00,
		AssetID:   2,
		NetworkID: 97,
		UserUID:   "user_uid",
		UID:       "withdrawal_uid",
		ToAddress: "0x4b2048ede5ca5d962795fec5edf8b41a860e2d4d",
	})
	fmt.Println(res, e)
	fmt.Println(fmt.Sprintf("RequestWithdrawal:%s", "withdrawal_uid"), map[string]interface{}{
		"input data": types.RequestWithdrawalRequest{
			Amount:    100.00,
			AssetID:   2,
			NetworkID: 97,
			UserUID:   "user_uid",
			UID:       "withdrawal_uid",
			ToAddress: "0x4b2048ede5ca5d962795fec5edf8b41a860e2d4d",
		},
		"response": res,
		"error":    e,
	})
}
