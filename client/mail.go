package client

import "context"

// TestPOP3Mailbox checks availability of a given POP3 mailbox
func (c *Client) TestPOP3Mailbox(ctx context.Context, req TestPOP3MailboxRequestDTO) (*TestPOP3MailboxResponseDTO, error) {
	var res TestPOP3MailboxResponseDTO
	if err := c.Call(ctx, "testPOP3Mailbox", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetEmailBounces returns the list of bounced emails
func (c *Client) GetEmailBounces(ctx context.Context) (*EmailBouncesResponseDTO, error) {
	var res EmailBouncesResponseDTO
	if err := c.Call(ctx, "getEmailBounces", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ResetEmailBounces resets the list of bounced emails
func (c *Client) ResetEmailBounces(ctx context.Context) error {
	return c.Call(ctx, "resetEmailBounces", &VoidMessage{}, &VoidMessage{})
}
