package types

type GetDepositAddressResponse struct {
	Message string `json:"message"`
	Address string `json:"address"`
	UID     string `json:"uid"`
}

type GetDepositAddressRequest struct {
	NetworkID uint
	UserUID   string
}
