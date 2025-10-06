package client

import "context"

// GetVcsHostingServices loads VCS Hosting services
func (c *Client) GetVcsHostingServices(ctx context.Context, req VcsHostingRequestDTO) (*VcsHostingResponseDTO, error) {
	var res VcsHostingResponseDTO
	if err := c.Call(ctx, "getVcsHostingServices", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TestVcsConnection checks connection to a given VCS repository
func (c *Client) TestVcsConnection(ctx context.Context, req TestConnectionRequestDTO) (*TestConnectionResponseDTO, error) {
	var res TestConnectionResponseDTO
	if err := c.Call(ctx, "testVcsConnection", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
