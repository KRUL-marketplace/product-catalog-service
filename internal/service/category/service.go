package category

import (
	"context"
	"product-catalog-service/client/db"
	repository "product-catalog-service/internal/repository/category"
	"product-catalog-service/internal/repository/category/model"
)

type categoryService struct {
	categoryRepository repository.Repository
	txManager          db.TxManager
}

type CategoryService interface {
	Create(ctx context.Context, info *model.CategoryInfo) (uint32, error)
	Update(ctx context.Context, id uint32, info *model.CategoryInfo) error
	GetAll(ctx context.Context) ([]*model.Category, error)
	GetById(ctx context.Context, id uint32) (*model.Category, error)
	GetBySlug(ctx context.Context, slug string) (*model.Category, error)
	DeleteById(ctx context.Context, id uint32) error
}

func NewService(productRepository repository.Repository, txManager db.TxManager) CategoryService {
	return &categoryService{
		categoryRepository: productRepository,
		txManager:          txManager,
	}
}
