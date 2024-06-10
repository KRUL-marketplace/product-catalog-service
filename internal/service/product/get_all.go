package product

import (
	"context"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/product/model"
)

func (s *productService) GetAll(ctx context.Context, filter *model.FilterProduct) ([]*model.GetProduct, error) {
	products, err := s.productRepository.GetAll(ctx, filter)
	if err != nil {
		return nil, err
	}

	return products, nil
}
