package tracks

import (
	"api-music/internal/models/spotify"
	"api-music/internal/models/trackacktivities"
	spotifyRepo "api-music/internal/repository/spotify"
	"context"

	"github.com/rs/zerolog/log"
)

func (s *service) Search(ctx context.Context, query string, pageSize, pageIndex int, userID uint) (*spotify.SearchResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	trackDetails, err := s.spotifyOutbond.Search(ctx, query, limit, offset)
	if err != nil {
		log.Error().Err(err).Msg("error search track to spotify")
		return nil, err
	}
	trackIDs := make([]string, len(trackDetails.Tracks.Items))

	for idx, item := range trackDetails.Tracks.Items {
		trackIDs[idx] = item.ID
	}

	trackActivities, err := s.trackActivitiesRepo.GetBulkSpotifyIDs(ctx, userID, trackIDs)
	if err != nil {
		log.Error().Err(err).Msg("error get track activities from db")
		return nil, err
	}
	return modelToresponse(trackDetails, trackActivities), nil
}

func modelToresponse(data *spotifyRepo.SpotifySearchResponse, mapTrackActivities map[string]trackacktivities.TrackActivity) *spotify.SearchResponse {

	if data == nil {
		return nil
	}
	items := make([]spotify.SpotifyTrackObject, 0)

	for _, item := range data.Tracks.Items {
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

	return &spotify.SearchResponse{
		Limit:  data.Tracks.Limit,
		Offset: data.Tracks.Offset,
		Total:  data.Tracks.Total,
		Items:  items,
	}

}
