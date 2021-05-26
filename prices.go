package coinbase

import (
	"context"
	"fmt"
)

// GetBuyPrice Get the total price to buy one bitcoin or ether.
// Endpoint: GET /prices/:currency_pair/buy
func (c *Client) GetBuyPrice(ctx context.Context, currencyPair string) (*Price, error) {
	priceResponse := &Price{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", c.APIBase, currencyPair, "buy"), nil)
	if err != nil {
		return priceResponse, err
	}

	if err = c.Send(req, priceResponse); err != nil {
		return priceResponse, err
	}

	return priceResponse, nil
}

// GetSellPrice Get the total price to sell one bitcoin or ether.
// Endpoint: GET /prices/:currency_pair/sell
func (c *Client) GetSellPrice(ctx context.Context, currencyPair string) (*Price, error) {
	priceResponse := &Price{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", c.APIBase, currencyPair, "sell"), nil)
	if err != nil {
		return priceResponse, err
	}

	if err = c.Send(req, priceResponse); err != nil {
		return priceResponse, err
	}

	return priceResponse, nil
}

// GetSpotPrice Get the current market price for bitcoin.
// Endpoint: GET /prices/:currency_pair/spot
func (c *Client) GetSpotPrice(ctx context.Context, currencyPair string) (*Price, error) {
	priceResponse := &Price{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", c.APIBase, currencyPair, "spot"), nil)
	if err != nil {
		return priceResponse, err
	}

	if err = c.Send(req, priceResponse); err != nil {
		return priceResponse, err
	}

	return priceResponse, nil
}
