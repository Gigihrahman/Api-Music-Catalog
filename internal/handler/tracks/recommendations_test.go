package tracks

import (
	"api-music/internal/models/spotify"
	"api-music/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestHandler_GetRecommendation(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()
	MockSvc := NewMockservice(mockCtrl)

	tests := []struct {
		name               string
		expectedStatusCode int
		expectedBody       *spotify.RecommendationsResponse
		wantErr            bool
		mockFn             func()
	}{
		{
			name:               "success",
			expectedStatusCode: http.StatusOK,
			expectedBody: &spotify.RecommendationsResponse{

				Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
					},
					{AlbumType: "album",
						AlbumTotalTracks: 12,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0"},
						AlbumName:        "A Night At The Opera (2011 Remaster)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "4u7EnebtmKWzUH433cf5Qv",
						Name:             "Bohemian Rhapsody - Remastered 2011"},
				},
			},
			wantErr: false,
			mockFn: func() {
				MockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(&spotify.RecommendationsResponse{Items: []spotify.SpotifyTrackObject{
					{
						AlbumType:        "album",
						AlbumTotalTracks: 22,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00001e02e8b066f70c206551210d902b", "https://i.scdn.co/image/ab67616d00004851e8b066f70c206551210d902b"},
						AlbumName:        "Bohemian Rhapsody (The Original Soundtrack)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "3z8h0TU7ReDPLIbEnYhWZb",
						Name:             "Bohemian Rhapsody",
					},
					{AlbumType: "album",
						AlbumTotalTracks: 12,
						AlbumImagesURL:   []string{"https://i.scdn.co/image/ab67616d0000b273e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00001e02e319baafd16e84f0408af2a0", "https://i.scdn.co/image/ab67616d00004851e319baafd16e84f0408af2a0"},
						AlbumName:        "A Night At The Opera (2011 Remaster)",
						ArtistsName:      []string{"Queen"},
						Explicit:         false,
						ID:               "4u7EnebtmKWzUH433cf5Qv",
						Name:             "Bohemian Rhapsody - Remastered 2011"},
				}}, nil)
			},
		},

		{
			name:               "failed",
			expectedStatusCode: http.StatusBadRequest,
			expectedBody:       nil,
			wantErr:            true,
			mockFn: func() {
				MockSvc.EXPECT().GetRecommendation(gomock.Any(), uint(1), 10, "trackID").Return(nil, assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			api := gin.New()
			h := &Handler{
				Engine:  api,
				service: MockSvc,
			}
			h.RegisterRoute()
			w := httptest.NewRecorder()

			endpoint := `/tracks/recommendations?limit=10&trackID=trackID`

			req, err := http.NewRequest(http.MethodGet, endpoint, nil)
			assert.NoError(t, err)
			token, err := jwt.CreateToken(1, "username", "")
			assert.NoError(t, err)
			req.Header.Set("Authorization", token)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)

			if !tt.wantErr {
				res := w.Result()
				defer res.Body.Close()
				response := spotify.RecommendationsResponse{}
				err = json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				fmt.Println("succes nih bang")
				assert.Equal(t, tt.expectedBody, &response)
			}
		})
	}
}
