package coinbase

import (
	"context"
	"fmt"
)

// ListAccounts Lists current user’s accounts to which the authentication method has access to.
// Endpoint: GET /accounts
func (c *Client) ListAccounts(ctx context.Context) (*[]Account, *Pagination, error) {
	accounts := &[]Account{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", c.APIBase, "accounts"), nil)
	if err != nil {
		return accounts, pagination, err
	}

	if err = c.SendWithAuth(req, accounts, pagination); err != nil {
		return accounts, pagination, err
	}

	return accounts, pagination, nil
}

// GetAccount Show current user’s account.
// Endpoint: GET /accounts/:account_id
func (c *Client) GetAccount(ctx context.Context, accountID string) (*Account, error) {
	account := &Account{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", c.APIBase, "accounts", accountID), nil)
	if err != nil {
		return account, err
	}

	if err = c.SendWithAuth(req, account); err != nil {
		return account, err
	}

	return account, nil
}

// UpdateAccount Modifies user’s account.
// Endpoint: PUT /accounts/:account_id
func (c *Client) UpdateAccount(ctx context.Context, accountID string, accountData UpdateAccount) (*Account, error) {
	account := &Account{}

	req, err := c.NewRequest(ctx, "PUT", fmt.Sprintf("%s/%s/%s", c.APIBase, "accounts", accountID), accountData)
	if err != nil {
		return account, err
	}

	if err = c.SendWithAuth(req, account); err != nil {
		return account, err
	}

	return account, nil
}

// DeleteAccount Removes user’s account.
// Endpoint: DELETE /accounts/:account_id
func (c *Client) DeleteAccount(ctx context.Context, accountID string) error {
	accounts := &[]Account{}

	req, err := c.NewRequest(ctx, "DELETE", fmt.Sprintf("%s/%s/%s", c.APIBase, "accounts", accountID), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, accounts); err != nil {
		return err
	}

	return nil
}
