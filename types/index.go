package types

type Client struct {
	Config Config
}

type Config struct {
	BaseURL              string
	ProjectID            string
	PublicKey            string
	PrivateKey           string
	PayPlatformPublicKey string
}

type WithdrawalResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	// Other fields
}

type Keys struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}
