package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	models "github.com/meiti-x/go-blog/internal/_models"
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

func (r *BlogRepo) Create(ctx context.Context, posts *models.Posts) (*models.Posts, error) {

	return nil, nil
}
