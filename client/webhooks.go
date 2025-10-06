package client

import "context"

// SetProjectWebhooks updates the set of webhooks in a project
func (c *Client) SetProjectWebhooks(ctx context.Context, req SetProjectWebhooksRequestDTO) error {
	return c.Call(ctx, "setProjectWebhooks", req, &VoidMessage{})
}
