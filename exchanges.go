package coinbase

import (
	"context"
	"fmt"
)

// ListExchangeRates Get current exchange rates.
// Endpoint: GET /exchange-rates
func (c *Client) ListExchangeRates(ctx context.Context, currency string) (*ExchangeRates, error) {
	exchangeRates := &ExchangeRates{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", c.APIBase, "exchange-rates"), currency)
	if err != nil {
		return exchangeRates, err
	}

	if err = c.Send(req, exchangeRates); err != nil {
		return exchangeRates, err
	}

	return exchangeRates, nil
}
