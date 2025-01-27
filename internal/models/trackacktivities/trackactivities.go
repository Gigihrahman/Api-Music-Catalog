package trackacktivities

import "gorm.io/gorm"

type (
	TrackActivity struct {
		gorm.Model
		UserID    uint   `gorm:"not null"`
		SpotifyID string `gorm:"not null"`
		IsLiked   *bool
		CreatedBy string `gorm:"not null"`
		UpdatedBy string `gorm:"not null"`
	}
)

type (
	TrackActivityRequest struct {
		SpotifyID string `json:"spotifyID"`
		IsLiked   *bool  `json:"isLiked"` //trues = liked, false = dislike, null= neutral

	}
)
