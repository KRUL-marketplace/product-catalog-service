package category

import (
	"context"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/category/model"
)

const (
	tableName       = "categories"
	idColumn        = "id"
	nameColumn      = "name"
	slugColumn      = "slug"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type Repository interface {
	Create(ctx context.Context, info *model.CategoryInfo) (uint32, error)
	Update(ctx context.Context, id uint32, info *model.CategoryInfo) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetById(ctx context.Context, id uint32) (*model.Category, error)
	GetBySlug(ctx context.Context, slug string) (*model.Category, error)
	DeleteById(ctx context.Context, id uint32) error
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repository {
	return &repo{
		db: db,
	}
}
