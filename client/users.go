package client

import "context"

// AcceptUserAgreement accepts the end user agreement for a given user
func (c *Client) AcceptUserAgreement(ctx context.Context) error {
	return c.Call(ctx, "acceptUserAgreement", &VoidMessage{}, &VoidMessage{})
}

// BindVcsUsername maps a VCS username/email to a Hub account
func (c *Client) BindVcsUsername(ctx context.Context, req BindVcsUsernameRequestDTO) error {
	return c.Call(ctx, "bindVcsUsername", req, &VoidMessage{})
}

// FindUsers returns the list of users matching a given search criteria
func (c *Client) FindUsers(ctx context.Context, req FindUsersRequestDTO) (*UserInfoResponseDTO, error) {
	var res UserInfoResponseDTO
	if err := c.Call(ctx, "findUsers", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCurrentUser returns the user info for a currently logged-in user
func (c *Client) GetCurrentUser(ctx context.Context) (*CurrentUserResponseDTO, error) {
	var res CurrentUserResponseDTO
	if err := c.Call(ctx, "getCurrentUser", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserAgreementText returns the text of the end user agreement
func (c *Client) GetUserAgreementText(ctx context.Context) (*UserAgreementTextDTO, error) {
	var res UserAgreementTextDTO
	if err := c.Call(ctx, "getUserAgreementText", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserGroupsByIds returns the list of user groups by given IDs
func (c *Client) GetUserGroupsByIds(ctx context.Context, req UserGroupsIdsListDTO) (*UserGroupsListDTO, error) {
	var res UserGroupsListDTO
	if err := c.Call(ctx, "getUserGroupsByIds", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserInfo returns user info for given users
func (c *Client) GetUserInfo(ctx context.Context, req UserInfoRequestDTO) (*UserInfoResponseDTO, error) {
	var res UserInfoResponseDTO
	if err := c.Call(ctx, "getUserInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserProjects returns the list of projects the specified user contributed to
func (c *Client) GetUserProjects(ctx context.Context, req UserProjectsRequestDTO) (*UserProjectsResponseDTO, error) {
	var res UserProjectsResponseDTO
	if err := c.Call(ctx, "getUserProjects", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUsersForMention returns the suggested users to be mentioned in a comment
func (c *Client) GetUsersForMention(ctx context.Context, req UsersForMentionRequestDTO) (*UsersListDTO, error) {
	var res UsersListDTO
	if err := c.Call(ctx, "getUsersForMention", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUsersForReview returns the suggested reviewers for a given review
func (c *Client) GetUsersForReview(ctx context.Context, req UsersForReviewRequestDTO) (*UsersForReviewDTO, error) {
	var res UsersForReviewDTO
	if err := c.Call(ctx, "getUsersForReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUsersPresenceInfo returns presence info for given users
func (c *Client) GetUsersPresenceInfo(ctx context.Context, req UserInfoRequestDTO) (*UsersPresenceInfoResponseDTO, error) {
	var res UsersPresenceInfoResponseDTO
	if err := c.Call(ctx, "getUsersPresenceInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// InviteUser searches for user in Hub by an email, and sends an invitation if not found
func (c *Client) InviteUser(ctx context.Context, req InviteUserRequestDTO) (*InviteUserResponseDTO, error) {
	var res InviteUserResponseDTO
	if err := c.Call(ctx, "inviteUser", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetUserAbsence sets or clears a user absence
func (c *Client) SetUserAbsence(ctx context.Context, req UserAbsenceRequestDTO) error {
	return c.Call(ctx, "setUserAbsence", req, &VoidMessage{})
}

// UpdateUserTimezone updates user timezone
func (c *Client) UpdateUserTimezone(ctx context.Context, req UpdateUserTimezoneDTO) error {
	return c.Call(ctx, "updateUserTimezone", req, &VoidMessage{})
}

// GetProjectUserGroups returns the list of non-empty user groups relevant to a given project
func (c *Client) GetProjectUserGroups(ctx context.Context, req ProjectUserGroupsRequestDTO) (*UserGroupsListDTO, error) {
	var res UserGroupsListDTO
	if err := c.Call(ctx, "getProjectUserGroups", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
