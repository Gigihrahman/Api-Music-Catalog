package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rs/zerolog/log"
)

func (o *outbound) GetRecommendation(ctx context.Context, limit int, trackID string) (*SpotifyRecommendationResponse, error) {
	params := url.Values{}

	params.Set("type", "track")
	params.Set("limit", strconv.Itoa(limit))
	params.Set("market", "ID")
	params.Set("seed_tracks", trackID)
	basePath := `https://api.spotify.com/v1/recomendations`
	urlPath := fmt.Sprintf("%s?%s", basePath, params.Encode())

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Error().Err(err).Msg("error create recomendations request for spotify")
		return nil, err
	}
	accesToken, tokenType, err := o.GetTokenDetails()
	if err != nil {
		log.Error().Err(err).Msg("error get token detail for recomendations")
		return nil, err
	}

	bearerToken := fmt.Sprintf("%s %s", tokenType, accesToken)
	req.Header.Set("Authorization", bearerToken)

	resp, err := o.client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("error execute recomendations request for spotify")
		return nil, err
	}
	defer resp.Body.Close()

	var response SpotifyRecommendationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("ini responsenya: ", resp)
		log.Error().Err(err).Msg("error unmarshal recomendations response from spotify")
		return nil, err
	}
	return &response, nil
}
