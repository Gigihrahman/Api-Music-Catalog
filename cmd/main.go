package main

import (
	configs "api-music/internal/configs"
	membershipHandler "api-music/internal/handler/memberships"
	tracksHandler "api-music/internal/handler/tracks"
	"api-music/internal/models/memberships"
	"api-music/internal/models/trackacktivities"
	membershipRepo "api-music/internal/repository/memberships"
	"api-music/internal/repository/spotify"
	membershipSvc "api-music/internal/service/memberships"
	"api-music/internal/service/tracks"
	"api-music/pkg/httpclient"
	"api-music/pkg/internalsql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("gagal", err)
	}
	cfg = configs.Get()
	fmt.Println(cfg)
	db, err := internalsql.Connect(cfg.Database.DatabaseSourceName)
	if err != nil {
		log.Fatalf("failed to connect to db: %+v", err)
	}
	db.AutoMigrate(&memberships.User{})
	db.AutoMigrate(&trackacktivities.TrackActivity{})

	httpClient := httpclient.NewClient(&http.Client{})

	spotifyOutbond := spotify.NewSpotyOutbound(cfg, httpClient)
	trackSvc := tracks.NewService(spotifyOutbond)

	membershipRepo := membershipRepo.NewRepositoy(db)
	membershipsSvc := membershipSvc.NewService(cfg, membershipRepo)
	membershipHandler := membershipHandler.NewHandler(r, membershipsSvc)
	membershipHandler.RegisterRoute()
	tracksHandler := tracksHandler.NewHandler(r, trackSvc)
	tracksHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
