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
	ok, err := linkepay.VerifyPlatformSignature("x", struct{}{}, "x")
	fmt.Println(ok)
	if err != nil {
		panic(err)
	}
}
