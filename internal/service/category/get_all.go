package category

import (
	"context"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/category/model"
)

func (s *categoryService) GetAll(ctx context.Context) ([]*model.Category, error) {
	categories, err := s.categoryRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
