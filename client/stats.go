package client

import "context"

// GetProjectActivity returns project activity distribution over time, i.e. the number of commits per time period
func (c *Client) GetProjectActivity(ctx context.Context, req ProjectActivityRequestDTO) (*ProjectActivityDTO, error) {
	var res ProjectActivityDTO
	if err := c.Call(ctx, "getProjectActivity", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetContributorsDistribution returns contributors distribution over time, i.e. the number of commits per time period per committer
func (c *Client) GetContributorsDistribution(ctx context.Context, req ContributorsDistributionRequestDTO) (*ContributorsDistributionDTO, error) {
	var res ContributorsDistributionDTO
	if err := c.Call(ctx, "getContributorsDistribution", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetResponsibilityDistribution returns commits distribution in a given date range over all committers and modules
func (c *Client) GetResponsibilityDistribution(ctx context.Context, req ResponsibilityDistributionRequestDTO) (*ResponsibilityDistributionDTO, error) {
	var res ResponsibilityDistributionDTO
	if err := c.Call(ctx, "getResponsibilityDistribution", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectCommitters returns all committers throughout the project history
func (c *Client) GetProjectCommitters(ctx context.Context, projectID string) (*ProjectCommittersDTO, error) {
	req := ProjectIdDTO{ProjectID: projectID}
	var res ProjectCommittersDTO
	if err := c.Call(ctx, "getProjectCommitters", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetModulesDistribution returns per-module activity distribution in the project
func (c *Client) GetModulesDistribution(ctx context.Context, req ModulesDistributionRequestDTO) (*ModulesDistributionDTO, error) {
	var res ModulesDistributionDTO
	if err := c.Call(ctx, "getModulesDistribution", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCommitsSummary returns the summary for the specified date range and set of committers: total number of commits, number of commits that aren't part of any module, total number of modules changed
func (c *Client) GetCommitsSummary(ctx context.Context, req CommitsSummaryRequestDTO) (*CommitsSummaryDTO, error) {
	var res CommitsSummaryDTO
	if err := c.Call(ctx, "getCommitsSummary", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetCommitsDetails returns the revisions in the specified module, in the specified date range and for the specified committers
func (c *Client) GetCommitsDetails(ctx context.Context, req CommitsDetailsRequestDTO) (*CommitsDetailsDTO, error) {
	var res CommitsDetailsDTO
	if err := c.Call(ctx, "getCommitsDetails", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileHistoryChart returns the data required to display an animated chart of distribution of individual files in project across "edits - age" plane
func (c *Client) GetFileHistoryChart(ctx context.Context, req FileHistoryChartRequestDTO) (*FileHistoryChartDTO, error) {
	var res FileHistoryChartDTO
	if err := c.Call(ctx, "getFileHistoryChart", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectTreeMap returns the list of files in a project and their content sizes required to build the interactive treemap chart
func (c *Client) GetProjectTreeMap(ctx context.Context, req ProjectTreeMapRequestDTO) (*ProjectTreeMapDTO, error) {
	var res ProjectTreeMapDTO
	if err := c.Call(ctx, "getProjectTreeMap", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectPulse returns the activity pulse for a specified project
func (c *Client) GetProjectPulse(ctx context.Context, req ProjectPulseRequestDTO) (*PulseResponseDTO, error) {
	var res PulseResponseDTO
	if err := c.Call(ctx, "getProjectPulse", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserActivity returns project activity distribution over time, i.e. the number of commits per time period for a specified set of committers
func (c *Client) GetUserActivity(ctx context.Context, req UserActivityRequestDTO) (*UserActivityDTO, error) {
	var res UserActivityDTO
	if err := c.Call(ctx, "getUserActivity", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserPulse returns the activity pulse of a given user across all projects (`allValues` time series is always empty)
func (c *Client) GetUserPulse(ctx context.Context, req UserPulseRequestDTO) (*PulseResponseDTO, error) {
	var res PulseResponseDTO
	if err := c.Call(ctx, "getUserPulse", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
