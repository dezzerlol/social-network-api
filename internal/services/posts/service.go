package posts

import (
	"context"
	"mime/multipart"
	"social-network-api/cfg"
	"social-network-api/internal/db/models"
	"social-network-api/internal/repository/media"
	"social-network-api/internal/repository/posts"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	CreatePost(files []*multipart.FileHeader, body string) error
}

type service struct {
	postsRepo *posts.Repo
	mediaRepo *media.Repo
}

func New(db *pgxpool.Pool) Service {
	return &service{
		postsRepo: posts.NewRepo(db),
		mediaRepo: media.NewRepo(cfg.Get().Cloud.Name, cfg.Get().Cloud.Key, cfg.Get().Cloud.Secret),
	}
}

func (s *service) CreatePost(files []*multipart.FileHeader, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Upload files to cloud
	media := make([]models.Media, len(files))
	for i, file := range files {
		file, _ := file.Open()
		defer file.Close()

		res, err := s.mediaRepo.Upload(ctx, file, "posts")
		if err != nil {
			return err
		}

		media[i].Url = res.PublicLink
	}

	// Create post
	post := &models.Post{
		Body:   body,
		UserId: 5,
		Images: media,
	}

	err := s.postsRepo.CreatePost(ctx, post)

	if err != nil {
		return err
	}

	return nil
}
