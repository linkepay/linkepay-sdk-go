package main

import (
	"fmt"

	"github.com/linkepay/linkepay-sdk-go"
	"github.com/linkepay/linkepay-sdk-go/types"
)

func main() {

	// pay:
	// project_uid: "bEApWREzo5"
	// consolidation_address: "0x9eE4C914A0c6E41E5F7c4c9febff9FEC846835eF"
	// public_key: "0x9e971acec71dd5d308ba0dc0cd533e15df6411b8adb02c72dac2c1951e6937b9d0f23db70e995dc089e5e72949e189d95e306e83401d7ac749864135b68d6584"
	// private_key: "0xbb33d2de2bd86d891ef5d71603ef93cd6a5214a0ed79bdb0a819d295c6e77690"
	// api_key: "601b24c6c991d3e82de7f804661c3e49"
	// host: "https://api.cheetos.club/"
	// vendor_public_key: "0xa38913ed7c0d9d170a47ee264de240faf1c5eba53a006f46bbd96a7907d801137adb1459c3c8041947ab76ae529bad0da378c1c5c6778a8676f862cdd528d30f"

	client := linkepay.NewClient(&types.Config{
		BaseURL:              "https://api.cheetos.club/",
		ProjectID:            "bEApWREzo5",
		PublicKey:            "0x9e971acec71dd5d308ba0dc0cd533e15df6411b8adb02c72dac2c1951e6937b9d0f23db70e995dc089e5e72949e189d95e306e83401d7ac749864135b68d6584",
		PrivateKey:           "0xbb33d2de2bd86d891ef5d71603ef93cd6a5214a0ed79bdb0a819d295c6e77690",
		PayPlatformPublicKey: "0xa38913ed7c0d9d170a47ee264de240faf1c5eba53a006f46bbd96a7907d801137adb1459c3c8041947ab76ae529bad0da378c1c5c6778a8676f862cdd528d30f",
	})

	resp, err := client.CreateMultipleDepositAddress(&types.CreateMultipleDepositAddressRequest{
		NetworkID: 11155111,
		Count:     2,
	})
	fmt.Println(resp)
	fmt.Println(err)

}
