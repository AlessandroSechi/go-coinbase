package coinbase

import (
	"context"
	"fmt"
)

// ListBuys Lists buys for an account.
// Endpoint: GET /accounts/:account_id/buys
func (c *Client) ListBuys(ctx context.Context, accountID string) (*[]Buy, *Pagination, error) {
	buys := &[]Buy{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "buys"), nil)
	if err != nil {
		return buys, pagination, err
	}

	if err = c.SendWithAuth(req, buys, pagination); err != nil {
		return buys, pagination, err
	}

	return buys, pagination, nil
}

// GetBuy Show an individual buy.
// Endpoint: GET /accounts/:account_id/buys/:buy_id
func (c *Client) GetBuy(ctx context.Context, accountID string, buyID string) (*Buy, error) {
	buy := &Buy{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "buys", buyID), nil)
	if err != nil {
		return buy, err
	}

	if err = c.SendWithAuth(req, buy); err != nil {
		return buy, err
	}

	return buy, nil
}

// PlaceBuy Buys a user-defined amount of bitcoin, bitcoin cash, litecoin or ethereum.
// Endpoint: POST /accounts/:account_id/buys
func (c *Client) PlaceBuy(ctx context.Context, accountID string, buyData PlaceBuy) (*Buy, error) {
	buy := &Buy{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "buys"), buyData)
	if err != nil {
		return buy, err
	}

	if err = c.SendWithAuth(req, buy); err != nil {
		return buy, err
	}

	return buy, nil
}

// CommitBuy Completes a buy that is created in commit: false state.
// Endpoint: POST /accounts/:account_id/buys/:buy_id/commit
func (c *Client) CommitBuy(ctx context.Context, accountID string, buyID string) (*Buy, error) {
	buy := &Buy{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "buys", buyID, "commit"), nil)
	if err != nil {
		return buy, err
	}

	if err = c.SendWithAuth(req, buy); err != nil {
		return buy, err
	}

	return buy, nil
}
