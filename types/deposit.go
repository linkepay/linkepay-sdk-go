package types

type GetDepositAddressResponse struct {
	Message string `json:"message"`
	Address string `json:"address"`
	UID     string `json:"uid"`
}

type GetDepositAddressRequest struct {
	NetworkID uint   `json:"network_id"`
	UserUID   string `json:"user_uid"`
}

type CreateDepositAddressRequest struct {
	NetworkID uint   `json:"network_id"`
	UserUID   string `json:"user_uid"`
}

type CreateDepositAddressResponse struct {
	Message string `json:"message"`
	Address string `json:"address"`
	UID     string `json:"uid"`
}

type CreateMultipleDepositAddressRequest struct {
	Count     uint `json:"count"`
	NetworkID uint `json:"network_id"`
}

type CreateMultipleDepositAddressResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Addresses []Address `json:"addresses"`
}

type Address struct {
	ID            uint    `json:"id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	DeletedAt     *string `json:"deleted_at"`
	UserUID       string  `json:"user_uid"`
	NetworkID     uint    `json:"network_id"`
	NetworkName   string  `json:"network_name"`
	Address       string  `json:"address"`
	Path          uint    `json:"path"`
	TokenBalances *string `json:"token_balances"`
	ProjectUID    string  `json:"project_uid"`
	Nonce         uint    `json:"nonce"`
}
