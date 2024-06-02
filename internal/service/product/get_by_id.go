package product

import (
	"context"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/product/model"
)

func (s *productService) GetById(ctx context.Context, ids []string) (*[]model.GetProduct, error) {
	product, err := s.productRepository.GetById(ctx, ids)
	if err != nil {
		return nil, err
	}

	return product, nil
}
