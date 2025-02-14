package main

import (
	"fmt"

	"github.com/linkepay/linkepay-sdk-go"
	"github.com/linkepay/linkepay-sdk-go/operations"
	"github.com/linkepay/linkepay-sdk-go/types"
)

func main() {
	client := linkepay.NewClient(&types.Config{
		BaseURL: "",
	})

	withdrawal, err := operations.CreateWithdrawal(client, &types.WithdrawalRequest{
		Amount:   "100.00",
		Currency: "USD",
	})
	fmt.Println(withdrawal)
	if err != nil {
		panic(err)
	}
}
