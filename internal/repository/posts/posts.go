package posts

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	DB *pgxpool.Pool
}

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{DB: db}
}
