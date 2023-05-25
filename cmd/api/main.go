package main

import (
	"social-network-api/config"
	"social-network-api/internal/db"
	"social-network-api/internal/http"
	"social-network-api/internal/redis"
	"social-network-api/pkg/logger"
)

func main() {
	logger := logger.New()
	cfg, err := config.Load(".")

	if err != nil {
		logger.Fatalf("Error reading config: %s", err)
	}

	db, err := db.New(cfg)

	if err != nil {
		logger.Fatalf("Error starting db: %s", err)
	}

	defer db.Close()

	cache := redis.New(cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Pass)

	defer cache.Close()

	httpServer := http.New(logger, db, cache)
	httpServer.Run()
}
