package main

import (
	"time"

	"social-network-api/config"
	"social-network-api/internal/http"
	"social-network-api/internal/repository/pg"
	"social-network-api/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.New()
	cfg, err := config.Load(".")

	if err != nil {
		logger.Fatalf("Error reading config: %s", err)
	}

	db, err := pg.New(cfg)

	if err != nil {
		logger.Fatalf("Error starting db: %s", err)
	}

	defer db.Close()

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		time.Sleep(10 * time.Second)
		ctx.JSON(200, gin.H{
			"health": "ok",
		})
	})

	httpServer := http.New(r, logger)
	httpServer.Run()
}
