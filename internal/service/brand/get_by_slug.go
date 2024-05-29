package brand

import (
	"context"
	"product-catalog-service/internal/repository/brand/model"
)

func (s *brandService) GetBySlug(ctx context.Context, slug string) (*model.Brand, error) {
	brand, err := s.brandRepository.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return brand, nil
}
