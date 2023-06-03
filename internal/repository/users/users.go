package users

import (
	"context"
	"errors"

	"social-network-api/internal/db/models"

	"github.com/jackc/pgx/v5"
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

func NewRepo(db *pgxpool.Pool) *Repo {
	return &Repo{DB: db}
}

func (r Repo) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users (email, password_hash, first_name, last_name, username, birthdate, activated)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at`

	args := []any{
		user.Email,
		user.Password.Hash,
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

func (r Repo) IsEmailUnique(ctx context.Context, email string) (bool, error) {
	var user models.User

	query := `
	SELECT id, email
	FROM users
	WHERE email = $1`

	err := r.DB.
		QueryRow(ctx, query, email).
		Scan(&user.Id, &user.Email)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (r Repo) IsUsernameUnique(ctx context.Context, username string) (bool, error) {
	var user models.User

	query := `
	SELECT id, username
	FROM users
	WHERE username = $1`

	err := r.DB.
		QueryRow(ctx, query, username).
		Scan(&user.Id, &user.Username)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}

func (r Repo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	query := `
	SELECT id, email, password_hash, first_name, last_name, username, birthdate, activated, created_at
	FROM users
	WHERE email = $1`

	err := r.DB.
		QueryRow(ctx, query, email).
		Scan(
			&user.Id,
			&user.Email,
			&user.Password.Hash,
			&user.Firstname,
			&user.Lastname,
			&user.Username,
			&user.Birthdate,
			&user.Activated,
			&user.Created_at,
		)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
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
	SELECT id, email, password_hash, first_name, last_name, username, birthdate, activated, created_at
	FROM users
	WHERE username = $1`

	err := r.DB.
		QueryRow(ctx, query, username).
		Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Firstname,
			&user.Lastname,
			&user.Username,
			&user.Birthdate,
			&user.Activated,
			&user.Created_at,
		)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
