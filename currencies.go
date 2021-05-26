package coinbase

import (
	"context"
	"fmt"
)

// ListCurrencies List known currencies.
// Endpoint: GET /currencies
func (c *Client) ListCurrencies(ctx context.Context) (*[]Currency, error) {
	currencies := &[]Currency{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", c.APIBase, "currencies"), nil)
	if err != nil {
		return currencies, err
	}

	if err = c.Send(req, currencies); err != nil {
		return currencies, err
	}

	return currencies, nil
}
