package linkepay

import (
	"github.com/linkepay/linkepay-sdk-go/operations"
	"github.com/linkepay/linkepay-sdk-go/types"
	"github.com/linkepay/linkepay-sdk-go/utils"
)

var Client = &types.Client{}

func SetClient(config *types.Config) *types.Client {
	Client = &types.Client{
		Config: *config,
	}
	return Client
}

func GetDepositAddress(req *types.GetDepositAddressRequest) (*types.GetDepositAddressResponse, error) {
	return operations.GetDepositAddress(Client, req)
}

func RequestWithdrawal(data types.WithdrawalRequest) (map[string]interface{}, error) {
	return operations.RequestWithdrawal(Client, data)
}

func VerifyPlatformSignature(platformPublicKey string, data interface{}, platformSignature string) (bool, error) {
	km := utils.NewKeyManager()
	km.LoadKeys(&Client.Config)
	return km.VerifyPlatformSignature(platformPublicKey, data, platformSignature)
}

func GenerateKeys() (types.Keys, error) {
	km := utils.NewKeyManager()
	km.LoadKeys(&Client.Config)
	return km.GenerateKeys()
}

func GetPlatformPublicKey() string {
	km := utils.NewKeyManager()
	km.LoadKeys(&Client.Config)
	return km.GetPlatformPublicKey()
}
