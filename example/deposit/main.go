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
		ApiKey:               "", // Required for GetDeposits / GetWithdrawals
	})

	// Create deposit addresses (uses X-Signature auth)
	response, err := client.CreateMultipleDepositAddress(&types.CreateMultipleDepositAddressRequest{
		NetworkID: 4,
		Count:     2,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(response)

	// Get deposit history (uses X-API-Key auth)
	deposits, err := client.GetDeposits(&types.GetDepositsRequest{
		Page:     1,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(deposits)
}
