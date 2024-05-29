package product

import (
	"context"
	"product-catalog-service/internal/repository/product/model"
)

func (s *productService) GetById(ctx context.Context, id string) (*model.GetProduct, error) {
	product, err := s.productRepository.GetByField(ctx, "id", id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
