package coinbase

import (
	"context"
	"fmt"
)

// ListWithdrawals Lists withdrawals for an account.
// Endpoint: GET /accounts/:account_id/withdrawals
func (c *Client) ListWithdrawals(ctx context.Context, accountID string) (*[]Withdrawal, *Pagination, error) {
	withdrawals := &[]Withdrawal{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "withdrawals"), nil)
	if err != nil {
		return withdrawals, pagination, err
	}

	if err = c.SendWithAuth(req, withdrawals, pagination); err != nil {
		return withdrawals, pagination, err
	}

	return withdrawals, pagination, nil
}

// GetWithdrawal Show an individual withdrawal.
// Endpoint: GET /accounts/:account_id/withdrawals/:withdrawal_id
func (c *Client) GetWithdrawal(ctx context.Context, accountID string, withdrawalID string) (*Withdrawal, error) {
	withdrawal := &Withdrawal{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "withdrawals", withdrawalID), nil)
	if err != nil {
		return withdrawal, err
	}

	if err = c.SendWithAuth(req, withdrawal); err != nil {
		return withdrawal, err
	}

	return withdrawal, nil
}

// Withdraw Withdraws user-defined amount of funds from a fiat account.
// Endpoint: POST /accounts/:account_id/withdrawals
func (c *Client) Withdraw(ctx context.Context, accountID string, withdrawData Withdraw) (*Withdrawal, error) {
	withdrawal := &Withdrawal{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "withdrawals"), withdrawData)
	if err != nil {
		return withdrawal, err
	}

	if err = c.SendWithAuth(req, withdrawal); err != nil {
		return withdrawal, err
	}

	return withdrawal, nil
}

// CommitWithdrawal Completes a withdrawal that is created in commit: false state.
// Endpoint: POST /accounts/:account_id/withdrawals/:withdrawal_id/commit
func (c *Client) CommitWithdrawal(ctx context.Context, accountID string, withdrawID string) (*Withdrawal, error) {
	withdrawal := &Withdrawal{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "withdrawals", withdrawID, "commit"), nil)
	if err != nil {
		return withdrawal, err
	}

	if err = c.SendWithAuth(req, withdrawal); err != nil {
		return withdrawal, err
	}

	return withdrawal, nil
}
