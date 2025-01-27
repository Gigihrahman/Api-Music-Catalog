package tracks

import (
	"api-music/internal/models/trackacktivities"
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func (s *service) UpsertTrackActivities(ctx context.Context, userID uint, request trackacktivities.TrackActivityRequest) error {
	activity, err := s.trackActivitiesRepo.Get(ctx, userID, request.SpotifyID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get record database")
		return err
	}
	if err == gorm.ErrRecordNotFound || activity == nil {
		err = s.trackActivitiesRepo.Create(ctx, trackacktivities.TrackActivity{
			UserID:    userID,
			SpotifyID: request.SpotifyID,
			IsLiked:   request.IsLiked,
			CreatedBy: fmt.Sprintf("%d", userID),
			UpdatedBy: fmt.Sprintf("%d", userID),
		})
		if err != nil {
			log.Error().Err(err).Msg("error create record to database")
			return err
		}
		return nil
	}
	activity.IsLiked = request.IsLiked
	err = s.trackActivitiesRepo.Update(ctx, *activity)
	if err != nil {
		log.Error().Err(err).Msg("error update record to database")
		return err
	}
	return nil
}
