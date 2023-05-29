package posts

import (
	"context"
	"mime/multipart"
	cfg "social-network-api/config"
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

	file_urls := make([]string, len(files))
	// Upload files to cloud
	for _, file := range files {
		file, _ := file.Open()

		res, err := s.mediaRepo.Upload(ctx, file, "posts")
		if err != nil {
			return err
		}

		file_urls = append(file_urls, res.PublicLink)
	}

	return nil
}
