package linkepay

import (
	"github.com/linkepay/linkepay-sdk-go/operations"
	"github.com/linkepay/linkepay-sdk-go/types"
	"github.com/linkepay/linkepay-sdk-go/utils"
)

type Client struct {
	Config types.Config
}

func NewClient(config *types.Config) *Client {
	return &Client{
		Config: *config,
	}
}

func (c *Client) GetDepositAddress(req *types.GetDepositAddressRequest) (*types.GetDepositAddressResponse, error) {
	return operations.GetDepositAddress(&types.Client{Config: c.Config}, req)
}

func (c *Client) CreateDepositAddress(req *types.CreateDepositAddressRequest) (*types.CreateDepositAddressResponse, error) {
	return operations.CreateDepositAddress(&types.Client{Config: c.Config}, req)
}

func (c *Client) CreateMultipleDepositAddress(req *types.CreateMultipleDepositAddressRequest) (*types.CreateMultipleDepositAddressResponse, error) {
	return operations.CreateMultipleDepositAddress(&types.Client{Config: c.Config}, req)
}

func (c *Client) RequestWithdrawal(data types.RequestWithdrawalRequest) (map[string]interface{}, error) {
	return operations.RequestWithdrawal(&types.Client{Config: c.Config}, data)
}

func (c *Client) VerifyPlatformSignature(platformPublicKey string, data interface{}, platformSignature string) (bool, error) {
	km := utils.NewKeyManager()
	km.LoadKeys(&c.Config)
	return km.VerifyPlatformSignature(platformPublicKey, data, platformSignature)
}

func (c *Client) VerifySignature(publicKey string, signature string, message string) (bool, error) {
	km := utils.NewKeyManager()
	km.LoadKeys(&c.Config)
	return km.VerifySignature(publicKey, signature, message)
}

func (c *Client) GenerateKeys() (types.Keys, error) {
	km := utils.NewKeyManager()
	km.LoadKeys(&c.Config)
	return km.GenerateKeys()
}

func (c *Client) GetPlatformPublicKey() string {
	km := utils.NewKeyManager()
	km.LoadKeys(&c.Config)
	return km.GetPlatformPublicKey()
}

func (c *Client) SignDataWithPrivateKey(data interface{}, privateKey string) (string, error) {
	km := utils.NewKeyManager()
	return km.SignDataWithPrivateKey(data, privateKey)
}
