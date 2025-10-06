package client

import "context"

// GetLabels returns the list of discussion labels
func (c *Client) GetLabels(ctx context.Context, req LabelsRequestDTO) (*LabelsListDTO, error) {
	var res LabelsListDTO
	if err := c.Call(ctx, "getLabels", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateOrEditLabel creates/updates a discussion label in a project
func (c *Client) CreateOrEditLabel(ctx context.Context, req EditLabelRequestDTO) (*LabelDTO, error) {
	var res LabelDTO
	if err := c.Call(ctx, "createOrEditLabel", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// CreateOrEditReviewLabel creates/updates a review label in a project
func (c *Client) CreateOrEditReviewLabel(ctx context.Context, req EditLabelRequestDTO) (*LabelDTO, error) {
	var res LabelDTO
	if err := c.Call(ctx, "createOrEditReviewLabel", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// HidePredefinedLabels hides predefined discussion labels in a project
func (c *Client) HidePredefinedLabels(ctx context.Context, req HidePredefinedLabelsRequestDTO) error {
	return c.Call(ctx, "hidePredefinedLabels", req, &VoidMessage{})
}

// HidePredefinedReviewLabels hides predefined review labels in a project
func (c *Client) HidePredefinedReviewLabels(ctx context.Context, req HidePredefinedLabelsRequestDTO) error {
	return c.Call(ctx, "hidePredefinedReviewLabels", req, &VoidMessage{})
}

// DeleteLabel deletes a discussion label from a project
func (c *Client) DeleteLabel(ctx context.Context, req EditLabelRequestDTO) error {
	return c.Call(ctx, "deleteLabel", req, &VoidMessage{})
}

// DeleteReviewLabel deletes a review label from a project
func (c *Client) DeleteReviewLabel(ctx context.Context, req EditLabelRequestDTO) error {
	return c.Call(ctx, "deleteReviewLabel", req, &VoidMessage{})
}

// MergeLabels merges two or more discussion labels into one
func (c *Client) MergeLabels(ctx context.Context, req EditLabelsRequestDTO) (*LabelDTO, error) {
	var res LabelDTO
	if err := c.Call(ctx, "mergeLabels", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
