package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	models "github.com/meiti-x/go-blog/internal/_models"
	"github.com/pkg/errors"
)

// BlogRepo Repository
type BlogRepo struct {
	db *sqlx.DB
}

func NewBlogRepo(db *sqlx.DB) *BlogRepo {
	return &BlogRepo{
		db: db,
	}
}

func (r *BlogRepo) Create(ctx context.Context, post *models.Posts) (*models.Posts, error) {
	query := `INSERT INTO blogs (title, content) VALUES ($1, $2)`

	var n models.Posts
	if err := r.db.QueryRowxContext(ctx, query, &post.Title, &post.Content).StructScan(&n); err != nil {
		return nil, errors.Wrap(err, "blogRepo.Create.QueryRowxContext")
	}

	return post, nil
}
