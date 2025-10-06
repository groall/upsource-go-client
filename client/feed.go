package client

import "context"

// GetFeed returns the news feed
func (c *Client) GetFeed(ctx context.Context, req FeedRequestDTO) (*FeedDTO, error) {
	var res FeedDTO
	if err := c.Call(ctx, "getFeed", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
