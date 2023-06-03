package posts

import (
	"context"
	"errors"
	"social-network-api/internal/db/models"
	"social-network-api/pkg/dbutil"

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

func (r *Repo) CreatePost(ctx context.Context, post *models.Post) error {
	query := `
	INSERT INTO posts (user_id, body)
	VALUES ($1, $2)
	RETURNING id, created_at
	`

	postArgs := []any{post.UserId, post.Body}

	err := r.DB.
		QueryRow(ctx, query, postArgs...).
		Scan(&post.Id, &post.Created_at)

	if err != nil {
		return err
	}

	if len(post.Images) > 0 {
		query = `
		INSERT INTO post_images (post_id, url)
		VALUES %s
		`

		argsPerRow := 2
		imageArgs := make([]interface{}, 0, argsPerRow*len(post.Images))

		for _, image := range post.Images {
			imageArgs = append(imageArgs, post.Id, image.Url)
		}

		batchSQLString := dbutil.GetBulkInsertSQLString(query, argsPerRow, len(post.Images))

		_, err = r.DB.Exec(ctx, batchSQLString, imageArgs...)
	}

	return err
}

func (r *Repo) DeletePost(ctx context.Context, userId, postId int) error {
	query := `
	DELETE FROM posts
	WHERE id = $1 AND user_id = $2
	`

	args := []any{postId, userId}

	_, err := r.DB.Exec(ctx, query, args...)

	return err
}
