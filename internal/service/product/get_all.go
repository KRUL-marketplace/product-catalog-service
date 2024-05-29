package product

import (
	"context"
	"product-catalog-service/internal/repository/product/model"
)

func (s *productService) GetAll(ctx context.Context) ([]*model.GetProduct, error) {
	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
