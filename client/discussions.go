package client

import "context"

// RemoveDiscussionLabel removes a label from a discussion
func (c *Client) RemoveDiscussionLabel(ctx context.Context, req UpdateDiscussionLabelRequestDTO) error {
	return c.Call(ctx, "removeDiscussionLabel", req, &VoidMessage{})
}

// AddComment adds a comment to the discussion
func (c *Client) AddComment(ctx context.Context, req AddCommentRequestDTO) (*CommentDTO, error) {
	var res CommentDTO
	if err := c.Call(ctx, "addComment", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateComment updates the comment
func (c *Client) UpdateComment(ctx context.Context, req UpdateCommentRequestDTO) (*UpdateCommentResponseDTO, error) {
	var res UpdateCommentResponseDTO
	if err := c.Call(ctx, "updateComment", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RemoveComment removes the comment
func (c *Client) RemoveComment(ctx context.Context, req RemoveCommentRequestDTO) (*RemoveCommentResponseDTO, error) {
	var res RemoveCommentResponseDTO
	if err := c.Call(ctx, "removeComment", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateTaskList updates the task list embedded in a comment
func (c *Client) UpdateTaskList(ctx context.Context, req UpdateTaskListRequestDTO) (*UpdateCommentResponseDTO, error) {
	var res UpdateCommentResponseDTO
	if err := c.Call(ctx, "updateTaskList", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ToggleReaction adds or removes a reaction to a specified target
func (c *Client) ToggleReaction(ctx context.Context, req ToggleReactionRequestDTO) error {
	return c.Call(ctx, "toggleReaction", req, &VoidMessage{})
}

// GetProjectDiscussions returns the discussions in the project, matching given query
func (c *Client) GetProjectDiscussions(ctx context.Context, req DiscussionsInProjectRequestDTO) (*DiscussionsInProjectDTO, error) {
	var res DiscussionsInProjectDTO
	if err := c.Call(ctx, "getProjectDiscussions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileDiscussions returns the discussions in the given file
func (c *Client) GetFileDiscussions(ctx context.Context, req FileInRevisionDTO) (*DiscussionsInFileDTO, error) {
	var res DiscussionsInFileDTO
	if err := c.Call(ctx, "getFileDiscussions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetInlineDiscussionsInRevision returns the discussions in the given revision
func (c *Client) GetInlineDiscussionsInRevision(ctx context.Context, req RevisionInProjectDTO) (*DiscussionsInFilesDTO, error) {
	var res DiscussionsInFilesDTO
	if err := c.Call(ctx, "getInlineDiscussionsInRevision", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateDiscussion creates a new discussion
func (c *Client) CreateDiscussion(ctx context.Context, req CreateDiscussionRequestDTO) (*DiscussionInFileDTO, error) {
	var res DiscussionInFileDTO
	if err := c.Call(ctx, "createDiscussion", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// ResolveDiscussion updates a discussion
func (c *Client) ResolveDiscussion(ctx context.Context, req ResolveDiscussionRequestDTO) (*ResolveDiscussionResponseDTO, error) {
	var res ResolveDiscussionResponseDTO
	if err := c.Call(ctx, "resolveDiscussion", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCurrentUserCanResolveDiscussion checks if current user can resolve the given discussion
func (c *Client) GetCurrentUserCanResolveDiscussion(ctx context.Context, req DiscussionIdDTO) error {
	return c.Call(ctx, "getCurrentUserCanResolveDiscussion", req, &VoidMessage{})
}

// AddDiscussionLabel adds a label to a discussion
func (c *Client) AddDiscussionLabel(ctx context.Context, req UpdateDiscussionLabelRequestDTO) (*LabelDTO, error) {
	var res LabelDTO
	if err := c.Call(ctx, "addDiscussionLabel", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// StarDiscussion stars a discussion
func (c *Client) StarDiscussion(ctx context.Context, req UpdateDiscussionFlagRequestDTO) error {
	return c.Call(ctx, "starDiscussion", req, &VoidMessage{})
}

// ReadDiscussion toggles the read/unread state of a discussion
func (c *Client) ReadDiscussion(ctx context.Context, req UpdateDiscussionFlagRequestDTO) error {
	return c.Call(ctx, "readDiscussion", req, &VoidMessage{})
}
