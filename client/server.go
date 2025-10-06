package client

import "context"

// GetReadOnlyMode checks if server is in read-only mode
func (c *Client) GetReadOnlyMode(ctx context.Context) (*ReadOnlyModeDTO, error) {
	var res ReadOnlyModeDTO
	if err := c.Call(ctx, "getReadOnlyMode", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetReadOnlyMode enables/disables the read-only mode
func (c *Client) SetReadOnlyMode(ctx context.Context, req ReadOnlyModeDTO) error {
	return c.Call(ctx, "setReadOnlyMode", req, &VoidMessage{})
}

// ExportData exports user-generated data - reviews, discussions, settings
func (c *Client) ExportData(ctx context.Context) (*ExportDataResponseDTO, error) {
	var res ExportDataResponseDTO
	if err := c.Call(ctx, "exportData", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
