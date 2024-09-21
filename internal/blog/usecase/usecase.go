package usecase

import (
	"context"
	"github.com/meiti-x/go-blog/config"
	models "github.com/meiti-x/go-blog/internal/_models"
	"github.com/meiti-x/go-blog/internal/blog"
)

// News UseCase
type newsUC struct {
	cfg      *config.Config
	blogRepo blog.Repository
}

// News UseCase constructor
func NewBlogUseCase(cfg *config.Config, newsRepo blog.Repository) blog.UseCase {
	return &newsUC{cfg: cfg, blogRepo: newsRepo}
}

func (uc *newsUC) Create(ctx context.Context, blog *models.Posts) (*models.Posts, error) {

}
