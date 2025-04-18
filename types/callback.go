package types

import "time"

type CallbackRespData struct {
	ProjectUid  string     `json:"project_uid"`
	FromAddress string     `json:"from_address"`
	ToAddress   string     `json:"to_address"`
	AddressUid  string     `json:"address_uid"`
	Amount      string     `json:"amount"`
	FloatAmount string     `json:"float_amount"`
	Decimal     uint       `json:"decimal"`
	TxHash      string     `json:"tx_hash"`
	Status      string     `json:"status"`
	AssetId     uint64     `json:"asset_id"`
	AssetName   string     `json:"asset_name"`
	NetworkId   uint64     `json:"network_id"`
	NetworkName string     `json:"network_name"`
	Confirmed   bool       `json:"confirmed"`
	ConfirmedAt *time.Time `json:"confirmed_at"`
	Type        string     `json:"type"` // "deposit" or "withdrawal"
}

type CallbackRequestDataWithSig struct {
	Data CallbackRespData `json:"verify_data"`
	Sig  string           `json:"sig"`
}
