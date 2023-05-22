package users

import (
	"context"
	"database/sql"
	"errors"
	"social-network-api/internal/db/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repo struct {
	DB *pgxpool.Pool
}

var (
	ErrRecordNotFound    = errors.New("record not found")
	ErrEditConflict      = errors.New("edit conflict")
	ErrDuplicateEmail    = errors.New("duplicate email")
	ErrDuplicateUsername = errors.New("duplicate username")
)

func (r Repo) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users (email, password_hash, first_name, last_name, username, birthdate, activated)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at
	`

	args := []any{
		user.Email,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Username,
		user.Birthdate,
		user.Activated,
	}

	err := r.DB.
		QueryRow(ctx, query, args...).
		Scan(&user.Id, &user.Created_at)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		case err.Error() == `pq: duplicate key value violates unique constraint "users_username_key"`:
			return ErrDuplicateUsername
		default:
			return err
		}
	}

	return nil
}

func (r Repo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := `
	SELECT id, username, email, password_hash, activated, created_at
	FROM users
	WHERE email = $1`

	err := r.DB.
		QueryRow(ctx, query, email).
		Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Activated,
			&user.Created_at,
		)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (r Repo) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	query := `
	SELECT id, username, email, password_hash, activated, created_at
	FROM users
	WHERE username = $1`

	err := r.DB.
		QueryRow(ctx, query, username).
		Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Activated,
			&user.Created_at,
		)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}
