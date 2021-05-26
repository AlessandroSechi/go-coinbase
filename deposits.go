package coinbase

import (
	"context"
	"fmt"
)

// ListDeposits Lists deposits for an account.
// Endpoint: GET /accounts/:account_id/deposits
func (c *Client) ListDeposits(ctx context.Context, accountID string) (*[]Deposit, *Pagination, error) {
	deposits := &[]Deposit{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "deposits"), nil)
	if err != nil {
		return deposits, pagination, err
	}

	if err = c.SendWithAuth(req, deposits, pagination); err != nil {
		return deposits, pagination, err
	}

	return deposits, pagination, nil
}

// GetDeposit Show an individual deposit.
// Endpoint: GET /accounts/:account_id/deposits/:deposit_id
func (c *Client) GetDeposit(ctx context.Context, accountID string, depositID string) (*Deposit, error) {
	deposit := &Deposit{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "deposits", depositID), nil)
	if err != nil {
		return deposit, err
	}

	if err = c.SendWithAuth(req, deposit); err != nil {
		return deposit, err
	}

	return deposit, nil
}

// DepositFunds Deposits user-defined amount of funds to a fiat account.
// Endpoint: POST /accounts/:account_id/deposits
func (c *Client) DepositFunds(ctx context.Context, accountID string, depositData DepositFunds) (*Deposit, error) {
	deposit := &Deposit{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "deposits"), depositData)
	if err != nil {
		return deposit, err
	}

	if err = c.SendWithAuth(req, deposit); err != nil {
		return deposit, err
	}

	return deposit, nil
}

// CommitDeposit Completes a deposit that is created in commit: false state.
// Endpoint: POST /accounts/:account_id/deposits/:deposit_id/commit
func (c *Client) CommitDeposit(ctx context.Context, accountID string, depositID string) (*Deposit, error) {
	deposit := &Deposit{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "deposits", depositID, "commit"), nil)
	if err != nil {
		return deposit, err
	}

	if err = c.SendWithAuth(req, deposit); err != nil {
		return deposit, err
	}

	return deposit, nil
}
