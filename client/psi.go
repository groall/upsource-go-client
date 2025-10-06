package client

import "context"

// GetFilePsi returns the semantic markup of the given file
func (c *Client) GetFilePsi(ctx context.Context, req FilePsiRequestDTO) (*FilePsiResponseDTO, error) {
	var res FilePsiResponseDTO
	if err := c.Call(ctx, "getFilePsi", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetStubNavigationRange returns the text range of the given PSI element
func (c *Client) GetStubNavigationRange(ctx context.Context, req StubIdDTO) (*NavigationTargetItemDTO, error) {
	var res NavigationTargetItemDTO
	if err := c.Call(ctx, "getStubNavigationRange", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetElementDescription returns the text representation of the given PSI element
func (c *Client) GetElementDescription(ctx context.Context, req PsiElementIdDTO) (*TargetDescriptionDTO, error) {
	var res TargetDescriptionDTO
	if err := c.Call(ctx, "getElementDescription", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetElementDocumentation returns the documentation (e.g. Javadoc) for the given PSI element
func (c *Client) GetElementDocumentation(ctx context.Context, req PsiElementIdDTO) (*ElementDocumentationDTO, error) {
	var res ElementDocumentationDTO
	if err := c.Call(ctx, "getElementDocumentation", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// FindUsages returns the list of usages for the given PSI element
func (c *Client) FindUsages(ctx context.Context, req PsiElementIdDTO) (*FindUsagesResponseDTO, error) {
	var res FindUsagesResponseDTO
	if err := c.Call(ctx, "findUsages", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUsagesDiff returns the usages diff for the given PSI element between two revisions
func (c *Client) GetUsagesDiff(ctx context.Context, req UsagesDiffRequestDTO) (*UsagesDiffResponseDTO, error) {
	var res UsagesDiffResponseDTO
	if err := c.Call(ctx, "getUsagesDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// FindHierarchy returns the list of inheritors and ancestors for the given PSI element
func (c *Client) FindHierarchy(ctx context.Context, req PsiElementIdDTO) (*FindHierarchyResultDTO, error) {
	var res FindHierarchyResultDTO
	if err := c.Call(ctx, "findHierarchy", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
