package posts

import (
	"context"
	"errors"
	"fmt"
	"social-network-api/internal/db/models"
	"strconv"
	"strings"

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

		batchSQLString := getBulkInsertSQLSimple(query, argsPerRow, len(post.Images))

		_, err = r.DB.Exec(ctx, batchSQLString, imageArgs...)
	}

	return err
}

func getBulkInsertSQL(SQLString string, rowValueSQL string, numRows int) string {
	// Combine the base SQL string and N value strings
	valueStrings := make([]string, 0, numRows)
	for i := 0; i < numRows; i++ {
		valueStrings = append(valueStrings, "("+rowValueSQL+")")
	}
	allValuesString := strings.Join(valueStrings, ",")
	SQLString = fmt.Sprintf(SQLString, allValuesString)

	// Convert all of the "?" to "$1", "$2", "$3", etc.
	// (which is the way that pgx expects query variables to be)
	numArgs := strings.Count(SQLString, "?")
	SQLString = strings.ReplaceAll(SQLString, "?", "$%v")
	numbers := make([]interface{}, 0, numRows)
	for i := 1; i <= numArgs; i++ {
		numbers = append(numbers, strconv.Itoa(i))
	}
	return fmt.Sprintf(SQLString, numbers...)
}

// getBulkInsertSQLSimple is a helper function to prepare a SQL query for a bulk insert.
// getBulkInsertSQLSimple is used over getBulkInsertSQL when all of the values are plain question
// marks (e.g. a 1-for-1 value insertion).
// The example given for getBulkInsertSQL is such a query.
func getBulkInsertSQLSimple(SQLString string, numArgsPerRow int, numRows int) string {
	questionMarks := make([]string, 0, numArgsPerRow)
	for i := 0; i < numArgsPerRow; i++ {
		questionMarks = append(questionMarks, "?")
	}
	rowValueSQL := strings.Join(questionMarks, ", ")
	return getBulkInsertSQL(SQLString, rowValueSQL, numRows)
}
