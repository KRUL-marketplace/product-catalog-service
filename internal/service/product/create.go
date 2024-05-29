package product

import (
	"context"
	"github.com/gosimple/slug"
	"product-catalog-service/internal/repository/product/model"
)

func (s *productService) Create(ctx context.Context, info *model.CreateProduct) (string, error) {
	info.Slug = slug.Make(info.Name)

	var id string
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.productRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return id, nil
}
