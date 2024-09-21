package blog

import (
	"context"
	models "github.com/meiti-x/go-blog/internal/_models"
)

type Repository interface {
	Create(ctx context.Context, news *models.Posts) (*models.Posts, error)
	//Update(ctx context.Context, news *models.News) (*models.News, error)
	//GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.NewsBase, error)
	//Delete(ctx context.Context, newsID uuid.UUID) error
	//GetNews(ctx context.Context, pq *utils.PaginationQuery) (*models.NewsList, error)
	//SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error)
}
