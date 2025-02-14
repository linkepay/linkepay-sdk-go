package operations

import (
	"fmt"

	"github.com/linkepay/linkepay-sdk-go/types"
)

// CreateWithdrawal creates a new withdrawal request
func CreateWithdrawal(c *types.Client, req *types.WithdrawalRequest) (*types.WithdrawalResponse, error) {
	resp := &types.WithdrawalResponse{}
	// err := c.Post("/withdrawals", req, resp)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("CreateWithdrawal resp", resp)
	return resp, nil
}

// GetWithdrawalStatus gets the status of a withdrawal by ID
func GetWithdrawalStatus(c *types.Client, id string) (*types.WithdrawalResponse, error) {
	resp := &types.WithdrawalResponse{}
	// err := c.Get("/withdrawals/"+id, resp)
	// if err != nil {
	// 	return nil, err
	// }
	return resp, nil
}
