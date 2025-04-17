package types

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel contains common columns for all tables.
type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey" sql:"AUTO_INCREMENT"`
	UID       string         `json:"uid" gorm:"uniqueIndex;type:varchar(36);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

type CallbackModel struct {
	RequestBody            string         `json:"request_body" gorm:"type:text;comment:'回调请求体内容'"`
	RequestError           string         `json:"request_error" gorm:"comment:'回调请求错误信息'"`
	CallBackStatus         CallBackStatus `json:"call_back_status" gorm:"index:idx_call_back_status;default:pending;comment:'回调状态'"`
	CallBackResponseCode   int            `json:"call_back_response_code" gorm:"comment:'回调响应状态码'"`
	CallBackResponseStatus string         `json:"call_back_response_status" gorm:"comment:'回调响应状态描述'"`
	CallBackResponseBody   string         `json:"call_back_response_body" gorm:"type:text;comment:'回调响应体内容'"`
	LastCallBackAt         *time.Time     `json:"last_call_back_at" gorm:"comment:'最近一次回调时间'"`
	CallBackRetries        uint           `json:"call_back_retries" gorm:"default:0;comment:'回调重试次数'"`
	NextCallBackAt         *time.Time     `json:"next_call_back_at" gorm:"index:idx_next_call_back_at;comment:'下一次回调时间'"`
}

// !all amount is in string and with decimals
type TxType string

const TrxShastaUsdcAddress = "todo"
const TrxShastaUsdtAddress = "TZALvaoytcCTdLzsVHU3Rvywt9FrnvHToU"

const SepoliaUsdcAddress = "0x80214AE442613Aa6371303501571A6430c067Ccf"
const SepoliaUsdtAddress = "0x0EC675b445Cf07D954aAcC33890D45A967e4CEB6"

const (
	DepositTx  TxType = "deposit"
	WithdrawTx TxType = "withdraw"
)

var SupportedNetworkIds = map[uint64]bool{
	EthereumMainnetNetworkID: true,
	BSCMainnetNetworkID:      true,
	SepoliaNetworkID:         true,
	TRXNetworkID:             true,
	TRXShastaNetworkID:       true,
}

const (
	EthereumMainnetNetworkID = 1
	BSCMainnetNetworkID      = 66
	BSCTestnetNetworkID      = 67
	SepoliaNetworkID         = 11155111
	TRXNetworkID             = 4
	TRXShastaNetworkID       = 5
	AllNetworkID             = 0
)

var SupportedNetworks = []uint64{
	EthereumMainnetNetworkID,
	BSCMainnetNetworkID,
	SepoliaNetworkID,
	TRXNetworkID,
	TRXShastaNetworkID,
}

type NetworkType string

const (
	EvmCompatibleNetwork NetworkType = "evm-compatible"
	TRXNetwork           NetworkType = "TRX"
)

type NetworkName string

const (
	EthereumMainnetNetworkName NetworkName = "Ethereum Mainnet"
	BSCMainnetNetworkName      NetworkName = "BSC Mainnet"
	SepoliaNetworkName         NetworkName = "Sepolia"
	TRXNetworkName             NetworkName = "TRX"
	TRXShastaNetworkName       NetworkName = "TRX Shasta"
	AllNetworkName             NetworkName = "All"
)

var NetworkNames = map[uint64]NetworkName{
	EthereumMainnetNetworkID: EthereumMainnetNetworkName,
	BSCMainnetNetworkID:      BSCMainnetNetworkName,
	SepoliaNetworkID:         SepoliaNetworkName,
	TRXNetworkID:             TRXNetworkName,
	TRXShastaNetworkID:       TRXShastaNetworkName,
	AllNetworkID:             AllNetworkName,
}

var NetworkTypes = map[uint64]NetworkType{
	EthereumMainnetNetworkID: EvmCompatibleNetwork,
	BSCMainnetNetworkID:      EvmCompatibleNetwork,
	SepoliaNetworkID:         EvmCompatibleNetwork,
	TRXNetworkID:             TRXNetwork,
	TRXShastaNetworkID:       TRXNetwork,
}

var NetworkConfirmationBlocks = map[uint64]int64{
	EthereumMainnetNetworkID: 6,
	BSCMainnetNetworkID:      6,
	SepoliaNetworkID:         6,
	TRXNetworkID:             6,
	TRXShastaNetworkID:       6,
}

var SupportedAssetIds = []uint64{UsdcAssetID, UsdtAssetID}

// networkId => address => assetId
var SupportedAssetAddresses = map[uint64]map[string]uint64{
	TRXNetworkID: {
		TrxShastaUsdcAddress: UsdcAssetID,
		TrxShastaUsdtAddress: UsdtAssetID,
	},
	TRXShastaNetworkID: {
		TrxShastaUsdcAddress: UsdcAssetID,
		TrxShastaUsdtAddress: UsdtAssetID,
	},
}

var SupportedAddressesListForDifferentNetworks = map[uint64][]string{
	TRXNetworkID: {
		TrxShastaUsdcAddress,
		TrxShastaUsdtAddress,
	},
	TRXShastaNetworkID: {
		TrxShastaUsdcAddress,
		TrxShastaUsdtAddress,
	},
	SepoliaNetworkID: {
		SepoliaUsdtAddress,
		SepoliaUsdcAddress,
	},
}

const AllAssetID uint64 = 0

const (
	UsdcAssetID uint64 = iota + 1
	UsdtAssetID
	NativeAssetID
)

type AssetName string

const (
	UsdcAsset   AssetName = "USDC"
	UsdtAsset   AssetName = "USDT"
	NativeAsset AssetName = "Native"
	AllAsset    AssetName = "All"
)

var AssetNames = map[uint64]AssetName{
	UsdcAssetID:   UsdcAsset,
	UsdtAssetID:   UsdtAsset,
	NativeAssetID: NativeAsset,
	AllAssetID:    AllAsset,
}

// networkId -> assetId -> address
var AssetAddressesForDifferentNetworks = map[uint64]map[uint64]string{
	SepoliaNetworkID: {
		UsdcAssetID:   SepoliaUsdcAddress,
		UsdtAssetID:   SepoliaUsdtAddress,
		NativeAssetID: "0x0000000000000000000000000000000000000000",
	},

	TRXNetworkID: {
		UsdcAssetID:   "TEkxiTehnzSmSe2XqrBj4w32RUN966rdz8",
		UsdtAssetID:   "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
		NativeAssetID: "0x0000000000000000000000000000000000000000",
	},

	TRXShastaNetworkID: {
		UsdcAssetID:   "TG3wrVRHL7rsdHFfdRopGXwGsiN6AFJZsx",
		UsdtAssetID:   "TZALvaoytcCTdLzsVHU3Rvywt9FrnvHToU",
		NativeAssetID: "0x0000000000000000000000000000000000000000",
	},
}

// networkId -> assetId -> decimal
var AssetDecimalForDifferentNetworks = map[uint64]map[uint64]int64{
	SepoliaNetworkID: {
		UsdcAssetID:   18,
		UsdtAssetID:   18,
		NativeAssetID: 18,
	},
	TRXNetworkID: {
		UsdcAssetID:   18,
		UsdtAssetID:   18,
		NativeAssetID: 6,
	},
	TRXShastaNetworkID: {
		UsdcAssetID:   18,
		UsdtAssetID:   18,
		NativeAssetID: 6,
	},
}

var BatchTransferContractForDifferentNetworks = map[uint64]string{
	TRXNetworkID:       "TQqNxL7Tkg1JvB1JPVAEgRVArqafUaAbgy",
	TRXShastaNetworkID: "TQqNxL7Tkg1JvB1JPVAEgRVArqafUaAbgy",
	SepoliaNetworkID:   "0xAa6beBE4cF31e21DAB0C4e90593b6C29dC8270AA",
}

const (
	Trc20Transfer2AccountWithBalanceEnergy    int64 = 13045
	Trc20Transfer2AccountWithoutBalanceEnergy int64 = 28036
	Trc20Transfer2AccountBandwidth            int64 = 346

	TrxTransferBandWith = 269
)

// const (
// 	ShastaTrc20TransferWithBalanceTrxDelegateAmount     int64 = 207380
// 	ShastaTrc20TransferWithoutBalanceTrxDelegateAmount  int64 = 420280
// 	MainnetTrc20TransferWithBalanceTrxDelegateAmount    int64 = 0
// 	MainnetTrc20TransferWithoutBalanceTrxDelegateAmount int64 = 0
// )

const (
	ShastaTotalStakedTrx  = 648802467567
	MainnetTotalStakedTrx = 48381744745
)

const (
	// 28036 energy to 11.775120 trx
	TrxToEnergy = float64(0.0023809524)
	// 267 Bandwidth to 0.267 TRX
	TrxToBandwidth = float64(0.001)
)

type TxStatus string
type WithdrawStatus string
type CollectTaskStatus string
type CallBackStatus string

const (
	PendingTxStatus   TxStatus = "pending"
	ConfirmedTxStatus TxStatus = "confirmed"
	TimeoutTxStatus   TxStatus = "timeout" // pending for too long(1 hour)
)

const (
	UnderReviewWithdrawStatus WithdrawStatus = "under_review" // waiting for review
	RejectedWithdrawStatus    WithdrawStatus = "rejected"     // review rejected
	PendingWithdrawStatus     WithdrawStatus = "pending"      // review approved
	ProcessingWithdrawStatus  WithdrawStatus = "processing"
	SuccessWithdrawStatus     WithdrawStatus = "success"
	ConfirmedWithdrawStatus   WithdrawStatus = "confirmed"
	ErrorWithdrawStatus       WithdrawStatus = "error"
	FailedWithdrawStatus      WithdrawStatus = "failed"
)

const (
	CollectTaskStatusPending    CollectTaskStatus = "pending"
	CollectTaskStatusProcessing CollectTaskStatus = "processing"
	CollectTaskStatusSuccess    CollectTaskStatus = "success"
	CollectTaskStatusConfirmed  CollectTaskStatus = "confirmed"
	CollectTaskStatusError      CollectTaskStatus = "error"
	CollectTaskStatusFailed     CollectTaskStatus = "failed"
)

const (
	CallBackStatusPending    CallBackStatus = "pending"
	CallBackStatusProcessing CallBackStatus = "processing"
	CallBackStatusSuccess    CallBackStatus = "success"
	CallBackStatusError      CallBackStatus = "error"
	CallBackStatusFailed     CallBackStatus = "failed"
)

type CollectPolicy string

const (
	CollectAfterEveryDeposit  CollectPolicy = "after_every_deposit"
	CollectEveryFixedDuration CollectPolicy = "every_fixed_duration"
	CollectAtFixedTime        CollectPolicy = "at_fixed_time"
	ThresholdCollect          CollectPolicy = "threshold"
	CollectManually           CollectPolicy = "manually"
)

var AsssetTokenIconsForDifferentNetworks = map[uint64]map[uint64]string{
	SepoliaNetworkID: {
		UsdcAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/usd-coin-usdc-logo.svg",
		UsdtAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tether-usdt-logo.svg",
		NativeAssetID: "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/ethereum-eth.svg",
	},
	BSCMainnetNetworkID: {
		UsdcAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/usd-coin-usdc-logo.svg",
		UsdtAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tether-usdt-logo.svg",
		NativeAssetID: "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/bnb-bnb-logo.svg",
	},
	BSCTestnetNetworkID: {
		UsdcAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/usd-coin-usdc-logo.svg",
		UsdtAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tether-usdt-logo.svg",
		NativeAssetID: "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/bnb-bnb-logo.svg",
	},
	TRXNetworkID: {
		UsdcAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/usd-coin-usdc-logo.svg",
		UsdtAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tether-usdt-logo.svg",
		NativeAssetID: "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tron-trx-icon.svg",
	},
	TRXShastaNetworkID: {
		UsdcAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/usd-coin-usdc-logo.svg",
		UsdtAssetID:   "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tether-usdt-logo.svg",
		NativeAssetID: "https://pub-e344fce3baef429cb82b43d6712df2db.r2.dev/tron-trx-icon.svg",
	},
}

type StakeResourceType string

const (
	StakeResourceTypeEnergy    StakeResourceType = "energy"
	StakeResourceTypeBandwidth StakeResourceType = "bandwidth"
)

var CurrencyNameForDifferentNetworks = map[uint64]map[uint64]string{
	SepoliaNetworkID: {
		UsdcAssetID:   "Sepolia USDC",
		UsdtAssetID:   "Sepolia USDT",
		NativeAssetID: "Sepolia ETH",
	},
	EthereumMainnetNetworkID: {
		UsdcAssetID:   "Ethereum USDC",
		UsdtAssetID:   "Ethereum USDT",
		NativeAssetID: "ETH",
	},
	TRXNetworkID: {
		UsdcAssetID:   "TRX USDC",
		UsdtAssetID:   "TRX USDT",
		NativeAssetID: "TRX",
	},
	TRXShastaNetworkID: {
		UsdcAssetID:   "TRX Shasta USDC",
		UsdtAssetID:   "TRX Shasta USDT",
		NativeAssetID: "TRX Shasta TRX",
	},
	BSCMainnetNetworkID: {
		UsdcAssetID:   "BSC USDC",
		UsdtAssetID:   "BSC USDT",
		NativeAssetID: "BNB",
	},
	BSCTestnetNetworkID: {
		UsdcAssetID:   "BSC Testnet USDC",
		UsdtAssetID:   "BSC Testnet USDT",
		NativeAssetID: "BSC Testnet BNB",
	},
}
