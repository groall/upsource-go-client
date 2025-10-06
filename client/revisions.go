package client

import "context"

// GetHeadRevision returns the revision descriptor for the latest revision - ID, date, commit message, authors, parent revisions, etc.
func (c *Client) GetHeadRevision(ctx context.Context, projectID string) (*RevisionInfoDTO, error) {
	req := ProjectIdDTO{ProjectID: projectID}
	var res RevisionInfoDTO
	if err := c.Call(ctx, "getHeadRevision", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionsListFiltered returns the list of revisions that match the given search query
func (c *Client) GetRevisionsListFiltered(ctx context.Context, req RevisionsListFilteredRequestDTO) (*RevisionDescriptorListDTO, error) {
	var res RevisionDescriptorListDTO
	if err := c.Call(ctx, "getRevisionsListFiltered", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserRevisionsList returns the list of revisions across all projects authored by the specified user and (optionally) matching the specified query
func (c *Client) GetUserRevisionsList(ctx context.Context, req UserRevisionsListRequestDTO) (*RevisionDescriptorListDTO, error) {
	var res RevisionDescriptorListDTO
	if err := c.Call(ctx, "getUserRevisionsList", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionsListUpdate returns the list of revisions since the given revision
func (c *Client) GetRevisionsListUpdate(ctx context.Context, req RevisionsListUpdateRequestDTO) (*RevisionsListUpdateResponseDTO, error) {
	var res RevisionsListUpdateResponseDTO
	if err := c.Call(ctx, "getRevisionsListUpdate", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionInfo returns the revision descriptor - ID, date, commit message, authors, parent revisions, etc.
func (c *Client) GetRevisionInfo(ctx context.Context, req RevisionInProjectDTO) (*RevisionInfoDTO, error) {
	var res RevisionInfoDTO
	if err := c.Call(ctx, "getRevisionInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionChanges returns the list of changes (files that were added, removed, or modified) in a revision
func (c *Client) GetRevisionChanges(ctx context.Context, req RevisionChangesRequestDTO) (*RevisionsDiffDTO, error) {
	var res RevisionsDiffDTO
	if err := c.Call(ctx, "getRevisionChanges", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionBranches returns the list of branches a revision is part of
func (c *Client) GetRevisionBranches(ctx context.Context, req RevisionInProjectDTO) (*RevisionBranchesResponseDTO, error) {
	var res RevisionBranchesResponseDTO
	if err := c.Call(ctx, "getRevisionBranches", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionsList fetches latest revisions for a project
func (c *Client) GetRevisionsList(ctx context.Context, projectID string, limit int) (*RevisionDescriptorListDTO, error) {
	if limit <= 0 {
		limit = 30
	}
	req := RevisionsListRequestDTO{ProjectID: projectID, Limit: int32(limit)}
	var res RevisionDescriptorListDTO
	if err := c.Call(ctx, "getRevisionsList", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionDiscussions returns the discussions in the given revision
func (c *Client) GetRevisionDiscussions(ctx context.Context, req RevisionInProjectDTO) (*DiscussionsInRevisionDTO, error) {
	var res DiscussionsInRevisionDTO
	if err := c.Call(ctx, "getRevisionDiscussions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionReviewInfo returns short review information for a set of revisions
func (c *Client) GetRevisionReviewInfo(ctx context.Context, req RevisionListDTO) (*RevisionReviewInfoListDTO, error) {
	var res RevisionReviewInfoListDTO
	if err := c.Call(ctx, "getRevisionReviewInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetSuggestedReviewsForRevisions returns review suggestions for a set of revisions
func (c *Client) GetSuggestedReviewsForRevisions(ctx context.Context, req RevisionListDTO) (*SuggestedReviewListDTO, error) {
	var res SuggestedReviewListDTO
	if err := c.Call(ctx, "getSuggestedReviewsForRevisions", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetSuggestedReviewForGhosts returns review suggestion for the uncommitted revision
func (c *Client) GetSuggestedReviewForGhosts(ctx context.Context, req UncommittedRevisionDTO) (*ReviewSuggestDTO, error) {
	var res ReviewSuggestDTO
	if err := c.Call(ctx, "getSuggestedReviewForGhosts", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionDiscussionCounters returns discussion counters for a set of revisions
func (c *Client) GetRevisionDiscussionCounters(ctx context.Context, req RevisionDiscussionCountersRequestDTO) (*RevisionDiscussionCountersListDTO, error) {
	var res RevisionDiscussionCountersListDTO
	if err := c.Call(ctx, "getRevisionDiscussionCounters", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionBuildStatus returns build status for revisions
func (c *Client) GetRevisionBuildStatus(ctx context.Context, req RevisionListDTO) (*RevisionBuildStatusListDTO, error) {
	var res RevisionBuildStatusListDTO
	if err := c.Call(ctx, "getRevisionBuildStatus", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionsOwnershipSummary returns the code ownership summary for given revisions
func (c *Client) GetRevisionsOwnershipSummary(ctx context.Context, req RevisionListDTO) (*RevisionsOwnershipSummaryDTO, error) {
	var res RevisionsOwnershipSummaryDTO
	if err := c.Call(ctx, "getRevisionsOwnershipSummary", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetRevisionsExternalInspectionsDiff returns the diff of external inspections compared to the previous run
func (c *Client) GetRevisionsExternalInspectionsDiff(ctx context.Context, req RevisionListDTO) (*RevisionsExternalInspectionsDiffResponseDTO, error) {
	var res RevisionsExternalInspectionsDiffResponseDTO
	if err := c.Call(ctx, "getRevisionsExternalInspectionsDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetMatchingRevisionsForStacktrace returns matching revisions for stacktrace
func (c *Client) GetMatchingRevisionsForStacktrace(ctx context.Context, req StacktraceDTO) (*MatchingRevisionsResponseDTO, error) {
	var res MatchingRevisionsResponseDTO
	if err := c.Call(ctx, "getMatchingRevisionsForStacktrace", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetStacktraceInContextOfRevision returns full paths and vcs commit ids for lines of stacktrace in context of revision
func (c *Client) GetStacktraceInContextOfRevision(ctx context.Context, req StacktraceDTO) (*StacktracePositionsDTO, error) {
	var res StacktracePositionsDTO
	if err := c.Call(ctx, "getStacktraceInContextOfRevision", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
