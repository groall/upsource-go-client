package client

import "context"

// GetAllRoles returns all roles available in Hub
func (c *Client) GetAllRoles(ctx context.Context) (*AllRolesResponseDTO, error) {
	var res AllRolesResponseDTO
	if err := c.Call(ctx, "getAllRoles", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUsersRoles returns users that have any access to the project and role keys associated with each user
func (c *Client) GetUsersRoles(ctx context.Context, req UsersRolesRequestDTO) (*UsersRolesResponseDTO, error) {
	var res UsersRolesResponseDTO
	if err := c.Call(ctx, "getUsersRoles", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AddUserRole adds a user role
func (c *Client) AddUserRole(ctx context.Context, req AddRoleRequestDTO) error {
	return c.Call(ctx, "addUserRole", req, &VoidMessage{})
}

// DeleteUserRole deletes a user role (not the role itself, but the association)
func (c *Client) DeleteUserRole(ctx context.Context, req DeleteRoleRequestDTO) error {
	return c.Call(ctx, "deleteUserRole", req, &VoidMessage{})
}
