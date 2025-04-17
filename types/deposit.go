package types

import "time"

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

type CallbackResponseData struct {
	Deposit  *Deposit         `json:"deposit"`
	Withdraw *WithdrawRequest `json:"withdraw"`
}
type CallbackResponse struct {
	Data CallbackResponseData `json:"data"`
	Sig  string               `json:"sig"`
}

type Deposit struct {

	// TxHash         string     `json:"tx_hash" gorm:"uniqueIndex"`
	TxHash         string     `json:"tx_hash" gorm:"comment:'交易哈希'"` // remove uniqueIndex for test batch transfer
	FromAddress    string     `json:"from_address" gorm:"index:idx_from_address;comment:'发送地址'"`
	ToAddress      string     `json:"to_address" gorm:"index:idx_to_address;comment:'接收地址'"`
	AddressUid     string     `json:"address_uid" gorm:"index:idx_address_uid;comment:'地址唯一标识'"`
	WalletPath     string     `json:"wallet_path" gorm:"comment:'钱包路径'"`
	Amount         string     `json:"amount" gorm:"comment:'金额（用于计算）'"` // always use this one for calculation
	Decimal        uint       `json:"decimal" gorm:"comment:'小数位数'"`
	FloatAmount    string     `json:"float_amount" gorm:"comment:'浮点金额（四位小数）'"` // four decimal places
	NetworkId      uint64     `json:"network_id" gorm:"index:idx_network_id;comment:'网络ID'"`
	NetworkName    string     `json:"network_name" gorm:"index:idx_network_name;comment:'网络名称'"`
	BlockNumber    uint64     `json:"block_number" gorm:"index:idx_block_number;comment:'区块高度'"`
	AssetId        uint64     `json:"asset_id" gorm:"index:idx_asset_id;comment:'资产ID'"`
	AssetName      AssetName  `json:"asset_name" gorm:"index:idx_asset_name;comment:'资产名称'"`
	Confirmed      bool       `json:"confirmed" gorm:"index:idx_confirmed;comment:'是否已确认'"`
	ConfirmedAt    *time.Time `json:"confirmed_at" gorm:"comment:'确认时间'"`
	TxStatus       TxStatus   `json:"tx_status" gorm:"index:idx_tx_status;comment:'交易状态'"`
	GasTxHash      string     `json:"gas_tx_hash" gorm:"comment:'Gas交易哈希'"`
	GasBaseFee     string     `json:"gas_base_fee" gorm:"comment:'Gas基础费用'"`
	GasPriorityFee string     `json:"gas_priority_fee" gorm:"comment:'Gas优先费用'"`
	GasNonce       uint64     `json:"gas_nonce" gorm:"comment:'Gas交易Nonce'"`
	// Bas
	IsCollected             bool              `json:"is_collected" gorm:"index:idx_is_collected;comment:'是否已归集'"`
	CollectTaskStatus       CollectTaskStatus `json:"collect_task_status" gorm:"index:idx_collect_task_status;default:pending;comment:'归集任务状态'"`
	CollectTxHash           string            `json:"collect_tx_hash" gorm:"index:idx_collect_tx_hash;comment:'归集交易哈希'"`
	CollectBaseFee          string            `json:"collect_base_fee" gorm:"comment:'归集基础费用'"`
	CollectPriorityFee      string            `json:"collect_priority_fee" gorm:"comment:'归集优先费用'"`
	CollectNonce            uint64            `json:"collect_nonce" gorm:"comment:'归集交易Nonce'"`
	CollectBlockNumber      uint64            `json:"collect_block_number" gorm:"index:idx_collect_block_number;comment:'归集区块高度'"`
	CollectTxConfirmed      bool              `json:"collect_tx_confirmed" gorm:"index:idx_collect_tx_confirmed;default:false;comment:'归集交易是否已确认'"`
	CollectConfirmedAt      *time.Time        `json:"collect_confirmed_at" gorm:"index:idx_collect_confirmed_at;comment:'归集确认时间'"`
	CollectedAt             *time.Time        `json:"collected_at" gorm:"comment:'归集完成时间'"`
	CollectTotalAmount      string            `json:"collect_total_amount" gorm:"comment:'归集总金额'"`
	CollectTotalAmountFloat string            `json:"collect_total_amount_float" gorm:"comment:'归集总金额'"`
	Retries                 uint              `json:"retries" gorm:"comment:'重试次数'"`
	LastError               string            `json:"last_error" gorm:"comment:'最近一次错误信息'"`
	ProjectUid              string            `json:"project_uid" gorm:"index:idx_project_uid;comment:'项目唯一标识'"`
	Fee                     string            `json:"fee" gorm:"comment:'手续费'"`
	CallbackModel
}

