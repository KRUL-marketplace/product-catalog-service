package product

import (
	"context"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/product/model"
)

const (
	tableName = "products"

	idColumn          = "id"
	nameColumn        = "name"
	slugColumn        = "slug"
	descriptionColumn = "description"
	priceColumn       = "price"
	brandIdColumn     = "brand_id"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

type Repository interface {
	Create(ctx context.Context, info *model.CreateProduct) (string, error)
	GetByField(ctx context.Context, field string, value string) (*model.GetProduct, error)
	GetAll(ctx context.Context) ([]*model.GetProduct, error)
	GetByBrand(ctx context.Context, brandId uint32) ([]*model.GetProduct, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, id string, info *model.CreateProduct) (string, error)
}

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) Repository {
	return &repo{
		db: db,
	}
}
