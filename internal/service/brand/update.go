package brand

import (
	"context"
	"github.com/KRUL-marketplace/product-catalog-service/internal/repository/brand/model"
	"github.com/gosimple/slug"
)

func (s *brandService) Update(ctx context.Context, id uint32, info *model.BrandInfo) error {
	info.Slug = slug.Make(info.Name)

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.brandRepository.Update(ctx, id, info)
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
