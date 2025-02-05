package trackactivities

import (
	"api-music/internal/models/trackacktivities"
	"context"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Test_repository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	now := time.Now()
	isLiked := true
	type args struct {
		model trackacktivities.TrackActivity
	}
	tests := []struct {
		name    string
		r       *repository
		args    args
		wantErr bool
		mockFn  func(args args)
	}{{
		name: "succes",
		args: args{
			model: trackacktivities.TrackActivity{
				Model: gorm.Model{
					CreatedAt: now,
					UpdatedAt: now,
				},
				UserID:    1,
				SpotifyID: "spotifyID",
				IsLiked:   &isLiked,
				CreatedBy: "1",
				UpdatedBy: "1",
			},
		},
		wantErr: false,
		mockFn: func(args args) {
			mock.ExpectBegin()

			mock.ExpectQuery(`INSERT INTO "track_activities" (.+) VALUES (.+)`).WithArgs(
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				args.model.UserID,
				args.model.SpotifyID,
				args.model.IsLiked,
				args.model.CreatedBy,
				args.model.UpdatedBy,
			).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uint(1)))

			mock.ExpectCommit()
		},
	},
		{
			name: "error",
			args: args{
				model: trackacktivities.TrackActivity{
					Model: gorm.Model{
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectQuery(`INSERT INTO "track_activities" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.SpotifyID,
					args.model.IsLiked,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			if err := r.Create(context.Background(), tt.args.model); (err != nil) != tt.wantErr {

				t.Errorf("repository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	now := time.Now()
	isLiked := true
	type args struct {
		model trackacktivities.TrackActivity
	}
	tests := []struct {
		name    string
		r       *repository
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "succes",
			args: args{
				model: trackacktivities.TrackActivity{
					Model: gorm.Model{
						ID:        123,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE  "track_activities" SET (.+) WHERE (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.SpotifyID,
					args.model.IsLiked,
					args.model.CreatedBy,
					args.model.UpdatedBy,
					args.model.ID,
				).WillReturnResult(sqlmock.NewResult(123, 1))

				mock.ExpectCommit()
			},
		},
		{
			name: "error",
			args: args{
				model: trackacktivities.TrackActivity{
					Model: gorm.Model{
						ID:        123,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
					CreatedBy: "1",
					UpdatedBy: "1",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()

				mock.ExpectExec(`UPDATE  "track_activities" SET (.+) WHERE (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.UserID,
					args.model.SpotifyID,
					args.model.IsLiked,
					args.model.CreatedBy,
					args.model.UpdatedBy,
					args.model.ID,
				).WillReturnError(assert.AnError)

				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := repository{
				db: gormDB,
			}
			if err := r.Update(context.Background(), tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("repository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_Get(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	now := time.Now()
	isLiked := true
	type args struct {
		userID    uint
		spotifyID string
	}
	tests := []struct {
		name string

		args    args
		want    *trackacktivities.TrackActivity
		wantErr bool
		mockFn  func(args args)
	}{{
		name: "succes",
		args: args{
			userID:    123,
			spotifyID: "spotifyID",
		},
		want: &trackacktivities.TrackActivity{
			Model: gorm.Model{
				ID:        1,
				CreatedAt: now,
				UpdatedAt: now,
			},
			UserID:    1,
			SpotifyID: "spotifyID",
			IsLiked:   &isLiked,
			CreatedBy: "test@gmail.com",
			UpdatedBy: "test@gmail.com",
		},
		wantErr: false,
		mockFn: func(args args) {
			mock.ExpectQuery(`SELECT \* FROM "track_activities" .+`).WithArgs(args.userID, args.spotifyID, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "spotify_id", "is_liked", "created_by", "updated_by"}).AddRow(1, now, now, 1, "spotifyID", true, "test@gmail.com", "test@gmail.com"))
		},
	}, {
		name: "Error",
		args: args{
			userID:    123,
			spotifyID: "spotifyID",
		},
		want:    nil,
		wantErr: true,
		mockFn: func(args args) {
			mock.ExpectQuery(`SELECT \* FROM "track_activities" .+`).WithArgs(args.userID, args.spotifyID, 1).WillReturnError(assert.AnError)
		},
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := repository{
				db: gormDB,
			}
			got, err := r.Get(context.Background(), tt.args.userID, tt.args.spotifyID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.Get() = %v, want %v", got, tt.want)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func Test_repository_GetBulkSpotifyIDs(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)
	now := time.Now()
	isLiked := true
	type args struct {
		userID     uint
		spotifyIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]trackacktivities.TrackActivity
		wantErr bool
		mockFn  func(args args)
	}{
		{
			name: "succes",
			args: args{
				userID:     123,
				spotifyIDs: []string{"spotifyID"},
			},
			want: map[string]trackacktivities.TrackActivity{

				"spotifyID": {
					Model: gorm.Model{
						ID:        1,
						CreatedAt: now,
						UpdatedAt: now,
					},
					UserID:    1,
					SpotifyID: "spotifyID",
					IsLiked:   &isLiked,
					CreatedBy: "test@gmail.com",
					UpdatedBy: "test@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "track_activities" .+`).WithArgs(args.userID, strings.Join(args.spotifyIDs, ",")).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "user_id", "spotify_id", "is_liked", "created_by", "updated_by"}).AddRow(1, now, now, 1, "spotifyID", true, "test@gmail.com", "test@gmail.com"))
			},
		},
		{
			name: "failed",
			args: args{
				userID:     123,
				spotifyIDs: []string{"spotifyID"},
			},
			want:    nil,
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectQuery(`SELECT \* FROM "track_activities" .+`).WithArgs(args.userID, strings.Join(args.spotifyIDs, ",")).WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &repository{
				db: gormDB,
			}
			got, err := r.GetBulkSpotifyIDs(context.Background(), tt.args.userID, tt.args.spotifyIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetBulkSpotifyIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetBulkSpotifyIDs() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
