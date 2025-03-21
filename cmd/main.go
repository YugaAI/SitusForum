package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"main.go/internal/configs"
	"main.go/internal/handlers/memberships"
	"main.go/internal/handlers/posts"
	membershipRepo "main.go/internal/repository/memberships"
	postRepo "main.go/internal/repository/posts"
	membershipSvc "main.go/internal/service/memberships"
	postSvc "main.go/internal/service/posts"
	"main.go/pkg/internalsql"
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
		log.Fatalf("failed to init config: %v", err)
	}
	cfg = configs.Get()
	log.Println("config: ", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)
	

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	membershipsHandler.RegisterRoute()

	postsHandler := posts.NewHandler(r, postService)
	postsHandler.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
