package product

import (
	"context"
	"product-catalog-service/internal/repository/product/model"
)

func (s *productService) GetBySlug(ctx context.Context, slug string) (*model.GetProduct, error) {
	product, err := s.productRepository.GetByField(ctx, "slug", slug)
	if err != nil {
		return nil, err
	}

	return product, nil
}
