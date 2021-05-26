package coinbase

import (
	"context"
	"fmt"
)

// ListSells Lists sells for an account.
// Endpoint: GET /accounts/:account_id/sells
func (c *Client) ListSells(ctx context.Context, accountID string) (*[]Sell, *Pagination, error) {
	sells := &[]Sell{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "sells"), nil)
	if err != nil {
		return sells, pagination, err
	}

	if err = c.SendWithAuth(req, sells, pagination); err != nil {
		return sells, pagination, err
	}

	return sells, pagination, nil
}

// GetSell Show an individual sell.
// Endpoint: GET /accounts/:account_id/sells/:sell_id
func (c *Client) GetSell(ctx context.Context, accountID string, sellID string) (*Sell, error) {
	sell := &Sell{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "sells", sellID), nil)
	if err != nil {
		return sell, err
	}

	if err = c.SendWithAuth(req, sell); err != nil {
		return sell, err
	}

	return sell, nil
}

// PlaceSell Sells a user-defined amount of bitcoin, bitcoin cash, litecoin or ethereum.
// Endpoint: POST /accounts/:account_id/sells
func (c *Client) PlaceSell(ctx context.Context, accountID string, sellData PlaceSell) (*Sell, error) {
	sell := &Sell{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "sells"), sellData)
	if err != nil {
		return sell, err
	}

	if err = c.SendWithAuth(req, sell); err != nil {
		return sell, err
	}

	return sell, nil
}

// CommitSell Completes a sell that is created in commit: false state.
// Endpoint: POST /accounts/:account_id/sells/:sell_id/commit
func (c *Client) CommitSell(ctx context.Context, accountID string, sellID string) (*Sell, error) {
	sell := &Sell{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "sells", sellID, "commit"), nil)
	if err != nil {
		return sell, err
	}

	if err = c.SendWithAuth(req, sell); err != nil {
		return sell, err
	}

	return sell, nil
}
