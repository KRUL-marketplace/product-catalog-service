package product

import (
	"context"
	"product-catalog-service/client/db"
	repository "product-catalog-service/internal/repository/product"
	"product-catalog-service/internal/repository/product/model"
)

type productService struct {
	productRepository repository.Repository
	txManager         db.TxManager
}

type ProductService interface {
	Create(ctx context.Context, info *model.CreateProduct) (string, error)
	GetById(ctx context.Context, id string) (*model.GetProduct, error)
	GetBySlug(ctx context.Context, slug string) (*model.GetProduct, error)
	GetAll(ctx context.Context) ([]*model.GetProduct, error)
	GetByBrand(ctx context.Context, brandId uint32) ([]*model.GetProduct, error)
	Update(ctx context.Context, id string, info *model.CreateProduct) (string, error)
	DeleteById(ctx context.Context, id string) error
}

func NewService(productRepository repository.Repository, txManager db.TxManager) ProductService {
	return &productService{
		productRepository: productRepository,
		txManager:         txManager,
	}
}
