package linkepay

import (
	"errors"

	"github.com/linkepay/linkepay-sdk-go/operations"
	"github.com/linkepay/linkepay-sdk-go/types"
	"github.com/linkepay/linkepay-sdk-go/utils"
)

type Client struct {
	Config types.Config
	km     *utils.KeyManager
}

func NewClient(config *types.Config) *Client {
	km := utils.NewKeyManager()
	km.LoadKeys(config)
	return &Client{
		Config: *config,
		km:     km,
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
	return c.km.VerifyPlatformSignature(platformPublicKey, data, platformSignature)
}

func (c *Client) ParseCallbackData(data types.CallbackRequestDataWithSig) (*types.CallbackRespData, error) {
	// verify signature
	ok, err := c.VerifyPlatformSignature(c.km.GetPlatformPublicKey(), data.Data, data.Sig)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("invalid signature")
	}

	return &data.Data, nil
}

func (c *Client) VerifySignature(publicKey string, signature string, message string) (bool, error) {
	return c.km.VerifySignature(publicKey, signature, message)
}

func (c *Client) GenerateKeys() (types.Keys, error) {
	return c.km.GenerateKeys()
}

func (c *Client) GetPlatformPublicKey() string {
	return c.km.GetPlatformPublicKey()
}

func (c *Client) SignDataWithPrivateKey(data interface{}, privateKey string) (string, error) {
	return c.km.SignDataWithPrivateKey(data, privateKey)
}
