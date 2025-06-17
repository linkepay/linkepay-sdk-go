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
	resp, err := operations.GetDepositAddress(&types.Client{Config: c.Config}, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateDepositAddress(req *types.CreateDepositAddressRequest) (*types.CreateDepositAddressResponse, error) {
	resp, err := operations.CreateDepositAddress(&types.Client{Config: c.Config}, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) CreateMultipleDepositAddress(req *types.CreateMultipleDepositAddressRequest) (*types.CreateMultipleDepositAddressResponse, error) {
	resp, err := operations.CreateMultipleDepositAddress(&types.Client{Config: c.Config}, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) RequestWithdrawal(data types.RequestWithdrawalRequest) (*types.RequestWithdrawalResponse, error) {
	resp, err := operations.RequestWithdrawal(&types.Client{Config: c.Config}, data)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) VerifyPlatformSignature(platformPublicKey string, data interface{}, platformSignature string) (bool, error) {
	resp, err := c.km.VerifyPlatformSignature(platformPublicKey, data, platformSignature)
	if err != nil {
		return false, err
	}
	return resp, nil
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
	resp, err := c.km.VerifySignature(publicKey, signature, message)
	if err != nil {
		return false, err
	}
	return resp, nil
}

func (c *Client) GenerateKeys() (types.Keys, error) {
	resp, err := c.km.GenerateKeys()
	if err != nil {
		return types.Keys{}, err
	}
	return resp, nil
}

func (c *Client) GetPlatformPublicKey() string {
	return c.km.GetPlatformPublicKey()
}

func (c *Client) SignDataWithPrivateKey(data interface{}, privateKey string) (string, error) {
	resp, err := c.km.SignDataWithPrivateKey(data, privateKey)
	if err != nil {
		return "", err
	}
	return resp, nil
}
