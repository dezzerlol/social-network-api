package db

import (
	"context"
	"social-network-api/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New(cfg *config.Config) (*pgxpool.Pool, error) {
	dbConf, err := pgxpool.ParseConfig(cfg.DB_DSN)

	if err != nil {
		return nil, err
	}

	dbConf.MaxConns = cfg.DB_CFG.MaxOpenConn
	dbConf.MinConns = cfg.DB_CFG.MaxIdleConn
	dbConf.MaxConnIdleTime = cfg.DB_CFG.ConnMaxLifeTime * time.Minute

	db, err := pgxpool.New(context.Background(), cfg.DB_DSN)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Ping(ctx)

	if err != nil {
		return nil, err
	}

	return db, nil
}
