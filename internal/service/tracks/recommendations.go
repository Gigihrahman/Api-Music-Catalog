package tracks

import (
	"api-music/internal/models/spotify"
	"api-music/internal/models/trackacktivities"
	spotifyRepo "api-music/internal/repository/spotify"
	"context"

	"github.com/rs/zerolog/log"
)

func (s *service) GetRecommendation(ctx context.Context, userID uint, limit int, trackID string) (*spotify.RecommendationsResponse, error) {
	trackDetails, err := s.spotifyOutbond.GetRecommendation(ctx, limit, trackID)
	if err != nil {
		log.Error().Err(err).Msg("Error get recomendation from spotify outbond")
		return nil, err
	}
	trackIDs := make([]string, len(trackDetails.Tracks))

	for idx, item := range trackDetails.Tracks {
		trackIDs[idx] = item.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from db")
		return nil, err
	}

	return modelToRecommendationsresponse(trackDetails, trackActivities), nil
}

func modelToRecommendationsresponse(data *spotifyRepo.SpotifyRecommendationResponse, mapTrackActivities map[string]trackacktivities.TrackActivity) *spotify.RecommendationsResponse {

	if data == nil {
		return nil
	}
	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks {
		artistsName := make([]string, len(item.Artists))

		for idx, artist := range item.Artists {
			artistsName[idx] = artist.Name
		}

		imageUrls := make([]string, len(item.Album.Images))

		for idx, image := range item.Album.Images {
			imageUrls[idx] = image.URL

		}

		items = append(items, spotify.SpotifyTrackObject{
			AlbumType:        item.Album.AlbumType,
			AlbumTotalTracks: item.Album.TotalTracks,
			AlbumImagesURL:   imageUrls,
			AlbumName:        item.Album.Name,

			ArtistsName: artistsName,
			Explicit:    item.Explicit,
			ID:          item.ID,
			Name:        item.Name,
			IsLiked:     mapTrackActivities[item.ID].IsLiked,
		})
	}

	return &spotify.RecommendationsResponse{

		Items: items,
	}

}
