package client

import "context"

// GotoFileByName returns the files matched by the search query in a given revision, review, project-wide, or service-wide
func (c *Client) GotoFileByName(ctx context.Context, req GotoFileRequestDTO) (*GotoFileResponseDTO, error) {
	var res GotoFileResponseDTO
	if err := c.Call(ctx, "gotoFileByName", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// TextSearch performs full-text search across all commits and files (either service-wide or project-wide)
func (c *Client) TextSearch(ctx context.Context, req TextSearchRequestDTO) (*TextSearchResponseDTO, error) {
	var res TextSearchResponseDTO
	if err := c.Call(ctx, "textSearch", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
