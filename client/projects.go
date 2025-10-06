package client

import "context"

// GetAllProjects wraps getAllProjects. If ids is empty, returns all accessible projects
func (c *Client) GetAllProjects(ctx context.Context, ids []string) (*ShortProjectInfoListDTO, error) {
	var req ProjectIdListDTO
	if len(ids) > 0 {
		req.ProjectID = ids
	}
	var res ShortProjectInfoListDTO
	if err := c.Call(ctx, "getAllProjects", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// FindProjects returns the list of projects  matching a given search criteria
func (c *Client) FindProjects(ctx context.Context, req FindProjectsRequestDTO) (*ShortProjectInfoListDTO, error) {
	var res ShortProjectInfoListDTO
	if err := c.Call(ctx, "findProjects", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectInfo returns basic information about the project - name, project model, latest revision, etc.
func (c *Client) GetProjectInfo(ctx context.Context, projectID string) (*ProjectInfoDTO, error) {
	req := ProjectIdDTO{ProjectID: projectID}
	var res ProjectInfoDTO
	if err := c.Call(ctx, "getProjectInfo", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectVcsLinks returns VCS repository URL(s) for a given project
func (c *Client) GetProjectVcsLinks(ctx context.Context, projectID string) (*VcsRepoListDTO, error) {
	req := ProjectIdDTO{ProjectID: projectID}
	var res VcsRepoListDTO
	if err := c.Call(ctx, "getProjectVcsLinks", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectReadme returns README (README.md) contents from the latest revision
func (c *Client) GetProjectReadme(ctx context.Context, projectID string) (*ProjectReadmeResponseDTO, error) {
	req := ProjectIdDTO{ProjectID: projectID}
	var res ProjectReadmeResponseDTO
	if err := c.Call(ctx, "getProjectReadme", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProjectSubtree returns the project structure tree starting from the given file
func (c *Client) GetProjectSubtree(ctx context.Context, req FileInRevisionDTO) (*ProjectItemsListDTO, error) {
	var res ProjectItemsListDTO
	if err := c.Call(ctx, "getProjectSubtree", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateProject creates a project
func (c *Client) CreateProject(ctx context.Context, req CreateProjectRequestDTO) error {
	return c.Call(ctx, "createProject", req, &VoidMessage{})
}

// LoadProject loads project settings
func (c *Client) LoadProject(ctx context.Context, req ProjectIdDTO) (*ProjectSettingsDTO, error) {
	var res ProjectSettingsDTO
	if err := c.Call(ctx, "loadProject", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// EditProject updates project settings
func (c *Client) EditProject(ctx context.Context, req EditProjectRequestDTO) error {
	return c.Call(ctx, "editProject", req, &VoidMessage{})
}

// DeleteProject deletes a project
func (c *Client) DeleteProject(ctx context.Context, req ProjectIdDTO) error {
	return c.Call(ctx, "deleteProject", req, &VoidMessage{})
}

// ResetProject completely reset a project
func (c *Client) ResetProject(ctx context.Context, req ProjectIdDTO) error {
	return c.Call(ctx, "resetProject", req, &VoidMessage{})
}

// GetProjectConfigurationParameters loads project settings
func (c *Client) GetProjectConfigurationParameters(ctx context.Context) (*ProjectConfigurationResponseDTO, error) {
	var res ProjectConfigurationResponseDTO
	if err := c.Call(ctx, "getProjectConfigurationParameters", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
