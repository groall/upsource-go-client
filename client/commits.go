package client

import "context"

// FindCommits finds commit(s) with the given commit hash
func (c *Client) FindCommits(ctx context.Context, req FindCommitsRequestDTO) (*FindCommitsResponseDTO, error) {
	var res FindCommitsResponseDTO
	if err := c.Call(ctx, "findCommits", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
