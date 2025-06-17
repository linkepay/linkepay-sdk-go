package types

type RequestWithdrawalRequest struct {
	AssetID     uint    `json:"asset_id"`
	Amount      float64 `json:"amount"`
	Destination string  `json:"destination"`
	NetworkID   uint    `json:"network_id"`
	UserUID     string  `json:"user_uid"`
	UID         string  `json:"uid"`
}

type RequestWithdrawalResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}
