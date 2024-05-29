package brand

import (
	"context"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/brand/model"
)

const (
	tableName         = "brands"
	idColumn          = "id"
	nameColumn        = "name"
	slugColumn        = "slug"
	descriptionColumn = "description"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

type Repository interface {
	Create(ctx context.Context, info *model.BrandInfo) (uint32, error)
	GetAll(ctx context.Context) ([]*model.Brand, error)
	DeleteById(ctx context.Context, id uint32) error
	Update(ctx context.Context, id uint32, info *model.BrandInfo) error
	GetById(ctx context.Context, id uint32) (*model.Brand, error)
	GetBySlug(ctx context.Context, slug string) (*model.Brand, error)
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repository {
	return &repo{
		db: db,
	}
}
