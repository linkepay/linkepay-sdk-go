package linkepay

import (
	"github.com/linkepay/linkepay-sdk-go/types"
)

func NewConfig(baseURL string) *types.Config {
	return &types.Config{
		BaseURL: baseURL,
	}
}
