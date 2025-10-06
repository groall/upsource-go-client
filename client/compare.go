package client

import "context"

// GetRevisionsDiff returns the list of changes between any two revisions
func (c *Client) GetRevisionsDiff(ctx context.Context, req RevisionsDiffRequestDTO) (*RevisionsDiffDTO, error) {
	var res RevisionsDiffDTO
	if err := c.Call(ctx, "getRevisionsDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCompareInfo Compare page
func (c *Client) GetCompareInfo(ctx context.Context, req CompareRequestDTO) (*CompareInfoDTO, error) {
	var res CompareInfoDTO
	if err := c.Call(ctx, "getCompareInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCompareGraph Revision graph for compare page
func (c *Client) GetCompareGraph(ctx context.Context, req RevisionsDiffRequestDTO) (*CompareGraphDTO, error) {
	var res CompareGraphDTO
	if err := c.Call(ctx, "getCompareGraph", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
