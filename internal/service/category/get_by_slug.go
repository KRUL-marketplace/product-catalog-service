package category

import (
	"context"
	"product-catalog-service/internal/repository/category/model"
)

func (s *categoryService) GetBySlug(ctx context.Context, slug string) (*model.Category, error) {
	category, err := s.categoryRepository.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	return category, nil
}
