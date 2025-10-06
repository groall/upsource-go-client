package client

import "context"

// GetAvailableIssueTrackerProviders returns the list of available issue trackers
func (c *Client) GetAvailableIssueTrackerProviders(ctx context.Context) (*IssueTrackerProvidersListDTO, error) {
	var res IssueTrackerProvidersListDTO
	if err := c.Call(ctx, "getAvailableIssueTrackerProviders", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateIssueFromDiscussion creates an issue from the discussion in a given issue tracker
func (c *Client) CreateIssueFromDiscussion(ctx context.Context, req CreateIssueFromDiscussionRequestDTO) (*IssueIdDTO, error) {
	var res IssueIdDTO
	if err := c.Call(ctx, "createIssueFromDiscussion", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetIssueInfo returns detailed issue information
func (c *Client) GetIssueInfo(ctx context.Context, req IssueInfoRequestDTO) (*IssueInfoDTO, error) {
	var res IssueInfoDTO
	if err := c.Call(ctx, "getIssueInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateIssueFromReview creates an issue from the review in a given issue tracker
func (c *Client) CreateIssueFromReview(ctx context.Context, req CreateIssueFromReviewRequestDTO) (*IssueIdDTO, error) {
	var res IssueIdDTO
	if err := c.Call(ctx, "createIssueFromReview", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
