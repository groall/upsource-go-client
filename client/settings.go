package client

import "context"

// GetUserSetting returns the value of a user setting by name
func (c *Client) GetUserSetting(ctx context.Context, req GetSettingRequestDTO) (*GetSettingResponseDTO, error) {
	var res GetSettingResponseDTO
	if err := c.Call(ctx, "getUserSetting", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetUserSetting updates the value of a user setting by name
func (c *Client) SetUserSetting(ctx context.Context, req SetSettingRequestDTO) error {
	return c.Call(ctx, "setUserSetting", req, &VoidMessage{})
}

// GetClusterSetting returns the value of a cluster-wide setting by name
func (c *Client) GetClusterSetting(ctx context.Context, req GetSettingRequestDTO) (*GetSettingResponseDTO, error) {
	var res GetSettingResponseDTO
	if err := c.Call(ctx, "getClusterSetting", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetClusterSetting updates the value of a cluster-wide setting by name
func (c *Client) SetClusterSetting(ctx context.Context, req SetSettingRequestDTO) error {
	return c.Call(ctx, "setClusterSetting", req, &VoidMessage{})
}

// SetMotto sets the server motto
func (c *Client) SetMotto(ctx context.Context, req SetMottoRequestDTO) error {
	return c.Call(ctx, "setMotto", req, &VoidMessage{})
}

// GetProjectSetting returns the value of a project setting by name
func (c *Client) GetProjectSetting(ctx context.Context, req GetProjectSettingRequestDTO) (*GetSettingResponseDTO, error) {
	var res GetSettingResponseDTO
	if err := c.Call(ctx, "getProjectSetting", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetProjectSetting updates the value of a project setting by name
func (c *Client) SetProjectSetting(ctx context.Context, req SetProjectSettingRequestDTO) error {
	return c.Call(ctx, "setProjectSetting", req, &VoidMessage{})
}
