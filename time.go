package coinbase

import (
	"context"
	"fmt"
)

// GetTime Get the API server time.
// Endpoint: GET /time
func (c *Client) GetTime(ctx context.Context) (*Time, error) {
	time := &Time{}

	req, err := c.NewRequest(ctx, "GET", fmt.Sprintf("%s/%s", c.APIBase, "time"), nil)
	if err != nil {
		return time, err
	}

	if err = c.SendWithAuth(req, time); err != nil {
		return time, err
	}

	return time, nil
}
