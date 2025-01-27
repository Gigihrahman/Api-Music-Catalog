package trackactivities

import (
	"api-music/internal/models/trackacktivities"
	"context"
)

func (r *repository) Create(ctx context.Context, model trackacktivities.TrackActivity) error {
	return r.db.Create(&model).Error
}

func (r *repository) Update(ctx context.Context, model trackacktivities.TrackActivity) error {
	return r.db.Save(&model).Error
}

func (r *repository) Get(ctx context.Context, UserID uint, spotifyID string) (*trackacktivities.TrackActivity, error) {
	activity := trackacktivities.TrackActivity{}
	res := r.db.Where("user_id = ?", UserID).Where("spotify_id = ?", spotifyID).First(&activity)
	if res.Error != nil {
		return nil, res.Error
	}
	return &activity, nil
}

func (r *repository) GetBulkSpotifyIDs(ctx context.Context, UserID uint, spotifyIDs []string) (map[string]trackacktivities.TrackActivity, error) {
	activities := make([]trackacktivities.TrackActivity, 0)
	res := r.db.Where("user_id = ?", UserID).Where("spotify_id IN ?", spotifyIDs).Find(&activities)
	if res.Error != nil {
		return nil, res.Error
	}
	result := make(map[string]trackacktivities.TrackActivity, 0)
	for _, activity := range activities {
		result[activity.SpotifyID] = activity
	}
	return result, nil
}
