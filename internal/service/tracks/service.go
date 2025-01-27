package tracks

import (
	"api-music/internal/models/trackacktivities"
	"api-music/internal/repository/spotify"
	"context"
)

//go:generate mockgen -source=service.go -destination=service_mock_test.go -package=tracks
type spotifyOutbond interface {
	Search(ctx context.Context, query string, limit, offset int) (*spotify.SpotifySearchResponse, error)
}
type trackActivitiesRepository interface {
	Create(ctx context.Context, model trackacktivities.TrackActivity) error
	Update(ctx context.Context, model trackacktivities.TrackActivity) error
	Get(ctx context.Context, UserID uint, spotifyID string) (*trackacktivities.TrackActivity, error)
	GetBulkSpotifyIDs(ctx context.Context, UserID uint, spotifyIDs []string) (map[string]trackacktivities.TrackActivity, error)
}
type service struct {
	spotifyOutbond      spotifyOutbond
	trackActivitiesRepo trackActivitiesRepository
}

func NewService(spotifyOutbond spotifyOutbond, trackActivitiesRepo trackActivitiesRepository) *service {
	return &service{spotifyOutbond: spotifyOutbond, trackActivitiesRepo: trackActivitiesRepo}
}
