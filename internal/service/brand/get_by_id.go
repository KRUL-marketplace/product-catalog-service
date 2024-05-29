package brand

import (
	"context"
	"product-catalog-service/internal/repository/brand/model"
)

func (s *brandService) GetById(ctx context.Context, id uint32) (*model.Brand, error) {
	brand, err := s.brandRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return brand, nil
}
