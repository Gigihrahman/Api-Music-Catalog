package tracks

import (
	"api-music/internal/models/trackacktivities"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_UpsertTrackActivities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTrackActivityRepo := NewMocktrackActivitiesRepository(mockCtrl)
	isLikedTrue := true
	isLikedfalse := false
	type args struct {
		userID  uint
		request trackacktivities.TrackActivityRequest
	}
	tests := []struct {
		name    string
		s       *service
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "succes: create",
			args: args{userID: 1, request: trackacktivities.TrackActivityRequest{
				SpotifyID: "spotifyID",
				IsLiked:   &isLikedTrue,
			}},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).Return(nil, gorm.ErrRecordNotFound)
				mockTrackActivityRepo.EXPECT().Create(gomock.Any(), trackacktivities.TrackActivity{
					UserID:    args.userID,
					SpotifyID: args.request.SpotifyID,
					IsLiked:   args.request.IsLiked,
					CreatedBy: fmt.Sprintf("%d", args.userID),
					UpdatedBy: fmt.Sprintf("%d", args.userID),
				}).Return(nil)

			},
		}, {
			name: "succes: update",
			args: args{userID: 1, request: trackacktivities.TrackActivityRequest{
				SpotifyID: "spotifyID",
				IsLiked:   &isLikedTrue,
			}},
			wantErr: false,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).Return(&trackacktivities.TrackActivity{
					IsLiked: &isLikedfalse,
				}, nil)
				mockTrackActivityRepo.EXPECT().Update(gomock.Any(), trackacktivities.TrackActivity{
					IsLiked: args.request.IsLiked}).Return(nil)

			},
		}, {
			name: "failed",
			args: args{userID: 1, request: trackacktivities.TrackActivityRequest{
				SpotifyID: "spotifyID",
				IsLiked:   &isLikedTrue,
			}},
			wantErr: true,
			mockFn: func(args args) {
				mockTrackActivityRepo.EXPECT().Get(gomock.Any(), args.userID, args.request.SpotifyID).Return(
					nil, assert.AnError)

			},
		},
	}
	for _, tt := range tests {
		tt.mockFn(tt.args)
		s := &service{

			trackActivitiesRepo: mockTrackActivityRepo,
		}
		t.Run(tt.name, func(t *testing.T) {
			if err := s.UpsertTrackActivities(context.Background(), tt.args.userID, tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.UpsertTrackActivities() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
