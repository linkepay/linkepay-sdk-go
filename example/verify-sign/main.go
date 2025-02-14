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
	ok, err := client.VerifyPlatformSignature("x", struct{}{}, "x")
	fmt.Println(ok)
	if err != nil {
		panic(err)
	}
}
