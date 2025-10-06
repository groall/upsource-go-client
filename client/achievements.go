package client

import "context"

// UpdateAchievementsLastSeen updates the time when the user last saw the list of his/her achievements
func (c *Client) UpdateAchievementsLastSeen(ctx context.Context, req UserActionNotificationDTO) error {
	return c.Call(ctx, "updateAchievementsLastSeen", req, &VoidMessage{})
}

// GetUnlockedUserAchievements returns the list of achievements the user has unlocked
func (c *Client) GetUnlockedUserAchievements(ctx context.Context, req UserAchievementsRequestDTO) (*AchievementsListDTO, error) {
	var res AchievementsListDTO
	if err := c.Call(ctx, "getUnlockedUserAchievements", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUnreadUnlockedUserAchievementsCount returns a count of unread achievements the user has unlocked
func (c *Client) GetUnreadUnlockedUserAchievementsCount(ctx context.Context) (*UserAchievementsCountDTO, error) {
	var res UserAchievementsCountDTO
	if err := c.Call(ctx, "getUnreadUnlockedUserAchievementsCount", &VoidMessage{}, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetUserAchievements returns the list of all achievements with information about whether the achievement was unlocked by the user or not
func (c *Client) GetUserAchievements(ctx context.Context, req UserAchievementsRequestDTO) (*AchievementsListDTO, error) {
	var res AchievementsListDTO
	if err := c.Call(ctx, "getUserAchievements", req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
