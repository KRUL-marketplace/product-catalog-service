package brand

import (
	"context"
	"github.com/gosimple/slug"
	"product-catalog-service/internal/repository/brand/model"
)

func (s *brandService) Create(ctx context.Context, info *model.BrandInfo) (uint32, error) {
	info.Slug = slug.Make(info.Name)

	var id uint32
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.brandRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
