package category

import (
	"context"
	"github.com/gosimple/slug"
	"product-catalog-service/internal/repository/category/model"
)

func (s *categoryService) Update(ctx context.Context, id uint32, info *model.CategoryInfo) error {
	info.Slug = slug.Make(info.Name)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.categoryRepository.Update(ctx, id, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return err
	}

	return err
}
