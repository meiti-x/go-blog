package blog

import (
	"context"
	models "github.com/meiti-x/go-blog/internal/_models"
)

type UseCase interface {
	Create(ctx context.Context, blog *models.Posts) (*models.Posts, error)
	//Update(ctx context.Context, blog *models.Posts) (*models.Posts, error)
	//GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.PostBase, error)
	//Delete(ctx context.Context, newsID uuid.UUID) error
	//GetPosts(ctx context.Context) (*models.PostList, error)
	//SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error)
}
