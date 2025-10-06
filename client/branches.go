package client

import "context"

// FindBranches performs project-wide search by branch name
func (c *Client) FindBranches(ctx context.Context, req FindBranchRequestDTO) (*FindBranchResponseDTO, error) {
	var res FindBranchResponseDTO
	if err := c.Call(ctx, "findBranches", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetBranchDiff returns the list of changes in the given branch
func (c *Client) GetBranchDiff(ctx context.Context, req BranchRequestDTO) (*RevisionsDiffDTO, error) {
	var res RevisionsDiffDTO
	if err := c.Call(ctx, "getBranchDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetBranchInfo Branch page
func (c *Client) GetBranchInfo(ctx context.Context, req BranchRequestDTO) (*BranchInfoDTO, error) {
	var res BranchInfoDTO
	if err := c.Call(ctx, "getBranchInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetBranchGraph Revision graph for branch page
func (c *Client) GetBranchGraph(ctx context.Context, req BranchRequestDTO) (*BranchGraphDTO, error) {
	var res BranchGraphDTO
	if err := c.Call(ctx, "getBranchGraph", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetBranches returns the list of branches in a project
func (c *Client) GetBranches(ctx context.Context, req BranchesRequestDTO) (*BranchListDTO, error) {
	var res BranchListDTO
	if err := c.Call(ctx, "getBranches", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// StartBranchTracking initiates branch tracking for a given review
func (c *Client) StartBranchTracking(ctx context.Context, req StartBranchTrackingRequestDTO) error {
	return c.Call(ctx, "startBranchTracking", req, &VoidMessage{})
}

// StopBranchTracking stops branch tracking for a given review
func (c *Client) StopBranchTracking(ctx context.Context, req StopBranchTrackingRequestDTO) error {
	return c.Call(ctx, "stopBranchTracking", req, &VoidMessage{})
}
