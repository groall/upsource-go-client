package client

import "context"

// AddGroupToReview adds a group of participants (reviewers or watchers) to a review
func (c *Client) AddGroupToReview(ctx context.Context, req AddGroupToReviewRequestDTO) (*AddGroupToReviewResponseDTO, error) {
	var res AddGroupToReviewResponseDTO
	if err := c.Call(ctx, "addGroupToReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AddParticipantToReview adds a participant (reviewer or watcher) to a review
func (c *Client) AddParticipantToReview(ctx context.Context, req ParticipantInReviewRequestDTO) error {
	return c.Call(ctx, "addParticipantToReview", req, &VoidMessage{})
}

// AddReviewLabel adds a label to a review
func (c *Client) AddReviewLabel(ctx context.Context, req UpdateReviewLabelRequestDTO) (*UpdateReviewLabelResponseDTO, error) {
	var res UpdateReviewLabelResponseDTO
	if err := c.Call(ctx, "addReviewLabel", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// AddRevisionToReview attaches a revision to a review
func (c *Client) AddRevisionToReview(ctx context.Context, req RevisionsInReviewDTO) error {
	return c.Call(ctx, "addRevisionToReview", req, &VoidMessage{})
}

// CloseReview closes a review
func (c *Client) CloseReview(ctx context.Context, req CloseReviewRequestDTO) (*CloseReviewResponseDTO, error) {
	var res CloseReviewResponseDTO
	if err := c.Call(ctx, "closeReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateReview creates a review
func (c *Client) CreateReview(ctx context.Context, req CreateReviewRequestDTO) (*ReviewDescriptorDTO, error) {
	var res ReviewDescriptorDTO
	if err := c.Call(ctx, "createReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// EditReviewDescription sets review description
func (c *Client) EditReviewDescription(ctx context.Context, req EditReviewDescriptionRequestDTO) (*EditReviewDescriptionResponseDTO, error) {
	var res EditReviewDescriptionResponseDTO
	if err := c.Call(ctx, "editReviewDescription", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCodeReviewPatterns returns all registered code review patterns across all projects
func (c *Client) GetCodeReviewPatterns(ctx context.Context) (*CodeReviewPatternsDTO, error) {
	var res CodeReviewPatternsDTO
	if err := c.Call(ctx, "getCodeReviewPatterns", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCurrentUserCanCloseReview checks if current user can close given review
func (c *Client) GetCurrentUserCanCloseReview(ctx context.Context, req ReviewIdDTO) error {
	return c.Call(ctx, "getCurrentUserCanCloseReview", req, &VoidMessage{})
}

// GetCurrentUserCanDeleteReview checks if current user can delete given review
func (c *Client) GetCurrentUserCanDeleteReview(ctx context.Context, req ReviewIdDTO) error {
	return c.Call(ctx, "getCurrentUserCanDeleteReview", req, &VoidMessage{})
}

// GetFileInReviewSummaryDiff returns the diff (sum of all revisions) of a file in review
func (c *Client) GetFileInReviewSummaryDiff(ctx context.Context, req FileInReviewDiffRequestDTO) (*FileDiffResponseDTO, error) {
	var res FileDiffResponseDTO
	if err := c.Call(ctx, "getFileInReviewSummaryDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileInReviewSummaryInlineChanges returns the diff (sum of all revisions) of a file in review
func (c *Client) GetFileInReviewSummaryInlineChanges(ctx context.Context, req FileInReviewDiffRequestDTO) (*FileInlineDiffResponseDTO, error) {
	var res FileInlineDiffResponseDTO
	if err := c.Call(ctx, "getFileInReviewSummaryInlineChanges", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewDetails returns review details
func (c *Client) GetReviewDetails(ctx context.Context, req ReviewIdDTO) (*ReviewDescriptorDTO, error) {
	var res ReviewDescriptorDTO
	if err := c.Call(ctx, "getReviewDetails", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewInspectionsDiff returns the diff of inspections between two revisions
func (c *Client) GetReviewInspectionsDiff(ctx context.Context, req ReviewInspectionsDiffRequestDTO) (*InspectionsDiffDTO, error) {
	var res InspectionsDiffDTO
	if err := c.Call(ctx, "getReviewInspectionsDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewLabels returns the list of review labels
func (c *Client) GetReviewLabels(ctx context.Context, req LabelsRequestDTO) (*LabelsListDTO, error) {
	var res LabelsListDTO
	if err := c.Call(ctx, "getReviewLabels", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewOwnershipSummary returns the code ownership summary for a given review
func (c *Client) GetReviewOwnershipSummary(ctx context.Context, req ReviewIdDTO) (*ReviewOwnershipSummaryDTO, error) {
	var res ReviewOwnershipSummaryDTO
	if err := c.Call(ctx, "getReviewOwnershipSummary", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewProgress returns participants' progress in a given review
func (c *Client) GetReviewProgress(ctx context.Context, req ReviewIdDTO) (*ReviewProgressDTO, error) {
	var res ReviewProgressDTO
	if err := c.Call(ctx, "getReviewProgress", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviews returns the list of reviews
func (c *Client) GetReviews(ctx context.Context, req ReviewsRequestDTO) (*ReviewListDTO, error) {
	var res ReviewListDTO
	if err := c.Call(ctx, "getReviews", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewSummaryChanges returns the list of changes (sum of all revisions) in a review
func (c *Client) GetReviewSummaryChanges(ctx context.Context, req ReviewSummaryChangesRequestDTO) (*ReviewSummaryChangesResponseDTO, error) {
	var res ReviewSummaryChangesResponseDTO
	if err := c.Call(ctx, "getReviewSummaryChanges", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetReviewSummaryDiscussions returns the review discussions
func (c *Client) GetReviewSummaryDiscussions(ctx context.Context, req ReviewSummaryDiscussionsRequestDTO) (*DiscussionsInFilesDTO, error) {
	var res DiscussionsInFilesDTO
	if err := c.Call(ctx, "getReviewSummaryDiscussions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionsInReview returns the list of revisions in a review
func (c *Client) GetRevisionsInReview(ctx context.Context, req ReviewIdDTO) (*RevisionsInReviewResponseDTO, error) {
	var res RevisionsInReviewResponseDTO
	if err := c.Call(ctx, "getRevisionsInReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetSuggestedRevisionsForReview returns the list of suggested revisions for a particular review
func (c *Client) GetSuggestedRevisionsForReview(ctx context.Context, req ReviewIdDTO) (*SuggestedRevisionListDTO, error) {
	var res SuggestedRevisionListDTO
	if err := c.Call(ctx, "getSuggestedRevisionsForReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RemoveParticipantFromReview removes a participant from a review
func (c *Client) RemoveParticipantFromReview(ctx context.Context, req ParticipantInReviewRequestDTO) error {
	return c.Call(ctx, "removeParticipantFromReview", req, &VoidMessage{})
}

// RemoveReview removes the review
func (c *Client) RemoveReview(ctx context.Context, req ReviewIdDTO) error {
	return c.Call(ctx, "removeReview", req, &VoidMessage{})
}

// RemoveReviewLabel removes a label from a review
func (c *Client) RemoveReviewLabel(ctx context.Context, req UpdateReviewLabelRequestDTO) (*UpdateReviewLabelResponseDTO, error) {
	var res UpdateReviewLabelResponseDTO
	if err := c.Call(ctx, "removeReviewLabel", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RemoveRevisionFromReview removes a revision from a review
func (c *Client) RemoveRevisionFromReview(ctx context.Context, req RevisionsInReviewDTO) error {
	return c.Call(ctx, "removeRevisionFromReview", req, &VoidMessage{})
}

// RenameReview renames a review
func (c *Client) RenameReview(ctx context.Context, req RenameReviewRequestDTO) (*RenameReviewResponseDTO, error) {
	var res RenameReviewResponseDTO
	if err := c.Call(ctx, "renameReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SetFileInReviewReadStatus Track file read status
func (c *Client) SetFileInReviewReadStatus(ctx context.Context, req FileInReviewReadStatusRequestDTO) error {
	return c.Call(ctx, "setFileInReviewReadStatus", req, &VoidMessage{})
}

// SetReviewDeadline sets/clears review due date
func (c *Client) SetReviewDeadline(ctx context.Context, req ReviewDeadlineRequestDTO) error {
	return c.Call(ctx, "setReviewDeadline", req, &VoidMessage{})
}

// SetReviewTargetBranch Merge review: sets target branch
func (c *Client) SetReviewTargetBranch(ctx context.Context, req ReviewTargetBranchDTO) (*SetReviewTargetBranchResponseDTO, error) {
	var res SetReviewTargetBranchResponseDTO
	if err := c.Call(ctx, "setReviewTargetBranch", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// SquashReviewRevisions Rebase the review to a single squashed revision
func (c *Client) SquashReviewRevisions(ctx context.Context, req ReviewIdDTO) error {
	return c.Call(ctx, "squashReviewRevisions", req, &VoidMessage{})
}

// ToggleReviewMuted mutes activities taking place in a certain review or cancels muting that was previously set
func (c *Client) ToggleReviewMuted(ctx context.Context, req ToggleReviewMutedRequestDTO) error {
	return c.Call(ctx, "toggleReviewMuted", req, &VoidMessage{})
}

// UpdateParticipantInReview updates the participant's state
func (c *Client) UpdateParticipantInReview(ctx context.Context, req UpdateParticipantInReviewRequestDTO) (*UpdateParticipantInReviewResponseDTO, error) {
	var res UpdateParticipantInReviewResponseDTO
	if err := c.Call(ctx, "updateParticipantInReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
