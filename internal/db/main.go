package db

import (
	"context"
	"social-network-api/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func New() (*pgxpool.Pool, error) {
	dbConf, err := pgxpool.ParseConfig(config.Get().DB_DSN)

	if err != nil {
		return nil, err
	}

	dbConf.MaxConns = config.Get().DB_CFG.MaxOpenConn
	dbConf.MinConns = config.Get().DB_CFG.MaxIdleConn
	dbConf.MaxConnIdleTime = config.Get().DB_CFG.ConnMaxLifeTime * time.Minute

	db, err := pgxpool.New(context.Background(), config.Get().DB_DSN)

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
