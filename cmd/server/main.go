package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/nishantshekhada/ezygaming/internal/config"
	"github.com/nishantshekhada/ezygaming/internal/database"
	"github.com/nishantshekhada/ezygaming/internal/routes"
)

func main() {
	cfg := config.Load()

	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := database.Connect(cfg)
	database.Migrate(db)
	database.Seed(db)

	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")

	// Register all routes
	routes.Setup(r, db, cfg.JWT.Secret)

	addr := ":" + cfg.AppPort
	log.Printf("Ezy Gaming server starting on http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
