package linkepay

import (
	"net/http"

	"github.com/linkepay/linkepay-sdk-go/types"
)

func NewClient(config *types.Config) *types.Client {
	return &types.Client{
		Config:     config,
		HTTPClient: &http.Client{},
	}
}
