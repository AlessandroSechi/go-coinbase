package coinbase

import (
	"context"
	"fmt"
)

// ListTransactions Lists account’s transactions.
// Endpoint: GET /accounts/:account_id/transactions
func (c *Client) ListTransactions(ctx context.Context, accountID string) (*[]Transaction, *Pagination, error) {
	transactions := &[]Transaction{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions"), nil)
	if err != nil {
		return transactions, pagination, err
	}

	if err = c.SendWithAuth(req, transactions, pagination); err != nil {
		return transactions, pagination, err
	}

	return transactions, pagination, nil
}

// GetTransaction Show an individual transaction for an account.
// Endpoint: GET /accounts/:account_id/transactions/:transaction_id
func (c *Client) GetTransaction(ctx context.Context, accountID string, transactionID string) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions", transactionID), nil)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

// SendMoney Send funds to a bitcoin address, bitcoin cash address, litecoin address, ethereum address, or email address.
// Endpoint: POST /accounts/:account_id/transactions
func (c *Client) SendMoney(ctx context.Context, accountID string, sendData SendMoney) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions"), sendData)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

// TransferMoney Transfer bitcoin, bitcoin cash, litecoin or ethereum between two of a user’s accounts.
// Endpoint: POST /accounts/:account_id/transactions
func (c *Client) TransferMoney(ctx context.Context, accountID string, transferData TransferMoney) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions"), transferData)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

// RequestMoney Requests money from an email address.
// Endpoint: POST /accounts/:account_id/transactions
func (c *Client) RequestMoney(ctx context.Context, accountID string, requestData RequestMoney) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions"), requestData)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

// CompleteRequestMoney Lets the recipient of a money request complete the request by sending money to the user who requested the money.
// Endpoint: POST /accounts/:account_id/transactions/:transaction_id/complete
func (c *Client) CompleteRequestMoney(ctx context.Context, accountID string, transactionID string) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions", transactionID, "complete"), nil)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

// ResendRequestMoney Lets the user resend a money request.
// Endpoint: POST /accounts/:account_id/transactions/:transaction_id/resend
func (c *Client) ResendRequestMoney(ctx context.Context, accountID string, transactionID string) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions", transactionID, "resend"), nil)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}

// CancelRequestMoney Lets a user cancel a money request.
// Endpoint: DELETE /accounts/:account_id/transactions/:transaction_id
func (c *Client) CancelRequestMoney(ctx context.Context, accountID string, transactionID string) (*Transaction, error) {
	transaction := &Transaction{}

	req, err := c.NewRequest(ctx, "DELETE", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "transactions", transactionID), nil)
	if err != nil {
		return transaction, err
	}

	if err = c.SendWithAuth(req, transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}
