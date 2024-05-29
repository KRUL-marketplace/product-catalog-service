package brand

import (
	"context"
	"product-catalog-service/internal/repository/brand/model"
)

func (s *brandService) GetAll(ctx context.Context) ([]*model.Brand, error) {
	brands, err := s.brandRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return brands, nil
}
