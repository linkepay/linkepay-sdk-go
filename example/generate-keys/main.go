package main

import (
	"fmt"

	"github.com/linkepay/linkepay-sdk-go"
	"github.com/linkepay/linkepay-sdk-go/types"
)

func main() {
	client := linkepay.NewClient(&types.Config{})

	keys, err := client.GenerateKeys()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Private Key: %s\n", keys.PrivateKey)
	fmt.Printf("Public Key: %s\n", keys.PublicKey)
	fmt.Printf("Address: %s\n", keys.Address)
}
