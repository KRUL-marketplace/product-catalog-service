package category

import (
	"context"
	"product-catalog-service/internal/repository/category/model"
)

func (s *categoryService) GetById(ctx context.Context, id uint32) (*model.Category, error) {
	category, err := s.categoryRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}
