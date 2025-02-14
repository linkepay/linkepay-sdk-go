package types

import "net/http"

type Client struct {
	Config     *Config
	HTTPClient *http.Client
}

type Config struct {
	BaseURL string
}

type WithdrawalRequest struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	// Other fields
}

type WithdrawalResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	// Other fields
}
