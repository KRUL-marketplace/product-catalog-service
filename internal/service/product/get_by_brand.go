package product

import (
	"context"
	"product-catalog-service/internal/repository/product/model"
)

func (s *productService) GetByBrand(ctx context.Context, brandId uint32) ([]*model.GetProduct, error) {
	products, err := s.productRepository.GetByBrand(ctx, brandId)
	if err != nil {
		return nil, err
	}

	return products, nil
}
