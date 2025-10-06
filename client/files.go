package client

import "context"

// GetFileMeta returns meta information about a file (is deleted, is latest revision, etc.)
func (c *Client) GetFileMeta(ctx context.Context, req FileInRevisionDTO) (*FileMetaResponseDTO, error) {
	var res FileMetaResponseDTO
	if err := c.Call(ctx, "getFileMeta", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileAnnotation returns the line-by-line file annotation
func (c *Client) GetFileAnnotation(ctx context.Context, req FileInRevisionDTO) (*FileAnnotationResponseDTO, error) {
	var res FileAnnotationResponseDTO
	if err := c.Call(ctx, "getFileAnnotation", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileContributors returns the file contributors
func (c *Client) GetFileContributors(ctx context.Context, req FileInRevisionDTO) (*FileContributorsResponseDTO, error) {
	var res FileContributorsResponseDTO
	if err := c.Call(ctx, "getFileContributors", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileFragmentAuthors returns the authors of the given file fragment
func (c *Client) GetFileFragmentAuthors(ctx context.Context, req FileFragmentAuthorsRequestDTO) (*UsersListDTO, error) {
	var res UsersListDTO
	if err := c.Call(ctx, "getFileFragmentAuthors", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileHistory returns the history (list of revision IDs and change types) of the file
func (c *Client) GetFileHistory(ctx context.Context, req FileHistoryRequestDTO) (*FileHistoryResponseDTO, error) {
	var res FileHistoryResponseDTO
	if err := c.Call(ctx, "getFileHistory", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileDiff returns the diff (changed lines and ranges) between the given revisions of a file
func (c *Client) GetFileDiff(ctx context.Context, req FileDiffRequestDTO) (*FileDiffResponseDTO, error) {
	var res FileDiffResponseDTO
	if err := c.Call(ctx, "getFileDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileInlineDiff returns the inline diff (changed lines and ranges, line numbers mapping) for the given file
func (c *Client) GetFileInlineDiff(ctx context.Context, req FileDiffRequestDTO) (*FileInlineDiffResponseDTO, error) {
	var res FileInlineDiffResponseDTO
	if err := c.Call(ctx, "getFileInlineDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileMergeInlineDiff returns the inline diff of a file after the merge with the base branch
func (c *Client) GetFileMergeInlineDiff(ctx context.Context, req FileMergeInlineDiffRequestDTO) (*FileInlineDiffResponseDTO, error) {
	var res FileInlineDiffResponseDTO
	if err := c.Call(ctx, "getFileMergeInlineDiff", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetFileContent returns the contents of the given file
func (c *Client) GetFileContent(ctx context.Context, req FileInRevisionDTO) (*FileContentResponseDTO, error) {
	var res FileContentResponseDTO
	if err := c.Call(ctx, "getFileContent", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
