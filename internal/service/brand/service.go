package brand

import (
	"context"
	"github.com/KRUL-marketplace/common-libs/pkg/client/db"
	repository "github.com/KRUL-marketplace/product-catalog-service/internal/repository/brand"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/brand/model"
)

type brandService struct {
	brandRepository repository.Repository
	txManager       db.TxManager
}

type BrandService interface {
	Create(ctx context.Context, info *model.BrandInfo) (uint32, error)
	GetAll(ctx context.Context) ([]*model.Brand, error)
	DeleteById(ctx context.Context, id uint32) error
	GetById(ctx context.Context, id uint32) (*model.Brand, error)
	Update(ctx context.Context, id uint32, info *model.BrandInfo) error
	GetBySlug(ctx context.Context, slug string) (*model.Brand, error)
}

func NewService(brandRepository repository.Repository, txManager db.TxManager) BrandService {
	return &brandService{
		brandRepository: brandRepository,
		txManager:       txManager,
	}
}
