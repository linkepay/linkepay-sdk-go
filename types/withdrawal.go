package types

type RequestWithdrawalRequest struct {
	AssetID   uint    `json:"asset_id"`
	Amount    float64 `json:"amount"`
	ToAddress string  `json:"to_address"`
	NetworkID uint    `json:"network_id"`
	UserUID   string  `json:"user_uid"`
	UID       string  `json:"uid"`
}

type RequestWithdrawalResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

// GetWithdrawalsRequest represents parameters for querying withdrawal history
type GetWithdrawalsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// WithdrawalRecord represents a single withdrawal entry
type WithdrawalRecord struct {
	ID             uint    `json:"id"`
	UID            string  `json:"uid"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	DeletedAt      *string `json:"deleted_at"`
	ProjectUID     string  `json:"project_uid"`
	NetworkID      uint64  `json:"network_id"`
	NetworkName    string  `json:"network_name"`
	AssetID        uint64  `json:"asset_id"`
	AssetName      string  `json:"asset_name"`
	Amount         string  `json:"amount"`
	FloatAmount    string  `json:"float_amount"`
	TxHash         string  `json:"tx_hash"`
	WithdrawStatus string  `json:"withdraw_status"`
	ToAddress      string  `json:"to_address"`
}

// GetWithdrawalsResponse represents the response from the get withdrawals API
type GetWithdrawalsResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Data struct {
			Withdraws []WithdrawalRecord `json:"withdraws"`
			Summary   []TokenSummary     `json:"summary"`
		} `json:"data"`
		Page     int   `json:"page"`
		PageSize int   `json:"page_size"`
		Total    int64 `json:"total"`
	} `json:"data"`
	Sig string `json:"sig"`
}
