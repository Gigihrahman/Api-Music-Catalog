package spotify

type SearchResponse struct {
	Offset int                  `json:"offset"`
	Limit  int                  `json:"limit"`
	Total  int                  `json:"total"`
	Items  []SpotifyTrackObject `json:"items"`
}

type SpotifyTrackObject struct {
	//album related field
	AlbumType        string   `json:"albumType"`
	AlbumTotalTracks int      `json:"totalTracks"`
	AlbumImagesURL   []string `json:"albumImagesURL"`
	AlbumName        string   `json:"albumName"`
	//artist related field
	ArtistsName []string `json:"artistsName"`
	Explicit    bool     `json:"explicit"`
	//track related field
	ID      string `json:"id"`
	Name    string `json:"name"`
	IsLiked *bool  `json:"isLiked"`
}

type RecommendationsResponse struct {
	Items []SpotifyTrackObject `json:"items"`
}
