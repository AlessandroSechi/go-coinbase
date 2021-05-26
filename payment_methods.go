package coinbase

import (
	"context"
	"fmt"
)

// ListPaymentMethods Lists current user’s payment methods.
// Endpoint: GET /payment-methods
func (c *Client) ListPaymentMethods(ctx context.Context) (*[]PaymentMethod, *Pagination, error) {
	paymentMethods := &[]PaymentMethod{}

	pagination := &Pagination{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", c.APIBase, "payment-methods"), nil)
	if err != nil {
		return paymentMethods, pagination, err
	}

	if err = c.SendWithAuth(req, paymentMethods, pagination); err != nil {
		return paymentMethods, pagination, err
	}

	return paymentMethods, pagination, nil
}

// ShowPaymentMethod Show current user’s payment method.
// Endpoint: GET /payment-methods/:payment_method_id/
func (c *Client) ShowPaymentMethod(ctx context.Context, paymentMethodID string) (*PaymentMethod, error) {
	paymentMethod := &PaymentMethod{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s/", c.APIBase, "payment-methods", paymentMethodID), nil)
	if err != nil {
		return paymentMethod, err
	}

	if err = c.SendWithAuth(req, paymentMethod); err != nil {
		return paymentMethod, err
	}

	return paymentMethod, nil
}