type WithdrawRequest struct {
	BaseModel
	Uid                           string         `json:"uid" gorm:"uniqueIndex;comment:'唯一标识'"`
	FromAddress                   string         `json:"from_address" gorm:"index:idx_from_address;comment:'发送地址'"`
	ToAddress                     string         `json:"to_address" gorm:"index:idx_to_address;comment:'接收地址'"`
	Amount                        string         `json:"amount" gorm:"comment:'金额'"`
	AmountToWithdraw              string         `json:"amount_to_withdraw" gorm:"comment:'提款金额'"`
	Fee                           string         `json:"fee" gorm:"comment:'手续费'"`
	Decimal                       uint           `json:"decimal" gorm:"comment:'小数位数'"`
	FloatAmount                   string         `json:"float_amount" gorm:"comment:'浮点金额'"`
	AssetId                       uint64         `json:"asset_id" gorm:"index:idx_asset_id;comment:'资产ID'"`
	AssetName                     AssetName      `json:"asset_name" gorm:"index:idx_asset_name;comment:'资产名称'"`
	NetworkId                     uint64         `json:"network_id" gorm:"index:idx_network_id;comment:'网络ID'"`
	NetworkName                   string         `json:"network_name" gorm:"index:idx_network_name;comment:'网络名称'"`
	WithdrawStatus                WithdrawStatus `json:"withdraw_status" gorm:"index:idx_withdraw_status;comment:'提款状态'"`
	Retries                       uint           `json:"retries" gorm:"comment:'重试次数'"`
	LastError                     string         `json:"last_error" gorm:"comment:'最近一次错误信息'"`
	Note                          string         `json:"note" gorm:"comment:'备注'"`
	AdminUid                      string         `json:"admin_uid" gorm:"index:idx_admin_uid;comment:'管理员唯一标识'"`
	ApproveTxHash                 string         `json:"approve_tx_hash" gorm:"index:idx_approve_tx_hash;comment:'批准交易哈希'"`
	BatchTransferTxHash           string         `json:"batch_transfer_tx_hash" gorm:"index:idx_batch_transfer_tx_hash;comment:'批量转账交易哈希'"`
	TotalBatchTransferAmount      string         `json:"total_batch_transfer_amount" gorm:"comment:'批量转账总金额'"`
	TotalBatchTransferAmountFloat string         `json:"total_batch_transfer_amount_float" gorm:"comment:'批量转账总金额'"`
	GasAmount                     string         `json:"gas_amount" gorm:"comment:'Gas金额'"`
	BaseFee                       string         `json:"base_fee" gorm:"comment:'基础费用'"`
	PriorityFee                   string         `json:"priority_fee" gorm:"comment:'优先费用'"`
	Nonce                         uint64         `json:"nonce" gorm:"comment:'交易Nonce'"`
	BlockNumber                   uint64         `json:"block_number" gorm:"index:idx_block_number;comment:'区块高度'"`
	Confirmed                     bool           `json:"confirmed" gorm:"index:idx_confirmed;default:false;comment:'是否已确认'"`
	ConfirmedAt                   *time.Time     `json:"confirmed_at" gorm:"comment:'确认时间'"`
	ProjectUid                    string         `json:"project_uid" gorm:"index:idx_project_uid;comment:'项目唯一标识'"`
	ReviewCount                   uint           `json:"review_count" gorm:"comment:'已经审核次数'"`
	IsColdWalletCollect           bool           `json:"is_cold_wallet_collect" gorm:"comment:'是否为冷钱包归集'default:false"`
	ColdWalletCollectTaskUid      string         `json:"cold_wallet_collect_task_uid" gorm:"index:idx_cold_wallet_collect_task_uid;comment:'冷钱包归集任务唯一标识'"`
	CallbackModel
}
