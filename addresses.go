package coinbase

import (
	"context"
	"fmt"
)

// ListAddresses Lists addresses for an account.
// Endpoint: GET /accounts/:account_id/addresses
func (c *Client) ListAddresses(ctx context.Context, accountID string) (*[]Address, *Pagination, error) {
	addresses := &[]Address{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "addresses"), nil)
	if err != nil {
		return addresses, pagination, err
	}

	if err = c.SendWithAuth(req, addresses, pagination); err != nil {
		return addresses, pagination, err
	}

	return addresses, pagination, nil
}

// ShowAddress Show an individual address for an account.
// Endpoint: GET /accounts/:account_id/addresses/:address_id
func (c *Client) ShowAddress(ctx context.Context, accountID string, addressID string) (*Address, error) {
	address := &Address{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "addresses", addressID), nil)
	if err != nil {
		return address, err
	}

	if err = c.SendWithAuth(req, address); err != nil {
		return address, err
	}

	return address, nil
}

// ListAddressTransactions List transactions that have been sent to a specific address.
// Endpoint: GET /accounts/:account_id/addresses/:address_id/transactions
func (c *Client) ListAddressTransactions(ctx context.Context, accountID string, addressID string) (*[]Transaction, *Pagination, error) {
	addresses := &[]Transaction{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/%s/%s/%s", c.APIBase, "accounts", accountID, "addresses", addressID, "transactions"), nil)
	if err != nil {
		return addresses, pagination, err
	}

	if err = c.SendWithAuth(req, addresses, pagination); err != nil {
		return addresses, pagination, err
	}

	return addresses, pagination, nil
}

// CreateAddress Creates a new address for an account.
// Endpoint: POST /accounts/:account_id/addresses
func (c *Client) CreateAddress(ctx context.Context, accountID string, addressData CreateAddress) (*Address, error) {
	address := &Address{}

	req, err := c.NewRequest(ctx, "POST", fmt.Sprintf("%s/%s/%s/%s", c.APIBase, "accounts", accountID, "addresses"), addressData)
	if err != nil {
		return address, err
	}

	if err = c.SendWithAuth(req, address); err != nil {
		return address, err
	}

	return address, nil
}
