package main

import (
	configs "api-music/internal/configs"
	membershipHandler "api-music/internal/handler/memberships"
	"api-music/internal/models/memberships"
	membershipRepo "api-music/internal/repository/memberships"
	membershipSvc "api-music/internal/service/memberships"
	"api-music/pkg/internalsql"
	"fmt"
	"log"

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
	
	if err!= nil{
		log.Fatal("gagal", err)
	}
	cfg = configs.Get()
	fmt.Println(cfg)
	db, err := internalsql.Connect(cfg.Database.DatabaseSourceName)
	if err !=nil{
		log.Fatalf("failed to connect to db: %+v",err)
	}
	db.AutoMigrate(&memberships.User{})
	
	membershipRepo:= membershipRepo.NewRepositoy(db)
	membershipsSvc := membershipSvc.NewService(cfg,membershipRepo)
	membershipHandler:= membershipHandler.NewHandler(r,membershipsSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}