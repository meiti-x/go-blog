package usecase

import (
	"context"
	"github.com/meiti-x/go-blog/config"
	models "github.com/meiti-x/go-blog/internal/_models"
	"github.com/meiti-x/go-blog/internal/blog"
)

// News UseCase
type BlogUC struct {
	cfg      *config.Config
	blogRepo blog.Repository
}

// News UseCase constructor
func NewBlogUseCase(cfg *config.Config, blogRepo blog.Repository) *BlogUC {
	return &BlogUC{cfg: cfg, blogRepo: blogRepo}
}

func (uc *BlogUC) Create(ctx context.Context, blog *models.Posts) (*models.Posts, error) {
	return uc.blogRepo.Create(ctx, blog)
}
