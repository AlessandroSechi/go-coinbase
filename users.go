package coinbase

import (
	"context"
	"fmt"
)

// GetUserByID Get any user’s public information with their ID.
// Endpoint: GET /users/:user_id
func (c *Client) GetUserByID(ctx context.Context, userID string) (*User, error) {
	user := &User{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s/%s", c.APIBase, "users", userID), nil)
	if err != nil {
		return user, err
	}

	if err = c.Send(req, user); err != nil {
		return user, err
	}

	return user, nil
}

// GetUser Get current user’s public information
// Endpoint: GET /user
func (c *Client) GetUser(ctx context.Context) (*User, error) {
	user := &User{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", c.APIBase, "user"), nil)
	if err != nil {
		return nil, err
	}

	if err = c.SendWithAuth(req, user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser Modify current user and their preferences.
// Endpoint: PUT /user
func (c *Client) UpdateUser(ctx context.Context, userData UpdateCurrentUser) (*User, error) {
	user := &User{}

	req, err := c.NewRequest(ctx, "PUT", fmt.Sprintf("%s/%s", c.APIBase, "user"), userData)
	if err != nil {
		return user, err
	}

	if err = c.SendWithAuth(req, user); err != nil {
		return user, err
	}

	return user, nil
}