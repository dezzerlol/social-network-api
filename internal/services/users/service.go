package users

import (
	"context"
	"social-network-api/internal/db/models"
	"social-network-api/internal/repository/users"
	"social-network-api/pkg/password"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}

type service struct {
	DB       *pgxpool.Pool
	userRepo *users.Repo
}

func New(db *pgxpool.Pool) Service {
	return &service{DB: db}
}

func (s *service) Create(ctx context.Context, user *models.User) error {
	// check if user with this email already exists
	user, err := s.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}

	if user != nil {
		return users.ErrDuplicateEmail
	}

	// check if user with this username already exists
	user, err = s.FindByUsername(ctx, user.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return users.ErrDuplicateUsername
	}

	// hash password
	hashedPass, err := password.Generate(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPass
	user.Activated = false

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)

	if err != nil {
		return nil, err
	}

	return user, nil
}
