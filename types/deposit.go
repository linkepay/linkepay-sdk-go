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
	Uid           string  `json:"uid"`
	NetworkID     uint    `json:"network_id"`
	NetworkName   string  `json:"network_name"`
	Address       string  `json:"address"`
	Path          uint    `json:"path"`
	TokenBalances *string `json:"token_balances"`
	ProjectUID    string  `json:"project_uid"`
	Nonce         uint    `json:"nonce"`
}

// GetDepositsRequest represents parameters for querying deposit history
type GetDepositsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// DepositRecord represents a single deposit entry
type DepositRecord struct {
	ID          uint    `json:"id"`
	UID         string  `json:"uid"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
	ProjectUID  string  `json:"project_uid"`
	NetworkID   uint64  `json:"network_id"`
	NetworkName string  `json:"network_name"`
	AssetID     uint64  `json:"asset_id"`
	AssetName   string  `json:"asset_name"`
	Amount      string  `json:"amount"`
	FloatAmount string  `json:"float_amount"`
	TxHash      string  `json:"tx_hash"`
	TxStatus    string  `json:"tx_status"`
	FromAddress string  `json:"from_address"`
	UserAddress string  `json:"user_address"`
	AddressUID  string  `json:"address_uid"`
}

// TokenSummary represents a summary of deposits/withdrawals by token
type TokenSummary struct {
	AssetID     uint64 `json:"asset_id"`
	AssetName   string `json:"asset_name"`
	NetworkID   uint64 `json:"network_id"`
	NetworkName string `json:"network_name"`
	TotalAmount string `json:"total_amount"`
	RecordCount int    `json:"record_count"`
}

// GetDepositsResponse represents the response from the get deposits API
type GetDepositsResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Data struct {
			Deposits []DepositRecord `json:"deposits"`
			Summary  []TokenSummary  `json:"summary"`
		} `json:"data"`
		Page     int   `json:"page"`
		PageSize int   `json:"page_size"`
		Total    int64 `json:"total"`
	} `json:"data"`
	Sig string `json:"sig"`
}
